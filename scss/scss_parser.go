// Code generated from ScssParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 521,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 3, 2, 7, 2, 108, 10,
	2, 12, 2, 14, 2, 111, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 124, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 129,
	10, 4, 12, 4, 14, 4, 132, 11, 4, 3, 4, 5, 4, 135, 10, 4, 3, 5, 3, 5, 5,
	5, 139, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 6, 7, 146, 10, 7, 13, 7, 14,
	7, 147, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 154, 10, 8, 3, 8, 5, 8, 157, 10,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 166, 10, 9, 3, 9, 3,
	9, 5, 9, 170, 10, 9, 5, 9, 172, 10, 9, 3, 9, 5, 9, 175, 10, 9, 5, 9, 177,
	10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 183, 10, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 188, 10, 10, 3, 10, 3, 10, 3, 11, 7, 11, 193, 10, 11, 12, 11,
	14, 11, 196, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 5, 13, 208, 10, 13, 3, 14, 6, 14, 211, 10, 14, 13, 14,
	14, 14, 212, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 219, 10, 14, 3, 14, 5,
	14, 222, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 237, 10, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 7, 18, 243, 10, 18, 12, 18, 14, 18, 246, 11, 18, 3, 18, 5, 18,
	249, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 21, 5, 21, 262, 10, 21, 3, 21, 5, 21, 265, 10, 21, 3, 22,
	3, 22, 3, 22, 5, 22, 270, 10, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 276,
	10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 282, 10, 23, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 306,
	10, 28, 12, 28, 14, 28, 309, 11, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 29, 7, 29, 318, 10, 29, 12, 29, 14, 29, 321, 11, 29, 3, 29, 3,
	29, 3, 29, 7, 29, 326, 10, 29, 12, 29, 14, 29, 329, 11, 29, 5, 29, 331,
	10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 337, 10, 30, 12, 30, 14, 30,
	340, 11, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 5, 31, 347, 10, 31, 3,
	32, 3, 32, 3, 32, 5, 32, 352, 10, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 33, 5, 33, 360, 10, 33, 3, 34, 3, 34, 3, 34, 7, 34, 365, 10, 34, 12,
	34, 14, 34, 368, 11, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 36, 3, 36, 7, 36, 379, 10, 36, 12, 36, 14, 36, 382, 11, 36, 3, 36, 7,
	36, 385, 10, 36, 12, 36, 14, 36, 388, 11, 36, 3, 37, 3, 37, 3, 37, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 7, 38, 398, 10, 38, 12, 38, 14, 38, 401, 11,
	38, 3, 38, 5, 38, 404, 10, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 7, 39,
	411, 10, 39, 12, 39, 14, 39, 414, 11, 39, 3, 40, 6, 40, 417, 10, 40, 13,
	40, 14, 40, 418, 3, 40, 3, 40, 3, 40, 7, 40, 424, 10, 40, 12, 40, 14, 40,
	427, 11, 40, 3, 40, 7, 40, 430, 10, 40, 12, 40, 14, 40, 433, 11, 40, 3,
	40, 5, 40, 436, 10, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 5, 42, 447, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 453,
	10, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 460, 10, 44, 3, 44, 3,
	44, 3, 45, 3, 45, 3, 46, 3, 46, 7, 46, 468, 10, 46, 12, 46, 14, 46, 471,
	11, 46, 3, 46, 3, 46, 3, 46, 3, 46, 7, 46, 477, 10, 46, 12, 46, 14, 46,
	480, 11, 46, 5, 46, 482, 10, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 5,
	47, 489, 10, 47, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50,
	3, 50, 3, 50, 7, 50, 501, 10, 50, 12, 50, 14, 50, 504, 11, 50, 3, 51, 3,
	51, 3, 51, 3, 51, 3, 52, 3, 52, 5, 52, 512, 10, 52, 3, 53, 3, 53, 3, 53,
	5, 53, 517, 10, 53, 3, 53, 3, 53, 3, 53, 2, 2, 54, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84,
	86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 2, 10, 3, 2, 27, 31, 5, 2, 15,
	15, 17, 17, 33, 34, 4, 2, 24, 24, 52, 52, 4, 2, 15, 16, 27, 27, 4, 2, 18,
	18, 26, 26, 3, 2, 52, 53, 3, 2, 35, 37, 4, 2, 52, 52, 63, 63, 2, 543, 2,
	109, 3, 2, 2, 2, 4, 123, 3, 2, 2, 2, 6, 125, 3, 2, 2, 2, 8, 136, 3, 2,
	2, 2, 10, 140, 3, 2, 2, 2, 12, 143, 3, 2, 2, 2, 14, 149, 3, 2, 2, 2, 16,
	160, 3, 2, 2, 2, 18, 178, 3, 2, 2, 2, 20, 194, 3, 2, 2, 2, 22, 199, 3,
	2, 2, 2, 24, 207, 3, 2, 2, 2, 26, 218, 3, 2, 2, 2, 28, 223, 3, 2, 2, 2,
	30, 225, 3, 2, 2, 2, 32, 236, 3, 2, 2, 2, 34, 238, 3, 2, 2, 2, 36, 250,
	3, 2, 2, 2, 38, 255, 3, 2, 2, 2, 40, 264, 3, 2, 2, 2, 42, 275, 3, 2, 2,
	2, 44, 277, 3, 2, 2, 2, 46, 285, 3, 2, 2, 2, 48, 293, 3, 2, 2, 2, 50, 295,
	3, 2, 2, 2, 52, 297, 3, 2, 2, 2, 54, 301, 3, 2, 2, 2, 56, 330, 3, 2, 2,
	2, 58, 332, 3, 2, 2, 2, 60, 343, 3, 2, 2, 2, 62, 348, 3, 2, 2, 2, 64, 359,
	3, 2, 2, 2, 66, 361, 3, 2, 2, 2, 68, 369, 3, 2, 2, 2, 70, 376, 3, 2, 2,
	2, 72, 389, 3, 2, 2, 2, 74, 392, 3, 2, 2, 2, 76, 407, 3, 2, 2, 2, 78, 416,
	3, 2, 2, 2, 80, 437, 3, 2, 2, 2, 82, 446, 3, 2, 2, 2, 84, 452, 3, 2, 2,
	2, 86, 454, 3, 2, 2, 2, 88, 463, 3, 2, 2, 2, 90, 481, 3, 2, 2, 2, 92, 488,
	3, 2, 2, 2, 94, 490, 3, 2, 2, 2, 96, 493, 3, 2, 2, 2, 98, 497, 3, 2, 2,
	2, 100, 505, 3, 2, 2, 2, 102, 509, 3, 2, 2, 2, 104, 513, 3, 2, 2, 2, 106,
	108, 5, 4, 3, 2, 107, 106, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107,
	3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 3, 3, 2, 2, 2, 111, 109, 3, 2, 2,
	2, 112, 124, 5, 62, 32, 2, 113, 124, 5, 68, 35, 2, 114, 124, 5, 72, 37,
	2, 115, 124, 5, 14, 8, 2, 116, 124, 5, 18, 10, 2, 117, 124, 5, 44, 23,
	2, 118, 124, 5, 16, 9, 2, 119, 124, 5, 34, 18, 2, 120, 124, 5, 46, 24,
	2, 121, 124, 5, 52, 27, 2, 122, 124, 5, 54, 28, 2, 123, 112, 3, 2, 2, 2,
	123, 113, 3, 2, 2, 2, 123, 114, 3, 2, 2, 2, 123, 115, 3, 2, 2, 2, 123,
	116, 3, 2, 2, 2, 123, 117, 3, 2, 2, 2, 123, 118, 3, 2, 2, 2, 123, 119,
	3, 2, 2, 2, 123, 120, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3, 2,
	2, 2, 124, 5, 3, 2, 2, 2, 125, 130, 5, 8, 5, 2, 126, 127, 7, 20, 2, 2,
	127, 129, 5, 8, 5, 2, 128, 126, 3, 2, 2, 2, 129, 132, 3, 2, 2, 2, 130,
	128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130,
	3, 2, 2, 2, 133, 135, 7, 7, 2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2,
	2, 2, 135, 7, 3, 2, 2, 2, 136, 138, 5, 10, 6, 2, 137, 139, 5, 12, 7, 2,
	138, 137, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 9, 3, 2, 2, 2, 140, 141,
	7, 22, 2, 2, 141, 142, 7, 52, 2, 2, 142, 11, 3, 2, 2, 2, 143, 145, 7, 18,
	2, 2, 144, 146, 5, 32, 17, 2, 145, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2,
	2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 13, 3, 2, 2, 2, 149,
	150, 7, 38, 2, 2, 150, 156, 7, 52, 2, 2, 151, 153, 7, 9, 2, 2, 152, 154,
	5, 6, 4, 2, 153, 152, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 3, 2,
	2, 2, 155, 157, 7, 10, 2, 2, 156, 151, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2,
	157, 158, 3, 2, 2, 2, 158, 159, 5, 74, 38, 2, 159, 15, 3, 2, 2, 2, 160,
	161, 7, 46, 2, 2, 161, 176, 7, 52, 2, 2, 162, 177, 7, 19, 2, 2, 163, 165,
	7, 9, 2, 2, 164, 166, 5, 98, 50, 2, 165, 164, 3, 2, 2, 2, 165, 166, 3,
	2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 169, 7, 10, 2, 2, 168, 170, 7, 19,
	2, 2, 169, 168, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 172, 3, 2, 2, 2,
	171, 163, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2, 173,
	175, 5, 74, 38, 2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177,
	3, 2, 2, 2, 176, 162, 3, 2, 2, 2, 176, 171, 3, 2, 2, 2, 177, 17, 3, 2,
	2, 2, 178, 179, 7, 39, 2, 2, 179, 180, 7, 52, 2, 2, 180, 182, 7, 9, 2,
	2, 181, 183, 5, 6, 4, 2, 182, 181, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183,
	184, 3, 2, 2, 2, 184, 185, 7, 10, 2, 2, 185, 187, 7, 11, 2, 2, 186, 188,
	5, 20, 11, 2, 187, 186, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 189, 3,
	2, 2, 2, 189, 190, 7, 12, 2, 2, 190, 19, 3, 2, 2, 2, 191, 193, 5, 24, 13,
	2, 192, 191, 3, 2, 2, 2, 193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194,
	195, 3, 2, 2, 2, 195, 197, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 198,
	5, 22, 12, 2, 198, 21, 3, 2, 2, 2, 199, 200, 7, 48, 2, 2, 200, 201, 5,
	26, 14, 2, 201, 202, 7, 19, 2, 2, 202, 23, 3, 2, 2, 2, 203, 204, 5, 26,
	14, 2, 204, 205, 7, 19, 2, 2, 205, 208, 3, 2, 2, 2, 206, 208, 5, 4, 3,
	2, 207, 203, 3, 2, 2, 2, 207, 206, 3, 2, 2, 2, 208, 25, 3, 2, 2, 2, 209,
	211, 5, 32, 17, 2, 210, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 210,
	3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 219, 3, 2, 2, 2, 214, 215, 7, 9,
	2, 2, 215, 216, 5, 26, 14, 2, 216, 217, 7, 10, 2, 2, 217, 219, 3, 2, 2,
	2, 218, 210, 3, 2, 2, 2, 218, 214, 3, 2, 2, 2, 219, 221, 3, 2, 2, 2, 220,
	222, 5, 30, 16, 2, 221, 220, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 27,
	3, 2, 2, 2, 223, 224, 9, 2, 2, 2, 224, 29, 3, 2, 2, 2, 225, 226, 5, 28,
	15, 2, 226, 227, 5, 26, 14, 2, 227, 31, 3, 2, 2, 2, 228, 237, 5, 102, 52,
	2, 229, 237, 5, 90, 46, 2, 230, 237, 7, 55, 2, 2, 231, 237, 7, 53, 2, 2,
	232, 237, 7, 3, 2, 2, 233, 237, 5, 100, 51, 2, 234, 237, 5, 10, 6, 2, 235,
	237, 5, 104, 53, 2, 236, 228, 3, 2, 2, 2, 236, 229, 3, 2, 2, 2, 236, 230,
	3, 2, 2, 2, 236, 231, 3, 2, 2, 2, 236, 232, 3, 2, 2, 2, 236, 233, 3, 2,
	2, 2, 236, 234, 3, 2, 2, 2, 236, 235, 3, 2, 2, 2, 237, 33, 3, 2, 2, 2,
	238, 239, 7, 42, 2, 2, 239, 240, 5, 40, 21, 2, 240, 244, 5, 74, 38, 2,
	241, 243, 5, 36, 19, 2, 242, 241, 3, 2, 2, 2, 243, 246, 3, 2, 2, 2, 244,
	242, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 248, 3, 2, 2, 2, 246, 244,
	3, 2, 2, 2, 247, 249, 5, 38, 20, 2, 248, 247, 3, 2, 2, 2, 248, 249, 3,
	2, 2, 2, 249, 35, 3, 2, 2, 2, 250, 251, 7, 40, 2, 2, 251, 252, 7, 41, 2,
	2, 252, 253, 5, 40, 21, 2, 253, 254, 5, 74, 38, 2, 254, 37, 3, 2, 2, 2,
	255, 256, 7, 40, 2, 2, 256, 257, 5, 74, 38, 2, 257, 39, 3, 2, 2, 2, 258,
	261, 5, 42, 22, 2, 259, 260, 7, 6, 2, 2, 260, 262, 5, 40, 21, 2, 261, 259,
	3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 265, 3, 2, 2, 2, 263, 265, 7, 3,
	2, 2, 264, 258, 3, 2, 2, 2, 264, 263, 3, 2, 2, 2, 265, 41, 3, 2, 2, 2,
	266, 269, 5, 26, 14, 2, 267, 268, 9, 3, 2, 2, 268, 270, 5, 40, 21, 2, 269,
	267, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 276, 3, 2, 2, 2, 271, 272,
	7, 9, 2, 2, 272, 273, 5, 40, 21, 2, 273, 274, 7, 10, 2, 2, 274, 276, 3,
	2, 2, 2, 275, 266, 3, 2, 2, 2, 275, 271, 3, 2, 2, 2, 276, 43, 3, 2, 2,
	2, 277, 278, 5, 10, 6, 2, 278, 279, 7, 18, 2, 2, 279, 281, 5, 98, 50, 2,
	280, 282, 7, 51, 2, 2, 281, 280, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282,
	283, 3, 2, 2, 2, 283, 284, 7, 19, 2, 2, 284, 45, 3, 2, 2, 2, 285, 286,
	7, 43, 2, 2, 286, 287, 5, 10, 6, 2, 287, 288, 7, 49, 2, 2, 288, 289, 5,
	48, 25, 2, 289, 290, 7, 50, 2, 2, 290, 291, 5, 50, 26, 2, 291, 292, 5,
	74, 38, 2, 292, 47, 3, 2, 2, 2, 293, 294, 7, 54, 2, 2, 294, 49, 3, 2, 2,
	2, 295, 296, 7, 54, 2, 2, 296, 51, 3, 2, 2, 2, 297, 298, 7, 44, 2, 2, 298,
	299, 5, 40, 21, 2, 299, 300, 5, 74, 38, 2, 300, 53, 3, 2, 2, 2, 301, 302,
	7, 45, 2, 2, 302, 307, 5, 10, 6, 2, 303, 304, 7, 20, 2, 2, 304, 306, 5,
	10, 6, 2, 305, 303, 3, 2, 2, 2, 306, 309, 3, 2, 2, 2, 307, 305, 3, 2, 2,
	2, 307, 308, 3, 2, 2, 2, 308, 310, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 310,
	311, 7, 4, 2, 2, 311, 312, 5, 56, 29, 2, 312, 313, 5, 74, 38, 2, 313, 55,
	3, 2, 2, 2, 314, 319, 7, 52, 2, 2, 315, 316, 7, 20, 2, 2, 316, 318, 7,
	52, 2, 2, 317, 315, 3, 2, 2, 2, 318, 321, 3, 2, 2, 2, 319, 317, 3, 2, 2,
	2, 319, 320, 3, 2, 2, 2, 320, 331, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 322,
	327, 5, 58, 30, 2, 323, 324, 7, 20, 2, 2, 324, 326, 5, 58, 30, 2, 325,
	323, 3, 2, 2, 2, 326, 329, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327, 328,
	3, 2, 2, 2, 328, 331, 3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 330, 314, 3, 2,
	2, 2, 330, 322, 3, 2, 2, 2, 331, 57, 3, 2, 2, 2, 332, 333, 7, 9, 2, 2,
	333, 338, 5, 60, 31, 2, 334, 335, 7, 20, 2, 2, 335, 337, 5, 60, 31, 2,
	336, 334, 3, 2, 2, 2, 337, 340, 3, 2, 2, 2, 338, 336, 3, 2, 2, 2, 338,
	339, 3, 2, 2, 2, 339, 341, 3, 2, 2, 2, 340, 338, 3, 2, 2, 2, 341, 342,
	7, 10, 2, 2, 342, 59, 3, 2, 2, 2, 343, 346, 5, 90, 46, 2, 344, 345, 7,
	18, 2, 2, 345, 347, 5, 98, 50, 2, 346, 344, 3, 2, 2, 2, 346, 347, 3, 2,
	2, 2, 347, 61, 3, 2, 2, 2, 348, 349, 7, 47, 2, 2, 349, 351, 5, 64, 33,
	2, 350, 352, 5, 66, 34, 2, 351, 350, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2,
	352, 353, 3, 2, 2, 2, 353, 354, 7, 19, 2, 2, 354, 63, 3, 2, 2, 2, 355,
	360, 7, 53, 2, 2, 356, 357, 7, 32, 2, 2, 357, 358, 7, 60, 2, 2, 358, 360,
	7, 59, 2, 2, 359, 355, 3, 2, 2, 2, 359, 356, 3, 2, 2, 2, 360, 65, 3, 2,
	2, 2, 361, 366, 7, 52, 2, 2, 362, 363, 7, 20, 2, 2, 363, 365, 7, 52, 2,
	2, 364, 362, 3, 2, 2, 2, 365, 368, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 366,
	367, 3, 2, 2, 2, 367, 67, 3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 369, 370, 7,
	23, 2, 2, 370, 371, 5, 70, 36, 2, 371, 372, 5, 76, 39, 2, 372, 373, 7,
	11, 2, 2, 373, 374, 5, 2, 2, 2, 374, 375, 7, 12, 2, 2, 375, 69, 3, 2, 2,
	2, 376, 380, 9, 4, 2, 2, 377, 379, 7, 52, 2, 2, 378, 377, 3, 2, 2, 2, 379,
	382, 3, 2, 2, 2, 380, 378, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 386,
	3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 383, 385, 5, 84, 43, 2, 384, 383, 3,
	2, 2, 2, 385, 388, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2,
	2, 387, 71, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 389, 390, 5, 76, 39, 2, 390,
	391, 5, 74, 38, 2, 391, 73, 3, 2, 2, 2, 392, 399, 7, 11, 2, 2, 393, 394,
	5, 96, 49, 2, 394, 395, 7, 19, 2, 2, 395, 398, 3, 2, 2, 2, 396, 398, 5,
	4, 3, 2, 397, 393, 3, 2, 2, 2, 397, 396, 3, 2, 2, 2, 398, 401, 3, 2, 2,
	2, 399, 397, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 403, 3, 2, 2, 2, 401,
	399, 3, 2, 2, 2, 402, 404, 5, 96, 49, 2, 403, 402, 3, 2, 2, 2, 403, 404,
	3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 406, 7, 12, 2, 2, 406, 75, 3, 2,
	2, 2, 407, 412, 5, 78, 40, 2, 408, 409, 7, 20, 2, 2, 409, 411, 5, 78, 40,
	2, 410, 408, 3, 2, 2, 2, 411, 414, 3, 2, 2, 2, 412, 410, 3, 2, 2, 2, 412,
	413, 3, 2, 2, 2, 413, 77, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2, 415, 417, 5,
	82, 42, 2, 416, 415, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 416, 3, 2,
	2, 2, 418, 419, 3, 2, 2, 2, 419, 425, 3, 2, 2, 2, 420, 421, 5, 80, 41,
	2, 421, 422, 5, 82, 42, 2, 422, 424, 3, 2, 2, 2, 423, 420, 3, 2, 2, 2,
	424, 427, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426,
	431, 3, 2, 2, 2, 427, 425, 3, 2, 2, 2, 428, 430, 5, 86, 44, 2, 429, 428,
	3, 2, 2, 2, 430, 433, 3, 2, 2, 2, 431, 429, 3, 2, 2, 2, 431, 432, 3, 2,
	2, 2, 432, 435, 3, 2, 2, 2, 433, 431, 3, 2, 2, 2, 434, 436, 5, 84, 43,
	2, 435, 434, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 79, 3, 2, 2, 2, 437,
	438, 9, 5, 2, 2, 438, 81, 3, 2, 2, 2, 439, 447, 5, 90, 46, 2, 440, 441,
	7, 25, 2, 2, 441, 447, 5, 90, 46, 2, 442, 443, 7, 21, 2, 2, 443, 447, 5,
	90, 46, 2, 444, 447, 7, 24, 2, 2, 445, 447, 7, 28, 2, 2, 446, 439, 3, 2,
	2, 2, 446, 440, 3, 2, 2, 2, 446, 442, 3, 2, 2, 2, 446, 444, 3, 2, 2, 2,
	446, 445, 3, 2, 2, 2, 447, 83, 3, 2, 2, 2, 448, 449, 9, 6, 2, 2, 449, 453,
	7, 52, 2, 2, 450, 451, 9, 6, 2, 2, 451, 453, 5, 104, 53, 2, 452, 448, 3,
	2, 2, 2, 452, 450, 3, 2, 2, 2, 453, 85, 3, 2, 2, 2, 454, 455, 7, 13, 2,
	2, 455, 459, 7, 52, 2, 2, 456, 457, 5, 88, 45, 2, 457, 458, 9, 7, 2, 2,
	458, 460, 3, 2, 2, 2, 459, 456, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2, 460,
	461, 3, 2, 2, 2, 461, 462, 7, 14, 2, 2, 462, 87, 3, 2, 2, 2, 463, 464,
	9, 8, 2, 2, 464, 89, 3, 2, 2, 2, 465, 469, 7, 52, 2, 2, 466, 468, 5, 92,
	47, 2, 467, 466, 3, 2, 2, 2, 468, 471, 3, 2, 2, 2, 469, 467, 3, 2, 2, 2,
	469, 470, 3, 2, 2, 2, 470, 482, 3, 2, 2, 2, 471, 469, 3, 2, 2, 2, 472,
	473, 7, 8, 2, 2, 473, 474, 5, 94, 48, 2, 474, 478, 7, 12, 2, 2, 475, 477,
	5, 92, 47, 2, 476, 475, 3, 2, 2, 2, 477, 480, 3, 2, 2, 2, 478, 476, 3,
	2, 2, 2, 478, 479, 3, 2, 2, 2, 479, 482, 3, 2, 2, 2, 480, 478, 3, 2, 2,
	2, 481, 465, 3, 2, 2, 2, 481, 472, 3, 2, 2, 2, 482, 91, 3, 2, 2, 2, 483,
	484, 7, 62, 2, 2, 484, 485, 5, 94, 48, 2, 485, 486, 7, 12, 2, 2, 486, 489,
	3, 2, 2, 2, 487, 489, 7, 63, 2, 2, 488, 483, 3, 2, 2, 2, 488, 487, 3, 2,
	2, 2, 489, 93, 3, 2, 2, 2, 490, 491, 7, 22, 2, 2, 491, 492, 9, 9, 2, 2,
	492, 95, 3, 2, 2, 2, 493, 494, 5, 90, 46, 2, 494, 495, 7, 18, 2, 2, 495,
	496, 5, 98, 50, 2, 496, 97, 3, 2, 2, 2, 497, 502, 5, 26, 14, 2, 498, 499,
	7, 20, 2, 2, 499, 501, 5, 26, 14, 2, 500, 498, 3, 2, 2, 2, 501, 504, 3,
	2, 2, 2, 502, 500, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 99, 3, 2, 2,
	2, 504, 502, 3, 2, 2, 2, 505, 506, 7, 32, 2, 2, 506, 507, 7, 60, 2, 2,
	507, 508, 7, 59, 2, 2, 508, 101, 3, 2, 2, 2, 509, 511, 7, 54, 2, 2, 510,
	512, 7, 5, 2, 2, 511, 510, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 512, 103,
	3, 2, 2, 2, 513, 514, 7, 52, 2, 2, 514, 516, 7, 9, 2, 2, 515, 517, 5, 98,
	50, 2, 516, 515, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 518, 3, 2, 2, 2,
	518, 519, 7, 10, 2, 2, 519, 105, 3, 2, 2, 2, 59, 109, 123, 130, 134, 138,
	147, 153, 156, 165, 169, 171, 174, 176, 182, 187, 194, 207, 212, 218, 221,
	236, 244, 248, 261, 264, 269, 275, 281, 307, 319, 327, 330, 338, 346, 351,
	359, 366, 380, 386, 397, 399, 403, 412, 418, 425, 431, 435, 446, 452, 459,
	469, 478, 481, 488, 502, 511, 516,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'null'", "'in'", "", "", "'...'", "", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "'>'", "'~'", "'<'", "':'", "';'", "','", "'.'", "'$'", "'@'",
	"'&'", "'#'", "'::'", "'+'", "'*'", "'/'", "'-'", "'%'", "", "'=='", "'!='",
	"'='", "'|='", "'~='", "'@mixin'", "'@function'", "'@else'", "'if'", "'@if'",
	"'@for'", "'@while'", "'@each'", "'@include'", "'@import'", "'@return'",
	"'from'", "'through'", "'!default'",
}
var symbolicNames = []string{
	"", "NULL", "IN", "Unit", "COMBINE_COMPARE", "Ellipsis", "InterpolationStart",
	"LPAREN", "RPAREN", "BlockStart", "BlockEnd", "LBRACK", "RBRACK", "GT",
	"TIL", "LT", "COLON", "SEMI", "COMMA", "DOT", "DOLLAR", "AT", "AND", "HASH",
	"COLONCOLON", "PLUS", "TIMES", "DIV", "MINUS", "PERC", "UrlStart", "EQEQ",
	"NOTEQ", "EQ", "PIPE_EQ", "TILD_EQ", "MIXIN", "FUNCTION", "AT_ELSE", "IF",
	"AT_IF", "AT_FOR", "AT_WHILE", "AT_EACH", "INCLUDE", "IMPORT", "RETURN",
	"FROM", "THROUGH", "POUND_DEFAULT", "Identifier", "StringLiteral", "Number",
	"Color", "WS", "SL_COMMENT", "COMMENT", "UrlEnd", "Url", "SPACE", "InterpolationStartAfter",
	"IdentifierAfter",
}

var ruleNames = []string{
	"stylesheet", "statement", "params", "param", "variableName", "paramOptionalValue",
	"mixinDeclaration", "includeDeclaration", "functionDeclaration", "functionBody",
	"functionReturn", "functionStatement", "commandStatement", "mathCharacter",
	"mathStatement", "expression", "ifDeclaration", "elseIfStatement", "elseStatement",
	"conditions", "condition", "variableDeclaration", "forDeclaration", "fromNumber",
	"throughNumber", "whileDeclaration", "eachDeclaration", "eachValueList",
	"identifierListOrMap", "identifierValue", "importDeclaration", "referenceUrl",
	"mediaTypes", "nested", "nest", "ruleset", "block", "selectors", "selector",
	"selectorPrefix", "element", "pseudo", "attrib", "attribRelate", "identifier",
	"identifierPart", "identifierVariableName", "property", "values", "url",
	"measurement", "functionCall",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ScssParser struct {
	*antlr.BaseParser
}

func NewScssParser(input antlr.TokenStream) *ScssParser {
	this := new(ScssParser)

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
	ScssParserNULL                    = 1
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
	ScssParserCOLONCOLON              = 24
	ScssParserPLUS                    = 25
	ScssParserTIMES                   = 26
	ScssParserDIV                     = 27
	ScssParserMINUS                   = 28
	ScssParserPERC                    = 29
	ScssParserUrlStart                = 30
	ScssParserEQEQ                    = 31
	ScssParserNOTEQ                   = 32
	ScssParserEQ                      = 33
	ScssParserPIPE_EQ                 = 34
	ScssParserTILD_EQ                 = 35
	ScssParserMIXIN                   = 36
	ScssParserFUNCTION                = 37
	ScssParserAT_ELSE                 = 38
	ScssParserIF                      = 39
	ScssParserAT_IF                   = 40
	ScssParserAT_FOR                  = 41
	ScssParserAT_WHILE                = 42
	ScssParserAT_EACH                 = 43
	ScssParserINCLUDE                 = 44
	ScssParserIMPORT                  = 45
	ScssParserRETURN                  = 46
	ScssParserFROM                    = 47
	ScssParserTHROUGH                 = 48
	ScssParserPOUND_DEFAULT           = 49
	ScssParserIdentifier              = 50
	ScssParserStringLiteral           = 51
	ScssParserNumber                  = 52
	ScssParserColor                   = 53
	ScssParserWS                      = 54
	ScssParserSL_COMMENT              = 55
	ScssParserCOMMENT                 = 56
	ScssParserUrlEnd                  = 57
	ScssParserUrl                     = 58
	ScssParserSPACE                   = 59
	ScssParserInterpolationStartAfter = 60
	ScssParserIdentifierAfter         = 61
)

// ScssParser rules.
const (
	ScssParserRULE_stylesheet             = 0
	ScssParserRULE_statement              = 1
	ScssParserRULE_params                 = 2
	ScssParserRULE_param                  = 3
	ScssParserRULE_variableName           = 4
	ScssParserRULE_paramOptionalValue     = 5
	ScssParserRULE_mixinDeclaration       = 6
	ScssParserRULE_includeDeclaration     = 7
	ScssParserRULE_functionDeclaration    = 8
	ScssParserRULE_functionBody           = 9
	ScssParserRULE_functionReturn         = 10
	ScssParserRULE_functionStatement      = 11
	ScssParserRULE_commandStatement       = 12
	ScssParserRULE_mathCharacter          = 13
	ScssParserRULE_mathStatement          = 14
	ScssParserRULE_expression             = 15
	ScssParserRULE_ifDeclaration          = 16
	ScssParserRULE_elseIfStatement        = 17
	ScssParserRULE_elseStatement          = 18
	ScssParserRULE_conditions             = 19
	ScssParserRULE_condition              = 20
	ScssParserRULE_variableDeclaration    = 21
	ScssParserRULE_forDeclaration         = 22
	ScssParserRULE_fromNumber             = 23
	ScssParserRULE_throughNumber          = 24
	ScssParserRULE_whileDeclaration       = 25
	ScssParserRULE_eachDeclaration        = 26
	ScssParserRULE_eachValueList          = 27
	ScssParserRULE_identifierListOrMap    = 28
	ScssParserRULE_identifierValue        = 29
	ScssParserRULE_importDeclaration      = 30
	ScssParserRULE_referenceUrl           = 31
	ScssParserRULE_mediaTypes             = 32
	ScssParserRULE_nested                 = 33
	ScssParserRULE_nest                   = 34
	ScssParserRULE_ruleset                = 35
	ScssParserRULE_block                  = 36
	ScssParserRULE_selectors              = 37
	ScssParserRULE_selector               = 38
	ScssParserRULE_selectorPrefix         = 39
	ScssParserRULE_element                = 40
	ScssParserRULE_pseudo                 = 41
	ScssParserRULE_attrib                 = 42
	ScssParserRULE_attribRelate           = 43
	ScssParserRULE_identifier             = 44
	ScssParserRULE_identifierPart         = 45
	ScssParserRULE_identifierVariableName = 46
	ScssParserRULE_property               = 47
	ScssParserRULE_values                 = 48
	ScssParserRULE_url                    = 49
	ScssParserRULE_measurement            = 50
	ScssParserRULE_functionCall           = 51
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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserInterpolationStart)|(1<<ScssParserDOT)|(1<<ScssParserDOLLAR)|(1<<ScssParserAT)|(1<<ScssParserAND)|(1<<ScssParserHASH)|(1<<ScssParserTIMES))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ScssParserMIXIN-36))|(1<<(ScssParserFUNCTION-36))|(1<<(ScssParserAT_IF-36))|(1<<(ScssParserAT_FOR-36))|(1<<(ScssParserAT_WHILE-36))|(1<<(ScssParserAT_EACH-36))|(1<<(ScssParserINCLUDE-36))|(1<<(ScssParserIMPORT-36))|(1<<(ScssParserIdentifier-36)))) != 0) {
		{
			p.SetState(104)
			p.Statement()
		}

		p.SetState(109)
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

func (s *StatementContext) Nested() INestedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INestedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INestedContext)
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

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.ImportDeclaration()
		}

	case ScssParserAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Nested()
		}

	case ScssParserInterpolationStart, ScssParserDOT, ScssParserAND, ScssParserHASH, ScssParserTIMES, ScssParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Ruleset()
		}

	case ScssParserMIXIN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.MixinDeclaration()
		}

	case ScssParserFUNCTION:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.FunctionDeclaration()
		}

	case ScssParserDOLLAR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(115)
			p.VariableDeclaration()
		}

	case ScssParserINCLUDE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(116)
			p.IncludeDeclaration()
		}

	case ScssParserAT_IF:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(117)
			p.IfDeclaration()
		}

	case ScssParserAT_FOR:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(118)
			p.ForDeclaration()
		}

	case ScssParserAT_WHILE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(119)
			p.WhileDeclaration()
		}

	case ScssParserAT_EACH:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(120)
			p.EachDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *ParamsContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *ParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *ParamsContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(ScssParserEllipsis, 0)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *ScssParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ScssParserRULE_params)
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
		p.SetState(123)
		p.Param()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserCOMMA {
		{
			p.SetState(124)
			p.Match(ScssParserCOMMA)
		}
		{
			p.SetState(125)
			p.Param()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserEllipsis {
		{
			p.SetState(131)
			p.Match(ScssParserEllipsis)
		}

	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *ParamContext) ParamOptionalValue() IParamOptionalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamOptionalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamOptionalValueContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *ScssParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ScssParserRULE_param)
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
		p.SetState(134)
		p.VariableName()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserCOLON {
		{
			p.SetState(135)
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

func (s *VariableNameContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(ScssParserDOLLAR, 0)
}

func (s *VariableNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
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
	localctx = NewVariableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ScssParserRULE_variableName)

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
		p.SetState(138)
		p.Match(ScssParserDOLLAR)
	}
	{
		p.SetState(139)
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
		p.SetState(141)
		p.Match(ScssParserCOLON)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL)|(1<<ScssParserInterpolationStart)|(1<<ScssParserDOLLAR)|(1<<ScssParserUrlStart))) != 0) || (((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(ScssParserIdentifier-50))|(1<<(ScssParserStringLiteral-50))|(1<<(ScssParserNumber-50))|(1<<(ScssParserColor-50)))) != 0) {
		{
			p.SetState(142)
			p.Expression()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *MixinDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *MixinDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MixinDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *MixinDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *MixinDeclarationContext) Params() IParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
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
	localctx = NewMixinDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ScssParserRULE_mixinDeclaration)
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
		p.SetState(147)
		p.Match(ScssParserMIXIN)
	}
	{
		p.SetState(148)
		p.Match(ScssParserIdentifier)
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserLPAREN {
		{
			p.SetState(149)
			p.Match(ScssParserLPAREN)
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserDOLLAR {
			{
				p.SetState(150)
				p.Params()
			}

		}
		{
			p.SetState(153)
			p.Match(ScssParserRPAREN)
		}

	}
	{
		p.SetState(156)
		p.Block()
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

func (s *IncludeDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, 0)
}

func (s *IncludeDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *IncludeDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *IncludeDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IncludeDeclarationContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
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
	localctx = NewIncludeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ScssParserRULE_includeDeclaration)
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
		p.SetState(158)
		p.Match(ScssParserINCLUDE)
	}
	{
		p.SetState(159)
		p.Match(ScssParserIdentifier)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserSEMI:
		{
			p.SetState(160)
			p.Match(ScssParserSEMI)
		}

	case ScssParserNULL, ScssParserInterpolationStart, ScssParserLPAREN, ScssParserBlockStart, ScssParserBlockEnd, ScssParserDOT, ScssParserDOLLAR, ScssParserAT, ScssParserAND, ScssParserHASH, ScssParserTIMES, ScssParserUrlStart, ScssParserMIXIN, ScssParserFUNCTION, ScssParserAT_IF, ScssParserAT_FOR, ScssParserAT_WHILE, ScssParserAT_EACH, ScssParserINCLUDE, ScssParserIMPORT, ScssParserRETURN, ScssParserIdentifier, ScssParserStringLiteral, ScssParserNumber, ScssParserColor:
		p.SetState(169)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(161)
				p.Match(ScssParserLPAREN)
			}
			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserDOLLAR)|(1<<ScssParserUrlStart))) != 0) || (((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(ScssParserIdentifier-50))|(1<<(ScssParserStringLiteral-50))|(1<<(ScssParserNumber-50))|(1<<(ScssParserColor-50)))) != 0) {
				{
					p.SetState(162)
					p.Values()
				}

			}
			{
				p.SetState(165)
				p.Match(ScssParserRPAREN)
			}
			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ScssParserSEMI {
				{
					p.SetState(166)
					p.Match(ScssParserSEMI)
				}

			}

		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserBlockStart {
			{
				p.SetState(171)
				p.Block()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
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

func (s *FunctionDeclarationContext) Params() IParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
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
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ScssParserRULE_functionDeclaration)
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
		p.SetState(176)
		p.Match(ScssParserFUNCTION)
	}
	{
		p.SetState(177)
		p.Match(ScssParserIdentifier)
	}
	{
		p.SetState(178)
		p.Match(ScssParserLPAREN)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserDOLLAR {
		{
			p.SetState(179)
			p.Params()
		}

	}
	{
		p.SetState(182)
		p.Match(ScssParserRPAREN)
	}
	{
		p.SetState(183)
		p.Match(ScssParserBlockStart)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserDOT)|(1<<ScssParserDOLLAR)|(1<<ScssParserAT)|(1<<ScssParserAND)|(1<<ScssParserHASH)|(1<<ScssParserTIMES)|(1<<ScssParserUrlStart))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ScssParserMIXIN-36))|(1<<(ScssParserFUNCTION-36))|(1<<(ScssParserAT_IF-36))|(1<<(ScssParserAT_FOR-36))|(1<<(ScssParserAT_WHILE-36))|(1<<(ScssParserAT_EACH-36))|(1<<(ScssParserINCLUDE-36))|(1<<(ScssParserIMPORT-36))|(1<<(ScssParserRETURN-36))|(1<<(ScssParserIdentifier-36))|(1<<(ScssParserStringLiteral-36))|(1<<(ScssParserNumber-36))|(1<<(ScssParserColor-36)))) != 0) {
		{
			p.SetState(184)
			p.FunctionBody()
		}

	}
	{
		p.SetState(187)
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
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ScssParserRULE_functionBody)
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
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserDOT)|(1<<ScssParserDOLLAR)|(1<<ScssParserAT)|(1<<ScssParserAND)|(1<<ScssParserHASH)|(1<<ScssParserTIMES)|(1<<ScssParserUrlStart))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ScssParserMIXIN-36))|(1<<(ScssParserFUNCTION-36))|(1<<(ScssParserAT_IF-36))|(1<<(ScssParserAT_FOR-36))|(1<<(ScssParserAT_WHILE-36))|(1<<(ScssParserAT_EACH-36))|(1<<(ScssParserINCLUDE-36))|(1<<(ScssParserIMPORT-36))|(1<<(ScssParserIdentifier-36))|(1<<(ScssParserStringLiteral-36))|(1<<(ScssParserNumber-36))|(1<<(ScssParserColor-36)))) != 0) {
		{
			p.SetState(189)
			p.FunctionStatement()
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(195)
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
	localctx = NewFunctionReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ScssParserRULE_functionReturn)

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
		p.SetState(197)
		p.Match(ScssParserRETURN)
	}
	{
		p.SetState(198)
		p.CommandStatement()
	}
	{
		p.SetState(199)
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
	localctx = NewFunctionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ScssParserRULE_functionStatement)

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

	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.CommandStatement()
		}
		{
			p.SetState(202)
			p.Match(ScssParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
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

func (s *CommandStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
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

func (s *CommandStatementContext) MathStatement() IMathStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathStatementContext)
}

func (s *CommandStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CommandStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
	localctx = NewCommandStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ScssParserRULE_commandStatement)
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
	p.SetState(216)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserNULL, ScssParserInterpolationStart, ScssParserDOLLAR, ScssParserUrlStart, ScssParserIdentifier, ScssParserStringLiteral, ScssParserNumber, ScssParserColor:
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL)|(1<<ScssParserInterpolationStart)|(1<<ScssParserDOLLAR)|(1<<ScssParserUrlStart))) != 0) || (((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(ScssParserIdentifier-50))|(1<<(ScssParserStringLiteral-50))|(1<<(ScssParserNumber-50))|(1<<(ScssParserColor-50)))) != 0) {
			{
				p.SetState(207)
				p.Expression()
			}

			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ScssParserLPAREN:
		{
			p.SetState(212)
			p.Match(ScssParserLPAREN)
		}
		{
			p.SetState(213)
			p.CommandStatement()
		}
		{
			p.SetState(214)
			p.Match(ScssParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserPLUS)|(1<<ScssParserTIMES)|(1<<ScssParserDIV)|(1<<ScssParserMINUS)|(1<<ScssParserPERC))) != 0 {
		{
			p.SetState(218)
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
	localctx = NewMathCharacterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ScssParserRULE_mathCharacter)
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
	localctx = NewMathStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ScssParserRULE_mathStatement)

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
		p.SetState(223)
		p.MathCharacter()
	}
	{
		p.SetState(224)
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

func (s *ExpressionContext) NULL() antlr.TerminalNode {
	return s.GetToken(ScssParserNULL, 0)
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
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ScssParserRULE_expression)

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

	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.Measurement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Identifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(228)
			p.Match(ScssParserColor)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(229)
			p.Match(ScssParserStringLiteral)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(230)
			p.Match(ScssParserNULL)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(231)
			p.Url()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(232)
			p.VariableName()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(233)
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
	localctx = NewIfDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ScssParserRULE_ifDeclaration)
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
		p.SetState(236)
		p.Match(ScssParserAT_IF)
	}
	{
		p.SetState(237)
		p.Conditions()
	}
	{
		p.SetState(238)
		p.Block()
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(239)
				p.ElseIfStatement()
			}

		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserAT_ELSE {
		{
			p.SetState(245)
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
	localctx = NewElseIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ScssParserRULE_elseIfStatement)

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
		p.SetState(248)
		p.Match(ScssParserAT_ELSE)
	}
	{
		p.SetState(249)
		p.Match(ScssParserIF)
	}
	{
		p.SetState(250)
		p.Conditions()
	}
	{
		p.SetState(251)
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
	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ScssParserRULE_elseStatement)

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
		p.SetState(253)
		p.Match(ScssParserAT_ELSE)
	}
	{
		p.SetState(254)
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

func (s *ConditionsContext) NULL() antlr.TerminalNode {
	return s.GetToken(ScssParserNULL, 0)
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
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ScssParserRULE_conditions)

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

	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Condition()
		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(257)
				p.Match(ScssParserCOMBINE_COMPARE)
			}
			{
				p.SetState(258)
				p.Conditions()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Match(ScssParserNULL)
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
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ScssParserRULE_condition)
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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.CommandStatement()
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(ScssParserGT-13))|(1<<(ScssParserLT-13))|(1<<(ScssParserEQEQ-13))|(1<<(ScssParserNOTEQ-13)))) != 0 {
			{
				p.SetState(265)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(ScssParserGT-13))|(1<<(ScssParserLT-13))|(1<<(ScssParserEQEQ-13))|(1<<(ScssParserNOTEQ-13)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(266)
				p.Conditions()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(269)
			p.Match(ScssParserLPAREN)
		}
		{
			p.SetState(270)
			p.Conditions()
		}
		{
			p.SetState(271)
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

func (s *VariableDeclarationContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *VariableDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, 0)
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
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ScssParserRULE_variableDeclaration)
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
		p.SetState(275)
		p.VariableName()
	}
	{
		p.SetState(276)
		p.Match(ScssParserCOLON)
	}
	{
		p.SetState(277)
		p.Values()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserPOUND_DEFAULT {
		{
			p.SetState(278)
			p.Match(ScssParserPOUND_DEFAULT)
		}

	}
	{
		p.SetState(281)
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

func (s *ForDeclarationContext) THROUGH() antlr.TerminalNode {
	return s.GetToken(ScssParserTHROUGH, 0)
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
	localctx = NewForDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ScssParserRULE_forDeclaration)

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
		p.SetState(283)
		p.Match(ScssParserAT_FOR)
	}
	{
		p.SetState(284)
		p.VariableName()
	}
	{
		p.SetState(285)
		p.Match(ScssParserFROM)
	}
	{
		p.SetState(286)
		p.FromNumber()
	}
	{
		p.SetState(287)
		p.Match(ScssParserTHROUGH)
	}
	{
		p.SetState(288)
		p.ThroughNumber()
	}
	{
		p.SetState(289)
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
	localctx = NewFromNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ScssParserRULE_fromNumber)

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
		p.SetState(291)
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
	localctx = NewThroughNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ScssParserRULE_throughNumber)

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
		p.SetState(293)
		p.Match(ScssParserNumber)
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
	localctx = NewWhileDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ScssParserRULE_whileDeclaration)

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
		p.SetState(295)
		p.Match(ScssParserAT_WHILE)
	}
	{
		p.SetState(296)
		p.Conditions()
	}
	{
		p.SetState(297)
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
	localctx = NewEachDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ScssParserRULE_eachDeclaration)
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
		p.Match(ScssParserAT_EACH)
	}
	{
		p.SetState(300)
		p.VariableName()
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserCOMMA {
		{
			p.SetState(301)
			p.Match(ScssParserCOMMA)
		}
		{
			p.SetState(302)
			p.VariableName()
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(308)
		p.Match(ScssParserIN)
	}
	{
		p.SetState(309)
		p.EachValueList()
	}
	{
		p.SetState(310)
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

func (s *EachValueListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(ScssParserIdentifier)
}

func (s *EachValueListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, i)
}

func (s *EachValueListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *EachValueListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *EachValueListContext) AllIdentifierListOrMap() []IIdentifierListOrMapContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierListOrMapContext)(nil)).Elem())
	var tst = make([]IIdentifierListOrMapContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierListOrMapContext)
		}
	}

	return tst
}

func (s *EachValueListContext) IdentifierListOrMap(i int) IIdentifierListOrMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListOrMapContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListOrMapContext)
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
	localctx = NewEachValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ScssParserRULE_eachValueList)
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

	p.SetState(328)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(312)
			p.Match(ScssParserIdentifier)
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScssParserCOMMA {
			{
				p.SetState(313)
				p.Match(ScssParserCOMMA)
			}
			{
				p.SetState(314)
				p.Match(ScssParserIdentifier)
			}

			p.SetState(319)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ScssParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.IdentifierListOrMap()
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScssParserCOMMA {
			{
				p.SetState(321)
				p.Match(ScssParserCOMMA)
			}
			{
				p.SetState(322)
				p.IdentifierListOrMap()
			}

			p.SetState(327)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierListOrMapContext is an interface to support dynamic dispatch.
type IIdentifierListOrMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListOrMapContext differentiates from other interfaces.
	IsIdentifierListOrMapContext()
}

type IdentifierListOrMapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListOrMapContext() *IdentifierListOrMapContext {
	var p = new(IdentifierListOrMapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_identifierListOrMap
	return p
}

func (*IdentifierListOrMapContext) IsIdentifierListOrMapContext() {}

func NewIdentifierListOrMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListOrMapContext {
	var p = new(IdentifierListOrMapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_identifierListOrMap

	return p
}

func (s *IdentifierListOrMapContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListOrMapContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *IdentifierListOrMapContext) AllIdentifierValue() []IIdentifierValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierValueContext)(nil)).Elem())
	var tst = make([]IIdentifierValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierValueContext)
		}
	}

	return tst
}

func (s *IdentifierListOrMapContext) IdentifierValue(i int) IIdentifierValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierValueContext)
}

func (s *IdentifierListOrMapContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *IdentifierListOrMapContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *IdentifierListOrMapContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *IdentifierListOrMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListOrMapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListOrMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterIdentifierListOrMap(s)
	}
}

func (s *IdentifierListOrMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitIdentifierListOrMap(s)
	}
}

func (p *ScssParser) IdentifierListOrMap() (localctx IIdentifierListOrMapContext) {
	localctx = NewIdentifierListOrMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ScssParserRULE_identifierListOrMap)
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
		p.Match(ScssParserLPAREN)
	}
	{
		p.SetState(331)
		p.IdentifierValue()
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserCOMMA {
		{
			p.SetState(332)
			p.Match(ScssParserCOMMA)
		}
		{
			p.SetState(333)
			p.IdentifierValue()
		}

		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(339)
		p.Match(ScssParserRPAREN)
	}

	return localctx
}

// IIdentifierValueContext is an interface to support dynamic dispatch.
type IIdentifierValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierValueContext differentiates from other interfaces.
	IsIdentifierValueContext()
}

type IdentifierValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierValueContext() *IdentifierValueContext {
	var p = new(IdentifierValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_identifierValue
	return p
}

func (*IdentifierValueContext) IsIdentifierValueContext() {}

func NewIdentifierValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierValueContext {
	var p = new(IdentifierValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_identifierValue

	return p
}

func (s *IdentifierValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierValueContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *IdentifierValueContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *IdentifierValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterIdentifierValue(s)
	}
}

func (s *IdentifierValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitIdentifierValue(s)
	}
}

func (p *ScssParser) IdentifierValue() (localctx IIdentifierValueContext) {
	localctx = NewIdentifierValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ScssParserRULE_identifierValue)
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
		p.SetState(341)
		p.Identifier()
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserCOLON {
		{
			p.SetState(342)
			p.Match(ScssParserCOLON)
		}
		{
			p.SetState(343)
			p.Values()
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

func (s *ImportDeclarationContext) MediaTypes() IMediaTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMediaTypesContext)
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
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ScssParserRULE_importDeclaration)
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
		p.Match(ScssParserIMPORT)
	}
	{
		p.SetState(347)
		p.ReferenceUrl()
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserIdentifier {
		{
			p.SetState(348)
			p.MediaTypes()
		}

	}
	{
		p.SetState(351)
		p.Match(ScssParserSEMI)
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
	localctx = NewReferenceUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ScssParserRULE_referenceUrl)

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

	p.SetState(357)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.Match(ScssParserStringLiteral)
		}

	case ScssParserUrlStart:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.Match(ScssParserUrlStart)
		}
		{
			p.SetState(355)
			p.Match(ScssParserUrl)
		}
		{
			p.SetState(356)
			p.Match(ScssParserUrlEnd)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMediaTypesContext is an interface to support dynamic dispatch.
type IMediaTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaTypesContext differentiates from other interfaces.
	IsMediaTypesContext()
}

type MediaTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaTypesContext() *MediaTypesContext {
	var p = new(MediaTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mediaTypes
	return p
}

func (*MediaTypesContext) IsMediaTypesContext() {}

func NewMediaTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaTypesContext {
	var p = new(MediaTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mediaTypes

	return p
}

func (s *MediaTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaTypesContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(ScssParserIdentifier)
}

func (s *MediaTypesContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, i)
}

func (s *MediaTypesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *MediaTypesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *MediaTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMediaTypes(s)
	}
}

func (s *MediaTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMediaTypes(s)
	}
}

func (p *ScssParser) MediaTypes() (localctx IMediaTypesContext) {
	localctx = NewMediaTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ScssParserRULE_mediaTypes)
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
		p.SetState(359)
		p.Match(ScssParserIdentifier)
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserCOMMA {
		{
			p.SetState(360)
			p.Match(ScssParserCOMMA)
		}
		{
			p.SetState(361)
			p.Match(ScssParserIdentifier)
		}

		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INestedContext is an interface to support dynamic dispatch.
type INestedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNestedContext differentiates from other interfaces.
	IsNestedContext()
}

type NestedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedContext() *NestedContext {
	var p = new(NestedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_nested
	return p
}

func (*NestedContext) IsNestedContext() {}

func NewNestedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedContext {
	var p = new(NestedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_nested

	return p
}

func (s *NestedContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedContext) AT() antlr.TerminalNode {
	return s.GetToken(ScssParserAT, 0)
}

func (s *NestedContext) Nest() INestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INestContext)
}

func (s *NestedContext) Selectors() ISelectorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorsContext)
}

func (s *NestedContext) BlockStart() antlr.TerminalNode {
	return s.GetToken(ScssParserBlockStart, 0)
}

func (s *NestedContext) Stylesheet() IStylesheetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStylesheetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStylesheetContext)
}

func (s *NestedContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(ScssParserBlockEnd, 0)
}

func (s *NestedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterNested(s)
	}
}

func (s *NestedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitNested(s)
	}
}

func (p *ScssParser) Nested() (localctx INestedContext) {
	localctx = NewNestedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ScssParserRULE_nested)

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
		p.SetState(367)
		p.Match(ScssParserAT)
	}
	{
		p.SetState(368)
		p.Nest()
	}
	{
		p.SetState(369)
		p.Selectors()
	}
	{
		p.SetState(370)
		p.Match(ScssParserBlockStart)
	}
	{
		p.SetState(371)
		p.Stylesheet()
	}
	{
		p.SetState(372)
		p.Match(ScssParserBlockEnd)
	}

	return localctx
}

// INestContext is an interface to support dynamic dispatch.
type INestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNestContext differentiates from other interfaces.
	IsNestContext()
}

type NestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestContext() *NestContext {
	var p = new(NestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_nest
	return p
}

func (*NestContext) IsNestContext() {}

func NewNestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestContext {
	var p = new(NestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_nest

	return p
}

func (s *NestContext) GetParser() antlr.Parser { return s.parser }

func (s *NestContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(ScssParserIdentifier)
}

func (s *NestContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, i)
}

func (s *NestContext) AND() antlr.TerminalNode {
	return s.GetToken(ScssParserAND, 0)
}

func (s *NestContext) AllPseudo() []IPseudoContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPseudoContext)(nil)).Elem())
	var tst = make([]IPseudoContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPseudoContext)
		}
	}

	return tst
}

func (s *NestContext) Pseudo(i int) IPseudoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudoContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPseudoContext)
}

func (s *NestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterNest(s)
	}
}

func (s *NestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitNest(s)
	}
}

func (p *ScssParser) Nest() (localctx INestContext) {
	localctx = NewNestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ScssParserRULE_nest)
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
		p.SetState(374)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ScssParserAND || _la == ScssParserIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(375)
				p.Match(ScssParserIdentifier)
			}

		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserCOLON || _la == ScssParserCOLONCOLON {
		{
			p.SetState(381)
			p.Pseudo()
		}

		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	localctx = NewRulesetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ScssParserRULE_ruleset)

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
		p.SetState(387)
		p.Selectors()
	}
	{
		p.SetState(388)
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

func (s *BlockContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *BlockContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *BlockContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(ScssParserSEMI)
}

func (s *BlockContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, i)
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
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ScssParserRULE_block)
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
		p.SetState(390)
		p.Match(ScssParserBlockStart)
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(395)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(391)
					p.Property()
				}
				{
					p.SetState(392)
					p.Match(ScssParserSEMI)
				}

			case 2:
				{
					p.SetState(394)
					p.Statement()
				}

			}

		}
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserInterpolationStart || _la == ScssParserIdentifier {
		{
			p.SetState(400)
			p.Property()
		}

	}
	{
		p.SetState(403)
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
	localctx = NewSelectorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ScssParserRULE_selectors)
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
		p.SetState(405)
		p.Selector()
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserCOMMA {
		{
			p.SetState(406)
			p.Match(ScssParserCOMMA)
		}
		{
			p.SetState(407)
			p.Selector()
		}

		p.SetState(412)
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

func (s *SelectorContext) AllSelectorPrefix() []ISelectorPrefixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectorPrefixContext)(nil)).Elem())
	var tst = make([]ISelectorPrefixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectorPrefixContext)
		}
	}

	return tst
}

func (s *SelectorContext) SelectorPrefix(i int) ISelectorPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorPrefixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectorPrefixContext)
}

func (s *SelectorContext) AllAttrib() []IAttribContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttribContext)(nil)).Elem())
	var tst = make([]IAttribContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttribContext)
		}
	}

	return tst
}

func (s *SelectorContext) Attrib(i int) IAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *SelectorContext) Pseudo() IPseudoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPseudoContext)
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
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ScssParserRULE_selector)
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

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserInterpolationStart)|(1<<ScssParserDOT)|(1<<ScssParserAND)|(1<<ScssParserHASH)|(1<<ScssParserTIMES))) != 0) || _la == ScssParserIdentifier {
		{
			p.SetState(413)
			p.Element()
		}

		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserGT)|(1<<ScssParserTIL)|(1<<ScssParserPLUS))) != 0 {
		{
			p.SetState(418)
			p.SelectorPrefix()
		}
		{
			p.SetState(419)
			p.Element()
		}

		p.SetState(425)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserLBRACK {
		{
			p.SetState(426)
			p.Attrib()
		}

		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserCOLON || _la == ScssParserCOLONCOLON {
		{
			p.SetState(432)
			p.Pseudo()
		}

	}

	return localctx
}

// ISelectorPrefixContext is an interface to support dynamic dispatch.
type ISelectorPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorPrefixContext differentiates from other interfaces.
	IsSelectorPrefixContext()
}

type SelectorPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorPrefixContext() *SelectorPrefixContext {
	var p = new(SelectorPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_selectorPrefix
	return p
}

func (*SelectorPrefixContext) IsSelectorPrefixContext() {}

func NewSelectorPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorPrefixContext {
	var p = new(SelectorPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_selectorPrefix

	return p
}

func (s *SelectorPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorPrefixContext) GT() antlr.TerminalNode {
	return s.GetToken(ScssParserGT, 0)
}

func (s *SelectorPrefixContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ScssParserPLUS, 0)
}

func (s *SelectorPrefixContext) TIL() antlr.TerminalNode {
	return s.GetToken(ScssParserTIL, 0)
}

func (s *SelectorPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterSelectorPrefix(s)
	}
}

func (s *SelectorPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitSelectorPrefix(s)
	}
}

func (p *ScssParser) SelectorPrefix() (localctx ISelectorPrefixContext) {
	localctx = NewSelectorPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ScssParserRULE_selectorPrefix)
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
		p.SetState(435)
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
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ScssParserRULE_element)

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

	p.SetState(444)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserInterpolationStart, ScssParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.Identifier()
		}

	case ScssParserHASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.Match(ScssParserHASH)
		}
		{
			p.SetState(439)
			p.Identifier()
		}

	case ScssParserDOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(440)
			p.Match(ScssParserDOT)
		}
		{
			p.SetState(441)
			p.Identifier()
		}

	case ScssParserAND:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(442)
			p.Match(ScssParserAND)
		}

	case ScssParserTIMES:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(443)
			p.Match(ScssParserTIMES)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *PseudoContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *PseudoContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *PseudoContext) COLONCOLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLONCOLON, 0)
}

func (s *PseudoContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
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
	localctx = NewPseudoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ScssParserRULE_pseudo)
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

	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(446)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ScssParserCOLON || _la == ScssParserCOLONCOLON) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(447)
			p.Match(ScssParserIdentifier)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(448)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ScssParserCOLON || _la == ScssParserCOLONCOLON) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(449)
			p.FunctionCall()
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
	localctx = NewAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ScssParserRULE_attrib)
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
		p.SetState(452)
		p.Match(ScssParserLBRACK)
	}
	{
		p.SetState(453)
		p.Match(ScssParserIdentifier)
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ScssParserEQ-33))|(1<<(ScssParserPIPE_EQ-33))|(1<<(ScssParserTILD_EQ-33)))) != 0 {
		{
			p.SetState(454)
			p.AttribRelate()
		}
		{
			p.SetState(455)
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
		p.SetState(459)
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

func (s *AttribRelateContext) TILD_EQ() antlr.TerminalNode {
	return s.GetToken(ScssParserTILD_EQ, 0)
}

func (s *AttribRelateContext) PIPE_EQ() antlr.TerminalNode {
	return s.GetToken(ScssParserPIPE_EQ, 0)
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
	localctx = NewAttribRelateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ScssParserRULE_attribRelate)
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
		p.SetState(461)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ScssParserEQ-33))|(1<<(ScssParserPIPE_EQ-33))|(1<<(ScssParserTILD_EQ-33)))) != 0) {
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
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ScssParserRULE_identifier)
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

	p.SetState(479)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(463)
			p.Match(ScssParserIdentifier)
		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScssParserInterpolationStartAfter || _la == ScssParserIdentifierAfter {
			{
				p.SetState(464)
				p.IdentifierPart()
			}

			p.SetState(469)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ScssParserInterpolationStart:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(470)
			p.Match(ScssParserInterpolationStart)
		}
		{
			p.SetState(471)
			p.IdentifierVariableName()
		}
		{
			p.SetState(472)
			p.Match(ScssParserBlockEnd)
		}
		p.SetState(476)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScssParserInterpolationStartAfter || _la == ScssParserIdentifierAfter {
			{
				p.SetState(473)
				p.IdentifierPart()
			}

			p.SetState(478)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	localctx = NewIdentifierPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ScssParserRULE_identifierPart)

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

	p.SetState(486)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserInterpolationStartAfter:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(481)
			p.Match(ScssParserInterpolationStartAfter)
		}
		{
			p.SetState(482)
			p.IdentifierVariableName()
		}
		{
			p.SetState(483)
			p.Match(ScssParserBlockEnd)
		}

	case ScssParserIdentifierAfter:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(485)
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
	localctx = NewIdentifierVariableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ScssParserRULE_identifierVariableName)
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
		p.SetState(488)
		p.Match(ScssParserDOLLAR)
	}
	{
		p.SetState(489)
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

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PropertyContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *PropertyContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *ScssParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ScssParserRULE_property)

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
		p.SetState(491)
		p.Identifier()
	}
	{
		p.SetState(492)
		p.Match(ScssParserCOLON)
	}
	{
		p.SetState(493)
		p.Values()
	}

	return localctx
}

// IValuesContext is an interface to support dynamic dispatch.
type IValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValuesContext differentiates from other interfaces.
	IsValuesContext()
}

type ValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesContext() *ValuesContext {
	var p = new(ValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_values
	return p
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_values

	return p
}

func (s *ValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesContext) AllCommandStatement() []ICommandStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem())
	var tst = make([]ICommandStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandStatementContext)
		}
	}

	return tst
}

func (s *ValuesContext) CommandStatement(i int) ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *ValuesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *ValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *ScssParser) Values() (localctx IValuesContext) {
	localctx = NewValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ScssParserRULE_values)

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
		p.SetState(495)
		p.CommandStatement()
	}
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(496)
				p.Match(ScssParserCOMMA)
			}
			{
				p.SetState(497)
				p.CommandStatement()
			}

		}
		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext())
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
	localctx = NewUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ScssParserRULE_url)

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
		p.SetState(503)
		p.Match(ScssParserUrlStart)
	}
	{
		p.SetState(504)
		p.Match(ScssParserUrl)
	}
	{
		p.SetState(505)
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
	localctx = NewMeasurementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ScssParserRULE_measurement)
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
		p.SetState(507)
		p.Match(ScssParserNumber)
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserUnit {
		{
			p.SetState(508)
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

func (s *FunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *FunctionCallContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
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
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ScssParserRULE_functionCall)
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
		p.SetState(511)
		p.Match(ScssParserIdentifier)
	}
	{
		p.SetState(512)
		p.Match(ScssParserLPAREN)
	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserDOLLAR)|(1<<ScssParserUrlStart))) != 0) || (((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(ScssParserIdentifier-50))|(1<<(ScssParserStringLiteral-50))|(1<<(ScssParserNumber-50))|(1<<(ScssParserColor-50)))) != 0) {
		{
			p.SetState(513)
			p.Values()
		}

	}
	{
		p.SetState(516)
		p.Match(ScssParserRPAREN)
	}

	return localctx
}
