// Code generated from matlab.g4 by ANTLR 4.7.2. DO NOT EDIT.

package matlab // matlab
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 430,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 84, 10, 2, 3, 3, 3, 3, 3, 3, 5, 3, 89,
	10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 95, 10, 3, 12, 3, 14, 3, 98, 11, 3,
	3, 4, 3, 4, 5, 4, 102, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	110, 10, 5, 12, 5, 14, 5, 113, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 5, 7, 124, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7,
	9, 155, 10, 9, 12, 9, 14, 9, 158, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 169, 10, 10, 12, 10, 14, 10, 172,
	11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 189, 10, 11, 12, 11, 14,
	11, 192, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 7, 12, 203, 10, 12, 12, 12, 14, 12, 206, 11, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 7, 13, 214, 10, 13, 12, 13, 14, 13, 217, 11, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 225, 10, 14, 12, 14, 14,
	14, 228, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 236,
	10, 15, 12, 15, 14, 15, 239, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 254, 10,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 261, 10, 19, 12, 19, 14,
	19, 264, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 271, 10, 20,
	12, 20, 14, 20, 274, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 288, 10, 23, 3, 24, 3,
	24, 3, 24, 3, 25, 3, 25, 5, 25, 295, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 7, 26, 302, 10, 26, 12, 26, 14, 26, 305, 11, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 337, 10, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 349,
	10, 28, 12, 28, 14, 28, 352, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 378,
	10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 384, 10, 30, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 5, 31, 392, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 7, 32, 400, 10, 32, 12, 32, 14, 32, 403, 11, 32, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 5, 33, 410, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 421, 10, 34, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 5, 35, 428, 10, 35, 3, 35, 2, 16, 4, 8, 16, 18, 20, 22,
	24, 26, 28, 36, 38, 50, 54, 62, 36, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 66, 68, 2, 4, 3, 2, 9, 11, 5, 2, 8, 8, 21, 21, 46, 46,
	2, 449, 2, 83, 3, 2, 2, 2, 4, 88, 3, 2, 2, 2, 6, 101, 3, 2, 2, 2, 8, 103,
	3, 2, 2, 2, 10, 114, 3, 2, 2, 2, 12, 123, 3, 2, 2, 2, 14, 125, 3, 2, 2,
	2, 16, 127, 3, 2, 2, 2, 18, 159, 3, 2, 2, 2, 20, 173, 3, 2, 2, 2, 22, 193,
	3, 2, 2, 2, 24, 207, 3, 2, 2, 2, 26, 218, 3, 2, 2, 2, 28, 229, 3, 2, 2,
	2, 30, 240, 3, 2, 2, 2, 32, 244, 3, 2, 2, 2, 34, 253, 3, 2, 2, 2, 36, 255,
	3, 2, 2, 2, 38, 265, 3, 2, 2, 2, 40, 275, 3, 2, 2, 2, 42, 279, 3, 2, 2,
	2, 44, 287, 3, 2, 2, 2, 46, 289, 3, 2, 2, 2, 48, 294, 3, 2, 2, 2, 50, 296,
	3, 2, 2, 2, 52, 336, 3, 2, 2, 2, 54, 338, 3, 2, 2, 2, 56, 377, 3, 2, 2,
	2, 58, 383, 3, 2, 2, 2, 60, 391, 3, 2, 2, 2, 62, 393, 3, 2, 2, 2, 64, 409,
	3, 2, 2, 2, 66, 420, 3, 2, 2, 2, 68, 427, 3, 2, 2, 2, 70, 84, 7, 44, 2,
	2, 71, 84, 7, 45, 2, 2, 72, 84, 7, 43, 2, 2, 73, 74, 7, 3, 2, 2, 74, 75,
	5, 28, 15, 2, 75, 76, 7, 4, 2, 2, 76, 84, 3, 2, 2, 2, 77, 78, 7, 5, 2,
	2, 78, 84, 7, 6, 2, 2, 79, 80, 7, 5, 2, 2, 80, 81, 5, 50, 26, 2, 81, 82,
	7, 6, 2, 2, 82, 84, 3, 2, 2, 2, 83, 70, 3, 2, 2, 2, 83, 71, 3, 2, 2, 2,
	83, 72, 3, 2, 2, 2, 83, 73, 3, 2, 2, 2, 83, 77, 3, 2, 2, 2, 83, 79, 3,
	2, 2, 2, 84, 3, 3, 2, 2, 2, 85, 86, 8, 3, 1, 2, 86, 89, 5, 2, 2, 2, 87,
	89, 5, 10, 6, 2, 88, 85, 3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89, 96, 3, 2,
	2, 2, 90, 91, 12, 4, 2, 2, 91, 95, 7, 41, 2, 2, 92, 93, 12, 3, 2, 2, 93,
	95, 7, 42, 2, 2, 94, 90, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 98, 3, 2,
	2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 5, 3, 2, 2, 2, 98, 96,
	3, 2, 2, 2, 99, 102, 7, 7, 2, 2, 100, 102, 5, 28, 15, 2, 101, 99, 3, 2,
	2, 2, 101, 100, 3, 2, 2, 2, 102, 7, 3, 2, 2, 2, 103, 104, 8, 5, 1, 2, 104,
	105, 5, 6, 4, 2, 105, 111, 3, 2, 2, 2, 106, 107, 12, 3, 2, 2, 107, 108,
	7, 8, 2, 2, 108, 110, 5, 6, 4, 2, 109, 106, 3, 2, 2, 2, 110, 113, 3, 2,
	2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 9, 3, 2, 2, 2, 113,
	111, 3, 2, 2, 2, 114, 115, 7, 44, 2, 2, 115, 116, 7, 3, 2, 2, 116, 117,
	5, 8, 5, 2, 117, 118, 7, 4, 2, 2, 118, 11, 3, 2, 2, 2, 119, 124, 5, 4,
	3, 2, 120, 121, 5, 14, 8, 2, 121, 122, 5, 4, 3, 2, 122, 124, 3, 2, 2, 2,
	123, 119, 3, 2, 2, 2, 123, 120, 3, 2, 2, 2, 124, 13, 3, 2, 2, 2, 125, 126,
	9, 2, 2, 2, 126, 15, 3, 2, 2, 2, 127, 128, 8, 9, 1, 2, 128, 129, 5, 12,
	7, 2, 129, 156, 3, 2, 2, 2, 130, 131, 12, 10, 2, 2, 131, 132, 7, 12, 2,
	2, 132, 155, 5, 12, 7, 2, 133, 134, 12, 9, 2, 2, 134, 135, 7, 13, 2, 2,
	135, 155, 5, 12, 7, 2, 136, 137, 12, 8, 2, 2, 137, 138, 7, 14, 2, 2, 138,
	155, 5, 12, 7, 2, 139, 140, 12, 7, 2, 2, 140, 141, 7, 15, 2, 2, 141, 155,
	5, 12, 7, 2, 142, 143, 12, 6, 2, 2, 143, 144, 7, 22, 2, 2, 144, 155, 5,
	12, 7, 2, 145, 146, 12, 5, 2, 2, 146, 147, 7, 23, 2, 2, 147, 155, 5, 12,
	7, 2, 148, 149, 12, 4, 2, 2, 149, 150, 7, 24, 2, 2, 150, 155, 5, 12, 7,
	2, 151, 152, 12, 3, 2, 2, 152, 153, 7, 25, 2, 2, 153, 155, 5, 12, 7, 2,
	154, 130, 3, 2, 2, 2, 154, 133, 3, 2, 2, 2, 154, 136, 3, 2, 2, 2, 154,
	139, 3, 2, 2, 2, 154, 142, 3, 2, 2, 2, 154, 145, 3, 2, 2, 2, 154, 148,
	3, 2, 2, 2, 154, 151, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2,
	2, 2, 156, 157, 3, 2, 2, 2, 157, 17, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2,
	159, 160, 8, 10, 1, 2, 160, 161, 5, 16, 9, 2, 161, 170, 3, 2, 2, 2, 162,
	163, 12, 4, 2, 2, 163, 164, 7, 9, 2, 2, 164, 169, 5, 16, 9, 2, 165, 166,
	12, 3, 2, 2, 166, 167, 7, 10, 2, 2, 167, 169, 5, 16, 9, 2, 168, 162, 3,
	2, 2, 2, 168, 165, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2,
	2, 170, 171, 3, 2, 2, 2, 171, 19, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173,
	174, 8, 11, 1, 2, 174, 175, 5, 18, 10, 2, 175, 190, 3, 2, 2, 2, 176, 177,
	12, 6, 2, 2, 177, 178, 7, 16, 2, 2, 178, 189, 5, 18, 10, 2, 179, 180, 12,
	5, 2, 2, 180, 181, 7, 17, 2, 2, 181, 189, 5, 18, 10, 2, 182, 183, 12, 4,
	2, 2, 183, 184, 7, 37, 2, 2, 184, 189, 5, 18, 10, 2, 185, 186, 12, 3, 2,
	2, 186, 187, 7, 38, 2, 2, 187, 189, 5, 18, 10, 2, 188, 176, 3, 2, 2, 2,
	188, 179, 3, 2, 2, 2, 188, 182, 3, 2, 2, 2, 188, 185, 3, 2, 2, 2, 189,
	192, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 21, 3,
	2, 2, 2, 192, 190, 3, 2, 2, 2, 193, 194, 8, 12, 1, 2, 194, 195, 5, 20,
	11, 2, 195, 204, 3, 2, 2, 2, 196, 197, 12, 4, 2, 2, 197, 198, 7, 39, 2,
	2, 198, 203, 5, 20, 11, 2, 199, 200, 12, 3, 2, 2, 200, 201, 7, 40, 2, 2,
	201, 203, 5, 20, 11, 2, 202, 196, 3, 2, 2, 2, 202, 199, 3, 2, 2, 2, 203,
	206, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 23, 3,
	2, 2, 2, 206, 204, 3, 2, 2, 2, 207, 208, 8, 13, 1, 2, 208, 209, 5, 22,
	12, 2, 209, 215, 3, 2, 2, 2, 210, 211, 12, 3, 2, 2, 211, 212, 7, 18, 2,
	2, 212, 214, 5, 22, 12, 2, 213, 210, 3, 2, 2, 2, 214, 217, 3, 2, 2, 2,
	215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 25, 3, 2, 2, 2, 217, 215,
	3, 2, 2, 2, 218, 219, 8, 14, 1, 2, 219, 220, 5, 24, 13, 2, 220, 226, 3,
	2, 2, 2, 221, 222, 12, 3, 2, 2, 222, 223, 7, 19, 2, 2, 223, 225, 5, 24,
	13, 2, 224, 221, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2,
	226, 227, 3, 2, 2, 2, 227, 27, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229, 230,
	8, 15, 1, 2, 230, 231, 5, 26, 14, 2, 231, 237, 3, 2, 2, 2, 232, 233, 12,
	3, 2, 2, 233, 234, 7, 7, 2, 2, 234, 236, 5, 26, 14, 2, 235, 232, 3, 2,
	2, 2, 236, 239, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2,
	238, 29, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 240, 241, 5, 4, 3, 2, 241, 242,
	7, 20, 2, 2, 242, 243, 5, 28, 15, 2, 243, 31, 3, 2, 2, 2, 244, 245, 9,
	3, 2, 2, 245, 33, 3, 2, 2, 2, 246, 254, 5, 40, 21, 2, 247, 254, 5, 42,
	22, 2, 248, 254, 5, 46, 24, 2, 249, 254, 5, 44, 23, 2, 250, 254, 5, 52,
	27, 2, 251, 254, 5, 56, 29, 2, 252, 254, 5, 58, 30, 2, 253, 246, 3, 2,
	2, 2, 253, 247, 3, 2, 2, 2, 253, 248, 3, 2, 2, 2, 253, 249, 3, 2, 2, 2,
	253, 250, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 252, 3, 2, 2, 2, 254,
	35, 3, 2, 2, 2, 255, 256, 8, 19, 1, 2, 256, 257, 5, 34, 18, 2, 257, 262,
	3, 2, 2, 2, 258, 259, 12, 3, 2, 2, 259, 261, 5, 34, 18, 2, 260, 258, 3,
	2, 2, 2, 261, 264, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 262, 263, 3, 2, 2,
	2, 263, 37, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 265, 266, 8, 20, 1, 2, 266,
	267, 7, 44, 2, 2, 267, 272, 3, 2, 2, 2, 268, 269, 12, 3, 2, 2, 269, 271,
	7, 44, 2, 2, 270, 268, 3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2,
	2, 2, 272, 273, 3, 2, 2, 2, 273, 39, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2,
	275, 276, 7, 32, 2, 2, 276, 277, 5, 38, 20, 2, 277, 278, 5, 32, 17, 2,
	278, 41, 3, 2, 2, 2, 279, 280, 7, 34, 2, 2, 280, 281, 5, 38, 20, 2, 281,
	282, 5, 32, 17, 2, 282, 43, 3, 2, 2, 2, 283, 288, 5, 32, 17, 2, 284, 285,
	5, 28, 15, 2, 285, 286, 5, 32, 17, 2, 286, 288, 3, 2, 2, 2, 287, 283, 3,
	2, 2, 2, 287, 284, 3, 2, 2, 2, 288, 45, 3, 2, 2, 2, 289, 290, 5, 30, 16,
	2, 290, 291, 5, 32, 17, 2, 291, 47, 3, 2, 2, 2, 292, 295, 5, 28, 15, 2,
	293, 295, 5, 44, 23, 2, 294, 292, 3, 2, 2, 2, 294, 293, 3, 2, 2, 2, 295,
	49, 3, 2, 2, 2, 296, 297, 8, 26, 1, 2, 297, 298, 5, 48, 25, 2, 298, 303,
	3, 2, 2, 2, 299, 300, 12, 3, 2, 2, 300, 302, 5, 48, 25, 2, 301, 299, 3,
	2, 2, 2, 302, 305, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 303, 304, 3, 2, 2,
	2, 304, 51, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 306, 307, 7, 33, 2, 2, 307,
	308, 5, 28, 15, 2, 308, 309, 5, 36, 19, 2, 309, 310, 7, 31, 2, 2, 310,
	311, 5, 32, 17, 2, 311, 337, 3, 2, 2, 2, 312, 313, 7, 33, 2, 2, 313, 314,
	5, 28, 15, 2, 314, 315, 5, 36, 19, 2, 315, 316, 7, 35, 2, 2, 316, 317,
	5, 36, 19, 2, 317, 318, 7, 31, 2, 2, 318, 319, 5, 32, 17, 2, 319, 337,
	3, 2, 2, 2, 320, 321, 7, 33, 2, 2, 321, 322, 5, 28, 15, 2, 322, 323, 5,
	36, 19, 2, 323, 324, 5, 54, 28, 2, 324, 325, 7, 31, 2, 2, 325, 326, 5,
	32, 17, 2, 326, 337, 3, 2, 2, 2, 327, 328, 7, 33, 2, 2, 328, 329, 5, 28,
	15, 2, 329, 330, 5, 36, 19, 2, 330, 331, 5, 54, 28, 2, 331, 332, 7, 35,
	2, 2, 332, 333, 5, 36, 19, 2, 333, 334, 7, 31, 2, 2, 334, 335, 5, 32, 17,
	2, 335, 337, 3, 2, 2, 2, 336, 306, 3, 2, 2, 2, 336, 312, 3, 2, 2, 2, 336,
	320, 3, 2, 2, 2, 336, 327, 3, 2, 2, 2, 337, 53, 3, 2, 2, 2, 338, 339, 8,
	28, 1, 2, 339, 340, 7, 36, 2, 2, 340, 341, 5, 28, 15, 2, 341, 342, 5, 36,
	19, 2, 342, 350, 3, 2, 2, 2, 343, 344, 12, 3, 2, 2, 344, 345, 7, 36, 2,
	2, 345, 346, 5, 28, 15, 2, 346, 347, 5, 36, 19, 2, 347, 349, 3, 2, 2, 2,
	348, 343, 3, 2, 2, 2, 349, 352, 3, 2, 2, 2, 350, 348, 3, 2, 2, 2, 350,
	351, 3, 2, 2, 2, 351, 55, 3, 2, 2, 2, 352, 350, 3, 2, 2, 2, 353, 354, 7,
	30, 2, 2, 354, 355, 5, 28, 15, 2, 355, 356, 5, 36, 19, 2, 356, 357, 7,
	31, 2, 2, 357, 358, 5, 32, 17, 2, 358, 378, 3, 2, 2, 2, 359, 360, 7, 29,
	2, 2, 360, 361, 7, 44, 2, 2, 361, 362, 7, 20, 2, 2, 362, 363, 5, 28, 15,
	2, 363, 364, 5, 36, 19, 2, 364, 365, 7, 31, 2, 2, 365, 366, 5, 32, 17,
	2, 366, 378, 3, 2, 2, 2, 367, 368, 7, 29, 2, 2, 368, 369, 7, 3, 2, 2, 369,
	370, 7, 44, 2, 2, 370, 371, 7, 20, 2, 2, 371, 372, 5, 28, 15, 2, 372, 373,
	7, 4, 2, 2, 373, 374, 5, 36, 19, 2, 374, 375, 7, 31, 2, 2, 375, 376, 5,
	32, 17, 2, 376, 378, 3, 2, 2, 2, 377, 353, 3, 2, 2, 2, 377, 359, 3, 2,
	2, 2, 377, 367, 3, 2, 2, 2, 378, 57, 3, 2, 2, 2, 379, 380, 7, 26, 2, 2,
	380, 384, 5, 32, 17, 2, 381, 382, 7, 27, 2, 2, 382, 384, 5, 32, 17, 2,
	383, 379, 3, 2, 2, 2, 383, 381, 3, 2, 2, 2, 384, 59, 3, 2, 2, 2, 385, 392,
	5, 36, 19, 2, 386, 387, 7, 28, 2, 2, 387, 388, 5, 68, 35, 2, 388, 389,
	5, 32, 17, 2, 389, 390, 5, 36, 19, 2, 390, 392, 3, 2, 2, 2, 391, 385, 3,
	2, 2, 2, 391, 386, 3, 2, 2, 2, 392, 61, 3, 2, 2, 2, 393, 394, 8, 32, 1,
	2, 394, 395, 7, 44, 2, 2, 395, 401, 3, 2, 2, 2, 396, 397, 12, 3, 2, 2,
	397, 398, 7, 8, 2, 2, 398, 400, 7, 44, 2, 2, 399, 396, 3, 2, 2, 2, 400,
	403, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2, 401, 402, 3, 2, 2, 2, 402, 63, 3,
	2, 2, 2, 403, 401, 3, 2, 2, 2, 404, 410, 7, 44, 2, 2, 405, 406, 7, 5, 2,
	2, 406, 407, 5, 62, 32, 2, 407, 408, 7, 6, 2, 2, 408, 410, 3, 2, 2, 2,
	409, 404, 3, 2, 2, 2, 409, 405, 3, 2, 2, 2, 410, 65, 3, 2, 2, 2, 411, 421,
	7, 44, 2, 2, 412, 413, 7, 44, 2, 2, 413, 414, 7, 3, 2, 2, 414, 421, 7,
	4, 2, 2, 415, 416, 7, 44, 2, 2, 416, 417, 7, 3, 2, 2, 417, 418, 5, 62,
	32, 2, 418, 419, 7, 4, 2, 2, 419, 421, 3, 2, 2, 2, 420, 411, 3, 2, 2, 2,
	420, 412, 3, 2, 2, 2, 420, 415, 3, 2, 2, 2, 421, 67, 3, 2, 2, 2, 422, 428,
	5, 66, 34, 2, 423, 424, 5, 64, 33, 2, 424, 425, 7, 20, 2, 2, 425, 426,
	5, 66, 34, 2, 426, 428, 3, 2, 2, 2, 427, 422, 3, 2, 2, 2, 427, 423, 3,
	2, 2, 2, 428, 69, 3, 2, 2, 2, 35, 83, 88, 94, 96, 101, 111, 123, 154, 156,
	168, 170, 188, 190, 202, 204, 215, 226, 237, 253, 262, 272, 287, 294, 303,
	336, 350, 377, 383, 391, 401, 409, 420, 427,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'['", "']'", "':'", "','", "'+'", "'-'", "'~'", "'*'",
	"'/'", "'\\'", "'^'", "'<'", "'>'", "'&'", "'|'", "'='", "';'", "'.*'",
	"'.\\'", "'./'", "'.^'", "'break'", "'return'", "'function'", "'for'",
	"'while'", "'end'", "'global'", "'if'", "'clear'", "'else'", "'elseif'",
	"'<='", "'>='", "'=='", "'~='", "'transpose'", "'.''",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "ARRAYMUL", "ARRAYDIV", "ARRAYRDIV", "ARRAYPOW", "BREAK", "RETURN",
	"FUNCTION", "FOR", "WHILE", "END", "GLOBAL", "IF", "CLEAR", "ELSE", "ELSEIF",
	"LE_OP", "GE_OP", "EQ_OP", "NE_OP", "TRANSPOSE", "NCTRANSPOSE", "STRING_LITERAL",
	"IDENTIFIER", "CONSTANT", "CR", "WS",
}

var ruleNames = []string{
	"primary_expression", "postfix_expression", "index_expression", "index_expression_list",
	"array_expression", "unary_expression", "unary_operator", "multiplicative_expression",
	"additive_expression", "relational_expression", "equality_expression",
	"and_expression", "or_expression", "expression", "assignment_expression",
	"eostmt", "statement", "statement_list", "identifier_list", "global_statement",
	"clear_statement", "expression_statement", "assignment_statement", "array_element",
	"array_list", "selection_statement", "elseif_clause", "iteration_statement",
	"jump_statement", "translation_unit", "func_ident_list", "func_return_list",
	"function_declare_lhs", "function_declare",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type matlabParser struct {
	*antlr.BaseParser
}

func NewmatlabParser(input antlr.TokenStream) *matlabParser {
	this := new(matlabParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "matlab.g4"

	return this
}

// matlabParser tokens.
const (
	matlabParserEOF            = antlr.TokenEOF
	matlabParserT__0           = 1
	matlabParserT__1           = 2
	matlabParserT__2           = 3
	matlabParserT__3           = 4
	matlabParserT__4           = 5
	matlabParserT__5           = 6
	matlabParserT__6           = 7
	matlabParserT__7           = 8
	matlabParserT__8           = 9
	matlabParserT__9           = 10
	matlabParserT__10          = 11
	matlabParserT__11          = 12
	matlabParserT__12          = 13
	matlabParserT__13          = 14
	matlabParserT__14          = 15
	matlabParserT__15          = 16
	matlabParserT__16          = 17
	matlabParserT__17          = 18
	matlabParserT__18          = 19
	matlabParserARRAYMUL       = 20
	matlabParserARRAYDIV       = 21
	matlabParserARRAYRDIV      = 22
	matlabParserARRAYPOW       = 23
	matlabParserBREAK          = 24
	matlabParserRETURN         = 25
	matlabParserFUNCTION       = 26
	matlabParserFOR            = 27
	matlabParserWHILE          = 28
	matlabParserEND            = 29
	matlabParserGLOBAL         = 30
	matlabParserIF             = 31
	matlabParserCLEAR          = 32
	matlabParserELSE           = 33
	matlabParserELSEIF         = 34
	matlabParserLE_OP          = 35
	matlabParserGE_OP          = 36
	matlabParserEQ_OP          = 37
	matlabParserNE_OP          = 38
	matlabParserTRANSPOSE      = 39
	matlabParserNCTRANSPOSE    = 40
	matlabParserSTRING_LITERAL = 41
	matlabParserIDENTIFIER     = 42
	matlabParserCONSTANT       = 43
	matlabParserCR             = 44
	matlabParserWS             = 45
)

// matlabParser rules.
const (
	matlabParserRULE_primary_expression        = 0
	matlabParserRULE_postfix_expression        = 1
	matlabParserRULE_index_expression          = 2
	matlabParserRULE_index_expression_list     = 3
	matlabParserRULE_array_expression          = 4
	matlabParserRULE_unary_expression          = 5
	matlabParserRULE_unary_operator            = 6
	matlabParserRULE_multiplicative_expression = 7
	matlabParserRULE_additive_expression       = 8
	matlabParserRULE_relational_expression     = 9
	matlabParserRULE_equality_expression       = 10
	matlabParserRULE_and_expression            = 11
	matlabParserRULE_or_expression             = 12
	matlabParserRULE_expression                = 13
	matlabParserRULE_assignment_expression     = 14
	matlabParserRULE_eostmt                    = 15
	matlabParserRULE_statement                 = 16
	matlabParserRULE_statement_list            = 17
	matlabParserRULE_identifier_list           = 18
	matlabParserRULE_global_statement          = 19
	matlabParserRULE_clear_statement           = 20
	matlabParserRULE_expression_statement      = 21
	matlabParserRULE_assignment_statement      = 22
	matlabParserRULE_array_element             = 23
	matlabParserRULE_array_list                = 24
	matlabParserRULE_selection_statement       = 25
	matlabParserRULE_elseif_clause             = 26
	matlabParserRULE_iteration_statement       = 27
	matlabParserRULE_jump_statement            = 28
	matlabParserRULE_translation_unit          = 29
	matlabParserRULE_func_ident_list           = 30
	matlabParserRULE_func_return_list          = 31
	matlabParserRULE_function_declare_lhs      = 32
	matlabParserRULE_function_declare          = 33
)

// IPrimary_expressionContext is an interface to support dynamic dispatch.
type IPrimary_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimary_expressionContext differentiates from other interfaces.
	IsPrimary_expressionContext()
}

type Primary_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_expressionContext() *Primary_expressionContext {
	var p = new(Primary_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_primary_expression
	return p
}

func (*Primary_expressionContext) IsPrimary_expressionContext() {}

func NewPrimary_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_expressionContext {
	var p = new(Primary_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_primary_expression

	return p
}

func (s *Primary_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_expressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(matlabParserIDENTIFIER, 0)
}

func (s *Primary_expressionContext) CONSTANT() antlr.TerminalNode {
	return s.GetToken(matlabParserCONSTANT, 0)
}

func (s *Primary_expressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(matlabParserSTRING_LITERAL, 0)
}

func (s *Primary_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Primary_expressionContext) Array_list() IArray_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_listContext)
}

func (s *Primary_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primary_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterPrimary_expression(s)
	}
}

func (s *Primary_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitPrimary_expression(s)
	}
}

func (p *matlabParser) Primary_expression() (localctx IPrimary_expressionContext) {
	localctx = NewPrimary_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, matlabParserRULE_primary_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(matlabParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Match(matlabParserCONSTANT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Match(matlabParserSTRING_LITERAL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(71)
			p.Match(matlabParserT__0)
		}
		{
			p.SetState(72)
			p.expression(0)
		}
		{
			p.SetState(73)
			p.Match(matlabParserT__1)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.Match(matlabParserT__2)
		}
		{
			p.SetState(76)
			p.Match(matlabParserT__3)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
			p.Match(matlabParserT__2)
		}
		{
			p.SetState(78)
			p.array_list(0)
		}
		{
			p.SetState(79)
			p.Match(matlabParserT__3)
		}

	}

	return localctx
}

// IPostfix_expressionContext is an interface to support dynamic dispatch.
type IPostfix_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostfix_expressionContext differentiates from other interfaces.
	IsPostfix_expressionContext()
}

type Postfix_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfix_expressionContext() *Postfix_expressionContext {
	var p = new(Postfix_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_postfix_expression
	return p
}

func (*Postfix_expressionContext) IsPostfix_expressionContext() {}

func NewPostfix_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Postfix_expressionContext {
	var p = new(Postfix_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_postfix_expression

	return p
}

func (s *Postfix_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Postfix_expressionContext) Primary_expression() IPrimary_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_expressionContext)
}

func (s *Postfix_expressionContext) Array_expression() IArray_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_expressionContext)
}

func (s *Postfix_expressionContext) Postfix_expression() IPostfix_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfix_expressionContext)
}

func (s *Postfix_expressionContext) TRANSPOSE() antlr.TerminalNode {
	return s.GetToken(matlabParserTRANSPOSE, 0)
}

func (s *Postfix_expressionContext) NCTRANSPOSE() antlr.TerminalNode {
	return s.GetToken(matlabParserNCTRANSPOSE, 0)
}

func (s *Postfix_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Postfix_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Postfix_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterPostfix_expression(s)
	}
}

func (s *Postfix_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitPostfix_expression(s)
	}
}

func (p *matlabParser) Postfix_expression() (localctx IPostfix_expressionContext) {
	return p.postfix_expression(0)
}

func (p *matlabParser) postfix_expression(_p int) (localctx IPostfix_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPostfix_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPostfix_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, matlabParserRULE_postfix_expression, _p)

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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(84)
			p.Primary_expression()
		}

	case 2:
		{
			p.SetState(85)
			p.Array_expression()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPostfix_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_postfix_expression)
				p.SetState(88)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(89)
					p.Match(matlabParserTRANSPOSE)
				}

			case 2:
				localctx = NewPostfix_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_postfix_expression)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(91)
					p.Match(matlabParserNCTRANSPOSE)
				}

			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IIndex_expressionContext is an interface to support dynamic dispatch.
type IIndex_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndex_expressionContext differentiates from other interfaces.
	IsIndex_expressionContext()
}

type Index_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_expressionContext() *Index_expressionContext {
	var p = new(Index_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_index_expression
	return p
}

func (*Index_expressionContext) IsIndex_expressionContext() {}

func NewIndex_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_expressionContext {
	var p = new(Index_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_index_expression

	return p
}

func (s *Index_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Index_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterIndex_expression(s)
	}
}

func (s *Index_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitIndex_expression(s)
	}
}

func (p *matlabParser) Index_expression() (localctx IIndex_expressionContext) {
	localctx = NewIndex_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, matlabParserRULE_index_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case matlabParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Match(matlabParserT__4)
		}

	case matlabParserT__0, matlabParserT__2, matlabParserT__6, matlabParserT__7, matlabParserT__8, matlabParserSTRING_LITERAL, matlabParserIDENTIFIER, matlabParserCONSTANT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIndex_expression_listContext is an interface to support dynamic dispatch.
type IIndex_expression_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndex_expression_listContext differentiates from other interfaces.
	IsIndex_expression_listContext()
}

type Index_expression_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_expression_listContext() *Index_expression_listContext {
	var p = new(Index_expression_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_index_expression_list
	return p
}

func (*Index_expression_listContext) IsIndex_expression_listContext() {}

func NewIndex_expression_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_expression_listContext {
	var p = new(Index_expression_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_index_expression_list

	return p
}

func (s *Index_expression_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_expression_listContext) Index_expression() IIndex_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndex_expressionContext)
}

func (s *Index_expression_listContext) Index_expression_list() IIndex_expression_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_expression_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndex_expression_listContext)
}

func (s *Index_expression_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_expression_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_expression_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterIndex_expression_list(s)
	}
}

func (s *Index_expression_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitIndex_expression_list(s)
	}
}

func (p *matlabParser) Index_expression_list() (localctx IIndex_expression_listContext) {
	return p.index_expression_list(0)
}

func (p *matlabParser) index_expression_list(_p int) (localctx IIndex_expression_listContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewIndex_expression_listContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IIndex_expression_listContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, matlabParserRULE_index_expression_list, _p)

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
		p.SetState(102)
		p.Index_expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIndex_expression_listContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_index_expression_list)
			p.SetState(104)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(105)
				p.Match(matlabParserT__5)
			}
			{
				p.SetState(106)
				p.Index_expression()
			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IArray_expressionContext is an interface to support dynamic dispatch.
type IArray_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_expressionContext differentiates from other interfaces.
	IsArray_expressionContext()
}

type Array_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_expressionContext() *Array_expressionContext {
	var p = new(Array_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_array_expression
	return p
}

func (*Array_expressionContext) IsArray_expressionContext() {}

func NewArray_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_expressionContext {
	var p = new(Array_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_array_expression

	return p
}

func (s *Array_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_expressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(matlabParserIDENTIFIER, 0)
}

func (s *Array_expressionContext) Index_expression_list() IIndex_expression_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_expression_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndex_expression_listContext)
}

func (s *Array_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterArray_expression(s)
	}
}

func (s *Array_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitArray_expression(s)
	}
}

func (p *matlabParser) Array_expression() (localctx IArray_expressionContext) {
	localctx = NewArray_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, matlabParserRULE_array_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(112)
		p.Match(matlabParserIDENTIFIER)
	}
	{
		p.SetState(113)
		p.Match(matlabParserT__0)
	}
	{
		p.SetState(114)
		p.index_expression_list(0)
	}
	{
		p.SetState(115)
		p.Match(matlabParserT__1)
	}

	return localctx
}

// IUnary_expressionContext is an interface to support dynamic dispatch.
type IUnary_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnary_expressionContext differentiates from other interfaces.
	IsUnary_expressionContext()
}

type Unary_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_expressionContext() *Unary_expressionContext {
	var p = new(Unary_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_unary_expression
	return p
}

func (*Unary_expressionContext) IsUnary_expressionContext() {}

func NewUnary_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_expressionContext {
	var p = new(Unary_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_unary_expression

	return p
}

func (s *Unary_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_expressionContext) Postfix_expression() IPostfix_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfix_expressionContext)
}

func (s *Unary_expressionContext) Unary_operator() IUnary_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnary_operatorContext)
}

func (s *Unary_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterUnary_expression(s)
	}
}

func (s *Unary_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitUnary_expression(s)
	}
}

func (p *matlabParser) Unary_expression() (localctx IUnary_expressionContext) {
	localctx = NewUnary_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, matlabParserRULE_unary_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case matlabParserT__0, matlabParserT__2, matlabParserSTRING_LITERAL, matlabParserIDENTIFIER, matlabParserCONSTANT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.postfix_expression(0)
		}

	case matlabParserT__6, matlabParserT__7, matlabParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Unary_operator()
		}
		{
			p.SetState(119)
			p.postfix_expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnary_operatorContext is an interface to support dynamic dispatch.
type IUnary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnary_operatorContext differentiates from other interfaces.
	IsUnary_operatorContext()
}

type Unary_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_operatorContext() *Unary_operatorContext {
	var p = new(Unary_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_unary_operator
	return p
}

func (*Unary_operatorContext) IsUnary_operatorContext() {}

func NewUnary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_operatorContext {
	var p = new(Unary_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_unary_operator

	return p
}

func (s *Unary_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Unary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterUnary_operator(s)
	}
}

func (s *Unary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitUnary_operator(s)
	}
}

func (p *matlabParser) Unary_operator() (localctx IUnary_operatorContext) {
	localctx = NewUnary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, matlabParserRULE_unary_operator)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<matlabParserT__6)|(1<<matlabParserT__7)|(1<<matlabParserT__8))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMultiplicative_expressionContext is an interface to support dynamic dispatch.
type IMultiplicative_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicative_expressionContext differentiates from other interfaces.
	IsMultiplicative_expressionContext()
}

type Multiplicative_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicative_expressionContext() *Multiplicative_expressionContext {
	var p = new(Multiplicative_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_multiplicative_expression
	return p
}

func (*Multiplicative_expressionContext) IsMultiplicative_expressionContext() {}

func NewMultiplicative_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Multiplicative_expressionContext {
	var p = new(Multiplicative_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_multiplicative_expression

	return p
}

func (s *Multiplicative_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Multiplicative_expressionContext) Unary_expression() IUnary_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnary_expressionContext)
}

func (s *Multiplicative_expressionContext) Multiplicative_expression() IMultiplicative_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicative_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicative_expressionContext)
}

func (s *Multiplicative_expressionContext) ARRAYMUL() antlr.TerminalNode {
	return s.GetToken(matlabParserARRAYMUL, 0)
}

func (s *Multiplicative_expressionContext) ARRAYDIV() antlr.TerminalNode {
	return s.GetToken(matlabParserARRAYDIV, 0)
}

func (s *Multiplicative_expressionContext) ARRAYRDIV() antlr.TerminalNode {
	return s.GetToken(matlabParserARRAYRDIV, 0)
}

func (s *Multiplicative_expressionContext) ARRAYPOW() antlr.TerminalNode {
	return s.GetToken(matlabParserARRAYPOW, 0)
}

func (s *Multiplicative_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Multiplicative_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Multiplicative_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterMultiplicative_expression(s)
	}
}

func (s *Multiplicative_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitMultiplicative_expression(s)
	}
}

func (p *matlabParser) Multiplicative_expression() (localctx IMultiplicative_expressionContext) {
	return p.multiplicative_expression(0)
}

func (p *matlabParser) multiplicative_expression(_p int) (localctx IMultiplicative_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultiplicative_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicative_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, matlabParserRULE_multiplicative_expression, _p)

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
		p.SetState(126)
		p.Unary_expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicative_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_multiplicative_expression)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(129)
					p.Match(matlabParserT__9)
				}
				{
					p.SetState(130)
					p.Unary_expression()
				}

			case 2:
				localctx = NewMultiplicative_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_multiplicative_expression)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(132)
					p.Match(matlabParserT__10)
				}
				{
					p.SetState(133)
					p.Unary_expression()
				}

			case 3:
				localctx = NewMultiplicative_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_multiplicative_expression)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(135)
					p.Match(matlabParserT__11)
				}
				{
					p.SetState(136)
					p.Unary_expression()
				}

			case 4:
				localctx = NewMultiplicative_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_multiplicative_expression)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(138)
					p.Match(matlabParserT__12)
				}
				{
					p.SetState(139)
					p.Unary_expression()
				}

			case 5:
				localctx = NewMultiplicative_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_multiplicative_expression)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(141)
					p.Match(matlabParserARRAYMUL)
				}
				{
					p.SetState(142)
					p.Unary_expression()
				}

			case 6:
				localctx = NewMultiplicative_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_multiplicative_expression)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(144)
					p.Match(matlabParserARRAYDIV)
				}
				{
					p.SetState(145)
					p.Unary_expression()
				}

			case 7:
				localctx = NewMultiplicative_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_multiplicative_expression)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(147)
					p.Match(matlabParserARRAYRDIV)
				}
				{
					p.SetState(148)
					p.Unary_expression()
				}

			case 8:
				localctx = NewMultiplicative_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_multiplicative_expression)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(150)
					p.Match(matlabParserARRAYPOW)
				}
				{
					p.SetState(151)
					p.Unary_expression()
				}

			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IAdditive_expressionContext is an interface to support dynamic dispatch.
type IAdditive_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditive_expressionContext differentiates from other interfaces.
	IsAdditive_expressionContext()
}

type Additive_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditive_expressionContext() *Additive_expressionContext {
	var p = new(Additive_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_additive_expression
	return p
}

func (*Additive_expressionContext) IsAdditive_expressionContext() {}

func NewAdditive_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Additive_expressionContext {
	var p = new(Additive_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_additive_expression

	return p
}

func (s *Additive_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Additive_expressionContext) Multiplicative_expression() IMultiplicative_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicative_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicative_expressionContext)
}

func (s *Additive_expressionContext) Additive_expression() IAdditive_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditive_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditive_expressionContext)
}

func (s *Additive_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Additive_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Additive_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterAdditive_expression(s)
	}
}

func (s *Additive_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitAdditive_expression(s)
	}
}

func (p *matlabParser) Additive_expression() (localctx IAdditive_expressionContext) {
	return p.additive_expression(0)
}

func (p *matlabParser) additive_expression(_p int) (localctx IAdditive_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAdditive_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditive_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, matlabParserRULE_additive_expression, _p)

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
		p.SetState(158)
		p.multiplicative_expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAdditive_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_additive_expression)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(161)
					p.Match(matlabParserT__6)
				}
				{
					p.SetState(162)
					p.multiplicative_expression(0)
				}

			case 2:
				localctx = NewAdditive_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_additive_expression)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(164)
					p.Match(matlabParserT__7)
				}
				{
					p.SetState(165)
					p.multiplicative_expression(0)
				}

			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IRelational_expressionContext is an interface to support dynamic dispatch.
type IRelational_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelational_expressionContext differentiates from other interfaces.
	IsRelational_expressionContext()
}

type Relational_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelational_expressionContext() *Relational_expressionContext {
	var p = new(Relational_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_relational_expression
	return p
}

func (*Relational_expressionContext) IsRelational_expressionContext() {}

func NewRelational_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Relational_expressionContext {
	var p = new(Relational_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_relational_expression

	return p
}

func (s *Relational_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Relational_expressionContext) Additive_expression() IAdditive_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditive_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditive_expressionContext)
}

func (s *Relational_expressionContext) Relational_expression() IRelational_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelational_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelational_expressionContext)
}

func (s *Relational_expressionContext) LE_OP() antlr.TerminalNode {
	return s.GetToken(matlabParserLE_OP, 0)
}

func (s *Relational_expressionContext) GE_OP() antlr.TerminalNode {
	return s.GetToken(matlabParserGE_OP, 0)
}

func (s *Relational_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Relational_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Relational_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterRelational_expression(s)
	}
}

func (s *Relational_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitRelational_expression(s)
	}
}

func (p *matlabParser) Relational_expression() (localctx IRelational_expressionContext) {
	return p.relational_expression(0)
}

func (p *matlabParser) relational_expression(_p int) (localctx IRelational_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelational_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelational_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, matlabParserRULE_relational_expression, _p)

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
		p.SetState(172)
		p.additive_expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewRelational_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_relational_expression)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(175)
					p.Match(matlabParserT__13)
				}
				{
					p.SetState(176)
					p.additive_expression(0)
				}

			case 2:
				localctx = NewRelational_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_relational_expression)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(178)
					p.Match(matlabParserT__14)
				}
				{
					p.SetState(179)
					p.additive_expression(0)
				}

			case 3:
				localctx = NewRelational_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_relational_expression)
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(181)
					p.Match(matlabParserLE_OP)
				}
				{
					p.SetState(182)
					p.additive_expression(0)
				}

			case 4:
				localctx = NewRelational_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_relational_expression)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(184)
					p.Match(matlabParserGE_OP)
				}
				{
					p.SetState(185)
					p.additive_expression(0)
				}

			}

		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IEquality_expressionContext is an interface to support dynamic dispatch.
type IEquality_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEquality_expressionContext differentiates from other interfaces.
	IsEquality_expressionContext()
}

type Equality_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquality_expressionContext() *Equality_expressionContext {
	var p = new(Equality_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_equality_expression
	return p
}

func (*Equality_expressionContext) IsEquality_expressionContext() {}

func NewEquality_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Equality_expressionContext {
	var p = new(Equality_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_equality_expression

	return p
}

func (s *Equality_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Equality_expressionContext) Relational_expression() IRelational_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelational_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelational_expressionContext)
}

func (s *Equality_expressionContext) Equality_expression() IEquality_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEquality_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEquality_expressionContext)
}

func (s *Equality_expressionContext) EQ_OP() antlr.TerminalNode {
	return s.GetToken(matlabParserEQ_OP, 0)
}

func (s *Equality_expressionContext) NE_OP() antlr.TerminalNode {
	return s.GetToken(matlabParserNE_OP, 0)
}

func (s *Equality_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Equality_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Equality_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterEquality_expression(s)
	}
}

func (s *Equality_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitEquality_expression(s)
	}
}

func (p *matlabParser) Equality_expression() (localctx IEquality_expressionContext) {
	return p.equality_expression(0)
}

func (p *matlabParser) equality_expression(_p int) (localctx IEquality_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEquality_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEquality_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, matlabParserRULE_equality_expression, _p)

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
		p.SetState(192)
		p.relational_expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEquality_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_equality_expression)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(195)
					p.Match(matlabParserEQ_OP)
				}
				{
					p.SetState(196)
					p.relational_expression(0)
				}

			case 2:
				localctx = NewEquality_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_equality_expression)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(198)
					p.Match(matlabParserNE_OP)
				}
				{
					p.SetState(199)
					p.relational_expression(0)
				}

			}

		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IAnd_expressionContext is an interface to support dynamic dispatch.
type IAnd_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_expressionContext differentiates from other interfaces.
	IsAnd_expressionContext()
}

type And_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_expressionContext() *And_expressionContext {
	var p = new(And_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_and_expression
	return p
}

func (*And_expressionContext) IsAnd_expressionContext() {}

func NewAnd_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_expressionContext {
	var p = new(And_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_and_expression

	return p
}

func (s *And_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *And_expressionContext) Equality_expression() IEquality_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEquality_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEquality_expressionContext)
}

func (s *And_expressionContext) And_expression() IAnd_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnd_expressionContext)
}

func (s *And_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterAnd_expression(s)
	}
}

func (s *And_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitAnd_expression(s)
	}
}

func (p *matlabParser) And_expression() (localctx IAnd_expressionContext) {
	return p.and_expression(0)
}

func (p *matlabParser) and_expression(_p int) (localctx IAnd_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAnd_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAnd_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, matlabParserRULE_and_expression, _p)

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
		p.SetState(206)
		p.equality_expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAnd_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_and_expression)
			p.SetState(208)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(209)
				p.Match(matlabParserT__15)
			}
			{
				p.SetState(210)
				p.equality_expression(0)
			}

		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IOr_expressionContext is an interface to support dynamic dispatch.
type IOr_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOr_expressionContext differentiates from other interfaces.
	IsOr_expressionContext()
}

type Or_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_expressionContext() *Or_expressionContext {
	var p = new(Or_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_or_expression
	return p
}

func (*Or_expressionContext) IsOr_expressionContext() {}

func NewOr_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_expressionContext {
	var p = new(Or_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_or_expression

	return p
}

func (s *Or_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_expressionContext) And_expression() IAnd_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnd_expressionContext)
}

func (s *Or_expressionContext) Or_expression() IOr_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_expressionContext)
}

func (s *Or_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterOr_expression(s)
	}
}

func (s *Or_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitOr_expression(s)
	}
}

func (p *matlabParser) Or_expression() (localctx IOr_expressionContext) {
	return p.or_expression(0)
}

func (p *matlabParser) or_expression(_p int) (localctx IOr_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewOr_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOr_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, matlabParserRULE_or_expression, _p)

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
		p.SetState(217)
		p.and_expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOr_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_or_expression)
			p.SetState(219)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(220)
				p.Match(matlabParserT__16)
			}
			{
				p.SetState(221)
				p.and_expression(0)
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
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
	p.RuleIndex = matlabParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Or_expression() IOr_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_expressionContext)
}

func (s *ExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *matlabParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *matlabParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, matlabParserRULE_expression, _p)

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
		p.SetState(228)
		p.or_expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_expression)
			p.SetState(230)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(231)
				p.Match(matlabParserT__4)
			}
			{
				p.SetState(232)
				p.or_expression(0)
			}

		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignment_expressionContext is an interface to support dynamic dispatch.
type IAssignment_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignment_expressionContext differentiates from other interfaces.
	IsAssignment_expressionContext()
}

type Assignment_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_expressionContext() *Assignment_expressionContext {
	var p = new(Assignment_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_assignment_expression
	return p
}

func (*Assignment_expressionContext) IsAssignment_expressionContext() {}

func NewAssignment_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_expressionContext {
	var p = new(Assignment_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_assignment_expression

	return p
}

func (s *Assignment_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_expressionContext) Postfix_expression() IPostfix_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfix_expressionContext)
}

func (s *Assignment_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Assignment_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterAssignment_expression(s)
	}
}

func (s *Assignment_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitAssignment_expression(s)
	}
}

func (p *matlabParser) Assignment_expression() (localctx IAssignment_expressionContext) {
	localctx = NewAssignment_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, matlabParserRULE_assignment_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(238)
		p.postfix_expression(0)
	}
	{
		p.SetState(239)
		p.Match(matlabParserT__17)
	}
	{
		p.SetState(240)
		p.expression(0)
	}

	return localctx
}

// IEostmtContext is an interface to support dynamic dispatch.
type IEostmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEostmtContext differentiates from other interfaces.
	IsEostmtContext()
}

type EostmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEostmtContext() *EostmtContext {
	var p = new(EostmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_eostmt
	return p
}

func (*EostmtContext) IsEostmtContext() {}

func NewEostmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EostmtContext {
	var p = new(EostmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_eostmt

	return p
}

func (s *EostmtContext) GetParser() antlr.Parser { return s.parser }

func (s *EostmtContext) CR() antlr.TerminalNode {
	return s.GetToken(matlabParserCR, 0)
}

func (s *EostmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EostmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EostmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterEostmt(s)
	}
}

func (s *EostmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitEostmt(s)
	}
}

func (p *matlabParser) Eostmt() (localctx IEostmtContext) {
	localctx = NewEostmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, matlabParserRULE_eostmt)
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
		p.SetState(242)
		_la = p.GetTokenStream().LA(1)

		if !(_la == matlabParserT__5 || _la == matlabParserT__18 || _la == matlabParserCR) {
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
	p.RuleIndex = matlabParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Global_statement() IGlobal_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobal_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlobal_statementContext)
}

func (s *StatementContext) Clear_statement() IClear_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClear_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClear_statementContext)
}

func (s *StatementContext) Assignment_statement() IAssignment_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignment_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignment_statementContext)
}

func (s *StatementContext) Expression_statement() IExpression_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_statementContext)
}

func (s *StatementContext) Selection_statement() ISelection_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelection_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelection_statementContext)
}

func (s *StatementContext) Iteration_statement() IIteration_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIteration_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIteration_statementContext)
}

func (s *StatementContext) Jump_statement() IJump_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJump_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJump_statementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *matlabParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, matlabParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(244)
			p.Global_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(245)
			p.Clear_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(246)
			p.Assignment_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(247)
			p.Expression_statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(248)
			p.Selection_statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(249)
			p.Iteration_statement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(250)
			p.Jump_statement()
		}

	}

	return localctx
}

// IStatement_listContext is an interface to support dynamic dispatch.
type IStatement_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatement_listContext differentiates from other interfaces.
	IsStatement_listContext()
}

type Statement_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatement_listContext() *Statement_listContext {
	var p = new(Statement_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_statement_list
	return p
}

func (*Statement_listContext) IsStatement_listContext() {}

func NewStatement_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_listContext {
	var p = new(Statement_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_statement_list

	return p
}

func (s *Statement_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Statement_listContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_listContext) Statement_list() IStatement_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatement_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Statement_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statement_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterStatement_list(s)
	}
}

func (s *Statement_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitStatement_list(s)
	}
}

func (p *matlabParser) Statement_list() (localctx IStatement_listContext) {
	return p.statement_list(0)
}

func (p *matlabParser) statement_list(_p int) (localctx IStatement_listContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewStatement_listContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatement_listContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, matlabParserRULE_statement_list, _p)

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
		p.SetState(254)
		p.Statement()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewStatement_listContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_statement_list)
			p.SetState(256)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(257)
				p.Statement()
			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IIdentifier_listContext is an interface to support dynamic dispatch.
type IIdentifier_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifier_listContext differentiates from other interfaces.
	IsIdentifier_listContext()
}

type Identifier_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifier_listContext() *Identifier_listContext {
	var p = new(Identifier_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_identifier_list
	return p
}

func (*Identifier_listContext) IsIdentifier_listContext() {}

func NewIdentifier_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Identifier_listContext {
	var p = new(Identifier_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_identifier_list

	return p
}

func (s *Identifier_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Identifier_listContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(matlabParserIDENTIFIER, 0)
}

func (s *Identifier_listContext) Identifier_list() IIdentifier_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_listContext)
}

func (s *Identifier_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Identifier_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Identifier_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterIdentifier_list(s)
	}
}

func (s *Identifier_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitIdentifier_list(s)
	}
}

func (p *matlabParser) Identifier_list() (localctx IIdentifier_listContext) {
	return p.identifier_list(0)
}

func (p *matlabParser) identifier_list(_p int) (localctx IIdentifier_listContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewIdentifier_listContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IIdentifier_listContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, matlabParserRULE_identifier_list, _p)

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
		p.SetState(264)
		p.Match(matlabParserIDENTIFIER)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIdentifier_listContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_identifier_list)
			p.SetState(266)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(267)
				p.Match(matlabParserIDENTIFIER)
			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IGlobal_statementContext is an interface to support dynamic dispatch.
type IGlobal_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobal_statementContext differentiates from other interfaces.
	IsGlobal_statementContext()
}

type Global_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobal_statementContext() *Global_statementContext {
	var p = new(Global_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_global_statement
	return p
}

func (*Global_statementContext) IsGlobal_statementContext() {}

func NewGlobal_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Global_statementContext {
	var p = new(Global_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_global_statement

	return p
}

func (s *Global_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Global_statementContext) GLOBAL() antlr.TerminalNode {
	return s.GetToken(matlabParserGLOBAL, 0)
}

func (s *Global_statementContext) Identifier_list() IIdentifier_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_listContext)
}

func (s *Global_statementContext) Eostmt() IEostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEostmtContext)
}

func (s *Global_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Global_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Global_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterGlobal_statement(s)
	}
}

func (s *Global_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitGlobal_statement(s)
	}
}

func (p *matlabParser) Global_statement() (localctx IGlobal_statementContext) {
	localctx = NewGlobal_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, matlabParserRULE_global_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(matlabParserGLOBAL)
	}
	{
		p.SetState(274)
		p.identifier_list(0)
	}
	{
		p.SetState(275)
		p.Eostmt()
	}

	return localctx
}

// IClear_statementContext is an interface to support dynamic dispatch.
type IClear_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClear_statementContext differentiates from other interfaces.
	IsClear_statementContext()
}

type Clear_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClear_statementContext() *Clear_statementContext {
	var p = new(Clear_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_clear_statement
	return p
}

func (*Clear_statementContext) IsClear_statementContext() {}

func NewClear_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Clear_statementContext {
	var p = new(Clear_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_clear_statement

	return p
}

func (s *Clear_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Clear_statementContext) CLEAR() antlr.TerminalNode {
	return s.GetToken(matlabParserCLEAR, 0)
}

func (s *Clear_statementContext) Identifier_list() IIdentifier_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_listContext)
}

func (s *Clear_statementContext) Eostmt() IEostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEostmtContext)
}

func (s *Clear_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Clear_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Clear_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterClear_statement(s)
	}
}

func (s *Clear_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitClear_statement(s)
	}
}

func (p *matlabParser) Clear_statement() (localctx IClear_statementContext) {
	localctx = NewClear_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, matlabParserRULE_clear_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(matlabParserCLEAR)
	}
	{
		p.SetState(278)
		p.identifier_list(0)
	}
	{
		p.SetState(279)
		p.Eostmt()
	}

	return localctx
}

// IExpression_statementContext is an interface to support dynamic dispatch.
type IExpression_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_statementContext differentiates from other interfaces.
	IsExpression_statementContext()
}

type Expression_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_statementContext() *Expression_statementContext {
	var p = new(Expression_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_expression_statement
	return p
}

func (*Expression_statementContext) IsExpression_statementContext() {}

func NewExpression_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_statementContext {
	var p = new(Expression_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_expression_statement

	return p
}

func (s *Expression_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_statementContext) Eostmt() IEostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEostmtContext)
}

func (s *Expression_statementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expression_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterExpression_statement(s)
	}
}

func (s *Expression_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitExpression_statement(s)
	}
}

func (p *matlabParser) Expression_statement() (localctx IExpression_statementContext) {
	localctx = NewExpression_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, matlabParserRULE_expression_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(285)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case matlabParserT__5, matlabParserT__18, matlabParserCR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Eostmt()
		}

	case matlabParserT__0, matlabParserT__2, matlabParserT__6, matlabParserT__7, matlabParserT__8, matlabParserSTRING_LITERAL, matlabParserIDENTIFIER, matlabParserCONSTANT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.expression(0)
		}
		{
			p.SetState(283)
			p.Eostmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignment_statementContext is an interface to support dynamic dispatch.
type IAssignment_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignment_statementContext differentiates from other interfaces.
	IsAssignment_statementContext()
}

type Assignment_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_statementContext() *Assignment_statementContext {
	var p = new(Assignment_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_assignment_statement
	return p
}

func (*Assignment_statementContext) IsAssignment_statementContext() {}

func NewAssignment_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_statementContext {
	var p = new(Assignment_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_assignment_statement

	return p
}

func (s *Assignment_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_statementContext) Assignment_expression() IAssignment_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignment_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignment_expressionContext)
}

func (s *Assignment_statementContext) Eostmt() IEostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEostmtContext)
}

func (s *Assignment_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterAssignment_statement(s)
	}
}

func (s *Assignment_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitAssignment_statement(s)
	}
}

func (p *matlabParser) Assignment_statement() (localctx IAssignment_statementContext) {
	localctx = NewAssignment_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, matlabParserRULE_assignment_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(287)
		p.Assignment_expression()
	}
	{
		p.SetState(288)
		p.Eostmt()
	}

	return localctx
}

// IArray_elementContext is an interface to support dynamic dispatch.
type IArray_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_elementContext differentiates from other interfaces.
	IsArray_elementContext()
}

type Array_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_elementContext() *Array_elementContext {
	var p = new(Array_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_array_element
	return p
}

func (*Array_elementContext) IsArray_elementContext() {}

func NewArray_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_elementContext {
	var p = new(Array_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_array_element

	return p
}

func (s *Array_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_elementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Array_elementContext) Expression_statement() IExpression_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_statementContext)
}

func (s *Array_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterArray_element(s)
	}
}

func (s *Array_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitArray_element(s)
	}
}

func (p *matlabParser) Array_element() (localctx IArray_elementContext) {
	localctx = NewArray_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, matlabParserRULE_array_element)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(291)
			p.Expression_statement()
		}

	}

	return localctx
}

// IArray_listContext is an interface to support dynamic dispatch.
type IArray_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_listContext differentiates from other interfaces.
	IsArray_listContext()
}

type Array_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_listContext() *Array_listContext {
	var p = new(Array_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_array_list
	return p
}

func (*Array_listContext) IsArray_listContext() {}

func NewArray_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_listContext {
	var p = new(Array_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_array_list

	return p
}

func (s *Array_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_listContext) Array_element() IArray_elementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_elementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_elementContext)
}

func (s *Array_listContext) Array_list() IArray_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_listContext)
}

func (s *Array_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterArray_list(s)
	}
}

func (s *Array_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitArray_list(s)
	}
}

func (p *matlabParser) Array_list() (localctx IArray_listContext) {
	return p.array_list(0)
}

func (p *matlabParser) array_list(_p int) (localctx IArray_listContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArray_listContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArray_listContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, matlabParserRULE_array_list, _p)

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
		p.SetState(295)
		p.Array_element()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArray_listContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_array_list)
			p.SetState(297)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(298)
				p.Array_element()
			}

		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// ISelection_statementContext is an interface to support dynamic dispatch.
type ISelection_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelection_statementContext differentiates from other interfaces.
	IsSelection_statementContext()
}

type Selection_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelection_statementContext() *Selection_statementContext {
	var p = new(Selection_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_selection_statement
	return p
}

func (*Selection_statementContext) IsSelection_statementContext() {}

func NewSelection_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Selection_statementContext {
	var p = new(Selection_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_selection_statement

	return p
}

func (s *Selection_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Selection_statementContext) IF() antlr.TerminalNode {
	return s.GetToken(matlabParserIF, 0)
}

func (s *Selection_statementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Selection_statementContext) AllStatement_list() []IStatement_listContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatement_listContext)(nil)).Elem())
	var tst = make([]IStatement_listContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatement_listContext)
		}
	}

	return tst
}

func (s *Selection_statementContext) Statement_list(i int) IStatement_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatement_listContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Selection_statementContext) END() antlr.TerminalNode {
	return s.GetToken(matlabParserEND, 0)
}

func (s *Selection_statementContext) Eostmt() IEostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEostmtContext)
}

func (s *Selection_statementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(matlabParserELSE, 0)
}

func (s *Selection_statementContext) Elseif_clause() IElseif_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseif_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseif_clauseContext)
}

func (s *Selection_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Selection_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Selection_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterSelection_statement(s)
	}
}

func (s *Selection_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitSelection_statement(s)
	}
}

func (p *matlabParser) Selection_statement() (localctx ISelection_statementContext) {
	localctx = NewSelection_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, matlabParserRULE_selection_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.Match(matlabParserIF)
		}
		{
			p.SetState(305)
			p.expression(0)
		}
		{
			p.SetState(306)
			p.statement_list(0)
		}
		{
			p.SetState(307)
			p.Match(matlabParserEND)
		}
		{
			p.SetState(308)
			p.Eostmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.Match(matlabParserIF)
		}
		{
			p.SetState(311)
			p.expression(0)
		}
		{
			p.SetState(312)
			p.statement_list(0)
		}
		{
			p.SetState(313)
			p.Match(matlabParserELSE)
		}
		{
			p.SetState(314)
			p.statement_list(0)
		}
		{
			p.SetState(315)
			p.Match(matlabParserEND)
		}
		{
			p.SetState(316)
			p.Eostmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(318)
			p.Match(matlabParserIF)
		}
		{
			p.SetState(319)
			p.expression(0)
		}
		{
			p.SetState(320)
			p.statement_list(0)
		}
		{
			p.SetState(321)
			p.elseif_clause(0)
		}
		{
			p.SetState(322)
			p.Match(matlabParserEND)
		}
		{
			p.SetState(323)
			p.Eostmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(325)
			p.Match(matlabParserIF)
		}
		{
			p.SetState(326)
			p.expression(0)
		}
		{
			p.SetState(327)
			p.statement_list(0)
		}
		{
			p.SetState(328)
			p.elseif_clause(0)
		}
		{
			p.SetState(329)
			p.Match(matlabParserELSE)
		}
		{
			p.SetState(330)
			p.statement_list(0)
		}
		{
			p.SetState(331)
			p.Match(matlabParserEND)
		}
		{
			p.SetState(332)
			p.Eostmt()
		}

	}

	return localctx
}

// IElseif_clauseContext is an interface to support dynamic dispatch.
type IElseif_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseif_clauseContext differentiates from other interfaces.
	IsElseif_clauseContext()
}

type Elseif_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseif_clauseContext() *Elseif_clauseContext {
	var p = new(Elseif_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_elseif_clause
	return p
}

func (*Elseif_clauseContext) IsElseif_clauseContext() {}

func NewElseif_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elseif_clauseContext {
	var p = new(Elseif_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_elseif_clause

	return p
}

func (s *Elseif_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Elseif_clauseContext) ELSEIF() antlr.TerminalNode {
	return s.GetToken(matlabParserELSEIF, 0)
}

func (s *Elseif_clauseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Elseif_clauseContext) Statement_list() IStatement_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatement_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Elseif_clauseContext) Elseif_clause() IElseif_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseif_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseif_clauseContext)
}

func (s *Elseif_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elseif_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elseif_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterElseif_clause(s)
	}
}

func (s *Elseif_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitElseif_clause(s)
	}
}

func (p *matlabParser) Elseif_clause() (localctx IElseif_clauseContext) {
	return p.elseif_clause(0)
}

func (p *matlabParser) elseif_clause(_p int) (localctx IElseif_clauseContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewElseif_clauseContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IElseif_clauseContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, matlabParserRULE_elseif_clause, _p)

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
		p.SetState(337)
		p.Match(matlabParserELSEIF)
	}
	{
		p.SetState(338)
		p.expression(0)
	}
	{
		p.SetState(339)
		p.statement_list(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewElseif_clauseContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_elseif_clause)
			p.SetState(341)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(342)
				p.Match(matlabParserELSEIF)
			}
			{
				p.SetState(343)
				p.expression(0)
			}
			{
				p.SetState(344)
				p.statement_list(0)
			}

		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IIteration_statementContext is an interface to support dynamic dispatch.
type IIteration_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIteration_statementContext differentiates from other interfaces.
	IsIteration_statementContext()
}

type Iteration_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIteration_statementContext() *Iteration_statementContext {
	var p = new(Iteration_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_iteration_statement
	return p
}

func (*Iteration_statementContext) IsIteration_statementContext() {}

func NewIteration_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iteration_statementContext {
	var p = new(Iteration_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_iteration_statement

	return p
}

func (s *Iteration_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Iteration_statementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(matlabParserWHILE, 0)
}

func (s *Iteration_statementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Iteration_statementContext) Statement_list() IStatement_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatement_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Iteration_statementContext) END() antlr.TerminalNode {
	return s.GetToken(matlabParserEND, 0)
}

func (s *Iteration_statementContext) Eostmt() IEostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEostmtContext)
}

func (s *Iteration_statementContext) FOR() antlr.TerminalNode {
	return s.GetToken(matlabParserFOR, 0)
}

func (s *Iteration_statementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(matlabParserIDENTIFIER, 0)
}

func (s *Iteration_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iteration_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iteration_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterIteration_statement(s)
	}
}

func (s *Iteration_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitIteration_statement(s)
	}
}

func (p *matlabParser) Iteration_statement() (localctx IIteration_statementContext) {
	localctx = NewIteration_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, matlabParserRULE_iteration_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.Match(matlabParserWHILE)
		}
		{
			p.SetState(352)
			p.expression(0)
		}
		{
			p.SetState(353)
			p.statement_list(0)
		}
		{
			p.SetState(354)
			p.Match(matlabParserEND)
		}
		{
			p.SetState(355)
			p.Eostmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Match(matlabParserFOR)
		}
		{
			p.SetState(358)
			p.Match(matlabParserIDENTIFIER)
		}
		{
			p.SetState(359)
			p.Match(matlabParserT__17)
		}
		{
			p.SetState(360)
			p.expression(0)
		}
		{
			p.SetState(361)
			p.statement_list(0)
		}
		{
			p.SetState(362)
			p.Match(matlabParserEND)
		}
		{
			p.SetState(363)
			p.Eostmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(365)
			p.Match(matlabParserFOR)
		}
		{
			p.SetState(366)
			p.Match(matlabParserT__0)
		}
		{
			p.SetState(367)
			p.Match(matlabParserIDENTIFIER)
		}
		{
			p.SetState(368)
			p.Match(matlabParserT__17)
		}
		{
			p.SetState(369)
			p.expression(0)
		}
		{
			p.SetState(370)
			p.Match(matlabParserT__1)
		}
		{
			p.SetState(371)
			p.statement_list(0)
		}
		{
			p.SetState(372)
			p.Match(matlabParserEND)
		}
		{
			p.SetState(373)
			p.Eostmt()
		}

	}

	return localctx
}

// IJump_statementContext is an interface to support dynamic dispatch.
type IJump_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJump_statementContext differentiates from other interfaces.
	IsJump_statementContext()
}

type Jump_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJump_statementContext() *Jump_statementContext {
	var p = new(Jump_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_jump_statement
	return p
}

func (*Jump_statementContext) IsJump_statementContext() {}

func NewJump_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Jump_statementContext {
	var p = new(Jump_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_jump_statement

	return p
}

func (s *Jump_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Jump_statementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(matlabParserBREAK, 0)
}

func (s *Jump_statementContext) Eostmt() IEostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEostmtContext)
}

func (s *Jump_statementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(matlabParserRETURN, 0)
}

func (s *Jump_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Jump_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Jump_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterJump_statement(s)
	}
}

func (s *Jump_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitJump_statement(s)
	}
}

func (p *matlabParser) Jump_statement() (localctx IJump_statementContext) {
	localctx = NewJump_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, matlabParserRULE_jump_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(381)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case matlabParserBREAK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(377)
			p.Match(matlabParserBREAK)
		}
		{
			p.SetState(378)
			p.Eostmt()
		}

	case matlabParserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)
			p.Match(matlabParserRETURN)
		}
		{
			p.SetState(380)
			p.Eostmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITranslation_unitContext is an interface to support dynamic dispatch.
type ITranslation_unitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranslation_unitContext differentiates from other interfaces.
	IsTranslation_unitContext()
}

type Translation_unitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranslation_unitContext() *Translation_unitContext {
	var p = new(Translation_unitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_translation_unit
	return p
}

func (*Translation_unitContext) IsTranslation_unitContext() {}

func NewTranslation_unitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Translation_unitContext {
	var p = new(Translation_unitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_translation_unit

	return p
}

func (s *Translation_unitContext) GetParser() antlr.Parser { return s.parser }

func (s *Translation_unitContext) Statement_list() IStatement_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatement_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Translation_unitContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(matlabParserFUNCTION, 0)
}

func (s *Translation_unitContext) Function_declare() IFunction_declareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declareContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_declareContext)
}

func (s *Translation_unitContext) Eostmt() IEostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEostmtContext)
}

func (s *Translation_unitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Translation_unitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Translation_unitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterTranslation_unit(s)
	}
}

func (s *Translation_unitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitTranslation_unit(s)
	}
}

func (p *matlabParser) Translation_unit() (localctx ITranslation_unitContext) {
	localctx = NewTranslation_unitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, matlabParserRULE_translation_unit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case matlabParserT__0, matlabParserT__2, matlabParserT__5, matlabParserT__6, matlabParserT__7, matlabParserT__8, matlabParserT__18, matlabParserBREAK, matlabParserRETURN, matlabParserFOR, matlabParserWHILE, matlabParserGLOBAL, matlabParserIF, matlabParserCLEAR, matlabParserSTRING_LITERAL, matlabParserIDENTIFIER, matlabParserCONSTANT, matlabParserCR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.statement_list(0)
		}

	case matlabParserFUNCTION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)
			p.Match(matlabParserFUNCTION)
		}
		{
			p.SetState(385)
			p.Function_declare()
		}
		{
			p.SetState(386)
			p.Eostmt()
		}
		{
			p.SetState(387)
			p.statement_list(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunc_ident_listContext is an interface to support dynamic dispatch.
type IFunc_ident_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_ident_listContext differentiates from other interfaces.
	IsFunc_ident_listContext()
}

type Func_ident_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_ident_listContext() *Func_ident_listContext {
	var p = new(Func_ident_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_func_ident_list
	return p
}

func (*Func_ident_listContext) IsFunc_ident_listContext() {}

func NewFunc_ident_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_ident_listContext {
	var p = new(Func_ident_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_func_ident_list

	return p
}

func (s *Func_ident_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_ident_listContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(matlabParserIDENTIFIER, 0)
}

func (s *Func_ident_listContext) Func_ident_list() IFunc_ident_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_ident_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_ident_listContext)
}

func (s *Func_ident_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_ident_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_ident_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterFunc_ident_list(s)
	}
}

func (s *Func_ident_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitFunc_ident_list(s)
	}
}

func (p *matlabParser) Func_ident_list() (localctx IFunc_ident_listContext) {
	return p.func_ident_list(0)
}

func (p *matlabParser) func_ident_list(_p int) (localctx IFunc_ident_listContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFunc_ident_listContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFunc_ident_listContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, matlabParserRULE_func_ident_list, _p)

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
		p.SetState(392)
		p.Match(matlabParserIDENTIFIER)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFunc_ident_listContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, matlabParserRULE_func_ident_list)
			p.SetState(394)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(395)
				p.Match(matlabParserT__5)
			}
			{
				p.SetState(396)
				p.Match(matlabParserIDENTIFIER)
			}

		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// IFunc_return_listContext is an interface to support dynamic dispatch.
type IFunc_return_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_return_listContext differentiates from other interfaces.
	IsFunc_return_listContext()
}

type Func_return_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_return_listContext() *Func_return_listContext {
	var p = new(Func_return_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_func_return_list
	return p
}

func (*Func_return_listContext) IsFunc_return_listContext() {}

func NewFunc_return_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_return_listContext {
	var p = new(Func_return_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_func_return_list

	return p
}

func (s *Func_return_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_return_listContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(matlabParserIDENTIFIER, 0)
}

func (s *Func_return_listContext) Func_ident_list() IFunc_ident_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_ident_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_ident_listContext)
}

func (s *Func_return_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_return_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_return_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterFunc_return_list(s)
	}
}

func (s *Func_return_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitFunc_return_list(s)
	}
}

func (p *matlabParser) Func_return_list() (localctx IFunc_return_listContext) {
	localctx = NewFunc_return_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, matlabParserRULE_func_return_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(407)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case matlabParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(402)
			p.Match(matlabParserIDENTIFIER)
		}

	case matlabParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(403)
			p.Match(matlabParserT__2)
		}
		{
			p.SetState(404)
			p.func_ident_list(0)
		}
		{
			p.SetState(405)
			p.Match(matlabParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunction_declare_lhsContext is an interface to support dynamic dispatch.
type IFunction_declare_lhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_declare_lhsContext differentiates from other interfaces.
	IsFunction_declare_lhsContext()
}

type Function_declare_lhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declare_lhsContext() *Function_declare_lhsContext {
	var p = new(Function_declare_lhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_function_declare_lhs
	return p
}

func (*Function_declare_lhsContext) IsFunction_declare_lhsContext() {}

func NewFunction_declare_lhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declare_lhsContext {
	var p = new(Function_declare_lhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_function_declare_lhs

	return p
}

func (s *Function_declare_lhsContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declare_lhsContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(matlabParserIDENTIFIER, 0)
}

func (s *Function_declare_lhsContext) Func_ident_list() IFunc_ident_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_ident_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_ident_listContext)
}

func (s *Function_declare_lhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declare_lhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declare_lhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterFunction_declare_lhs(s)
	}
}

func (s *Function_declare_lhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitFunction_declare_lhs(s)
	}
}

func (p *matlabParser) Function_declare_lhs() (localctx IFunction_declare_lhsContext) {
	localctx = NewFunction_declare_lhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, matlabParserRULE_function_declare_lhs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.Match(matlabParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.Match(matlabParserIDENTIFIER)
		}
		{
			p.SetState(411)
			p.Match(matlabParserT__0)
		}
		{
			p.SetState(412)
			p.Match(matlabParserT__1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(413)
			p.Match(matlabParserIDENTIFIER)
		}
		{
			p.SetState(414)
			p.Match(matlabParserT__0)
		}
		{
			p.SetState(415)
			p.func_ident_list(0)
		}
		{
			p.SetState(416)
			p.Match(matlabParserT__1)
		}

	}

	return localctx
}

// IFunction_declareContext is an interface to support dynamic dispatch.
type IFunction_declareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_declareContext differentiates from other interfaces.
	IsFunction_declareContext()
}

type Function_declareContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declareContext() *Function_declareContext {
	var p = new(Function_declareContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = matlabParserRULE_function_declare
	return p
}

func (*Function_declareContext) IsFunction_declareContext() {}

func NewFunction_declareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declareContext {
	var p = new(Function_declareContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = matlabParserRULE_function_declare

	return p
}

func (s *Function_declareContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declareContext) Function_declare_lhs() IFunction_declare_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declare_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_declare_lhsContext)
}

func (s *Function_declareContext) Func_return_list() IFunc_return_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_return_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_return_listContext)
}

func (s *Function_declareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.EnterFunction_declare(s)
	}
}

func (s *Function_declareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(matlabListener); ok {
		listenerT.ExitFunction_declare(s)
	}
}

func (p *matlabParser) Function_declare() (localctx IFunction_declareContext) {
	localctx = NewFunction_declareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, matlabParserRULE_function_declare)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(420)
			p.Function_declare_lhs()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(421)
			p.Func_return_list()
		}
		{
			p.SetState(422)
			p.Match(matlabParserT__17)
		}
		{
			p.SetState(423)
			p.Function_declare_lhs()
		}

	}

	return localctx
}

func (p *matlabParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *Postfix_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Postfix_expressionContext)
		}
		return p.Postfix_expression_Sempred(t, predIndex)

	case 3:
		var t *Index_expression_listContext = nil
		if localctx != nil {
			t = localctx.(*Index_expression_listContext)
		}
		return p.Index_expression_list_Sempred(t, predIndex)

	case 7:
		var t *Multiplicative_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Multiplicative_expressionContext)
		}
		return p.Multiplicative_expression_Sempred(t, predIndex)

	case 8:
		var t *Additive_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Additive_expressionContext)
		}
		return p.Additive_expression_Sempred(t, predIndex)

	case 9:
		var t *Relational_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Relational_expressionContext)
		}
		return p.Relational_expression_Sempred(t, predIndex)

	case 10:
		var t *Equality_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Equality_expressionContext)
		}
		return p.Equality_expression_Sempred(t, predIndex)

	case 11:
		var t *And_expressionContext = nil
		if localctx != nil {
			t = localctx.(*And_expressionContext)
		}
		return p.And_expression_Sempred(t, predIndex)

	case 12:
		var t *Or_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Or_expressionContext)
		}
		return p.Or_expression_Sempred(t, predIndex)

	case 13:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 17:
		var t *Statement_listContext = nil
		if localctx != nil {
			t = localctx.(*Statement_listContext)
		}
		return p.Statement_list_Sempred(t, predIndex)

	case 18:
		var t *Identifier_listContext = nil
		if localctx != nil {
			t = localctx.(*Identifier_listContext)
		}
		return p.Identifier_list_Sempred(t, predIndex)

	case 24:
		var t *Array_listContext = nil
		if localctx != nil {
			t = localctx.(*Array_listContext)
		}
		return p.Array_list_Sempred(t, predIndex)

	case 26:
		var t *Elseif_clauseContext = nil
		if localctx != nil {
			t = localctx.(*Elseif_clauseContext)
		}
		return p.Elseif_clause_Sempred(t, predIndex)

	case 30:
		var t *Func_ident_listContext = nil
		if localctx != nil {
			t = localctx.(*Func_ident_listContext)
		}
		return p.Func_ident_list_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *matlabParser) Postfix_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Index_expression_list_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Multiplicative_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Additive_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Relational_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Equality_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 17:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) And_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 19:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Or_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 20:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 21:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Statement_list_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 22:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Identifier_list_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 23:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Array_list_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 24:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Elseif_clause_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 25:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *matlabParser) Func_ident_list_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 26:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
