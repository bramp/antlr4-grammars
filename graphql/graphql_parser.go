// Code generated from GraphQL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package graphql // GraphQL
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 66, 680,
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
	71, 9, 71, 4, 72, 9, 72, 3, 2, 6, 2, 146, 10, 2, 13, 2, 14, 2, 147, 3,
	3, 3, 3, 3, 3, 5, 3, 153, 10, 3, 3, 4, 3, 4, 5, 4, 157, 10, 4, 3, 5, 3,
	5, 5, 5, 161, 10, 5, 3, 5, 5, 5, 164, 10, 5, 3, 5, 5, 5, 167, 10, 5, 3,
	5, 3, 5, 3, 5, 5, 5, 172, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 6, 7, 178, 10,
	7, 13, 7, 14, 7, 179, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 187, 10, 8, 3,
	9, 5, 9, 190, 10, 9, 3, 9, 3, 9, 5, 9, 194, 10, 9, 3, 9, 5, 9, 197, 10,
	9, 3, 9, 5, 9, 200, 10, 9, 3, 10, 3, 10, 6, 10, 204, 10, 10, 13, 10, 14,
	10, 205, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 5, 13, 220, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 5,
	14, 226, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 5, 17, 237, 10, 17, 3, 17, 5, 17, 240, 10, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 253,
	10, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 6, 25, 271, 10, 25, 13,
	25, 14, 25, 272, 3, 25, 3, 25, 5, 25, 277, 10, 25, 3, 26, 3, 26, 7, 26,
	281, 10, 26, 12, 26, 14, 26, 284, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 6, 29, 297, 10, 29, 13, 29,
	14, 29, 298, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 307, 10,
	30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 5, 32, 314, 10, 32, 3, 32, 3, 32,
	5, 32, 318, 10, 32, 5, 32, 320, 10, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 35, 6, 35, 329, 10, 35, 13, 35, 14, 35, 330, 3, 36, 3, 36,
	3, 36, 5, 36, 336, 10, 36, 3, 37, 3, 37, 3, 37, 5, 37, 341, 10, 37, 3,
	38, 3, 38, 5, 38, 345, 10, 38, 3, 39, 3, 39, 5, 39, 349, 10, 39, 3, 39,
	3, 39, 6, 39, 353, 10, 39, 13, 39, 14, 39, 354, 3, 39, 3, 39, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 5, 41, 366, 10, 41, 3, 41, 3, 41,
	6, 41, 370, 10, 41, 13, 41, 14, 41, 371, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	41, 5, 41, 379, 10, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 393, 10, 44, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 5, 45, 401, 10, 45, 3, 46, 5, 46, 404, 10, 46,
	3, 46, 3, 46, 3, 46, 5, 46, 409, 10, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3,
	47, 3, 48, 5, 48, 417, 10, 48, 3, 48, 3, 48, 3, 48, 5, 48, 422, 10, 48,
	3, 48, 5, 48, 425, 10, 48, 3, 48, 5, 48, 428, 10, 48, 3, 49, 3, 49, 3,
	49, 5, 49, 433, 10, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 7, 49, 440,
	10, 49, 12, 49, 14, 49, 443, 11, 49, 3, 50, 3, 50, 6, 50, 447, 10, 50,
	13, 50, 14, 50, 448, 3, 50, 3, 50, 3, 51, 5, 51, 454, 10, 51, 3, 51, 3,
	51, 5, 51, 458, 10, 51, 3, 51, 3, 51, 3, 51, 5, 51, 463, 10, 51, 3, 52,
	3, 52, 6, 52, 467, 10, 52, 13, 52, 14, 52, 468, 3, 52, 3, 52, 3, 53, 5,
	53, 474, 10, 53, 3, 53, 3, 53, 3, 53, 3, 53, 5, 53, 480, 10, 53, 3, 53,
	5, 53, 483, 10, 53, 3, 54, 3, 54, 3, 54, 3, 54, 5, 54, 489, 10, 54, 3,
	54, 5, 54, 492, 10, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 5, 54,
	500, 10, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 5, 54, 509,
	10, 54, 3, 55, 5, 55, 512, 10, 55, 3, 55, 3, 55, 3, 55, 5, 55, 517, 10,
	55, 3, 55, 5, 55, 520, 10, 55, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56, 526,
	10, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56, 535, 10,
	56, 3, 57, 5, 57, 538, 10, 57, 3, 57, 3, 57, 3, 57, 5, 57, 543, 10, 57,
	3, 57, 5, 57, 546, 10, 57, 3, 58, 3, 58, 5, 58, 550, 10, 58, 3, 58, 3,
	58, 3, 58, 7, 58, 555, 10, 58, 12, 58, 14, 58, 558, 11, 58, 3, 59, 3, 59,
	3, 59, 3, 59, 5, 59, 564, 10, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3,
	59, 3, 59, 5, 59, 573, 10, 59, 3, 60, 5, 60, 576, 10, 60, 3, 60, 3, 60,
	3, 60, 5, 60, 581, 10, 60, 3, 60, 5, 60, 584, 10, 60, 3, 61, 3, 61, 6,
	61, 588, 10, 61, 13, 61, 14, 61, 589, 3, 61, 3, 61, 3, 62, 5, 62, 595,
	10, 62, 3, 62, 3, 62, 5, 62, 599, 10, 62, 3, 63, 3, 63, 3, 63, 3, 63, 5,
	63, 605, 10, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 5, 63,
	614, 10, 63, 3, 64, 5, 64, 617, 10, 64, 3, 64, 3, 64, 3, 64, 5, 64, 622,
	10, 64, 3, 64, 5, 64, 625, 10, 64, 3, 65, 3, 65, 6, 65, 629, 10, 65, 13,
	65, 14, 65, 630, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 5, 66, 639,
	10, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 5, 66, 648, 10,
	66, 3, 67, 5, 67, 651, 10, 67, 3, 67, 3, 67, 3, 67, 3, 67, 5, 67, 657,
	10, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 7, 68, 665, 10, 68, 12,
	68, 14, 68, 668, 11, 68, 3, 69, 3, 69, 5, 69, 672, 10, 69, 3, 70, 3, 70,
	3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 2, 3, 96, 73, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
	88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118,
	120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 2, 7, 3, 2,
	3, 5, 3, 2, 14, 15, 3, 2, 54, 55, 3, 2, 35, 41, 3, 2, 42, 52, 2, 715, 2,
	145, 3, 2, 2, 2, 4, 152, 3, 2, 2, 2, 6, 156, 3, 2, 2, 2, 8, 171, 3, 2,
	2, 2, 10, 173, 3, 2, 2, 2, 12, 175, 3, 2, 2, 2, 14, 186, 3, 2, 2, 2, 16,
	189, 3, 2, 2, 2, 18, 201, 3, 2, 2, 2, 20, 209, 3, 2, 2, 2, 22, 213, 3,
	2, 2, 2, 24, 216, 3, 2, 2, 2, 26, 221, 3, 2, 2, 2, 28, 229, 3, 2, 2, 2,
	30, 231, 3, 2, 2, 2, 32, 234, 3, 2, 2, 2, 34, 252, 3, 2, 2, 2, 36, 254,
	3, 2, 2, 2, 38, 256, 3, 2, 2, 2, 40, 258, 3, 2, 2, 2, 42, 260, 3, 2, 2,
	2, 44, 262, 3, 2, 2, 2, 46, 264, 3, 2, 2, 2, 48, 276, 3, 2, 2, 2, 50, 278,
	3, 2, 2, 2, 52, 287, 3, 2, 2, 2, 54, 291, 3, 2, 2, 2, 56, 294, 3, 2, 2,
	2, 58, 302, 3, 2, 2, 2, 60, 308, 3, 2, 2, 2, 62, 319, 3, 2, 2, 2, 64, 321,
	3, 2, 2, 2, 66, 323, 3, 2, 2, 2, 68, 328, 3, 2, 2, 2, 70, 332, 3, 2, 2,
	2, 72, 340, 3, 2, 2, 2, 74, 344, 3, 2, 2, 2, 76, 346, 3, 2, 2, 2, 78, 358,
	3, 2, 2, 2, 80, 378, 3, 2, 2, 2, 82, 380, 3, 2, 2, 2, 84, 384, 3, 2, 2,
	2, 86, 392, 3, 2, 2, 2, 88, 400, 3, 2, 2, 2, 90, 403, 3, 2, 2, 2, 92, 410,
	3, 2, 2, 2, 94, 416, 3, 2, 2, 2, 96, 429, 3, 2, 2, 2, 98, 444, 3, 2, 2,
	2, 100, 453, 3, 2, 2, 2, 102, 464, 3, 2, 2, 2, 104, 473, 3, 2, 2, 2, 106,
	508, 3, 2, 2, 2, 108, 511, 3, 2, 2, 2, 110, 534, 3, 2, 2, 2, 112, 537,
	3, 2, 2, 2, 114, 547, 3, 2, 2, 2, 116, 572, 3, 2, 2, 2, 118, 575, 3, 2,
	2, 2, 120, 585, 3, 2, 2, 2, 122, 594, 3, 2, 2, 2, 124, 613, 3, 2, 2, 2,
	126, 616, 3, 2, 2, 2, 128, 626, 3, 2, 2, 2, 130, 647, 3, 2, 2, 2, 132,
	650, 3, 2, 2, 2, 134, 661, 3, 2, 2, 2, 136, 671, 3, 2, 2, 2, 138, 673,
	3, 2, 2, 2, 140, 675, 3, 2, 2, 2, 142, 677, 3, 2, 2, 2, 144, 146, 5, 4,
	3, 2, 145, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2,
	147, 148, 3, 2, 2, 2, 148, 3, 3, 2, 2, 2, 149, 153, 5, 6, 4, 2, 150, 153,
	5, 72, 37, 2, 151, 153, 5, 74, 38, 2, 152, 149, 3, 2, 2, 2, 152, 150, 3,
	2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 5, 3, 2, 2, 2, 154, 157, 5, 8, 5, 2,
	155, 157, 5, 26, 14, 2, 156, 154, 3, 2, 2, 2, 156, 155, 3, 2, 2, 2, 157,
	7, 3, 2, 2, 2, 158, 160, 5, 10, 6, 2, 159, 161, 5, 142, 72, 2, 160, 159,
	3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 163, 3, 2, 2, 2, 162, 164, 5, 56,
	29, 2, 163, 162, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 166, 3, 2, 2, 2,
	165, 167, 5, 68, 35, 2, 166, 165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167,
	168, 3, 2, 2, 2, 168, 169, 5, 12, 7, 2, 169, 172, 3, 2, 2, 2, 170, 172,
	5, 12, 7, 2, 171, 158, 3, 2, 2, 2, 171, 170, 3, 2, 2, 2, 172, 9, 3, 2,
	2, 2, 173, 174, 9, 2, 2, 2, 174, 11, 3, 2, 2, 2, 175, 177, 7, 6, 2, 2,
	176, 178, 5, 14, 8, 2, 177, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179,
	177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182,
	7, 7, 2, 2, 182, 13, 3, 2, 2, 2, 183, 187, 5, 16, 9, 2, 184, 187, 5, 24,
	13, 2, 185, 187, 5, 32, 17, 2, 186, 183, 3, 2, 2, 2, 186, 184, 3, 2, 2,
	2, 186, 185, 3, 2, 2, 2, 187, 15, 3, 2, 2, 2, 188, 190, 5, 22, 12, 2, 189,
	188, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 193,
	5, 142, 72, 2, 192, 194, 5, 18, 10, 2, 193, 192, 3, 2, 2, 2, 193, 194,
	3, 2, 2, 2, 194, 196, 3, 2, 2, 2, 195, 197, 5, 68, 35, 2, 196, 195, 3,
	2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 199, 3, 2, 2, 2, 198, 200, 5, 12, 7,
	2, 199, 198, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 17, 3, 2, 2, 2, 201,
	203, 7, 8, 2, 2, 202, 204, 5, 20, 11, 2, 203, 202, 3, 2, 2, 2, 204, 205,
	3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 3, 2,
	2, 2, 207, 208, 7, 9, 2, 2, 208, 19, 3, 2, 2, 2, 209, 210, 5, 142, 72,
	2, 210, 211, 7, 10, 2, 2, 211, 212, 5, 34, 18, 2, 212, 21, 3, 2, 2, 2,
	213, 214, 5, 142, 72, 2, 214, 215, 7, 10, 2, 2, 215, 23, 3, 2, 2, 2, 216,
	217, 7, 11, 2, 2, 217, 219, 5, 28, 15, 2, 218, 220, 5, 68, 35, 2, 219,
	218, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 25, 3, 2, 2, 2, 221, 222, 7,
	12, 2, 2, 222, 223, 5, 28, 15, 2, 223, 225, 5, 30, 16, 2, 224, 226, 5,
	68, 35, 2, 225, 224, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 227, 3, 2,
	2, 2, 227, 228, 5, 12, 7, 2, 228, 27, 3, 2, 2, 2, 229, 230, 5, 142, 72,
	2, 230, 29, 3, 2, 2, 2, 231, 232, 7, 13, 2, 2, 232, 233, 5, 64, 33, 2,
	233, 31, 3, 2, 2, 2, 234, 236, 7, 11, 2, 2, 235, 237, 5, 30, 16, 2, 236,
	235, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 239, 3, 2, 2, 2, 238, 240,
	5, 68, 35, 2, 239, 238, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 241, 3,
	2, 2, 2, 241, 242, 5, 12, 7, 2, 242, 33, 3, 2, 2, 2, 243, 253, 5, 54, 28,
	2, 244, 253, 5, 36, 19, 2, 245, 253, 5, 38, 20, 2, 246, 253, 5, 42, 22,
	2, 247, 253, 5, 40, 21, 2, 248, 253, 5, 44, 23, 2, 249, 253, 5, 46, 24,
	2, 250, 253, 5, 48, 25, 2, 251, 253, 5, 50, 26, 2, 252, 243, 3, 2, 2, 2,
	252, 244, 3, 2, 2, 2, 252, 245, 3, 2, 2, 2, 252, 246, 3, 2, 2, 2, 252,
	247, 3, 2, 2, 2, 252, 248, 3, 2, 2, 2, 252, 249, 3, 2, 2, 2, 252, 250,
	3, 2, 2, 2, 252, 251, 3, 2, 2, 2, 253, 35, 3, 2, 2, 2, 254, 255, 7, 58,
	2, 2, 255, 37, 3, 2, 2, 2, 256, 257, 7, 57, 2, 2, 257, 39, 3, 2, 2, 2,
	258, 259, 9, 3, 2, 2, 259, 41, 3, 2, 2, 2, 260, 261, 9, 4, 2, 2, 261, 43,
	3, 2, 2, 2, 262, 263, 7, 16, 2, 2, 263, 45, 3, 2, 2, 2, 264, 265, 5, 142,
	72, 2, 265, 47, 3, 2, 2, 2, 266, 267, 7, 17, 2, 2, 267, 277, 7, 18, 2,
	2, 268, 270, 7, 17, 2, 2, 269, 271, 5, 34, 18, 2, 270, 269, 3, 2, 2, 2,
	271, 272, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273,
	274, 3, 2, 2, 2, 274, 275, 7, 18, 2, 2, 275, 277, 3, 2, 2, 2, 276, 266,
	3, 2, 2, 2, 276, 268, 3, 2, 2, 2, 277, 49, 3, 2, 2, 2, 278, 282, 7, 6,
	2, 2, 279, 281, 5, 52, 27, 2, 280, 279, 3, 2, 2, 2, 281, 284, 3, 2, 2,
	2, 282, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 285, 3, 2, 2, 2, 284,
	282, 3, 2, 2, 2, 285, 286, 7, 7, 2, 2, 286, 51, 3, 2, 2, 2, 287, 288, 5,
	142, 72, 2, 288, 289, 7, 10, 2, 2, 289, 290, 5, 34, 18, 2, 290, 53, 3,
	2, 2, 2, 291, 292, 7, 19, 2, 2, 292, 293, 5, 142, 72, 2, 293, 55, 3, 2,
	2, 2, 294, 296, 7, 8, 2, 2, 295, 297, 5, 58, 30, 2, 296, 295, 3, 2, 2,
	2, 297, 298, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299,
	300, 3, 2, 2, 2, 300, 301, 7, 9, 2, 2, 301, 57, 3, 2, 2, 2, 302, 303, 5,
	54, 28, 2, 303, 304, 7, 10, 2, 2, 304, 306, 5, 62, 32, 2, 305, 307, 5,
	60, 31, 2, 306, 305, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 59, 3, 2, 2,
	2, 308, 309, 7, 20, 2, 2, 309, 310, 5, 34, 18, 2, 310, 61, 3, 2, 2, 2,
	311, 313, 5, 64, 33, 2, 312, 314, 7, 21, 2, 2, 313, 312, 3, 2, 2, 2, 313,
	314, 3, 2, 2, 2, 314, 320, 3, 2, 2, 2, 315, 317, 5, 66, 34, 2, 316, 318,
	7, 21, 2, 2, 317, 316, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 320, 3, 2,
	2, 2, 319, 311, 3, 2, 2, 2, 319, 315, 3, 2, 2, 2, 320, 63, 3, 2, 2, 2,
	321, 322, 5, 142, 72, 2, 322, 65, 3, 2, 2, 2, 323, 324, 7, 17, 2, 2, 324,
	325, 5, 62, 32, 2, 325, 326, 7, 18, 2, 2, 326, 67, 3, 2, 2, 2, 327, 329,
	5, 70, 36, 2, 328, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 328, 3,
	2, 2, 2, 330, 331, 3, 2, 2, 2, 331, 69, 3, 2, 2, 2, 332, 333, 7, 22, 2,
	2, 333, 335, 5, 142, 72, 2, 334, 336, 5, 18, 10, 2, 335, 334, 3, 2, 2,
	2, 335, 336, 3, 2, 2, 2, 336, 71, 3, 2, 2, 2, 337, 341, 5, 76, 39, 2, 338,
	341, 5, 86, 44, 2, 339, 341, 5, 132, 67, 2, 340, 337, 3, 2, 2, 2, 340,
	338, 3, 2, 2, 2, 340, 339, 3, 2, 2, 2, 341, 73, 3, 2, 2, 2, 342, 345, 5,
	80, 41, 2, 343, 345, 5, 88, 45, 2, 344, 342, 3, 2, 2, 2, 344, 343, 3, 2,
	2, 2, 345, 75, 3, 2, 2, 2, 346, 348, 7, 23, 2, 2, 347, 349, 5, 68, 35,
	2, 348, 347, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350,
	352, 7, 6, 2, 2, 351, 353, 5, 78, 40, 2, 352, 351, 3, 2, 2, 2, 353, 354,
	3, 2, 2, 2, 354, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 356, 3, 2,
	2, 2, 356, 357, 7, 7, 2, 2, 357, 77, 3, 2, 2, 2, 358, 359, 5, 10, 6, 2,
	359, 360, 7, 10, 2, 2, 360, 361, 5, 64, 33, 2, 361, 79, 3, 2, 2, 2, 362,
	363, 7, 24, 2, 2, 363, 365, 7, 23, 2, 2, 364, 366, 5, 68, 35, 2, 365, 364,
	3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 369, 7, 6,
	2, 2, 368, 370, 5, 82, 42, 2, 369, 368, 3, 2, 2, 2, 370, 371, 3, 2, 2,
	2, 371, 369, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373,
	374, 7, 7, 2, 2, 374, 379, 3, 2, 2, 2, 375, 376, 7, 24, 2, 2, 376, 377,
	7, 23, 2, 2, 377, 379, 5, 68, 35, 2, 378, 362, 3, 2, 2, 2, 378, 375, 3,
	2, 2, 2, 379, 81, 3, 2, 2, 2, 380, 381, 5, 10, 6, 2, 381, 382, 7, 10, 2,
	2, 382, 383, 5, 64, 33, 2, 383, 83, 3, 2, 2, 2, 384, 385, 5, 42, 22, 2,
	385, 85, 3, 2, 2, 2, 386, 393, 5, 90, 46, 2, 387, 393, 5, 94, 48, 2, 388,
	393, 5, 108, 55, 2, 389, 393, 5, 112, 57, 2, 390, 393, 5, 118, 60, 2, 391,
	393, 5, 126, 64, 2, 392, 386, 3, 2, 2, 2, 392, 387, 3, 2, 2, 2, 392, 388,
	3, 2, 2, 2, 392, 389, 3, 2, 2, 2, 392, 390, 3, 2, 2, 2, 392, 391, 3, 2,
	2, 2, 393, 87, 3, 2, 2, 2, 394, 401, 5, 92, 47, 2, 395, 401, 5, 106, 54,
	2, 396, 401, 5, 110, 56, 2, 397, 401, 5, 116, 59, 2, 398, 401, 5, 124,
	63, 2, 399, 401, 5, 130, 66, 2, 400, 394, 3, 2, 2, 2, 400, 395, 3, 2, 2,
	2, 400, 396, 3, 2, 2, 2, 400, 397, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2, 400,
	399, 3, 2, 2, 2, 401, 89, 3, 2, 2, 2, 402, 404, 5, 84, 43, 2, 403, 402,
	3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 406, 7, 25,
	2, 2, 406, 408, 5, 142, 72, 2, 407, 409, 5, 68, 35, 2, 408, 407, 3, 2,
	2, 2, 408, 409, 3, 2, 2, 2, 409, 91, 3, 2, 2, 2, 410, 411, 7, 24, 2, 2,
	411, 412, 7, 25, 2, 2, 412, 413, 5, 142, 72, 2, 413, 414, 5, 68, 35, 2,
	414, 93, 3, 2, 2, 2, 415, 417, 5, 84, 43, 2, 416, 415, 3, 2, 2, 2, 416,
	417, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 419, 7, 26, 2, 2, 419, 421,
	5, 142, 72, 2, 420, 422, 5, 96, 49, 2, 421, 420, 3, 2, 2, 2, 421, 422,
	3, 2, 2, 2, 422, 424, 3, 2, 2, 2, 423, 425, 5, 68, 35, 2, 424, 423, 3,
	2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 427, 3, 2, 2, 2, 426, 428, 5, 98, 50,
	2, 427, 426, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 95, 3, 2, 2, 2, 429,
	430, 8, 49, 1, 2, 430, 432, 7, 27, 2, 2, 431, 433, 7, 28, 2, 2, 432, 431,
	3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 434, 3, 2, 2, 2, 434, 435, 5, 64,
	33, 2, 435, 441, 3, 2, 2, 2, 436, 437, 12, 3, 2, 2, 437, 438, 7, 28, 2,
	2, 438, 440, 5, 64, 33, 2, 439, 436, 3, 2, 2, 2, 440, 443, 3, 2, 2, 2,
	441, 439, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 97, 3, 2, 2, 2, 443, 441,
	3, 2, 2, 2, 444, 446, 7, 6, 2, 2, 445, 447, 5, 100, 51, 2, 446, 445, 3,
	2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 446, 3, 2, 2, 2, 448, 449, 3, 2, 2,
	2, 449, 450, 3, 2, 2, 2, 450, 451, 7, 7, 2, 2, 451, 99, 3, 2, 2, 2, 452,
	454, 5, 84, 43, 2, 453, 452, 3, 2, 2, 2, 453, 454, 3, 2, 2, 2, 454, 455,
	3, 2, 2, 2, 455, 457, 5, 142, 72, 2, 456, 458, 5, 102, 52, 2, 457, 456,
	3, 2, 2, 2, 457, 458, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 460, 7, 10,
	2, 2, 460, 462, 5, 62, 32, 2, 461, 463, 5, 68, 35, 2, 462, 461, 3, 2, 2,
	2, 462, 463, 3, 2, 2, 2, 463, 101, 3, 2, 2, 2, 464, 466, 7, 8, 2, 2, 465,
	467, 5, 104, 53, 2, 466, 465, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 468, 466,
	3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470, 471, 7, 9,
	2, 2, 471, 103, 3, 2, 2, 2, 472, 474, 5, 84, 43, 2, 473, 472, 3, 2, 2,
	2, 473, 474, 3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475, 476, 5, 142, 72, 2,
	476, 477, 7, 10, 2, 2, 477, 479, 5, 62, 32, 2, 478, 480, 5, 60, 31, 2,
	479, 478, 3, 2, 2, 2, 479, 480, 3, 2, 2, 2, 480, 482, 3, 2, 2, 2, 481,
	483, 5, 68, 35, 2, 482, 481, 3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 105,
	3, 2, 2, 2, 484, 485, 7, 24, 2, 2, 485, 486, 7, 26, 2, 2, 486, 488, 5,
	142, 72, 2, 487, 489, 5, 96, 49, 2, 488, 487, 3, 2, 2, 2, 488, 489, 3,
	2, 2, 2, 489, 491, 3, 2, 2, 2, 490, 492, 5, 68, 35, 2, 491, 490, 3, 2,
	2, 2, 491, 492, 3, 2, 2, 2, 492, 493, 3, 2, 2, 2, 493, 494, 5, 98, 50,
	2, 494, 509, 3, 2, 2, 2, 495, 496, 7, 24, 2, 2, 496, 497, 7, 26, 2, 2,
	497, 499, 5, 142, 72, 2, 498, 500, 5, 96, 49, 2, 499, 498, 3, 2, 2, 2,
	499, 500, 3, 2, 2, 2, 500, 501, 3, 2, 2, 2, 501, 502, 5, 68, 35, 2, 502,
	509, 3, 2, 2, 2, 503, 504, 7, 24, 2, 2, 504, 505, 7, 26, 2, 2, 505, 506,
	5, 142, 72, 2, 506, 507, 5, 96, 49, 2, 507, 509, 3, 2, 2, 2, 508, 484,
	3, 2, 2, 2, 508, 495, 3, 2, 2, 2, 508, 503, 3, 2, 2, 2, 509, 107, 3, 2,
	2, 2, 510, 512, 5, 84, 43, 2, 511, 510, 3, 2, 2, 2, 511, 512, 3, 2, 2,
	2, 512, 513, 3, 2, 2, 2, 513, 514, 7, 29, 2, 2, 514, 516, 5, 142, 72, 2,
	515, 517, 5, 68, 35, 2, 516, 515, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517,
	519, 3, 2, 2, 2, 518, 520, 5, 98, 50, 2, 519, 518, 3, 2, 2, 2, 519, 520,
	3, 2, 2, 2, 520, 109, 3, 2, 2, 2, 521, 522, 7, 24, 2, 2, 522, 523, 7, 29,
	2, 2, 523, 525, 5, 142, 72, 2, 524, 526, 5, 68, 35, 2, 525, 524, 3, 2,
	2, 2, 525, 526, 3, 2, 2, 2, 526, 527, 3, 2, 2, 2, 527, 528, 5, 98, 50,
	2, 528, 535, 3, 2, 2, 2, 529, 530, 7, 24, 2, 2, 530, 531, 7, 29, 2, 2,
	531, 532, 5, 142, 72, 2, 532, 533, 5, 68, 35, 2, 533, 535, 3, 2, 2, 2,
	534, 521, 3, 2, 2, 2, 534, 529, 3, 2, 2, 2, 535, 111, 3, 2, 2, 2, 536,
	538, 5, 84, 43, 2, 537, 536, 3, 2, 2, 2, 537, 538, 3, 2, 2, 2, 538, 539,
	3, 2, 2, 2, 539, 540, 7, 30, 2, 2, 540, 542, 5, 142, 72, 2, 541, 543, 5,
	68, 35, 2, 542, 541, 3, 2, 2, 2, 542, 543, 3, 2, 2, 2, 543, 545, 3, 2,
	2, 2, 544, 546, 5, 114, 58, 2, 545, 544, 3, 2, 2, 2, 545, 546, 3, 2, 2,
	2, 546, 113, 3, 2, 2, 2, 547, 549, 7, 20, 2, 2, 548, 550, 7, 31, 2, 2,
	549, 548, 3, 2, 2, 2, 549, 550, 3, 2, 2, 2, 550, 551, 3, 2, 2, 2, 551,
	556, 5, 64, 33, 2, 552, 553, 7, 31, 2, 2, 553, 555, 5, 64, 33, 2, 554,
	552, 3, 2, 2, 2, 555, 558, 3, 2, 2, 2, 556, 554, 3, 2, 2, 2, 556, 557,
	3, 2, 2, 2, 557, 115, 3, 2, 2, 2, 558, 556, 3, 2, 2, 2, 559, 560, 7, 24,
	2, 2, 560, 561, 7, 30, 2, 2, 561, 563, 5, 142, 72, 2, 562, 564, 5, 68,
	35, 2, 563, 562, 3, 2, 2, 2, 563, 564, 3, 2, 2, 2, 564, 565, 3, 2, 2, 2,
	565, 566, 5, 114, 58, 2, 566, 573, 3, 2, 2, 2, 567, 568, 7, 24, 2, 2, 568,
	569, 7, 30, 2, 2, 569, 570, 5, 142, 72, 2, 570, 571, 5, 68, 35, 2, 571,
	573, 3, 2, 2, 2, 572, 559, 3, 2, 2, 2, 572, 567, 3, 2, 2, 2, 573, 117,
	3, 2, 2, 2, 574, 576, 5, 84, 43, 2, 575, 574, 3, 2, 2, 2, 575, 576, 3,
	2, 2, 2, 576, 577, 3, 2, 2, 2, 577, 578, 7, 32, 2, 2, 578, 580, 5, 142,
	72, 2, 579, 581, 5, 68, 35, 2, 580, 579, 3, 2, 2, 2, 580, 581, 3, 2, 2,
	2, 581, 583, 3, 2, 2, 2, 582, 584, 5, 120, 61, 2, 583, 582, 3, 2, 2, 2,
	583, 584, 3, 2, 2, 2, 584, 119, 3, 2, 2, 2, 585, 587, 7, 6, 2, 2, 586,
	588, 5, 122, 62, 2, 587, 586, 3, 2, 2, 2, 588, 589, 3, 2, 2, 2, 589, 587,
	3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590, 591, 3, 2, 2, 2, 591, 592, 7, 7,
	2, 2, 592, 121, 3, 2, 2, 2, 593, 595, 5, 84, 43, 2, 594, 593, 3, 2, 2,
	2, 594, 595, 3, 2, 2, 2, 595, 596, 3, 2, 2, 2, 596, 598, 5, 46, 24, 2,
	597, 599, 5, 68, 35, 2, 598, 597, 3, 2, 2, 2, 598, 599, 3, 2, 2, 2, 599,
	123, 3, 2, 2, 2, 600, 601, 7, 24, 2, 2, 601, 602, 7, 32, 2, 2, 602, 604,
	5, 142, 72, 2, 603, 605, 5, 68, 35, 2, 604, 603, 3, 2, 2, 2, 604, 605,
	3, 2, 2, 2, 605, 606, 3, 2, 2, 2, 606, 607, 5, 120, 61, 2, 607, 614, 3,
	2, 2, 2, 608, 609, 7, 24, 2, 2, 609, 610, 7, 32, 2, 2, 610, 611, 5, 142,
	72, 2, 611, 612, 5, 68, 35, 2, 612, 614, 3, 2, 2, 2, 613, 600, 3, 2, 2,
	2, 613, 608, 3, 2, 2, 2, 614, 125, 3, 2, 2, 2, 615, 617, 5, 84, 43, 2,
	616, 615, 3, 2, 2, 2, 616, 617, 3, 2, 2, 2, 617, 618, 3, 2, 2, 2, 618,
	619, 7, 33, 2, 2, 619, 621, 5, 142, 72, 2, 620, 622, 5, 68, 35, 2, 621,
	620, 3, 2, 2, 2, 621, 622, 3, 2, 2, 2, 622, 624, 3, 2, 2, 2, 623, 625,
	5, 128, 65, 2, 624, 623, 3, 2, 2, 2, 624, 625, 3, 2, 2, 2, 625, 127, 3,
	2, 2, 2, 626, 628, 7, 6, 2, 2, 627, 629, 5, 104, 53, 2, 628, 627, 3, 2,
	2, 2, 629, 630, 3, 2, 2, 2, 630, 628, 3, 2, 2, 2, 630, 631, 3, 2, 2, 2,
	631, 632, 3, 2, 2, 2, 632, 633, 7, 7, 2, 2, 633, 129, 3, 2, 2, 2, 634,
	635, 7, 24, 2, 2, 635, 636, 7, 33, 2, 2, 636, 638, 5, 142, 72, 2, 637,
	639, 5, 68, 35, 2, 638, 637, 3, 2, 2, 2, 638, 639, 3, 2, 2, 2, 639, 640,
	3, 2, 2, 2, 640, 641, 5, 128, 65, 2, 641, 648, 3, 2, 2, 2, 642, 643, 7,
	24, 2, 2, 643, 644, 7, 33, 2, 2, 644, 645, 5, 142, 72, 2, 645, 646, 5,
	68, 35, 2, 646, 648, 3, 2, 2, 2, 647, 634, 3, 2, 2, 2, 647, 642, 3, 2,
	2, 2, 648, 131, 3, 2, 2, 2, 649, 651, 5, 84, 43, 2, 650, 649, 3, 2, 2,
	2, 650, 651, 3, 2, 2, 2, 651, 652, 3, 2, 2, 2, 652, 653, 7, 34, 2, 2, 653,
	654, 7, 22, 2, 2, 654, 656, 5, 142, 72, 2, 655, 657, 5, 102, 52, 2, 656,
	655, 3, 2, 2, 2, 656, 657, 3, 2, 2, 2, 657, 658, 3, 2, 2, 2, 658, 659,
	7, 13, 2, 2, 659, 660, 5, 134, 68, 2, 660, 133, 3, 2, 2, 2, 661, 666, 5,
	136, 69, 2, 662, 663, 7, 31, 2, 2, 663, 665, 5, 136, 69, 2, 664, 662, 3,
	2, 2, 2, 665, 668, 3, 2, 2, 2, 666, 664, 3, 2, 2, 2, 666, 667, 3, 2, 2,
	2, 667, 135, 3, 2, 2, 2, 668, 666, 3, 2, 2, 2, 669, 672, 5, 138, 70, 2,
	670, 672, 5, 140, 71, 2, 671, 669, 3, 2, 2, 2, 671, 670, 3, 2, 2, 2, 672,
	137, 3, 2, 2, 2, 673, 674, 9, 5, 2, 2, 674, 139, 3, 2, 2, 2, 675, 676,
	9, 6, 2, 2, 676, 141, 3, 2, 2, 2, 677, 678, 7, 53, 2, 2, 678, 143, 3, 2,
	2, 2, 90, 147, 152, 156, 160, 163, 166, 171, 179, 186, 189, 193, 196, 199,
	205, 219, 225, 236, 239, 252, 272, 276, 282, 298, 306, 313, 317, 319, 330,
	335, 340, 344, 348, 354, 365, 371, 378, 392, 400, 403, 408, 416, 421, 424,
	427, 432, 441, 448, 453, 457, 462, 468, 473, 479, 482, 488, 491, 499, 508,
	511, 516, 519, 525, 534, 537, 542, 545, 549, 556, 563, 572, 575, 580, 583,
	589, 594, 598, 604, 613, 616, 621, 624, 630, 638, 647, 650, 656, 666, 671,
}
var literalNames = []string{
	"", "'query'", "'mutation'", "'subscription'", "'{'", "'}'", "'('", "')'",
	"':'", "'...'", "'fragment'", "'on'", "'true'", "'false'", "'null'", "'['",
	"']'", "'$'", "'='", "'!'", "'@'", "'schema'", "'extend'", "'scalar'",
	"'type'", "'implements'", "'&'", "'interface'", "'union'", "'|'", "'enum'",
	"'input'", "'directive'", "'QUERY'", "'MUTATION'", "'SUBSCRIPTION'", "'FIELD'",
	"'FRAGMENT_DEFINITION'", "'FRAGMENT_SPREAD'", "'INLINE_FRAGMENT'", "'SCHEMA'",
	"'SCALAR'", "'OBJECT'", "'FIELD_DEFINITION'", "'ARGUMENT_DEFINITION'",
	"'INTERFACE'", "'UNION'", "'ENUM'", "'ENUM_VALUE'", "'INPUT_OBJECT'", "'INPUT_FIELD_DEFINITION'",
	"", "", "", "", "", "", "", "", "','", "", "", "'\uEFBBBF'", "'\uFEFF'",
	"'\u0000FEFF'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NAME", "STRING",
	"BLOCK_STRING", "ID", "FLOAT", "INT", "PUNCTUATOR", "WS", "COMMA", "LineComment",
	"UNICODE_BOM", "UTF8_BOM", "UTF16_BOM", "UTF32_BOM",
}

var ruleNames = []string{
	"document", "definition", "executableDefinition", "operationDefinition",
	"operationType", "selectionSet", "selection", "field", "arguments", "argument",
	"alias", "fragmentSpread", "fragmentDefinition", "fragmentName", "typeCondition",
	"inlineFragment", "value", "intValue", "floatValue", "booleanValue", "stringValue",
	"nullValue", "enumValue", "listValue", "objectValue", "objectField", "variable",
	"variableDefinitions", "variableDefinition", "defaultValue", "type_", "namedType",
	"listType", "directives", "directive", "typeSystemDefinition", "typeSystemExtension",
	"schemaDefinition", "rootOperationTypeDefinition", "schemaExtension", "operationTypeDefinition",
	"description", "typeDefinition", "typeExtension", "scalarTypeDefinition",
	"scalarTypeExtension", "objectTypeDefinition", "implementsInterfaces",
	"fieldsDefinition", "fieldDefinition", "argumentsDefinition", "inputValueDefinition",
	"objectTypeExtension", "interfaceTypeDefinition", "interfaceTypeExtension",
	"unionTypeDefinition", "unionMemberTypes", "unionTypeExtension", "enumTypeDefinition",
	"enumValuesDefinition", "enumValueDefinition", "enumTypeExtension", "inputObjectTypeDefinition",
	"inputFieldsDefinition", "inputObjectTypeExtension", "directiveDefinition",
	"directiveLocations", "directiveLocation", "executableDirectiveLocation",
	"typeSystemDirectiveLocation", "name",
}

type GraphQLParser struct {
	*antlr.BaseParser
}

// NewGraphQLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GraphQLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGraphQLParser(input antlr.TokenStream) *GraphQLParser {
	this := new(GraphQLParser)
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
	this.GrammarFileName = "GraphQL.g4"

	return this
}

// GraphQLParser tokens.
const (
	GraphQLParserEOF          = antlr.TokenEOF
	GraphQLParserT__0         = 1
	GraphQLParserT__1         = 2
	GraphQLParserT__2         = 3
	GraphQLParserT__3         = 4
	GraphQLParserT__4         = 5
	GraphQLParserT__5         = 6
	GraphQLParserT__6         = 7
	GraphQLParserT__7         = 8
	GraphQLParserT__8         = 9
	GraphQLParserT__9         = 10
	GraphQLParserT__10        = 11
	GraphQLParserT__11        = 12
	GraphQLParserT__12        = 13
	GraphQLParserT__13        = 14
	GraphQLParserT__14        = 15
	GraphQLParserT__15        = 16
	GraphQLParserT__16        = 17
	GraphQLParserT__17        = 18
	GraphQLParserT__18        = 19
	GraphQLParserT__19        = 20
	GraphQLParserT__20        = 21
	GraphQLParserT__21        = 22
	GraphQLParserT__22        = 23
	GraphQLParserT__23        = 24
	GraphQLParserT__24        = 25
	GraphQLParserT__25        = 26
	GraphQLParserT__26        = 27
	GraphQLParserT__27        = 28
	GraphQLParserT__28        = 29
	GraphQLParserT__29        = 30
	GraphQLParserT__30        = 31
	GraphQLParserT__31        = 32
	GraphQLParserT__32        = 33
	GraphQLParserT__33        = 34
	GraphQLParserT__34        = 35
	GraphQLParserT__35        = 36
	GraphQLParserT__36        = 37
	GraphQLParserT__37        = 38
	GraphQLParserT__38        = 39
	GraphQLParserT__39        = 40
	GraphQLParserT__40        = 41
	GraphQLParserT__41        = 42
	GraphQLParserT__42        = 43
	GraphQLParserT__43        = 44
	GraphQLParserT__44        = 45
	GraphQLParserT__45        = 46
	GraphQLParserT__46        = 47
	GraphQLParserT__47        = 48
	GraphQLParserT__48        = 49
	GraphQLParserT__49        = 50
	GraphQLParserNAME         = 51
	GraphQLParserSTRING       = 52
	GraphQLParserBLOCK_STRING = 53
	GraphQLParserID           = 54
	GraphQLParserFLOAT        = 55
	GraphQLParserINT          = 56
	GraphQLParserPUNCTUATOR   = 57
	GraphQLParserWS           = 58
	GraphQLParserCOMMA        = 59
	GraphQLParserLineComment  = 60
	GraphQLParserUNICODE_BOM  = 61
	GraphQLParserUTF8_BOM     = 62
	GraphQLParserUTF16_BOM    = 63
	GraphQLParserUTF32_BOM    = 64
)

// GraphQLParser rules.
const (
	GraphQLParserRULE_document                    = 0
	GraphQLParserRULE_definition                  = 1
	GraphQLParserRULE_executableDefinition        = 2
	GraphQLParserRULE_operationDefinition         = 3
	GraphQLParserRULE_operationType               = 4
	GraphQLParserRULE_selectionSet                = 5
	GraphQLParserRULE_selection                   = 6
	GraphQLParserRULE_field                       = 7
	GraphQLParserRULE_arguments                   = 8
	GraphQLParserRULE_argument                    = 9
	GraphQLParserRULE_alias                       = 10
	GraphQLParserRULE_fragmentSpread              = 11
	GraphQLParserRULE_fragmentDefinition          = 12
	GraphQLParserRULE_fragmentName                = 13
	GraphQLParserRULE_typeCondition               = 14
	GraphQLParserRULE_inlineFragment              = 15
	GraphQLParserRULE_value                       = 16
	GraphQLParserRULE_intValue                    = 17
	GraphQLParserRULE_floatValue                  = 18
	GraphQLParserRULE_booleanValue                = 19
	GraphQLParserRULE_stringValue                 = 20
	GraphQLParserRULE_nullValue                   = 21
	GraphQLParserRULE_enumValue                   = 22
	GraphQLParserRULE_listValue                   = 23
	GraphQLParserRULE_objectValue                 = 24
	GraphQLParserRULE_objectField                 = 25
	GraphQLParserRULE_variable                    = 26
	GraphQLParserRULE_variableDefinitions         = 27
	GraphQLParserRULE_variableDefinition          = 28
	GraphQLParserRULE_defaultValue                = 29
	GraphQLParserRULE_type_                       = 30
	GraphQLParserRULE_namedType                   = 31
	GraphQLParserRULE_listType                    = 32
	GraphQLParserRULE_directives                  = 33
	GraphQLParserRULE_directive                   = 34
	GraphQLParserRULE_typeSystemDefinition        = 35
	GraphQLParserRULE_typeSystemExtension         = 36
	GraphQLParserRULE_schemaDefinition            = 37
	GraphQLParserRULE_rootOperationTypeDefinition = 38
	GraphQLParserRULE_schemaExtension             = 39
	GraphQLParserRULE_operationTypeDefinition     = 40
	GraphQLParserRULE_description                 = 41
	GraphQLParserRULE_typeDefinition              = 42
	GraphQLParserRULE_typeExtension               = 43
	GraphQLParserRULE_scalarTypeDefinition        = 44
	GraphQLParserRULE_scalarTypeExtension         = 45
	GraphQLParserRULE_objectTypeDefinition        = 46
	GraphQLParserRULE_implementsInterfaces        = 47
	GraphQLParserRULE_fieldsDefinition            = 48
	GraphQLParserRULE_fieldDefinition             = 49
	GraphQLParserRULE_argumentsDefinition         = 50
	GraphQLParserRULE_inputValueDefinition        = 51
	GraphQLParserRULE_objectTypeExtension         = 52
	GraphQLParserRULE_interfaceTypeDefinition     = 53
	GraphQLParserRULE_interfaceTypeExtension      = 54
	GraphQLParserRULE_unionTypeDefinition         = 55
	GraphQLParserRULE_unionMemberTypes            = 56
	GraphQLParserRULE_unionTypeExtension          = 57
	GraphQLParserRULE_enumTypeDefinition          = 58
	GraphQLParserRULE_enumValuesDefinition        = 59
	GraphQLParserRULE_enumValueDefinition         = 60
	GraphQLParserRULE_enumTypeExtension           = 61
	GraphQLParserRULE_inputObjectTypeDefinition   = 62
	GraphQLParserRULE_inputFieldsDefinition       = 63
	GraphQLParserRULE_inputObjectTypeExtension    = 64
	GraphQLParserRULE_directiveDefinition         = 65
	GraphQLParserRULE_directiveLocations          = 66
	GraphQLParserRULE_directiveLocation           = 67
	GraphQLParserRULE_executableDirectiveLocation = 68
	GraphQLParserRULE_typeSystemDirectiveLocation = 69
	GraphQLParserRULE_name                        = 70
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *GraphQLParser) Document() (localctx IDocumentContext) {
	this := p
	_ = this

	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphQLParserRULE_document)
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
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__20)|(1<<GraphQLParserT__21)|(1<<GraphQLParserT__22)|(1<<GraphQLParserT__23)|(1<<GraphQLParserT__26)|(1<<GraphQLParserT__27)|(1<<GraphQLParserT__29)|(1<<GraphQLParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GraphQLParserT__31-32))|(1<<(GraphQLParserSTRING-32))|(1<<(GraphQLParserBLOCK_STRING-32)))) != 0) {
		{
			p.SetState(142)
			p.Definition()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) ExecutableDefinition() IExecutableDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecutableDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecutableDefinitionContext)
}

func (s *DefinitionContext) TypeSystemDefinition() ITypeSystemDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSystemDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSystemDefinitionContext)
}

func (s *DefinitionContext) TypeSystemExtension() ITypeSystemExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSystemExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSystemExtensionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *GraphQLParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphQLParserRULE_definition)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__0, GraphQLParserT__1, GraphQLParserT__2, GraphQLParserT__3, GraphQLParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.ExecutableDefinition()
		}

	case GraphQLParserT__20, GraphQLParserT__22, GraphQLParserT__23, GraphQLParserT__26, GraphQLParserT__27, GraphQLParserT__29, GraphQLParserT__30, GraphQLParserT__31, GraphQLParserSTRING, GraphQLParserBLOCK_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.TypeSystemDefinition()
		}

	case GraphQLParserT__21:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.TypeSystemExtension()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExecutableDefinitionContext is an interface to support dynamic dispatch.
type IExecutableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecutableDefinitionContext differentiates from other interfaces.
	IsExecutableDefinitionContext()
}

type ExecutableDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutableDefinitionContext() *ExecutableDefinitionContext {
	var p = new(ExecutableDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_executableDefinition
	return p
}

func (*ExecutableDefinitionContext) IsExecutableDefinitionContext() {}

func NewExecutableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutableDefinitionContext {
	var p = new(ExecutableDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_executableDefinition

	return p
}

func (s *ExecutableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecutableDefinitionContext) OperationDefinition() IOperationDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationDefinitionContext)
}

func (s *ExecutableDefinitionContext) FragmentDefinition() IFragmentDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentDefinitionContext)
}

func (s *ExecutableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterExecutableDefinition(s)
	}
}

func (s *ExecutableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitExecutableDefinition(s)
	}
}

func (p *GraphQLParser) ExecutableDefinition() (localctx IExecutableDefinitionContext) {
	this := p
	_ = this

	localctx = NewExecutableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraphQLParserRULE_executableDefinition)

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

	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__0, GraphQLParserT__1, GraphQLParserT__2, GraphQLParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.OperationDefinition()
		}

	case GraphQLParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.FragmentDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationDefinitionContext is an interface to support dynamic dispatch.
type IOperationDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationDefinitionContext differentiates from other interfaces.
	IsOperationDefinitionContext()
}

type OperationDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationDefinitionContext() *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationDefinition
	return p
}

func (*OperationDefinitionContext) IsOperationDefinitionContext() {}

func NewOperationDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationDefinition

	return p
}

func (s *OperationDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *OperationDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *OperationDefinitionContext) VariableDefinitions() IVariableDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionsContext)
}

func (s *OperationDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *OperationDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationDefinition(s)
	}
}

func (s *OperationDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationDefinition(s)
	}
}

func (p *GraphQLParser) OperationDefinition() (localctx IOperationDefinitionContext) {
	this := p
	_ = this

	localctx = NewOperationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraphQLParserRULE_operationDefinition)
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

	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__0, GraphQLParserT__1, GraphQLParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.OperationType()
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserNAME {
			{
				p.SetState(157)
				p.Name()
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__5 {
			{
				p.SetState(160)
				p.VariableDefinitions()
			}

		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(163)
				p.Directives()
			}

		}
		{
			p.SetState(166)
			p.SelectionSet()
		}

	case GraphQLParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.SelectionSet()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationTypeContext is an interface to support dynamic dispatch.
type IOperationTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeContext differentiates from other interfaces.
	IsOperationTypeContext()
}

type OperationTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeContext() *OperationTypeContext {
	var p = new(OperationTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationType
	return p
}

func (*OperationTypeContext) IsOperationTypeContext() {}

func NewOperationTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeContext {
	var p = new(OperationTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationType

	return p
}

func (s *OperationTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *OperationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationType(s)
	}
}

func (s *OperationTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationType(s)
	}
}

func (p *GraphQLParser) OperationType() (localctx IOperationTypeContext) {
	this := p
	_ = this

	localctx = NewOperationTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraphQLParserRULE_operationType)
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
		p.SetState(171)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISelectionSetContext is an interface to support dynamic dispatch.
type ISelectionSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionSetContext differentiates from other interfaces.
	IsSelectionSetContext()
}

type SelectionSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionSetContext() *SelectionSetContext {
	var p = new(SelectionSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_selectionSet
	return p
}

func (*SelectionSetContext) IsSelectionSetContext() {}

func NewSelectionSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionSetContext {
	var p = new(SelectionSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selectionSet

	return p
}

func (s *SelectionSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionSetContext) AllSelection() []ISelectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectionContext)(nil)).Elem())
	var tst = make([]ISelectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectionContext)
		}
	}

	return tst
}

func (s *SelectionSetContext) Selection(i int) ISelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectionContext)
}

func (s *SelectionSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelectionSet(s)
	}
}

func (s *SelectionSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelectionSet(s)
	}
}

func (p *GraphQLParser) SelectionSet() (localctx ISelectionSetContext) {
	this := p
	_ = this

	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraphQLParserRULE_selectionSet)
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
		p.Match(GraphQLParserT__3)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__8 || _la == GraphQLParserNAME {
		{
			p.SetState(174)
			p.Selection()
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(179)
		p.Match(GraphQLParserT__4)
	}

	return localctx
}

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_selection
	return p
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *SelectionContext) FragmentSpread() IFragmentSpreadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentSpreadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentSpreadContext)
}

func (s *SelectionContext) InlineFragment() IInlineFragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineFragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineFragmentContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelection(s)
	}
}

func (s *SelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelection(s)
	}
}

func (p *GraphQLParser) Selection() (localctx ISelectionContext) {
	this := p
	_ = this

	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GraphQLParserRULE_selection)

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

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.FragmentSpread()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(183)
			p.InlineFragment()
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
	p.RuleIndex = GraphQLParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *FieldContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FieldContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *GraphQLParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraphQLParserRULE_field)
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
	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(186)
			p.Alias()
		}

	}
	{
		p.SetState(189)
		p.Name()
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(190)
			p.Arguments()
		}

	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(193)
			p.Directives()
		}

	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__3 {
		{
			p.SetState(196)
			p.SelectionSet()
		}

	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GraphQLParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraphQLParserRULE_arguments)
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
		p.SetState(199)
		p.Match(GraphQLParserT__5)
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserNAME {
		{
			p.SetState(200)
			p.Argument()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(205)
		p.Match(GraphQLParserT__6)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ArgumentContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *GraphQLParser) Argument() (localctx IArgumentContext) {
	this := p
	_ = this

	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GraphQLParserRULE_argument)

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
		p.SetState(207)
		p.Name()
	}
	{
		p.SetState(208)
		p.Match(GraphQLParserT__7)
	}
	{
		p.SetState(209)
		p.Value()
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (p *GraphQLParser) Alias() (localctx IAliasContext) {
	this := p
	_ = this

	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GraphQLParserRULE_alias)

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
		p.Name()
	}
	{
		p.SetState(212)
		p.Match(GraphQLParserT__7)
	}

	return localctx
}

// IFragmentSpreadContext is an interface to support dynamic dispatch.
type IFragmentSpreadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentSpreadContext differentiates from other interfaces.
	IsFragmentSpreadContext()
}

type FragmentSpreadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentSpreadContext() *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentSpread
	return p
}

func (*FragmentSpreadContext) IsFragmentSpreadContext() {}

func NewFragmentSpreadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentSpread

	return p
}

func (s *FragmentSpreadContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentSpreadContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentSpreadContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentSpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentSpreadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentSpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentSpread(s)
	}
}

func (s *FragmentSpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentSpread(s)
	}
}

func (p *GraphQLParser) FragmentSpread() (localctx IFragmentSpreadContext) {
	this := p
	_ = this

	localctx = NewFragmentSpreadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GraphQLParserRULE_fragmentSpread)
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
		p.SetState(214)
		p.Match(GraphQLParserT__8)
	}
	{
		p.SetState(215)
		p.FragmentName()
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(216)
			p.Directives()
		}

	}

	return localctx
}

// IFragmentDefinitionContext is an interface to support dynamic dispatch.
type IFragmentDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentDefinitionContext differentiates from other interfaces.
	IsFragmentDefinitionContext()
}

type FragmentDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentDefinitionContext() *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition
	return p
}

func (*FragmentDefinitionContext) IsFragmentDefinitionContext() {}

func NewFragmentDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition

	return p
}

func (s *FragmentDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentDefinitionContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentDefinitionContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *FragmentDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FragmentDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentDefinition(s)
	}
}

func (s *FragmentDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentDefinition(s)
	}
}

func (p *GraphQLParser) FragmentDefinition() (localctx IFragmentDefinitionContext) {
	this := p
	_ = this

	localctx = NewFragmentDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GraphQLParserRULE_fragmentDefinition)
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
		p.SetState(219)
		p.Match(GraphQLParserT__9)
	}
	{
		p.SetState(220)
		p.FragmentName()
	}
	{
		p.SetState(221)
		p.TypeCondition()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(222)
			p.Directives()
		}

	}
	{
		p.SetState(225)
		p.SelectionSet()
	}

	return localctx
}

// IFragmentNameContext is an interface to support dynamic dispatch.
type IFragmentNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentNameContext differentiates from other interfaces.
	IsFragmentNameContext()
}

type FragmentNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentNameContext() *FragmentNameContext {
	var p = new(FragmentNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentName
	return p
}

func (*FragmentNameContext) IsFragmentNameContext() {}

func NewFragmentNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentNameContext {
	var p = new(FragmentNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentName

	return p
}

func (s *FragmentNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentNameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FragmentNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentName(s)
	}
}

func (s *FragmentNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentName(s)
	}
}

func (p *GraphQLParser) FragmentName() (localctx IFragmentNameContext) {
	this := p
	_ = this

	localctx = NewFragmentNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GraphQLParserRULE_fragmentName)

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
		p.SetState(227)
		p.Name()
	}

	return localctx
}

// ITypeConditionContext is an interface to support dynamic dispatch.
type ITypeConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeConditionContext differentiates from other interfaces.
	IsTypeConditionContext()
}

type TypeConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeConditionContext() *TypeConditionContext {
	var p = new(TypeConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeCondition
	return p
}

func (*TypeConditionContext) IsTypeConditionContext() {}

func NewTypeConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeConditionContext {
	var p = new(TypeConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeCondition

	return p
}

func (s *TypeConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeConditionContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *TypeConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeCondition(s)
	}
}

func (s *TypeConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeCondition(s)
	}
}

func (p *GraphQLParser) TypeCondition() (localctx ITypeConditionContext) {
	this := p
	_ = this

	localctx = NewTypeConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GraphQLParserRULE_typeCondition)

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
		p.Match(GraphQLParserT__10)
	}
	{
		p.SetState(230)
		p.NamedType()
	}

	return localctx
}

// IInlineFragmentContext is an interface to support dynamic dispatch.
type IInlineFragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineFragmentContext differentiates from other interfaces.
	IsInlineFragmentContext()
}

type InlineFragmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineFragmentContext() *InlineFragmentContext {
	var p = new(InlineFragmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inlineFragment
	return p
}

func (*InlineFragmentContext) IsInlineFragmentContext() {}

func NewInlineFragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineFragmentContext {
	var p = new(InlineFragmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inlineFragment

	return p
}

func (s *InlineFragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineFragmentContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *InlineFragmentContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *InlineFragmentContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InlineFragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineFragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineFragmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInlineFragment(s)
	}
}

func (s *InlineFragmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInlineFragment(s)
	}
}

func (p *GraphQLParser) InlineFragment() (localctx IInlineFragmentContext) {
	this := p
	_ = this

	localctx = NewInlineFragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GraphQLParserRULE_inlineFragment)
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
		p.SetState(232)
		p.Match(GraphQLParserT__8)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__10 {
		{
			p.SetState(233)
			p.TypeCondition()
		}

	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(236)
			p.Directives()
		}

	}
	{
		p.SetState(239)
		p.SelectionSet()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ValueContext) IntValue() IIntValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntValueContext)
}

func (s *ValueContext) FloatValue() IFloatValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatValueContext)
}

func (s *ValueContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *ValueContext) BooleanValue() IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *ValueContext) NullValue() INullValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullValueContext)
}

func (s *ValueContext) EnumValue() IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *ValueContext) ListValue() IListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListValueContext)
}

func (s *ValueContext) ObjectValue() IObjectValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectValueContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *GraphQLParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GraphQLParserRULE_value)

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

	p.SetState(250)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Variable()
		}

	case GraphQLParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.IntValue()
		}

	case GraphQLParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)
			p.FloatValue()
		}

	case GraphQLParserSTRING, GraphQLParserBLOCK_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(244)
			p.StringValue()
		}

	case GraphQLParserT__11, GraphQLParserT__12:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(245)
			p.BooleanValue()
		}

	case GraphQLParserT__13:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(246)
			p.NullValue()
		}

	case GraphQLParserNAME:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(247)
			p.EnumValue()
		}

	case GraphQLParserT__14:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(248)
			p.ListValue()
		}

	case GraphQLParserT__3:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(249)
			p.ObjectValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntValueContext is an interface to support dynamic dispatch.
type IIntValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntValueContext differentiates from other interfaces.
	IsIntValueContext()
}

type IntValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntValueContext() *IntValueContext {
	var p = new(IntValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_intValue
	return p
}

func (*IntValueContext) IsIntValueContext() {}

func NewIntValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntValueContext {
	var p = new(IntValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_intValue

	return p
}

func (s *IntValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntValueContext) INT() antlr.TerminalNode {
	return s.GetToken(GraphQLParserINT, 0)
}

func (s *IntValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (p *GraphQLParser) IntValue() (localctx IIntValueContext) {
	this := p
	_ = this

	localctx = NewIntValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GraphQLParserRULE_intValue)

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
		p.SetState(252)
		p.Match(GraphQLParserINT)
	}

	return localctx
}

// IFloatValueContext is an interface to support dynamic dispatch.
type IFloatValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatValueContext differentiates from other interfaces.
	IsFloatValueContext()
}

type FloatValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatValueContext() *FloatValueContext {
	var p = new(FloatValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_floatValue
	return p
}

func (*FloatValueContext) IsFloatValueContext() {}

func NewFloatValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatValueContext {
	var p = new(FloatValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_floatValue

	return p
}

func (s *FloatValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GraphQLParserFLOAT, 0)
}

func (s *FloatValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFloatValue(s)
	}
}

func (s *FloatValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFloatValue(s)
	}
}

func (p *GraphQLParser) FloatValue() (localctx IFloatValueContext) {
	this := p
	_ = this

	localctx = NewFloatValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GraphQLParserRULE_floatValue)

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
		p.SetState(254)
		p.Match(GraphQLParserFLOAT)
	}

	return localctx
}

// IBooleanValueContext is an interface to support dynamic dispatch.
type IBooleanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanValueContext differentiates from other interfaces.
	IsBooleanValueContext()
}

type BooleanValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanValueContext() *BooleanValueContext {
	var p = new(BooleanValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_booleanValue
	return p
}

func (*BooleanValueContext) IsBooleanValueContext() {}

func NewBooleanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueContext {
	var p = new(BooleanValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_booleanValue

	return p
}

func (s *BooleanValueContext) GetParser() antlr.Parser { return s.parser }
func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterBooleanValue(s)
	}
}

func (s *BooleanValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitBooleanValue(s)
	}
}

func (p *GraphQLParser) BooleanValue() (localctx IBooleanValueContext) {
	this := p
	_ = this

	localctx = NewBooleanValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GraphQLParserRULE_booleanValue)
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
		p.SetState(256)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GraphQLParserT__11 || _la == GraphQLParserT__12) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSTRING, 0)
}

func (s *StringValueContext) BLOCK_STRING() antlr.TerminalNode {
	return s.GetToken(GraphQLParserBLOCK_STRING, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *GraphQLParser) StringValue() (localctx IStringValueContext) {
	this := p
	_ = this

	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GraphQLParserRULE_stringValue)
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
		p.SetState(258)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INullValueContext is an interface to support dynamic dispatch.
type INullValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullValueContext differentiates from other interfaces.
	IsNullValueContext()
}

type NullValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullValueContext() *NullValueContext {
	var p = new(NullValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_nullValue
	return p
}

func (*NullValueContext) IsNullValueContext() {}

func NewNullValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullValueContext {
	var p = new(NullValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_nullValue

	return p
}

func (s *NullValueContext) GetParser() antlr.Parser { return s.parser }
func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNullValue(s)
	}
}

func (s *NullValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNullValue(s)
	}
}

func (p *GraphQLParser) NullValue() (localctx INullValueContext) {
	this := p
	_ = this

	localctx = NewNullValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GraphQLParserRULE_nullValue)

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
		p.SetState(260)
		p.Match(GraphQLParserT__13)
	}

	return localctx
}

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValue
	return p
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValue(s)
	}
}

func (p *GraphQLParser) EnumValue() (localctx IEnumValueContext) {
	this := p
	_ = this

	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GraphQLParserRULE_enumValue)

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
		p.SetState(262)
		p.Name()
	}

	return localctx
}

// IListValueContext is an interface to support dynamic dispatch.
type IListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListValueContext differentiates from other interfaces.
	IsListValueContext()
}

type ListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListValueContext() *ListValueContext {
	var p = new(ListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_listValue
	return p
}

func (*ListValueContext) IsListValueContext() {}

func NewListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListValueContext {
	var p = new(ListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_listValue

	return p
}

func (s *ListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ListValueContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ListValueContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterListValue(s)
	}
}

func (s *ListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitListValue(s)
	}
}

func (p *GraphQLParser) ListValue() (localctx IListValueContext) {
	this := p
	_ = this

	localctx = NewListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GraphQLParserRULE_listValue)
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

	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Match(GraphQLParserT__14)
		}
		{
			p.SetState(265)
			p.Match(GraphQLParserT__15)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(GraphQLParserT__14)
		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__3)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__16))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(GraphQLParserNAME-51))|(1<<(GraphQLParserSTRING-51))|(1<<(GraphQLParserBLOCK_STRING-51))|(1<<(GraphQLParserFLOAT-51))|(1<<(GraphQLParserINT-51)))) != 0) {
			{
				p.SetState(267)
				p.Value()
			}

			p.SetState(270)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(272)
			p.Match(GraphQLParserT__15)
		}

	}

	return localctx
}

// IObjectValueContext is an interface to support dynamic dispatch.
type IObjectValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectValueContext differentiates from other interfaces.
	IsObjectValueContext()
}

type ObjectValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectValueContext() *ObjectValueContext {
	var p = new(ObjectValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectValue
	return p
}

func (*ObjectValueContext) IsObjectValueContext() {}

func NewObjectValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectValueContext {
	var p = new(ObjectValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectValue

	return p
}

func (s *ObjectValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectValueContext) AllObjectField() []IObjectFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem())
	var tst = make([]IObjectFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectFieldContext)
		}
	}

	return tst
}

func (s *ObjectValueContext) ObjectField(i int) IObjectFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectFieldContext)
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

func (p *GraphQLParser) ObjectValue() (localctx IObjectValueContext) {
	this := p
	_ = this

	localctx = NewObjectValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GraphQLParserRULE_objectValue)
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
		p.SetState(276)
		p.Match(GraphQLParserT__3)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLParserNAME {
		{
			p.SetState(277)
			p.ObjectField()
		}

		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(283)
		p.Match(GraphQLParserT__4)
	}

	return localctx
}

// IObjectFieldContext is an interface to support dynamic dispatch.
type IObjectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectFieldContext differentiates from other interfaces.
	IsObjectFieldContext()
}

type ObjectFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectFieldContext() *ObjectFieldContext {
	var p = new(ObjectFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectField
	return p
}

func (*ObjectFieldContext) IsObjectFieldContext() {}

func NewObjectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectFieldContext {
	var p = new(ObjectFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectField

	return p
}

func (s *ObjectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectFieldContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectFieldContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ObjectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectField(s)
	}
}

func (s *ObjectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectField(s)
	}
}

func (p *GraphQLParser) ObjectField() (localctx IObjectFieldContext) {
	this := p
	_ = this

	localctx = NewObjectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GraphQLParserRULE_objectField)

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
		p.SetState(285)
		p.Name()
	}
	{
		p.SetState(286)
		p.Match(GraphQLParserT__7)
	}
	{
		p.SetState(287)
		p.Value()
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
	p.RuleIndex = GraphQLParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *GraphQLParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GraphQLParserRULE_variable)

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
		p.SetState(289)
		p.Match(GraphQLParserT__16)
	}
	{
		p.SetState(290)
		p.Name()
	}

	return localctx
}

// IVariableDefinitionsContext is an interface to support dynamic dispatch.
type IVariableDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionsContext differentiates from other interfaces.
	IsVariableDefinitionsContext()
}

type VariableDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionsContext() *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinitions
	return p
}

func (*VariableDefinitionsContext) IsVariableDefinitionsContext() {}

func NewVariableDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinitions

	return p
}

func (s *VariableDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionsContext) AllVariableDefinition() []IVariableDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem())
	var tst = make([]IVariableDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDefinitionContext)
		}
	}

	return tst
}

func (s *VariableDefinitionsContext) VariableDefinition(i int) IVariableDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionContext)
}

func (s *VariableDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinitions(s)
	}
}

func (s *VariableDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinitions(s)
	}
}

func (p *GraphQLParser) VariableDefinitions() (localctx IVariableDefinitionsContext) {
	this := p
	_ = this

	localctx = NewVariableDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GraphQLParserRULE_variableDefinitions)
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
		p.SetState(292)
		p.Match(GraphQLParserT__5)
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__16 {
		{
			p.SetState(293)
			p.VariableDefinition()
		}

		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(298)
		p.Match(GraphQLParserT__6)
	}

	return localctx
}

// IVariableDefinitionContext is an interface to support dynamic dispatch.
type IVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionContext differentiates from other interfaces.
	IsVariableDefinitionContext()
}

type VariableDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionContext() *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinition
	return p
}

func (*VariableDefinitionContext) IsVariableDefinitionContext() {}

func NewVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinition

	return p
}

func (s *VariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableDefinitionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VariableDefinitionContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinition(s)
	}
}

func (p *GraphQLParser) VariableDefinition() (localctx IVariableDefinitionContext) {
	this := p
	_ = this

	localctx = NewVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GraphQLParserRULE_variableDefinition)
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
		p.SetState(300)
		p.Variable()
	}
	{
		p.SetState(301)
		p.Match(GraphQLParserT__7)
	}
	{
		p.SetState(302)
		p.Type_()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__17 {
		{
			p.SetState(303)
			p.DefaultValue()
		}

	}

	return localctx
}

// IDefaultValueContext is an interface to support dynamic dispatch.
type IDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValueContext differentiates from other interfaces.
	IsDefaultValueContext()
}

type DefaultValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueContext() *DefaultValueContext {
	var p = new(DefaultValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_defaultValue
	return p
}

func (*DefaultValueContext) IsDefaultValueContext() {}

func NewDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueContext {
	var p = new(DefaultValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_defaultValue

	return p
}

func (s *DefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefaultValue(s)
	}
}

func (s *DefaultValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefaultValue(s)
	}
}

func (p *GraphQLParser) DefaultValue() (localctx IDefaultValueContext) {
	this := p
	_ = this

	localctx = NewDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GraphQLParserRULE_defaultValue)

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
		p.SetState(306)
		p.Match(GraphQLParserT__17)
	}
	{
		p.SetState(307)
		p.Value()
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
	p.RuleIndex = GraphQLParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *Type_Context) ListType() IListTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *GraphQLParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GraphQLParserRULE_type_)
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

	p.SetState(317)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.NamedType()
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__18 {
			{
				p.SetState(310)
				p.Match(GraphQLParserT__18)
			}

		}

	case GraphQLParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.ListType()
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__18 {
			{
				p.SetState(314)
				p.Match(GraphQLParserT__18)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INamedTypeContext is an interface to support dynamic dispatch.
type INamedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedTypeContext differentiates from other interfaces.
	IsNamedTypeContext()
}

type NamedTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedTypeContext() *NamedTypeContext {
	var p = new(NamedTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_namedType
	return p
}

func (*NamedTypeContext) IsNamedTypeContext() {}

func NewNamedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedTypeContext {
	var p = new(NamedTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_namedType

	return p
}

func (s *NamedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedTypeContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NamedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNamedType(s)
	}
}

func (s *NamedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNamedType(s)
	}
}

func (p *GraphQLParser) NamedType() (localctx INamedTypeContext) {
	this := p
	_ = this

	localctx = NewNamedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GraphQLParserRULE_namedType)

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
		p.SetState(319)
		p.Name()
	}

	return localctx
}

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_listType
	return p
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitListType(s)
	}
}

func (p *GraphQLParser) ListType() (localctx IListTypeContext) {
	this := p
	_ = this

	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GraphQLParserRULE_listType)

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
		p.SetState(321)
		p.Match(GraphQLParserT__14)
	}
	{
		p.SetState(322)
		p.Type_()
	}
	{
		p.SetState(323)
		p.Match(GraphQLParserT__15)
	}

	return localctx
}

// IDirectivesContext is an interface to support dynamic dispatch.
type IDirectivesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectivesContext differentiates from other interfaces.
	IsDirectivesContext()
}

type DirectivesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectivesContext() *DirectivesContext {
	var p = new(DirectivesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directives
	return p
}

func (*DirectivesContext) IsDirectivesContext() {}

func NewDirectivesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectivesContext {
	var p = new(DirectivesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directives

	return p
}

func (s *DirectivesContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectivesContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *DirectivesContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *DirectivesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectivesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectivesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectives(s)
	}
}

func (s *DirectivesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectives(s)
	}
}

func (p *GraphQLParser) Directives() (localctx IDirectivesContext) {
	this := p
	_ = this

	localctx = NewDirectivesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GraphQLParserRULE_directives)
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
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__19 {
		{
			p.SetState(325)
			p.Directive()
		}

		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *GraphQLParser) Directive() (localctx IDirectiveContext) {
	this := p
	_ = this

	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GraphQLParserRULE_directive)
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
		p.SetState(330)
		p.Match(GraphQLParserT__19)
	}
	{
		p.SetState(331)
		p.Name()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(332)
			p.Arguments()
		}

	}

	return localctx
}

// ITypeSystemDefinitionContext is an interface to support dynamic dispatch.
type ITypeSystemDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSystemDefinitionContext differentiates from other interfaces.
	IsTypeSystemDefinitionContext()
}

type TypeSystemDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemDefinitionContext() *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDefinition
	return p
}

func (*TypeSystemDefinitionContext) IsTypeSystemDefinitionContext() {}

func NewTypeSystemDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemDefinition

	return p
}

func (s *TypeSystemDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemDefinitionContext) SchemaDefinition() ISchemaDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaDefinitionContext)
}

func (s *TypeSystemDefinitionContext) TypeDefinition() ITypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefinitionContext)
}

func (s *TypeSystemDefinitionContext) DirectiveDefinition() IDirectiveDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveDefinitionContext)
}

func (s *TypeSystemDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemDefinition(s)
	}
}

func (s *TypeSystemDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemDefinition(s)
	}
}

func (p *GraphQLParser) TypeSystemDefinition() (localctx ITypeSystemDefinitionContext) {
	this := p
	_ = this

	localctx = NewTypeSystemDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GraphQLParserRULE_typeSystemDefinition)

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

	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.SchemaDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.TypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(337)
			p.DirectiveDefinition()
		}

	}

	return localctx
}

// ITypeSystemExtensionContext is an interface to support dynamic dispatch.
type ITypeSystemExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSystemExtensionContext differentiates from other interfaces.
	IsTypeSystemExtensionContext()
}

type TypeSystemExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemExtensionContext() *TypeSystemExtensionContext {
	var p = new(TypeSystemExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemExtension
	return p
}

func (*TypeSystemExtensionContext) IsTypeSystemExtensionContext() {}

func NewTypeSystemExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemExtensionContext {
	var p = new(TypeSystemExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemExtension

	return p
}

func (s *TypeSystemExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemExtensionContext) SchemaExtension() ISchemaExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaExtensionContext)
}

func (s *TypeSystemExtensionContext) TypeExtension() ITypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExtensionContext)
}

func (s *TypeSystemExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemExtension(s)
	}
}

func (s *TypeSystemExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemExtension(s)
	}
}

func (p *GraphQLParser) TypeSystemExtension() (localctx ITypeSystemExtensionContext) {
	this := p
	_ = this

	localctx = NewTypeSystemExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GraphQLParserRULE_typeSystemExtension)

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

	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(340)
			p.SchemaExtension()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(341)
			p.TypeExtension()
		}

	}

	return localctx
}

// ISchemaDefinitionContext is an interface to support dynamic dispatch.
type ISchemaDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaDefinitionContext differentiates from other interfaces.
	IsSchemaDefinitionContext()
}

type SchemaDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaDefinitionContext() *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_schemaDefinition
	return p
}

func (*SchemaDefinitionContext) IsSchemaDefinitionContext() {}

func NewSchemaDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_schemaDefinition

	return p
}

func (s *SchemaDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *SchemaDefinitionContext) AllRootOperationTypeDefinition() []IRootOperationTypeDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRootOperationTypeDefinitionContext)(nil)).Elem())
	var tst = make([]IRootOperationTypeDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRootOperationTypeDefinitionContext)
		}
	}

	return tst
}

func (s *SchemaDefinitionContext) RootOperationTypeDefinition(i int) IRootOperationTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRootOperationTypeDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRootOperationTypeDefinitionContext)
}

func (s *SchemaDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSchemaDefinition(s)
	}
}

func (s *SchemaDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSchemaDefinition(s)
	}
}

func (p *GraphQLParser) SchemaDefinition() (localctx ISchemaDefinitionContext) {
	this := p
	_ = this

	localctx = NewSchemaDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GraphQLParserRULE_schemaDefinition)
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
		p.SetState(344)
		p.Match(GraphQLParserT__20)
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(345)
			p.Directives()
		}

	}
	{
		p.SetState(348)
		p.Match(GraphQLParserT__3)
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2))) != 0) {
		{
			p.SetState(349)
			p.RootOperationTypeDefinition()
		}

		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(354)
		p.Match(GraphQLParserT__4)
	}

	return localctx
}

// IRootOperationTypeDefinitionContext is an interface to support dynamic dispatch.
type IRootOperationTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootOperationTypeDefinitionContext differentiates from other interfaces.
	IsRootOperationTypeDefinitionContext()
}

type RootOperationTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootOperationTypeDefinitionContext() *RootOperationTypeDefinitionContext {
	var p = new(RootOperationTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_rootOperationTypeDefinition
	return p
}

func (*RootOperationTypeDefinitionContext) IsRootOperationTypeDefinitionContext() {}

func NewRootOperationTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootOperationTypeDefinitionContext {
	var p = new(RootOperationTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_rootOperationTypeDefinition

	return p
}

func (s *RootOperationTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *RootOperationTypeDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *RootOperationTypeDefinitionContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *RootOperationTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootOperationTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootOperationTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterRootOperationTypeDefinition(s)
	}
}

func (s *RootOperationTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitRootOperationTypeDefinition(s)
	}
}

func (p *GraphQLParser) RootOperationTypeDefinition() (localctx IRootOperationTypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewRootOperationTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GraphQLParserRULE_rootOperationTypeDefinition)

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
		p.SetState(356)
		p.OperationType()
	}
	{
		p.SetState(357)
		p.Match(GraphQLParserT__7)
	}
	{
		p.SetState(358)
		p.NamedType()
	}

	return localctx
}

// ISchemaExtensionContext is an interface to support dynamic dispatch.
type ISchemaExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaExtensionContext differentiates from other interfaces.
	IsSchemaExtensionContext()
}

type SchemaExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaExtensionContext() *SchemaExtensionContext {
	var p = new(SchemaExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_schemaExtension
	return p
}

func (*SchemaExtensionContext) IsSchemaExtensionContext() {}

func NewSchemaExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaExtensionContext {
	var p = new(SchemaExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_schemaExtension

	return p
}

func (s *SchemaExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *SchemaExtensionContext) AllOperationTypeDefinition() []IOperationTypeDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperationTypeDefinitionContext)(nil)).Elem())
	var tst = make([]IOperationTypeDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperationTypeDefinitionContext)
		}
	}

	return tst
}

func (s *SchemaExtensionContext) OperationTypeDefinition(i int) IOperationTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeDefinitionContext)
}

func (s *SchemaExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSchemaExtension(s)
	}
}

func (s *SchemaExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSchemaExtension(s)
	}
}

func (p *GraphQLParser) SchemaExtension() (localctx ISchemaExtensionContext) {
	this := p
	_ = this

	localctx = NewSchemaExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GraphQLParserRULE_schemaExtension)
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

	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(361)
			p.Match(GraphQLParserT__20)
		}
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(362)
				p.Directives()
			}

		}
		{
			p.SetState(365)
			p.Match(GraphQLParserT__3)
		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2))) != 0) {
			{
				p.SetState(366)
				p.OperationTypeDefinition()
			}

			p.SetState(369)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(371)
			p.Match(GraphQLParserT__4)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(374)
			p.Match(GraphQLParserT__20)
		}
		{
			p.SetState(375)
			p.Directives()
		}

	}

	return localctx
}

// IOperationTypeDefinitionContext is an interface to support dynamic dispatch.
type IOperationTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeDefinitionContext differentiates from other interfaces.
	IsOperationTypeDefinitionContext()
}

type OperationTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeDefinitionContext() *OperationTypeDefinitionContext {
	var p = new(OperationTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationTypeDefinition
	return p
}

func (*OperationTypeDefinitionContext) IsOperationTypeDefinitionContext() {}

func NewOperationTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeDefinitionContext {
	var p = new(OperationTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationTypeDefinition

	return p
}

func (s *OperationTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationTypeDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationTypeDefinitionContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *OperationTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationTypeDefinition(s)
	}
}

func (s *OperationTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationTypeDefinition(s)
	}
}

func (p *GraphQLParser) OperationTypeDefinition() (localctx IOperationTypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewOperationTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GraphQLParserRULE_operationTypeDefinition)

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
		p.SetState(378)
		p.OperationType()
	}
	{
		p.SetState(379)
		p.Match(GraphQLParserT__7)
	}
	{
		p.SetState(380)
		p.NamedType()
	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (p *GraphQLParser) Description() (localctx IDescriptionContext) {
	this := p
	_ = this

	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GraphQLParserRULE_description)

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
		p.SetState(382)
		p.StringValue()
	}

	return localctx
}

// ITypeDefinitionContext is an interface to support dynamic dispatch.
type ITypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefinitionContext differentiates from other interfaces.
	IsTypeDefinitionContext()
}

type TypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefinitionContext() *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeDefinition
	return p
}

func (*TypeDefinitionContext) IsTypeDefinitionContext() {}

func NewTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeDefinition

	return p
}

func (s *TypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefinitionContext) ScalarTypeDefinition() IScalarTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarTypeDefinitionContext)
}

func (s *TypeDefinitionContext) ObjectTypeDefinition() IObjectTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectTypeDefinitionContext)
}

func (s *TypeDefinitionContext) InterfaceTypeDefinition() IInterfaceTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeDefinitionContext)
}

func (s *TypeDefinitionContext) UnionTypeDefinition() IUnionTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeDefinitionContext)
}

func (s *TypeDefinitionContext) EnumTypeDefinition() IEnumTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeDefinitionContext)
}

func (s *TypeDefinitionContext) InputObjectTypeDefinition() IInputObjectTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeDefinitionContext)
}

func (s *TypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeDefinition(s)
	}
}

func (p *GraphQLParser) TypeDefinition() (localctx ITypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GraphQLParserRULE_typeDefinition)

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

	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)
			p.ScalarTypeDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.ObjectTypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(386)
			p.InterfaceTypeDefinition()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(387)
			p.UnionTypeDefinition()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(388)
			p.EnumTypeDefinition()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(389)
			p.InputObjectTypeDefinition()
		}

	}

	return localctx
}

// ITypeExtensionContext is an interface to support dynamic dispatch.
type ITypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExtensionContext differentiates from other interfaces.
	IsTypeExtensionContext()
}

type TypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExtensionContext() *TypeExtensionContext {
	var p = new(TypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeExtension
	return p
}

func (*TypeExtensionContext) IsTypeExtensionContext() {}

func NewTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExtensionContext {
	var p = new(TypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeExtension

	return p
}

func (s *TypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExtensionContext) ScalarTypeExtension() IScalarTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarTypeExtensionContext)
}

func (s *TypeExtensionContext) ObjectTypeExtension() IObjectTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectTypeExtensionContext)
}

func (s *TypeExtensionContext) InterfaceTypeExtension() IInterfaceTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeExtensionContext)
}

func (s *TypeExtensionContext) UnionTypeExtension() IUnionTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeExtensionContext)
}

func (s *TypeExtensionContext) EnumTypeExtension() IEnumTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeExtensionContext)
}

func (s *TypeExtensionContext) InputObjectTypeExtension() IInputObjectTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeExtensionContext)
}

func (s *TypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeExtension(s)
	}
}

func (s *TypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeExtension(s)
	}
}

func (p *GraphQLParser) TypeExtension() (localctx ITypeExtensionContext) {
	this := p
	_ = this

	localctx = NewTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GraphQLParserRULE_typeExtension)

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

	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.ScalarTypeExtension()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.ObjectTypeExtension()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(394)
			p.InterfaceTypeExtension()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(395)
			p.UnionTypeExtension()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(396)
			p.EnumTypeExtension()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(397)
			p.InputObjectTypeExtension()
		}

	}

	return localctx
}

// IScalarTypeDefinitionContext is an interface to support dynamic dispatch.
type IScalarTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarTypeDefinitionContext differentiates from other interfaces.
	IsScalarTypeDefinitionContext()
}

type ScalarTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeDefinitionContext() *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeDefinition
	return p
}

func (*ScalarTypeDefinitionContext) IsScalarTypeDefinitionContext() {}

func NewScalarTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_scalarTypeDefinition

	return p
}

func (s *ScalarTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ScalarTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterScalarTypeDefinition(s)
	}
}

func (s *ScalarTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitScalarTypeDefinition(s)
	}
}

func (p *GraphQLParser) ScalarTypeDefinition() (localctx IScalarTypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewScalarTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GraphQLParserRULE_scalarTypeDefinition)
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
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(400)
			p.Description()
		}

	}
	{
		p.SetState(403)
		p.Match(GraphQLParserT__22)
	}
	{
		p.SetState(404)
		p.Name()
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(405)
			p.Directives()
		}

	}

	return localctx
}

// IScalarTypeExtensionContext is an interface to support dynamic dispatch.
type IScalarTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarTypeExtensionContext differentiates from other interfaces.
	IsScalarTypeExtensionContext()
}

type ScalarTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeExtensionContext() *ScalarTypeExtensionContext {
	var p = new(ScalarTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeExtension
	return p
}

func (*ScalarTypeExtensionContext) IsScalarTypeExtensionContext() {}

func NewScalarTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeExtensionContext {
	var p = new(ScalarTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_scalarTypeExtension

	return p
}

func (s *ScalarTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterScalarTypeExtension(s)
	}
}

func (s *ScalarTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitScalarTypeExtension(s)
	}
}

func (p *GraphQLParser) ScalarTypeExtension() (localctx IScalarTypeExtensionContext) {
	this := p
	_ = this

	localctx = NewScalarTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GraphQLParserRULE_scalarTypeExtension)

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
		p.SetState(408)
		p.Match(GraphQLParserT__21)
	}
	{
		p.SetState(409)
		p.Match(GraphQLParserT__22)
	}
	{
		p.SetState(410)
		p.Name()
	}
	{
		p.SetState(411)
		p.Directives()
	}

	return localctx
}

// IObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectTypeDefinitionContext differentiates from other interfaces.
	IsObjectTypeDefinitionContext()
}

type ObjectTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeDefinitionContext() *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeDefinition
	return p
}

func (*ObjectTypeDefinitionContext) IsObjectTypeDefinitionContext() {}

func NewObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectTypeDefinition

	return p
}

func (s *ObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ObjectTypeDefinitionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectTypeDefinition(s)
	}
}

func (s *ObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectTypeDefinition(s)
	}
}

func (p *GraphQLParser) ObjectTypeDefinition() (localctx IObjectTypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GraphQLParserRULE_objectTypeDefinition)
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
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(413)
			p.Description()
		}

	}
	{
		p.SetState(416)
		p.Match(GraphQLParserT__23)
	}
	{
		p.SetState(417)
		p.Name()
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__24 {
		{
			p.SetState(418)
			p.implementsInterfaces(0)
		}

	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(421)
			p.Directives()
		}

	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(424)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IImplementsInterfacesContext is an interface to support dynamic dispatch.
type IImplementsInterfacesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImplementsInterfacesContext differentiates from other interfaces.
	IsImplementsInterfacesContext()
}

type ImplementsInterfacesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplementsInterfacesContext() *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_implementsInterfaces
	return p
}

func (*ImplementsInterfacesContext) IsImplementsInterfacesContext() {}

func NewImplementsInterfacesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_implementsInterfaces

	return p
}

func (s *ImplementsInterfacesContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplementsInterfacesContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *ImplementsInterfacesContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ImplementsInterfacesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementsInterfacesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplementsInterfacesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterImplementsInterfaces(s)
	}
}

func (s *ImplementsInterfacesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitImplementsInterfaces(s)
	}
}

func (p *GraphQLParser) ImplementsInterfaces() (localctx IImplementsInterfacesContext) {
	return p.implementsInterfaces(0)
}

func (p *GraphQLParser) implementsInterfaces(_p int) (localctx IImplementsInterfacesContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewImplementsInterfacesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IImplementsInterfacesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 94
	p.EnterRecursionRule(localctx, 94, GraphQLParserRULE_implementsInterfaces, _p)
	var _la int

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
	{
		p.SetState(428)
		p.Match(GraphQLParserT__24)
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__25 {
		{
			p.SetState(429)
			p.Match(GraphQLParserT__25)
		}

	}
	{
		p.SetState(432)
		p.NamedType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewImplementsInterfacesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GraphQLParserRULE_implementsInterfaces)
			p.SetState(434)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(435)
				p.Match(GraphQLParserT__25)
			}
			{
				p.SetState(436)
				p.NamedType()
			}

		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
	}

	return localctx
}

// IFieldsDefinitionContext is an interface to support dynamic dispatch.
type IFieldsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsDefinitionContext differentiates from other interfaces.
	IsFieldsDefinitionContext()
}

type FieldsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsDefinitionContext() *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldsDefinition
	return p
}

func (*FieldsDefinitionContext) IsFieldsDefinitionContext() {}

func NewFieldsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fieldsDefinition

	return p
}

func (s *FieldsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsDefinitionContext) AllFieldDefinition() []IFieldDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem())
	var tst = make([]IFieldDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldDefinitionContext)
		}
	}

	return tst
}

func (s *FieldsDefinitionContext) FieldDefinition(i int) IFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldDefinitionContext)
}

func (s *FieldsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFieldsDefinition(s)
	}
}

func (s *FieldsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFieldsDefinition(s)
	}
}

func (p *GraphQLParser) FieldsDefinition() (localctx IFieldsDefinitionContext) {
	this := p
	_ = this

	localctx = NewFieldsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, GraphQLParserRULE_fieldsDefinition)
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
		p.SetState(442)
		p.Match(GraphQLParserT__3)
	}
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(GraphQLParserNAME-51))|(1<<(GraphQLParserSTRING-51))|(1<<(GraphQLParserBLOCK_STRING-51)))) != 0) {
		{
			p.SetState(443)
			p.FieldDefinition()
		}

		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(448)
		p.Match(GraphQLParserT__4)
	}

	return localctx
}

// IFieldDefinitionContext is an interface to support dynamic dispatch.
type IFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDefinitionContext differentiates from other interfaces.
	IsFieldDefinitionContext()
}

type FieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefinitionContext() *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldDefinition
	return p
}

func (*FieldDefinitionContext) IsFieldDefinitionContext() {}

func NewFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fieldDefinition

	return p
}

func (s *FieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldDefinitionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *FieldDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *FieldDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFieldDefinition(s)
	}
}

func (s *FieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFieldDefinition(s)
	}
}

func (p *GraphQLParser) FieldDefinition() (localctx IFieldDefinitionContext) {
	this := p
	_ = this

	localctx = NewFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, GraphQLParserRULE_fieldDefinition)
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
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(450)
			p.Description()
		}

	}
	{
		p.SetState(453)
		p.Name()
	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(454)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(457)
		p.Match(GraphQLParserT__7)
	}
	{
		p.SetState(458)
		p.Type_()
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(459)
			p.Directives()
		}

	}

	return localctx
}

// IArgumentsDefinitionContext is an interface to support dynamic dispatch.
type IArgumentsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsDefinitionContext differentiates from other interfaces.
	IsArgumentsDefinitionContext()
}

type ArgumentsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsDefinitionContext() *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_argumentsDefinition
	return p
}

func (*ArgumentsDefinitionContext) IsArgumentsDefinitionContext() {}

func NewArgumentsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_argumentsDefinition

	return p
}

func (s *ArgumentsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsDefinitionContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem())
	var tst = make([]IInputValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputValueDefinitionContext)
		}
	}

	return tst
}

func (s *ArgumentsDefinitionContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInputValueDefinitionContext)
}

func (s *ArgumentsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArgumentsDefinition(s)
	}
}

func (s *ArgumentsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArgumentsDefinition(s)
	}
}

func (p *GraphQLParser) ArgumentsDefinition() (localctx IArgumentsDefinitionContext) {
	this := p
	_ = this

	localctx = NewArgumentsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, GraphQLParserRULE_argumentsDefinition)
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
		p.SetState(462)
		p.Match(GraphQLParserT__5)
	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(GraphQLParserNAME-51))|(1<<(GraphQLParserSTRING-51))|(1<<(GraphQLParserBLOCK_STRING-51)))) != 0) {
		{
			p.SetState(463)
			p.InputValueDefinition()
		}

		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(468)
		p.Match(GraphQLParserT__6)
	}

	return localctx
}

// IInputValueDefinitionContext is an interface to support dynamic dispatch.
type IInputValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputValueDefinitionContext differentiates from other interfaces.
	IsInputValueDefinitionContext()
}

type InputValueDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputValueDefinitionContext() *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputValueDefinition
	return p
}

func (*InputValueDefinitionContext) IsInputValueDefinitionContext() {}

func NewInputValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputValueDefinition

	return p
}

func (s *InputValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputValueDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputValueDefinitionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *InputValueDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputValueDefinitionContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *InputValueDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputValueDefinition(s)
	}
}

func (s *InputValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputValueDefinition(s)
	}
}

func (p *GraphQLParser) InputValueDefinition() (localctx IInputValueDefinitionContext) {
	this := p
	_ = this

	localctx = NewInputValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, GraphQLParserRULE_inputValueDefinition)
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
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(470)
			p.Description()
		}

	}
	{
		p.SetState(473)
		p.Name()
	}
	{
		p.SetState(474)
		p.Match(GraphQLParserT__7)
	}
	{
		p.SetState(475)
		p.Type_()
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__17 {
		{
			p.SetState(476)
			p.DefaultValue()
		}

	}
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(479)
			p.Directives()
		}

	}

	return localctx
}

// IObjectTypeExtensionContext is an interface to support dynamic dispatch.
type IObjectTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectTypeExtensionContext differentiates from other interfaces.
	IsObjectTypeExtensionContext()
}

type ObjectTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeExtensionContext() *ObjectTypeExtensionContext {
	var p = new(ObjectTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeExtension
	return p
}

func (*ObjectTypeExtensionContext) IsObjectTypeExtensionContext() {}

func NewObjectTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeExtensionContext {
	var p = new(ObjectTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectTypeExtension

	return p
}

func (s *ObjectTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeExtensionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeExtensionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectTypeExtension(s)
	}
}

func (s *ObjectTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectTypeExtension(s)
	}
}

func (p *GraphQLParser) ObjectTypeExtension() (localctx IObjectTypeExtensionContext) {
	this := p
	_ = this

	localctx = NewObjectTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, GraphQLParserRULE_objectTypeExtension)
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

	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(482)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(483)
			p.Match(GraphQLParserT__23)
		}
		{
			p.SetState(484)
			p.Name()
		}
		p.SetState(486)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__24 {
			{
				p.SetState(485)
				p.implementsInterfaces(0)
			}

		}
		p.SetState(489)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(488)
				p.Directives()
			}

		}
		{
			p.SetState(491)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(493)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(494)
			p.Match(GraphQLParserT__23)
		}
		{
			p.SetState(495)
			p.Name()
		}
		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__24 {
			{
				p.SetState(496)
				p.implementsInterfaces(0)
			}

		}
		{
			p.SetState(499)
			p.Directives()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(501)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(502)
			p.Match(GraphQLParserT__23)
		}
		{
			p.SetState(503)
			p.Name()
		}
		{
			p.SetState(504)
			p.implementsInterfaces(0)
		}

	}

	return localctx
}

// IInterfaceTypeDefinitionContext is an interface to support dynamic dispatch.
type IInterfaceTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceTypeDefinitionContext differentiates from other interfaces.
	IsInterfaceTypeDefinitionContext()
}

type InterfaceTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeDefinitionContext() *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeDefinition
	return p
}

func (*InterfaceTypeDefinitionContext) IsInterfaceTypeDefinitionContext() {}

func NewInterfaceTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_interfaceTypeDefinition

	return p
}

func (s *InterfaceTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InterfaceTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInterfaceTypeDefinition(s)
	}
}

func (s *InterfaceTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInterfaceTypeDefinition(s)
	}
}

func (p *GraphQLParser) InterfaceTypeDefinition() (localctx IInterfaceTypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewInterfaceTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, GraphQLParserRULE_interfaceTypeDefinition)
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
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(508)
			p.Description()
		}

	}
	{
		p.SetState(511)
		p.Match(GraphQLParserT__26)
	}
	{
		p.SetState(512)
		p.Name()
	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(513)
			p.Directives()
		}

	}
	p.SetState(517)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(516)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IInterfaceTypeExtensionContext is an interface to support dynamic dispatch.
type IInterfaceTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceTypeExtensionContext differentiates from other interfaces.
	IsInterfaceTypeExtensionContext()
}

type InterfaceTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeExtensionContext() *InterfaceTypeExtensionContext {
	var p = new(InterfaceTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeExtension
	return p
}

func (*InterfaceTypeExtensionContext) IsInterfaceTypeExtensionContext() {}

func NewInterfaceTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeExtensionContext {
	var p = new(InterfaceTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_interfaceTypeExtension

	return p
}

func (s *InterfaceTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeExtensionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInterfaceTypeExtension(s)
	}
}

func (s *InterfaceTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInterfaceTypeExtension(s)
	}
}

func (p *GraphQLParser) InterfaceTypeExtension() (localctx IInterfaceTypeExtensionContext) {
	this := p
	_ = this

	localctx = NewInterfaceTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, GraphQLParserRULE_interfaceTypeExtension)
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

	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(519)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(520)
			p.Match(GraphQLParserT__26)
		}
		{
			p.SetState(521)
			p.Name()
		}
		p.SetState(523)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(522)
				p.Directives()
			}

		}
		{
			p.SetState(525)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(527)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(528)
			p.Match(GraphQLParserT__26)
		}
		{
			p.SetState(529)
			p.Name()
		}
		{
			p.SetState(530)
			p.Directives()
		}

	}

	return localctx
}

// IUnionTypeDefinitionContext is an interface to support dynamic dispatch.
type IUnionTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeDefinitionContext differentiates from other interfaces.
	IsUnionTypeDefinitionContext()
}

type UnionTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeDefinitionContext() *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeDefinition
	return p
}

func (*UnionTypeDefinitionContext) IsUnionTypeDefinitionContext() {}

func NewUnionTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionTypeDefinition

	return p
}

func (s *UnionTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *UnionTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeDefinitionContext) UnionMemberTypes() IUnionMemberTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMemberTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMemberTypesContext)
}

func (s *UnionTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionTypeDefinition(s)
	}
}

func (s *UnionTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionTypeDefinition(s)
	}
}

func (p *GraphQLParser) UnionTypeDefinition() (localctx IUnionTypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewUnionTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, GraphQLParserRULE_unionTypeDefinition)
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
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(534)
			p.Description()
		}

	}
	{
		p.SetState(537)
		p.Match(GraphQLParserT__27)
	}
	{
		p.SetState(538)
		p.Name()
	}
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(539)
			p.Directives()
		}

	}
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__17 {
		{
			p.SetState(542)
			p.UnionMemberTypes()
		}

	}

	return localctx
}

// IUnionMemberTypesContext is an interface to support dynamic dispatch.
type IUnionMemberTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionMemberTypesContext differentiates from other interfaces.
	IsUnionMemberTypesContext()
}

type UnionMemberTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionMemberTypesContext() *UnionMemberTypesContext {
	var p = new(UnionMemberTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionMemberTypes
	return p
}

func (*UnionMemberTypesContext) IsUnionMemberTypesContext() {}

func NewUnionMemberTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionMemberTypesContext {
	var p = new(UnionMemberTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionMemberTypes

	return p
}

func (s *UnionMemberTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionMemberTypesContext) AllNamedType() []INamedTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INamedTypeContext)(nil)).Elem())
	var tst = make([]INamedTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INamedTypeContext)
		}
	}

	return tst
}

func (s *UnionMemberTypesContext) NamedType(i int) INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *UnionMemberTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionMemberTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionMemberTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionMemberTypes(s)
	}
}

func (s *UnionMemberTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionMemberTypes(s)
	}
}

func (p *GraphQLParser) UnionMemberTypes() (localctx IUnionMemberTypesContext) {
	this := p
	_ = this

	localctx = NewUnionMemberTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, GraphQLParserRULE_unionMemberTypes)
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
		p.Match(GraphQLParserT__17)
	}
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__28 {
		{
			p.SetState(546)
			p.Match(GraphQLParserT__28)
		}

	}
	{
		p.SetState(549)
		p.NamedType()
	}
	p.SetState(554)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLParserT__28 {
		{
			p.SetState(550)
			p.Match(GraphQLParserT__28)
		}
		{
			p.SetState(551)
			p.NamedType()
		}

		p.SetState(556)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnionTypeExtensionContext is an interface to support dynamic dispatch.
type IUnionTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeExtensionContext differentiates from other interfaces.
	IsUnionTypeExtensionContext()
}

type UnionTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeExtensionContext() *UnionTypeExtensionContext {
	var p = new(UnionTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeExtension
	return p
}

func (*UnionTypeExtensionContext) IsUnionTypeExtensionContext() {}

func NewUnionTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeExtensionContext {
	var p = new(UnionTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionTypeExtension

	return p
}

func (s *UnionTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeExtensionContext) UnionMemberTypes() IUnionMemberTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMemberTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMemberTypesContext)
}

func (s *UnionTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionTypeExtension(s)
	}
}

func (s *UnionTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionTypeExtension(s)
	}
}

func (p *GraphQLParser) UnionTypeExtension() (localctx IUnionTypeExtensionContext) {
	this := p
	_ = this

	localctx = NewUnionTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, GraphQLParserRULE_unionTypeExtension)
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

	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(557)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(558)
			p.Match(GraphQLParserT__27)
		}
		{
			p.SetState(559)
			p.Name()
		}
		p.SetState(561)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(560)
				p.Directives()
			}

		}
		{
			p.SetState(563)
			p.UnionMemberTypes()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(565)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(566)
			p.Match(GraphQLParserT__27)
		}
		{
			p.SetState(567)
			p.Name()
		}
		{
			p.SetState(568)
			p.Directives()
		}

	}

	return localctx
}

// IEnumTypeDefinitionContext is an interface to support dynamic dispatch.
type IEnumTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeDefinitionContext differentiates from other interfaces.
	IsEnumTypeDefinitionContext()
}

type EnumTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeDefinitionContext() *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeDefinition
	return p
}

func (*EnumTypeDefinitionContext) IsEnumTypeDefinitionContext() {}

func NewEnumTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumTypeDefinition

	return p
}

func (s *EnumTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeDefinitionContext) EnumValuesDefinition() IEnumValuesDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValuesDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValuesDefinitionContext)
}

func (s *EnumTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumTypeDefinition(s)
	}
}

func (s *EnumTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumTypeDefinition(s)
	}
}

func (p *GraphQLParser) EnumTypeDefinition() (localctx IEnumTypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewEnumTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, GraphQLParserRULE_enumTypeDefinition)
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
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(572)
			p.Description()
		}

	}
	{
		p.SetState(575)
		p.Match(GraphQLParserT__29)
	}
	{
		p.SetState(576)
		p.Name()
	}
	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(577)
			p.Directives()
		}

	}
	p.SetState(581)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(580)
			p.EnumValuesDefinition()
		}

	}

	return localctx
}

// IEnumValuesDefinitionContext is an interface to support dynamic dispatch.
type IEnumValuesDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValuesDefinitionContext differentiates from other interfaces.
	IsEnumValuesDefinitionContext()
}

type EnumValuesDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValuesDefinitionContext() *EnumValuesDefinitionContext {
	var p = new(EnumValuesDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValuesDefinition
	return p
}

func (*EnumValuesDefinitionContext) IsEnumValuesDefinitionContext() {}

func NewEnumValuesDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValuesDefinitionContext {
	var p = new(EnumValuesDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValuesDefinition

	return p
}

func (s *EnumValuesDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValuesDefinitionContext) AllEnumValueDefinition() []IEnumValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueDefinitionContext)(nil)).Elem())
	var tst = make([]IEnumValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueDefinitionContext)
		}
	}

	return tst
}

func (s *EnumValuesDefinitionContext) EnumValueDefinition(i int) IEnumValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueDefinitionContext)
}

func (s *EnumValuesDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValuesDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValuesDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValuesDefinition(s)
	}
}

func (s *EnumValuesDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValuesDefinition(s)
	}
}

func (p *GraphQLParser) EnumValuesDefinition() (localctx IEnumValuesDefinitionContext) {
	this := p
	_ = this

	localctx = NewEnumValuesDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, GraphQLParserRULE_enumValuesDefinition)
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
		p.SetState(583)
		p.Match(GraphQLParserT__3)
	}
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(GraphQLParserNAME-51))|(1<<(GraphQLParserSTRING-51))|(1<<(GraphQLParserBLOCK_STRING-51)))) != 0) {
		{
			p.SetState(584)
			p.EnumValueDefinition()
		}

		p.SetState(587)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(589)
		p.Match(GraphQLParserT__4)
	}

	return localctx
}

// IEnumValueDefinitionContext is an interface to support dynamic dispatch.
type IEnumValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueDefinitionContext differentiates from other interfaces.
	IsEnumValueDefinitionContext()
}

type EnumValueDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueDefinitionContext() *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValueDefinition
	return p
}

func (*EnumValueDefinitionContext) IsEnumValueDefinitionContext() {}

func NewEnumValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValueDefinition

	return p
}

func (s *EnumValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueDefinitionContext) EnumValue() IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumValueDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumValueDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValueDefinition(s)
	}
}

func (s *EnumValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValueDefinition(s)
	}
}

func (p *GraphQLParser) EnumValueDefinition() (localctx IEnumValueDefinitionContext) {
	this := p
	_ = this

	localctx = NewEnumValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, GraphQLParserRULE_enumValueDefinition)
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
	p.SetState(592)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(591)
			p.Description()
		}

	}
	{
		p.SetState(594)
		p.EnumValue()
	}
	p.SetState(596)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(595)
			p.Directives()
		}

	}

	return localctx
}

// IEnumTypeExtensionContext is an interface to support dynamic dispatch.
type IEnumTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeExtensionContext differentiates from other interfaces.
	IsEnumTypeExtensionContext()
}

type EnumTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeExtensionContext() *EnumTypeExtensionContext {
	var p = new(EnumTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeExtension
	return p
}

func (*EnumTypeExtensionContext) IsEnumTypeExtensionContext() {}

func NewEnumTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeExtensionContext {
	var p = new(EnumTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumTypeExtension

	return p
}

func (s *EnumTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeExtensionContext) EnumValuesDefinition() IEnumValuesDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValuesDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValuesDefinitionContext)
}

func (s *EnumTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumTypeExtension(s)
	}
}

func (s *EnumTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumTypeExtension(s)
	}
}

func (p *GraphQLParser) EnumTypeExtension() (localctx IEnumTypeExtensionContext) {
	this := p
	_ = this

	localctx = NewEnumTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, GraphQLParserRULE_enumTypeExtension)
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

	p.SetState(611)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 77, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(598)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(599)
			p.Match(GraphQLParserT__29)
		}
		{
			p.SetState(600)
			p.Name()
		}
		p.SetState(602)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(601)
				p.Directives()
			}

		}
		{
			p.SetState(604)
			p.EnumValuesDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(606)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(607)
			p.Match(GraphQLParserT__29)
		}
		{
			p.SetState(608)
			p.Name()
		}
		{
			p.SetState(609)
			p.Directives()
		}

	}

	return localctx
}

// IInputObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IInputObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputObjectTypeDefinitionContext differentiates from other interfaces.
	IsInputObjectTypeDefinitionContext()
}

type InputObjectTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeDefinitionContext() *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeDefinition
	return p
}

func (*InputObjectTypeDefinitionContext) IsInputObjectTypeDefinitionContext() {}

func NewInputObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeDefinition

	return p
}

func (s *InputObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeDefinitionContext) InputFieldsDefinition() IInputFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputFieldsDefinitionContext)
}

func (s *InputObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputObjectTypeDefinition(s)
	}
}

func (s *InputObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputObjectTypeDefinition(s)
	}
}

func (p *GraphQLParser) InputObjectTypeDefinition() (localctx IInputObjectTypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewInputObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, GraphQLParserRULE_inputObjectTypeDefinition)
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
	p.SetState(614)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(613)
			p.Description()
		}

	}
	{
		p.SetState(616)
		p.Match(GraphQLParserT__30)
	}
	{
		p.SetState(617)
		p.Name()
	}
	p.SetState(619)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(618)
			p.Directives()
		}

	}
	p.SetState(622)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 80, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(621)
			p.InputFieldsDefinition()
		}

	}

	return localctx
}

// IInputFieldsDefinitionContext is an interface to support dynamic dispatch.
type IInputFieldsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputFieldsDefinitionContext differentiates from other interfaces.
	IsInputFieldsDefinitionContext()
}

type InputFieldsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputFieldsDefinitionContext() *InputFieldsDefinitionContext {
	var p = new(InputFieldsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputFieldsDefinition
	return p
}

func (*InputFieldsDefinitionContext) IsInputFieldsDefinitionContext() {}

func NewInputFieldsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputFieldsDefinitionContext {
	var p = new(InputFieldsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputFieldsDefinition

	return p
}

func (s *InputFieldsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputFieldsDefinitionContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem())
	var tst = make([]IInputValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputValueDefinitionContext)
		}
	}

	return tst
}

func (s *InputFieldsDefinitionContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInputValueDefinitionContext)
}

func (s *InputFieldsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputFieldsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputFieldsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputFieldsDefinition(s)
	}
}

func (s *InputFieldsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputFieldsDefinition(s)
	}
}

func (p *GraphQLParser) InputFieldsDefinition() (localctx IInputFieldsDefinitionContext) {
	this := p
	_ = this

	localctx = NewInputFieldsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, GraphQLParserRULE_inputFieldsDefinition)
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
		p.SetState(624)
		p.Match(GraphQLParserT__3)
	}
	p.SetState(626)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(GraphQLParserNAME-51))|(1<<(GraphQLParserSTRING-51))|(1<<(GraphQLParserBLOCK_STRING-51)))) != 0) {
		{
			p.SetState(625)
			p.InputValueDefinition()
		}

		p.SetState(628)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(630)
		p.Match(GraphQLParserT__4)
	}

	return localctx
}

// IInputObjectTypeExtensionContext is an interface to support dynamic dispatch.
type IInputObjectTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputObjectTypeExtensionContext differentiates from other interfaces.
	IsInputObjectTypeExtensionContext()
}

type InputObjectTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeExtensionContext() *InputObjectTypeExtensionContext {
	var p = new(InputObjectTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeExtension
	return p
}

func (*InputObjectTypeExtensionContext) IsInputObjectTypeExtensionContext() {}

func NewInputObjectTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeExtensionContext {
	var p = new(InputObjectTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeExtension

	return p
}

func (s *InputObjectTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeExtensionContext) InputFieldsDefinition() IInputFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputFieldsDefinitionContext)
}

func (s *InputObjectTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputObjectTypeExtension(s)
	}
}

func (s *InputObjectTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputObjectTypeExtension(s)
	}
}

func (p *GraphQLParser) InputObjectTypeExtension() (localctx IInputObjectTypeExtensionContext) {
	this := p
	_ = this

	localctx = NewInputObjectTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, GraphQLParserRULE_inputObjectTypeExtension)
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

	p.SetState(645)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(632)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(633)
			p.Match(GraphQLParserT__30)
		}
		{
			p.SetState(634)
			p.Name()
		}
		p.SetState(636)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(635)
				p.Directives()
			}

		}
		{
			p.SetState(638)
			p.InputFieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(640)
			p.Match(GraphQLParserT__21)
		}
		{
			p.SetState(641)
			p.Match(GraphQLParserT__30)
		}
		{
			p.SetState(642)
			p.Name()
		}
		{
			p.SetState(643)
			p.Directives()
		}

	}

	return localctx
}

// IDirectiveDefinitionContext is an interface to support dynamic dispatch.
type IDirectiveDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveDefinitionContext differentiates from other interfaces.
	IsDirectiveDefinitionContext()
}

type DirectiveDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveDefinitionContext() *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveDefinition
	return p
}

func (*DirectiveDefinitionContext) IsDirectiveDefinitionContext() {}

func NewDirectiveDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveDefinition

	return p
}

func (s *DirectiveDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveDefinitionContext) DirectiveLocations() IDirectiveLocationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveLocationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationsContext)
}

func (s *DirectiveDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *DirectiveDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *DirectiveDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveDefinition(s)
	}
}

func (s *DirectiveDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveDefinition(s)
	}
}

func (p *GraphQLParser) DirectiveDefinition() (localctx IDirectiveDefinitionContext) {
	this := p
	_ = this

	localctx = NewDirectiveDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, GraphQLParserRULE_directiveDefinition)
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

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(647)
			p.Description()
		}

	}
	{
		p.SetState(650)
		p.Match(GraphQLParserT__31)
	}
	{
		p.SetState(651)
		p.Match(GraphQLParserT__19)
	}
	{
		p.SetState(652)
		p.Name()
	}
	p.SetState(654)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(653)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(656)
		p.Match(GraphQLParserT__10)
	}
	{
		p.SetState(657)
		p.DirectiveLocations()
	}

	return localctx
}

// IDirectiveLocationsContext is an interface to support dynamic dispatch.
type IDirectiveLocationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveLocationsContext differentiates from other interfaces.
	IsDirectiveLocationsContext()
}

type DirectiveLocationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveLocationsContext() *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveLocations
	return p
}

func (*DirectiveLocationsContext) IsDirectiveLocationsContext() {}

func NewDirectiveLocationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveLocations

	return p
}

func (s *DirectiveLocationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveLocationsContext) AllDirectiveLocation() []IDirectiveLocationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveLocationContext)(nil)).Elem())
	var tst = make([]IDirectiveLocationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveLocationContext)
		}
	}

	return tst
}

func (s *DirectiveLocationsContext) DirectiveLocation(i int) IDirectiveLocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveLocationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationContext)
}

func (s *DirectiveLocationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveLocationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveLocationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveLocations(s)
	}
}

func (s *DirectiveLocationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveLocations(s)
	}
}

func (p *GraphQLParser) DirectiveLocations() (localctx IDirectiveLocationsContext) {
	this := p
	_ = this

	localctx = NewDirectiveLocationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, GraphQLParserRULE_directiveLocations)
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
		p.SetState(659)
		p.DirectiveLocation()
	}
	p.SetState(664)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLParserT__28 {
		{
			p.SetState(660)
			p.Match(GraphQLParserT__28)
		}
		{
			p.SetState(661)
			p.DirectiveLocation()
		}

		p.SetState(666)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDirectiveLocationContext is an interface to support dynamic dispatch.
type IDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveLocationContext differentiates from other interfaces.
	IsDirectiveLocationContext()
}

type DirectiveLocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveLocationContext() *DirectiveLocationContext {
	var p = new(DirectiveLocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveLocation
	return p
}

func (*DirectiveLocationContext) IsDirectiveLocationContext() {}

func NewDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveLocationContext {
	var p = new(DirectiveLocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveLocation

	return p
}

func (s *DirectiveLocationContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveLocationContext) ExecutableDirectiveLocation() IExecutableDirectiveLocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecutableDirectiveLocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecutableDirectiveLocationContext)
}

func (s *DirectiveLocationContext) TypeSystemDirectiveLocation() ITypeSystemDirectiveLocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSystemDirectiveLocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSystemDirectiveLocationContext)
}

func (s *DirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveLocation(s)
	}
}

func (s *DirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveLocation(s)
	}
}

func (p *GraphQLParser) DirectiveLocation() (localctx IDirectiveLocationContext) {
	this := p
	_ = this

	localctx = NewDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, GraphQLParserRULE_directiveLocation)

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

	p.SetState(669)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__32, GraphQLParserT__33, GraphQLParserT__34, GraphQLParserT__35, GraphQLParserT__36, GraphQLParserT__37, GraphQLParserT__38:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(667)
			p.ExecutableDirectiveLocation()
		}

	case GraphQLParserT__39, GraphQLParserT__40, GraphQLParserT__41, GraphQLParserT__42, GraphQLParserT__43, GraphQLParserT__44, GraphQLParserT__45, GraphQLParserT__46, GraphQLParserT__47, GraphQLParserT__48, GraphQLParserT__49:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(668)
			p.TypeSystemDirectiveLocation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExecutableDirectiveLocationContext is an interface to support dynamic dispatch.
type IExecutableDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecutableDirectiveLocationContext differentiates from other interfaces.
	IsExecutableDirectiveLocationContext()
}

type ExecutableDirectiveLocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutableDirectiveLocationContext() *ExecutableDirectiveLocationContext {
	var p = new(ExecutableDirectiveLocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_executableDirectiveLocation
	return p
}

func (*ExecutableDirectiveLocationContext) IsExecutableDirectiveLocationContext() {}

func NewExecutableDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutableDirectiveLocationContext {
	var p = new(ExecutableDirectiveLocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_executableDirectiveLocation

	return p
}

func (s *ExecutableDirectiveLocationContext) GetParser() antlr.Parser { return s.parser }
func (s *ExecutableDirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutableDirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutableDirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterExecutableDirectiveLocation(s)
	}
}

func (s *ExecutableDirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitExecutableDirectiveLocation(s)
	}
}

func (p *GraphQLParser) ExecutableDirectiveLocation() (localctx IExecutableDirectiveLocationContext) {
	this := p
	_ = this

	localctx = NewExecutableDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, GraphQLParserRULE_executableDirectiveLocation)
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
		p.SetState(671)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GraphQLParserT__32-33))|(1<<(GraphQLParserT__33-33))|(1<<(GraphQLParserT__34-33))|(1<<(GraphQLParserT__35-33))|(1<<(GraphQLParserT__36-33))|(1<<(GraphQLParserT__37-33))|(1<<(GraphQLParserT__38-33)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeSystemDirectiveLocationContext is an interface to support dynamic dispatch.
type ITypeSystemDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSystemDirectiveLocationContext differentiates from other interfaces.
	IsTypeSystemDirectiveLocationContext()
}

type TypeSystemDirectiveLocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemDirectiveLocationContext() *TypeSystemDirectiveLocationContext {
	var p = new(TypeSystemDirectiveLocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDirectiveLocation
	return p
}

func (*TypeSystemDirectiveLocationContext) IsTypeSystemDirectiveLocationContext() {}

func NewTypeSystemDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemDirectiveLocationContext {
	var p = new(TypeSystemDirectiveLocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemDirectiveLocation

	return p
}

func (s *TypeSystemDirectiveLocationContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeSystemDirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemDirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemDirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemDirectiveLocation(s)
	}
}

func (s *TypeSystemDirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemDirectiveLocation(s)
	}
}

func (p *GraphQLParser) TypeSystemDirectiveLocation() (localctx ITypeSystemDirectiveLocationContext) {
	this := p
	_ = this

	localctx = NewTypeSystemDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, GraphQLParserRULE_typeSystemDirectiveLocation)
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
		p.SetState(673)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(GraphQLParserT__39-40))|(1<<(GraphQLParserT__40-40))|(1<<(GraphQLParserT__41-40))|(1<<(GraphQLParserT__42-40))|(1<<(GraphQLParserT__43-40))|(1<<(GraphQLParserT__44-40))|(1<<(GraphQLParserT__45-40))|(1<<(GraphQLParserT__46-40))|(1<<(GraphQLParserT__47-40))|(1<<(GraphQLParserT__48-40))|(1<<(GraphQLParserT__49-40)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *GraphQLParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, GraphQLParserRULE_name)

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
		p.Match(GraphQLParserNAME)
	}

	return localctx
}

func (p *GraphQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 47:
		var t *ImplementsInterfacesContext = nil
		if localctx != nil {
			t = localctx.(*ImplementsInterfacesContext)
		}
		return p.ImplementsInterfaces_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GraphQLParser) ImplementsInterfaces_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
