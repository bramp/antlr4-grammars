// Generated from IRI.g4 by ANTLR 4.7.

package iri // IRI
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 65, 561,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 95, 10, 3, 3, 3, 3, 3, 5, 3,
	99, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 108, 10, 4,
	3, 5, 3, 5, 5, 5, 112, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 119,
	10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 124, 10, 7, 3, 7, 3, 7, 5, 7, 128, 10, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 137, 10, 8, 3, 9, 3, 9,
	3, 9, 5, 9, 142, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 147, 10, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 7, 10, 153, 10, 10, 12, 10, 14, 10, 156, 11, 10, 3, 11, 3,
	11, 3, 11, 5, 11, 161, 10, 11, 3, 12, 3, 12, 3, 12, 7, 12, 166, 10, 12,
	12, 12, 14, 12, 169, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13,
	176, 10, 13, 3, 14, 3, 14, 7, 14, 180, 10, 14, 12, 14, 14, 14, 183, 11,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 189, 10, 15, 12, 15, 14, 15, 192,
	11, 15, 5, 15, 194, 10, 15, 3, 16, 3, 16, 3, 16, 7, 16, 199, 10, 16, 12,
	16, 14, 16, 202, 11, 16, 3, 17, 3, 17, 3, 17, 7, 17, 207, 10, 17, 12, 17,
	14, 17, 210, 11, 17, 3, 18, 3, 18, 3, 19, 7, 19, 215, 10, 19, 12, 19, 14,
	19, 218, 11, 19, 3, 20, 6, 20, 221, 10, 20, 13, 20, 14, 20, 222, 3, 21,
	3, 21, 3, 21, 3, 21, 6, 21, 229, 10, 21, 13, 21, 14, 21, 230, 3, 22, 3,
	22, 3, 22, 3, 22, 5, 22, 237, 10, 22, 3, 23, 3, 23, 7, 23, 241, 10, 23,
	12, 23, 14, 23, 244, 11, 23, 3, 24, 3, 24, 7, 24, 248, 10, 24, 12, 24,
	14, 24, 251, 11, 24, 3, 25, 3, 25, 3, 25, 5, 25, 256, 10, 25, 3, 26, 3,
	26, 3, 26, 3, 26, 7, 26, 262, 10, 26, 12, 26, 14, 26, 265, 11, 26, 3, 27,
	7, 27, 268, 10, 27, 12, 27, 14, 27, 271, 11, 27, 3, 28, 3, 28, 3, 28, 5,
	28, 276, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 6, 29, 282, 10, 29, 13, 29,
	14, 29, 283, 3, 29, 3, 29, 3, 29, 3, 29, 6, 29, 290, 10, 29, 13, 29, 14,
	29, 291, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 322,
	10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 338, 10, 30, 3, 30, 5, 30, 341,
	10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 5, 30, 355, 10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 360,
	10, 30, 3, 30, 5, 30, 363, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 375, 10, 30, 3, 30, 3, 30, 3, 30,
	5, 30, 380, 10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 385, 10, 30, 3, 30, 5,
	30, 388, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	5, 30, 398, 10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 403, 10, 30, 3, 30, 3,
	30, 3, 30, 5, 30, 408, 10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 413, 10, 30,
	3, 30, 5, 30, 416, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 423,
	10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 428, 10, 30, 3, 30, 3, 30, 3, 30, 5,
	30, 433, 10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 438, 10, 30, 3, 30, 3, 30,
	3, 30, 5, 30, 443, 10, 30, 3, 30, 5, 30, 446, 10, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 5, 30, 453, 10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 458,
	10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 463, 10, 30, 3, 30, 3, 30, 3, 30, 5,
	30, 468, 10, 30, 3, 30, 3, 30, 3, 30, 5, 30, 473, 10, 30, 3, 30, 3, 30,
	3, 30, 5, 30, 478, 10, 30, 3, 30, 5, 30, 481, 10, 30, 3, 30, 5, 30, 484,
	10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 499, 10, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 5, 32, 506, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 530, 10, 34, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 5, 36, 539, 10, 36, 3, 37, 3, 37, 5,
	37, 543, 10, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41,
	5, 41, 553, 10, 41, 3, 42, 3, 42, 5, 42, 557, 10, 42, 3, 43, 3, 43, 3,
	43, 2, 2, 44, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 72, 74, 76, 78, 80, 82, 84, 2, 15, 4, 2, 42, 42, 65, 65, 4, 2, 4, 4,
	60, 61, 3, 2, 60, 61, 5, 2, 3, 3, 43, 43, 45, 47, 5, 2, 43, 43, 45, 45,
	55, 55, 3, 2, 5, 9, 3, 2, 5, 10, 4, 2, 43, 43, 45, 47, 4, 2, 42, 42, 60,
	65, 3, 2, 48, 58, 3, 2, 15, 40, 3, 2, 15, 20, 3, 2, 6, 14, 2, 623, 2, 86,
	3, 2, 2, 2, 4, 89, 3, 2, 2, 2, 6, 107, 3, 2, 2, 2, 8, 111, 3, 2, 2, 2,
	10, 113, 3, 2, 2, 2, 12, 120, 3, 2, 2, 2, 14, 136, 3, 2, 2, 2, 16, 141,
	3, 2, 2, 2, 18, 154, 3, 2, 2, 2, 20, 160, 3, 2, 2, 2, 22, 167, 3, 2, 2,
	2, 24, 175, 3, 2, 2, 2, 26, 181, 3, 2, 2, 2, 28, 184, 3, 2, 2, 2, 30, 195,
	3, 2, 2, 2, 32, 203, 3, 2, 2, 2, 34, 211, 3, 2, 2, 2, 36, 216, 3, 2, 2,
	2, 38, 220, 3, 2, 2, 2, 40, 228, 3, 2, 2, 2, 42, 236, 3, 2, 2, 2, 44, 242,
	3, 2, 2, 2, 46, 249, 3, 2, 2, 2, 48, 255, 3, 2, 2, 2, 50, 257, 3, 2, 2,
	2, 52, 269, 3, 2, 2, 2, 54, 272, 3, 2, 2, 2, 56, 279, 3, 2, 2, 2, 58, 483,
	3, 2, 2, 2, 60, 498, 3, 2, 2, 2, 62, 505, 3, 2, 2, 2, 64, 507, 3, 2, 2,
	2, 66, 529, 3, 2, 2, 2, 68, 531, 3, 2, 2, 2, 70, 538, 3, 2, 2, 2, 72, 542,
	3, 2, 2, 2, 74, 544, 3, 2, 2, 2, 76, 546, 3, 2, 2, 2, 78, 548, 3, 2, 2,
	2, 80, 552, 3, 2, 2, 2, 82, 556, 3, 2, 2, 2, 84, 558, 3, 2, 2, 2, 86, 87,
	5, 4, 3, 2, 87, 88, 7, 2, 2, 3, 88, 3, 3, 2, 2, 2, 89, 90, 5, 50, 26, 2,
	90, 91, 7, 42, 2, 2, 91, 94, 5, 6, 4, 2, 92, 93, 7, 61, 2, 2, 93, 95, 5,
	44, 23, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2,
	96, 97, 7, 62, 2, 2, 97, 99, 5, 46, 24, 2, 98, 96, 3, 2, 2, 2, 98, 99,
	3, 2, 2, 2, 99, 5, 3, 2, 2, 2, 100, 101, 7, 59, 2, 2, 101, 102, 5, 16,
	9, 2, 102, 103, 5, 26, 14, 2, 103, 108, 3, 2, 2, 2, 104, 108, 5, 28, 15,
	2, 105, 108, 5, 32, 17, 2, 106, 108, 5, 34, 18, 2, 107, 100, 3, 2, 2, 2,
	107, 104, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108,
	7, 3, 2, 2, 2, 109, 112, 5, 4, 3, 2, 110, 112, 5, 12, 7, 2, 111, 109, 3,
	2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 9, 3, 2, 2, 2, 113, 114, 5, 50, 26,
	2, 114, 115, 7, 42, 2, 2, 115, 118, 5, 6, 4, 2, 116, 117, 7, 61, 2, 2,
	117, 119, 5, 44, 23, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119,
	11, 3, 2, 2, 2, 120, 123, 5, 14, 8, 2, 121, 122, 7, 61, 2, 2, 122, 124,
	5, 44, 23, 2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 127, 3,
	2, 2, 2, 125, 126, 7, 62, 2, 2, 126, 128, 5, 46, 24, 2, 127, 125, 3, 2,
	2, 2, 127, 128, 3, 2, 2, 2, 128, 13, 3, 2, 2, 2, 129, 130, 7, 59, 2, 2,
	130, 131, 5, 16, 9, 2, 131, 132, 5, 26, 14, 2, 132, 137, 3, 2, 2, 2, 133,
	137, 5, 28, 15, 2, 134, 137, 5, 30, 16, 2, 135, 137, 5, 34, 18, 2, 136,
	129, 3, 2, 2, 2, 136, 133, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 135,
	3, 2, 2, 2, 137, 15, 3, 2, 2, 2, 138, 139, 5, 18, 10, 2, 139, 140, 7, 65,
	2, 2, 140, 142, 3, 2, 2, 2, 141, 138, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2,
	142, 143, 3, 2, 2, 2, 143, 146, 5, 20, 11, 2, 144, 145, 7, 42, 2, 2, 145,
	147, 5, 52, 27, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 17,
	3, 2, 2, 2, 148, 153, 5, 48, 25, 2, 149, 153, 5, 68, 35, 2, 150, 153, 5,
	76, 39, 2, 151, 153, 7, 42, 2, 2, 152, 148, 3, 2, 2, 2, 152, 149, 3, 2,
	2, 2, 152, 150, 3, 2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 156, 3, 2, 2, 2,
	154, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 19, 3, 2, 2, 2, 156, 154,
	3, 2, 2, 2, 157, 161, 5, 54, 28, 2, 158, 161, 5, 64, 33, 2, 159, 161, 5,
	22, 12, 2, 160, 157, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 159, 3, 2,
	2, 2, 161, 21, 3, 2, 2, 2, 162, 166, 5, 48, 25, 2, 163, 166, 5, 68, 35,
	2, 164, 166, 5, 76, 39, 2, 165, 162, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2,
	165, 164, 3, 2, 2, 2, 166, 169, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167,
	168, 3, 2, 2, 2, 168, 23, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 170, 176, 5,
	26, 14, 2, 171, 176, 5, 28, 15, 2, 172, 176, 5, 30, 16, 2, 173, 176, 5,
	32, 17, 2, 174, 176, 5, 34, 18, 2, 175, 170, 3, 2, 2, 2, 175, 171, 3, 2,
	2, 2, 175, 172, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2,
	176, 25, 3, 2, 2, 2, 177, 178, 7, 60, 2, 2, 178, 180, 5, 36, 19, 2, 179,
	177, 3, 2, 2, 2, 180, 183, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182,
	3, 2, 2, 2, 182, 27, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184, 193, 7, 60,
	2, 2, 185, 190, 5, 38, 20, 2, 186, 187, 7, 60, 2, 2, 187, 189, 5, 36, 19,
	2, 188, 186, 3, 2, 2, 2, 189, 192, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190,
	191, 3, 2, 2, 2, 191, 194, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 193, 185,
	3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 29, 3, 2, 2, 2, 195, 200, 5, 40,
	21, 2, 196, 197, 7, 60, 2, 2, 197, 199, 5, 36, 19, 2, 198, 196, 3, 2, 2,
	2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201,
	31, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 208, 5, 38, 20, 2, 204, 205,
	7, 60, 2, 2, 205, 207, 5, 36, 19, 2, 206, 204, 3, 2, 2, 2, 207, 210, 3,
	2, 2, 2, 208, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 33, 3, 2, 2,
	2, 210, 208, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 35, 3, 2, 2, 2, 213,
	215, 5, 42, 22, 2, 214, 213, 3, 2, 2, 2, 215, 218, 3, 2, 2, 2, 216, 214,
	3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 37, 3, 2, 2, 2, 218, 216, 3, 2,
	2, 2, 219, 221, 5, 42, 22, 2, 220, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2,
	2, 222, 220, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 39, 3, 2, 2, 2, 224,
	229, 5, 48, 25, 2, 225, 229, 5, 68, 35, 2, 226, 229, 5, 76, 39, 2, 227,
	229, 7, 65, 2, 2, 228, 224, 3, 2, 2, 2, 228, 225, 3, 2, 2, 2, 228, 226,
	3, 2, 2, 2, 228, 227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 228, 3, 2,
	2, 2, 230, 231, 3, 2, 2, 2, 231, 41, 3, 2, 2, 2, 232, 237, 5, 48, 25, 2,
	233, 237, 5, 68, 35, 2, 234, 237, 5, 76, 39, 2, 235, 237, 9, 2, 2, 2, 236,
	232, 3, 2, 2, 2, 236, 233, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236, 235,
	3, 2, 2, 2, 237, 43, 3, 2, 2, 2, 238, 241, 5, 42, 22, 2, 239, 241, 9, 3,
	2, 2, 240, 238, 3, 2, 2, 2, 240, 239, 3, 2, 2, 2, 241, 244, 3, 2, 2, 2,
	242, 240, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 45, 3, 2, 2, 2, 244, 242,
	3, 2, 2, 2, 245, 248, 5, 42, 22, 2, 246, 248, 9, 4, 2, 2, 247, 245, 3,
	2, 2, 2, 247, 246, 3, 2, 2, 2, 248, 251, 3, 2, 2, 2, 249, 247, 3, 2, 2,
	2, 249, 250, 3, 2, 2, 2, 250, 47, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 252,
	256, 5, 78, 40, 2, 253, 256, 5, 82, 42, 2, 254, 256, 9, 5, 2, 2, 255, 252,
	3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255, 254, 3, 2, 2, 2, 256, 49, 3, 2,
	2, 2, 257, 263, 5, 78, 40, 2, 258, 262, 5, 78, 40, 2, 259, 262, 5, 82,
	42, 2, 260, 262, 9, 6, 2, 2, 261, 258, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2,
	261, 260, 3, 2, 2, 2, 262, 265, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 263,
	264, 3, 2, 2, 2, 264, 51, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 266, 268, 5,
	82, 42, 2, 267, 266, 3, 2, 2, 2, 268, 271, 3, 2, 2, 2, 269, 267, 3, 2,
	2, 2, 269, 270, 3, 2, 2, 2, 270, 53, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2,
	272, 275, 7, 63, 2, 2, 273, 276, 5, 58, 30, 2, 274, 276, 5, 56, 29, 2,
	275, 273, 3, 2, 2, 2, 275, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277,
	278, 7, 64, 2, 2, 278, 55, 3, 2, 2, 2, 279, 281, 7, 36, 2, 2, 280, 282,
	5, 80, 41, 2, 281, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 281, 3,
	2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 289, 7, 43, 2,
	2, 286, 290, 5, 70, 36, 2, 287, 290, 5, 76, 39, 2, 288, 290, 7, 42, 2,
	2, 289, 286, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289, 288, 3, 2, 2, 2, 290,
	291, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 57, 3,
	2, 2, 2, 293, 294, 5, 60, 31, 2, 294, 295, 7, 42, 2, 2, 295, 296, 5, 60,
	31, 2, 296, 297, 7, 42, 2, 2, 297, 298, 5, 60, 31, 2, 298, 299, 7, 42,
	2, 2, 299, 300, 5, 60, 31, 2, 300, 301, 7, 42, 2, 2, 301, 302, 5, 60, 31,
	2, 302, 303, 7, 42, 2, 2, 303, 304, 5, 60, 31, 2, 304, 305, 7, 42, 2, 2,
	305, 306, 5, 62, 32, 2, 306, 484, 3, 2, 2, 2, 307, 308, 7, 41, 2, 2, 308,
	309, 5, 60, 31, 2, 309, 310, 7, 42, 2, 2, 310, 311, 5, 60, 31, 2, 311,
	312, 7, 42, 2, 2, 312, 313, 5, 60, 31, 2, 313, 314, 7, 42, 2, 2, 314, 315,
	5, 60, 31, 2, 315, 316, 7, 42, 2, 2, 316, 317, 5, 60, 31, 2, 317, 318,
	7, 42, 2, 2, 318, 319, 5, 62, 32, 2, 319, 484, 3, 2, 2, 2, 320, 322, 5,
	60, 31, 2, 321, 320, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 3, 2,
	2, 2, 323, 324, 7, 41, 2, 2, 324, 325, 5, 60, 31, 2, 325, 326, 7, 42, 2,
	2, 326, 327, 5, 60, 31, 2, 327, 328, 7, 42, 2, 2, 328, 329, 5, 60, 31,
	2, 329, 330, 7, 42, 2, 2, 330, 331, 5, 60, 31, 2, 331, 332, 7, 42, 2, 2,
	332, 333, 5, 62, 32, 2, 333, 484, 3, 2, 2, 2, 334, 335, 5, 60, 31, 2, 335,
	336, 7, 42, 2, 2, 336, 338, 3, 2, 2, 2, 337, 334, 3, 2, 2, 2, 337, 338,
	3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 341, 5, 60, 31, 2, 340, 337, 3,
	2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342, 343, 7, 41, 2,
	2, 343, 344, 5, 60, 31, 2, 344, 345, 7, 42, 2, 2, 345, 346, 5, 60, 31,
	2, 346, 347, 7, 42, 2, 2, 347, 348, 5, 60, 31, 2, 348, 349, 7, 42, 2, 2,
	349, 350, 5, 62, 32, 2, 350, 484, 3, 2, 2, 2, 351, 352, 5, 60, 31, 2, 352,
	353, 7, 42, 2, 2, 353, 355, 3, 2, 2, 2, 354, 351, 3, 2, 2, 2, 354, 355,
	3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 357, 5, 60, 31, 2, 357, 358, 7,
	42, 2, 2, 358, 360, 3, 2, 2, 2, 359, 354, 3, 2, 2, 2, 359, 360, 3, 2, 2,
	2, 360, 361, 3, 2, 2, 2, 361, 363, 5, 60, 31, 2, 362, 359, 3, 2, 2, 2,
	362, 363, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 365, 7, 41, 2, 2, 365,
	366, 5, 60, 31, 2, 366, 367, 7, 42, 2, 2, 367, 368, 5, 60, 31, 2, 368,
	369, 7, 42, 2, 2, 369, 370, 5, 62, 32, 2, 370, 484, 3, 2, 2, 2, 371, 372,
	5, 60, 31, 2, 372, 373, 7, 42, 2, 2, 373, 375, 3, 2, 2, 2, 374, 371, 3,
	2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 377, 5, 60, 31,
	2, 377, 378, 7, 42, 2, 2, 378, 380, 3, 2, 2, 2, 379, 374, 3, 2, 2, 2, 379,
	380, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 382, 5, 60, 31, 2, 382, 383,
	7, 42, 2, 2, 383, 385, 3, 2, 2, 2, 384, 379, 3, 2, 2, 2, 384, 385, 3, 2,
	2, 2, 385, 386, 3, 2, 2, 2, 386, 388, 5, 60, 31, 2, 387, 384, 3, 2, 2,
	2, 387, 388, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389, 390, 7, 41, 2, 2, 390,
	391, 5, 60, 31, 2, 391, 392, 7, 42, 2, 2, 392, 393, 5, 62, 32, 2, 393,
	484, 3, 2, 2, 2, 394, 395, 5, 60, 31, 2, 395, 396, 7, 42, 2, 2, 396, 398,
	3, 2, 2, 2, 397, 394, 3, 2, 2, 2, 397, 398, 3, 2, 2, 2, 398, 399, 3, 2,
	2, 2, 399, 400, 5, 60, 31, 2, 400, 401, 7, 42, 2, 2, 401, 403, 3, 2, 2,
	2, 402, 397, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404,
	405, 5, 60, 31, 2, 405, 406, 7, 42, 2, 2, 406, 408, 3, 2, 2, 2, 407, 402,
	3, 2, 2, 2, 407, 408, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 410, 5, 60,
	31, 2, 410, 411, 7, 42, 2, 2, 411, 413, 3, 2, 2, 2, 412, 407, 3, 2, 2,
	2, 412, 413, 3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 416, 5, 60, 31, 2,
	415, 412, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417,
	418, 7, 41, 2, 2, 418, 484, 5, 62, 32, 2, 419, 420, 5, 60, 31, 2, 420,
	421, 7, 42, 2, 2, 421, 423, 3, 2, 2, 2, 422, 419, 3, 2, 2, 2, 422, 423,
	3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 425, 5, 60, 31, 2, 425, 426, 7,
	42, 2, 2, 426, 428, 3, 2, 2, 2, 427, 422, 3, 2, 2, 2, 427, 428, 3, 2, 2,
	2, 428, 429, 3, 2, 2, 2, 429, 430, 5, 60, 31, 2, 430, 431, 7, 42, 2, 2,
	431, 433, 3, 2, 2, 2, 432, 427, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433,
	434, 3, 2, 2, 2, 434, 435, 5, 60, 31, 2, 435, 436, 7, 42, 2, 2, 436, 438,
	3, 2, 2, 2, 437, 432, 3, 2, 2, 2, 437, 438, 3, 2, 2, 2, 438, 439, 3, 2,
	2, 2, 439, 440, 5, 60, 31, 2, 440, 441, 7, 42, 2, 2, 441, 443, 3, 2, 2,
	2, 442, 437, 3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444,
	446, 5, 60, 31, 2, 445, 442, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2, 446, 447,
	3, 2, 2, 2, 447, 448, 7, 41, 2, 2, 448, 484, 5, 60, 31, 2, 449, 450, 5,
	60, 31, 2, 450, 451, 7, 42, 2, 2, 451, 453, 3, 2, 2, 2, 452, 449, 3, 2,
	2, 2, 452, 453, 3, 2, 2, 2, 453, 454, 3, 2, 2, 2, 454, 455, 5, 60, 31,
	2, 455, 456, 7, 42, 2, 2, 456, 458, 3, 2, 2, 2, 457, 452, 3, 2, 2, 2, 457,
	458, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 460, 5, 60, 31, 2, 460, 461,
	7, 42, 2, 2, 461, 463, 3, 2, 2, 2, 462, 457, 3, 2, 2, 2, 462, 463, 3, 2,
	2, 2, 463, 464, 3, 2, 2, 2, 464, 465, 5, 60, 31, 2, 465, 466, 7, 42, 2,
	2, 466, 468, 3, 2, 2, 2, 467, 462, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 468,
	469, 3, 2, 2, 2, 469, 470, 5, 60, 31, 2, 470, 471, 7, 42, 2, 2, 471, 473,
	3, 2, 2, 2, 472, 467, 3, 2, 2, 2, 472, 473, 3, 2, 2, 2, 473, 474, 3, 2,
	2, 2, 474, 475, 5, 60, 31, 2, 475, 476, 7, 42, 2, 2, 476, 478, 3, 2, 2,
	2, 477, 472, 3, 2, 2, 2, 477, 478, 3, 2, 2, 2, 478, 479, 3, 2, 2, 2, 479,
	481, 5, 60, 31, 2, 480, 477, 3, 2, 2, 2, 480, 481, 3, 2, 2, 2, 481, 482,
	3, 2, 2, 2, 482, 484, 7, 41, 2, 2, 483, 293, 3, 2, 2, 2, 483, 307, 3, 2,
	2, 2, 483, 321, 3, 2, 2, 2, 483, 340, 3, 2, 2, 2, 483, 362, 3, 2, 2, 2,
	483, 387, 3, 2, 2, 2, 483, 415, 3, 2, 2, 2, 483, 445, 3, 2, 2, 2, 483,
	480, 3, 2, 2, 2, 484, 59, 3, 2, 2, 2, 485, 486, 5, 80, 41, 2, 486, 487,
	5, 80, 41, 2, 487, 488, 5, 80, 41, 2, 488, 489, 5, 80, 41, 2, 489, 499,
	3, 2, 2, 2, 490, 491, 5, 80, 41, 2, 491, 492, 5, 80, 41, 2, 492, 493, 5,
	80, 41, 2, 493, 499, 3, 2, 2, 2, 494, 495, 5, 80, 41, 2, 495, 496, 5, 80,
	41, 2, 496, 499, 3, 2, 2, 2, 497, 499, 5, 80, 41, 2, 498, 485, 3, 2, 2,
	2, 498, 490, 3, 2, 2, 2, 498, 494, 3, 2, 2, 2, 498, 497, 3, 2, 2, 2, 499,
	61, 3, 2, 2, 2, 500, 501, 5, 60, 31, 2, 501, 502, 7, 42, 2, 2, 502, 503,
	5, 60, 31, 2, 503, 506, 3, 2, 2, 2, 504, 506, 5, 64, 33, 2, 505, 500, 3,
	2, 2, 2, 505, 504, 3, 2, 2, 2, 506, 63, 3, 2, 2, 2, 507, 508, 5, 66, 34,
	2, 508, 509, 7, 43, 2, 2, 509, 510, 5, 66, 34, 2, 510, 511, 7, 43, 2, 2,
	511, 512, 5, 66, 34, 2, 512, 513, 7, 43, 2, 2, 513, 514, 5, 66, 34, 2,
	514, 65, 3, 2, 2, 2, 515, 530, 5, 82, 42, 2, 516, 517, 5, 84, 43, 2, 517,
	518, 5, 82, 42, 2, 518, 530, 3, 2, 2, 2, 519, 520, 7, 6, 2, 2, 520, 521,
	5, 82, 42, 2, 521, 522, 5, 82, 42, 2, 522, 530, 3, 2, 2, 2, 523, 524, 7,
	7, 2, 2, 524, 525, 9, 7, 2, 2, 525, 530, 5, 82, 42, 2, 526, 527, 7, 7,
	2, 2, 527, 528, 7, 10, 2, 2, 528, 530, 9, 8, 2, 2, 529, 515, 3, 2, 2, 2,
	529, 516, 3, 2, 2, 2, 529, 519, 3, 2, 2, 2, 529, 523, 3, 2, 2, 2, 529,
	526, 3, 2, 2, 2, 530, 67, 3, 2, 2, 2, 531, 532, 7, 44, 2, 2, 532, 533,
	5, 80, 41, 2, 533, 534, 5, 80, 41, 2, 534, 69, 3, 2, 2, 2, 535, 539, 5,
	78, 40, 2, 536, 539, 5, 82, 42, 2, 537, 539, 9, 9, 2, 2, 538, 535, 3, 2,
	2, 2, 538, 536, 3, 2, 2, 2, 538, 537, 3, 2, 2, 2, 539, 71, 3, 2, 2, 2,
	540, 543, 5, 74, 38, 2, 541, 543, 5, 76, 39, 2, 542, 540, 3, 2, 2, 2, 542,
	541, 3, 2, 2, 2, 543, 73, 3, 2, 2, 2, 544, 545, 9, 10, 2, 2, 545, 75, 3,
	2, 2, 2, 546, 547, 9, 11, 2, 2, 547, 77, 3, 2, 2, 2, 548, 549, 9, 12, 2,
	2, 549, 79, 3, 2, 2, 2, 550, 553, 5, 82, 42, 2, 551, 553, 9, 13, 2, 2,
	552, 550, 3, 2, 2, 2, 552, 551, 3, 2, 2, 2, 553, 81, 3, 2, 2, 2, 554, 557,
	7, 5, 2, 2, 555, 557, 5, 84, 43, 2, 556, 554, 3, 2, 2, 2, 556, 555, 3,
	2, 2, 2, 557, 83, 3, 2, 2, 2, 558, 559, 9, 14, 2, 2, 559, 85, 3, 2, 2,
	2, 76, 94, 98, 107, 111, 118, 123, 127, 136, 141, 146, 152, 154, 160, 165,
	167, 175, 181, 190, 193, 200, 208, 216, 222, 228, 230, 236, 240, 242, 247,
	249, 255, 261, 263, 269, 275, 283, 289, 291, 321, 337, 340, 354, 359, 362,
	374, 379, 384, 387, 397, 402, 407, 412, 415, 422, 427, 432, 437, 442, 445,
	452, 457, 462, 467, 472, 477, 480, 483, 498, 505, 529, 538, 542, 552, 556,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "'0'", "'1'", "'2'", "'3'", "'4'", "'5'", "'6'", "'7'", "'8'",
	"'9'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "'::'", "':'", "'.'", "'%'", "'-'",
	"'~'", "'_'", "'!'", "'$'", "'&'", "'''", "'('", "')'", "'*'", "'+'", "','",
	"';'", "'='", "'//'", "'/'", "'?'", "'#'", "'['", "']'", "'@'",
}
var symbolicNames = []string{
	"", "UCSCHAR", "IPRIVATE", "D0", "D1", "D2", "D3", "D4", "D5", "D6", "D7",
	"D8", "D9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "COL2",
	"COL", "DOT", "PERCENT", "HYPHEN", "TILDE", "USCORE", "EXCL", "DOLLAR",
	"AMP", "SQUOTE", "OPAREN", "CPAREN", "STAR", "PLUS", "COMMA", "SCOL", "EQUALS",
	"FSLASH2", "FSLASH", "QMARK", "HASH", "OBRACK", "CBRACK", "AT",
}

var ruleNames = []string{
	"parse", "iri", "ihier_part", "iri_reference", "absolute_iri", "irelative_ref",
	"irelative_part", "iauthority", "iuserinfo", "ihost", "ireg_name", "ipath",
	"ipath_abempty", "ipath_absolute", "ipath_noscheme", "ipath_rootless",
	"ipath_empty", "isegment", "isegment_nz", "isegment_nz_nc", "ipchar", "iquery",
	"ifragment", "iunreserved", "scheme", "port", "ip_literal", "ip_v_future",
	"ip_v6_address", "h16", "ls32", "ip_v4_address", "dec_octet", "pct_encoded",
	"unreserved", "reserved", "gen_delims", "sub_delims", "alpha", "hexdig",
	"digit", "non_zero_digit",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type IRIParser struct {
	*antlr.BaseParser
}

func NewIRIParser(input antlr.TokenStream) *IRIParser {
	this := new(IRIParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "IRI.g4"

	return this
}

// IRIParser tokens.
const (
	IRIParserEOF      = antlr.TokenEOF
	IRIParserUCSCHAR  = 1
	IRIParserIPRIVATE = 2
	IRIParserD0       = 3
	IRIParserD1       = 4
	IRIParserD2       = 5
	IRIParserD3       = 6
	IRIParserD4       = 7
	IRIParserD5       = 8
	IRIParserD6       = 9
	IRIParserD7       = 10
	IRIParserD8       = 11
	IRIParserD9       = 12
	IRIParserA        = 13
	IRIParserB        = 14
	IRIParserC        = 15
	IRIParserD        = 16
	IRIParserE        = 17
	IRIParserF        = 18
	IRIParserG        = 19
	IRIParserH        = 20
	IRIParserI        = 21
	IRIParserJ        = 22
	IRIParserK        = 23
	IRIParserL        = 24
	IRIParserM        = 25
	IRIParserN        = 26
	IRIParserO        = 27
	IRIParserP        = 28
	IRIParserQ        = 29
	IRIParserR        = 30
	IRIParserS        = 31
	IRIParserT        = 32
	IRIParserU        = 33
	IRIParserV        = 34
	IRIParserW        = 35
	IRIParserX        = 36
	IRIParserY        = 37
	IRIParserZ        = 38
	IRIParserCOL2     = 39
	IRIParserCOL      = 40
	IRIParserDOT      = 41
	IRIParserPERCENT  = 42
	IRIParserHYPHEN   = 43
	IRIParserTILDE    = 44
	IRIParserUSCORE   = 45
	IRIParserEXCL     = 46
	IRIParserDOLLAR   = 47
	IRIParserAMP      = 48
	IRIParserSQUOTE   = 49
	IRIParserOPAREN   = 50
	IRIParserCPAREN   = 51
	IRIParserSTAR     = 52
	IRIParserPLUS     = 53
	IRIParserCOMMA    = 54
	IRIParserSCOL     = 55
	IRIParserEQUALS   = 56
	IRIParserFSLASH2  = 57
	IRIParserFSLASH   = 58
	IRIParserQMARK    = 59
	IRIParserHASH     = 60
	IRIParserOBRACK   = 61
	IRIParserCBRACK   = 62
	IRIParserAT       = 63
)

// IRIParser rules.
const (
	IRIParserRULE_parse          = 0
	IRIParserRULE_iri            = 1
	IRIParserRULE_ihier_part     = 2
	IRIParserRULE_iri_reference  = 3
	IRIParserRULE_absolute_iri   = 4
	IRIParserRULE_irelative_ref  = 5
	IRIParserRULE_irelative_part = 6
	IRIParserRULE_iauthority     = 7
	IRIParserRULE_iuserinfo      = 8
	IRIParserRULE_ihost          = 9
	IRIParserRULE_ireg_name      = 10
	IRIParserRULE_ipath          = 11
	IRIParserRULE_ipath_abempty  = 12
	IRIParserRULE_ipath_absolute = 13
	IRIParserRULE_ipath_noscheme = 14
	IRIParserRULE_ipath_rootless = 15
	IRIParserRULE_ipath_empty    = 16
	IRIParserRULE_isegment       = 17
	IRIParserRULE_isegment_nz    = 18
	IRIParserRULE_isegment_nz_nc = 19
	IRIParserRULE_ipchar         = 20
	IRIParserRULE_iquery         = 21
	IRIParserRULE_ifragment      = 22
	IRIParserRULE_iunreserved    = 23
	IRIParserRULE_scheme         = 24
	IRIParserRULE_port           = 25
	IRIParserRULE_ip_literal     = 26
	IRIParserRULE_ip_v_future    = 27
	IRIParserRULE_ip_v6_address  = 28
	IRIParserRULE_h16            = 29
	IRIParserRULE_ls32           = 30
	IRIParserRULE_ip_v4_address  = 31
	IRIParserRULE_dec_octet      = 32
	IRIParserRULE_pct_encoded    = 33
	IRIParserRULE_unreserved     = 34
	IRIParserRULE_reserved       = 35
	IRIParserRULE_gen_delims     = 36
	IRIParserRULE_sub_delims     = 37
	IRIParserRULE_alpha          = 38
	IRIParserRULE_hexdig         = 39
	IRIParserRULE_digit          = 40
	IRIParserRULE_non_zero_digit = 41
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(IRIParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IRIParserRULE_parse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(84)
		p.Iri()
	}
	{
		p.SetState(85)
		p.Match(IRIParserEOF)
	}

	return localctx
}

// IIriContext is an interface to support dynamic dispatch.
type IIriContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIriContext differentiates from other interfaces.
	IsIriContext()
}

type IriContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIriContext() *IriContext {
	var p = new(IriContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_iri
	return p
}

func (*IriContext) IsIriContext() {}

func NewIriContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriContext {
	var p = new(IriContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_iri

	return p
}

func (s *IriContext) GetParser() antlr.Parser { return s.parser }

func (s *IriContext) Scheme() ISchemeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemeContext)
}

func (s *IriContext) Ihier_part() IIhier_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIhier_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIhier_partContext)
}

func (s *IriContext) Iquery() IIqueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIqueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIqueryContext)
}

func (s *IriContext) Ifragment() IIfragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfragmentContext)
}

func (s *IriContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IriContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IriContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIri(s)
	}
}

func (s *IriContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIri(s)
	}
}

func (s *IriContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIri(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Iri() (localctx IIriContext) {
	localctx = NewIriContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IRIParserRULE_iri)
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
		p.SetState(87)
		p.Scheme()
	}
	{
		p.SetState(88)
		p.Match(IRIParserCOL)
	}
	{
		p.SetState(89)
		p.Ihier_part()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRIParserQMARK {
		{
			p.SetState(90)
			p.Match(IRIParserQMARK)
		}
		{
			p.SetState(91)
			p.Iquery()
		}

	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRIParserHASH {
		{
			p.SetState(94)
			p.Match(IRIParserHASH)
		}
		{
			p.SetState(95)
			p.Ifragment()
		}

	}

	return localctx
}

// IIhier_partContext is an interface to support dynamic dispatch.
type IIhier_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIhier_partContext differentiates from other interfaces.
	IsIhier_partContext()
}

type Ihier_partContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIhier_partContext() *Ihier_partContext {
	var p = new(Ihier_partContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ihier_part
	return p
}

func (*Ihier_partContext) IsIhier_partContext() {}

func NewIhier_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ihier_partContext {
	var p = new(Ihier_partContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ihier_part

	return p
}

func (s *Ihier_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Ihier_partContext) Iauthority() IIauthorityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIauthorityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIauthorityContext)
}

func (s *Ihier_partContext) Ipath_abempty() IIpath_abemptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_abemptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_abemptyContext)
}

func (s *Ihier_partContext) Ipath_absolute() IIpath_absoluteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_absoluteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_absoluteContext)
}

func (s *Ihier_partContext) Ipath_rootless() IIpath_rootlessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_rootlessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_rootlessContext)
}

func (s *Ihier_partContext) Ipath_empty() IIpath_emptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_emptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_emptyContext)
}

func (s *Ihier_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ihier_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ihier_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIhier_part(s)
	}
}

func (s *Ihier_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIhier_part(s)
	}
}

func (s *Ihier_partContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIhier_part(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ihier_part() (localctx IIhier_partContext) {
	localctx = NewIhier_partContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, IRIParserRULE_ihier_part)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRIParserFSLASH2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(IRIParserFSLASH2)
		}
		{
			p.SetState(99)
			p.Iauthority()
		}
		{
			p.SetState(100)
			p.Ipath_abempty()
		}

	case IRIParserFSLASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Ipath_absolute()
		}

	case IRIParserUCSCHAR, IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ, IRIParserCOL, IRIParserDOT, IRIParserPERCENT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE, IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS, IRIParserAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.Ipath_rootless()
		}

	case IRIParserEOF, IRIParserQMARK, IRIParserHASH:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(104)
			p.Ipath_empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIri_referenceContext is an interface to support dynamic dispatch.
type IIri_referenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIri_referenceContext differentiates from other interfaces.
	IsIri_referenceContext()
}

type Iri_referenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIri_referenceContext() *Iri_referenceContext {
	var p = new(Iri_referenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_iri_reference
	return p
}

func (*Iri_referenceContext) IsIri_referenceContext() {}

func NewIri_referenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iri_referenceContext {
	var p = new(Iri_referenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_iri_reference

	return p
}

func (s *Iri_referenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Iri_referenceContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *Iri_referenceContext) Irelative_ref() IIrelative_refContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIrelative_refContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIrelative_refContext)
}

func (s *Iri_referenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iri_referenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iri_referenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIri_reference(s)
	}
}

func (s *Iri_referenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIri_reference(s)
	}
}

func (s *Iri_referenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIri_reference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Iri_reference() (localctx IIri_referenceContext) {
	localctx = NewIri_referenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, IRIParserRULE_iri_reference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Iri()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Irelative_ref()
		}

	}

	return localctx
}

// IAbsolute_iriContext is an interface to support dynamic dispatch.
type IAbsolute_iriContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbsolute_iriContext differentiates from other interfaces.
	IsAbsolute_iriContext()
}

type Absolute_iriContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsolute_iriContext() *Absolute_iriContext {
	var p = new(Absolute_iriContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_absolute_iri
	return p
}

func (*Absolute_iriContext) IsAbsolute_iriContext() {}

func NewAbsolute_iriContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Absolute_iriContext {
	var p = new(Absolute_iriContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_absolute_iri

	return p
}

func (s *Absolute_iriContext) GetParser() antlr.Parser { return s.parser }

func (s *Absolute_iriContext) Scheme() ISchemeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemeContext)
}

func (s *Absolute_iriContext) Ihier_part() IIhier_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIhier_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIhier_partContext)
}

func (s *Absolute_iriContext) Iquery() IIqueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIqueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIqueryContext)
}

func (s *Absolute_iriContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Absolute_iriContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Absolute_iriContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterAbsolute_iri(s)
	}
}

func (s *Absolute_iriContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitAbsolute_iri(s)
	}
}

func (s *Absolute_iriContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitAbsolute_iri(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Absolute_iri() (localctx IAbsolute_iriContext) {
	localctx = NewAbsolute_iriContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, IRIParserRULE_absolute_iri)
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
		p.SetState(111)
		p.Scheme()
	}
	{
		p.SetState(112)
		p.Match(IRIParserCOL)
	}
	{
		p.SetState(113)
		p.Ihier_part()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRIParserQMARK {
		{
			p.SetState(114)
			p.Match(IRIParserQMARK)
		}
		{
			p.SetState(115)
			p.Iquery()
		}

	}

	return localctx
}

// IIrelative_refContext is an interface to support dynamic dispatch.
type IIrelative_refContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIrelative_refContext differentiates from other interfaces.
	IsIrelative_refContext()
}

type Irelative_refContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIrelative_refContext() *Irelative_refContext {
	var p = new(Irelative_refContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_irelative_ref
	return p
}

func (*Irelative_refContext) IsIrelative_refContext() {}

func NewIrelative_refContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Irelative_refContext {
	var p = new(Irelative_refContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_irelative_ref

	return p
}

func (s *Irelative_refContext) GetParser() antlr.Parser { return s.parser }

func (s *Irelative_refContext) Irelative_part() IIrelative_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIrelative_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIrelative_partContext)
}

func (s *Irelative_refContext) Iquery() IIqueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIqueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIqueryContext)
}

func (s *Irelative_refContext) Ifragment() IIfragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfragmentContext)
}

func (s *Irelative_refContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Irelative_refContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Irelative_refContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIrelative_ref(s)
	}
}

func (s *Irelative_refContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIrelative_ref(s)
	}
}

func (s *Irelative_refContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIrelative_ref(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Irelative_ref() (localctx IIrelative_refContext) {
	localctx = NewIrelative_refContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, IRIParserRULE_irelative_ref)
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
		p.SetState(118)
		p.Irelative_part()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRIParserQMARK {
		{
			p.SetState(119)
			p.Match(IRIParserQMARK)
		}
		{
			p.SetState(120)
			p.Iquery()
		}

	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRIParserHASH {
		{
			p.SetState(123)
			p.Match(IRIParserHASH)
		}
		{
			p.SetState(124)
			p.Ifragment()
		}

	}

	return localctx
}

// IIrelative_partContext is an interface to support dynamic dispatch.
type IIrelative_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIrelative_partContext differentiates from other interfaces.
	IsIrelative_partContext()
}

type Irelative_partContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIrelative_partContext() *Irelative_partContext {
	var p = new(Irelative_partContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_irelative_part
	return p
}

func (*Irelative_partContext) IsIrelative_partContext() {}

func NewIrelative_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Irelative_partContext {
	var p = new(Irelative_partContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_irelative_part

	return p
}

func (s *Irelative_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Irelative_partContext) Iauthority() IIauthorityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIauthorityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIauthorityContext)
}

func (s *Irelative_partContext) Ipath_abempty() IIpath_abemptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_abemptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_abemptyContext)
}

func (s *Irelative_partContext) Ipath_absolute() IIpath_absoluteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_absoluteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_absoluteContext)
}

func (s *Irelative_partContext) Ipath_noscheme() IIpath_noschemeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_noschemeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_noschemeContext)
}

func (s *Irelative_partContext) Ipath_empty() IIpath_emptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_emptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_emptyContext)
}

func (s *Irelative_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Irelative_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Irelative_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIrelative_part(s)
	}
}

func (s *Irelative_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIrelative_part(s)
	}
}

func (s *Irelative_partContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIrelative_part(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Irelative_part() (localctx IIrelative_partContext) {
	localctx = NewIrelative_partContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, IRIParserRULE_irelative_part)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRIParserFSLASH2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(IRIParserFSLASH2)
		}
		{
			p.SetState(128)
			p.Iauthority()
		}
		{
			p.SetState(129)
			p.Ipath_abempty()
		}

	case IRIParserFSLASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Ipath_absolute()
		}

	case IRIParserUCSCHAR, IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ, IRIParserDOT, IRIParserPERCENT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE, IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS, IRIParserAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Ipath_noscheme()
		}

	case IRIParserEOF, IRIParserQMARK, IRIParserHASH:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(133)
			p.Ipath_empty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIauthorityContext is an interface to support dynamic dispatch.
type IIauthorityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIauthorityContext differentiates from other interfaces.
	IsIauthorityContext()
}

type IauthorityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIauthorityContext() *IauthorityContext {
	var p = new(IauthorityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_iauthority
	return p
}

func (*IauthorityContext) IsIauthorityContext() {}

func NewIauthorityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IauthorityContext {
	var p = new(IauthorityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_iauthority

	return p
}

func (s *IauthorityContext) GetParser() antlr.Parser { return s.parser }

func (s *IauthorityContext) Ihost() IIhostContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIhostContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIhostContext)
}

func (s *IauthorityContext) Iuserinfo() IIuserinfoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIuserinfoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIuserinfoContext)
}

func (s *IauthorityContext) Port() IPortContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortContext)
}

func (s *IauthorityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IauthorityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IauthorityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIauthority(s)
	}
}

func (s *IauthorityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIauthority(s)
	}
}

func (s *IauthorityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIauthority(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Iauthority() (localctx IIauthorityContext) {
	localctx = NewIauthorityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, IRIParserRULE_iauthority)
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
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(136)
			p.Iuserinfo()
		}
		{
			p.SetState(137)
			p.Match(IRIParserAT)
		}

	}
	{
		p.SetState(141)
		p.Ihost()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IRIParserCOL {
		{
			p.SetState(142)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(143)
			p.Port()
		}

	}

	return localctx
}

// IIuserinfoContext is an interface to support dynamic dispatch.
type IIuserinfoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIuserinfoContext differentiates from other interfaces.
	IsIuserinfoContext()
}

type IuserinfoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIuserinfoContext() *IuserinfoContext {
	var p = new(IuserinfoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_iuserinfo
	return p
}

func (*IuserinfoContext) IsIuserinfoContext() {}

func NewIuserinfoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IuserinfoContext {
	var p = new(IuserinfoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_iuserinfo

	return p
}

func (s *IuserinfoContext) GetParser() antlr.Parser { return s.parser }

func (s *IuserinfoContext) AllIunreserved() []IIunreservedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIunreservedContext)(nil)).Elem())
	var tst = make([]IIunreservedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIunreservedContext)
		}
	}

	return tst
}

func (s *IuserinfoContext) Iunreserved(i int) IIunreservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIunreservedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIunreservedContext)
}

func (s *IuserinfoContext) AllPct_encoded() []IPct_encodedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPct_encodedContext)(nil)).Elem())
	var tst = make([]IPct_encodedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPct_encodedContext)
		}
	}

	return tst
}

func (s *IuserinfoContext) Pct_encoded(i int) IPct_encodedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPct_encodedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPct_encodedContext)
}

func (s *IuserinfoContext) AllSub_delims() []ISub_delimsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem())
	var tst = make([]ISub_delimsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISub_delimsContext)
		}
	}

	return tst
}

func (s *IuserinfoContext) Sub_delims(i int) ISub_delimsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISub_delimsContext)
}

func (s *IuserinfoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IuserinfoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IuserinfoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIuserinfo(s)
	}
}

func (s *IuserinfoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIuserinfo(s)
	}
}

func (s *IuserinfoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIuserinfo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Iuserinfo() (localctx IIuserinfoContext) {
	localctx = NewIuserinfoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, IRIParserRULE_iuserinfo)
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
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserUCSCHAR)|(1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserCOL-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserPERCENT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserTILDE-32))|(1<<(IRIParserUSCORE-32))|(1<<(IRIParserEXCL-32))|(1<<(IRIParserDOLLAR-32))|(1<<(IRIParserAMP-32))|(1<<(IRIParserSQUOTE-32))|(1<<(IRIParserOPAREN-32))|(1<<(IRIParserCPAREN-32))|(1<<(IRIParserSTAR-32))|(1<<(IRIParserPLUS-32))|(1<<(IRIParserCOMMA-32))|(1<<(IRIParserSCOL-32))|(1<<(IRIParserEQUALS-32)))) != 0) {
		p.SetState(150)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IRIParserUCSCHAR, IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ, IRIParserDOT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE:
			{
				p.SetState(146)
				p.Iunreserved()
			}

		case IRIParserPERCENT:
			{
				p.SetState(147)
				p.Pct_encoded()
			}

		case IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS:
			{
				p.SetState(148)
				p.Sub_delims()
			}

		case IRIParserCOL:
			{
				p.SetState(149)
				p.Match(IRIParserCOL)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIhostContext is an interface to support dynamic dispatch.
type IIhostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIhostContext differentiates from other interfaces.
	IsIhostContext()
}

type IhostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIhostContext() *IhostContext {
	var p = new(IhostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ihost
	return p
}

func (*IhostContext) IsIhostContext() {}

func NewIhostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IhostContext {
	var p = new(IhostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ihost

	return p
}

func (s *IhostContext) GetParser() antlr.Parser { return s.parser }

func (s *IhostContext) Ip_literal() IIp_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIp_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIp_literalContext)
}

func (s *IhostContext) Ip_v4_address() IIp_v4_addressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIp_v4_addressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIp_v4_addressContext)
}

func (s *IhostContext) Ireg_name() IIreg_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIreg_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIreg_nameContext)
}

func (s *IhostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IhostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IhostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIhost(s)
	}
}

func (s *IhostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIhost(s)
	}
}

func (s *IhostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIhost(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ihost() (localctx IIhostContext) {
	localctx = NewIhostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, IRIParserRULE_ihost)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Ip_literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Ip_v4_address()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.Ireg_name()
		}

	}

	return localctx
}

// IIreg_nameContext is an interface to support dynamic dispatch.
type IIreg_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIreg_nameContext differentiates from other interfaces.
	IsIreg_nameContext()
}

type Ireg_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIreg_nameContext() *Ireg_nameContext {
	var p = new(Ireg_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ireg_name
	return p
}

func (*Ireg_nameContext) IsIreg_nameContext() {}

func NewIreg_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ireg_nameContext {
	var p = new(Ireg_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ireg_name

	return p
}

func (s *Ireg_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Ireg_nameContext) AllIunreserved() []IIunreservedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIunreservedContext)(nil)).Elem())
	var tst = make([]IIunreservedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIunreservedContext)
		}
	}

	return tst
}

func (s *Ireg_nameContext) Iunreserved(i int) IIunreservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIunreservedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIunreservedContext)
}

func (s *Ireg_nameContext) AllPct_encoded() []IPct_encodedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPct_encodedContext)(nil)).Elem())
	var tst = make([]IPct_encodedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPct_encodedContext)
		}
	}

	return tst
}

func (s *Ireg_nameContext) Pct_encoded(i int) IPct_encodedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPct_encodedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPct_encodedContext)
}

func (s *Ireg_nameContext) AllSub_delims() []ISub_delimsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem())
	var tst = make([]ISub_delimsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISub_delimsContext)
		}
	}

	return tst
}

func (s *Ireg_nameContext) Sub_delims(i int) ISub_delimsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISub_delimsContext)
}

func (s *Ireg_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ireg_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ireg_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIreg_name(s)
	}
}

func (s *Ireg_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIreg_name(s)
	}
}

func (s *Ireg_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIreg_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ireg_name() (localctx IIreg_nameContext) {
	localctx = NewIreg_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, IRIParserRULE_ireg_name)
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
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserUCSCHAR)|(1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserPERCENT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserTILDE-32))|(1<<(IRIParserUSCORE-32))|(1<<(IRIParserEXCL-32))|(1<<(IRIParserDOLLAR-32))|(1<<(IRIParserAMP-32))|(1<<(IRIParserSQUOTE-32))|(1<<(IRIParserOPAREN-32))|(1<<(IRIParserCPAREN-32))|(1<<(IRIParserSTAR-32))|(1<<(IRIParserPLUS-32))|(1<<(IRIParserCOMMA-32))|(1<<(IRIParserSCOL-32))|(1<<(IRIParserEQUALS-32)))) != 0) {
		p.SetState(163)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IRIParserUCSCHAR, IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ, IRIParserDOT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE:
			{
				p.SetState(160)
				p.Iunreserved()
			}

		case IRIParserPERCENT:
			{
				p.SetState(161)
				p.Pct_encoded()
			}

		case IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS:
			{
				p.SetState(162)
				p.Sub_delims()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIpathContext is an interface to support dynamic dispatch.
type IIpathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpathContext differentiates from other interfaces.
	IsIpathContext()
}

type IpathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpathContext() *IpathContext {
	var p = new(IpathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ipath
	return p
}

func (*IpathContext) IsIpathContext() {}

func NewIpathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IpathContext {
	var p = new(IpathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ipath

	return p
}

func (s *IpathContext) GetParser() antlr.Parser { return s.parser }

func (s *IpathContext) Ipath_abempty() IIpath_abemptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_abemptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_abemptyContext)
}

func (s *IpathContext) Ipath_absolute() IIpath_absoluteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_absoluteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_absoluteContext)
}

func (s *IpathContext) Ipath_noscheme() IIpath_noschemeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_noschemeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_noschemeContext)
}

func (s *IpathContext) Ipath_rootless() IIpath_rootlessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_rootlessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_rootlessContext)
}

func (s *IpathContext) Ipath_empty() IIpath_emptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpath_emptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpath_emptyContext)
}

func (s *IpathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IpathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIpath(s)
	}
}

func (s *IpathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIpath(s)
	}
}

func (s *IpathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIpath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ipath() (localctx IIpathContext) {
	localctx = NewIpathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, IRIParserRULE_ipath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Ipath_abempty()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)
			p.Ipath_absolute()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(170)
			p.Ipath_noscheme()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(171)
			p.Ipath_rootless()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(172)
			p.Ipath_empty()
		}

	}

	return localctx
}

// IIpath_abemptyContext is an interface to support dynamic dispatch.
type IIpath_abemptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpath_abemptyContext differentiates from other interfaces.
	IsIpath_abemptyContext()
}

type Ipath_abemptyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpath_abemptyContext() *Ipath_abemptyContext {
	var p = new(Ipath_abemptyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ipath_abempty
	return p
}

func (*Ipath_abemptyContext) IsIpath_abemptyContext() {}

func NewIpath_abemptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipath_abemptyContext {
	var p = new(Ipath_abemptyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ipath_abempty

	return p
}

func (s *Ipath_abemptyContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipath_abemptyContext) AllIsegment() []IIsegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIsegmentContext)(nil)).Elem())
	var tst = make([]IIsegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIsegmentContext)
		}
	}

	return tst
}

func (s *Ipath_abemptyContext) Isegment(i int) IIsegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIsegmentContext)
}

func (s *Ipath_abemptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipath_abemptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipath_abemptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIpath_abempty(s)
	}
}

func (s *Ipath_abemptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIpath_abempty(s)
	}
}

func (s *Ipath_abemptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIpath_abempty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ipath_abempty() (localctx IIpath_abemptyContext) {
	localctx = NewIpath_abemptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, IRIParserRULE_ipath_abempty)
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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IRIParserFSLASH {
		{
			p.SetState(175)
			p.Match(IRIParserFSLASH)
		}
		{
			p.SetState(176)
			p.Isegment()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIpath_absoluteContext is an interface to support dynamic dispatch.
type IIpath_absoluteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpath_absoluteContext differentiates from other interfaces.
	IsIpath_absoluteContext()
}

type Ipath_absoluteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpath_absoluteContext() *Ipath_absoluteContext {
	var p = new(Ipath_absoluteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ipath_absolute
	return p
}

func (*Ipath_absoluteContext) IsIpath_absoluteContext() {}

func NewIpath_absoluteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipath_absoluteContext {
	var p = new(Ipath_absoluteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ipath_absolute

	return p
}

func (s *Ipath_absoluteContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipath_absoluteContext) Isegment_nz() IIsegment_nzContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsegment_nzContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsegment_nzContext)
}

func (s *Ipath_absoluteContext) AllIsegment() []IIsegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIsegmentContext)(nil)).Elem())
	var tst = make([]IIsegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIsegmentContext)
		}
	}

	return tst
}

func (s *Ipath_absoluteContext) Isegment(i int) IIsegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIsegmentContext)
}

func (s *Ipath_absoluteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipath_absoluteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipath_absoluteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIpath_absolute(s)
	}
}

func (s *Ipath_absoluteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIpath_absolute(s)
	}
}

func (s *Ipath_absoluteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIpath_absolute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ipath_absolute() (localctx IIpath_absoluteContext) {
	localctx = NewIpath_absoluteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, IRIParserRULE_ipath_absolute)
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
		p.SetState(182)
		p.Match(IRIParserFSLASH)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserUCSCHAR)|(1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserCOL-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserPERCENT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserTILDE-32))|(1<<(IRIParserUSCORE-32))|(1<<(IRIParserEXCL-32))|(1<<(IRIParserDOLLAR-32))|(1<<(IRIParserAMP-32))|(1<<(IRIParserSQUOTE-32))|(1<<(IRIParserOPAREN-32))|(1<<(IRIParserCPAREN-32))|(1<<(IRIParserSTAR-32))|(1<<(IRIParserPLUS-32))|(1<<(IRIParserCOMMA-32))|(1<<(IRIParserSCOL-32))|(1<<(IRIParserEQUALS-32))|(1<<(IRIParserAT-32)))) != 0) {
		{
			p.SetState(183)
			p.Isegment_nz()
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IRIParserFSLASH {
			{
				p.SetState(184)
				p.Match(IRIParserFSLASH)
			}
			{
				p.SetState(185)
				p.Isegment()
			}

			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IIpath_noschemeContext is an interface to support dynamic dispatch.
type IIpath_noschemeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpath_noschemeContext differentiates from other interfaces.
	IsIpath_noschemeContext()
}

type Ipath_noschemeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpath_noschemeContext() *Ipath_noschemeContext {
	var p = new(Ipath_noschemeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ipath_noscheme
	return p
}

func (*Ipath_noschemeContext) IsIpath_noschemeContext() {}

func NewIpath_noschemeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipath_noschemeContext {
	var p = new(Ipath_noschemeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ipath_noscheme

	return p
}

func (s *Ipath_noschemeContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipath_noschemeContext) Isegment_nz_nc() IIsegment_nz_ncContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsegment_nz_ncContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsegment_nz_ncContext)
}

func (s *Ipath_noschemeContext) AllIsegment() []IIsegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIsegmentContext)(nil)).Elem())
	var tst = make([]IIsegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIsegmentContext)
		}
	}

	return tst
}

func (s *Ipath_noschemeContext) Isegment(i int) IIsegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIsegmentContext)
}

func (s *Ipath_noschemeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipath_noschemeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipath_noschemeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIpath_noscheme(s)
	}
}

func (s *Ipath_noschemeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIpath_noscheme(s)
	}
}

func (s *Ipath_noschemeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIpath_noscheme(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ipath_noscheme() (localctx IIpath_noschemeContext) {
	localctx = NewIpath_noschemeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, IRIParserRULE_ipath_noscheme)
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
		p.SetState(193)
		p.Isegment_nz_nc()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IRIParserFSLASH {
		{
			p.SetState(194)
			p.Match(IRIParserFSLASH)
		}
		{
			p.SetState(195)
			p.Isegment()
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIpath_rootlessContext is an interface to support dynamic dispatch.
type IIpath_rootlessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpath_rootlessContext differentiates from other interfaces.
	IsIpath_rootlessContext()
}

type Ipath_rootlessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpath_rootlessContext() *Ipath_rootlessContext {
	var p = new(Ipath_rootlessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ipath_rootless
	return p
}

func (*Ipath_rootlessContext) IsIpath_rootlessContext() {}

func NewIpath_rootlessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipath_rootlessContext {
	var p = new(Ipath_rootlessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ipath_rootless

	return p
}

func (s *Ipath_rootlessContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipath_rootlessContext) Isegment_nz() IIsegment_nzContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsegment_nzContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsegment_nzContext)
}

func (s *Ipath_rootlessContext) AllIsegment() []IIsegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIsegmentContext)(nil)).Elem())
	var tst = make([]IIsegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIsegmentContext)
		}
	}

	return tst
}

func (s *Ipath_rootlessContext) Isegment(i int) IIsegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIsegmentContext)
}

func (s *Ipath_rootlessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipath_rootlessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipath_rootlessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIpath_rootless(s)
	}
}

func (s *Ipath_rootlessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIpath_rootless(s)
	}
}

func (s *Ipath_rootlessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIpath_rootless(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ipath_rootless() (localctx IIpath_rootlessContext) {
	localctx = NewIpath_rootlessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, IRIParserRULE_ipath_rootless)
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
		p.SetState(201)
		p.Isegment_nz()
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IRIParserFSLASH {
		{
			p.SetState(202)
			p.Match(IRIParserFSLASH)
		}
		{
			p.SetState(203)
			p.Isegment()
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIpath_emptyContext is an interface to support dynamic dispatch.
type IIpath_emptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpath_emptyContext differentiates from other interfaces.
	IsIpath_emptyContext()
}

type Ipath_emptyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpath_emptyContext() *Ipath_emptyContext {
	var p = new(Ipath_emptyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ipath_empty
	return p
}

func (*Ipath_emptyContext) IsIpath_emptyContext() {}

func NewIpath_emptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipath_emptyContext {
	var p = new(Ipath_emptyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ipath_empty

	return p
}

func (s *Ipath_emptyContext) GetParser() antlr.Parser { return s.parser }
func (s *Ipath_emptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipath_emptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipath_emptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIpath_empty(s)
	}
}

func (s *Ipath_emptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIpath_empty(s)
	}
}

func (s *Ipath_emptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIpath_empty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ipath_empty() (localctx IIpath_emptyContext) {
	localctx = NewIpath_emptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, IRIParserRULE_ipath_empty)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	return localctx
}

// IIsegmentContext is an interface to support dynamic dispatch.
type IIsegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsegmentContext differentiates from other interfaces.
	IsIsegmentContext()
}

type IsegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsegmentContext() *IsegmentContext {
	var p = new(IsegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_isegment
	return p
}

func (*IsegmentContext) IsIsegmentContext() {}

func NewIsegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsegmentContext {
	var p = new(IsegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_isegment

	return p
}

func (s *IsegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *IsegmentContext) AllIpchar() []IIpcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIpcharContext)(nil)).Elem())
	var tst = make([]IIpcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIpcharContext)
		}
	}

	return tst
}

func (s *IsegmentContext) Ipchar(i int) IIpcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIpcharContext)
}

func (s *IsegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIsegment(s)
	}
}

func (s *IsegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIsegment(s)
	}
}

func (s *IsegmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIsegment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Isegment() (localctx IIsegmentContext) {
	localctx = NewIsegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, IRIParserRULE_isegment)
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
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserUCSCHAR)|(1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserCOL-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserPERCENT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserTILDE-32))|(1<<(IRIParserUSCORE-32))|(1<<(IRIParserEXCL-32))|(1<<(IRIParserDOLLAR-32))|(1<<(IRIParserAMP-32))|(1<<(IRIParserSQUOTE-32))|(1<<(IRIParserOPAREN-32))|(1<<(IRIParserCPAREN-32))|(1<<(IRIParserSTAR-32))|(1<<(IRIParserPLUS-32))|(1<<(IRIParserCOMMA-32))|(1<<(IRIParserSCOL-32))|(1<<(IRIParserEQUALS-32))|(1<<(IRIParserAT-32)))) != 0) {
		{
			p.SetState(211)
			p.Ipchar()
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIsegment_nzContext is an interface to support dynamic dispatch.
type IIsegment_nzContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsegment_nzContext differentiates from other interfaces.
	IsIsegment_nzContext()
}

type Isegment_nzContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsegment_nzContext() *Isegment_nzContext {
	var p = new(Isegment_nzContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_isegment_nz
	return p
}

func (*Isegment_nzContext) IsIsegment_nzContext() {}

func NewIsegment_nzContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Isegment_nzContext {
	var p = new(Isegment_nzContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_isegment_nz

	return p
}

func (s *Isegment_nzContext) GetParser() antlr.Parser { return s.parser }

func (s *Isegment_nzContext) AllIpchar() []IIpcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIpcharContext)(nil)).Elem())
	var tst = make([]IIpcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIpcharContext)
		}
	}

	return tst
}

func (s *Isegment_nzContext) Ipchar(i int) IIpcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIpcharContext)
}

func (s *Isegment_nzContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Isegment_nzContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Isegment_nzContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIsegment_nz(s)
	}
}

func (s *Isegment_nzContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIsegment_nz(s)
	}
}

func (s *Isegment_nzContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIsegment_nz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Isegment_nz() (localctx IIsegment_nzContext) {
	localctx = NewIsegment_nzContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, IRIParserRULE_isegment_nz)
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
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserUCSCHAR)|(1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserCOL-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserPERCENT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserTILDE-32))|(1<<(IRIParserUSCORE-32))|(1<<(IRIParserEXCL-32))|(1<<(IRIParserDOLLAR-32))|(1<<(IRIParserAMP-32))|(1<<(IRIParserSQUOTE-32))|(1<<(IRIParserOPAREN-32))|(1<<(IRIParserCPAREN-32))|(1<<(IRIParserSTAR-32))|(1<<(IRIParserPLUS-32))|(1<<(IRIParserCOMMA-32))|(1<<(IRIParserSCOL-32))|(1<<(IRIParserEQUALS-32))|(1<<(IRIParserAT-32)))) != 0) {
		{
			p.SetState(217)
			p.Ipchar()
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIsegment_nz_ncContext is an interface to support dynamic dispatch.
type IIsegment_nz_ncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsegment_nz_ncContext differentiates from other interfaces.
	IsIsegment_nz_ncContext()
}

type Isegment_nz_ncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsegment_nz_ncContext() *Isegment_nz_ncContext {
	var p = new(Isegment_nz_ncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_isegment_nz_nc
	return p
}

func (*Isegment_nz_ncContext) IsIsegment_nz_ncContext() {}

func NewIsegment_nz_ncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Isegment_nz_ncContext {
	var p = new(Isegment_nz_ncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_isegment_nz_nc

	return p
}

func (s *Isegment_nz_ncContext) GetParser() antlr.Parser { return s.parser }

func (s *Isegment_nz_ncContext) AllIunreserved() []IIunreservedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIunreservedContext)(nil)).Elem())
	var tst = make([]IIunreservedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIunreservedContext)
		}
	}

	return tst
}

func (s *Isegment_nz_ncContext) Iunreserved(i int) IIunreservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIunreservedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIunreservedContext)
}

func (s *Isegment_nz_ncContext) AllPct_encoded() []IPct_encodedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPct_encodedContext)(nil)).Elem())
	var tst = make([]IPct_encodedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPct_encodedContext)
		}
	}

	return tst
}

func (s *Isegment_nz_ncContext) Pct_encoded(i int) IPct_encodedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPct_encodedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPct_encodedContext)
}

func (s *Isegment_nz_ncContext) AllSub_delims() []ISub_delimsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem())
	var tst = make([]ISub_delimsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISub_delimsContext)
		}
	}

	return tst
}

func (s *Isegment_nz_ncContext) Sub_delims(i int) ISub_delimsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISub_delimsContext)
}

func (s *Isegment_nz_ncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Isegment_nz_ncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Isegment_nz_ncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIsegment_nz_nc(s)
	}
}

func (s *Isegment_nz_ncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIsegment_nz_nc(s)
	}
}

func (s *Isegment_nz_ncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIsegment_nz_nc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Isegment_nz_nc() (localctx IIsegment_nz_ncContext) {
	localctx = NewIsegment_nz_ncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, IRIParserRULE_isegment_nz_nc)
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
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserUCSCHAR)|(1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserPERCENT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserTILDE-32))|(1<<(IRIParserUSCORE-32))|(1<<(IRIParserEXCL-32))|(1<<(IRIParserDOLLAR-32))|(1<<(IRIParserAMP-32))|(1<<(IRIParserSQUOTE-32))|(1<<(IRIParserOPAREN-32))|(1<<(IRIParserCPAREN-32))|(1<<(IRIParserSTAR-32))|(1<<(IRIParserPLUS-32))|(1<<(IRIParserCOMMA-32))|(1<<(IRIParserSCOL-32))|(1<<(IRIParserEQUALS-32))|(1<<(IRIParserAT-32)))) != 0) {
		p.SetState(226)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IRIParserUCSCHAR, IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ, IRIParserDOT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE:
			{
				p.SetState(222)
				p.Iunreserved()
			}

		case IRIParserPERCENT:
			{
				p.SetState(223)
				p.Pct_encoded()
			}

		case IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS:
			{
				p.SetState(224)
				p.Sub_delims()
			}

		case IRIParserAT:
			{
				p.SetState(225)
				p.Match(IRIParserAT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIpcharContext is an interface to support dynamic dispatch.
type IIpcharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpcharContext differentiates from other interfaces.
	IsIpcharContext()
}

type IpcharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpcharContext() *IpcharContext {
	var p = new(IpcharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ipchar
	return p
}

func (*IpcharContext) IsIpcharContext() {}

func NewIpcharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IpcharContext {
	var p = new(IpcharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ipchar

	return p
}

func (s *IpcharContext) GetParser() antlr.Parser { return s.parser }

func (s *IpcharContext) Iunreserved() IIunreservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIunreservedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIunreservedContext)
}

func (s *IpcharContext) Pct_encoded() IPct_encodedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPct_encodedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPct_encodedContext)
}

func (s *IpcharContext) Sub_delims() ISub_delimsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISub_delimsContext)
}

func (s *IpcharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpcharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IpcharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIpchar(s)
	}
}

func (s *IpcharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIpchar(s)
	}
}

func (s *IpcharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIpchar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ipchar() (localctx IIpcharContext) {
	localctx = NewIpcharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, IRIParserRULE_ipchar)
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

	p.SetState(234)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRIParserUCSCHAR, IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ, IRIParserDOT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Iunreserved()
		}

	case IRIParserPERCENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.Pct_encoded()
		}

	case IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(232)
			p.Sub_delims()
		}

	case IRIParserCOL, IRIParserAT:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(233)
		_la = p.GetTokenStream().LA(1)

		if !(_la == IRIParserCOL || _la == IRIParserAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIqueryContext is an interface to support dynamic dispatch.
type IIqueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIqueryContext differentiates from other interfaces.
	IsIqueryContext()
}

type IqueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIqueryContext() *IqueryContext {
	var p = new(IqueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_iquery
	return p
}

func (*IqueryContext) IsIqueryContext() {}

func NewIqueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IqueryContext {
	var p = new(IqueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_iquery

	return p
}

func (s *IqueryContext) GetParser() antlr.Parser { return s.parser }

func (s *IqueryContext) AllIpchar() []IIpcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIpcharContext)(nil)).Elem())
	var tst = make([]IIpcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIpcharContext)
		}
	}

	return tst
}

func (s *IqueryContext) Ipchar(i int) IIpcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIpcharContext)
}

func (s *IqueryContext) AllIPRIVATE() []antlr.TerminalNode {
	return s.GetTokens(IRIParserIPRIVATE)
}

func (s *IqueryContext) IPRIVATE(i int) antlr.TerminalNode {
	return s.GetToken(IRIParserIPRIVATE, i)
}

func (s *IqueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IqueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IqueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIquery(s)
	}
}

func (s *IqueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIquery(s)
	}
}

func (s *IqueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIquery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Iquery() (localctx IIqueryContext) {
	localctx = NewIqueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, IRIParserRULE_iquery)
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
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserUCSCHAR)|(1<<IRIParserIPRIVATE)|(1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserCOL-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserPERCENT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserTILDE-32))|(1<<(IRIParserUSCORE-32))|(1<<(IRIParserEXCL-32))|(1<<(IRIParserDOLLAR-32))|(1<<(IRIParserAMP-32))|(1<<(IRIParserSQUOTE-32))|(1<<(IRIParserOPAREN-32))|(1<<(IRIParserCPAREN-32))|(1<<(IRIParserSTAR-32))|(1<<(IRIParserPLUS-32))|(1<<(IRIParserCOMMA-32))|(1<<(IRIParserSCOL-32))|(1<<(IRIParserEQUALS-32))|(1<<(IRIParserFSLASH-32))|(1<<(IRIParserQMARK-32))|(1<<(IRIParserAT-32)))) != 0) {
		p.SetState(238)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IRIParserUCSCHAR, IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ, IRIParserCOL, IRIParserDOT, IRIParserPERCENT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE, IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS, IRIParserAT:
			{
				p.SetState(236)
				p.Ipchar()
			}

		case IRIParserIPRIVATE, IRIParserFSLASH, IRIParserQMARK:
			p.SetState(237)
			_la = p.GetTokenStream().LA(1)

			if !(_la == IRIParserIPRIVATE || _la == IRIParserFSLASH || _la == IRIParserQMARK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIfragmentContext is an interface to support dynamic dispatch.
type IIfragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfragmentContext differentiates from other interfaces.
	IsIfragmentContext()
}

type IfragmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfragmentContext() *IfragmentContext {
	var p = new(IfragmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ifragment
	return p
}

func (*IfragmentContext) IsIfragmentContext() {}

func NewIfragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfragmentContext {
	var p = new(IfragmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ifragment

	return p
}

func (s *IfragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *IfragmentContext) AllIpchar() []IIpcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIpcharContext)(nil)).Elem())
	var tst = make([]IIpcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIpcharContext)
		}
	}

	return tst
}

func (s *IfragmentContext) Ipchar(i int) IIpcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIpcharContext)
}

func (s *IfragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfragmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIfragment(s)
	}
}

func (s *IfragmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIfragment(s)
	}
}

func (s *IfragmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIfragment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ifragment() (localctx IIfragmentContext) {
	localctx = NewIfragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, IRIParserRULE_ifragment)
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
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserUCSCHAR)|(1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserCOL-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserPERCENT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserTILDE-32))|(1<<(IRIParserUSCORE-32))|(1<<(IRIParserEXCL-32))|(1<<(IRIParserDOLLAR-32))|(1<<(IRIParserAMP-32))|(1<<(IRIParserSQUOTE-32))|(1<<(IRIParserOPAREN-32))|(1<<(IRIParserCPAREN-32))|(1<<(IRIParserSTAR-32))|(1<<(IRIParserPLUS-32))|(1<<(IRIParserCOMMA-32))|(1<<(IRIParserSCOL-32))|(1<<(IRIParserEQUALS-32))|(1<<(IRIParserFSLASH-32))|(1<<(IRIParserQMARK-32))|(1<<(IRIParserAT-32)))) != 0) {
		p.SetState(245)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IRIParserUCSCHAR, IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ, IRIParserCOL, IRIParserDOT, IRIParserPERCENT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE, IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS, IRIParserAT:
			{
				p.SetState(243)
				p.Ipchar()
			}

		case IRIParserFSLASH, IRIParserQMARK:
			p.SetState(244)
			_la = p.GetTokenStream().LA(1)

			if !(_la == IRIParserFSLASH || _la == IRIParserQMARK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIunreservedContext is an interface to support dynamic dispatch.
type IIunreservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIunreservedContext differentiates from other interfaces.
	IsIunreservedContext()
}

type IunreservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIunreservedContext() *IunreservedContext {
	var p = new(IunreservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_iunreserved
	return p
}

func (*IunreservedContext) IsIunreservedContext() {}

func NewIunreservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IunreservedContext {
	var p = new(IunreservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_iunreserved

	return p
}

func (s *IunreservedContext) GetParser() antlr.Parser { return s.parser }

func (s *IunreservedContext) Alpha() IAlphaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlphaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlphaContext)
}

func (s *IunreservedContext) Digit() IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *IunreservedContext) UCSCHAR() antlr.TerminalNode {
	return s.GetToken(IRIParserUCSCHAR, 0)
}

func (s *IunreservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IunreservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IunreservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIunreserved(s)
	}
}

func (s *IunreservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIunreserved(s)
	}
}

func (s *IunreservedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIunreserved(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Iunreserved() (localctx IIunreservedContext) {
	localctx = NewIunreservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, IRIParserRULE_iunreserved)
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

	p.SetState(253)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Alpha()
		}

	case IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Digit()
		}

	case IRIParserUCSCHAR, IRIParserDOT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(252)
		_la = p.GetTokenStream().LA(1)

		if !(_la == IRIParserUCSCHAR || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(IRIParserDOT-41))|(1<<(IRIParserHYPHEN-41))|(1<<(IRIParserTILDE-41))|(1<<(IRIParserUSCORE-41)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISchemeContext is an interface to support dynamic dispatch.
type ISchemeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemeContext differentiates from other interfaces.
	IsSchemeContext()
}

type SchemeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemeContext() *SchemeContext {
	var p = new(SchemeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_scheme
	return p
}

func (*SchemeContext) IsSchemeContext() {}

func NewSchemeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemeContext {
	var p = new(SchemeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_scheme

	return p
}

func (s *SchemeContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemeContext) AllAlpha() []IAlphaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAlphaContext)(nil)).Elem())
	var tst = make([]IAlphaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAlphaContext)
		}
	}

	return tst
}

func (s *SchemeContext) Alpha(i int) IAlphaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlphaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAlphaContext)
}

func (s *SchemeContext) AllDigit() []IDigitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDigitContext)(nil)).Elem())
	var tst = make([]IDigitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDigitContext)
		}
	}

	return tst
}

func (s *SchemeContext) Digit(i int) IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *SchemeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterScheme(s)
	}
}

func (s *SchemeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitScheme(s)
	}
}

func (s *SchemeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitScheme(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Scheme() (localctx ISchemeContext) {
	localctx = NewSchemeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, IRIParserRULE_scheme)
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
		p.Alpha()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserPLUS-32)))) != 0) {
		p.SetState(259)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ:
			{
				p.SetState(256)
				p.Alpha()
			}

		case IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9:
			{
				p.SetState(257)
				p.Digit()
			}

		case IRIParserDOT, IRIParserHYPHEN, IRIParserPLUS:
			p.SetState(258)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(IRIParserDOT-41))|(1<<(IRIParserHYPHEN-41))|(1<<(IRIParserPLUS-41)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPortContext is an interface to support dynamic dispatch.
type IPortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPortContext differentiates from other interfaces.
	IsPortContext()
}

type PortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortContext() *PortContext {
	var p = new(PortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_port
	return p
}

func (*PortContext) IsPortContext() {}

func NewPortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortContext {
	var p = new(PortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_port

	return p
}

func (s *PortContext) GetParser() antlr.Parser { return s.parser }

func (s *PortContext) AllDigit() []IDigitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDigitContext)(nil)).Elem())
	var tst = make([]IDigitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDigitContext)
		}
	}

	return tst
}

func (s *PortContext) Digit(i int) IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *PortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterPort(s)
	}
}

func (s *PortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitPort(s)
	}
}

func (s *PortContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitPort(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Port() (localctx IPortContext) {
	localctx = NewPortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, IRIParserRULE_port)
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
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9))) != 0 {
		{
			p.SetState(264)
			p.Digit()
		}

		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIp_literalContext is an interface to support dynamic dispatch.
type IIp_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIp_literalContext differentiates from other interfaces.
	IsIp_literalContext()
}

type Ip_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIp_literalContext() *Ip_literalContext {
	var p = new(Ip_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ip_literal
	return p
}

func (*Ip_literalContext) IsIp_literalContext() {}

func NewIp_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ip_literalContext {
	var p = new(Ip_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ip_literal

	return p
}

func (s *Ip_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Ip_literalContext) Ip_v6_address() IIp_v6_addressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIp_v6_addressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIp_v6_addressContext)
}

func (s *Ip_literalContext) Ip_v_future() IIp_v_futureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIp_v_futureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIp_v_futureContext)
}

func (s *Ip_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ip_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ip_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIp_literal(s)
	}
}

func (s *Ip_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIp_literal(s)
	}
}

func (s *Ip_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIp_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ip_literal() (localctx IIp_literalContext) {
	localctx = NewIp_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, IRIParserRULE_ip_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(270)
		p.Match(IRIParserOBRACK)
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserCOL2:
		{
			p.SetState(271)
			p.Ip_v6_address()
		}

	case IRIParserV:
		{
			p.SetState(272)
			p.Ip_v_future()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(275)
		p.Match(IRIParserCBRACK)
	}

	return localctx
}

// IIp_v_futureContext is an interface to support dynamic dispatch.
type IIp_v_futureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIp_v_futureContext differentiates from other interfaces.
	IsIp_v_futureContext()
}

type Ip_v_futureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIp_v_futureContext() *Ip_v_futureContext {
	var p = new(Ip_v_futureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ip_v_future
	return p
}

func (*Ip_v_futureContext) IsIp_v_futureContext() {}

func NewIp_v_futureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ip_v_futureContext {
	var p = new(Ip_v_futureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ip_v_future

	return p
}

func (s *Ip_v_futureContext) GetParser() antlr.Parser { return s.parser }

func (s *Ip_v_futureContext) V() antlr.TerminalNode {
	return s.GetToken(IRIParserV, 0)
}

func (s *Ip_v_futureContext) AllHexdig() []IHexdigContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHexdigContext)(nil)).Elem())
	var tst = make([]IHexdigContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHexdigContext)
		}
	}

	return tst
}

func (s *Ip_v_futureContext) Hexdig(i int) IHexdigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexdigContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHexdigContext)
}

func (s *Ip_v_futureContext) AllUnreserved() []IUnreservedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnreservedContext)(nil)).Elem())
	var tst = make([]IUnreservedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnreservedContext)
		}
	}

	return tst
}

func (s *Ip_v_futureContext) Unreserved(i int) IUnreservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnreservedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnreservedContext)
}

func (s *Ip_v_futureContext) AllSub_delims() []ISub_delimsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem())
	var tst = make([]ISub_delimsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISub_delimsContext)
		}
	}

	return tst
}

func (s *Ip_v_futureContext) Sub_delims(i int) ISub_delimsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISub_delimsContext)
}

func (s *Ip_v_futureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ip_v_futureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ip_v_futureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIp_v_future(s)
	}
}

func (s *Ip_v_futureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIp_v_future(s)
	}
}

func (s *Ip_v_futureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIp_v_future(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ip_v_future() (localctx IIp_v_futureContext) {
	localctx = NewIp_v_futureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, IRIParserRULE_ip_v_future)
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
		p.SetState(277)
		p.Match(IRIParserV)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF))) != 0) {
		{
			p.SetState(278)
			p.Hexdig()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(283)
		p.Match(IRIParserDOT)
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF)|(1<<IRIParserG)|(1<<IRIParserH)|(1<<IRIParserI)|(1<<IRIParserJ)|(1<<IRIParserK)|(1<<IRIParserL)|(1<<IRIParserM)|(1<<IRIParserN)|(1<<IRIParserO)|(1<<IRIParserP)|(1<<IRIParserQ)|(1<<IRIParserR)|(1<<IRIParserS))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IRIParserT-32))|(1<<(IRIParserU-32))|(1<<(IRIParserV-32))|(1<<(IRIParserW-32))|(1<<(IRIParserX-32))|(1<<(IRIParserY-32))|(1<<(IRIParserZ-32))|(1<<(IRIParserCOL-32))|(1<<(IRIParserDOT-32))|(1<<(IRIParserHYPHEN-32))|(1<<(IRIParserTILDE-32))|(1<<(IRIParserUSCORE-32))|(1<<(IRIParserEXCL-32))|(1<<(IRIParserDOLLAR-32))|(1<<(IRIParserAMP-32))|(1<<(IRIParserSQUOTE-32))|(1<<(IRIParserOPAREN-32))|(1<<(IRIParserCPAREN-32))|(1<<(IRIParserSTAR-32))|(1<<(IRIParserPLUS-32))|(1<<(IRIParserCOMMA-32))|(1<<(IRIParserSCOL-32))|(1<<(IRIParserEQUALS-32)))) != 0) {
		p.SetState(287)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9, IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ, IRIParserDOT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE:
			{
				p.SetState(284)
				p.Unreserved()
			}

		case IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS:
			{
				p.SetState(285)
				p.Sub_delims()
			}

		case IRIParserCOL:
			{
				p.SetState(286)
				p.Match(IRIParserCOL)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIp_v6_addressContext is an interface to support dynamic dispatch.
type IIp_v6_addressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIp_v6_addressContext differentiates from other interfaces.
	IsIp_v6_addressContext()
}

type Ip_v6_addressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIp_v6_addressContext() *Ip_v6_addressContext {
	var p = new(Ip_v6_addressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ip_v6_address
	return p
}

func (*Ip_v6_addressContext) IsIp_v6_addressContext() {}

func NewIp_v6_addressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ip_v6_addressContext {
	var p = new(Ip_v6_addressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ip_v6_address

	return p
}

func (s *Ip_v6_addressContext) GetParser() antlr.Parser { return s.parser }

func (s *Ip_v6_addressContext) AllH16() []IH16Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IH16Context)(nil)).Elem())
	var tst = make([]IH16Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IH16Context)
		}
	}

	return tst
}

func (s *Ip_v6_addressContext) H16(i int) IH16Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IH16Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IH16Context)
}

func (s *Ip_v6_addressContext) Ls32() ILs32Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILs32Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILs32Context)
}

func (s *Ip_v6_addressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ip_v6_addressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ip_v6_addressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIp_v6_address(s)
	}
}

func (s *Ip_v6_addressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIp_v6_address(s)
	}
}

func (s *Ip_v6_addressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIp_v6_address(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ip_v6_address() (localctx IIp_v6_addressContext) {
	localctx = NewIp_v6_addressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, IRIParserRULE_ip_v6_address)
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

	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)
			p.H16()
		}
		{
			p.SetState(292)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(293)
			p.H16()
		}
		{
			p.SetState(294)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(295)
			p.H16()
		}
		{
			p.SetState(296)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(297)
			p.H16()
		}
		{
			p.SetState(298)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(299)
			p.H16()
		}
		{
			p.SetState(300)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(301)
			p.H16()
		}
		{
			p.SetState(302)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(303)
			p.Ls32()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Match(IRIParserCOL2)
		}
		{
			p.SetState(306)
			p.H16()
		}
		{
			p.SetState(307)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(308)
			p.H16()
		}
		{
			p.SetState(309)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(310)
			p.H16()
		}
		{
			p.SetState(311)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(312)
			p.H16()
		}
		{
			p.SetState(313)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(314)
			p.H16()
		}
		{
			p.SetState(315)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(316)
			p.Ls32()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF))) != 0 {
			{
				p.SetState(318)
				p.H16()
			}

		}
		{
			p.SetState(321)
			p.Match(IRIParserCOL2)
		}
		{
			p.SetState(322)
			p.H16()
		}
		{
			p.SetState(323)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(324)
			p.H16()
		}
		{
			p.SetState(325)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(326)
			p.H16()
		}
		{
			p.SetState(327)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(328)
			p.H16()
		}
		{
			p.SetState(329)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(330)
			p.Ls32()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF))) != 0 {
			p.SetState(335)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(332)
					p.H16()
				}
				{
					p.SetState(333)
					p.Match(IRIParserCOL)
				}

			}
			{
				p.SetState(337)
				p.H16()
			}

		}
		{
			p.SetState(340)
			p.Match(IRIParserCOL2)
		}
		{
			p.SetState(341)
			p.H16()
		}
		{
			p.SetState(342)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(343)
			p.H16()
		}
		{
			p.SetState(344)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(345)
			p.H16()
		}
		{
			p.SetState(346)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(347)
			p.Ls32()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF))) != 0 {
			p.SetState(357)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
				p.SetState(352)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(349)
						p.H16()
					}
					{
						p.SetState(350)
						p.Match(IRIParserCOL)
					}

				}
				{
					p.SetState(354)
					p.H16()
				}
				{
					p.SetState(355)
					p.Match(IRIParserCOL)
				}

			}
			{
				p.SetState(359)
				p.H16()
			}

		}
		{
			p.SetState(362)
			p.Match(IRIParserCOL2)
		}
		{
			p.SetState(363)
			p.H16()
		}
		{
			p.SetState(364)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(365)
			p.H16()
		}
		{
			p.SetState(366)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(367)
			p.Ls32()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF))) != 0 {
			p.SetState(382)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
				p.SetState(377)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
					p.SetState(372)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
						{
							p.SetState(369)
							p.H16()
						}
						{
							p.SetState(370)
							p.Match(IRIParserCOL)
						}

					}
					{
						p.SetState(374)
						p.H16()
					}
					{
						p.SetState(375)
						p.Match(IRIParserCOL)
					}

				}
				{
					p.SetState(379)
					p.H16()
				}
				{
					p.SetState(380)
					p.Match(IRIParserCOL)
				}

			}
			{
				p.SetState(384)
				p.H16()
			}

		}
		{
			p.SetState(387)
			p.Match(IRIParserCOL2)
		}
		{
			p.SetState(388)
			p.H16()
		}
		{
			p.SetState(389)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(390)
			p.Ls32()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF))) != 0 {
			p.SetState(410)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
				p.SetState(405)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
					p.SetState(400)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) == 1 {
						p.SetState(395)
						p.GetErrorHandler().Sync(p)

						if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
							{
								p.SetState(392)
								p.H16()
							}
							{
								p.SetState(393)
								p.Match(IRIParserCOL)
							}

						}
						{
							p.SetState(397)
							p.H16()
						}
						{
							p.SetState(398)
							p.Match(IRIParserCOL)
						}

					}
					{
						p.SetState(402)
						p.H16()
					}
					{
						p.SetState(403)
						p.Match(IRIParserCOL)
					}

				}
				{
					p.SetState(407)
					p.H16()
				}
				{
					p.SetState(408)
					p.Match(IRIParserCOL)
				}

			}
			{
				p.SetState(412)
				p.H16()
			}

		}
		{
			p.SetState(415)
			p.Match(IRIParserCOL2)
		}
		{
			p.SetState(416)
			p.Ls32()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF))) != 0 {
			p.SetState(440)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
				p.SetState(435)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) == 1 {
					p.SetState(430)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
						p.SetState(425)
						p.GetErrorHandler().Sync(p)

						if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
							p.SetState(420)
							p.GetErrorHandler().Sync(p)

							if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
								{
									p.SetState(417)
									p.H16()
								}
								{
									p.SetState(418)
									p.Match(IRIParserCOL)
								}

							}
							{
								p.SetState(422)
								p.H16()
							}
							{
								p.SetState(423)
								p.Match(IRIParserCOL)
							}

						}
						{
							p.SetState(427)
							p.H16()
						}
						{
							p.SetState(428)
							p.Match(IRIParserCOL)
						}

					}
					{
						p.SetState(432)
						p.H16()
					}
					{
						p.SetState(433)
						p.Match(IRIParserCOL)
					}

				}
				{
					p.SetState(437)
					p.H16()
				}
				{
					p.SetState(438)
					p.Match(IRIParserCOL)
				}

			}
			{
				p.SetState(442)
				p.H16()
			}

		}
		{
			p.SetState(445)
			p.Match(IRIParserCOL2)
		}
		{
			p.SetState(446)
			p.H16()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		p.SetState(478)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9)|(1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF))) != 0 {
			p.SetState(475)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext()) == 1 {
				p.SetState(470)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) == 1 {
					p.SetState(465)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
						p.SetState(460)
						p.GetErrorHandler().Sync(p)

						if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) == 1 {
							p.SetState(455)
							p.GetErrorHandler().Sync(p)

							if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
								p.SetState(450)
								p.GetErrorHandler().Sync(p)

								if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) == 1 {
									{
										p.SetState(447)
										p.H16()
									}
									{
										p.SetState(448)
										p.Match(IRIParserCOL)
									}

								}
								{
									p.SetState(452)
									p.H16()
								}
								{
									p.SetState(453)
									p.Match(IRIParserCOL)
								}

							}
							{
								p.SetState(457)
								p.H16()
							}
							{
								p.SetState(458)
								p.Match(IRIParserCOL)
							}

						}
						{
							p.SetState(462)
							p.H16()
						}
						{
							p.SetState(463)
							p.Match(IRIParserCOL)
						}

					}
					{
						p.SetState(467)
						p.H16()
					}
					{
						p.SetState(468)
						p.Match(IRIParserCOL)
					}

				}
				{
					p.SetState(472)
					p.H16()
				}
				{
					p.SetState(473)
					p.Match(IRIParserCOL)
				}

			}
			{
				p.SetState(477)
				p.H16()
			}

		}
		{
			p.SetState(480)
			p.Match(IRIParserCOL2)
		}

	}

	return localctx
}

// IH16Context is an interface to support dynamic dispatch.
type IH16Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsH16Context differentiates from other interfaces.
	IsH16Context()
}

type H16Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyH16Context() *H16Context {
	var p = new(H16Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_h16
	return p
}

func (*H16Context) IsH16Context() {}

func NewH16Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H16Context {
	var p = new(H16Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_h16

	return p
}

func (s *H16Context) GetParser() antlr.Parser { return s.parser }

func (s *H16Context) AllHexdig() []IHexdigContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHexdigContext)(nil)).Elem())
	var tst = make([]IHexdigContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHexdigContext)
		}
	}

	return tst
}

func (s *H16Context) Hexdig(i int) IHexdigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexdigContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHexdigContext)
}

func (s *H16Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H16Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H16Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterH16(s)
	}
}

func (s *H16Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitH16(s)
	}
}

func (s *H16Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitH16(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) H16() (localctx IH16Context) {
	localctx = NewH16Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, IRIParserRULE_h16)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)
			p.Hexdig()
		}
		{
			p.SetState(484)
			p.Hexdig()
		}
		{
			p.SetState(485)
			p.Hexdig()
		}
		{
			p.SetState(486)
			p.Hexdig()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(488)
			p.Hexdig()
		}
		{
			p.SetState(489)
			p.Hexdig()
		}
		{
			p.SetState(490)
			p.Hexdig()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(492)
			p.Hexdig()
		}
		{
			p.SetState(493)
			p.Hexdig()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(495)
			p.Hexdig()
		}

	}

	return localctx
}

// ILs32Context is an interface to support dynamic dispatch.
type ILs32Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLs32Context differentiates from other interfaces.
	IsLs32Context()
}

type Ls32Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLs32Context() *Ls32Context {
	var p = new(Ls32Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ls32
	return p
}

func (*Ls32Context) IsLs32Context() {}

func NewLs32Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ls32Context {
	var p = new(Ls32Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ls32

	return p
}

func (s *Ls32Context) GetParser() antlr.Parser { return s.parser }

func (s *Ls32Context) AllH16() []IH16Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IH16Context)(nil)).Elem())
	var tst = make([]IH16Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IH16Context)
		}
	}

	return tst
}

func (s *Ls32Context) H16(i int) IH16Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IH16Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IH16Context)
}

func (s *Ls32Context) Ip_v4_address() IIp_v4_addressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIp_v4_addressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIp_v4_addressContext)
}

func (s *Ls32Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ls32Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ls32Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterLs32(s)
	}
}

func (s *Ls32Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitLs32(s)
	}
}

func (s *Ls32Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitLs32(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ls32() (localctx ILs32Context) {
	localctx = NewLs32Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, IRIParserRULE_ls32)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(498)
			p.H16()
		}
		{
			p.SetState(499)
			p.Match(IRIParserCOL)
		}
		{
			p.SetState(500)
			p.H16()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(502)
			p.Ip_v4_address()
		}

	}

	return localctx
}

// IIp_v4_addressContext is an interface to support dynamic dispatch.
type IIp_v4_addressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIp_v4_addressContext differentiates from other interfaces.
	IsIp_v4_addressContext()
}

type Ip_v4_addressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIp_v4_addressContext() *Ip_v4_addressContext {
	var p = new(Ip_v4_addressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_ip_v4_address
	return p
}

func (*Ip_v4_addressContext) IsIp_v4_addressContext() {}

func NewIp_v4_addressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ip_v4_addressContext {
	var p = new(Ip_v4_addressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_ip_v4_address

	return p
}

func (s *Ip_v4_addressContext) GetParser() antlr.Parser { return s.parser }

func (s *Ip_v4_addressContext) AllDec_octet() []IDec_octetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDec_octetContext)(nil)).Elem())
	var tst = make([]IDec_octetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDec_octetContext)
		}
	}

	return tst
}

func (s *Ip_v4_addressContext) Dec_octet(i int) IDec_octetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDec_octetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDec_octetContext)
}

func (s *Ip_v4_addressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ip_v4_addressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ip_v4_addressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterIp_v4_address(s)
	}
}

func (s *Ip_v4_addressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitIp_v4_address(s)
	}
}

func (s *Ip_v4_addressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitIp_v4_address(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Ip_v4_address() (localctx IIp_v4_addressContext) {
	localctx = NewIp_v4_addressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, IRIParserRULE_ip_v4_address)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(505)
		p.Dec_octet()
	}
	{
		p.SetState(506)
		p.Match(IRIParserDOT)
	}
	{
		p.SetState(507)
		p.Dec_octet()
	}
	{
		p.SetState(508)
		p.Match(IRIParserDOT)
	}
	{
		p.SetState(509)
		p.Dec_octet()
	}
	{
		p.SetState(510)
		p.Match(IRIParserDOT)
	}
	{
		p.SetState(511)
		p.Dec_octet()
	}

	return localctx
}

// IDec_octetContext is an interface to support dynamic dispatch.
type IDec_octetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDec_octetContext differentiates from other interfaces.
	IsDec_octetContext()
}

type Dec_octetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDec_octetContext() *Dec_octetContext {
	var p = new(Dec_octetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_dec_octet
	return p
}

func (*Dec_octetContext) IsDec_octetContext() {}

func NewDec_octetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_octetContext {
	var p = new(Dec_octetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_dec_octet

	return p
}

func (s *Dec_octetContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_octetContext) AllDigit() []IDigitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDigitContext)(nil)).Elem())
	var tst = make([]IDigitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDigitContext)
		}
	}

	return tst
}

func (s *Dec_octetContext) Digit(i int) IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *Dec_octetContext) Non_zero_digit() INon_zero_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INon_zero_digitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INon_zero_digitContext)
}

func (s *Dec_octetContext) D1() antlr.TerminalNode {
	return s.GetToken(IRIParserD1, 0)
}

func (s *Dec_octetContext) AllD2() []antlr.TerminalNode {
	return s.GetTokens(IRIParserD2)
}

func (s *Dec_octetContext) D2(i int) antlr.TerminalNode {
	return s.GetToken(IRIParserD2, i)
}

func (s *Dec_octetContext) D0() antlr.TerminalNode {
	return s.GetToken(IRIParserD0, 0)
}

func (s *Dec_octetContext) D3() antlr.TerminalNode {
	return s.GetToken(IRIParserD3, 0)
}

func (s *Dec_octetContext) D4() antlr.TerminalNode {
	return s.GetToken(IRIParserD4, 0)
}

func (s *Dec_octetContext) AllD5() []antlr.TerminalNode {
	return s.GetTokens(IRIParserD5)
}

func (s *Dec_octetContext) D5(i int) antlr.TerminalNode {
	return s.GetToken(IRIParserD5, i)
}

func (s *Dec_octetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_octetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dec_octetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterDec_octet(s)
	}
}

func (s *Dec_octetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitDec_octet(s)
	}
}

func (s *Dec_octetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitDec_octet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Dec_octet() (localctx IDec_octetContext) {
	localctx = NewDec_octetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, IRIParserRULE_dec_octet)
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

	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(513)
			p.Digit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(514)
			p.Non_zero_digit()
		}
		{
			p.SetState(515)
			p.Digit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(517)
			p.Match(IRIParserD1)
		}
		{
			p.SetState(518)
			p.Digit()
		}
		{
			p.SetState(519)
			p.Digit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(521)
			p.Match(IRIParserD2)
		}
		p.SetState(522)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(523)
			p.Digit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(524)
			p.Match(IRIParserD2)
		}
		{
			p.SetState(525)
			p.Match(IRIParserD5)
		}
		p.SetState(526)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD0)|(1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// IPct_encodedContext is an interface to support dynamic dispatch.
type IPct_encodedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPct_encodedContext differentiates from other interfaces.
	IsPct_encodedContext()
}

type Pct_encodedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPct_encodedContext() *Pct_encodedContext {
	var p = new(Pct_encodedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_pct_encoded
	return p
}

func (*Pct_encodedContext) IsPct_encodedContext() {}

func NewPct_encodedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pct_encodedContext {
	var p = new(Pct_encodedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_pct_encoded

	return p
}

func (s *Pct_encodedContext) GetParser() antlr.Parser { return s.parser }

func (s *Pct_encodedContext) AllHexdig() []IHexdigContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHexdigContext)(nil)).Elem())
	var tst = make([]IHexdigContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHexdigContext)
		}
	}

	return tst
}

func (s *Pct_encodedContext) Hexdig(i int) IHexdigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexdigContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHexdigContext)
}

func (s *Pct_encodedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pct_encodedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pct_encodedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterPct_encoded(s)
	}
}

func (s *Pct_encodedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitPct_encoded(s)
	}
}

func (s *Pct_encodedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitPct_encoded(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Pct_encoded() (localctx IPct_encodedContext) {
	localctx = NewPct_encodedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, IRIParserRULE_pct_encoded)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(529)
		p.Match(IRIParserPERCENT)
	}
	{
		p.SetState(530)
		p.Hexdig()
	}
	{
		p.SetState(531)
		p.Hexdig()
	}

	return localctx
}

// IUnreservedContext is an interface to support dynamic dispatch.
type IUnreservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnreservedContext differentiates from other interfaces.
	IsUnreservedContext()
}

type UnreservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnreservedContext() *UnreservedContext {
	var p = new(UnreservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_unreserved
	return p
}

func (*UnreservedContext) IsUnreservedContext() {}

func NewUnreservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnreservedContext {
	var p = new(UnreservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_unreserved

	return p
}

func (s *UnreservedContext) GetParser() antlr.Parser { return s.parser }

func (s *UnreservedContext) Alpha() IAlphaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlphaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlphaContext)
}

func (s *UnreservedContext) Digit() IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *UnreservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnreservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnreservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterUnreserved(s)
	}
}

func (s *UnreservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitUnreserved(s)
	}
}

func (s *UnreservedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitUnreserved(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Unreserved() (localctx IUnreservedContext) {
	localctx = NewUnreservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, IRIParserRULE_unreserved)
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

	p.SetState(536)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF, IRIParserG, IRIParserH, IRIParserI, IRIParserJ, IRIParserK, IRIParserL, IRIParserM, IRIParserN, IRIParserO, IRIParserP, IRIParserQ, IRIParserR, IRIParserS, IRIParserT, IRIParserU, IRIParserV, IRIParserW, IRIParserX, IRIParserY, IRIParserZ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(533)
			p.Alpha()
		}

	case IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(534)
			p.Digit()
		}

	case IRIParserDOT, IRIParserHYPHEN, IRIParserTILDE, IRIParserUSCORE:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(535)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(IRIParserDOT-41))|(1<<(IRIParserHYPHEN-41))|(1<<(IRIParserTILDE-41))|(1<<(IRIParserUSCORE-41)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReservedContext is an interface to support dynamic dispatch.
type IReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedContext differentiates from other interfaces.
	IsReservedContext()
}

type ReservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedContext() *ReservedContext {
	var p = new(ReservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_reserved
	return p
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) Gen_delims() IGen_delimsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGen_delimsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGen_delimsContext)
}

func (s *ReservedContext) Sub_delims() ISub_delimsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISub_delimsContext)
}

func (s *ReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterReserved(s)
	}
}

func (s *ReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitReserved(s)
	}
}

func (s *ReservedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitReserved(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Reserved() (localctx IReservedContext) {
	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, IRIParserRULE_reserved)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(540)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRIParserCOL, IRIParserFSLASH, IRIParserQMARK, IRIParserHASH, IRIParserOBRACK, IRIParserCBRACK, IRIParserAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(538)
			p.Gen_delims()
		}

	case IRIParserEXCL, IRIParserDOLLAR, IRIParserAMP, IRIParserSQUOTE, IRIParserOPAREN, IRIParserCPAREN, IRIParserSTAR, IRIParserPLUS, IRIParserCOMMA, IRIParserSCOL, IRIParserEQUALS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(539)
			p.Sub_delims()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGen_delimsContext is an interface to support dynamic dispatch.
type IGen_delimsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGen_delimsContext differentiates from other interfaces.
	IsGen_delimsContext()
}

type Gen_delimsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGen_delimsContext() *Gen_delimsContext {
	var p = new(Gen_delimsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_gen_delims
	return p
}

func (*Gen_delimsContext) IsGen_delimsContext() {}

func NewGen_delimsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gen_delimsContext {
	var p = new(Gen_delimsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_gen_delims

	return p
}

func (s *Gen_delimsContext) GetParser() antlr.Parser { return s.parser }
func (s *Gen_delimsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gen_delimsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gen_delimsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterGen_delims(s)
	}
}

func (s *Gen_delimsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitGen_delims(s)
	}
}

func (s *Gen_delimsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitGen_delims(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Gen_delims() (localctx IGen_delimsContext) {
	localctx = NewGen_delimsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, IRIParserRULE_gen_delims)
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
	p.SetState(542)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(IRIParserCOL-40))|(1<<(IRIParserFSLASH-40))|(1<<(IRIParserQMARK-40))|(1<<(IRIParserHASH-40))|(1<<(IRIParserOBRACK-40))|(1<<(IRIParserCBRACK-40))|(1<<(IRIParserAT-40)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ISub_delimsContext is an interface to support dynamic dispatch.
type ISub_delimsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSub_delimsContext differentiates from other interfaces.
	IsSub_delimsContext()
}

type Sub_delimsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySub_delimsContext() *Sub_delimsContext {
	var p = new(Sub_delimsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_sub_delims
	return p
}

func (*Sub_delimsContext) IsSub_delimsContext() {}

func NewSub_delimsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sub_delimsContext {
	var p = new(Sub_delimsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_sub_delims

	return p
}

func (s *Sub_delimsContext) GetParser() antlr.Parser { return s.parser }
func (s *Sub_delimsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sub_delimsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sub_delimsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterSub_delims(s)
	}
}

func (s *Sub_delimsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitSub_delims(s)
	}
}

func (s *Sub_delimsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitSub_delims(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Sub_delims() (localctx ISub_delimsContext) {
	localctx = NewSub_delimsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, IRIParserRULE_sub_delims)
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
	p.SetState(544)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(IRIParserEXCL-46))|(1<<(IRIParserDOLLAR-46))|(1<<(IRIParserAMP-46))|(1<<(IRIParserSQUOTE-46))|(1<<(IRIParserOPAREN-46))|(1<<(IRIParserCPAREN-46))|(1<<(IRIParserSTAR-46))|(1<<(IRIParserPLUS-46))|(1<<(IRIParserCOMMA-46))|(1<<(IRIParserSCOL-46))|(1<<(IRIParserEQUALS-46)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IAlphaContext is an interface to support dynamic dispatch.
type IAlphaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlphaContext differentiates from other interfaces.
	IsAlphaContext()
}

type AlphaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlphaContext() *AlphaContext {
	var p = new(AlphaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_alpha
	return p
}

func (*AlphaContext) IsAlphaContext() {}

func NewAlphaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlphaContext {
	var p = new(AlphaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_alpha

	return p
}

func (s *AlphaContext) GetParser() antlr.Parser { return s.parser }

func (s *AlphaContext) A() antlr.TerminalNode {
	return s.GetToken(IRIParserA, 0)
}

func (s *AlphaContext) B() antlr.TerminalNode {
	return s.GetToken(IRIParserB, 0)
}

func (s *AlphaContext) C() antlr.TerminalNode {
	return s.GetToken(IRIParserC, 0)
}

func (s *AlphaContext) D() antlr.TerminalNode {
	return s.GetToken(IRIParserD, 0)
}

func (s *AlphaContext) E() antlr.TerminalNode {
	return s.GetToken(IRIParserE, 0)
}

func (s *AlphaContext) F() antlr.TerminalNode {
	return s.GetToken(IRIParserF, 0)
}

func (s *AlphaContext) G() antlr.TerminalNode {
	return s.GetToken(IRIParserG, 0)
}

func (s *AlphaContext) H() antlr.TerminalNode {
	return s.GetToken(IRIParserH, 0)
}

func (s *AlphaContext) I() antlr.TerminalNode {
	return s.GetToken(IRIParserI, 0)
}

func (s *AlphaContext) J() antlr.TerminalNode {
	return s.GetToken(IRIParserJ, 0)
}

func (s *AlphaContext) K() antlr.TerminalNode {
	return s.GetToken(IRIParserK, 0)
}

func (s *AlphaContext) L() antlr.TerminalNode {
	return s.GetToken(IRIParserL, 0)
}

func (s *AlphaContext) M() antlr.TerminalNode {
	return s.GetToken(IRIParserM, 0)
}

func (s *AlphaContext) N() antlr.TerminalNode {
	return s.GetToken(IRIParserN, 0)
}

func (s *AlphaContext) O() antlr.TerminalNode {
	return s.GetToken(IRIParserO, 0)
}

func (s *AlphaContext) P() antlr.TerminalNode {
	return s.GetToken(IRIParserP, 0)
}

func (s *AlphaContext) Q() antlr.TerminalNode {
	return s.GetToken(IRIParserQ, 0)
}

func (s *AlphaContext) R() antlr.TerminalNode {
	return s.GetToken(IRIParserR, 0)
}

func (s *AlphaContext) S() antlr.TerminalNode {
	return s.GetToken(IRIParserS, 0)
}

func (s *AlphaContext) T() antlr.TerminalNode {
	return s.GetToken(IRIParserT, 0)
}

func (s *AlphaContext) U() antlr.TerminalNode {
	return s.GetToken(IRIParserU, 0)
}

func (s *AlphaContext) V() antlr.TerminalNode {
	return s.GetToken(IRIParserV, 0)
}

func (s *AlphaContext) W() antlr.TerminalNode {
	return s.GetToken(IRIParserW, 0)
}

func (s *AlphaContext) X() antlr.TerminalNode {
	return s.GetToken(IRIParserX, 0)
}

func (s *AlphaContext) Y() antlr.TerminalNode {
	return s.GetToken(IRIParserY, 0)
}

func (s *AlphaContext) Z() antlr.TerminalNode {
	return s.GetToken(IRIParserZ, 0)
}

func (s *AlphaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlphaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlphaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterAlpha(s)
	}
}

func (s *AlphaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitAlpha(s)
	}
}

func (s *AlphaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitAlpha(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Alpha() (localctx IAlphaContext) {
	localctx = NewAlphaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, IRIParserRULE_alpha)
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
	p.SetState(546)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(IRIParserA-13))|(1<<(IRIParserB-13))|(1<<(IRIParserC-13))|(1<<(IRIParserD-13))|(1<<(IRIParserE-13))|(1<<(IRIParserF-13))|(1<<(IRIParserG-13))|(1<<(IRIParserH-13))|(1<<(IRIParserI-13))|(1<<(IRIParserJ-13))|(1<<(IRIParserK-13))|(1<<(IRIParserL-13))|(1<<(IRIParserM-13))|(1<<(IRIParserN-13))|(1<<(IRIParserO-13))|(1<<(IRIParserP-13))|(1<<(IRIParserQ-13))|(1<<(IRIParserR-13))|(1<<(IRIParserS-13))|(1<<(IRIParserT-13))|(1<<(IRIParserU-13))|(1<<(IRIParserV-13))|(1<<(IRIParserW-13))|(1<<(IRIParserX-13))|(1<<(IRIParserY-13))|(1<<(IRIParserZ-13)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IHexdigContext is an interface to support dynamic dispatch.
type IHexdigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHexdigContext differentiates from other interfaces.
	IsHexdigContext()
}

type HexdigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexdigContext() *HexdigContext {
	var p = new(HexdigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_hexdig
	return p
}

func (*HexdigContext) IsHexdigContext() {}

func NewHexdigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexdigContext {
	var p = new(HexdigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_hexdig

	return p
}

func (s *HexdigContext) GetParser() antlr.Parser { return s.parser }

func (s *HexdigContext) Digit() IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *HexdigContext) A() antlr.TerminalNode {
	return s.GetToken(IRIParserA, 0)
}

func (s *HexdigContext) B() antlr.TerminalNode {
	return s.GetToken(IRIParserB, 0)
}

func (s *HexdigContext) C() antlr.TerminalNode {
	return s.GetToken(IRIParserC, 0)
}

func (s *HexdigContext) D() antlr.TerminalNode {
	return s.GetToken(IRIParserD, 0)
}

func (s *HexdigContext) E() antlr.TerminalNode {
	return s.GetToken(IRIParserE, 0)
}

func (s *HexdigContext) F() antlr.TerminalNode {
	return s.GetToken(IRIParserF, 0)
}

func (s *HexdigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexdigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HexdigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterHexdig(s)
	}
}

func (s *HexdigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitHexdig(s)
	}
}

func (s *HexdigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitHexdig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Hexdig() (localctx IHexdigContext) {
	localctx = NewHexdigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, IRIParserRULE_hexdig)
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

	p.SetState(550)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRIParserD0, IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(548)
			p.Digit()
		}

	case IRIParserA, IRIParserB, IRIParserC, IRIParserD, IRIParserE, IRIParserF:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(549)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserA)|(1<<IRIParserB)|(1<<IRIParserC)|(1<<IRIParserD)|(1<<IRIParserE)|(1<<IRIParserF))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDigitContext is an interface to support dynamic dispatch.
type IDigitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDigitContext differentiates from other interfaces.
	IsDigitContext()
}

type DigitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDigitContext() *DigitContext {
	var p = new(DigitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_digit
	return p
}

func (*DigitContext) IsDigitContext() {}

func NewDigitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DigitContext {
	var p = new(DigitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_digit

	return p
}

func (s *DigitContext) GetParser() antlr.Parser { return s.parser }

func (s *DigitContext) D0() antlr.TerminalNode {
	return s.GetToken(IRIParserD0, 0)
}

func (s *DigitContext) Non_zero_digit() INon_zero_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INon_zero_digitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INon_zero_digitContext)
}

func (s *DigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DigitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterDigit(s)
	}
}

func (s *DigitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitDigit(s)
	}
}

func (s *DigitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitDigit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Digit() (localctx IDigitContext) {
	localctx = NewDigitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, IRIParserRULE_digit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(554)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IRIParserD0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(552)
			p.Match(IRIParserD0)
		}

	case IRIParserD1, IRIParserD2, IRIParserD3, IRIParserD4, IRIParserD5, IRIParserD6, IRIParserD7, IRIParserD8, IRIParserD9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(553)
			p.Non_zero_digit()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INon_zero_digitContext is an interface to support dynamic dispatch.
type INon_zero_digitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNon_zero_digitContext differentiates from other interfaces.
	IsNon_zero_digitContext()
}

type Non_zero_digitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNon_zero_digitContext() *Non_zero_digitContext {
	var p = new(Non_zero_digitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IRIParserRULE_non_zero_digit
	return p
}

func (*Non_zero_digitContext) IsNon_zero_digitContext() {}

func NewNon_zero_digitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Non_zero_digitContext {
	var p = new(Non_zero_digitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IRIParserRULE_non_zero_digit

	return p
}

func (s *Non_zero_digitContext) GetParser() antlr.Parser { return s.parser }

func (s *Non_zero_digitContext) D1() antlr.TerminalNode {
	return s.GetToken(IRIParserD1, 0)
}

func (s *Non_zero_digitContext) D2() antlr.TerminalNode {
	return s.GetToken(IRIParserD2, 0)
}

func (s *Non_zero_digitContext) D3() antlr.TerminalNode {
	return s.GetToken(IRIParserD3, 0)
}

func (s *Non_zero_digitContext) D4() antlr.TerminalNode {
	return s.GetToken(IRIParserD4, 0)
}

func (s *Non_zero_digitContext) D5() antlr.TerminalNode {
	return s.GetToken(IRIParserD5, 0)
}

func (s *Non_zero_digitContext) D6() antlr.TerminalNode {
	return s.GetToken(IRIParserD6, 0)
}

func (s *Non_zero_digitContext) D7() antlr.TerminalNode {
	return s.GetToken(IRIParserD7, 0)
}

func (s *Non_zero_digitContext) D8() antlr.TerminalNode {
	return s.GetToken(IRIParserD8, 0)
}

func (s *Non_zero_digitContext) D9() antlr.TerminalNode {
	return s.GetToken(IRIParserD9, 0)
}

func (s *Non_zero_digitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Non_zero_digitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Non_zero_digitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.EnterNon_zero_digit(s)
	}
}

func (s *Non_zero_digitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IRIListener); ok {
		listenerT.ExitNon_zero_digit(s)
	}
}

func (s *Non_zero_digitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IRIVisitor:
		return t.VisitNon_zero_digit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IRIParser) Non_zero_digit() (localctx INon_zero_digitContext) {
	localctx = NewNon_zero_digitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, IRIParserRULE_non_zero_digit)
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
	p.SetState(556)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IRIParserD1)|(1<<IRIParserD2)|(1<<IRIParserD3)|(1<<IRIParserD4)|(1<<IRIParserD5)|(1<<IRIParserD6)|(1<<IRIParserD7)|(1<<IRIParserD8)|(1<<IRIParserD9))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
