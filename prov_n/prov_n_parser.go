// Code generated from PROV_N.g4 by ANTLR 4.9.3. DO NOT EDIT.

package prov_n // PROV_N
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 506,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 3, 2, 5, 2, 107, 10, 2, 3,
	2, 7, 2, 110, 10, 2, 12, 2, 14, 2, 113, 11, 2, 3, 2, 3, 2, 7, 2, 117, 10,
	2, 12, 2, 14, 2, 120, 11, 2, 5, 2, 122, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	5, 3, 128, 10, 3, 3, 3, 7, 3, 131, 10, 3, 12, 3, 14, 3, 134, 11, 3, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 5,
	7, 148, 10, 7, 3, 7, 7, 7, 151, 10, 7, 12, 7, 14, 7, 154, 11, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 178, 10, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 5, 11, 191, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 197, 10, 12,
	12, 12, 14, 12, 200, 11, 12, 5, 12, 202, 10, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 212, 10, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 5, 18, 222, 10, 18, 3, 18, 3, 18, 5,
	18, 226, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	5, 19, 236, 10, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 243, 10,
	20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 5, 22, 256, 10, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 5,
	23, 264, 10, 23, 3, 24, 3, 24, 5, 24, 268, 10, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 5, 26, 274, 10, 26, 3, 27, 3, 27, 5, 27, 278, 10, 27, 3, 28, 3,
	28, 3, 29, 3, 29, 5, 29, 284, 10, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 33, 3, 33, 5, 33, 294, 10, 33, 3, 34, 3, 34, 3, 35, 3, 35, 5,
	35, 300, 10, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 5, 36, 311, 10, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 327, 10, 37,
	3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 38, 3, 38, 5, 38, 343, 10, 38, 3, 38, 3, 38, 3, 38, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 357, 10,
	39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 386, 10, 42, 3, 42,
	3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 408,
	10, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 426, 10, 45, 3,
	45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 7, 50, 467, 10, 50, 12,
	50, 14, 50, 470, 11, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51,
	3, 51, 5, 51, 480, 10, 51, 3, 52, 3, 52, 3, 52, 3, 52, 7, 52, 486, 10,
	52, 12, 52, 14, 52, 489, 11, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3,
	52, 7, 52, 497, 10, 52, 12, 52, 14, 52, 500, 11, 52, 3, 52, 3, 52, 5, 52,
	504, 10, 52, 3, 52, 2, 2, 53, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96,
	98, 100, 102, 2, 3, 3, 2, 44, 45, 2, 511, 2, 104, 3, 2, 2, 2, 4, 127, 3,
	2, 2, 2, 6, 135, 3, 2, 2, 2, 8, 138, 3, 2, 2, 2, 10, 142, 3, 2, 2, 2, 12,
	144, 3, 2, 2, 2, 14, 157, 3, 2, 2, 2, 16, 177, 3, 2, 2, 2, 18, 179, 3,
	2, 2, 2, 20, 190, 3, 2, 2, 2, 22, 201, 3, 2, 2, 2, 24, 203, 3, 2, 2, 2,
	26, 207, 3, 2, 2, 2, 28, 211, 3, 2, 2, 2, 30, 213, 3, 2, 2, 2, 32, 217,
	3, 2, 2, 2, 34, 225, 3, 2, 2, 2, 36, 227, 3, 2, 2, 2, 38, 242, 3, 2, 2,
	2, 40, 244, 3, 2, 2, 2, 42, 246, 3, 2, 2, 2, 44, 263, 3, 2, 2, 2, 46, 267,
	3, 2, 2, 2, 48, 269, 3, 2, 2, 2, 50, 273, 3, 2, 2, 2, 52, 277, 3, 2, 2,
	2, 54, 279, 3, 2, 2, 2, 56, 283, 3, 2, 2, 2, 58, 285, 3, 2, 2, 2, 60, 287,
	3, 2, 2, 2, 62, 289, 3, 2, 2, 2, 64, 293, 3, 2, 2, 2, 66, 295, 3, 2, 2,
	2, 68, 299, 3, 2, 2, 2, 70, 301, 3, 2, 2, 2, 72, 315, 3, 2, 2, 2, 74, 331,
	3, 2, 2, 2, 76, 347, 3, 2, 2, 2, 78, 361, 3, 2, 2, 2, 80, 370, 3, 2, 2,
	2, 82, 376, 3, 2, 2, 2, 84, 390, 3, 2, 2, 2, 86, 399, 3, 2, 2, 2, 88, 412,
	3, 2, 2, 2, 90, 430, 3, 2, 2, 2, 92, 439, 3, 2, 2, 2, 94, 446, 3, 2, 2,
	2, 96, 453, 3, 2, 2, 2, 98, 460, 3, 2, 2, 2, 100, 479, 3, 2, 2, 2, 102,
	503, 3, 2, 2, 2, 104, 106, 7, 32, 2, 2, 105, 107, 5, 4, 3, 2, 106, 105,
	3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 111, 3, 2, 2, 2, 108, 110, 5, 16,
	9, 2, 109, 108, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2,
	111, 112, 3, 2, 2, 2, 112, 121, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114,
	118, 5, 12, 7, 2, 115, 117, 5, 12, 7, 2, 116, 115, 3, 2, 2, 2, 117, 120,
	3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 122, 3, 2,
	2, 2, 120, 118, 3, 2, 2, 2, 121, 114, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2,
	122, 123, 3, 2, 2, 2, 123, 124, 7, 33, 2, 2, 124, 3, 3, 2, 2, 2, 125, 128,
	5, 6, 4, 2, 126, 128, 5, 8, 5, 2, 127, 125, 3, 2, 2, 2, 127, 126, 3, 2,
	2, 2, 128, 132, 3, 2, 2, 2, 129, 131, 5, 8, 5, 2, 130, 129, 3, 2, 2, 2,
	131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133,
	5, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 136, 7, 3, 2, 2, 136, 137, 7,
	39, 2, 2, 137, 7, 3, 2, 2, 2, 138, 139, 7, 4, 2, 2, 139, 140, 7, 44, 2,
	2, 140, 141, 5, 10, 6, 2, 141, 9, 3, 2, 2, 2, 142, 143, 7, 39, 2, 2, 143,
	11, 3, 2, 2, 2, 144, 145, 7, 34, 2, 2, 145, 147, 5, 14, 8, 2, 146, 148,
	5, 4, 3, 2, 147, 146, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 152, 3, 2,
	2, 2, 149, 151, 5, 16, 9, 2, 150, 149, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2,
	152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 155, 3, 2, 2, 2, 154,
	152, 3, 2, 2, 2, 155, 156, 7, 35, 2, 2, 156, 13, 3, 2, 2, 2, 157, 158,
	9, 2, 2, 2, 158, 15, 3, 2, 2, 2, 159, 178, 5, 18, 10, 2, 160, 178, 5, 36,
	19, 2, 161, 178, 5, 42, 22, 2, 162, 178, 5, 70, 36, 2, 163, 178, 5, 72,
	37, 2, 164, 178, 5, 74, 38, 2, 165, 178, 5, 76, 39, 2, 166, 178, 5, 78,
	40, 2, 167, 178, 5, 80, 41, 2, 168, 178, 5, 82, 42, 2, 169, 178, 5, 84,
	43, 2, 170, 178, 5, 86, 44, 2, 171, 178, 5, 88, 45, 2, 172, 178, 5, 90,
	46, 2, 173, 178, 5, 92, 47, 2, 174, 178, 5, 94, 48, 2, 175, 178, 5, 96,
	49, 2, 176, 178, 5, 98, 50, 2, 177, 159, 3, 2, 2, 2, 177, 160, 3, 2, 2,
	2, 177, 161, 3, 2, 2, 2, 177, 162, 3, 2, 2, 2, 177, 163, 3, 2, 2, 2, 177,
	164, 3, 2, 2, 2, 177, 165, 3, 2, 2, 2, 177, 166, 3, 2, 2, 2, 177, 167,
	3, 2, 2, 2, 177, 168, 3, 2, 2, 2, 177, 169, 3, 2, 2, 2, 177, 170, 3, 2,
	2, 2, 177, 171, 3, 2, 2, 2, 177, 172, 3, 2, 2, 2, 177, 173, 3, 2, 2, 2,
	177, 174, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 176, 3, 2, 2, 2, 178,
	17, 3, 2, 2, 2, 179, 180, 7, 5, 2, 2, 180, 181, 7, 6, 2, 2, 181, 182, 5,
	14, 8, 2, 182, 183, 5, 20, 11, 2, 183, 184, 7, 7, 2, 2, 184, 19, 3, 2,
	2, 2, 185, 186, 7, 8, 2, 2, 186, 187, 7, 9, 2, 2, 187, 188, 5, 22, 12,
	2, 188, 189, 7, 10, 2, 2, 189, 191, 3, 2, 2, 2, 190, 185, 3, 2, 2, 2, 190,
	191, 3, 2, 2, 2, 191, 21, 3, 2, 2, 2, 192, 202, 3, 2, 2, 2, 193, 198, 5,
	24, 13, 2, 194, 195, 7, 8, 2, 2, 195, 197, 5, 24, 13, 2, 196, 194, 3, 2,
	2, 2, 197, 200, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2,
	199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 201, 192, 3, 2, 2, 2, 201,
	193, 3, 2, 2, 2, 202, 23, 3, 2, 2, 2, 203, 204, 5, 26, 14, 2, 204, 205,
	7, 11, 2, 2, 205, 206, 5, 28, 15, 2, 206, 25, 3, 2, 2, 2, 207, 208, 9,
	2, 2, 2, 208, 27, 3, 2, 2, 2, 209, 212, 5, 30, 16, 2, 210, 212, 5, 34,
	18, 2, 211, 209, 3, 2, 2, 2, 211, 210, 3, 2, 2, 2, 212, 29, 3, 2, 2, 2,
	213, 214, 7, 47, 2, 2, 214, 215, 7, 12, 2, 2, 215, 216, 5, 32, 17, 2, 216,
	31, 3, 2, 2, 2, 217, 218, 9, 2, 2, 2, 218, 33, 3, 2, 2, 2, 219, 221, 7,
	47, 2, 2, 220, 222, 7, 54, 2, 2, 221, 220, 3, 2, 2, 2, 221, 222, 3, 2,
	2, 2, 222, 226, 3, 2, 2, 2, 223, 226, 7, 48, 2, 2, 224, 226, 7, 49, 2,
	2, 225, 219, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 224, 3, 2, 2, 2, 226,
	35, 3, 2, 2, 2, 227, 228, 7, 13, 2, 2, 228, 229, 7, 6, 2, 2, 229, 235,
	5, 14, 8, 2, 230, 231, 7, 8, 2, 2, 231, 232, 5, 38, 20, 2, 232, 233, 7,
	8, 2, 2, 233, 234, 5, 38, 20, 2, 234, 236, 3, 2, 2, 2, 235, 230, 3, 2,
	2, 2, 235, 236, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 238, 5, 20, 11,
	2, 238, 239, 7, 7, 2, 2, 239, 37, 3, 2, 2, 2, 240, 243, 5, 40, 21, 2, 241,
	243, 7, 43, 2, 2, 242, 240, 3, 2, 2, 2, 242, 241, 3, 2, 2, 2, 243, 39,
	3, 2, 2, 2, 244, 245, 7, 53, 2, 2, 245, 41, 3, 2, 2, 2, 246, 247, 7, 14,
	2, 2, 247, 248, 7, 6, 2, 2, 248, 249, 5, 44, 23, 2, 249, 255, 5, 48, 25,
	2, 250, 251, 7, 8, 2, 2, 251, 252, 5, 52, 27, 2, 252, 253, 7, 8, 2, 2,
	253, 254, 5, 38, 20, 2, 254, 256, 3, 2, 2, 2, 255, 250, 3, 2, 2, 2, 255,
	256, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 258, 5, 20, 11, 2, 258, 259,
	7, 7, 2, 2, 259, 43, 3, 2, 2, 2, 260, 261, 5, 46, 24, 2, 261, 262, 7, 15,
	2, 2, 262, 264, 3, 2, 2, 2, 263, 260, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2,
	264, 45, 3, 2, 2, 2, 265, 268, 5, 14, 8, 2, 266, 268, 7, 43, 2, 2, 267,
	265, 3, 2, 2, 2, 267, 266, 3, 2, 2, 2, 268, 47, 3, 2, 2, 2, 269, 270, 5,
	14, 8, 2, 270, 49, 3, 2, 2, 2, 271, 274, 5, 48, 25, 2, 272, 274, 7, 43,
	2, 2, 273, 271, 3, 2, 2, 2, 273, 272, 3, 2, 2, 2, 274, 51, 3, 2, 2, 2,
	275, 278, 5, 54, 28, 2, 276, 278, 7, 43, 2, 2, 277, 275, 3, 2, 2, 2, 277,
	276, 3, 2, 2, 2, 278, 53, 3, 2, 2, 2, 279, 280, 5, 14, 8, 2, 280, 55, 3,
	2, 2, 2, 281, 284, 5, 58, 30, 2, 282, 284, 7, 43, 2, 2, 283, 281, 3, 2,
	2, 2, 283, 282, 3, 2, 2, 2, 284, 57, 3, 2, 2, 2, 285, 286, 5, 14, 8, 2,
	286, 59, 3, 2, 2, 2, 287, 288, 5, 14, 8, 2, 288, 61, 3, 2, 2, 2, 289, 290,
	5, 14, 8, 2, 290, 63, 3, 2, 2, 2, 291, 294, 5, 62, 32, 2, 292, 294, 7,
	43, 2, 2, 293, 291, 3, 2, 2, 2, 293, 292, 3, 2, 2, 2, 294, 65, 3, 2, 2,
	2, 295, 296, 5, 14, 8, 2, 296, 67, 3, 2, 2, 2, 297, 300, 5, 66, 34, 2,
	298, 300, 7, 43, 2, 2, 299, 297, 3, 2, 2, 2, 299, 298, 3, 2, 2, 2, 300,
	69, 3, 2, 2, 2, 301, 302, 7, 16, 2, 2, 302, 303, 7, 6, 2, 2, 303, 304,
	5, 44, 23, 2, 304, 310, 5, 54, 28, 2, 305, 306, 7, 8, 2, 2, 306, 307, 5,
	50, 26, 2, 307, 308, 7, 8, 2, 2, 308, 309, 5, 38, 20, 2, 309, 311, 3, 2,
	2, 2, 310, 305, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2,
	312, 313, 5, 20, 11, 2, 313, 314, 7, 7, 2, 2, 314, 71, 3, 2, 2, 2, 315,
	316, 7, 17, 2, 2, 316, 317, 7, 6, 2, 2, 317, 318, 5, 44, 23, 2, 318, 326,
	5, 54, 28, 2, 319, 320, 7, 8, 2, 2, 320, 321, 5, 50, 26, 2, 321, 322, 7,
	8, 2, 2, 322, 323, 5, 52, 27, 2, 323, 324, 7, 8, 2, 2, 324, 325, 5, 38,
	20, 2, 325, 327, 3, 2, 2, 2, 326, 319, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2,
	327, 328, 3, 2, 2, 2, 328, 329, 5, 20, 11, 2, 329, 330, 7, 7, 2, 2, 330,
	73, 3, 2, 2, 2, 331, 332, 7, 18, 2, 2, 332, 333, 7, 6, 2, 2, 333, 334,
	5, 44, 23, 2, 334, 342, 5, 54, 28, 2, 335, 336, 7, 8, 2, 2, 336, 337, 5,
	50, 26, 2, 337, 338, 7, 8, 2, 2, 338, 339, 5, 52, 27, 2, 339, 340, 7, 8,
	2, 2, 340, 341, 5, 38, 20, 2, 341, 343, 3, 2, 2, 2, 342, 335, 3, 2, 2,
	2, 342, 343, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 345, 5, 20, 11, 2,
	345, 346, 7, 7, 2, 2, 346, 75, 3, 2, 2, 2, 347, 348, 7, 19, 2, 2, 348,
	349, 7, 6, 2, 2, 349, 350, 5, 44, 23, 2, 350, 356, 5, 48, 25, 2, 351, 352,
	7, 8, 2, 2, 352, 353, 5, 52, 27, 2, 353, 354, 7, 8, 2, 2, 354, 355, 5,
	38, 20, 2, 355, 357, 3, 2, 2, 2, 356, 351, 3, 2, 2, 2, 356, 357, 3, 2,
	2, 2, 357, 358, 3, 2, 2, 2, 358, 359, 5, 20, 11, 2, 359, 360, 7, 7, 2,
	2, 360, 77, 3, 2, 2, 2, 361, 362, 7, 20, 2, 2, 362, 363, 7, 6, 2, 2, 363,
	364, 5, 44, 23, 2, 364, 365, 5, 54, 28, 2, 365, 366, 7, 8, 2, 2, 366, 367,
	5, 54, 28, 2, 367, 368, 5, 20, 11, 2, 368, 369, 7, 7, 2, 2, 369, 79, 3,
	2, 2, 2, 370, 371, 7, 21, 2, 2, 371, 372, 7, 6, 2, 2, 372, 373, 5, 14,
	8, 2, 373, 374, 5, 20, 11, 2, 374, 375, 7, 7, 2, 2, 375, 81, 3, 2, 2, 2,
	376, 377, 7, 22, 2, 2, 377, 378, 7, 6, 2, 2, 378, 379, 5, 44, 23, 2, 379,
	385, 5, 54, 28, 2, 380, 381, 7, 8, 2, 2, 381, 382, 5, 56, 29, 2, 382, 383,
	7, 8, 2, 2, 383, 384, 5, 50, 26, 2, 384, 386, 3, 2, 2, 2, 385, 380, 3,
	2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 388, 5, 20, 11,
	2, 388, 389, 7, 7, 2, 2, 389, 83, 3, 2, 2, 2, 390, 391, 7, 23, 2, 2, 391,
	392, 7, 6, 2, 2, 392, 393, 5, 44, 23, 2, 393, 394, 5, 48, 25, 2, 394, 395,
	7, 8, 2, 2, 395, 396, 5, 58, 30, 2, 396, 397, 5, 20, 11, 2, 397, 398, 7,
	7, 2, 2, 398, 85, 3, 2, 2, 2, 399, 400, 7, 24, 2, 2, 400, 401, 7, 6, 2,
	2, 401, 402, 5, 44, 23, 2, 402, 403, 5, 58, 30, 2, 403, 404, 7, 8, 2, 2,
	404, 407, 5, 58, 30, 2, 405, 406, 7, 8, 2, 2, 406, 408, 5, 52, 27, 2, 407,
	405, 3, 2, 2, 2, 407, 408, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 410,
	5, 20, 11, 2, 410, 411, 7, 7, 2, 2, 411, 87, 3, 2, 2, 2, 412, 413, 7, 25,
	2, 2, 413, 414, 7, 6, 2, 2, 414, 415, 5, 44, 23, 2, 415, 416, 5, 48, 25,
	2, 416, 417, 7, 8, 2, 2, 417, 425, 5, 48, 25, 2, 418, 419, 7, 8, 2, 2,
	419, 420, 5, 52, 27, 2, 420, 421, 7, 8, 2, 2, 421, 422, 5, 64, 33, 2, 422,
	423, 7, 8, 2, 2, 423, 424, 5, 68, 35, 2, 424, 426, 3, 2, 2, 2, 425, 418,
	3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 427, 3, 2, 2, 2, 427, 428, 5, 20,
	11, 2, 428, 429, 7, 7, 2, 2, 429, 89, 3, 2, 2, 2, 430, 431, 7, 26, 2, 2,
	431, 432, 7, 6, 2, 2, 432, 433, 5, 44, 23, 2, 433, 434, 5, 48, 25, 2, 434,
	435, 7, 8, 2, 2, 435, 436, 5, 48, 25, 2, 436, 437, 5, 20, 11, 2, 437, 438,
	7, 7, 2, 2, 438, 91, 3, 2, 2, 2, 439, 440, 7, 27, 2, 2, 440, 441, 7, 6,
	2, 2, 441, 442, 5, 48, 25, 2, 442, 443, 7, 8, 2, 2, 443, 444, 5, 48, 25,
	2, 444, 445, 7, 7, 2, 2, 445, 93, 3, 2, 2, 2, 446, 447, 7, 28, 2, 2, 447,
	448, 7, 6, 2, 2, 448, 449, 5, 48, 25, 2, 449, 450, 7, 8, 2, 2, 450, 451,
	5, 48, 25, 2, 451, 452, 7, 7, 2, 2, 452, 95, 3, 2, 2, 2, 453, 454, 7, 29,
	2, 2, 454, 455, 7, 6, 2, 2, 455, 456, 5, 60, 31, 2, 456, 457, 7, 8, 2,
	2, 457, 458, 5, 48, 25, 2, 458, 459, 7, 7, 2, 2, 459, 97, 3, 2, 2, 2, 460,
	461, 9, 2, 2, 2, 461, 462, 7, 6, 2, 2, 462, 463, 5, 44, 23, 2, 463, 468,
	5, 100, 51, 2, 464, 465, 7, 8, 2, 2, 465, 467, 5, 100, 51, 2, 466, 464,
	3, 2, 2, 2, 467, 470, 3, 2, 2, 2, 468, 466, 3, 2, 2, 2, 468, 469, 3, 2,
	2, 2, 469, 471, 3, 2, 2, 2, 470, 468, 3, 2, 2, 2, 471, 472, 5, 20, 11,
	2, 472, 473, 7, 7, 2, 2, 473, 99, 3, 2, 2, 2, 474, 480, 5, 46, 24, 2, 475,
	480, 5, 28, 15, 2, 476, 480, 5, 40, 21, 2, 477, 480, 5, 98, 50, 2, 478,
	480, 5, 102, 52, 2, 479, 474, 3, 2, 2, 2, 479, 475, 3, 2, 2, 2, 479, 476,
	3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 479, 478, 3, 2, 2, 2, 480, 101, 3, 2,
	2, 2, 481, 482, 7, 30, 2, 2, 482, 487, 5, 100, 51, 2, 483, 484, 7, 8, 2,
	2, 484, 486, 5, 100, 51, 2, 485, 483, 3, 2, 2, 2, 486, 489, 3, 2, 2, 2,
	487, 485, 3, 2, 2, 2, 487, 488, 3, 2, 2, 2, 488, 490, 3, 2, 2, 2, 489,
	487, 3, 2, 2, 2, 490, 491, 7, 31, 2, 2, 491, 504, 3, 2, 2, 2, 492, 493,
	7, 6, 2, 2, 493, 498, 5, 100, 51, 2, 494, 495, 7, 8, 2, 2, 495, 497, 5,
	100, 51, 2, 496, 494, 3, 2, 2, 2, 497, 500, 3, 2, 2, 2, 498, 496, 3, 2,
	2, 2, 498, 499, 3, 2, 2, 2, 499, 501, 3, 2, 2, 2, 500, 498, 3, 2, 2, 2,
	501, 502, 7, 7, 2, 2, 502, 504, 3, 2, 2, 2, 503, 481, 3, 2, 2, 2, 503,
	492, 3, 2, 2, 2, 504, 103, 3, 2, 2, 2, 39, 106, 111, 118, 121, 127, 132,
	147, 152, 177, 190, 198, 201, 211, 221, 225, 235, 242, 255, 263, 267, 273,
	277, 283, 293, 299, 310, 326, 342, 356, 385, 407, 425, 468, 479, 487, 498,
	503,
}
var literalNames = []string{
	"", "'default'", "'prefix'", "'entity'", "'('", "')'", "','", "'['", "']'",
	"'='", "'%%'", "'activity'", "'wasGeneratedBy'", "';'", "'used'", "'wasStartedBy'",
	"'wasEndedBy'", "'wasInvalidatedBy'", "'wasInformedBy'", "'agent'", "'wasAssociatedWith'",
	"'wasAttributedTo'", "'actedOnBehalfOf'", "'wasDerivedFrom'", "'wasInfluencedBy'",
	"'alternateOf'", "'specializationOf'", "'hadMember'", "'{'", "'}'", "'document'",
	"'endDocument'", "'bundle'", "'endBundle'", "", "", "", "", "'<'", "'>'",
	"'.'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "DOCUMENT", "ENDDOCUMENT",
	"BUNDLE", "ENDBUNDLE", "WS", "COMMENT", "LINE_COMMENT", "IRI_REF", "LESS",
	"GREATER", "DOT", "MINUS", "PREFX", "QUALIFIED_NAME", "HEX", "STRING_LITERAL",
	"INT_LITERAL", "QUALIFIED_NAME_LITERAL", "ECHAR", "STRING_LITERAL2", "STRING_LITERAL_LONG2",
	"DATETIME", "LANGTAG",
}

var ruleNames = []string{
	"document", "namespaceDeclarations", "defaultNamespaceDeclaration", "namespaceDeclaration",
	"namespace_", "bundle", "identifier", "expression", "entityExpression",
	"optionalAttributeValuePairs", "attributeValuePairs", "attributeValuePair",
	"attribute", "literal", "typedLiteral", "datatype", "convenienceNotation",
	"activityExpression", "timeOrMarker", "time", "generationExpression", "optionalIdentifier",
	"identifierOrMarker", "eIdentifier", "eIdentifierOrMarker", "aIdentifierOrMarker",
	"aIdentifier", "agIdentifierOrMarker", "agIdentifier", "cIdentifier", "gIdentifier",
	"gIdentifierOrMarker", "uIdentifier", "uIdentifierOrMarker", "usageExpression",
	"startExpression", "endExpression", "invalidationExpression", "communicationExpression",
	"agentExpression", "associationExpression", "attributionExpression", "delegationExpression",
	"derivationExpression", "influenceExpression", "alternateExpression", "specializationExpression",
	"membershipExpression", "extensibilityExpression", "extensibilityArgument",
	"extensibilityTuple",
}

type PROV_NParser struct {
	*antlr.BaseParser
}

// NewPROV_NParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *PROV_NParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPROV_NParser(input antlr.TokenStream) *PROV_NParser {
	this := new(PROV_NParser)
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
	this.GrammarFileName = "PROV_N.g4"

	return this
}

// PROV_NParser tokens.
const (
	PROV_NParserEOF                    = antlr.TokenEOF
	PROV_NParserT__0                   = 1
	PROV_NParserT__1                   = 2
	PROV_NParserT__2                   = 3
	PROV_NParserT__3                   = 4
	PROV_NParserT__4                   = 5
	PROV_NParserT__5                   = 6
	PROV_NParserT__6                   = 7
	PROV_NParserT__7                   = 8
	PROV_NParserT__8                   = 9
	PROV_NParserT__9                   = 10
	PROV_NParserT__10                  = 11
	PROV_NParserT__11                  = 12
	PROV_NParserT__12                  = 13
	PROV_NParserT__13                  = 14
	PROV_NParserT__14                  = 15
	PROV_NParserT__15                  = 16
	PROV_NParserT__16                  = 17
	PROV_NParserT__17                  = 18
	PROV_NParserT__18                  = 19
	PROV_NParserT__19                  = 20
	PROV_NParserT__20                  = 21
	PROV_NParserT__21                  = 22
	PROV_NParserT__22                  = 23
	PROV_NParserT__23                  = 24
	PROV_NParserT__24                  = 25
	PROV_NParserT__25                  = 26
	PROV_NParserT__26                  = 27
	PROV_NParserT__27                  = 28
	PROV_NParserT__28                  = 29
	PROV_NParserDOCUMENT               = 30
	PROV_NParserENDDOCUMENT            = 31
	PROV_NParserBUNDLE                 = 32
	PROV_NParserENDBUNDLE              = 33
	PROV_NParserWS                     = 34
	PROV_NParserCOMMENT                = 35
	PROV_NParserLINE_COMMENT           = 36
	PROV_NParserIRI_REF                = 37
	PROV_NParserLESS                   = 38
	PROV_NParserGREATER                = 39
	PROV_NParserDOT                    = 40
	PROV_NParserMINUS                  = 41
	PROV_NParserPREFX                  = 42
	PROV_NParserQUALIFIED_NAME         = 43
	PROV_NParserHEX                    = 44
	PROV_NParserSTRING_LITERAL         = 45
	PROV_NParserINT_LITERAL            = 46
	PROV_NParserQUALIFIED_NAME_LITERAL = 47
	PROV_NParserECHAR                  = 48
	PROV_NParserSTRING_LITERAL2        = 49
	PROV_NParserSTRING_LITERAL_LONG2   = 50
	PROV_NParserDATETIME               = 51
	PROV_NParserLANGTAG                = 52
)

// PROV_NParser rules.
const (
	PROV_NParserRULE_document                    = 0
	PROV_NParserRULE_namespaceDeclarations       = 1
	PROV_NParserRULE_defaultNamespaceDeclaration = 2
	PROV_NParserRULE_namespaceDeclaration        = 3
	PROV_NParserRULE_namespace_                  = 4
	PROV_NParserRULE_bundle                      = 5
	PROV_NParserRULE_identifier                  = 6
	PROV_NParserRULE_expression                  = 7
	PROV_NParserRULE_entityExpression            = 8
	PROV_NParserRULE_optionalAttributeValuePairs = 9
	PROV_NParserRULE_attributeValuePairs         = 10
	PROV_NParserRULE_attributeValuePair          = 11
	PROV_NParserRULE_attribute                   = 12
	PROV_NParserRULE_literal                     = 13
	PROV_NParserRULE_typedLiteral                = 14
	PROV_NParserRULE_datatype                    = 15
	PROV_NParserRULE_convenienceNotation         = 16
	PROV_NParserRULE_activityExpression          = 17
	PROV_NParserRULE_timeOrMarker                = 18
	PROV_NParserRULE_time                        = 19
	PROV_NParserRULE_generationExpression        = 20
	PROV_NParserRULE_optionalIdentifier          = 21
	PROV_NParserRULE_identifierOrMarker          = 22
	PROV_NParserRULE_eIdentifier                 = 23
	PROV_NParserRULE_eIdentifierOrMarker         = 24
	PROV_NParserRULE_aIdentifierOrMarker         = 25
	PROV_NParserRULE_aIdentifier                 = 26
	PROV_NParserRULE_agIdentifierOrMarker        = 27
	PROV_NParserRULE_agIdentifier                = 28
	PROV_NParserRULE_cIdentifier                 = 29
	PROV_NParserRULE_gIdentifier                 = 30
	PROV_NParserRULE_gIdentifierOrMarker         = 31
	PROV_NParserRULE_uIdentifier                 = 32
	PROV_NParserRULE_uIdentifierOrMarker         = 33
	PROV_NParserRULE_usageExpression             = 34
	PROV_NParserRULE_startExpression             = 35
	PROV_NParserRULE_endExpression               = 36
	PROV_NParserRULE_invalidationExpression      = 37
	PROV_NParserRULE_communicationExpression     = 38
	PROV_NParserRULE_agentExpression             = 39
	PROV_NParserRULE_associationExpression       = 40
	PROV_NParserRULE_attributionExpression       = 41
	PROV_NParserRULE_delegationExpression        = 42
	PROV_NParserRULE_derivationExpression        = 43
	PROV_NParserRULE_influenceExpression         = 44
	PROV_NParserRULE_alternateExpression         = 45
	PROV_NParserRULE_specializationExpression    = 46
	PROV_NParserRULE_membershipExpression        = 47
	PROV_NParserRULE_extensibilityExpression     = 48
	PROV_NParserRULE_extensibilityArgument       = 49
	PROV_NParserRULE_extensibilityTuple          = 50
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
	p.RuleIndex = PROV_NParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) DOCUMENT() antlr.TerminalNode {
	return s.GetToken(PROV_NParserDOCUMENT, 0)
}

func (s *DocumentContext) ENDDOCUMENT() antlr.TerminalNode {
	return s.GetToken(PROV_NParserENDDOCUMENT, 0)
}

func (s *DocumentContext) NamespaceDeclarations() INamespaceDeclarationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceDeclarationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceDeclarationsContext)
}

func (s *DocumentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DocumentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DocumentContext) AllBundle() []IBundleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBundleContext)(nil)).Elem())
	var tst = make([]IBundleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBundleContext)
		}
	}

	return tst
}

func (s *DocumentContext) Bundle(i int) IBundleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBundleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBundleContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *PROV_NParser) Document() (localctx IDocumentContext) {
	this := p
	_ = this

	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PROV_NParserRULE_document)
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
		p.Match(PROV_NParserDOCUMENT)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PROV_NParserT__0 || _la == PROV_NParserT__1 {
		{
			p.SetState(103)
			p.NamespaceDeclarations()
		}

	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PROV_NParserT__2)|(1<<PROV_NParserT__10)|(1<<PROV_NParserT__11)|(1<<PROV_NParserT__13)|(1<<PROV_NParserT__14)|(1<<PROV_NParserT__15)|(1<<PROV_NParserT__16)|(1<<PROV_NParserT__17)|(1<<PROV_NParserT__18)|(1<<PROV_NParserT__19)|(1<<PROV_NParserT__20)|(1<<PROV_NParserT__21)|(1<<PROV_NParserT__22)|(1<<PROV_NParserT__23)|(1<<PROV_NParserT__24)|(1<<PROV_NParserT__25)|(1<<PROV_NParserT__26))) != 0) || _la == PROV_NParserPREFX || _la == PROV_NParserQUALIFIED_NAME {
		{
			p.SetState(106)
			p.Expression()
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PROV_NParserBUNDLE {
		{
			p.SetState(112)
			p.Bundle()
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PROV_NParserBUNDLE {
			{
				p.SetState(113)
				p.Bundle()
			}

			p.SetState(118)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(121)
		p.Match(PROV_NParserENDDOCUMENT)
	}

	return localctx
}

// INamespaceDeclarationsContext is an interface to support dynamic dispatch.
type INamespaceDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceDeclarationsContext differentiates from other interfaces.
	IsNamespaceDeclarationsContext()
}

type NamespaceDeclarationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceDeclarationsContext() *NamespaceDeclarationsContext {
	var p = new(NamespaceDeclarationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_namespaceDeclarations
	return p
}

func (*NamespaceDeclarationsContext) IsNamespaceDeclarationsContext() {}

func NewNamespaceDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDeclarationsContext {
	var p = new(NamespaceDeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_namespaceDeclarations

	return p
}

func (s *NamespaceDeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDeclarationsContext) DefaultNamespaceDeclaration() IDefaultNamespaceDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultNamespaceDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultNamespaceDeclarationContext)
}

func (s *NamespaceDeclarationsContext) AllNamespaceDeclaration() []INamespaceDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INamespaceDeclarationContext)(nil)).Elem())
	var tst = make([]INamespaceDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INamespaceDeclarationContext)
		}
	}

	return tst
}

func (s *NamespaceDeclarationsContext) NamespaceDeclaration(i int) INamespaceDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INamespaceDeclarationContext)
}

func (s *NamespaceDeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterNamespaceDeclarations(s)
	}
}

func (s *NamespaceDeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitNamespaceDeclarations(s)
	}
}

func (p *PROV_NParser) NamespaceDeclarations() (localctx INamespaceDeclarationsContext) {
	this := p
	_ = this

	localctx = NewNamespaceDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PROV_NParserRULE_namespaceDeclarations)
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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserT__0:
		{
			p.SetState(123)
			p.DefaultNamespaceDeclaration()
		}

	case PROV_NParserT__1:
		{
			p.SetState(124)
			p.NamespaceDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PROV_NParserT__1 {
		{
			p.SetState(127)
			p.NamespaceDeclaration()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefaultNamespaceDeclarationContext is an interface to support dynamic dispatch.
type IDefaultNamespaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultNamespaceDeclarationContext differentiates from other interfaces.
	IsDefaultNamespaceDeclarationContext()
}

type DefaultNamespaceDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultNamespaceDeclarationContext() *DefaultNamespaceDeclarationContext {
	var p = new(DefaultNamespaceDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_defaultNamespaceDeclaration
	return p
}

func (*DefaultNamespaceDeclarationContext) IsDefaultNamespaceDeclarationContext() {}

func NewDefaultNamespaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultNamespaceDeclarationContext {
	var p = new(DefaultNamespaceDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_defaultNamespaceDeclaration

	return p
}

func (s *DefaultNamespaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultNamespaceDeclarationContext) IRI_REF() antlr.TerminalNode {
	return s.GetToken(PROV_NParserIRI_REF, 0)
}

func (s *DefaultNamespaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultNamespaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultNamespaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterDefaultNamespaceDeclaration(s)
	}
}

func (s *DefaultNamespaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitDefaultNamespaceDeclaration(s)
	}
}

func (p *PROV_NParser) DefaultNamespaceDeclaration() (localctx IDefaultNamespaceDeclarationContext) {
	this := p
	_ = this

	localctx = NewDefaultNamespaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PROV_NParserRULE_defaultNamespaceDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(133)
		p.Match(PROV_NParserT__0)
	}
	{
		p.SetState(134)
		p.Match(PROV_NParserIRI_REF)
	}

	return localctx
}

// INamespaceDeclarationContext is an interface to support dynamic dispatch.
type INamespaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceDeclarationContext differentiates from other interfaces.
	IsNamespaceDeclarationContext()
}

type NamespaceDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceDeclarationContext() *NamespaceDeclarationContext {
	var p = new(NamespaceDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_namespaceDeclaration
	return p
}

func (*NamespaceDeclarationContext) IsNamespaceDeclarationContext() {}

func NewNamespaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDeclarationContext {
	var p = new(NamespaceDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_namespaceDeclaration

	return p
}

func (s *NamespaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDeclarationContext) PREFX() antlr.TerminalNode {
	return s.GetToken(PROV_NParserPREFX, 0)
}

func (s *NamespaceDeclarationContext) Namespace_() INamespace_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespace_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespace_Context)
}

func (s *NamespaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterNamespaceDeclaration(s)
	}
}

func (s *NamespaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitNamespaceDeclaration(s)
	}
}

func (p *PROV_NParser) NamespaceDeclaration() (localctx INamespaceDeclarationContext) {
	this := p
	_ = this

	localctx = NewNamespaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PROV_NParserRULE_namespaceDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(136)
		p.Match(PROV_NParserT__1)
	}
	{
		p.SetState(137)
		p.Match(PROV_NParserPREFX)
	}
	{
		p.SetState(138)
		p.Namespace_()
	}

	return localctx
}

// INamespace_Context is an interface to support dynamic dispatch.
type INamespace_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespace_Context differentiates from other interfaces.
	IsNamespace_Context()
}

type Namespace_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespace_Context() *Namespace_Context {
	var p = new(Namespace_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_namespace_
	return p
}

func (*Namespace_Context) IsNamespace_Context() {}

func NewNamespace_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Namespace_Context {
	var p = new(Namespace_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_namespace_

	return p
}

func (s *Namespace_Context) GetParser() antlr.Parser { return s.parser }

func (s *Namespace_Context) IRI_REF() antlr.TerminalNode {
	return s.GetToken(PROV_NParserIRI_REF, 0)
}

func (s *Namespace_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Namespace_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Namespace_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterNamespace_(s)
	}
}

func (s *Namespace_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitNamespace_(s)
	}
}

func (p *PROV_NParser) Namespace_() (localctx INamespace_Context) {
	this := p
	_ = this

	localctx = NewNamespace_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PROV_NParserRULE_namespace_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(140)
		p.Match(PROV_NParserIRI_REF)
	}

	return localctx
}

// IBundleContext is an interface to support dynamic dispatch.
type IBundleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBundleContext differentiates from other interfaces.
	IsBundleContext()
}

type BundleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBundleContext() *BundleContext {
	var p = new(BundleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_bundle
	return p
}

func (*BundleContext) IsBundleContext() {}

func NewBundleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BundleContext {
	var p = new(BundleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_bundle

	return p
}

func (s *BundleContext) GetParser() antlr.Parser { return s.parser }

func (s *BundleContext) BUNDLE() antlr.TerminalNode {
	return s.GetToken(PROV_NParserBUNDLE, 0)
}

func (s *BundleContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *BundleContext) ENDBUNDLE() antlr.TerminalNode {
	return s.GetToken(PROV_NParserENDBUNDLE, 0)
}

func (s *BundleContext) NamespaceDeclarations() INamespaceDeclarationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceDeclarationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceDeclarationsContext)
}

func (s *BundleContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BundleContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BundleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BundleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BundleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterBundle(s)
	}
}

func (s *BundleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitBundle(s)
	}
}

func (p *PROV_NParser) Bundle() (localctx IBundleContext) {
	this := p
	_ = this

	localctx = NewBundleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PROV_NParserRULE_bundle)
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
		p.SetState(142)
		p.Match(PROV_NParserBUNDLE)
	}
	{
		p.SetState(143)
		p.Identifier()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PROV_NParserT__0 || _la == PROV_NParserT__1 {
		{
			p.SetState(144)
			p.NamespaceDeclarations()
		}

	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PROV_NParserT__2)|(1<<PROV_NParserT__10)|(1<<PROV_NParserT__11)|(1<<PROV_NParserT__13)|(1<<PROV_NParserT__14)|(1<<PROV_NParserT__15)|(1<<PROV_NParserT__16)|(1<<PROV_NParserT__17)|(1<<PROV_NParserT__18)|(1<<PROV_NParserT__19)|(1<<PROV_NParserT__20)|(1<<PROV_NParserT__21)|(1<<PROV_NParserT__22)|(1<<PROV_NParserT__23)|(1<<PROV_NParserT__24)|(1<<PROV_NParserT__25)|(1<<PROV_NParserT__26))) != 0) || _la == PROV_NParserPREFX || _la == PROV_NParserQUALIFIED_NAME {
		{
			p.SetState(147)
			p.Expression()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(PROV_NParserENDBUNDLE)
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
	p.RuleIndex = PROV_NParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) PREFX() antlr.TerminalNode {
	return s.GetToken(PROV_NParserPREFX, 0)
}

func (s *IdentifierContext) QUALIFIED_NAME() antlr.TerminalNode {
	return s.GetToken(PROV_NParserQUALIFIED_NAME, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *PROV_NParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PROV_NParserRULE_identifier)
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
		p.SetState(155)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PROV_NParserPREFX || _la == PROV_NParserQUALIFIED_NAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = PROV_NParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) EntityExpression() IEntityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntityExpressionContext)
}

func (s *ExpressionContext) ActivityExpression() IActivityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActivityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActivityExpressionContext)
}

func (s *ExpressionContext) GenerationExpression() IGenerationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGenerationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGenerationExpressionContext)
}

func (s *ExpressionContext) UsageExpression() IUsageExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUsageExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUsageExpressionContext)
}

func (s *ExpressionContext) StartExpression() IStartExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStartExpressionContext)
}

func (s *ExpressionContext) EndExpression() IEndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndExpressionContext)
}

func (s *ExpressionContext) InvalidationExpression() IInvalidationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvalidationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvalidationExpressionContext)
}

func (s *ExpressionContext) CommunicationExpression() ICommunicationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommunicationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommunicationExpressionContext)
}

func (s *ExpressionContext) AgentExpression() IAgentExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAgentExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAgentExpressionContext)
}

func (s *ExpressionContext) AssociationExpression() IAssociationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssociationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssociationExpressionContext)
}

func (s *ExpressionContext) AttributionExpression() IAttributionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributionExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributionExpressionContext)
}

func (s *ExpressionContext) DelegationExpression() IDelegationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelegationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelegationExpressionContext)
}

func (s *ExpressionContext) DerivationExpression() IDerivationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDerivationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDerivationExpressionContext)
}

func (s *ExpressionContext) InfluenceExpression() IInfluenceExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfluenceExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfluenceExpressionContext)
}

func (s *ExpressionContext) AlternateExpression() IAlternateExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlternateExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlternateExpressionContext)
}

func (s *ExpressionContext) SpecializationExpression() ISpecializationExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecializationExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecializationExpressionContext)
}

func (s *ExpressionContext) MembershipExpression() IMembershipExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMembershipExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMembershipExpressionContext)
}

func (s *ExpressionContext) ExtensibilityExpression() IExtensibilityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtensibilityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtensibilityExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PROV_NParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PROV_NParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserT__2:
		{
			p.SetState(157)
			p.EntityExpression()
		}

	case PROV_NParserT__10:
		{
			p.SetState(158)
			p.ActivityExpression()
		}

	case PROV_NParserT__11:
		{
			p.SetState(159)
			p.GenerationExpression()
		}

	case PROV_NParserT__13:
		{
			p.SetState(160)
			p.UsageExpression()
		}

	case PROV_NParserT__14:
		{
			p.SetState(161)
			p.StartExpression()
		}

	case PROV_NParserT__15:
		{
			p.SetState(162)
			p.EndExpression()
		}

	case PROV_NParserT__16:
		{
			p.SetState(163)
			p.InvalidationExpression()
		}

	case PROV_NParserT__17:
		{
			p.SetState(164)
			p.CommunicationExpression()
		}

	case PROV_NParserT__18:
		{
			p.SetState(165)
			p.AgentExpression()
		}

	case PROV_NParserT__19:
		{
			p.SetState(166)
			p.AssociationExpression()
		}

	case PROV_NParserT__20:
		{
			p.SetState(167)
			p.AttributionExpression()
		}

	case PROV_NParserT__21:
		{
			p.SetState(168)
			p.DelegationExpression()
		}

	case PROV_NParserT__22:
		{
			p.SetState(169)
			p.DerivationExpression()
		}

	case PROV_NParserT__23:
		{
			p.SetState(170)
			p.InfluenceExpression()
		}

	case PROV_NParserT__24:
		{
			p.SetState(171)
			p.AlternateExpression()
		}

	case PROV_NParserT__25:
		{
			p.SetState(172)
			p.SpecializationExpression()
		}

	case PROV_NParserT__26:
		{
			p.SetState(173)
			p.MembershipExpression()
		}

	case PROV_NParserPREFX, PROV_NParserQUALIFIED_NAME:
		{
			p.SetState(174)
			p.ExtensibilityExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEntityExpressionContext is an interface to support dynamic dispatch.
type IEntityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntityExpressionContext differentiates from other interfaces.
	IsEntityExpressionContext()
}

type EntityExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityExpressionContext() *EntityExpressionContext {
	var p = new(EntityExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_entityExpression
	return p
}

func (*EntityExpressionContext) IsEntityExpressionContext() {}

func NewEntityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityExpressionContext {
	var p = new(EntityExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_entityExpression

	return p
}

func (s *EntityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EntityExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *EntityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterEntityExpression(s)
	}
}

func (s *EntityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitEntityExpression(s)
	}
}

func (p *PROV_NParser) EntityExpression() (localctx IEntityExpressionContext) {
	this := p
	_ = this

	localctx = NewEntityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PROV_NParserRULE_entityExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(177)
		p.Match(PROV_NParserT__2)
	}
	{
		p.SetState(178)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(179)
		p.Identifier()
	}
	{
		p.SetState(180)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(181)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IOptionalAttributeValuePairsContext is an interface to support dynamic dispatch.
type IOptionalAttributeValuePairsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionalAttributeValuePairsContext differentiates from other interfaces.
	IsOptionalAttributeValuePairsContext()
}

type OptionalAttributeValuePairsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalAttributeValuePairsContext() *OptionalAttributeValuePairsContext {
	var p = new(OptionalAttributeValuePairsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_optionalAttributeValuePairs
	return p
}

func (*OptionalAttributeValuePairsContext) IsOptionalAttributeValuePairsContext() {}

func NewOptionalAttributeValuePairsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalAttributeValuePairsContext {
	var p = new(OptionalAttributeValuePairsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_optionalAttributeValuePairs

	return p
}

func (s *OptionalAttributeValuePairsContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalAttributeValuePairsContext) AttributeValuePairs() IAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeValuePairsContext)
}

func (s *OptionalAttributeValuePairsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalAttributeValuePairsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalAttributeValuePairsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterOptionalAttributeValuePairs(s)
	}
}

func (s *OptionalAttributeValuePairsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitOptionalAttributeValuePairs(s)
	}
}

func (p *PROV_NParser) OptionalAttributeValuePairs() (localctx IOptionalAttributeValuePairsContext) {
	this := p
	_ = this

	localctx = NewOptionalAttributeValuePairsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PROV_NParserRULE_optionalAttributeValuePairs)
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
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PROV_NParserT__5 {
		{
			p.SetState(183)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(184)
			p.Match(PROV_NParserT__6)
		}
		{
			p.SetState(185)
			p.AttributeValuePairs()
		}
		{
			p.SetState(186)
			p.Match(PROV_NParserT__7)
		}

	}

	return localctx
}

// IAttributeValuePairsContext is an interface to support dynamic dispatch.
type IAttributeValuePairsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeValuePairsContext differentiates from other interfaces.
	IsAttributeValuePairsContext()
}

type AttributeValuePairsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeValuePairsContext() *AttributeValuePairsContext {
	var p = new(AttributeValuePairsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_attributeValuePairs
	return p
}

func (*AttributeValuePairsContext) IsAttributeValuePairsContext() {}

func NewAttributeValuePairsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeValuePairsContext {
	var p = new(AttributeValuePairsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_attributeValuePairs

	return p
}

func (s *AttributeValuePairsContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeValuePairsContext) AllAttributeValuePair() []IAttributeValuePairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeValuePairContext)(nil)).Elem())
	var tst = make([]IAttributeValuePairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeValuePairContext)
		}
	}

	return tst
}

func (s *AttributeValuePairsContext) AttributeValuePair(i int) IAttributeValuePairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeValuePairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeValuePairContext)
}

func (s *AttributeValuePairsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeValuePairsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeValuePairsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAttributeValuePairs(s)
	}
}

func (s *AttributeValuePairsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAttributeValuePairs(s)
	}
}

func (p *PROV_NParser) AttributeValuePairs() (localctx IAttributeValuePairsContext) {
	this := p
	_ = this

	localctx = NewAttributeValuePairsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PROV_NParserRULE_attributeValuePairs)
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
	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserT__7:

	case PROV_NParserPREFX, PROV_NParserQUALIFIED_NAME:
		{
			p.SetState(191)
			p.AttributeValuePair()
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PROV_NParserT__5 {
			{
				p.SetState(192)
				p.Match(PROV_NParserT__5)
			}
			{
				p.SetState(193)
				p.AttributeValuePair()
			}

			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAttributeValuePairContext is an interface to support dynamic dispatch.
type IAttributeValuePairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeValuePairContext differentiates from other interfaces.
	IsAttributeValuePairContext()
}

type AttributeValuePairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeValuePairContext() *AttributeValuePairContext {
	var p = new(AttributeValuePairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_attributeValuePair
	return p
}

func (*AttributeValuePairContext) IsAttributeValuePairContext() {}

func NewAttributeValuePairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeValuePairContext {
	var p = new(AttributeValuePairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_attributeValuePair

	return p
}

func (s *AttributeValuePairContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeValuePairContext) Attribute() IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *AttributeValuePairContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *AttributeValuePairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeValuePairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeValuePairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAttributeValuePair(s)
	}
}

func (s *AttributeValuePairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAttributeValuePair(s)
	}
}

func (p *PROV_NParser) AttributeValuePair() (localctx IAttributeValuePairContext) {
	this := p
	_ = this

	localctx = NewAttributeValuePairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PROV_NParserRULE_attributeValuePair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Attribute()
	}
	{
		p.SetState(202)
		p.Match(PROV_NParserT__8)
	}
	{
		p.SetState(203)
		p.Literal()
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) PREFX() antlr.TerminalNode {
	return s.GetToken(PROV_NParserPREFX, 0)
}

func (s *AttributeContext) QUALIFIED_NAME() antlr.TerminalNode {
	return s.GetToken(PROV_NParserQUALIFIED_NAME, 0)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *PROV_NParser) Attribute() (localctx IAttributeContext) {
	this := p
	_ = this

	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PROV_NParserRULE_attribute)
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
		p.SetState(205)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PROV_NParserPREFX || _la == PROV_NParserQUALIFIED_NAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) TypedLiteral() ITypedLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedLiteralContext)
}

func (s *LiteralContext) ConvenienceNotation() IConvenienceNotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConvenienceNotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConvenienceNotationContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *PROV_NParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PROV_NParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.TypedLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.ConvenienceNotation()
		}

	}

	return localctx
}

// ITypedLiteralContext is an interface to support dynamic dispatch.
type ITypedLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypedLiteralContext differentiates from other interfaces.
	IsTypedLiteralContext()
}

type TypedLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedLiteralContext() *TypedLiteralContext {
	var p = new(TypedLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_typedLiteral
	return p
}

func (*TypedLiteralContext) IsTypedLiteralContext() {}

func NewTypedLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedLiteralContext {
	var p = new(TypedLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_typedLiteral

	return p
}

func (s *TypedLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(PROV_NParserSTRING_LITERAL, 0)
}

func (s *TypedLiteralContext) Datatype() IDatatypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatatypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *TypedLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterTypedLiteral(s)
	}
}

func (s *TypedLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitTypedLiteral(s)
	}
}

func (p *PROV_NParser) TypedLiteral() (localctx ITypedLiteralContext) {
	this := p
	_ = this

	localctx = NewTypedLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PROV_NParserRULE_typedLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PROV_NParserSTRING_LITERAL)
	}
	{
		p.SetState(212)
		p.Match(PROV_NParserT__9)
	}
	{
		p.SetState(213)
		p.Datatype()
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
	p.RuleIndex = PROV_NParserRULE_datatype
	return p
}

func (*DatatypeContext) IsDatatypeContext() {}

func NewDatatypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatatypeContext {
	var p = new(DatatypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_datatype

	return p
}

func (s *DatatypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatatypeContext) PREFX() antlr.TerminalNode {
	return s.GetToken(PROV_NParserPREFX, 0)
}

func (s *DatatypeContext) QUALIFIED_NAME() antlr.TerminalNode {
	return s.GetToken(PROV_NParserQUALIFIED_NAME, 0)
}

func (s *DatatypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatatypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatatypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterDatatype(s)
	}
}

func (s *DatatypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitDatatype(s)
	}
}

func (p *PROV_NParser) Datatype() (localctx IDatatypeContext) {
	this := p
	_ = this

	localctx = NewDatatypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PROV_NParserRULE_datatype)
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
		p.SetState(215)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PROV_NParserPREFX || _la == PROV_NParserQUALIFIED_NAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConvenienceNotationContext is an interface to support dynamic dispatch.
type IConvenienceNotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConvenienceNotationContext differentiates from other interfaces.
	IsConvenienceNotationContext()
}

type ConvenienceNotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConvenienceNotationContext() *ConvenienceNotationContext {
	var p = new(ConvenienceNotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_convenienceNotation
	return p
}

func (*ConvenienceNotationContext) IsConvenienceNotationContext() {}

func NewConvenienceNotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConvenienceNotationContext {
	var p = new(ConvenienceNotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_convenienceNotation

	return p
}

func (s *ConvenienceNotationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConvenienceNotationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(PROV_NParserSTRING_LITERAL, 0)
}

func (s *ConvenienceNotationContext) LANGTAG() antlr.TerminalNode {
	return s.GetToken(PROV_NParserLANGTAG, 0)
}

func (s *ConvenienceNotationContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PROV_NParserINT_LITERAL, 0)
}

func (s *ConvenienceNotationContext) QUALIFIED_NAME_LITERAL() antlr.TerminalNode {
	return s.GetToken(PROV_NParserQUALIFIED_NAME_LITERAL, 0)
}

func (s *ConvenienceNotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConvenienceNotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConvenienceNotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterConvenienceNotation(s)
	}
}

func (s *ConvenienceNotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitConvenienceNotation(s)
	}
}

func (p *PROV_NParser) ConvenienceNotation() (localctx IConvenienceNotationContext) {
	this := p
	_ = this

	localctx = NewConvenienceNotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PROV_NParserRULE_convenienceNotation)
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

	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(PROV_NParserSTRING_LITERAL)
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PROV_NParserLANGTAG {
			{
				p.SetState(218)
				p.Match(PROV_NParserLANGTAG)
			}

		}

	case PROV_NParserINT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Match(PROV_NParserINT_LITERAL)
		}

	case PROV_NParserQUALIFIED_NAME_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(222)
			p.Match(PROV_NParserQUALIFIED_NAME_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IActivityExpressionContext is an interface to support dynamic dispatch.
type IActivityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActivityExpressionContext differentiates from other interfaces.
	IsActivityExpressionContext()
}

type ActivityExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActivityExpressionContext() *ActivityExpressionContext {
	var p = new(ActivityExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_activityExpression
	return p
}

func (*ActivityExpressionContext) IsActivityExpressionContext() {}

func NewActivityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActivityExpressionContext {
	var p = new(ActivityExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_activityExpression

	return p
}

func (s *ActivityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ActivityExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ActivityExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *ActivityExpressionContext) AllTimeOrMarker() []ITimeOrMarkerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITimeOrMarkerContext)(nil)).Elem())
	var tst = make([]ITimeOrMarkerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITimeOrMarkerContext)
		}
	}

	return tst
}

func (s *ActivityExpressionContext) TimeOrMarker(i int) ITimeOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeOrMarkerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITimeOrMarkerContext)
}

func (s *ActivityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActivityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActivityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterActivityExpression(s)
	}
}

func (s *ActivityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitActivityExpression(s)
	}
}

func (p *PROV_NParser) ActivityExpression() (localctx IActivityExpressionContext) {
	this := p
	_ = this

	localctx = NewActivityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PROV_NParserRULE_activityExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(225)
		p.Match(PROV_NParserT__10)
	}
	{
		p.SetState(226)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(227)
		p.Identifier()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(228)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(229)
			p.TimeOrMarker()
		}
		{
			p.SetState(230)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(231)
			p.TimeOrMarker()
		}

	}
	{
		p.SetState(235)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(236)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// ITimeOrMarkerContext is an interface to support dynamic dispatch.
type ITimeOrMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeOrMarkerContext differentiates from other interfaces.
	IsTimeOrMarkerContext()
}

type TimeOrMarkerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeOrMarkerContext() *TimeOrMarkerContext {
	var p = new(TimeOrMarkerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_timeOrMarker
	return p
}

func (*TimeOrMarkerContext) IsTimeOrMarkerContext() {}

func NewTimeOrMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeOrMarkerContext {
	var p = new(TimeOrMarkerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_timeOrMarker

	return p
}

func (s *TimeOrMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeOrMarkerContext) Time() ITimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeContext)
}

func (s *TimeOrMarkerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PROV_NParserMINUS, 0)
}

func (s *TimeOrMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeOrMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeOrMarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterTimeOrMarker(s)
	}
}

func (s *TimeOrMarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitTimeOrMarker(s)
	}
}

func (p *PROV_NParser) TimeOrMarker() (localctx ITimeOrMarkerContext) {
	this := p
	_ = this

	localctx = NewTimeOrMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PROV_NParserRULE_timeOrMarker)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserDATETIME:
		{
			p.SetState(238)
			p.Time()
		}

	case PROV_NParserMINUS:
		{
			p.SetState(239)
			p.Match(PROV_NParserMINUS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITimeContext is an interface to support dynamic dispatch.
type ITimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeContext differentiates from other interfaces.
	IsTimeContext()
}

type TimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeContext() *TimeContext {
	var p = new(TimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_time
	return p
}

func (*TimeContext) IsTimeContext() {}

func NewTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeContext {
	var p = new(TimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_time

	return p
}

func (s *TimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(PROV_NParserDATETIME, 0)
}

func (s *TimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitTime(s)
	}
}

func (p *PROV_NParser) Time() (localctx ITimeContext) {
	this := p
	_ = this

	localctx = NewTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PROV_NParserRULE_time)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PROV_NParserDATETIME)
	}

	return localctx
}

// IGenerationExpressionContext is an interface to support dynamic dispatch.
type IGenerationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenerationExpressionContext differentiates from other interfaces.
	IsGenerationExpressionContext()
}

type GenerationExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenerationExpressionContext() *GenerationExpressionContext {
	var p = new(GenerationExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_generationExpression
	return p
}

func (*GenerationExpressionContext) IsGenerationExpressionContext() {}

func NewGenerationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenerationExpressionContext {
	var p = new(GenerationExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_generationExpression

	return p
}

func (s *GenerationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *GenerationExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *GenerationExpressionContext) EIdentifier() IEIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierContext)
}

func (s *GenerationExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *GenerationExpressionContext) AIdentifierOrMarker() IAIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierOrMarkerContext)
}

func (s *GenerationExpressionContext) TimeOrMarker() ITimeOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeOrMarkerContext)
}

func (s *GenerationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenerationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenerationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterGenerationExpression(s)
	}
}

func (s *GenerationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitGenerationExpression(s)
	}
}

func (p *PROV_NParser) GenerationExpression() (localctx IGenerationExpressionContext) {
	this := p
	_ = this

	localctx = NewGenerationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PROV_NParserRULE_generationExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PROV_NParserT__11)
	}
	{
		p.SetState(245)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(246)
		p.OptionalIdentifier()
	}
	{
		p.SetState(247)
		p.EIdentifier()
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(248)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(249)
			p.AIdentifierOrMarker()
		}
		{
			p.SetState(250)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(251)
			p.TimeOrMarker()
		}

	}
	{
		p.SetState(255)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(256)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IOptionalIdentifierContext is an interface to support dynamic dispatch.
type IOptionalIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionalIdentifierContext differentiates from other interfaces.
	IsOptionalIdentifierContext()
}

type OptionalIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalIdentifierContext() *OptionalIdentifierContext {
	var p = new(OptionalIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_optionalIdentifier
	return p
}

func (*OptionalIdentifierContext) IsOptionalIdentifierContext() {}

func NewOptionalIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalIdentifierContext {
	var p = new(OptionalIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_optionalIdentifier

	return p
}

func (s *OptionalIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalIdentifierContext) IdentifierOrMarker() IIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierOrMarkerContext)
}

func (s *OptionalIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterOptionalIdentifier(s)
	}
}

func (s *OptionalIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitOptionalIdentifier(s)
	}
}

func (p *PROV_NParser) OptionalIdentifier() (localctx IOptionalIdentifierContext) {
	this := p
	_ = this

	localctx = NewOptionalIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PROV_NParserRULE_optionalIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(258)
			p.IdentifierOrMarker()
		}
		{
			p.SetState(259)
			p.Match(PROV_NParserT__12)
		}

	}

	return localctx
}

// IIdentifierOrMarkerContext is an interface to support dynamic dispatch.
type IIdentifierOrMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierOrMarkerContext differentiates from other interfaces.
	IsIdentifierOrMarkerContext()
}

type IdentifierOrMarkerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierOrMarkerContext() *IdentifierOrMarkerContext {
	var p = new(IdentifierOrMarkerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_identifierOrMarker
	return p
}

func (*IdentifierOrMarkerContext) IsIdentifierOrMarkerContext() {}

func NewIdentifierOrMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierOrMarkerContext {
	var p = new(IdentifierOrMarkerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_identifierOrMarker

	return p
}

func (s *IdentifierOrMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierOrMarkerContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierOrMarkerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PROV_NParserMINUS, 0)
}

func (s *IdentifierOrMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierOrMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierOrMarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterIdentifierOrMarker(s)
	}
}

func (s *IdentifierOrMarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitIdentifierOrMarker(s)
	}
}

func (p *PROV_NParser) IdentifierOrMarker() (localctx IIdentifierOrMarkerContext) {
	this := p
	_ = this

	localctx = NewIdentifierOrMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PROV_NParserRULE_identifierOrMarker)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(265)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserPREFX, PROV_NParserQUALIFIED_NAME:
		{
			p.SetState(263)
			p.Identifier()
		}

	case PROV_NParserMINUS:
		{
			p.SetState(264)
			p.Match(PROV_NParserMINUS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEIdentifierContext is an interface to support dynamic dispatch.
type IEIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEIdentifierContext differentiates from other interfaces.
	IsEIdentifierContext()
}

type EIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEIdentifierContext() *EIdentifierContext {
	var p = new(EIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_eIdentifier
	return p
}

func (*EIdentifierContext) IsEIdentifierContext() {}

func NewEIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EIdentifierContext {
	var p = new(EIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_eIdentifier

	return p
}

func (s *EIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *EIdentifierContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *EIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterEIdentifier(s)
	}
}

func (s *EIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitEIdentifier(s)
	}
}

func (p *PROV_NParser) EIdentifier() (localctx IEIdentifierContext) {
	this := p
	_ = this

	localctx = NewEIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PROV_NParserRULE_eIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(267)
		p.Identifier()
	}

	return localctx
}

// IEIdentifierOrMarkerContext is an interface to support dynamic dispatch.
type IEIdentifierOrMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEIdentifierOrMarkerContext differentiates from other interfaces.
	IsEIdentifierOrMarkerContext()
}

type EIdentifierOrMarkerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEIdentifierOrMarkerContext() *EIdentifierOrMarkerContext {
	var p = new(EIdentifierOrMarkerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_eIdentifierOrMarker
	return p
}

func (*EIdentifierOrMarkerContext) IsEIdentifierOrMarkerContext() {}

func NewEIdentifierOrMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EIdentifierOrMarkerContext {
	var p = new(EIdentifierOrMarkerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_eIdentifierOrMarker

	return p
}

func (s *EIdentifierOrMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *EIdentifierOrMarkerContext) EIdentifier() IEIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierContext)
}

func (s *EIdentifierOrMarkerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PROV_NParserMINUS, 0)
}

func (s *EIdentifierOrMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIdentifierOrMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EIdentifierOrMarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterEIdentifierOrMarker(s)
	}
}

func (s *EIdentifierOrMarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitEIdentifierOrMarker(s)
	}
}

func (p *PROV_NParser) EIdentifierOrMarker() (localctx IEIdentifierOrMarkerContext) {
	this := p
	_ = this

	localctx = NewEIdentifierOrMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PROV_NParserRULE_eIdentifierOrMarker)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(271)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserPREFX, PROV_NParserQUALIFIED_NAME:
		{
			p.SetState(269)
			p.EIdentifier()
		}

	case PROV_NParserMINUS:
		{
			p.SetState(270)
			p.Match(PROV_NParserMINUS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAIdentifierOrMarkerContext is an interface to support dynamic dispatch.
type IAIdentifierOrMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAIdentifierOrMarkerContext differentiates from other interfaces.
	IsAIdentifierOrMarkerContext()
}

type AIdentifierOrMarkerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAIdentifierOrMarkerContext() *AIdentifierOrMarkerContext {
	var p = new(AIdentifierOrMarkerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_aIdentifierOrMarker
	return p
}

func (*AIdentifierOrMarkerContext) IsAIdentifierOrMarkerContext() {}

func NewAIdentifierOrMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AIdentifierOrMarkerContext {
	var p = new(AIdentifierOrMarkerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_aIdentifierOrMarker

	return p
}

func (s *AIdentifierOrMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *AIdentifierOrMarkerContext) AIdentifier() IAIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierContext)
}

func (s *AIdentifierOrMarkerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PROV_NParserMINUS, 0)
}

func (s *AIdentifierOrMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AIdentifierOrMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AIdentifierOrMarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAIdentifierOrMarker(s)
	}
}

func (s *AIdentifierOrMarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAIdentifierOrMarker(s)
	}
}

func (p *PROV_NParser) AIdentifierOrMarker() (localctx IAIdentifierOrMarkerContext) {
	this := p
	_ = this

	localctx = NewAIdentifierOrMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PROV_NParserRULE_aIdentifierOrMarker)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserPREFX, PROV_NParserQUALIFIED_NAME:
		{
			p.SetState(273)
			p.AIdentifier()
		}

	case PROV_NParserMINUS:
		{
			p.SetState(274)
			p.Match(PROV_NParserMINUS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAIdentifierContext is an interface to support dynamic dispatch.
type IAIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAIdentifierContext differentiates from other interfaces.
	IsAIdentifierContext()
}

type AIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAIdentifierContext() *AIdentifierContext {
	var p = new(AIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_aIdentifier
	return p
}

func (*AIdentifierContext) IsAIdentifierContext() {}

func NewAIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AIdentifierContext {
	var p = new(AIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_aIdentifier

	return p
}

func (s *AIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *AIdentifierContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAIdentifier(s)
	}
}

func (s *AIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAIdentifier(s)
	}
}

func (p *PROV_NParser) AIdentifier() (localctx IAIdentifierContext) {
	this := p
	_ = this

	localctx = NewAIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PROV_NParserRULE_aIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Identifier()
	}

	return localctx
}

// IAgIdentifierOrMarkerContext is an interface to support dynamic dispatch.
type IAgIdentifierOrMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAgIdentifierOrMarkerContext differentiates from other interfaces.
	IsAgIdentifierOrMarkerContext()
}

type AgIdentifierOrMarkerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgIdentifierOrMarkerContext() *AgIdentifierOrMarkerContext {
	var p = new(AgIdentifierOrMarkerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_agIdentifierOrMarker
	return p
}

func (*AgIdentifierOrMarkerContext) IsAgIdentifierOrMarkerContext() {}

func NewAgIdentifierOrMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgIdentifierOrMarkerContext {
	var p = new(AgIdentifierOrMarkerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_agIdentifierOrMarker

	return p
}

func (s *AgIdentifierOrMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *AgIdentifierOrMarkerContext) AgIdentifier() IAgIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAgIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAgIdentifierContext)
}

func (s *AgIdentifierOrMarkerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PROV_NParserMINUS, 0)
}

func (s *AgIdentifierOrMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgIdentifierOrMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgIdentifierOrMarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAgIdentifierOrMarker(s)
	}
}

func (s *AgIdentifierOrMarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAgIdentifierOrMarker(s)
	}
}

func (p *PROV_NParser) AgIdentifierOrMarker() (localctx IAgIdentifierOrMarkerContext) {
	this := p
	_ = this

	localctx = NewAgIdentifierOrMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PROV_NParserRULE_agIdentifierOrMarker)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(281)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserPREFX, PROV_NParserQUALIFIED_NAME:
		{
			p.SetState(279)
			p.AgIdentifier()
		}

	case PROV_NParserMINUS:
		{
			p.SetState(280)
			p.Match(PROV_NParserMINUS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAgIdentifierContext is an interface to support dynamic dispatch.
type IAgIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAgIdentifierContext differentiates from other interfaces.
	IsAgIdentifierContext()
}

type AgIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgIdentifierContext() *AgIdentifierContext {
	var p = new(AgIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_agIdentifier
	return p
}

func (*AgIdentifierContext) IsAgIdentifierContext() {}

func NewAgIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgIdentifierContext {
	var p = new(AgIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_agIdentifier

	return p
}

func (s *AgIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *AgIdentifierContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AgIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAgIdentifier(s)
	}
}

func (s *AgIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAgIdentifier(s)
	}
}

func (p *PROV_NParser) AgIdentifier() (localctx IAgIdentifierContext) {
	this := p
	_ = this

	localctx = NewAgIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PROV_NParserRULE_agIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Identifier()
	}

	return localctx
}

// ICIdentifierContext is an interface to support dynamic dispatch.
type ICIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCIdentifierContext differentiates from other interfaces.
	IsCIdentifierContext()
}

type CIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCIdentifierContext() *CIdentifierContext {
	var p = new(CIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_cIdentifier
	return p
}

func (*CIdentifierContext) IsCIdentifierContext() {}

func NewCIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CIdentifierContext {
	var p = new(CIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_cIdentifier

	return p
}

func (s *CIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *CIdentifierContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterCIdentifier(s)
	}
}

func (s *CIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitCIdentifier(s)
	}
}

func (p *PROV_NParser) CIdentifier() (localctx ICIdentifierContext) {
	this := p
	_ = this

	localctx = NewCIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PROV_NParserRULE_cIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Identifier()
	}

	return localctx
}

// IGIdentifierContext is an interface to support dynamic dispatch.
type IGIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGIdentifierContext differentiates from other interfaces.
	IsGIdentifierContext()
}

type GIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGIdentifierContext() *GIdentifierContext {
	var p = new(GIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_gIdentifier
	return p
}

func (*GIdentifierContext) IsGIdentifierContext() {}

func NewGIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GIdentifierContext {
	var p = new(GIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_gIdentifier

	return p
}

func (s *GIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *GIdentifierContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *GIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterGIdentifier(s)
	}
}

func (s *GIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitGIdentifier(s)
	}
}

func (p *PROV_NParser) GIdentifier() (localctx IGIdentifierContext) {
	this := p
	_ = this

	localctx = NewGIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PROV_NParserRULE_gIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Identifier()
	}

	return localctx
}

// IGIdentifierOrMarkerContext is an interface to support dynamic dispatch.
type IGIdentifierOrMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGIdentifierOrMarkerContext differentiates from other interfaces.
	IsGIdentifierOrMarkerContext()
}

type GIdentifierOrMarkerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGIdentifierOrMarkerContext() *GIdentifierOrMarkerContext {
	var p = new(GIdentifierOrMarkerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_gIdentifierOrMarker
	return p
}

func (*GIdentifierOrMarkerContext) IsGIdentifierOrMarkerContext() {}

func NewGIdentifierOrMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GIdentifierOrMarkerContext {
	var p = new(GIdentifierOrMarkerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_gIdentifierOrMarker

	return p
}

func (s *GIdentifierOrMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *GIdentifierOrMarkerContext) GIdentifier() IGIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGIdentifierContext)
}

func (s *GIdentifierOrMarkerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PROV_NParserMINUS, 0)
}

func (s *GIdentifierOrMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GIdentifierOrMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GIdentifierOrMarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterGIdentifierOrMarker(s)
	}
}

func (s *GIdentifierOrMarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitGIdentifierOrMarker(s)
	}
}

func (p *PROV_NParser) GIdentifierOrMarker() (localctx IGIdentifierOrMarkerContext) {
	this := p
	_ = this

	localctx = NewGIdentifierOrMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PROV_NParserRULE_gIdentifierOrMarker)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(291)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserPREFX, PROV_NParserQUALIFIED_NAME:
		{
			p.SetState(289)
			p.GIdentifier()
		}

	case PROV_NParserMINUS:
		{
			p.SetState(290)
			p.Match(PROV_NParserMINUS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUIdentifierContext is an interface to support dynamic dispatch.
type IUIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUIdentifierContext differentiates from other interfaces.
	IsUIdentifierContext()
}

type UIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUIdentifierContext() *UIdentifierContext {
	var p = new(UIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_uIdentifier
	return p
}

func (*UIdentifierContext) IsUIdentifierContext() {}

func NewUIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UIdentifierContext {
	var p = new(UIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_uIdentifier

	return p
}

func (s *UIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *UIdentifierContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *UIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterUIdentifier(s)
	}
}

func (s *UIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitUIdentifier(s)
	}
}

func (p *PROV_NParser) UIdentifier() (localctx IUIdentifierContext) {
	this := p
	_ = this

	localctx = NewUIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PROV_NParserRULE_uIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Identifier()
	}

	return localctx
}

// IUIdentifierOrMarkerContext is an interface to support dynamic dispatch.
type IUIdentifierOrMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUIdentifierOrMarkerContext differentiates from other interfaces.
	IsUIdentifierOrMarkerContext()
}

type UIdentifierOrMarkerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUIdentifierOrMarkerContext() *UIdentifierOrMarkerContext {
	var p = new(UIdentifierOrMarkerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_uIdentifierOrMarker
	return p
}

func (*UIdentifierOrMarkerContext) IsUIdentifierOrMarkerContext() {}

func NewUIdentifierOrMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UIdentifierOrMarkerContext {
	var p = new(UIdentifierOrMarkerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_uIdentifierOrMarker

	return p
}

func (s *UIdentifierOrMarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *UIdentifierOrMarkerContext) UIdentifier() IUIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUIdentifierContext)
}

func (s *UIdentifierOrMarkerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PROV_NParserMINUS, 0)
}

func (s *UIdentifierOrMarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UIdentifierOrMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UIdentifierOrMarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterUIdentifierOrMarker(s)
	}
}

func (s *UIdentifierOrMarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitUIdentifierOrMarker(s)
	}
}

func (p *PROV_NParser) UIdentifierOrMarker() (localctx IUIdentifierOrMarkerContext) {
	this := p
	_ = this

	localctx = NewUIdentifierOrMarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PROV_NParserRULE_uIdentifierOrMarker)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserPREFX, PROV_NParserQUALIFIED_NAME:
		{
			p.SetState(295)
			p.UIdentifier()
		}

	case PROV_NParserMINUS:
		{
			p.SetState(296)
			p.Match(PROV_NParserMINUS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUsageExpressionContext is an interface to support dynamic dispatch.
type IUsageExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUsageExpressionContext differentiates from other interfaces.
	IsUsageExpressionContext()
}

type UsageExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUsageExpressionContext() *UsageExpressionContext {
	var p = new(UsageExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_usageExpression
	return p
}

func (*UsageExpressionContext) IsUsageExpressionContext() {}

func NewUsageExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UsageExpressionContext {
	var p = new(UsageExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_usageExpression

	return p
}

func (s *UsageExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UsageExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *UsageExpressionContext) AIdentifier() IAIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierContext)
}

func (s *UsageExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *UsageExpressionContext) EIdentifierOrMarker() IEIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierOrMarkerContext)
}

func (s *UsageExpressionContext) TimeOrMarker() ITimeOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeOrMarkerContext)
}

func (s *UsageExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UsageExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UsageExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterUsageExpression(s)
	}
}

func (s *UsageExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitUsageExpression(s)
	}
}

func (p *PROV_NParser) UsageExpression() (localctx IUsageExpressionContext) {
	this := p
	_ = this

	localctx = NewUsageExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PROV_NParserRULE_usageExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PROV_NParserT__13)
	}
	{
		p.SetState(300)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(301)
		p.OptionalIdentifier()
	}
	{
		p.SetState(302)
		p.AIdentifier()
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(303)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(304)
			p.EIdentifierOrMarker()
		}
		{
			p.SetState(305)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(306)
			p.TimeOrMarker()
		}

	}
	{
		p.SetState(310)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(311)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IStartExpressionContext is an interface to support dynamic dispatch.
type IStartExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartExpressionContext differentiates from other interfaces.
	IsStartExpressionContext()
}

type StartExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartExpressionContext() *StartExpressionContext {
	var p = new(StartExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_startExpression
	return p
}

func (*StartExpressionContext) IsStartExpressionContext() {}

func NewStartExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartExpressionContext {
	var p = new(StartExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_startExpression

	return p
}

func (s *StartExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *StartExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *StartExpressionContext) AIdentifier() IAIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierContext)
}

func (s *StartExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *StartExpressionContext) EIdentifierOrMarker() IEIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierOrMarkerContext)
}

func (s *StartExpressionContext) AIdentifierOrMarker() IAIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierOrMarkerContext)
}

func (s *StartExpressionContext) TimeOrMarker() ITimeOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeOrMarkerContext)
}

func (s *StartExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterStartExpression(s)
	}
}

func (s *StartExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitStartExpression(s)
	}
}

func (p *PROV_NParser) StartExpression() (localctx IStartExpressionContext) {
	this := p
	_ = this

	localctx = NewStartExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PROV_NParserRULE_startExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PROV_NParserT__14)
	}
	{
		p.SetState(314)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(315)
		p.OptionalIdentifier()
	}
	{
		p.SetState(316)
		p.AIdentifier()
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(317)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(318)
			p.EIdentifierOrMarker()
		}
		{
			p.SetState(319)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(320)
			p.AIdentifierOrMarker()
		}
		{
			p.SetState(321)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(322)
			p.TimeOrMarker()
		}

	}
	{
		p.SetState(326)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(327)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IEndExpressionContext is an interface to support dynamic dispatch.
type IEndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndExpressionContext differentiates from other interfaces.
	IsEndExpressionContext()
}

type EndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndExpressionContext() *EndExpressionContext {
	var p = new(EndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_endExpression
	return p
}

func (*EndExpressionContext) IsEndExpressionContext() {}

func NewEndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndExpressionContext {
	var p = new(EndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_endExpression

	return p
}

func (s *EndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EndExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *EndExpressionContext) AIdentifier() IAIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierContext)
}

func (s *EndExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *EndExpressionContext) EIdentifierOrMarker() IEIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierOrMarkerContext)
}

func (s *EndExpressionContext) AIdentifierOrMarker() IAIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierOrMarkerContext)
}

func (s *EndExpressionContext) TimeOrMarker() ITimeOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeOrMarkerContext)
}

func (s *EndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterEndExpression(s)
	}
}

func (s *EndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitEndExpression(s)
	}
}

func (p *PROV_NParser) EndExpression() (localctx IEndExpressionContext) {
	this := p
	_ = this

	localctx = NewEndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PROV_NParserRULE_endExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(329)
		p.Match(PROV_NParserT__15)
	}
	{
		p.SetState(330)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(331)
		p.OptionalIdentifier()
	}
	{
		p.SetState(332)
		p.AIdentifier()
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(333)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(334)
			p.EIdentifierOrMarker()
		}
		{
			p.SetState(335)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(336)
			p.AIdentifierOrMarker()
		}
		{
			p.SetState(337)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(338)
			p.TimeOrMarker()
		}

	}
	{
		p.SetState(342)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(343)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IInvalidationExpressionContext is an interface to support dynamic dispatch.
type IInvalidationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvalidationExpressionContext differentiates from other interfaces.
	IsInvalidationExpressionContext()
}

type InvalidationExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvalidationExpressionContext() *InvalidationExpressionContext {
	var p = new(InvalidationExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_invalidationExpression
	return p
}

func (*InvalidationExpressionContext) IsInvalidationExpressionContext() {}

func NewInvalidationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvalidationExpressionContext {
	var p = new(InvalidationExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_invalidationExpression

	return p
}

func (s *InvalidationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InvalidationExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *InvalidationExpressionContext) EIdentifier() IEIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierContext)
}

func (s *InvalidationExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *InvalidationExpressionContext) AIdentifierOrMarker() IAIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierOrMarkerContext)
}

func (s *InvalidationExpressionContext) TimeOrMarker() ITimeOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeOrMarkerContext)
}

func (s *InvalidationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvalidationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvalidationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterInvalidationExpression(s)
	}
}

func (s *InvalidationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitInvalidationExpression(s)
	}
}

func (p *PROV_NParser) InvalidationExpression() (localctx IInvalidationExpressionContext) {
	this := p
	_ = this

	localctx = NewInvalidationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PROV_NParserRULE_invalidationExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(345)
		p.Match(PROV_NParserT__16)
	}
	{
		p.SetState(346)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(347)
		p.OptionalIdentifier()
	}
	{
		p.SetState(348)
		p.EIdentifier()
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(349)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(350)
			p.AIdentifierOrMarker()
		}
		{
			p.SetState(351)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(352)
			p.TimeOrMarker()
		}

	}
	{
		p.SetState(356)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(357)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// ICommunicationExpressionContext is an interface to support dynamic dispatch.
type ICommunicationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommunicationExpressionContext differentiates from other interfaces.
	IsCommunicationExpressionContext()
}

type CommunicationExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommunicationExpressionContext() *CommunicationExpressionContext {
	var p = new(CommunicationExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_communicationExpression
	return p
}

func (*CommunicationExpressionContext) IsCommunicationExpressionContext() {}

func NewCommunicationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommunicationExpressionContext {
	var p = new(CommunicationExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_communicationExpression

	return p
}

func (s *CommunicationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CommunicationExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *CommunicationExpressionContext) AllAIdentifier() []IAIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAIdentifierContext)(nil)).Elem())
	var tst = make([]IAIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAIdentifierContext)
		}
	}

	return tst
}

func (s *CommunicationExpressionContext) AIdentifier(i int) IAIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierContext)
}

func (s *CommunicationExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *CommunicationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommunicationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommunicationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterCommunicationExpression(s)
	}
}

func (s *CommunicationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitCommunicationExpression(s)
	}
}

func (p *PROV_NParser) CommunicationExpression() (localctx ICommunicationExpressionContext) {
	this := p
	_ = this

	localctx = NewCommunicationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PROV_NParserRULE_communicationExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PROV_NParserT__17)
	}
	{
		p.SetState(360)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(361)
		p.OptionalIdentifier()
	}
	{
		p.SetState(362)
		p.AIdentifier()
	}
	{
		p.SetState(363)
		p.Match(PROV_NParserT__5)
	}
	{
		p.SetState(364)
		p.AIdentifier()
	}
	{
		p.SetState(365)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(366)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IAgentExpressionContext is an interface to support dynamic dispatch.
type IAgentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAgentExpressionContext differentiates from other interfaces.
	IsAgentExpressionContext()
}

type AgentExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgentExpressionContext() *AgentExpressionContext {
	var p = new(AgentExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_agentExpression
	return p
}

func (*AgentExpressionContext) IsAgentExpressionContext() {}

func NewAgentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgentExpressionContext {
	var p = new(AgentExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_agentExpression

	return p
}

func (s *AgentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AgentExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AgentExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *AgentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAgentExpression(s)
	}
}

func (s *AgentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAgentExpression(s)
	}
}

func (p *PROV_NParser) AgentExpression() (localctx IAgentExpressionContext) {
	this := p
	_ = this

	localctx = NewAgentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PROV_NParserRULE_agentExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(368)
		p.Match(PROV_NParserT__18)
	}
	{
		p.SetState(369)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(370)
		p.Identifier()
	}
	{
		p.SetState(371)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(372)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IAssociationExpressionContext is an interface to support dynamic dispatch.
type IAssociationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssociationExpressionContext differentiates from other interfaces.
	IsAssociationExpressionContext()
}

type AssociationExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssociationExpressionContext() *AssociationExpressionContext {
	var p = new(AssociationExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_associationExpression
	return p
}

func (*AssociationExpressionContext) IsAssociationExpressionContext() {}

func NewAssociationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssociationExpressionContext {
	var p = new(AssociationExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_associationExpression

	return p
}

func (s *AssociationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssociationExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *AssociationExpressionContext) AIdentifier() IAIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierContext)
}

func (s *AssociationExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *AssociationExpressionContext) AgIdentifierOrMarker() IAgIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAgIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAgIdentifierOrMarkerContext)
}

func (s *AssociationExpressionContext) EIdentifierOrMarker() IEIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierOrMarkerContext)
}

func (s *AssociationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssociationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssociationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAssociationExpression(s)
	}
}

func (s *AssociationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAssociationExpression(s)
	}
}

func (p *PROV_NParser) AssociationExpression() (localctx IAssociationExpressionContext) {
	this := p
	_ = this

	localctx = NewAssociationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PROV_NParserRULE_associationExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(374)
		p.Match(PROV_NParserT__19)
	}
	{
		p.SetState(375)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(376)
		p.OptionalIdentifier()
	}
	{
		p.SetState(377)
		p.AIdentifier()
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(378)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(379)
			p.AgIdentifierOrMarker()
		}
		{
			p.SetState(380)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(381)
			p.EIdentifierOrMarker()
		}

	}
	{
		p.SetState(385)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(386)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IAttributionExpressionContext is an interface to support dynamic dispatch.
type IAttributionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributionExpressionContext differentiates from other interfaces.
	IsAttributionExpressionContext()
}

type AttributionExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributionExpressionContext() *AttributionExpressionContext {
	var p = new(AttributionExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_attributionExpression
	return p
}

func (*AttributionExpressionContext) IsAttributionExpressionContext() {}

func NewAttributionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributionExpressionContext {
	var p = new(AttributionExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_attributionExpression

	return p
}

func (s *AttributionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributionExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *AttributionExpressionContext) EIdentifier() IEIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierContext)
}

func (s *AttributionExpressionContext) AgIdentifier() IAgIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAgIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAgIdentifierContext)
}

func (s *AttributionExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *AttributionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAttributionExpression(s)
	}
}

func (s *AttributionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAttributionExpression(s)
	}
}

func (p *PROV_NParser) AttributionExpression() (localctx IAttributionExpressionContext) {
	this := p
	_ = this

	localctx = NewAttributionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, PROV_NParserRULE_attributionExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PROV_NParserT__20)
	}
	{
		p.SetState(389)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(390)
		p.OptionalIdentifier()
	}
	{
		p.SetState(391)
		p.EIdentifier()
	}
	{
		p.SetState(392)
		p.Match(PROV_NParserT__5)
	}
	{
		p.SetState(393)
		p.AgIdentifier()
	}
	{
		p.SetState(394)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(395)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IDelegationExpressionContext is an interface to support dynamic dispatch.
type IDelegationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelegationExpressionContext differentiates from other interfaces.
	IsDelegationExpressionContext()
}

type DelegationExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelegationExpressionContext() *DelegationExpressionContext {
	var p = new(DelegationExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_delegationExpression
	return p
}

func (*DelegationExpressionContext) IsDelegationExpressionContext() {}

func NewDelegationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelegationExpressionContext {
	var p = new(DelegationExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_delegationExpression

	return p
}

func (s *DelegationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *DelegationExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *DelegationExpressionContext) AllAgIdentifier() []IAgIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAgIdentifierContext)(nil)).Elem())
	var tst = make([]IAgIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAgIdentifierContext)
		}
	}

	return tst
}

func (s *DelegationExpressionContext) AgIdentifier(i int) IAgIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAgIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAgIdentifierContext)
}

func (s *DelegationExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *DelegationExpressionContext) AIdentifierOrMarker() IAIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierOrMarkerContext)
}

func (s *DelegationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelegationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelegationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterDelegationExpression(s)
	}
}

func (s *DelegationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitDelegationExpression(s)
	}
}

func (p *PROV_NParser) DelegationExpression() (localctx IDelegationExpressionContext) {
	this := p
	_ = this

	localctx = NewDelegationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, PROV_NParserRULE_delegationExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PROV_NParserT__21)
	}
	{
		p.SetState(398)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(399)
		p.OptionalIdentifier()
	}
	{
		p.SetState(400)
		p.AgIdentifier()
	}
	{
		p.SetState(401)
		p.Match(PROV_NParserT__5)
	}
	{
		p.SetState(402)
		p.AgIdentifier()
	}
	p.SetState(405)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(403)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(404)
			p.AIdentifierOrMarker()
		}

	}
	{
		p.SetState(407)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(408)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IDerivationExpressionContext is an interface to support dynamic dispatch.
type IDerivationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDerivationExpressionContext differentiates from other interfaces.
	IsDerivationExpressionContext()
}

type DerivationExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDerivationExpressionContext() *DerivationExpressionContext {
	var p = new(DerivationExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_derivationExpression
	return p
}

func (*DerivationExpressionContext) IsDerivationExpressionContext() {}

func NewDerivationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DerivationExpressionContext {
	var p = new(DerivationExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_derivationExpression

	return p
}

func (s *DerivationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *DerivationExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *DerivationExpressionContext) AllEIdentifier() []IEIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem())
	var tst = make([]IEIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEIdentifierContext)
		}
	}

	return tst
}

func (s *DerivationExpressionContext) EIdentifier(i int) IEIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierContext)
}

func (s *DerivationExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *DerivationExpressionContext) AIdentifierOrMarker() IAIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAIdentifierOrMarkerContext)
}

func (s *DerivationExpressionContext) GIdentifierOrMarker() IGIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGIdentifierOrMarkerContext)
}

func (s *DerivationExpressionContext) UIdentifierOrMarker() IUIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUIdentifierOrMarkerContext)
}

func (s *DerivationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerivationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DerivationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterDerivationExpression(s)
	}
}

func (s *DerivationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitDerivationExpression(s)
	}
}

func (p *PROV_NParser) DerivationExpression() (localctx IDerivationExpressionContext) {
	this := p
	_ = this

	localctx = NewDerivationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, PROV_NParserRULE_derivationExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(410)
		p.Match(PROV_NParserT__22)
	}
	{
		p.SetState(411)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(412)
		p.OptionalIdentifier()
	}
	{
		p.SetState(413)
		p.EIdentifier()
	}
	{
		p.SetState(414)
		p.Match(PROV_NParserT__5)
	}
	{
		p.SetState(415)
		p.EIdentifier()
	}
	p.SetState(423)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(416)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(417)
			p.AIdentifierOrMarker()
		}
		{
			p.SetState(418)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(419)
			p.GIdentifierOrMarker()
		}
		{
			p.SetState(420)
			p.Match(PROV_NParserT__5)
		}
		{
			p.SetState(421)
			p.UIdentifierOrMarker()
		}

	}
	{
		p.SetState(425)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(426)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IInfluenceExpressionContext is an interface to support dynamic dispatch.
type IInfluenceExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfluenceExpressionContext differentiates from other interfaces.
	IsInfluenceExpressionContext()
}

type InfluenceExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfluenceExpressionContext() *InfluenceExpressionContext {
	var p = new(InfluenceExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_influenceExpression
	return p
}

func (*InfluenceExpressionContext) IsInfluenceExpressionContext() {}

func NewInfluenceExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfluenceExpressionContext {
	var p = new(InfluenceExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_influenceExpression

	return p
}

func (s *InfluenceExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InfluenceExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *InfluenceExpressionContext) AllEIdentifier() []IEIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem())
	var tst = make([]IEIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEIdentifierContext)
		}
	}

	return tst
}

func (s *InfluenceExpressionContext) EIdentifier(i int) IEIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierContext)
}

func (s *InfluenceExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *InfluenceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfluenceExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfluenceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterInfluenceExpression(s)
	}
}

func (s *InfluenceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitInfluenceExpression(s)
	}
}

func (p *PROV_NParser) InfluenceExpression() (localctx IInfluenceExpressionContext) {
	this := p
	_ = this

	localctx = NewInfluenceExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, PROV_NParserRULE_influenceExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(428)
		p.Match(PROV_NParserT__23)
	}
	{
		p.SetState(429)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(430)
		p.OptionalIdentifier()
	}
	{
		p.SetState(431)
		p.EIdentifier()
	}
	{
		p.SetState(432)
		p.Match(PROV_NParserT__5)
	}
	{
		p.SetState(433)
		p.EIdentifier()
	}
	{
		p.SetState(434)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(435)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IAlternateExpressionContext is an interface to support dynamic dispatch.
type IAlternateExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlternateExpressionContext differentiates from other interfaces.
	IsAlternateExpressionContext()
}

type AlternateExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlternateExpressionContext() *AlternateExpressionContext {
	var p = new(AlternateExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_alternateExpression
	return p
}

func (*AlternateExpressionContext) IsAlternateExpressionContext() {}

func NewAlternateExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlternateExpressionContext {
	var p = new(AlternateExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_alternateExpression

	return p
}

func (s *AlternateExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AlternateExpressionContext) AllEIdentifier() []IEIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem())
	var tst = make([]IEIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEIdentifierContext)
		}
	}

	return tst
}

func (s *AlternateExpressionContext) EIdentifier(i int) IEIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierContext)
}

func (s *AlternateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlternateExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlternateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterAlternateExpression(s)
	}
}

func (s *AlternateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitAlternateExpression(s)
	}
}

func (p *PROV_NParser) AlternateExpression() (localctx IAlternateExpressionContext) {
	this := p
	_ = this

	localctx = NewAlternateExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, PROV_NParserRULE_alternateExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(437)
		p.Match(PROV_NParserT__24)
	}
	{
		p.SetState(438)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(439)
		p.EIdentifier()
	}
	{
		p.SetState(440)
		p.Match(PROV_NParserT__5)
	}
	{
		p.SetState(441)
		p.EIdentifier()
	}
	{
		p.SetState(442)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// ISpecializationExpressionContext is an interface to support dynamic dispatch.
type ISpecializationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecializationExpressionContext differentiates from other interfaces.
	IsSpecializationExpressionContext()
}

type SpecializationExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecializationExpressionContext() *SpecializationExpressionContext {
	var p = new(SpecializationExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_specializationExpression
	return p
}

func (*SpecializationExpressionContext) IsSpecializationExpressionContext() {}

func NewSpecializationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecializationExpressionContext {
	var p = new(SpecializationExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_specializationExpression

	return p
}

func (s *SpecializationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecializationExpressionContext) AllEIdentifier() []IEIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem())
	var tst = make([]IEIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEIdentifierContext)
		}
	}

	return tst
}

func (s *SpecializationExpressionContext) EIdentifier(i int) IEIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierContext)
}

func (s *SpecializationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecializationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecializationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterSpecializationExpression(s)
	}
}

func (s *SpecializationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitSpecializationExpression(s)
	}
}

func (p *PROV_NParser) SpecializationExpression() (localctx ISpecializationExpressionContext) {
	this := p
	_ = this

	localctx = NewSpecializationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, PROV_NParserRULE_specializationExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(444)
		p.Match(PROV_NParserT__25)
	}
	{
		p.SetState(445)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(446)
		p.EIdentifier()
	}
	{
		p.SetState(447)
		p.Match(PROV_NParserT__5)
	}
	{
		p.SetState(448)
		p.EIdentifier()
	}
	{
		p.SetState(449)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IMembershipExpressionContext is an interface to support dynamic dispatch.
type IMembershipExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMembershipExpressionContext differentiates from other interfaces.
	IsMembershipExpressionContext()
}

type MembershipExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMembershipExpressionContext() *MembershipExpressionContext {
	var p = new(MembershipExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_membershipExpression
	return p
}

func (*MembershipExpressionContext) IsMembershipExpressionContext() {}

func NewMembershipExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MembershipExpressionContext {
	var p = new(MembershipExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_membershipExpression

	return p
}

func (s *MembershipExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MembershipExpressionContext) CIdentifier() ICIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICIdentifierContext)
}

func (s *MembershipExpressionContext) EIdentifier() IEIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEIdentifierContext)
}

func (s *MembershipExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembershipExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterMembershipExpression(s)
	}
}

func (s *MembershipExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitMembershipExpression(s)
	}
}

func (p *PROV_NParser) MembershipExpression() (localctx IMembershipExpressionContext) {
	this := p
	_ = this

	localctx = NewMembershipExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, PROV_NParserRULE_membershipExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(451)
		p.Match(PROV_NParserT__26)
	}
	{
		p.SetState(452)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(453)
		p.CIdentifier()
	}
	{
		p.SetState(454)
		p.Match(PROV_NParserT__5)
	}
	{
		p.SetState(455)
		p.EIdentifier()
	}
	{
		p.SetState(456)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IExtensibilityExpressionContext is an interface to support dynamic dispatch.
type IExtensibilityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtensibilityExpressionContext differentiates from other interfaces.
	IsExtensibilityExpressionContext()
}

type ExtensibilityExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensibilityExpressionContext() *ExtensibilityExpressionContext {
	var p = new(ExtensibilityExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_extensibilityExpression
	return p
}

func (*ExtensibilityExpressionContext) IsExtensibilityExpressionContext() {}

func NewExtensibilityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensibilityExpressionContext {
	var p = new(ExtensibilityExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_extensibilityExpression

	return p
}

func (s *ExtensibilityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensibilityExpressionContext) OptionalIdentifier() IOptionalIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalIdentifierContext)
}

func (s *ExtensibilityExpressionContext) AllExtensibilityArgument() []IExtensibilityArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExtensibilityArgumentContext)(nil)).Elem())
	var tst = make([]IExtensibilityArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExtensibilityArgumentContext)
		}
	}

	return tst
}

func (s *ExtensibilityExpressionContext) ExtensibilityArgument(i int) IExtensibilityArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtensibilityArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExtensibilityArgumentContext)
}

func (s *ExtensibilityExpressionContext) OptionalAttributeValuePairs() IOptionalAttributeValuePairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAttributeValuePairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAttributeValuePairsContext)
}

func (s *ExtensibilityExpressionContext) PREFX() antlr.TerminalNode {
	return s.GetToken(PROV_NParserPREFX, 0)
}

func (s *ExtensibilityExpressionContext) QUALIFIED_NAME() antlr.TerminalNode {
	return s.GetToken(PROV_NParserQUALIFIED_NAME, 0)
}

func (s *ExtensibilityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensibilityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensibilityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterExtensibilityExpression(s)
	}
}

func (s *ExtensibilityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitExtensibilityExpression(s)
	}
}

func (p *PROV_NParser) ExtensibilityExpression() (localctx IExtensibilityExpressionContext) {
	this := p
	_ = this

	localctx = NewExtensibilityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, PROV_NParserRULE_extensibilityExpression)
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
		p.SetState(458)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PROV_NParserPREFX || _la == PROV_NParserQUALIFIED_NAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(459)
		p.Match(PROV_NParserT__3)
	}
	{
		p.SetState(460)
		p.OptionalIdentifier()
	}
	{
		p.SetState(461)
		p.ExtensibilityArgument()
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(462)
				p.Match(PROV_NParserT__5)
			}
			{
				p.SetState(463)
				p.ExtensibilityArgument()
			}

		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
	}
	{
		p.SetState(469)
		p.OptionalAttributeValuePairs()
	}
	{
		p.SetState(470)
		p.Match(PROV_NParserT__4)
	}

	return localctx
}

// IExtensibilityArgumentContext is an interface to support dynamic dispatch.
type IExtensibilityArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtensibilityArgumentContext differentiates from other interfaces.
	IsExtensibilityArgumentContext()
}

type ExtensibilityArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensibilityArgumentContext() *ExtensibilityArgumentContext {
	var p = new(ExtensibilityArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_extensibilityArgument
	return p
}

func (*ExtensibilityArgumentContext) IsExtensibilityArgumentContext() {}

func NewExtensibilityArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensibilityArgumentContext {
	var p = new(ExtensibilityArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_extensibilityArgument

	return p
}

func (s *ExtensibilityArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensibilityArgumentContext) IdentifierOrMarker() IIdentifierOrMarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierOrMarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierOrMarkerContext)
}

func (s *ExtensibilityArgumentContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExtensibilityArgumentContext) Time() ITimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeContext)
}

func (s *ExtensibilityArgumentContext) ExtensibilityExpression() IExtensibilityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtensibilityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtensibilityExpressionContext)
}

func (s *ExtensibilityArgumentContext) ExtensibilityTuple() IExtensibilityTupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtensibilityTupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtensibilityTupleContext)
}

func (s *ExtensibilityArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensibilityArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensibilityArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterExtensibilityArgument(s)
	}
}

func (s *ExtensibilityArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitExtensibilityArgument(s)
	}
}

func (p *PROV_NParser) ExtensibilityArgument() (localctx IExtensibilityArgumentContext) {
	this := p
	_ = this

	localctx = NewExtensibilityArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, PROV_NParserRULE_extensibilityArgument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(472)
			p.IdentifierOrMarker()
		}

	case 2:
		{
			p.SetState(473)
			p.Literal()
		}

	case 3:
		{
			p.SetState(474)
			p.Time()
		}

	case 4:
		{
			p.SetState(475)
			p.ExtensibilityExpression()
		}

	case 5:
		{
			p.SetState(476)
			p.ExtensibilityTuple()
		}

	}

	return localctx
}

// IExtensibilityTupleContext is an interface to support dynamic dispatch.
type IExtensibilityTupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtensibilityTupleContext differentiates from other interfaces.
	IsExtensibilityTupleContext()
}

type ExtensibilityTupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensibilityTupleContext() *ExtensibilityTupleContext {
	var p = new(ExtensibilityTupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PROV_NParserRULE_extensibilityTuple
	return p
}

func (*ExtensibilityTupleContext) IsExtensibilityTupleContext() {}

func NewExtensibilityTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensibilityTupleContext {
	var p = new(ExtensibilityTupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PROV_NParserRULE_extensibilityTuple

	return p
}

func (s *ExtensibilityTupleContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensibilityTupleContext) AllExtensibilityArgument() []IExtensibilityArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExtensibilityArgumentContext)(nil)).Elem())
	var tst = make([]IExtensibilityArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExtensibilityArgumentContext)
		}
	}

	return tst
}

func (s *ExtensibilityTupleContext) ExtensibilityArgument(i int) IExtensibilityArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtensibilityArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExtensibilityArgumentContext)
}

func (s *ExtensibilityTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensibilityTupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensibilityTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.EnterExtensibilityTuple(s)
	}
}

func (s *ExtensibilityTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PROV_NListener); ok {
		listenerT.ExitExtensibilityTuple(s)
	}
}

func (p *PROV_NParser) ExtensibilityTuple() (localctx IExtensibilityTupleContext) {
	this := p
	_ = this

	localctx = NewExtensibilityTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, PROV_NParserRULE_extensibilityTuple)
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

	p.SetState(501)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PROV_NParserT__27:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(479)
			p.Match(PROV_NParserT__27)
		}
		{
			p.SetState(480)
			p.ExtensibilityArgument()
		}
		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PROV_NParserT__5 {
			{
				p.SetState(481)
				p.Match(PROV_NParserT__5)
			}
			{
				p.SetState(482)
				p.ExtensibilityArgument()
			}

			p.SetState(487)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(488)
			p.Match(PROV_NParserT__28)
		}

	case PROV_NParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(490)
			p.Match(PROV_NParserT__3)
		}
		{
			p.SetState(491)
			p.ExtensibilityArgument()
		}
		p.SetState(496)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PROV_NParserT__5 {
			{
				p.SetState(492)
				p.Match(PROV_NParserT__5)
			}
			{
				p.SetState(493)
				p.ExtensibilityArgument()
			}

			p.SetState(498)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(499)
			p.Match(PROV_NParserT__4)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
