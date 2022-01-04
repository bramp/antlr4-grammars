// Code generated from Protobuf3.g4 by ANTLR 4.9.3. DO NOT EDIT.

package protobuf3 // Protobuf3
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 60, 476,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 7, 2, 111, 10, 2, 12, 2, 14, 2, 114, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 5, 4, 123, 10, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 5, 7, 144, 10, 7, 5, 7, 146, 10, 7, 3, 8, 5, 8, 149, 10, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 159, 10, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 7, 9, 166, 10, 9, 12, 9, 14, 9, 169, 11, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	7, 12, 183, 10, 12, 12, 12, 14, 12, 186, 11, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 198, 10, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 5, 14, 215, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 238, 10, 16, 3, 17,
	3, 17, 3, 17, 5, 17, 243, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7,
	18, 250, 10, 18, 12, 18, 14, 18, 253, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	5, 19, 259, 10, 19, 5, 19, 261, 10, 19, 3, 20, 3, 20, 3, 20, 7, 20, 266,
	10, 20, 12, 20, 14, 20, 269, 11, 20, 3, 21, 3, 21, 3, 21, 5, 21, 274, 10,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 7, 23, 282, 10, 23, 12, 23,
	14, 23, 285, 11, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 292, 10,
	24, 3, 25, 3, 25, 3, 25, 5, 25, 297, 10, 25, 3, 25, 3, 25, 5, 25, 301,
	10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 309, 10, 26, 12,
	26, 14, 26, 312, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 7, 29, 326, 10, 29, 12, 29, 14, 29,
	329, 11, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 5, 30, 341, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 347,
	10, 31, 12, 31, 14, 31, 350, 11, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32,
	5, 32, 357, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 363, 10, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 370, 10, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 7, 33, 377, 10, 33, 12, 33, 14, 33, 380, 11, 33, 3, 33, 3,
	33, 5, 33, 384, 10, 33, 3, 34, 3, 34, 5, 34, 388, 10, 34, 3, 34, 3, 34,
	5, 34, 392, 10, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 398, 10, 34, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35, 405, 10, 35, 12, 35, 14, 35, 408,
	11, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 5, 37, 416, 10, 37, 3,
	38, 3, 38, 3, 38, 7, 38, 421, 10, 38, 12, 38, 14, 38, 424, 11, 38, 3, 39,
	3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3,
	44, 3, 45, 3, 45, 3, 46, 5, 46, 441, 10, 46, 3, 46, 3, 46, 3, 46, 7, 46,
	446, 10, 46, 12, 46, 14, 46, 449, 11, 46, 3, 46, 3, 46, 3, 47, 5, 47, 454,
	10, 47, 3, 47, 3, 47, 3, 47, 7, 47, 459, 10, 47, 12, 47, 14, 47, 462, 11,
	47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51,
	3, 52, 3, 52, 3, 52, 2, 2, 53, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
	96, 98, 100, 102, 2, 8, 3, 2, 36, 37, 3, 2, 5, 6, 3, 2, 12, 23, 3, 2, 51,
	52, 4, 2, 36, 37, 53, 53, 4, 2, 3, 35, 54, 54, 2, 499, 2, 104, 3, 2, 2,
	2, 4, 115, 3, 2, 2, 2, 6, 120, 3, 2, 2, 2, 8, 127, 3, 2, 2, 2, 10, 131,
	3, 2, 2, 2, 12, 145, 3, 2, 2, 2, 14, 148, 3, 2, 2, 2, 16, 162, 3, 2, 2,
	2, 18, 170, 3, 2, 2, 2, 20, 174, 3, 2, 2, 2, 22, 176, 3, 2, 2, 2, 24, 189,
	3, 2, 2, 2, 26, 201, 3, 2, 2, 2, 28, 218, 3, 2, 2, 2, 30, 237, 3, 2, 2,
	2, 32, 239, 3, 2, 2, 2, 34, 246, 3, 2, 2, 2, 36, 254, 3, 2, 2, 2, 38, 262,
	3, 2, 2, 2, 40, 273, 3, 2, 2, 2, 42, 275, 3, 2, 2, 2, 44, 279, 3, 2, 2,
	2, 46, 291, 3, 2, 2, 2, 48, 293, 3, 2, 2, 2, 50, 304, 3, 2, 2, 2, 52, 315,
	3, 2, 2, 2, 54, 319, 3, 2, 2, 2, 56, 323, 3, 2, 2, 2, 58, 340, 3, 2, 2,
	2, 60, 342, 3, 2, 2, 2, 62, 356, 3, 2, 2, 2, 64, 358, 3, 2, 2, 2, 66, 397,
	3, 2, 2, 2, 68, 399, 3, 2, 2, 2, 70, 411, 3, 2, 2, 2, 72, 415, 3, 2, 2,
	2, 74, 417, 3, 2, 2, 2, 76, 425, 3, 2, 2, 2, 78, 427, 3, 2, 2, 2, 80, 429,
	3, 2, 2, 2, 82, 431, 3, 2, 2, 2, 84, 433, 3, 2, 2, 2, 86, 435, 3, 2, 2,
	2, 88, 437, 3, 2, 2, 2, 90, 440, 3, 2, 2, 2, 92, 453, 3, 2, 2, 2, 94, 465,
	3, 2, 2, 2, 96, 467, 3, 2, 2, 2, 98, 469, 3, 2, 2, 2, 100, 471, 3, 2, 2,
	2, 102, 473, 3, 2, 2, 2, 104, 112, 5, 4, 3, 2, 105, 111, 5, 6, 4, 2, 106,
	111, 5, 8, 5, 2, 107, 111, 5, 10, 6, 2, 108, 111, 5, 40, 21, 2, 109, 111,
	5, 70, 36, 2, 110, 105, 3, 2, 2, 2, 110, 106, 3, 2, 2, 2, 110, 107, 3,
	2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 114, 3, 2, 2,
	2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 3, 3, 2, 2, 2, 114,
	112, 3, 2, 2, 2, 115, 116, 7, 3, 2, 2, 116, 117, 7, 39, 2, 2, 117, 118,
	9, 2, 2, 2, 118, 119, 7, 38, 2, 2, 119, 5, 3, 2, 2, 2, 120, 122, 7, 4,
	2, 2, 121, 123, 9, 3, 2, 2, 122, 121, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2,
	123, 124, 3, 2, 2, 2, 124, 125, 5, 96, 49, 2, 125, 126, 7, 38, 2, 2, 126,
	7, 3, 2, 2, 2, 127, 128, 7, 7, 2, 2, 128, 129, 5, 74, 38, 2, 129, 130,
	7, 38, 2, 2, 130, 9, 3, 2, 2, 2, 131, 132, 7, 8, 2, 2, 132, 133, 5, 12,
	7, 2, 133, 134, 7, 39, 2, 2, 134, 135, 5, 66, 34, 2, 135, 136, 7, 38, 2,
	2, 136, 11, 3, 2, 2, 2, 137, 146, 5, 74, 38, 2, 138, 139, 7, 40, 2, 2,
	139, 140, 5, 74, 38, 2, 140, 143, 7, 41, 2, 2, 141, 142, 7, 48, 2, 2, 142,
	144, 5, 74, 38, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 146,
	3, 2, 2, 2, 145, 137, 3, 2, 2, 2, 145, 138, 3, 2, 2, 2, 146, 13, 3, 2,
	2, 2, 147, 149, 7, 9, 2, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2,
	149, 150, 3, 2, 2, 2, 150, 151, 5, 30, 16, 2, 151, 152, 5, 80, 41, 2, 152,
	153, 7, 39, 2, 2, 153, 158, 5, 20, 11, 2, 154, 155, 7, 42, 2, 2, 155, 156,
	5, 16, 9, 2, 156, 157, 7, 43, 2, 2, 157, 159, 3, 2, 2, 2, 158, 154, 3,
	2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161, 7, 38, 2,
	2, 161, 15, 3, 2, 2, 2, 162, 167, 5, 18, 10, 2, 163, 164, 7, 49, 2, 2,
	164, 166, 5, 18, 10, 2, 165, 163, 3, 2, 2, 2, 166, 169, 3, 2, 2, 2, 167,
	165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 17, 3, 2, 2, 2, 169, 167, 3,
	2, 2, 2, 170, 171, 5, 12, 7, 2, 171, 172, 7, 39, 2, 2, 172, 173, 5, 66,
	34, 2, 173, 19, 3, 2, 2, 2, 174, 175, 5, 94, 48, 2, 175, 21, 3, 2, 2, 2,
	176, 177, 7, 10, 2, 2, 177, 178, 5, 82, 42, 2, 178, 184, 7, 44, 2, 2, 179,
	183, 5, 10, 6, 2, 180, 183, 5, 24, 13, 2, 181, 183, 5, 70, 36, 2, 182,
	179, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 181, 3, 2, 2, 2, 183, 186,
	3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 187, 3, 2,
	2, 2, 186, 184, 3, 2, 2, 2, 187, 188, 7, 45, 2, 2, 188, 23, 3, 2, 2, 2,
	189, 190, 5, 30, 16, 2, 190, 191, 5, 80, 41, 2, 191, 192, 7, 39, 2, 2,
	192, 197, 5, 20, 11, 2, 193, 194, 7, 42, 2, 2, 194, 195, 5, 16, 9, 2, 195,
	196, 7, 43, 2, 2, 196, 198, 3, 2, 2, 2, 197, 193, 3, 2, 2, 2, 197, 198,
	3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 7, 38, 2, 2, 200, 25, 3, 2,
	2, 2, 201, 202, 7, 11, 2, 2, 202, 203, 7, 46, 2, 2, 203, 204, 5, 28, 15,
	2, 204, 205, 7, 49, 2, 2, 205, 206, 5, 30, 16, 2, 206, 207, 7, 47, 2, 2,
	207, 208, 5, 84, 43, 2, 208, 209, 7, 39, 2, 2, 209, 214, 5, 20, 11, 2,
	210, 211, 7, 42, 2, 2, 211, 212, 5, 16, 9, 2, 212, 213, 7, 43, 2, 2, 213,
	215, 3, 2, 2, 2, 214, 210, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216,
	3, 2, 2, 2, 216, 217, 7, 38, 2, 2, 217, 27, 3, 2, 2, 2, 218, 219, 9, 4,
	2, 2, 219, 29, 3, 2, 2, 2, 220, 238, 7, 24, 2, 2, 221, 238, 7, 25, 2, 2,
	222, 238, 7, 12, 2, 2, 223, 238, 7, 13, 2, 2, 224, 238, 7, 14, 2, 2, 225,
	238, 7, 15, 2, 2, 226, 238, 7, 16, 2, 2, 227, 238, 7, 17, 2, 2, 228, 238,
	7, 18, 2, 2, 229, 238, 7, 19, 2, 2, 230, 238, 7, 20, 2, 2, 231, 238, 7,
	21, 2, 2, 232, 238, 7, 22, 2, 2, 233, 238, 7, 23, 2, 2, 234, 238, 7, 26,
	2, 2, 235, 238, 5, 90, 46, 2, 236, 238, 5, 92, 47, 2, 237, 220, 3, 2, 2,
	2, 237, 221, 3, 2, 2, 2, 237, 222, 3, 2, 2, 2, 237, 223, 3, 2, 2, 2, 237,
	224, 3, 2, 2, 2, 237, 225, 3, 2, 2, 2, 237, 226, 3, 2, 2, 2, 237, 227,
	3, 2, 2, 2, 237, 228, 3, 2, 2, 2, 237, 229, 3, 2, 2, 2, 237, 230, 3, 2,
	2, 2, 237, 231, 3, 2, 2, 2, 237, 232, 3, 2, 2, 2, 237, 233, 3, 2, 2, 2,
	237, 234, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 236, 3, 2, 2, 2, 238,
	31, 3, 2, 2, 2, 239, 242, 7, 27, 2, 2, 240, 243, 5, 34, 18, 2, 241, 243,
	5, 38, 20, 2, 242, 240, 3, 2, 2, 2, 242, 241, 3, 2, 2, 2, 243, 244, 3,
	2, 2, 2, 244, 245, 7, 38, 2, 2, 245, 33, 3, 2, 2, 2, 246, 251, 5, 36, 19,
	2, 247, 248, 7, 49, 2, 2, 248, 250, 5, 36, 19, 2, 249, 247, 3, 2, 2, 2,
	250, 253, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252,
	35, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 254, 260, 5, 94, 48, 2, 255, 258,
	7, 28, 2, 2, 256, 259, 5, 94, 48, 2, 257, 259, 7, 29, 2, 2, 258, 256, 3,
	2, 2, 2, 258, 257, 3, 2, 2, 2, 259, 261, 3, 2, 2, 2, 260, 255, 3, 2, 2,
	2, 260, 261, 3, 2, 2, 2, 261, 37, 3, 2, 2, 2, 262, 267, 5, 96, 49, 2, 263,
	264, 7, 49, 2, 2, 264, 266, 5, 96, 49, 2, 265, 263, 3, 2, 2, 2, 266, 269,
	3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 39, 3, 2,
	2, 2, 269, 267, 3, 2, 2, 2, 270, 274, 5, 54, 28, 2, 271, 274, 5, 42, 22,
	2, 272, 274, 5, 60, 31, 2, 273, 270, 3, 2, 2, 2, 273, 271, 3, 2, 2, 2,
	273, 272, 3, 2, 2, 2, 274, 41, 3, 2, 2, 2, 275, 276, 7, 30, 2, 2, 276,
	277, 5, 78, 40, 2, 277, 278, 5, 44, 23, 2, 278, 43, 3, 2, 2, 2, 279, 283,
	7, 44, 2, 2, 280, 282, 5, 46, 24, 2, 281, 280, 3, 2, 2, 2, 282, 285, 3,
	2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 286, 3, 2, 2,
	2, 285, 283, 3, 2, 2, 2, 286, 287, 7, 45, 2, 2, 287, 45, 3, 2, 2, 2, 288,
	292, 5, 10, 6, 2, 289, 292, 5, 48, 25, 2, 290, 292, 5, 70, 36, 2, 291,
	288, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 291, 290, 3, 2, 2, 2, 292, 47, 3,
	2, 2, 2, 293, 294, 5, 72, 37, 2, 294, 296, 7, 39, 2, 2, 295, 297, 7, 52,
	2, 2, 296, 295, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2,
	298, 300, 5, 94, 48, 2, 299, 301, 5, 50, 26, 2, 300, 299, 3, 2, 2, 2, 300,
	301, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 303, 7, 38, 2, 2, 303, 49,
	3, 2, 2, 2, 304, 305, 7, 42, 2, 2, 305, 310, 5, 52, 27, 2, 306, 307, 7,
	49, 2, 2, 307, 309, 5, 52, 27, 2, 308, 306, 3, 2, 2, 2, 309, 312, 3, 2,
	2, 2, 310, 308, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 313, 3, 2, 2, 2,
	312, 310, 3, 2, 2, 2, 313, 314, 7, 43, 2, 2, 314, 51, 3, 2, 2, 2, 315,
	316, 5, 12, 7, 2, 316, 317, 7, 39, 2, 2, 317, 318, 5, 66, 34, 2, 318, 53,
	3, 2, 2, 2, 319, 320, 7, 31, 2, 2, 320, 321, 5, 76, 39, 2, 321, 322, 5,
	56, 29, 2, 322, 55, 3, 2, 2, 2, 323, 327, 7, 44, 2, 2, 324, 326, 5, 58,
	30, 2, 325, 324, 3, 2, 2, 2, 326, 329, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2,
	327, 328, 3, 2, 2, 2, 328, 330, 3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 330,
	331, 7, 45, 2, 2, 331, 57, 3, 2, 2, 2, 332, 341, 5, 14, 8, 2, 333, 341,
	5, 42, 22, 2, 334, 341, 5, 54, 28, 2, 335, 341, 5, 10, 6, 2, 336, 341,
	5, 22, 12, 2, 337, 341, 5, 26, 14, 2, 338, 341, 5, 32, 17, 2, 339, 341,
	5, 70, 36, 2, 340, 332, 3, 2, 2, 2, 340, 333, 3, 2, 2, 2, 340, 334, 3,
	2, 2, 2, 340, 335, 3, 2, 2, 2, 340, 336, 3, 2, 2, 2, 340, 337, 3, 2, 2,
	2, 340, 338, 3, 2, 2, 2, 340, 339, 3, 2, 2, 2, 341, 59, 3, 2, 2, 2, 342,
	343, 7, 32, 2, 2, 343, 344, 5, 86, 44, 2, 344, 348, 7, 44, 2, 2, 345, 347,
	5, 62, 32, 2, 346, 345, 3, 2, 2, 2, 347, 350, 3, 2, 2, 2, 348, 346, 3,
	2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 351, 3, 2, 2, 2, 350, 348, 3, 2, 2,
	2, 351, 352, 7, 45, 2, 2, 352, 61, 3, 2, 2, 2, 353, 357, 5, 10, 6, 2, 354,
	357, 5, 64, 33, 2, 355, 357, 5, 70, 36, 2, 356, 353, 3, 2, 2, 2, 356, 354,
	3, 2, 2, 2, 356, 355, 3, 2, 2, 2, 357, 63, 3, 2, 2, 2, 358, 359, 7, 33,
	2, 2, 359, 360, 5, 88, 45, 2, 360, 362, 7, 40, 2, 2, 361, 363, 7, 34, 2,
	2, 362, 361, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364,
	365, 5, 90, 46, 2, 365, 366, 7, 41, 2, 2, 366, 367, 7, 35, 2, 2, 367, 369,
	7, 40, 2, 2, 368, 370, 7, 34, 2, 2, 369, 368, 3, 2, 2, 2, 369, 370, 3,
	2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 372, 5, 90, 46, 2, 372, 383, 7, 41,
	2, 2, 373, 378, 7, 44, 2, 2, 374, 377, 5, 10, 6, 2, 375, 377, 5, 70, 36,
	2, 376, 374, 3, 2, 2, 2, 376, 375, 3, 2, 2, 2, 377, 380, 3, 2, 2, 2, 378,
	376, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 381, 3, 2, 2, 2, 380, 378,
	3, 2, 2, 2, 381, 384, 7, 45, 2, 2, 382, 384, 7, 38, 2, 2, 383, 373, 3,
	2, 2, 2, 383, 382, 3, 2, 2, 2, 384, 65, 3, 2, 2, 2, 385, 398, 5, 74, 38,
	2, 386, 388, 9, 5, 2, 2, 387, 386, 3, 2, 2, 2, 387, 388, 3, 2, 2, 2, 388,
	389, 3, 2, 2, 2, 389, 398, 5, 94, 48, 2, 390, 392, 9, 5, 2, 2, 391, 390,
	3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 393, 398, 5, 100,
	51, 2, 394, 398, 5, 96, 49, 2, 395, 398, 5, 98, 50, 2, 396, 398, 5, 68,
	35, 2, 397, 385, 3, 2, 2, 2, 397, 387, 3, 2, 2, 2, 397, 391, 3, 2, 2, 2,
	397, 394, 3, 2, 2, 2, 397, 395, 3, 2, 2, 2, 397, 396, 3, 2, 2, 2, 398,
	67, 3, 2, 2, 2, 399, 406, 7, 44, 2, 2, 400, 401, 5, 72, 37, 2, 401, 402,
	7, 50, 2, 2, 402, 403, 5, 66, 34, 2, 403, 405, 3, 2, 2, 2, 404, 400, 3,
	2, 2, 2, 405, 408, 3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 406, 407, 3, 2, 2,
	2, 407, 409, 3, 2, 2, 2, 408, 406, 3, 2, 2, 2, 409, 410, 7, 45, 2, 2, 410,
	69, 3, 2, 2, 2, 411, 412, 7, 38, 2, 2, 412, 71, 3, 2, 2, 2, 413, 416, 7,
	57, 2, 2, 414, 416, 5, 102, 52, 2, 415, 413, 3, 2, 2, 2, 415, 414, 3, 2,
	2, 2, 416, 73, 3, 2, 2, 2, 417, 422, 5, 72, 37, 2, 418, 419, 7, 48, 2,
	2, 419, 421, 5, 72, 37, 2, 420, 418, 3, 2, 2, 2, 421, 424, 3, 2, 2, 2,
	422, 420, 3, 2, 2, 2, 422, 423, 3, 2, 2, 2, 423, 75, 3, 2, 2, 2, 424, 422,
	3, 2, 2, 2, 425, 426, 5, 72, 37, 2, 426, 77, 3, 2, 2, 2, 427, 428, 5, 72,
	37, 2, 428, 79, 3, 2, 2, 2, 429, 430, 5, 72, 37, 2, 430, 81, 3, 2, 2, 2,
	431, 432, 5, 72, 37, 2, 432, 83, 3, 2, 2, 2, 433, 434, 5, 72, 37, 2, 434,
	85, 3, 2, 2, 2, 435, 436, 5, 72, 37, 2, 436, 87, 3, 2, 2, 2, 437, 438,
	5, 72, 37, 2, 438, 89, 3, 2, 2, 2, 439, 441, 7, 48, 2, 2, 440, 439, 3,
	2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 447, 3, 2, 2, 2, 442, 443, 5, 72, 37,
	2, 443, 444, 7, 48, 2, 2, 444, 446, 3, 2, 2, 2, 445, 442, 3, 2, 2, 2, 446,
	449, 3, 2, 2, 2, 447, 445, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 450,
	3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 450, 451, 5, 76, 39, 2, 451, 91, 3, 2,
	2, 2, 452, 454, 7, 48, 2, 2, 453, 452, 3, 2, 2, 2, 453, 454, 3, 2, 2, 2,
	454, 460, 3, 2, 2, 2, 455, 456, 5, 72, 37, 2, 456, 457, 7, 48, 2, 2, 457,
	459, 3, 2, 2, 2, 458, 455, 3, 2, 2, 2, 459, 462, 3, 2, 2, 2, 460, 458,
	3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461, 463, 3, 2, 2, 2, 462, 460, 3, 2,
	2, 2, 463, 464, 5, 78, 40, 2, 464, 93, 3, 2, 2, 2, 465, 466, 7, 56, 2,
	2, 466, 95, 3, 2, 2, 2, 467, 468, 9, 6, 2, 2, 468, 97, 3, 2, 2, 2, 469,
	470, 7, 54, 2, 2, 470, 99, 3, 2, 2, 2, 471, 472, 7, 55, 2, 2, 472, 101,
	3, 2, 2, 2, 473, 474, 9, 7, 2, 2, 474, 103, 3, 2, 2, 2, 45, 110, 112, 122,
	143, 145, 148, 158, 167, 182, 184, 197, 214, 237, 242, 251, 258, 260, 267,
	273, 283, 291, 296, 300, 310, 327, 340, 348, 356, 362, 369, 376, 378, 383,
	387, 391, 397, 406, 415, 422, 440, 447, 453, 460,
}
var literalNames = []string{
	"", "'syntax'", "'import'", "'weak'", "'public'", "'package'", "'option'",
	"'repeated'", "'oneof'", "'map'", "'int32'", "'int64'", "'uint32'", "'uint64'",
	"'sint32'", "'sint64'", "'fixed32'", "'fixed64'", "'sfixed32'", "'sfixed64'",
	"'bool'", "'string'", "'double'", "'float'", "'bytes'", "'reserved'", "'to'",
	"'max'", "'enum'", "'message'", "'service'", "'rpc'", "'stream'", "'returns'",
	"'\"proto3\"'", "''proto3''", "';'", "'='", "'('", "')'", "'['", "']'",
	"'{'", "'}'", "'<'", "'>'", "'.'", "','", "':'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "SYNTAX", "IMPORT", "WEAK", "PUBLIC", "PACKAGE", "OPTION", "REPEATED",
	"ONEOF", "MAP", "INT32", "INT64", "UINT32", "UINT64", "SINT32", "SINT64",
	"FIXED32", "FIXED64", "SFIXED32", "SFIXED64", "BOOL", "STRING", "DOUBLE",
	"FLOAT", "BYTES", "RESERVED", "TO", "MAX", "ENUM", "MESSAGE", "SERVICE",
	"RPC", "STREAM", "RETURNS", "PROTO3_LIT_SINGLE", "PROTO3_LIT_DOBULE", "SEMI",
	"EQ", "LP", "RP", "LB", "RB", "LC", "RC", "LT", "GT", "DOT", "COMMA", "COLON",
	"PLUS", "MINUS", "STR_LIT", "BOOL_LIT", "FLOAT_LIT", "INT_LIT", "IDENTIFIER",
	"WS", "LINE_COMMENT", "COMMENT",
}

var ruleNames = []string{
	"proto", "syntax", "importStatement", "packageStatement", "optionStatement",
	"optionName", "field", "fieldOptions", "fieldOption", "fieldNumber", "oneof",
	"oneofField", "mapField", "keyType", "type_", "reserved", "ranges", "range_",
	"reservedFieldNames", "topLevelDef", "enumDef", "enumBody", "enumElement",
	"enumField", "enumValueOptions", "enumValueOption", "messageDef", "messageBody",
	"messageElement", "serviceDef", "serviceElement", "rpc", "constant", "blockLit",
	"emptyStatement_", "ident", "fullIdent", "messageName", "enumName", "fieldName",
	"oneofName", "mapName", "serviceName", "rpcName", "messageType", "enumType",
	"intLit", "strLit", "boolLit", "floatLit", "keywords",
}

type Protobuf3Parser struct {
	*antlr.BaseParser
}

// NewProtobuf3Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Protobuf3Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewProtobuf3Parser(input antlr.TokenStream) *Protobuf3Parser {
	this := new(Protobuf3Parser)
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
	this.GrammarFileName = "Protobuf3.g4"

	return this
}

// Protobuf3Parser tokens.
const (
	Protobuf3ParserEOF               = antlr.TokenEOF
	Protobuf3ParserSYNTAX            = 1
	Protobuf3ParserIMPORT            = 2
	Protobuf3ParserWEAK              = 3
	Protobuf3ParserPUBLIC            = 4
	Protobuf3ParserPACKAGE           = 5
	Protobuf3ParserOPTION            = 6
	Protobuf3ParserREPEATED          = 7
	Protobuf3ParserONEOF             = 8
	Protobuf3ParserMAP               = 9
	Protobuf3ParserINT32             = 10
	Protobuf3ParserINT64             = 11
	Protobuf3ParserUINT32            = 12
	Protobuf3ParserUINT64            = 13
	Protobuf3ParserSINT32            = 14
	Protobuf3ParserSINT64            = 15
	Protobuf3ParserFIXED32           = 16
	Protobuf3ParserFIXED64           = 17
	Protobuf3ParserSFIXED32          = 18
	Protobuf3ParserSFIXED64          = 19
	Protobuf3ParserBOOL              = 20
	Protobuf3ParserSTRING            = 21
	Protobuf3ParserDOUBLE            = 22
	Protobuf3ParserFLOAT             = 23
	Protobuf3ParserBYTES             = 24
	Protobuf3ParserRESERVED          = 25
	Protobuf3ParserTO                = 26
	Protobuf3ParserMAX               = 27
	Protobuf3ParserENUM              = 28
	Protobuf3ParserMESSAGE           = 29
	Protobuf3ParserSERVICE           = 30
	Protobuf3ParserRPC               = 31
	Protobuf3ParserSTREAM            = 32
	Protobuf3ParserRETURNS           = 33
	Protobuf3ParserPROTO3_LIT_SINGLE = 34
	Protobuf3ParserPROTO3_LIT_DOBULE = 35
	Protobuf3ParserSEMI              = 36
	Protobuf3ParserEQ                = 37
	Protobuf3ParserLP                = 38
	Protobuf3ParserRP                = 39
	Protobuf3ParserLB                = 40
	Protobuf3ParserRB                = 41
	Protobuf3ParserLC                = 42
	Protobuf3ParserRC                = 43
	Protobuf3ParserLT                = 44
	Protobuf3ParserGT                = 45
	Protobuf3ParserDOT               = 46
	Protobuf3ParserCOMMA             = 47
	Protobuf3ParserCOLON             = 48
	Protobuf3ParserPLUS              = 49
	Protobuf3ParserMINUS             = 50
	Protobuf3ParserSTR_LIT           = 51
	Protobuf3ParserBOOL_LIT          = 52
	Protobuf3ParserFLOAT_LIT         = 53
	Protobuf3ParserINT_LIT           = 54
	Protobuf3ParserIDENTIFIER        = 55
	Protobuf3ParserWS                = 56
	Protobuf3ParserLINE_COMMENT      = 57
	Protobuf3ParserCOMMENT           = 58
)

// Protobuf3Parser rules.
const (
	Protobuf3ParserRULE_proto              = 0
	Protobuf3ParserRULE_syntax             = 1
	Protobuf3ParserRULE_importStatement    = 2
	Protobuf3ParserRULE_packageStatement   = 3
	Protobuf3ParserRULE_optionStatement    = 4
	Protobuf3ParserRULE_optionName         = 5
	Protobuf3ParserRULE_field              = 6
	Protobuf3ParserRULE_fieldOptions       = 7
	Protobuf3ParserRULE_fieldOption        = 8
	Protobuf3ParserRULE_fieldNumber        = 9
	Protobuf3ParserRULE_oneof              = 10
	Protobuf3ParserRULE_oneofField         = 11
	Protobuf3ParserRULE_mapField           = 12
	Protobuf3ParserRULE_keyType            = 13
	Protobuf3ParserRULE_type_              = 14
	Protobuf3ParserRULE_reserved           = 15
	Protobuf3ParserRULE_ranges             = 16
	Protobuf3ParserRULE_range_             = 17
	Protobuf3ParserRULE_reservedFieldNames = 18
	Protobuf3ParserRULE_topLevelDef        = 19
	Protobuf3ParserRULE_enumDef            = 20
	Protobuf3ParserRULE_enumBody           = 21
	Protobuf3ParserRULE_enumElement        = 22
	Protobuf3ParserRULE_enumField          = 23
	Protobuf3ParserRULE_enumValueOptions   = 24
	Protobuf3ParserRULE_enumValueOption    = 25
	Protobuf3ParserRULE_messageDef         = 26
	Protobuf3ParserRULE_messageBody        = 27
	Protobuf3ParserRULE_messageElement     = 28
	Protobuf3ParserRULE_serviceDef         = 29
	Protobuf3ParserRULE_serviceElement     = 30
	Protobuf3ParserRULE_rpc                = 31
	Protobuf3ParserRULE_constant           = 32
	Protobuf3ParserRULE_blockLit           = 33
	Protobuf3ParserRULE_emptyStatement_    = 34
	Protobuf3ParserRULE_ident              = 35
	Protobuf3ParserRULE_fullIdent          = 36
	Protobuf3ParserRULE_messageName        = 37
	Protobuf3ParserRULE_enumName           = 38
	Protobuf3ParserRULE_fieldName          = 39
	Protobuf3ParserRULE_oneofName          = 40
	Protobuf3ParserRULE_mapName            = 41
	Protobuf3ParserRULE_serviceName        = 42
	Protobuf3ParserRULE_rpcName            = 43
	Protobuf3ParserRULE_messageType        = 44
	Protobuf3ParserRULE_enumType           = 45
	Protobuf3ParserRULE_intLit             = 46
	Protobuf3ParserRULE_strLit             = 47
	Protobuf3ParserRULE_boolLit            = 48
	Protobuf3ParserRULE_floatLit           = 49
	Protobuf3ParserRULE_keywords           = 50
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_proto
	return p
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) AllImportStatement() []IImportStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportStatementContext)(nil)).Elem())
	var tst = make([]IImportStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportStatementContext)
		}
	}

	return tst
}

func (s *ProtoContext) ImportStatement(i int) IImportStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *ProtoContext) AllPackageStatement() []IPackageStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPackageStatementContext)(nil)).Elem())
	var tst = make([]IPackageStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPackageStatementContext)
		}
	}

	return tst
}

func (s *ProtoContext) PackageStatement(i int) IPackageStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *ProtoContext) AllOptionStatement() []IOptionStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem())
	var tst = make([]IOptionStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionStatementContext)
		}
	}

	return tst
}

func (s *ProtoContext) OptionStatement(i int) IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *ProtoContext) AllTopLevelDef() []ITopLevelDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopLevelDefContext)(nil)).Elem())
	var tst = make([]ITopLevelDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopLevelDefContext)
		}
	}

	return tst
}

func (s *ProtoContext) TopLevelDef(i int) ITopLevelDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopLevelDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopLevelDefContext)
}

func (s *ProtoContext) AllEmptyStatement_() []IEmptyStatement_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem())
	var tst = make([]IEmptyStatement_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatement_Context)
		}
	}

	return tst
}

func (s *ProtoContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitProto(s)
	}
}

func (p *Protobuf3Parser) Proto() (localctx IProtoContext) {
	this := p
	_ = this

	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Protobuf3ParserRULE_proto)
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
		p.SetState(102)
		p.Syntax()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserIMPORT)|(1<<Protobuf3ParserPACKAGE)|(1<<Protobuf3ParserOPTION)|(1<<Protobuf3ParserENUM)|(1<<Protobuf3ParserMESSAGE)|(1<<Protobuf3ParserSERVICE))) != 0) || _la == Protobuf3ParserSEMI {
		p.SetState(108)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserIMPORT:
			{
				p.SetState(103)
				p.ImportStatement()
			}

		case Protobuf3ParserPACKAGE:
			{
				p.SetState(104)
				p.PackageStatement()
			}

		case Protobuf3ParserOPTION:
			{
				p.SetState(105)
				p.OptionStatement()
			}

		case Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE:
			{
				p.SetState(106)
				p.TopLevelDef()
			}

		case Protobuf3ParserSEMI:
			{
				p.SetState(107)
				p.EmptyStatement_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSYNTAX, 0)
}

func (s *SyntaxContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *SyntaxContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *SyntaxContext) PROTO3_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_SINGLE, 0)
}

func (s *SyntaxContext) PROTO3_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_DOBULE, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (p *Protobuf3Parser) Syntax() (localctx ISyntaxContext) {
	this := p
	_ = this

	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Protobuf3ParserRULE_syntax)
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
		p.SetState(113)
		p.Match(Protobuf3ParserSYNTAX)
	}
	{
		p.SetState(114)
		p.Match(Protobuf3ParserEQ)
	}
	{
		p.SetState(115)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Protobuf3ParserPROTO3_LIT_SINGLE || _la == Protobuf3ParserPROTO3_LIT_DOBULE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(116)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIMPORT, 0)
}

func (s *ImportStatementContext) StrLit() IStrLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *ImportStatementContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserWEAK, 0)
}

func (s *ImportStatementContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPUBLIC, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *Protobuf3Parser) ImportStatement() (localctx IImportStatementContext) {
	this := p
	_ = this

	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Protobuf3ParserRULE_importStatement)
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
		p.Match(Protobuf3ParserIMPORT)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserWEAK || _la == Protobuf3ParserPUBLIC {
		{
			p.SetState(119)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Protobuf3ParserWEAK || _la == Protobuf3ParserPUBLIC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(122)
		p.StrLit()
	}
	{
		p.SetState(123)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_packageStatement
	return p
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPACKAGE, 0)
}

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (p *Protobuf3Parser) PackageStatement() (localctx IPackageStatementContext) {
	this := p
	_ = this

	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Protobuf3ParserRULE_packageStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(125)
		p.Match(Protobuf3ParserPACKAGE)
	}
	{
		p.SetState(126)
		p.FullIdent()
	}
	{
		p.SetState(127)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IOptionStatementContext is an interface to support dynamic dispatch.
type IOptionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionStatementContext differentiates from other interfaces.
	IsOptionStatementContext()
}

type OptionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionStatementContext() *OptionStatementContext {
	var p = new(OptionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionStatement
	return p
}

func (*OptionStatementContext) IsOptionStatementContext() {}

func NewOptionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionStatementContext {
	var p = new(OptionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionStatement

	return p
}

func (s *OptionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionStatementContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTION, 0)
}

func (s *OptionStatementContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *OptionStatementContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *OptionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOptionStatement(s)
	}
}

func (s *OptionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOptionStatement(s)
	}
}

func (p *Protobuf3Parser) OptionStatement() (localctx IOptionStatementContext) {
	this := p
	_ = this

	localctx = NewOptionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Protobuf3ParserRULE_optionStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(129)
		p.Match(Protobuf3ParserOPTION)
	}
	{
		p.SetState(130)
		p.OptionName()
	}
	{
		p.SetState(131)
		p.Match(Protobuf3ParserEQ)
	}
	{
		p.SetState(132)
		p.Constant()
	}
	{
		p.SetState(133)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionName
	return p
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllFullIdent() []IFullIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFullIdentContext)(nil)).Elem())
	var tst = make([]IFullIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFullIdentContext)
		}
	}

	return tst
}

func (s *OptionNameContext) FullIdent(i int) IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) LP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLP, 0)
}

func (s *OptionNameContext) RP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRP, 0)
}

func (s *OptionNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, 0)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (p *Protobuf3Parser) OptionName() (localctx IOptionNameContext) {
	this := p
	_ = this

	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Protobuf3ParserRULE_optionName)
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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserSYNTAX, Protobuf3ParserIMPORT, Protobuf3ParserWEAK, Protobuf3ParserPUBLIC, Protobuf3ParserPACKAGE, Protobuf3ParserOPTION, Protobuf3ParserREPEATED, Protobuf3ParserONEOF, Protobuf3ParserMAP, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserBOOL, Protobuf3ParserSTRING, Protobuf3ParserDOUBLE, Protobuf3ParserFLOAT, Protobuf3ParserBYTES, Protobuf3ParserRESERVED, Protobuf3ParserTO, Protobuf3ParserMAX, Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE, Protobuf3ParserRPC, Protobuf3ParserSTREAM, Protobuf3ParserRETURNS, Protobuf3ParserBOOL_LIT, Protobuf3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.FullIdent()
		}

	case Protobuf3ParserLP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Match(Protobuf3ParserLP)
		}
		{
			p.SetState(137)
			p.FullIdent()
		}
		{
			p.SetState(138)
			p.Match(Protobuf3ParserRP)
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserDOT {
			{
				p.SetState(139)
				p.Match(Protobuf3ParserDOT)
			}
			{
				p.SetState(140)
				p.FullIdent()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = Protobuf3ParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *FieldContext) FieldNumber() IFieldNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *FieldContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserREPEATED, 0)
}

func (s *FieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
}

func (s *FieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitField(s)
	}
}

func (p *Protobuf3Parser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Protobuf3ParserRULE_field)
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
	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(145)
			p.Match(Protobuf3ParserREPEATED)
		}

	}
	{
		p.SetState(148)
		p.Type_()
	}
	{
		p.SetState(149)
		p.FieldName()
	}
	{
		p.SetState(150)
		p.Match(Protobuf3ParserEQ)
	}
	{
		p.SetState(151)
		p.FieldNumber()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(152)
			p.Match(Protobuf3ParserLB)
		}
		{
			p.SetState(153)
			p.FieldOptions()
		}
		{
			p.SetState(154)
			p.Match(Protobuf3ParserRB)
		}

	}
	{
		p.SetState(158)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IFieldOptionsContext is an interface to support dynamic dispatch.
type IFieldOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionsContext differentiates from other interfaces.
	IsFieldOptionsContext()
}

type FieldOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionsContext() *FieldOptionsContext {
	var p = new(FieldOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions
	return p
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) AllFieldOption() []IFieldOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldOptionContext)(nil)).Elem())
	var tst = make([]IFieldOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldOptionContext)
		}
	}

	return tst
}

func (s *FieldOptionsContext) FieldOption(i int) IFieldOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionContext)
}

func (s *FieldOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldOptions(s)
	}
}

func (s *FieldOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldOptions(s)
	}
}

func (p *Protobuf3Parser) FieldOptions() (localctx IFieldOptionsContext) {
	this := p
	_ = this

	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Protobuf3ParserRULE_fieldOptions)
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
		p.SetState(160)
		p.FieldOption()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(161)
			p.Match(Protobuf3ParserCOMMA)
		}
		{
			p.SetState(162)
			p.FieldOption()
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldOptionContext is an interface to support dynamic dispatch.
type IFieldOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldOptionContext differentiates from other interfaces.
	IsFieldOptionContext()
}

type FieldOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionContext() *FieldOptionContext {
	var p = new(FieldOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOption
	return p
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOption

	return p
}

func (s *FieldOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FieldOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *FieldOptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldOption(s)
	}
}

func (s *FieldOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldOption(s)
	}
}

func (p *Protobuf3Parser) FieldOption() (localctx IFieldOptionContext) {
	this := p
	_ = this

	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Protobuf3ParserRULE_fieldOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(168)
		p.OptionName()
	}
	{
		p.SetState(169)
		p.Match(Protobuf3ParserEQ)
	}
	{
		p.SetState(170)
		p.Constant()
	}

	return localctx
}

// IFieldNumberContext is an interface to support dynamic dispatch.
type IFieldNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNumberContext differentiates from other interfaces.
	IsFieldNumberContext()
}

type FieldNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNumberContext() *FieldNumberContext {
	var p = new(FieldNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber
	return p
}

func (*FieldNumberContext) IsFieldNumberContext() {}

func NewFieldNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNumberContext {
	var p = new(FieldNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber

	return p
}

func (s *FieldNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNumberContext) IntLit() IIntLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *FieldNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldNumber(s)
	}
}

func (s *FieldNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldNumber(s)
	}
}

func (p *Protobuf3Parser) FieldNumber() (localctx IFieldNumberContext) {
	this := p
	_ = this

	localctx = NewFieldNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Protobuf3ParserRULE_fieldNumber)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(172)
		p.IntLit()
	}

	return localctx
}

// IOneofContext is an interface to support dynamic dispatch.
type IOneofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofContext differentiates from other interfaces.
	IsOneofContext()
}

type OneofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofContext() *OneofContext {
	var p = new(OneofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneof
	return p
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserONEOF, 0)
}

func (s *OneofContext) OneofName() IOneofNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneofNameContext)
}

func (s *OneofContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *OneofContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *OneofContext) AllOptionStatement() []IOptionStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem())
	var tst = make([]IOptionStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionStatementContext)
		}
	}

	return tst
}

func (s *OneofContext) OptionStatement(i int) IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *OneofContext) AllOneofField() []IOneofFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem())
	var tst = make([]IOneofFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOneofFieldContext)
		}
	}

	return tst
}

func (s *OneofContext) OneofField(i int) IOneofFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOneofFieldContext)
}

func (s *OneofContext) AllEmptyStatement_() []IEmptyStatement_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem())
	var tst = make([]IEmptyStatement_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatement_Context)
		}
	}

	return tst
}

func (s *OneofContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *OneofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneof(s)
	}
}

func (s *OneofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneof(s)
	}
}

func (p *Protobuf3Parser) Oneof() (localctx IOneofContext) {
	this := p
	_ = this

	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Protobuf3ParserRULE_oneof)
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
		p.SetState(174)
		p.Match(Protobuf3ParserONEOF)
	}
	{
		p.SetState(175)
		p.OneofName()
	}
	{
		p.SetState(176)
		p.Match(Protobuf3ParserLC)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserSYNTAX)|(1<<Protobuf3ParserIMPORT)|(1<<Protobuf3ParserWEAK)|(1<<Protobuf3ParserPUBLIC)|(1<<Protobuf3ParserPACKAGE)|(1<<Protobuf3ParserOPTION)|(1<<Protobuf3ParserREPEATED)|(1<<Protobuf3ParserONEOF)|(1<<Protobuf3ParserMAP)|(1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserUINT32)|(1<<Protobuf3ParserUINT64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserSTRING)|(1<<Protobuf3ParserDOUBLE)|(1<<Protobuf3ParserFLOAT)|(1<<Protobuf3ParserBYTES)|(1<<Protobuf3ParserRESERVED)|(1<<Protobuf3ParserTO)|(1<<Protobuf3ParserMAX)|(1<<Protobuf3ParserENUM)|(1<<Protobuf3ParserMESSAGE)|(1<<Protobuf3ParserSERVICE)|(1<<Protobuf3ParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(Protobuf3ParserSTREAM-32))|(1<<(Protobuf3ParserRETURNS-32))|(1<<(Protobuf3ParserSEMI-32))|(1<<(Protobuf3ParserDOT-32))|(1<<(Protobuf3ParserBOOL_LIT-32))|(1<<(Protobuf3ParserIDENTIFIER-32)))) != 0) {
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(177)
				p.OptionStatement()
			}

		case 2:
			{
				p.SetState(178)
				p.OneofField()
			}

		case 3:
			{
				p.SetState(179)
				p.EmptyStatement_()
			}

		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(185)
		p.Match(Protobuf3ParserRC)
	}

	return localctx
}

// IOneofFieldContext is an interface to support dynamic dispatch.
type IOneofFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofFieldContext differentiates from other interfaces.
	IsOneofFieldContext()
}

type OneofFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofFieldContext() *OneofFieldContext {
	var p = new(OneofFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofField
	return p
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofField

	return p
}

func (s *OneofFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofFieldContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *OneofFieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *OneofFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *OneofFieldContext) FieldNumber() IFieldNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *OneofFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *OneofFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
}

func (s *OneofFieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *OneofFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *OneofFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneofField(s)
	}
}

func (s *OneofFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneofField(s)
	}
}

func (p *Protobuf3Parser) OneofField() (localctx IOneofFieldContext) {
	this := p
	_ = this

	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Protobuf3ParserRULE_oneofField)
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
		p.SetState(187)
		p.Type_()
	}
	{
		p.SetState(188)
		p.FieldName()
	}
	{
		p.SetState(189)
		p.Match(Protobuf3ParserEQ)
	}
	{
		p.SetState(190)
		p.FieldNumber()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(191)
			p.Match(Protobuf3ParserLB)
		}
		{
			p.SetState(192)
			p.FieldOptions()
		}
		{
			p.SetState(193)
			p.Match(Protobuf3ParserRB)
		}

	}
	{
		p.SetState(197)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapField
	return p
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAP, 0)
}

func (s *MapFieldContext) LT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLT, 0)
}

func (s *MapFieldContext) KeyType() IKeyTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, 0)
}

func (s *MapFieldContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MapFieldContext) GT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserGT, 0)
}

func (s *MapFieldContext) MapName() IMapNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapNameContext)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *MapFieldContext) FieldNumber() IFieldNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *MapFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *MapFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
}

func (s *MapFieldContext) FieldOptions() IFieldOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *MapFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMapField(s)
	}
}

func (p *Protobuf3Parser) MapField() (localctx IMapFieldContext) {
	this := p
	_ = this

	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Protobuf3ParserRULE_mapField)
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
		p.Match(Protobuf3ParserMAP)
	}
	{
		p.SetState(200)
		p.Match(Protobuf3ParserLT)
	}
	{
		p.SetState(201)
		p.KeyType()
	}
	{
		p.SetState(202)
		p.Match(Protobuf3ParserCOMMA)
	}
	{
		p.SetState(203)
		p.Type_()
	}
	{
		p.SetState(204)
		p.Match(Protobuf3ParserGT)
	}
	{
		p.SetState(205)
		p.MapName()
	}
	{
		p.SetState(206)
		p.Match(Protobuf3ParserEQ)
	}
	{
		p.SetState(207)
		p.FieldNumber()
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(208)
			p.Match(Protobuf3ParserLB)
		}
		{
			p.SetState(209)
			p.FieldOptions()
		}
		{
			p.SetState(210)
			p.Match(Protobuf3ParserRB)
		}

	}
	{
		p.SetState(214)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keyType
	return p
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (p *Protobuf3Parser) KeyType() (localctx IKeyTypeContext) {
	this := p
	_ = this

	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Protobuf3ParserRULE_keyType)
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
		p.SetState(216)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserUINT32)|(1<<Protobuf3ParserUINT64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.RuleIndex = Protobuf3ParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOUBLE, 0)
}

func (s *Type_Context) FLOAT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT, 0)
}

func (s *Type_Context) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *Type_Context) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *Type_Context) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *Type_Context) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *Type_Context) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *Type_Context) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *Type_Context) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *Type_Context) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *Type_Context) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *Type_Context) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *Type_Context) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *Type_Context) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *Type_Context) BYTES() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBYTES, 0)
}

func (s *Type_Context) MessageType() IMessageTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *Type_Context) EnumType() IEnumTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *Protobuf3Parser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Protobuf3ParserRULE_type_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.Match(Protobuf3ParserDOUBLE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Match(Protobuf3ParserFLOAT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.Match(Protobuf3ParserINT32)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(221)
			p.Match(Protobuf3ParserINT64)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(222)
			p.Match(Protobuf3ParserUINT32)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(223)
			p.Match(Protobuf3ParserUINT64)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(224)
			p.Match(Protobuf3ParserSINT32)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(225)
			p.Match(Protobuf3ParserSINT64)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(226)
			p.Match(Protobuf3ParserFIXED32)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(227)
			p.Match(Protobuf3ParserFIXED64)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(228)
			p.Match(Protobuf3ParserSFIXED32)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(229)
			p.Match(Protobuf3ParserSFIXED64)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(230)
			p.Match(Protobuf3ParserBOOL)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(231)
			p.Match(Protobuf3ParserSTRING)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(232)
			p.Match(Protobuf3ParserBYTES)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(233)
			p.MessageType()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(234)
			p.EnumType()
		}

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
	p.RuleIndex = Protobuf3ParserRULE_reserved
	return p
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRESERVED, 0)
}

func (s *ReservedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *ReservedContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ReservedContext) ReservedFieldNames() IReservedFieldNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedFieldNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedFieldNamesContext)
}

func (s *ReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterReserved(s)
	}
}

func (s *ReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitReserved(s)
	}
}

func (p *Protobuf3Parser) Reserved() (localctx IReservedContext) {
	this := p
	_ = this

	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Protobuf3ParserRULE_reserved)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(237)
		p.Match(Protobuf3ParserRESERVED)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserINT_LIT:
		{
			p.SetState(238)
			p.Ranges()
		}

	case Protobuf3ParserPROTO3_LIT_SINGLE, Protobuf3ParserPROTO3_LIT_DOBULE, Protobuf3ParserSTR_LIT:
		{
			p.SetState(239)
			p.ReservedFieldNames()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(242)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) AllRange_() []IRange_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRange_Context)(nil)).Elem())
	var tst = make([]IRange_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRange_Context)
		}
	}

	return tst
}

func (s *RangesContext) Range_(i int) IRange_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRange_Context)
}

func (s *RangesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *RangesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRanges(s)
	}
}

func (p *Protobuf3Parser) Ranges() (localctx IRangesContext) {
	this := p
	_ = this

	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Protobuf3ParserRULE_ranges)
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
		p.SetState(244)
		p.Range_()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(245)
			p.Match(Protobuf3ParserCOMMA)
		}
		{
			p.SetState(246)
			p.Range_()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRange_Context is an interface to support dynamic dispatch.
type IRange_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRange_Context differentiates from other interfaces.
	IsRange_Context()
}

type Range_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_Context() *Range_Context {
	var p = new(Range_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_range_
	return p
}

func (*Range_Context) IsRange_Context() {}

func NewRange_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_Context {
	var p = new(Range_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_range_

	return p
}

func (s *Range_Context) GetParser() antlr.Parser { return s.parser }

func (s *Range_Context) AllIntLit() []IIntLitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntLitContext)(nil)).Elem())
	var tst = make([]IIntLitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntLitContext)
		}
	}

	return tst
}

func (s *Range_Context) IntLit(i int) IIntLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *Range_Context) TO() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserTO, 0)
}

func (s *Range_Context) MAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAX, 0)
}

func (s *Range_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRange_(s)
	}
}

func (s *Range_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRange_(s)
	}
}

func (p *Protobuf3Parser) Range_() (localctx IRange_Context) {
	this := p
	_ = this

	localctx = NewRange_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Protobuf3ParserRULE_range_)
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
		p.SetState(252)
		p.IntLit()
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserTO {
		{
			p.SetState(253)
			p.Match(Protobuf3ParserTO)
		}
		p.SetState(256)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserINT_LIT:
			{
				p.SetState(254)
				p.IntLit()
			}

		case Protobuf3ParserMAX:
			{
				p.SetState(255)
				p.Match(Protobuf3ParserMAX)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IReservedFieldNamesContext is an interface to support dynamic dispatch.
type IReservedFieldNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedFieldNamesContext differentiates from other interfaces.
	IsReservedFieldNamesContext()
}

type ReservedFieldNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedFieldNamesContext() *ReservedFieldNamesContext {
	var p = new(ReservedFieldNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reservedFieldNames
	return p
}

func (*ReservedFieldNamesContext) IsReservedFieldNamesContext() {}

func NewReservedFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedFieldNamesContext {
	var p = new(ReservedFieldNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_reservedFieldNames

	return p
}

func (s *ReservedFieldNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedFieldNamesContext) AllStrLit() []IStrLitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStrLitContext)(nil)).Elem())
	var tst = make([]IStrLitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStrLitContext)
		}
	}

	return tst
}

func (s *ReservedFieldNamesContext) StrLit(i int) IStrLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrLitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ReservedFieldNamesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *ReservedFieldNamesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *ReservedFieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedFieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedFieldNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterReservedFieldNames(s)
	}
}

func (s *ReservedFieldNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitReservedFieldNames(s)
	}
}

func (p *Protobuf3Parser) ReservedFieldNames() (localctx IReservedFieldNamesContext) {
	this := p
	_ = this

	localctx = NewReservedFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Protobuf3ParserRULE_reservedFieldNames)
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
		p.SetState(260)
		p.StrLit()
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(261)
			p.Match(Protobuf3ParserCOMMA)
		}
		{
			p.SetState(262)
			p.StrLit()
		}

		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITopLevelDefContext is an interface to support dynamic dispatch.
type ITopLevelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopLevelDefContext differentiates from other interfaces.
	IsTopLevelDefContext()
}

type TopLevelDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDefContext() *TopLevelDefContext {
	var p = new(TopLevelDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef
	return p
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef

	return p
}

func (s *TopLevelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDefContext) MessageDef() IMessageDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageDefContext)
}

func (s *TopLevelDefContext) EnumDef() IEnumDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *TopLevelDefContext) ServiceDef() IServiceDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceDefContext)
}

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitTopLevelDef(s)
	}
}

func (p *Protobuf3Parser) TopLevelDef() (localctx ITopLevelDefContext) {
	this := p
	_ = this

	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Protobuf3ParserRULE_topLevelDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(271)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserMESSAGE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.MessageDef()
		}

	case Protobuf3ParserENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(269)
			p.EnumDef()
		}

	case Protobuf3ParserSERVICE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(270)
			p.ServiceDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumDef
	return p
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserENUM, 0)
}

func (s *EnumDefContext) EnumName() IEnumNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (p *Protobuf3Parser) EnumDef() (localctx IEnumDefContext) {
	this := p
	_ = this

	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Protobuf3ParserRULE_enumDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(273)
		p.Match(Protobuf3ParserENUM)
	}
	{
		p.SetState(274)
		p.EnumName()
	}
	{
		p.SetState(275)
		p.EnumBody()
	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *EnumBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *EnumBodyContext) AllEnumElement() []IEnumElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumElementContext)(nil)).Elem())
	var tst = make([]IEnumElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumElementContext)
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumElement(i int) IEnumElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumElementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *Protobuf3Parser) EnumBody() (localctx IEnumBodyContext) {
	this := p
	_ = this

	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Protobuf3ParserRULE_enumBody)
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
		p.Match(Protobuf3ParserLC)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserSYNTAX)|(1<<Protobuf3ParserIMPORT)|(1<<Protobuf3ParserWEAK)|(1<<Protobuf3ParserPUBLIC)|(1<<Protobuf3ParserPACKAGE)|(1<<Protobuf3ParserOPTION)|(1<<Protobuf3ParserREPEATED)|(1<<Protobuf3ParserONEOF)|(1<<Protobuf3ParserMAP)|(1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserUINT32)|(1<<Protobuf3ParserUINT64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserSTRING)|(1<<Protobuf3ParserDOUBLE)|(1<<Protobuf3ParserFLOAT)|(1<<Protobuf3ParserBYTES)|(1<<Protobuf3ParserRESERVED)|(1<<Protobuf3ParserTO)|(1<<Protobuf3ParserMAX)|(1<<Protobuf3ParserENUM)|(1<<Protobuf3ParserMESSAGE)|(1<<Protobuf3ParserSERVICE)|(1<<Protobuf3ParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(Protobuf3ParserSTREAM-32))|(1<<(Protobuf3ParserRETURNS-32))|(1<<(Protobuf3ParserSEMI-32))|(1<<(Protobuf3ParserBOOL_LIT-32))|(1<<(Protobuf3ParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(278)
			p.EnumElement()
		}

		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(284)
		p.Match(Protobuf3ParserRC)
	}

	return localctx
}

// IEnumElementContext is an interface to support dynamic dispatch.
type IEnumElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumElementContext differentiates from other interfaces.
	IsEnumElementContext()
}

type EnumElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumElementContext() *EnumElementContext {
	var p = new(EnumElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumElement
	return p
}

func (*EnumElementContext) IsEnumElementContext() {}

func NewEnumElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumElementContext {
	var p = new(EnumElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumElement

	return p
}

func (s *EnumElementContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumElementContext) OptionStatement() IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *EnumElementContext) EnumField() IEnumFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *EnumElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumElement(s)
	}
}

func (s *EnumElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumElement(s)
	}
}

func (p *Protobuf3Parser) EnumElement() (localctx IEnumElementContext) {
	this := p
	_ = this

	localctx = NewEnumElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Protobuf3ParserRULE_enumElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.OptionStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.EnumField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(288)
			p.EmptyStatement_()
		}

	}

	return localctx
}

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumField
	return p
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *EnumFieldContext) IntLit() IIntLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *EnumFieldContext) MINUS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMINUS, 0)
}

func (s *EnumFieldContext) EnumValueOptions() IEnumValueOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionsContext)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (p *Protobuf3Parser) EnumField() (localctx IEnumFieldContext) {
	this := p
	_ = this

	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Protobuf3ParserRULE_enumField)
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
		p.SetState(291)
		p.Ident()
	}
	{
		p.SetState(292)
		p.Match(Protobuf3ParserEQ)
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserMINUS {
		{
			p.SetState(293)
			p.Match(Protobuf3ParserMINUS)
		}

	}
	{
		p.SetState(296)
		p.IntLit()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(297)
			p.EnumValueOptions()
		}

	}
	{
		p.SetState(300)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IEnumValueOptionsContext is an interface to support dynamic dispatch.
type IEnumValueOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueOptionsContext differentiates from other interfaces.
	IsEnumValueOptionsContext()
}

type EnumValueOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionsContext() *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOptions
	return p
}

func (*EnumValueOptionsContext) IsEnumValueOptionsContext() {}

func NewEnumValueOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumValueOptions

	return p
}

func (s *EnumValueOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionsContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
}

func (s *EnumValueOptionsContext) AllEnumValueOption() []IEnumValueOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem())
	var tst = make([]IEnumValueOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueOptionContext)
		}
	}

	return tst
}

func (s *EnumValueOptionsContext) EnumValueOption(i int) IEnumValueOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionContext)
}

func (s *EnumValueOptionsContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *EnumValueOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *EnumValueOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *EnumValueOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumValueOptions(s)
	}
}

func (s *EnumValueOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumValueOptions(s)
	}
}

func (p *Protobuf3Parser) EnumValueOptions() (localctx IEnumValueOptionsContext) {
	this := p
	_ = this

	localctx = NewEnumValueOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, Protobuf3ParserRULE_enumValueOptions)
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
		p.SetState(302)
		p.Match(Protobuf3ParserLB)
	}
	{
		p.SetState(303)
		p.EnumValueOption()
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(304)
			p.Match(Protobuf3ParserCOMMA)
		}
		{
			p.SetState(305)
			p.EnumValueOption()
		}

		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(311)
		p.Match(Protobuf3ParserRB)
	}

	return localctx
}

// IEnumValueOptionContext is an interface to support dynamic dispatch.
type IEnumValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueOptionContext differentiates from other interfaces.
	IsEnumValueOptionContext()
}

type EnumValueOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionContext() *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption
	return p
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption

	return p
}

func (s *EnumValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionContext) OptionName() IOptionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *EnumValueOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *EnumValueOptionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumValueOption(s)
	}
}

func (p *Protobuf3Parser) EnumValueOption() (localctx IEnumValueOptionContext) {
	this := p
	_ = this

	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Protobuf3ParserRULE_enumValueOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(313)
		p.OptionName()
	}
	{
		p.SetState(314)
		p.Match(Protobuf3ParserEQ)
	}
	{
		p.SetState(315)
		p.Constant()
	}

	return localctx
}

// IMessageDefContext is an interface to support dynamic dispatch.
type IMessageDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageDefContext differentiates from other interfaces.
	IsMessageDefContext()
}

type MessageDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageDefContext() *MessageDefContext {
	var p = new(MessageDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageDef
	return p
}

func (*MessageDefContext) IsMessageDefContext() {}

func NewMessageDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDefContext {
	var p = new(MessageDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageDef

	return p
}

func (s *MessageDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDefContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMESSAGE, 0)
}

func (s *MessageDefContext) MessageName() IMessageNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageDefContext) MessageBody() IMessageBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageDef(s)
	}
}

func (s *MessageDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageDef(s)
	}
}

func (p *Protobuf3Parser) MessageDef() (localctx IMessageDefContext) {
	this := p
	_ = this

	localctx = NewMessageDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Protobuf3ParserRULE_messageDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(317)
		p.Match(Protobuf3ParserMESSAGE)
	}
	{
		p.SetState(318)
		p.MessageName()
	}
	{
		p.SetState(319)
		p.MessageBody()
	}

	return localctx
}

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageBody
	return p
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *MessageBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *MessageBodyContext) AllMessageElement() []IMessageElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageElementContext)(nil)).Elem())
	var tst = make([]IMessageElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageElementContext)
		}
	}

	return tst
}

func (s *MessageBodyContext) MessageElement(i int) IMessageElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageElementContext)
}

func (s *MessageBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageBody(s)
	}
}

func (s *MessageBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageBody(s)
	}
}

func (p *Protobuf3Parser) MessageBody() (localctx IMessageBodyContext) {
	this := p
	_ = this

	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Protobuf3ParserRULE_messageBody)
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
		p.SetState(321)
		p.Match(Protobuf3ParserLC)
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserSYNTAX)|(1<<Protobuf3ParserIMPORT)|(1<<Protobuf3ParserWEAK)|(1<<Protobuf3ParserPUBLIC)|(1<<Protobuf3ParserPACKAGE)|(1<<Protobuf3ParserOPTION)|(1<<Protobuf3ParserREPEATED)|(1<<Protobuf3ParserONEOF)|(1<<Protobuf3ParserMAP)|(1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserUINT32)|(1<<Protobuf3ParserUINT64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserSTRING)|(1<<Protobuf3ParserDOUBLE)|(1<<Protobuf3ParserFLOAT)|(1<<Protobuf3ParserBYTES)|(1<<Protobuf3ParserRESERVED)|(1<<Protobuf3ParserTO)|(1<<Protobuf3ParserMAX)|(1<<Protobuf3ParserENUM)|(1<<Protobuf3ParserMESSAGE)|(1<<Protobuf3ParserSERVICE)|(1<<Protobuf3ParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(Protobuf3ParserSTREAM-32))|(1<<(Protobuf3ParserRETURNS-32))|(1<<(Protobuf3ParserSEMI-32))|(1<<(Protobuf3ParserDOT-32))|(1<<(Protobuf3ParserBOOL_LIT-32))|(1<<(Protobuf3ParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(322)
			p.MessageElement()
		}

		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(328)
		p.Match(Protobuf3ParserRC)
	}

	return localctx
}

// IMessageElementContext is an interface to support dynamic dispatch.
type IMessageElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageElementContext differentiates from other interfaces.
	IsMessageElementContext()
}

type MessageElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageElementContext() *MessageElementContext {
	var p = new(MessageElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageElement
	return p
}

func (*MessageElementContext) IsMessageElementContext() {}

func NewMessageElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageElementContext {
	var p = new(MessageElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageElement

	return p
}

func (s *MessageElementContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageElementContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageElementContext) EnumDef() IEnumDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *MessageElementContext) MessageDef() IMessageDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageDefContext)
}

func (s *MessageElementContext) OptionStatement() IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *MessageElementContext) Oneof() IOneofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneofContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneofContext)
}

func (s *MessageElementContext) MapField() IMapFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *MessageElementContext) Reserved() IReservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedContext)
}

func (s *MessageElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *MessageElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageElement(s)
	}
}

func (s *MessageElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageElement(s)
	}
}

func (p *Protobuf3Parser) MessageElement() (localctx IMessageElementContext) {
	this := p
	_ = this

	localctx = NewMessageElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Protobuf3ParserRULE_messageElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.EnumDef()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(332)
			p.MessageDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(333)
			p.OptionStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(334)
			p.Oneof()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(335)
			p.MapField()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(336)
			p.Reserved()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(337)
			p.EmptyStatement_()
		}

	}

	return localctx
}

// IServiceDefContext is an interface to support dynamic dispatch.
type IServiceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceDefContext differentiates from other interfaces.
	IsServiceDefContext()
}

type ServiceDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceDefContext() *ServiceDefContext {
	var p = new(ServiceDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceDef
	return p
}

func (*ServiceDefContext) IsServiceDefContext() {}

func NewServiceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDefContext {
	var p = new(ServiceDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceDef

	return p
}

func (s *ServiceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDefContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSERVICE, 0)
}

func (s *ServiceDefContext) ServiceName() IServiceNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceDefContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *ServiceDefContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *ServiceDefContext) AllServiceElement() []IServiceElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IServiceElementContext)(nil)).Elem())
	var tst = make([]IServiceElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IServiceElementContext)
		}
	}

	return tst
}

func (s *ServiceDefContext) ServiceElement(i int) IServiceElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IServiceElementContext)
}

func (s *ServiceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterServiceDef(s)
	}
}

func (s *ServiceDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitServiceDef(s)
	}
}

func (p *Protobuf3Parser) ServiceDef() (localctx IServiceDefContext) {
	this := p
	_ = this

	localctx = NewServiceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Protobuf3ParserRULE_serviceDef)
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
		p.SetState(340)
		p.Match(Protobuf3ParserSERVICE)
	}
	{
		p.SetState(341)
		p.ServiceName()
	}
	{
		p.SetState(342)
		p.Match(Protobuf3ParserLC)
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(Protobuf3ParserOPTION-6))|(1<<(Protobuf3ParserRPC-6))|(1<<(Protobuf3ParserSEMI-6)))) != 0 {
		{
			p.SetState(343)
			p.ServiceElement()
		}

		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(349)
		p.Match(Protobuf3ParserRC)
	}

	return localctx
}

// IServiceElementContext is an interface to support dynamic dispatch.
type IServiceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceElementContext differentiates from other interfaces.
	IsServiceElementContext()
}

type ServiceElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceElementContext() *ServiceElementContext {
	var p = new(ServiceElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceElement
	return p
}

func (*ServiceElementContext) IsServiceElementContext() {}

func NewServiceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceElementContext {
	var p = new(ServiceElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceElement

	return p
}

func (s *ServiceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceElementContext) OptionStatement() IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *ServiceElementContext) Rpc() IRpcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcContext)
}

func (s *ServiceElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ServiceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterServiceElement(s)
	}
}

func (s *ServiceElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitServiceElement(s)
	}
}

func (p *Protobuf3Parser) ServiceElement() (localctx IServiceElementContext) {
	this := p
	_ = this

	localctx = NewServiceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Protobuf3ParserRULE_serviceElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(354)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserOPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.OptionStatement()
		}

	case Protobuf3ParserRPC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(352)
			p.Rpc()
		}

	case Protobuf3ParserSEMI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(353)
			p.EmptyStatement_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpc
	return p
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRPC, 0)
}

func (s *RpcContext) RpcName() IRpcNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpcNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpcNameContext)
}

func (s *RpcContext) AllLP() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserLP)
}

func (s *RpcContext) LP(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLP, i)
}

func (s *RpcContext) AllMessageType() []IMessageTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem())
	var tst = make([]IMessageTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageTypeContext)
		}
	}

	return tst
}

func (s *RpcContext) MessageType(i int) IMessageTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *RpcContext) AllRP() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserRP)
}

func (s *RpcContext) RP(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRP, i)
}

func (s *RpcContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRETURNS, 0)
}

func (s *RpcContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *RpcContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *RpcContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *RpcContext) AllSTREAM() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserSTREAM)
}

func (s *RpcContext) STREAM(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTREAM, i)
}

func (s *RpcContext) AllOptionStatement() []IOptionStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem())
	var tst = make([]IOptionStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionStatementContext)
		}
	}

	return tst
}

func (s *RpcContext) OptionStatement(i int) IOptionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *RpcContext) AllEmptyStatement_() []IEmptyStatement_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem())
	var tst = make([]IEmptyStatement_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEmptyStatement_Context)
		}
	}

	return tst
}

func (s *RpcContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatement_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRpc(s)
	}
}

func (p *Protobuf3Parser) Rpc() (localctx IRpcContext) {
	this := p
	_ = this

	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Protobuf3ParserRULE_rpc)
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
		p.SetState(356)
		p.Match(Protobuf3ParserRPC)
	}
	{
		p.SetState(357)
		p.RpcName()
	}
	{
		p.SetState(358)
		p.Match(Protobuf3ParserLP)
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(359)
			p.Match(Protobuf3ParserSTREAM)
		}

	}
	{
		p.SetState(362)
		p.MessageType()
	}
	{
		p.SetState(363)
		p.Match(Protobuf3ParserRP)
	}
	{
		p.SetState(364)
		p.Match(Protobuf3ParserRETURNS)
	}
	{
		p.SetState(365)
		p.Match(Protobuf3ParserLP)
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(366)
			p.Match(Protobuf3ParserSTREAM)
		}

	}
	{
		p.SetState(369)
		p.MessageType()
	}
	{
		p.SetState(370)
		p.Match(Protobuf3ParserRP)
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserLC:
		{
			p.SetState(371)
			p.Match(Protobuf3ParserLC)
		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Protobuf3ParserOPTION || _la == Protobuf3ParserSEMI {
			p.SetState(374)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case Protobuf3ParserOPTION:
				{
					p.SetState(372)
					p.OptionStatement()
				}

			case Protobuf3ParserSEMI:
				{
					p.SetState(373)
					p.EmptyStatement_()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(378)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(379)
			p.Match(Protobuf3ParserRC)
		}

	case Protobuf3ParserSEMI:
		{
			p.SetState(380)
			p.Match(Protobuf3ParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) IntLit() IIntLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMINUS, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPLUS, 0)
}

func (s *ConstantContext) FloatLit() IFloatLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatLitContext)
}

func (s *ConstantContext) StrLit() IStrLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ConstantContext) BoolLit() IBoolLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolLitContext)
}

func (s *ConstantContext) BlockLit() IBlockLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockLitContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *Protobuf3Parser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Protobuf3ParserRULE_constant)
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

	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS {
			{
				p.SetState(384)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(387)
			p.IntLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS {
			{
				p.SetState(388)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(391)
			p.FloatLit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(392)
			p.StrLit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(393)
			p.BoolLit()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(394)
			p.BlockLit()
		}

	}

	return localctx
}

// IBlockLitContext is an interface to support dynamic dispatch.
type IBlockLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockLitContext differentiates from other interfaces.
	IsBlockLitContext()
}

type BlockLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockLitContext() *BlockLitContext {
	var p = new(BlockLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_blockLit
	return p
}

func (*BlockLitContext) IsBlockLitContext() {}

func NewBlockLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockLitContext {
	var p = new(BlockLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_blockLit

	return p
}

func (s *BlockLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockLitContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *BlockLitContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *BlockLitContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *BlockLitContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *BlockLitContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOLON)
}

func (s *BlockLitContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOLON, i)
}

func (s *BlockLitContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *BlockLitContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *BlockLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterBlockLit(s)
	}
}

func (s *BlockLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitBlockLit(s)
	}
}

func (p *Protobuf3Parser) BlockLit() (localctx IBlockLitContext) {
	this := p
	_ = this

	localctx = NewBlockLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, Protobuf3ParserRULE_blockLit)
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
		p.SetState(397)
		p.Match(Protobuf3ParserLC)
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserSYNTAX)|(1<<Protobuf3ParserIMPORT)|(1<<Protobuf3ParserWEAK)|(1<<Protobuf3ParserPUBLIC)|(1<<Protobuf3ParserPACKAGE)|(1<<Protobuf3ParserOPTION)|(1<<Protobuf3ParserREPEATED)|(1<<Protobuf3ParserONEOF)|(1<<Protobuf3ParserMAP)|(1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserUINT32)|(1<<Protobuf3ParserUINT64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserSTRING)|(1<<Protobuf3ParserDOUBLE)|(1<<Protobuf3ParserFLOAT)|(1<<Protobuf3ParserBYTES)|(1<<Protobuf3ParserRESERVED)|(1<<Protobuf3ParserTO)|(1<<Protobuf3ParserMAX)|(1<<Protobuf3ParserENUM)|(1<<Protobuf3ParserMESSAGE)|(1<<Protobuf3ParserSERVICE)|(1<<Protobuf3ParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(Protobuf3ParserSTREAM-32))|(1<<(Protobuf3ParserRETURNS-32))|(1<<(Protobuf3ParserBOOL_LIT-32))|(1<<(Protobuf3ParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(398)
			p.Ident()
		}
		{
			p.SetState(399)
			p.Match(Protobuf3ParserCOLON)
		}
		{
			p.SetState(400)
			p.Constant()
		}

		p.SetState(406)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(407)
		p.Match(Protobuf3ParserRC)
	}

	return localctx
}

// IEmptyStatement_Context is an interface to support dynamic dispatch.
type IEmptyStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyStatement_Context differentiates from other interfaces.
	IsEmptyStatement_Context()
}

type EmptyStatement_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatement_Context() *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement_
	return p
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEmptyStatement_(s)
	}
}

func (p *Protobuf3Parser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	this := p
	_ = this

	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, Protobuf3ParserRULE_emptyStatement_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(409)
		p.Match(Protobuf3ParserSEMI)
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENTIFIER, 0)
}

func (s *IdentContext) Keywords() IKeywordsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordsContext)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitIdent(s)
	}
}

func (p *Protobuf3Parser) Ident() (localctx IIdentContext) {
	this := p
	_ = this

	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Protobuf3ParserRULE_ident)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(413)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.Match(Protobuf3ParserIDENTIFIER)
		}

	case Protobuf3ParserSYNTAX, Protobuf3ParserIMPORT, Protobuf3ParserWEAK, Protobuf3ParserPUBLIC, Protobuf3ParserPACKAGE, Protobuf3ParserOPTION, Protobuf3ParserREPEATED, Protobuf3ParserONEOF, Protobuf3ParserMAP, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserBOOL, Protobuf3ParserSTRING, Protobuf3ParserDOUBLE, Protobuf3ParserFLOAT, Protobuf3ParserBYTES, Protobuf3ParserRESERVED, Protobuf3ParserTO, Protobuf3ParserMAX, Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE, Protobuf3ParserRPC, Protobuf3ParserSTREAM, Protobuf3ParserRETURNS, Protobuf3ParserBOOL_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(412)
			p.Keywords()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fullIdent
	return p
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *FullIdentContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (p *Protobuf3Parser) FullIdent() (localctx IFullIdentContext) {
	this := p
	_ = this

	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, Protobuf3ParserRULE_fullIdent)
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
		p.SetState(415)
		p.Ident()
	}
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserDOT {
		{
			p.SetState(416)
			p.Match(Protobuf3ParserDOT)
		}
		{
			p.SetState(417)
			p.Ident()
		}

		p.SetState(422)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMessageNameContext is an interface to support dynamic dispatch.
type IMessageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageNameContext differentiates from other interfaces.
	IsMessageNameContext()
}

type MessageNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageNameContext() *MessageNameContext {
	var p = new(MessageNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageName
	return p
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageName

	return p
}

func (s *MessageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageName(s)
	}
}

func (s *MessageNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageName(s)
	}
}

func (p *Protobuf3Parser) MessageName() (localctx IMessageNameContext) {
	this := p
	_ = this

	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, Protobuf3ParserRULE_messageName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Ident()
	}

	return localctx
}

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumName
	return p
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumName(s)
	}
}

func (s *EnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumName(s)
	}
}

func (p *Protobuf3Parser) EnumName() (localctx IEnumNameContext) {
	this := p
	_ = this

	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, Protobuf3ParserRULE_enumName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(425)
		p.Ident()
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *Protobuf3Parser) FieldName() (localctx IFieldNameContext) {
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, Protobuf3ParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(427)
		p.Ident()
	}

	return localctx
}

// IOneofNameContext is an interface to support dynamic dispatch.
type IOneofNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneofNameContext differentiates from other interfaces.
	IsOneofNameContext()
}

type OneofNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofNameContext() *OneofNameContext {
	var p = new(OneofNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofName
	return p
}

func (*OneofNameContext) IsOneofNameContext() {}

func NewOneofNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofNameContext {
	var p = new(OneofNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofName

	return p
}

func (s *OneofNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *OneofNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneofName(s)
	}
}

func (s *OneofNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneofName(s)
	}
}

func (p *Protobuf3Parser) OneofName() (localctx IOneofNameContext) {
	this := p
	_ = this

	localctx = NewOneofNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, Protobuf3ParserRULE_oneofName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(429)
		p.Ident()
	}

	return localctx
}

// IMapNameContext is an interface to support dynamic dispatch.
type IMapNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapNameContext differentiates from other interfaces.
	IsMapNameContext()
}

type MapNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapNameContext() *MapNameContext {
	var p = new(MapNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapName
	return p
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapName

	return p
}

func (s *MapNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MapNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MapNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMapName(s)
	}
}

func (s *MapNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMapName(s)
	}
}

func (p *Protobuf3Parser) MapName() (localctx IMapNameContext) {
	this := p
	_ = this

	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, Protobuf3ParserRULE_mapName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(431)
		p.Ident()
	}

	return localctx
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterServiceName(s)
	}
}

func (s *ServiceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitServiceName(s)
	}
}

func (p *Protobuf3Parser) ServiceName() (localctx IServiceNameContext) {
	this := p
	_ = this

	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, Protobuf3ParserRULE_serviceName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(433)
		p.Ident()
	}

	return localctx
}

// IRpcNameContext is an interface to support dynamic dispatch.
type IRpcNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpcNameContext differentiates from other interfaces.
	IsRpcNameContext()
}

type RpcNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcNameContext() *RpcNameContext {
	var p = new(RpcNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpcName
	return p
}

func (*RpcNameContext) IsRpcNameContext() {}

func NewRpcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcNameContext {
	var p = new(RpcNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpcName

	return p
}

func (s *RpcNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcNameContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *RpcNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRpcName(s)
	}
}

func (s *RpcNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRpcName(s)
	}
}

func (p *Protobuf3Parser) RpcName() (localctx IRpcNameContext) {
	this := p
	_ = this

	localctx = NewRpcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, Protobuf3ParserRULE_rpcName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Ident()
	}

	return localctx
}

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageType
	return p
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) MessageName() IMessageNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, i)
}

func (s *MessageTypeContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *MessageTypeContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (p *Protobuf3Parser) MessageType() (localctx IMessageTypeContext) {
	this := p
	_ = this

	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, Protobuf3ParserRULE_messageType)
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
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserDOT {
		{
			p.SetState(437)
			p.Match(Protobuf3ParserDOT)
		}

	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(440)
				p.Ident()
			}
			{
				p.SetState(441)
				p.Match(Protobuf3ParserDOT)
			}

		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}
	{
		p.SetState(448)
		p.MessageName()
	}

	return localctx
}

// IEnumTypeContext is an interface to support dynamic dispatch.
type IEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeContext differentiates from other interfaces.
	IsEnumTypeContext()
}

type EnumTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeContext() *EnumTypeContext {
	var p = new(EnumTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumType
	return p
}

func (*EnumTypeContext) IsEnumTypeContext() {}

func NewEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeContext {
	var p = new(EnumTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumType

	return p
}

func (s *EnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeContext) EnumName() IEnumNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserDOT)
}

func (s *EnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, i)
}

func (s *EnumTypeContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *EnumTypeContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumType(s)
	}
}

func (s *EnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumType(s)
	}
}

func (p *Protobuf3Parser) EnumType() (localctx IEnumTypeContext) {
	this := p
	_ = this

	localctx = NewEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, Protobuf3ParserRULE_enumType)
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
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserDOT {
		{
			p.SetState(450)
			p.Match(Protobuf3ParserDOT)
		}

	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(453)
				p.Ident()
			}
			{
				p.SetState(454)
				p.Match(Protobuf3ParserDOT)
			}

		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
	}
	{
		p.SetState(461)
		p.EnumName()
	}

	return localctx
}

// IIntLitContext is an interface to support dynamic dispatch.
type IIntLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntLitContext differentiates from other interfaces.
	IsIntLitContext()
}

type IntLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLitContext() *IntLitContext {
	var p = new(IntLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_intLit
	return p
}

func (*IntLitContext) IsIntLitContext() {}

func NewIntLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLitContext {
	var p = new(IntLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_intLit

	return p
}

func (s *IntLitContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT_LIT, 0)
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterIntLit(s)
	}
}

func (s *IntLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitIntLit(s)
	}
}

func (p *Protobuf3Parser) IntLit() (localctx IIntLitContext) {
	this := p
	_ = this

	localctx = NewIntLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, Protobuf3ParserRULE_intLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(463)
		p.Match(Protobuf3ParserINT_LIT)
	}

	return localctx
}

// IStrLitContext is an interface to support dynamic dispatch.
type IStrLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrLitContext differentiates from other interfaces.
	IsStrLitContext()
}

type StrLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrLitContext() *StrLitContext {
	var p = new(StrLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_strLit
	return p
}

func (*StrLitContext) IsStrLitContext() {}

func NewStrLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrLitContext {
	var p = new(StrLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_strLit

	return p
}

func (s *StrLitContext) GetParser() antlr.Parser { return s.parser }

func (s *StrLitContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTR_LIT, 0)
}

func (s *StrLitContext) PROTO3_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_SINGLE, 0)
}

func (s *StrLitContext) PROTO3_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_DOBULE, 0)
}

func (s *StrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterStrLit(s)
	}
}

func (s *StrLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitStrLit(s)
	}
}

func (p *Protobuf3Parser) StrLit() (localctx IStrLitContext) {
	this := p
	_ = this

	localctx = NewStrLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, Protobuf3ParserRULE_strLit)
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
		p.SetState(465)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(Protobuf3ParserPROTO3_LIT_SINGLE-34))|(1<<(Protobuf3ParserPROTO3_LIT_DOBULE-34))|(1<<(Protobuf3ParserSTR_LIT-34)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBoolLitContext is an interface to support dynamic dispatch.
type IBoolLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolLitContext differentiates from other interfaces.
	IsBoolLitContext()
}

type BoolLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLitContext() *BoolLitContext {
	var p = new(BoolLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_boolLit
	return p
}

func (*BoolLitContext) IsBoolLitContext() {}

func NewBoolLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLitContext {
	var p = new(BoolLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_boolLit

	return p
}

func (s *BoolLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLitContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL_LIT, 0)
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterBoolLit(s)
	}
}

func (s *BoolLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitBoolLit(s)
	}
}

func (p *Protobuf3Parser) BoolLit() (localctx IBoolLitContext) {
	this := p
	_ = this

	localctx = NewBoolLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, Protobuf3ParserRULE_boolLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(467)
		p.Match(Protobuf3ParserBOOL_LIT)
	}

	return localctx
}

// IFloatLitContext is an interface to support dynamic dispatch.
type IFloatLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLitContext differentiates from other interfaces.
	IsFloatLitContext()
}

type FloatLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLitContext() *FloatLitContext {
	var p = new(FloatLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_floatLit
	return p
}

func (*FloatLitContext) IsFloatLitContext() {}

func NewFloatLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLitContext {
	var p = new(FloatLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_floatLit

	return p
}

func (s *FloatLitContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT_LIT, 0)
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFloatLit(s)
	}
}

func (s *FloatLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFloatLit(s)
	}
}

func (p *Protobuf3Parser) FloatLit() (localctx IFloatLitContext) {
	this := p
	_ = this

	localctx = NewFloatLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, Protobuf3ParserRULE_floatLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(469)
		p.Match(Protobuf3ParserFLOAT_LIT)
	}

	return localctx
}

// IKeywordsContext is an interface to support dynamic dispatch.
type IKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordsContext differentiates from other interfaces.
	IsKeywordsContext()
}

type KeywordsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordsContext() *KeywordsContext {
	var p = new(KeywordsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keywords
	return p
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSYNTAX, 0)
}

func (s *KeywordsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIMPORT, 0)
}

func (s *KeywordsContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserWEAK, 0)
}

func (s *KeywordsContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPUBLIC, 0)
}

func (s *KeywordsContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPACKAGE, 0)
}

func (s *KeywordsContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTION, 0)
}

func (s *KeywordsContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserREPEATED, 0)
}

func (s *KeywordsContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserONEOF, 0)
}

func (s *KeywordsContext) MAP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAP, 0)
}

func (s *KeywordsContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *KeywordsContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *KeywordsContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *KeywordsContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *KeywordsContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *KeywordsContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *KeywordsContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *KeywordsContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *KeywordsContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *KeywordsContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *KeywordsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *KeywordsContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *KeywordsContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOUBLE, 0)
}

func (s *KeywordsContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT, 0)
}

func (s *KeywordsContext) BYTES() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBYTES, 0)
}

func (s *KeywordsContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRESERVED, 0)
}

func (s *KeywordsContext) TO() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserTO, 0)
}

func (s *KeywordsContext) MAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAX, 0)
}

func (s *KeywordsContext) ENUM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserENUM, 0)
}

func (s *KeywordsContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMESSAGE, 0)
}

func (s *KeywordsContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSERVICE, 0)
}

func (s *KeywordsContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRPC, 0)
}

func (s *KeywordsContext) STREAM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTREAM, 0)
}

func (s *KeywordsContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRETURNS, 0)
}

func (s *KeywordsContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL_LIT, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterKeywords(s)
	}
}

func (s *KeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitKeywords(s)
	}
}

func (p *Protobuf3Parser) Keywords() (localctx IKeywordsContext) {
	this := p
	_ = this

	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, Protobuf3ParserRULE_keywords)
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
		p.SetState(471)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Protobuf3ParserSYNTAX)|(1<<Protobuf3ParserIMPORT)|(1<<Protobuf3ParserWEAK)|(1<<Protobuf3ParserPUBLIC)|(1<<Protobuf3ParserPACKAGE)|(1<<Protobuf3ParserOPTION)|(1<<Protobuf3ParserREPEATED)|(1<<Protobuf3ParserONEOF)|(1<<Protobuf3ParserMAP)|(1<<Protobuf3ParserINT32)|(1<<Protobuf3ParserINT64)|(1<<Protobuf3ParserUINT32)|(1<<Protobuf3ParserUINT64)|(1<<Protobuf3ParserSINT32)|(1<<Protobuf3ParserSINT64)|(1<<Protobuf3ParserFIXED32)|(1<<Protobuf3ParserFIXED64)|(1<<Protobuf3ParserSFIXED32)|(1<<Protobuf3ParserSFIXED64)|(1<<Protobuf3ParserBOOL)|(1<<Protobuf3ParserSTRING)|(1<<Protobuf3ParserDOUBLE)|(1<<Protobuf3ParserFLOAT)|(1<<Protobuf3ParserBYTES)|(1<<Protobuf3ParserRESERVED)|(1<<Protobuf3ParserTO)|(1<<Protobuf3ParserMAX)|(1<<Protobuf3ParserENUM)|(1<<Protobuf3ParserMESSAGE)|(1<<Protobuf3ParserSERVICE)|(1<<Protobuf3ParserRPC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(Protobuf3ParserSTREAM-32))|(1<<(Protobuf3ParserRETURNS-32))|(1<<(Protobuf3ParserBOOL_LIT-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
