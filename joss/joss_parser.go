// Code generated from joss.g4 by ANTLR 4.9.3. DO NOT EDIT.

package joss // joss
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 77, 501,
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
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 3, 2, 6, 2, 148, 10, 2, 13, 2, 14,
	2, 149, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 160, 10,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 166, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 176, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 184, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 5, 11, 197, 10, 11, 3, 12, 3, 12, 3, 12, 5, 12, 202,
	10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 5, 18, 218, 10, 18, 3, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 229, 10, 20, 3, 21,
	3, 21, 5, 21, 233, 10, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 5, 23, 242, 10, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 255, 10, 26, 3, 27, 3, 27, 3, 27, 3,
	27, 5, 27, 261, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 267, 10, 28,
	3, 29, 3, 29, 3, 29, 5, 29, 272, 10, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	31, 5, 31, 279, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 33, 5, 33, 289, 10, 33, 3, 34, 3, 34, 5, 34, 293, 10, 34, 3, 35, 3,
	35, 5, 35, 297, 10, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 5, 37,
	305, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 6, 39, 313, 10,
	39, 13, 39, 14, 39, 314, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42,
	3, 42, 3, 42, 3, 43, 3, 43, 5, 43, 328, 10, 43, 3, 44, 3, 44, 3, 44, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 5, 45, 346, 10, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3,
	47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51,
	3, 51, 3, 51, 3, 51, 3, 51, 5, 51, 368, 10, 51, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 52, 3, 52, 5, 52, 377, 10, 52, 3, 53, 3, 53, 3, 53, 3, 53,
	3, 53, 5, 53, 384, 10, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 5, 54, 391,
	10, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55,
	5, 55, 402, 10, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 5,
	57, 411, 10, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58,
	3, 58, 3, 58, 3, 58, 3, 58, 5, 58, 425, 10, 58, 3, 59, 3, 59, 3, 59, 3,
	59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62,
	3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3,
	65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67,
	3, 67, 3, 68, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3,
	69, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3, 71,
	3, 71, 5, 71, 484, 10, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 5,
	72, 492, 10, 72, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 5, 73, 499, 10, 73,
	3, 73, 2, 2, 74, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102,
	104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132,
	134, 136, 138, 140, 142, 144, 2, 7, 3, 2, 8, 10, 3, 2, 18, 20, 4, 2, 8,
	10, 24, 24, 3, 2, 39, 40, 3, 2, 41, 46, 2, 492, 2, 147, 3, 2, 2, 2, 4,
	159, 3, 2, 2, 2, 6, 165, 3, 2, 2, 2, 8, 167, 3, 2, 2, 2, 10, 175, 3, 2,
	2, 2, 12, 183, 3, 2, 2, 2, 14, 185, 3, 2, 2, 2, 16, 187, 3, 2, 2, 2, 18,
	189, 3, 2, 2, 2, 20, 196, 3, 2, 2, 2, 22, 201, 3, 2, 2, 2, 24, 203, 3,
	2, 2, 2, 26, 205, 3, 2, 2, 2, 28, 207, 3, 2, 2, 2, 30, 209, 3, 2, 2, 2,
	32, 212, 3, 2, 2, 2, 34, 217, 3, 2, 2, 2, 36, 219, 3, 2, 2, 2, 38, 228,
	3, 2, 2, 2, 40, 232, 3, 2, 2, 2, 42, 234, 3, 2, 2, 2, 44, 241, 3, 2, 2,
	2, 46, 243, 3, 2, 2, 2, 48, 245, 3, 2, 2, 2, 50, 254, 3, 2, 2, 2, 52, 260,
	3, 2, 2, 2, 54, 266, 3, 2, 2, 2, 56, 271, 3, 2, 2, 2, 58, 273, 3, 2, 2,
	2, 60, 278, 3, 2, 2, 2, 62, 280, 3, 2, 2, 2, 64, 288, 3, 2, 2, 2, 66, 292,
	3, 2, 2, 2, 68, 296, 3, 2, 2, 2, 70, 298, 3, 2, 2, 2, 72, 304, 3, 2, 2,
	2, 74, 306, 3, 2, 2, 2, 76, 312, 3, 2, 2, 2, 78, 316, 3, 2, 2, 2, 80, 319,
	3, 2, 2, 2, 82, 322, 3, 2, 2, 2, 84, 327, 3, 2, 2, 2, 86, 329, 3, 2, 2,
	2, 88, 345, 3, 2, 2, 2, 90, 347, 3, 2, 2, 2, 92, 349, 3, 2, 2, 2, 94, 353,
	3, 2, 2, 2, 96, 355, 3, 2, 2, 2, 98, 358, 3, 2, 2, 2, 100, 367, 3, 2, 2,
	2, 102, 376, 3, 2, 2, 2, 104, 383, 3, 2, 2, 2, 106, 390, 3, 2, 2, 2, 108,
	401, 3, 2, 2, 2, 110, 403, 3, 2, 2, 2, 112, 410, 3, 2, 2, 2, 114, 424,
	3, 2, 2, 2, 116, 426, 3, 2, 2, 2, 118, 430, 3, 2, 2, 2, 120, 434, 3, 2,
	2, 2, 122, 438, 3, 2, 2, 2, 124, 442, 3, 2, 2, 2, 126, 446, 3, 2, 2, 2,
	128, 450, 3, 2, 2, 2, 130, 454, 3, 2, 2, 2, 132, 458, 3, 2, 2, 2, 134,
	462, 3, 2, 2, 2, 136, 466, 3, 2, 2, 2, 138, 472, 3, 2, 2, 2, 140, 483,
	3, 2, 2, 2, 142, 491, 3, 2, 2, 2, 144, 498, 3, 2, 2, 2, 146, 148, 5, 4,
	3, 2, 147, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2,
	149, 150, 3, 2, 2, 2, 150, 3, 3, 2, 2, 2, 151, 152, 5, 6, 4, 2, 152, 153,
	7, 3, 2, 2, 153, 160, 3, 2, 2, 2, 154, 155, 5, 8, 5, 2, 155, 156, 7, 3,
	2, 2, 156, 160, 3, 2, 2, 2, 157, 160, 5, 62, 32, 2, 158, 160, 5, 110, 56,
	2, 159, 151, 3, 2, 2, 2, 159, 154, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159,
	158, 3, 2, 2, 2, 160, 5, 3, 2, 2, 2, 161, 166, 5, 16, 9, 2, 162, 166, 5,
	18, 10, 2, 163, 166, 5, 14, 8, 2, 164, 166, 5, 12, 7, 2, 165, 161, 3, 2,
	2, 2, 165, 162, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 164, 3, 2, 2, 2,
	166, 7, 3, 2, 2, 2, 167, 168, 7, 76, 2, 2, 168, 169, 5, 10, 6, 2, 169,
	9, 3, 2, 2, 2, 170, 176, 5, 32, 17, 2, 171, 176, 5, 26, 14, 2, 172, 176,
	5, 28, 15, 2, 173, 176, 5, 30, 16, 2, 174, 176, 5, 12, 7, 2, 175, 170,
	3, 2, 2, 2, 175, 171, 3, 2, 2, 2, 175, 172, 3, 2, 2, 2, 175, 173, 3, 2,
	2, 2, 175, 174, 3, 2, 2, 2, 176, 11, 3, 2, 2, 2, 177, 184, 5, 36, 19, 2,
	178, 184, 5, 38, 20, 2, 179, 184, 5, 42, 22, 2, 180, 184, 5, 60, 31, 2,
	181, 184, 7, 60, 2, 2, 182, 184, 7, 61, 2, 2, 183, 177, 3, 2, 2, 2, 183,
	178, 3, 2, 2, 2, 183, 179, 3, 2, 2, 2, 183, 180, 3, 2, 2, 2, 183, 181,
	3, 2, 2, 2, 183, 182, 3, 2, 2, 2, 184, 13, 3, 2, 2, 2, 185, 186, 7, 4,
	2, 2, 186, 15, 3, 2, 2, 2, 187, 188, 7, 5, 2, 2, 188, 17, 3, 2, 2, 2, 189,
	190, 7, 6, 2, 2, 190, 191, 5, 20, 11, 2, 191, 19, 3, 2, 2, 2, 192, 197,
	5, 78, 40, 2, 193, 197, 5, 80, 41, 2, 194, 197, 5, 82, 42, 2, 195, 197,
	5, 22, 12, 2, 196, 192, 3, 2, 2, 2, 196, 193, 3, 2, 2, 2, 196, 194, 3,
	2, 2, 2, 196, 195, 3, 2, 2, 2, 197, 21, 3, 2, 2, 2, 198, 202, 7, 7, 2,
	2, 199, 200, 7, 7, 2, 2, 200, 202, 5, 24, 13, 2, 201, 198, 3, 2, 2, 2,
	201, 199, 3, 2, 2, 2, 202, 23, 3, 2, 2, 2, 203, 204, 9, 2, 2, 2, 204, 25,
	3, 2, 2, 2, 205, 206, 7, 11, 2, 2, 206, 27, 3, 2, 2, 2, 207, 208, 7, 12,
	2, 2, 208, 29, 3, 2, 2, 2, 209, 210, 7, 13, 2, 2, 210, 211, 5, 142, 72,
	2, 211, 31, 3, 2, 2, 2, 212, 213, 7, 14, 2, 2, 213, 214, 5, 34, 18, 2,
	214, 33, 3, 2, 2, 2, 215, 218, 5, 78, 40, 2, 216, 218, 5, 80, 41, 2, 217,
	215, 3, 2, 2, 2, 217, 216, 3, 2, 2, 2, 218, 35, 3, 2, 2, 2, 219, 220, 7,
	15, 2, 2, 220, 221, 5, 110, 56, 2, 221, 37, 3, 2, 2, 2, 222, 223, 7, 16,
	2, 2, 223, 229, 5, 40, 21, 2, 224, 225, 7, 16, 2, 2, 225, 226, 5, 40, 21,
	2, 226, 227, 5, 84, 43, 2, 227, 229, 3, 2, 2, 2, 228, 222, 3, 2, 2, 2,
	228, 224, 3, 2, 2, 2, 229, 39, 3, 2, 2, 2, 230, 233, 5, 78, 40, 2, 231,
	233, 5, 80, 41, 2, 232, 230, 3, 2, 2, 2, 232, 231, 3, 2, 2, 2, 233, 41,
	3, 2, 2, 2, 234, 235, 7, 17, 2, 2, 235, 236, 5, 44, 23, 2, 236, 43, 3,
	2, 2, 2, 237, 242, 5, 48, 25, 2, 238, 242, 5, 50, 26, 2, 239, 242, 5, 54,
	28, 2, 240, 242, 5, 46, 24, 2, 241, 237, 3, 2, 2, 2, 241, 238, 3, 2, 2,
	2, 241, 239, 3, 2, 2, 2, 241, 240, 3, 2, 2, 2, 242, 45, 3, 2, 2, 2, 243,
	244, 9, 3, 2, 2, 244, 47, 3, 2, 2, 2, 245, 246, 5, 144, 73, 2, 246, 247,
	7, 21, 2, 2, 247, 248, 5, 82, 42, 2, 248, 49, 3, 2, 2, 2, 249, 255, 5,
	52, 27, 2, 250, 251, 5, 52, 27, 2, 251, 252, 7, 22, 2, 2, 252, 253, 5,
	50, 26, 2, 253, 255, 3, 2, 2, 2, 254, 249, 3, 2, 2, 2, 254, 250, 3, 2,
	2, 2, 255, 51, 3, 2, 2, 2, 256, 257, 7, 23, 2, 2, 257, 258, 7, 64, 2, 2,
	258, 261, 7, 23, 2, 2, 259, 261, 5, 142, 72, 2, 260, 256, 3, 2, 2, 2, 260,
	259, 3, 2, 2, 2, 261, 53, 3, 2, 2, 2, 262, 267, 5, 78, 40, 2, 263, 267,
	5, 80, 41, 2, 264, 267, 5, 82, 42, 2, 265, 267, 5, 56, 29, 2, 266, 262,
	3, 2, 2, 2, 266, 263, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 266, 265, 3, 2,
	2, 2, 267, 55, 3, 2, 2, 2, 268, 272, 7, 7, 2, 2, 269, 270, 7, 7, 2, 2,
	270, 272, 5, 58, 30, 2, 271, 268, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 272,
	57, 3, 2, 2, 2, 273, 274, 9, 4, 2, 2, 274, 59, 3, 2, 2, 2, 275, 279, 5,
	144, 73, 2, 276, 277, 7, 7, 2, 2, 277, 279, 7, 24, 2, 2, 278, 275, 3, 2,
	2, 2, 278, 276, 3, 2, 2, 2, 279, 61, 3, 2, 2, 2, 280, 281, 7, 25, 2, 2,
	281, 282, 7, 75, 2, 2, 282, 283, 7, 26, 2, 2, 283, 63, 3, 2, 2, 2, 284,
	289, 5, 66, 34, 2, 285, 286, 5, 66, 34, 2, 286, 287, 5, 64, 33, 2, 287,
	289, 3, 2, 2, 2, 288, 284, 3, 2, 2, 2, 288, 285, 3, 2, 2, 2, 289, 65, 3,
	2, 2, 2, 290, 293, 5, 68, 35, 2, 291, 293, 7, 64, 2, 2, 292, 290, 3, 2,
	2, 2, 292, 291, 3, 2, 2, 2, 293, 67, 3, 2, 2, 2, 294, 297, 5, 70, 36, 2,
	295, 297, 5, 74, 38, 2, 296, 294, 3, 2, 2, 2, 296, 295, 3, 2, 2, 2, 297,
	69, 3, 2, 2, 2, 298, 299, 7, 27, 2, 2, 299, 300, 5, 72, 37, 2, 300, 71,
	3, 2, 2, 2, 301, 305, 7, 3, 2, 2, 302, 303, 7, 3, 2, 2, 303, 305, 5, 72,
	37, 2, 304, 301, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 305, 73, 3, 2, 2, 2,
	306, 307, 7, 28, 2, 2, 307, 308, 5, 76, 39, 2, 308, 309, 7, 29, 2, 2, 309,
	310, 5, 76, 39, 2, 310, 75, 3, 2, 2, 2, 311, 313, 7, 30, 2, 2, 312, 311,
	3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 314, 315, 3, 2,
	2, 2, 315, 77, 3, 2, 2, 2, 316, 317, 7, 31, 2, 2, 317, 318, 7, 76, 2, 2,
	318, 79, 3, 2, 2, 2, 319, 320, 7, 32, 2, 2, 320, 321, 7, 75, 2, 2, 321,
	81, 3, 2, 2, 2, 322, 323, 7, 33, 2, 2, 323, 324, 7, 75, 2, 2, 324, 83,
	3, 2, 2, 2, 325, 328, 5, 86, 44, 2, 326, 328, 5, 96, 49, 2, 327, 325, 3,
	2, 2, 2, 327, 326, 3, 2, 2, 2, 328, 85, 3, 2, 2, 2, 329, 330, 7, 34, 2,
	2, 330, 331, 5, 88, 45, 2, 331, 87, 3, 2, 2, 2, 332, 333, 7, 35, 2, 2,
	333, 334, 5, 88, 45, 2, 334, 335, 7, 36, 2, 2, 335, 346, 3, 2, 2, 2, 336,
	337, 7, 37, 2, 2, 337, 338, 5, 88, 45, 2, 338, 339, 7, 38, 2, 2, 339, 346,
	3, 2, 2, 2, 340, 346, 5, 92, 47, 2, 341, 342, 5, 92, 47, 2, 342, 343, 5,
	90, 46, 2, 343, 344, 5, 88, 45, 2, 344, 346, 3, 2, 2, 2, 345, 332, 3, 2,
	2, 2, 345, 336, 3, 2, 2, 2, 345, 340, 3, 2, 2, 2, 345, 341, 3, 2, 2, 2,
	346, 89, 3, 2, 2, 2, 347, 348, 9, 5, 2, 2, 348, 91, 3, 2, 2, 2, 349, 350,
	5, 104, 53, 2, 350, 351, 5, 94, 48, 2, 351, 352, 5, 104, 53, 2, 352, 93,
	3, 2, 2, 2, 353, 354, 9, 6, 2, 2, 354, 95, 3, 2, 2, 2, 355, 356, 7, 47,
	2, 2, 356, 357, 5, 98, 50, 2, 357, 97, 3, 2, 2, 2, 358, 359, 5, 142, 72,
	2, 359, 360, 7, 41, 2, 2, 360, 361, 5, 100, 51, 2, 361, 99, 3, 2, 2, 2,
	362, 368, 5, 102, 52, 2, 363, 364, 5, 102, 52, 2, 364, 365, 7, 22, 2, 2,
	365, 366, 5, 100, 51, 2, 366, 368, 3, 2, 2, 2, 367, 362, 3, 2, 2, 2, 367,
	363, 3, 2, 2, 2, 368, 101, 3, 2, 2, 2, 369, 370, 5, 104, 53, 2, 370, 371,
	7, 35, 2, 2, 371, 372, 5, 104, 53, 2, 372, 373, 7, 36, 2, 2, 373, 374,
	5, 102, 52, 2, 374, 377, 3, 2, 2, 2, 375, 377, 5, 104, 53, 2, 376, 369,
	3, 2, 2, 2, 376, 375, 3, 2, 2, 2, 377, 103, 3, 2, 2, 2, 378, 384, 5, 106,
	54, 2, 379, 380, 5, 106, 54, 2, 380, 381, 7, 63, 2, 2, 381, 382, 5, 104,
	53, 2, 382, 384, 3, 2, 2, 2, 383, 378, 3, 2, 2, 2, 383, 379, 3, 2, 2, 2,
	384, 105, 3, 2, 2, 2, 385, 391, 5, 108, 55, 2, 386, 387, 5, 108, 55, 2,
	387, 388, 7, 62, 2, 2, 388, 389, 5, 106, 54, 2, 389, 391, 3, 2, 2, 2, 390,
	385, 3, 2, 2, 2, 390, 386, 3, 2, 2, 2, 391, 107, 3, 2, 2, 2, 392, 393,
	7, 35, 2, 2, 393, 394, 5, 104, 53, 2, 394, 395, 7, 36, 2, 2, 395, 402,
	3, 2, 2, 2, 396, 397, 7, 37, 2, 2, 397, 398, 5, 104, 53, 2, 398, 399, 7,
	38, 2, 2, 399, 402, 3, 2, 2, 2, 400, 402, 5, 112, 57, 2, 401, 392, 3, 2,
	2, 2, 401, 396, 3, 2, 2, 2, 401, 400, 3, 2, 2, 2, 402, 109, 3, 2, 2, 2,
	403, 404, 5, 142, 72, 2, 404, 405, 7, 41, 2, 2, 405, 406, 5, 104, 53, 2,
	406, 111, 3, 2, 2, 2, 407, 411, 7, 68, 2, 2, 408, 411, 5, 142, 72, 2, 409,
	411, 5, 114, 58, 2, 410, 407, 3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 410, 409,
	3, 2, 2, 2, 411, 113, 3, 2, 2, 2, 412, 425, 5, 116, 59, 2, 413, 425, 5,
	118, 60, 2, 414, 425, 5, 120, 61, 2, 415, 425, 5, 122, 62, 2, 416, 425,
	5, 124, 63, 2, 417, 425, 5, 126, 64, 2, 418, 425, 5, 128, 65, 2, 419, 425,
	5, 130, 66, 2, 420, 425, 5, 132, 67, 2, 421, 425, 5, 134, 68, 2, 422, 425,
	5, 136, 69, 2, 423, 425, 5, 138, 70, 2, 424, 412, 3, 2, 2, 2, 424, 413,
	3, 2, 2, 2, 424, 414, 3, 2, 2, 2, 424, 415, 3, 2, 2, 2, 424, 416, 3, 2,
	2, 2, 424, 417, 3, 2, 2, 2, 424, 418, 3, 2, 2, 2, 424, 419, 3, 2, 2, 2,
	424, 420, 3, 2, 2, 2, 424, 421, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 424,
	423, 3, 2, 2, 2, 425, 115, 3, 2, 2, 2, 426, 427, 7, 48, 2, 2, 427, 428,
	5, 104, 53, 2, 428, 429, 7, 36, 2, 2, 429, 117, 3, 2, 2, 2, 430, 431, 7,
	49, 2, 2, 431, 432, 5, 104, 53, 2, 432, 433, 7, 36, 2, 2, 433, 119, 3,
	2, 2, 2, 434, 435, 7, 50, 2, 2, 435, 436, 5, 104, 53, 2, 436, 437, 7, 36,
	2, 2, 437, 121, 3, 2, 2, 2, 438, 439, 7, 51, 2, 2, 439, 440, 5, 104, 53,
	2, 440, 441, 7, 36, 2, 2, 441, 123, 3, 2, 2, 2, 442, 443, 7, 52, 2, 2,
	443, 444, 5, 104, 53, 2, 444, 445, 7, 36, 2, 2, 445, 125, 3, 2, 2, 2, 446,
	447, 7, 53, 2, 2, 447, 448, 5, 104, 53, 2, 448, 449, 7, 36, 2, 2, 449,
	127, 3, 2, 2, 2, 450, 451, 7, 54, 2, 2, 451, 452, 5, 104, 53, 2, 452, 453,
	7, 36, 2, 2, 453, 129, 3, 2, 2, 2, 454, 455, 7, 55, 2, 2, 455, 456, 5,
	104, 53, 2, 456, 457, 7, 36, 2, 2, 457, 131, 3, 2, 2, 2, 458, 459, 7, 56,
	2, 2, 459, 460, 5, 104, 53, 2, 460, 461, 7, 36, 2, 2, 461, 133, 3, 2, 2,
	2, 462, 463, 7, 57, 2, 2, 463, 464, 5, 104, 53, 2, 464, 465, 7, 36, 2,
	2, 465, 135, 3, 2, 2, 2, 466, 467, 7, 58, 2, 2, 467, 468, 5, 104, 53, 2,
	468, 469, 7, 22, 2, 2, 469, 470, 5, 140, 71, 2, 470, 471, 7, 36, 2, 2,
	471, 137, 3, 2, 2, 2, 472, 473, 7, 59, 2, 2, 473, 474, 5, 104, 53, 2, 474,
	475, 7, 22, 2, 2, 475, 476, 5, 140, 71, 2, 476, 477, 7, 36, 2, 2, 477,
	139, 3, 2, 2, 2, 478, 484, 5, 104, 53, 2, 479, 480, 5, 104, 53, 2, 480,
	481, 7, 22, 2, 2, 481, 482, 5, 140, 71, 2, 482, 484, 3, 2, 2, 2, 483, 478,
	3, 2, 2, 2, 483, 479, 3, 2, 2, 2, 484, 141, 3, 2, 2, 2, 485, 492, 7, 66,
	2, 2, 486, 487, 7, 66, 2, 2, 487, 488, 7, 35, 2, 2, 488, 489, 5, 104, 53,
	2, 489, 490, 7, 36, 2, 2, 490, 492, 3, 2, 2, 2, 491, 485, 3, 2, 2, 2, 491,
	486, 3, 2, 2, 2, 492, 143, 3, 2, 2, 2, 493, 499, 5, 142, 72, 2, 494, 495,
	5, 142, 72, 2, 495, 496, 7, 22, 2, 2, 496, 497, 5, 144, 73, 2, 497, 499,
	3, 2, 2, 2, 498, 493, 3, 2, 2, 2, 498, 494, 3, 2, 2, 2, 499, 145, 3, 2,
	2, 2, 35, 149, 159, 165, 175, 183, 196, 201, 217, 228, 232, 241, 254, 260,
	266, 271, 278, 288, 292, 296, 304, 314, 327, 345, 367, 376, 383, 390, 401,
	410, 424, 483, 491, 498,
}
var literalNames = []string{
	"", "'.'", "'Go'", "'Cancel'", "'Delete'", "'all'", "'steps'", "'parts'",
	"'forms'", "'Done'", "'Stop'", "'Demand'", "'To'", "'Set'", "'Do'", "'Type'",
	"'size'", "'time'", "'users'", "'in'", "','", "'''", "'values'", "'Form'",
	"':'", "'..'", "'__'", "'.__'", "'_'", "'step'", "'part'", "'form'", "'if'",
	"'('", "')'", "'['", "']'", "'and'", "'or'", "'='", "'!='", "'<='", "'=>'",
	"'<'", "'>'", "'for'", "'sqrt('", "'log('", "'exp('", "'sin('", "'cos('",
	"'ip('", "'fp('", "'dp('", "'xp('", "'sgn('", "'max('", "'min('", "'Line'",
	"'Page'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "LINECMD", "PAGECMD", "MULOP", "ADDOP", "STRING", "CHAR",
	"ALPHA", "SPECIALCHAR", "NUMBER", "NUMBERPART", "INTPART", "DECIMALPART",
	"DIGIT", "NZDIGIT", "SIGN", "PARTNUMBER", "STEPNUMBER", "WS",
}

var ruleNames = []string{
	"prog", "statement", "direct", "indirect_", "indirectCMD", "arbitraryCMD",
	"goCMD", "cancelCMD", "deleteCmdCMD", "delCmdSel", "delCmdAllSel", "delCmdAllType",
	"doneCMD", "stopCMD", "demandCMD", "toCMD", "toSel", "setCMD", "doCMD",
	"doSel", "typeCMD", "typeContent", "typeSys", "typeForm", "typeData", "typeElem",
	"typeSel", "typeAllSel", "typeAllType", "deleteValCMD", "formCMD", "formContent",
	"formObject", "formPH", "sciNotation", "dot", "fixedNotation", "uScore",
	"stepSel", "partSel", "formSel", "conditional", "if_", "boolExp", "boolOp",
	"comparison", "boolComp", "for_", "range_", "rangeExp", "rangeVal", "mathExp",
	"term", "factor", "assignment", "value", "function_", "funcSqrt", "funcLog",
	"funcExp", "funcSin", "funcCos", "funcIp", "funcFp", "funcDp", "funcXp",
	"funcSgn", "funcMax", "funcMin", "argList", "variable", "varList",
}

type jossParser struct {
	*antlr.BaseParser
}

// NewjossParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *jossParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewjossParser(input antlr.TokenStream) *jossParser {
	this := new(jossParser)
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
	this.GrammarFileName = "joss.g4"

	return this
}

// jossParser tokens.
const (
	jossParserEOF         = antlr.TokenEOF
	jossParserT__0        = 1
	jossParserT__1        = 2
	jossParserT__2        = 3
	jossParserT__3        = 4
	jossParserT__4        = 5
	jossParserT__5        = 6
	jossParserT__6        = 7
	jossParserT__7        = 8
	jossParserT__8        = 9
	jossParserT__9        = 10
	jossParserT__10       = 11
	jossParserT__11       = 12
	jossParserT__12       = 13
	jossParserT__13       = 14
	jossParserT__14       = 15
	jossParserT__15       = 16
	jossParserT__16       = 17
	jossParserT__17       = 18
	jossParserT__18       = 19
	jossParserT__19       = 20
	jossParserT__20       = 21
	jossParserT__21       = 22
	jossParserT__22       = 23
	jossParserT__23       = 24
	jossParserT__24       = 25
	jossParserT__25       = 26
	jossParserT__26       = 27
	jossParserT__27       = 28
	jossParserT__28       = 29
	jossParserT__29       = 30
	jossParserT__30       = 31
	jossParserT__31       = 32
	jossParserT__32       = 33
	jossParserT__33       = 34
	jossParserT__34       = 35
	jossParserT__35       = 36
	jossParserT__36       = 37
	jossParserT__37       = 38
	jossParserT__38       = 39
	jossParserT__39       = 40
	jossParserT__40       = 41
	jossParserT__41       = 42
	jossParserT__42       = 43
	jossParserT__43       = 44
	jossParserT__44       = 45
	jossParserT__45       = 46
	jossParserT__46       = 47
	jossParserT__47       = 48
	jossParserT__48       = 49
	jossParserT__49       = 50
	jossParserT__50       = 51
	jossParserT__51       = 52
	jossParserT__52       = 53
	jossParserT__53       = 54
	jossParserT__54       = 55
	jossParserT__55       = 56
	jossParserT__56       = 57
	jossParserLINECMD     = 58
	jossParserPAGECMD     = 59
	jossParserMULOP       = 60
	jossParserADDOP       = 61
	jossParserSTRING      = 62
	jossParserCHAR        = 63
	jossParserALPHA       = 64
	jossParserSPECIALCHAR = 65
	jossParserNUMBER      = 66
	jossParserNUMBERPART  = 67
	jossParserINTPART     = 68
	jossParserDECIMALPART = 69
	jossParserDIGIT       = 70
	jossParserNZDIGIT     = 71
	jossParserSIGN        = 72
	jossParserPARTNUMBER  = 73
	jossParserSTEPNUMBER  = 74
	jossParserWS          = 75
)

// jossParser rules.
const (
	jossParserRULE_prog          = 0
	jossParserRULE_statement     = 1
	jossParserRULE_direct        = 2
	jossParserRULE_indirect_     = 3
	jossParserRULE_indirectCMD   = 4
	jossParserRULE_arbitraryCMD  = 5
	jossParserRULE_goCMD         = 6
	jossParserRULE_cancelCMD     = 7
	jossParserRULE_deleteCmdCMD  = 8
	jossParserRULE_delCmdSel     = 9
	jossParserRULE_delCmdAllSel  = 10
	jossParserRULE_delCmdAllType = 11
	jossParserRULE_doneCMD       = 12
	jossParserRULE_stopCMD       = 13
	jossParserRULE_demandCMD     = 14
	jossParserRULE_toCMD         = 15
	jossParserRULE_toSel         = 16
	jossParserRULE_setCMD        = 17
	jossParserRULE_doCMD         = 18
	jossParserRULE_doSel         = 19
	jossParserRULE_typeCMD       = 20
	jossParserRULE_typeContent   = 21
	jossParserRULE_typeSys       = 22
	jossParserRULE_typeForm      = 23
	jossParserRULE_typeData      = 24
	jossParserRULE_typeElem      = 25
	jossParserRULE_typeSel       = 26
	jossParserRULE_typeAllSel    = 27
	jossParserRULE_typeAllType   = 28
	jossParserRULE_deleteValCMD  = 29
	jossParserRULE_formCMD       = 30
	jossParserRULE_formContent   = 31
	jossParserRULE_formObject    = 32
	jossParserRULE_formPH        = 33
	jossParserRULE_sciNotation   = 34
	jossParserRULE_dot           = 35
	jossParserRULE_fixedNotation = 36
	jossParserRULE_uScore        = 37
	jossParserRULE_stepSel       = 38
	jossParserRULE_partSel       = 39
	jossParserRULE_formSel       = 40
	jossParserRULE_conditional   = 41
	jossParserRULE_if_           = 42
	jossParserRULE_boolExp       = 43
	jossParserRULE_boolOp        = 44
	jossParserRULE_comparison    = 45
	jossParserRULE_boolComp      = 46
	jossParserRULE_for_          = 47
	jossParserRULE_range_        = 48
	jossParserRULE_rangeExp      = 49
	jossParserRULE_rangeVal      = 50
	jossParserRULE_mathExp       = 51
	jossParserRULE_term          = 52
	jossParserRULE_factor        = 53
	jossParserRULE_assignment    = 54
	jossParserRULE_value         = 55
	jossParserRULE_function_     = 56
	jossParserRULE_funcSqrt      = 57
	jossParserRULE_funcLog       = 58
	jossParserRULE_funcExp       = 59
	jossParserRULE_funcSin       = 60
	jossParserRULE_funcCos       = 61
	jossParserRULE_funcIp        = 62
	jossParserRULE_funcFp        = 63
	jossParserRULE_funcDp        = 64
	jossParserRULE_funcXp        = 65
	jossParserRULE_funcSgn       = 66
	jossParserRULE_funcMax       = 67
	jossParserRULE_funcMin       = 68
	jossParserRULE_argList       = 69
	jossParserRULE_variable      = 70
	jossParserRULE_varList       = 71
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *jossParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, jossParserRULE_prog)
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

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<jossParserT__1)|(1<<jossParserT__2)|(1<<jossParserT__3)|(1<<jossParserT__4)|(1<<jossParserT__12)|(1<<jossParserT__13)|(1<<jossParserT__14)|(1<<jossParserT__22))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(jossParserLINECMD-58))|(1<<(jossParserPAGECMD-58))|(1<<(jossParserALPHA-58))|(1<<(jossParserSTEPNUMBER-58)))) != 0) {
		{
			p.SetState(144)
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
	p.RuleIndex = jossParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Direct() IDirectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectContext)
}

func (s *StatementContext) Indirect_() IIndirect_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndirect_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndirect_Context)
}

func (s *StatementContext) FormCMD() IFormCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormCMDContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *jossParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, jossParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Direct()
		}
		{
			p.SetState(150)
			p.Match(jossParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Indirect_()
		}
		{
			p.SetState(153)
			p.Match(jossParserT__0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.FormCMD()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(156)
			p.Assignment()
		}

	}

	return localctx
}

// IDirectContext is an interface to support dynamic dispatch.
type IDirectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectContext differentiates from other interfaces.
	IsDirectContext()
}

type DirectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectContext() *DirectContext {
	var p = new(DirectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_direct
	return p
}

func (*DirectContext) IsDirectContext() {}

func NewDirectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectContext {
	var p = new(DirectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_direct

	return p
}

func (s *DirectContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectContext) CancelCMD() ICancelCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICancelCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICancelCMDContext)
}

func (s *DirectContext) DeleteCmdCMD() IDeleteCmdCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteCmdCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteCmdCMDContext)
}

func (s *DirectContext) GoCMD() IGoCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGoCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGoCMDContext)
}

func (s *DirectContext) ArbitraryCMD() IArbitraryCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArbitraryCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArbitraryCMDContext)
}

func (s *DirectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDirect(s)
	}
}

func (s *DirectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDirect(s)
	}
}

func (p *jossParser) Direct() (localctx IDirectContext) {
	this := p
	_ = this

	localctx = NewDirectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, jossParserRULE_direct)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.CancelCMD()
		}

	case jossParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.DeleteCmdCMD()
		}

	case jossParserT__1:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.GoCMD()
		}

	case jossParserT__4, jossParserT__12, jossParserT__13, jossParserT__14, jossParserLINECMD, jossParserPAGECMD, jossParserALPHA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(162)
			p.ArbitraryCMD()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIndirect_Context is an interface to support dynamic dispatch.
type IIndirect_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndirect_Context differentiates from other interfaces.
	IsIndirect_Context()
}

type Indirect_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndirect_Context() *Indirect_Context {
	var p = new(Indirect_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_indirect_
	return p
}

func (*Indirect_Context) IsIndirect_Context() {}

func NewIndirect_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Indirect_Context {
	var p = new(Indirect_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_indirect_

	return p
}

func (s *Indirect_Context) GetParser() antlr.Parser { return s.parser }

func (s *Indirect_Context) STEPNUMBER() antlr.TerminalNode {
	return s.GetToken(jossParserSTEPNUMBER, 0)
}

func (s *Indirect_Context) IndirectCMD() IIndirectCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndirectCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndirectCMDContext)
}

func (s *Indirect_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Indirect_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Indirect_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterIndirect_(s)
	}
}

func (s *Indirect_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitIndirect_(s)
	}
}

func (p *jossParser) Indirect_() (localctx IIndirect_Context) {
	this := p
	_ = this

	localctx = NewIndirect_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, jossParserRULE_indirect_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(165)
		p.Match(jossParserSTEPNUMBER)
	}
	{
		p.SetState(166)
		p.IndirectCMD()
	}

	return localctx
}

// IIndirectCMDContext is an interface to support dynamic dispatch.
type IIndirectCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndirectCMDContext differentiates from other interfaces.
	IsIndirectCMDContext()
}

type IndirectCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndirectCMDContext() *IndirectCMDContext {
	var p = new(IndirectCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_indirectCMD
	return p
}

func (*IndirectCMDContext) IsIndirectCMDContext() {}

func NewIndirectCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndirectCMDContext {
	var p = new(IndirectCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_indirectCMD

	return p
}

func (s *IndirectCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *IndirectCMDContext) ToCMD() IToCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IToCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IToCMDContext)
}

func (s *IndirectCMDContext) DoneCMD() IDoneCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoneCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoneCMDContext)
}

func (s *IndirectCMDContext) StopCMD() IStopCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStopCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStopCMDContext)
}

func (s *IndirectCMDContext) DemandCMD() IDemandCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDemandCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDemandCMDContext)
}

func (s *IndirectCMDContext) ArbitraryCMD() IArbitraryCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArbitraryCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArbitraryCMDContext)
}

func (s *IndirectCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndirectCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndirectCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterIndirectCMD(s)
	}
}

func (s *IndirectCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitIndirectCMD(s)
	}
}

func (p *jossParser) IndirectCMD() (localctx IIndirectCMDContext) {
	this := p
	_ = this

	localctx = NewIndirectCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, jossParserRULE_indirectCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	switch p.GetTokenStream().LA(1) {
	case jossParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.ToCMD()
		}

	case jossParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)
			p.DoneCMD()
		}

	case jossParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(170)
			p.StopCMD()
		}

	case jossParserT__10:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(171)
			p.DemandCMD()
		}

	case jossParserT__4, jossParserT__12, jossParserT__13, jossParserT__14, jossParserLINECMD, jossParserPAGECMD, jossParserALPHA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(172)
			p.ArbitraryCMD()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArbitraryCMDContext is an interface to support dynamic dispatch.
type IArbitraryCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArbitraryCMDContext differentiates from other interfaces.
	IsArbitraryCMDContext()
}

type ArbitraryCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArbitraryCMDContext() *ArbitraryCMDContext {
	var p = new(ArbitraryCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_arbitraryCMD
	return p
}

func (*ArbitraryCMDContext) IsArbitraryCMDContext() {}

func NewArbitraryCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArbitraryCMDContext {
	var p = new(ArbitraryCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_arbitraryCMD

	return p
}

func (s *ArbitraryCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *ArbitraryCMDContext) SetCMD() ISetCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetCMDContext)
}

func (s *ArbitraryCMDContext) DoCMD() IDoCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoCMDContext)
}

func (s *ArbitraryCMDContext) TypeCMD() ITypeCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeCMDContext)
}

func (s *ArbitraryCMDContext) DeleteValCMD() IDeleteValCMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteValCMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteValCMDContext)
}

func (s *ArbitraryCMDContext) LINECMD() antlr.TerminalNode {
	return s.GetToken(jossParserLINECMD, 0)
}

func (s *ArbitraryCMDContext) PAGECMD() antlr.TerminalNode {
	return s.GetToken(jossParserPAGECMD, 0)
}

func (s *ArbitraryCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArbitraryCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArbitraryCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterArbitraryCMD(s)
	}
}

func (s *ArbitraryCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitArbitraryCMD(s)
	}
}

func (p *jossParser) ArbitraryCMD() (localctx IArbitraryCMDContext) {
	this := p
	_ = this

	localctx = NewArbitraryCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, jossParserRULE_arbitraryCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__12:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.SetCMD()
		}

	case jossParserT__13:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.DoCMD()
		}

	case jossParserT__14:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.TypeCMD()
		}

	case jossParserT__4, jossParserALPHA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(178)
			p.DeleteValCMD()
		}

	case jossParserLINECMD:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(179)
			p.Match(jossParserLINECMD)
		}

	case jossParserPAGECMD:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(180)
			p.Match(jossParserPAGECMD)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGoCMDContext is an interface to support dynamic dispatch.
type IGoCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGoCMDContext differentiates from other interfaces.
	IsGoCMDContext()
}

type GoCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoCMDContext() *GoCMDContext {
	var p = new(GoCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_goCMD
	return p
}

func (*GoCMDContext) IsGoCMDContext() {}

func NewGoCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoCMDContext {
	var p = new(GoCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_goCMD

	return p
}

func (s *GoCMDContext) GetParser() antlr.Parser { return s.parser }
func (s *GoCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterGoCMD(s)
	}
}

func (s *GoCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitGoCMD(s)
	}
}

func (p *jossParser) GoCMD() (localctx IGoCMDContext) {
	this := p
	_ = this

	localctx = NewGoCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, jossParserRULE_goCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__1)
	}

	return localctx
}

// ICancelCMDContext is an interface to support dynamic dispatch.
type ICancelCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCancelCMDContext differentiates from other interfaces.
	IsCancelCMDContext()
}

type CancelCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCancelCMDContext() *CancelCMDContext {
	var p = new(CancelCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_cancelCMD
	return p
}

func (*CancelCMDContext) IsCancelCMDContext() {}

func NewCancelCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CancelCMDContext {
	var p = new(CancelCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_cancelCMD

	return p
}

func (s *CancelCMDContext) GetParser() antlr.Parser { return s.parser }
func (s *CancelCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CancelCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CancelCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterCancelCMD(s)
	}
}

func (s *CancelCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitCancelCMD(s)
	}
}

func (p *jossParser) CancelCMD() (localctx ICancelCMDContext) {
	this := p
	_ = this

	localctx = NewCancelCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, jossParserRULE_cancelCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(185)
		p.Match(jossParserT__2)
	}

	return localctx
}

// IDeleteCmdCMDContext is an interface to support dynamic dispatch.
type IDeleteCmdCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteCmdCMDContext differentiates from other interfaces.
	IsDeleteCmdCMDContext()
}

type DeleteCmdCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteCmdCMDContext() *DeleteCmdCMDContext {
	var p = new(DeleteCmdCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_deleteCmdCMD
	return p
}

func (*DeleteCmdCMDContext) IsDeleteCmdCMDContext() {}

func NewDeleteCmdCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteCmdCMDContext {
	var p = new(DeleteCmdCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_deleteCmdCMD

	return p
}

func (s *DeleteCmdCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteCmdCMDContext) DelCmdSel() IDelCmdSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelCmdSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelCmdSelContext)
}

func (s *DeleteCmdCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteCmdCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteCmdCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDeleteCmdCMD(s)
	}
}

func (s *DeleteCmdCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDeleteCmdCMD(s)
	}
}

func (p *jossParser) DeleteCmdCMD() (localctx IDeleteCmdCMDContext) {
	this := p
	_ = this

	localctx = NewDeleteCmdCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, jossParserRULE_deleteCmdCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__3)
	}
	{
		p.SetState(188)
		p.DelCmdSel()
	}

	return localctx
}

// IDelCmdSelContext is an interface to support dynamic dispatch.
type IDelCmdSelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelCmdSelContext differentiates from other interfaces.
	IsDelCmdSelContext()
}

type DelCmdSelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelCmdSelContext() *DelCmdSelContext {
	var p = new(DelCmdSelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_delCmdSel
	return p
}

func (*DelCmdSelContext) IsDelCmdSelContext() {}

func NewDelCmdSelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelCmdSelContext {
	var p = new(DelCmdSelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_delCmdSel

	return p
}

func (s *DelCmdSelContext) GetParser() antlr.Parser { return s.parser }

func (s *DelCmdSelContext) StepSel() IStepSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepSelContext)
}

func (s *DelCmdSelContext) PartSel() IPartSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartSelContext)
}

func (s *DelCmdSelContext) FormSel() IFormSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormSelContext)
}

func (s *DelCmdSelContext) DelCmdAllSel() IDelCmdAllSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelCmdAllSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelCmdAllSelContext)
}

func (s *DelCmdSelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelCmdSelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelCmdSelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDelCmdSel(s)
	}
}

func (s *DelCmdSelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDelCmdSel(s)
	}
}

func (p *jossParser) DelCmdSel() (localctx IDelCmdSelContext) {
	this := p
	_ = this

	localctx = NewDelCmdSelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, jossParserRULE_delCmdSel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(194)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.StepSel()
		}

	case jossParserT__29:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.PartSel()
		}

	case jossParserT__30:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(192)
			p.FormSel()
		}

	case jossParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(193)
			p.DelCmdAllSel()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDelCmdAllSelContext is an interface to support dynamic dispatch.
type IDelCmdAllSelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelCmdAllSelContext differentiates from other interfaces.
	IsDelCmdAllSelContext()
}

type DelCmdAllSelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelCmdAllSelContext() *DelCmdAllSelContext {
	var p = new(DelCmdAllSelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_delCmdAllSel
	return p
}

func (*DelCmdAllSelContext) IsDelCmdAllSelContext() {}

func NewDelCmdAllSelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelCmdAllSelContext {
	var p = new(DelCmdAllSelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_delCmdAllSel

	return p
}

func (s *DelCmdAllSelContext) GetParser() antlr.Parser { return s.parser }

func (s *DelCmdAllSelContext) DelCmdAllType() IDelCmdAllTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelCmdAllTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelCmdAllTypeContext)
}

func (s *DelCmdAllSelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelCmdAllSelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelCmdAllSelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDelCmdAllSel(s)
	}
}

func (s *DelCmdAllSelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDelCmdAllSel(s)
	}
}

func (p *jossParser) DelCmdAllSel() (localctx IDelCmdAllSelContext) {
	this := p
	_ = this

	localctx = NewDelCmdAllSelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, jossParserRULE_delCmdAllSel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
			p.Match(jossParserT__4)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Match(jossParserT__4)
		}
		{
			p.SetState(198)
			p.DelCmdAllType()
		}

	}

	return localctx
}

// IDelCmdAllTypeContext is an interface to support dynamic dispatch.
type IDelCmdAllTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelCmdAllTypeContext differentiates from other interfaces.
	IsDelCmdAllTypeContext()
}

type DelCmdAllTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelCmdAllTypeContext() *DelCmdAllTypeContext {
	var p = new(DelCmdAllTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_delCmdAllType
	return p
}

func (*DelCmdAllTypeContext) IsDelCmdAllTypeContext() {}

func NewDelCmdAllTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelCmdAllTypeContext {
	var p = new(DelCmdAllTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_delCmdAllType

	return p
}

func (s *DelCmdAllTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *DelCmdAllTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelCmdAllTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelCmdAllTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDelCmdAllType(s)
	}
}

func (s *DelCmdAllTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDelCmdAllType(s)
	}
}

func (p *jossParser) DelCmdAllType() (localctx IDelCmdAllTypeContext) {
	this := p
	_ = this

	localctx = NewDelCmdAllTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, jossParserRULE_delCmdAllType)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<jossParserT__5)|(1<<jossParserT__6)|(1<<jossParserT__7))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDoneCMDContext is an interface to support dynamic dispatch.
type IDoneCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoneCMDContext differentiates from other interfaces.
	IsDoneCMDContext()
}

type DoneCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoneCMDContext() *DoneCMDContext {
	var p = new(DoneCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_doneCMD
	return p
}

func (*DoneCMDContext) IsDoneCMDContext() {}

func NewDoneCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoneCMDContext {
	var p = new(DoneCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_doneCMD

	return p
}

func (s *DoneCMDContext) GetParser() antlr.Parser { return s.parser }
func (s *DoneCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoneCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoneCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDoneCMD(s)
	}
}

func (s *DoneCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDoneCMD(s)
	}
}

func (p *jossParser) DoneCMD() (localctx IDoneCMDContext) {
	this := p
	_ = this

	localctx = NewDoneCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, jossParserRULE_doneCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(203)
		p.Match(jossParserT__8)
	}

	return localctx
}

// IStopCMDContext is an interface to support dynamic dispatch.
type IStopCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStopCMDContext differentiates from other interfaces.
	IsStopCMDContext()
}

type StopCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStopCMDContext() *StopCMDContext {
	var p = new(StopCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_stopCMD
	return p
}

func (*StopCMDContext) IsStopCMDContext() {}

func NewStopCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StopCMDContext {
	var p = new(StopCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_stopCMD

	return p
}

func (s *StopCMDContext) GetParser() antlr.Parser { return s.parser }
func (s *StopCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StopCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StopCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterStopCMD(s)
	}
}

func (s *StopCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitStopCMD(s)
	}
}

func (p *jossParser) StopCMD() (localctx IStopCMDContext) {
	this := p
	_ = this

	localctx = NewStopCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, jossParserRULE_stopCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__9)
	}

	return localctx
}

// IDemandCMDContext is an interface to support dynamic dispatch.
type IDemandCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDemandCMDContext differentiates from other interfaces.
	IsDemandCMDContext()
}

type DemandCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDemandCMDContext() *DemandCMDContext {
	var p = new(DemandCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_demandCMD
	return p
}

func (*DemandCMDContext) IsDemandCMDContext() {}

func NewDemandCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DemandCMDContext {
	var p = new(DemandCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_demandCMD

	return p
}

func (s *DemandCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *DemandCMDContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *DemandCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DemandCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DemandCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDemandCMD(s)
	}
}

func (s *DemandCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDemandCMD(s)
	}
}

func (p *jossParser) DemandCMD() (localctx IDemandCMDContext) {
	this := p
	_ = this

	localctx = NewDemandCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, jossParserRULE_demandCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__10)
	}
	{
		p.SetState(208)
		p.Variable()
	}

	return localctx
}

// IToCMDContext is an interface to support dynamic dispatch.
type IToCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsToCMDContext differentiates from other interfaces.
	IsToCMDContext()
}

type ToCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToCMDContext() *ToCMDContext {
	var p = new(ToCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_toCMD
	return p
}

func (*ToCMDContext) IsToCMDContext() {}

func NewToCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToCMDContext {
	var p = new(ToCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_toCMD

	return p
}

func (s *ToCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *ToCMDContext) ToSel() IToSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IToSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IToSelContext)
}

func (s *ToCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ToCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterToCMD(s)
	}
}

func (s *ToCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitToCMD(s)
	}
}

func (p *jossParser) ToCMD() (localctx IToCMDContext) {
	this := p
	_ = this

	localctx = NewToCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, jossParserRULE_toCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(210)
		p.Match(jossParserT__11)
	}
	{
		p.SetState(211)
		p.ToSel()
	}

	return localctx
}

// IToSelContext is an interface to support dynamic dispatch.
type IToSelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsToSelContext differentiates from other interfaces.
	IsToSelContext()
}

type ToSelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToSelContext() *ToSelContext {
	var p = new(ToSelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_toSel
	return p
}

func (*ToSelContext) IsToSelContext() {}

func NewToSelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToSelContext {
	var p = new(ToSelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_toSel

	return p
}

func (s *ToSelContext) GetParser() antlr.Parser { return s.parser }

func (s *ToSelContext) StepSel() IStepSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepSelContext)
}

func (s *ToSelContext) PartSel() IPartSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartSelContext)
}

func (s *ToSelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToSelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ToSelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterToSel(s)
	}
}

func (s *ToSelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitToSel(s)
	}
}

func (p *jossParser) ToSel() (localctx IToSelContext) {
	this := p
	_ = this

	localctx = NewToSelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, jossParserRULE_toSel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.StepSel()
		}

	case jossParserT__29:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.PartSel()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISetCMDContext is an interface to support dynamic dispatch.
type ISetCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetCMDContext differentiates from other interfaces.
	IsSetCMDContext()
}

type SetCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetCMDContext() *SetCMDContext {
	var p = new(SetCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_setCMD
	return p
}

func (*SetCMDContext) IsSetCMDContext() {}

func NewSetCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetCMDContext {
	var p = new(SetCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_setCMD

	return p
}

func (s *SetCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *SetCMDContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *SetCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterSetCMD(s)
	}
}

func (s *SetCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitSetCMD(s)
	}
}

func (p *jossParser) SetCMD() (localctx ISetCMDContext) {
	this := p
	_ = this

	localctx = NewSetCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, jossParserRULE_setCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(217)
		p.Match(jossParserT__12)
	}
	{
		p.SetState(218)
		p.Assignment()
	}

	return localctx
}

// IDoCMDContext is an interface to support dynamic dispatch.
type IDoCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoCMDContext differentiates from other interfaces.
	IsDoCMDContext()
}

type DoCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoCMDContext() *DoCMDContext {
	var p = new(DoCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_doCMD
	return p
}

func (*DoCMDContext) IsDoCMDContext() {}

func NewDoCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoCMDContext {
	var p = new(DoCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_doCMD

	return p
}

func (s *DoCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *DoCMDContext) DoSel() IDoSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoSelContext)
}

func (s *DoCMDContext) Conditional() IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *DoCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDoCMD(s)
	}
}

func (s *DoCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDoCMD(s)
	}
}

func (p *jossParser) DoCMD() (localctx IDoCMDContext) {
	this := p
	_ = this

	localctx = NewDoCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, jossParserRULE_doCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Match(jossParserT__13)
		}
		{
			p.SetState(221)
			p.DoSel()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(jossParserT__13)
		}
		{
			p.SetState(223)
			p.DoSel()
		}
		{
			p.SetState(224)
			p.Conditional()
		}

	}

	return localctx
}

// IDoSelContext is an interface to support dynamic dispatch.
type IDoSelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoSelContext differentiates from other interfaces.
	IsDoSelContext()
}

type DoSelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoSelContext() *DoSelContext {
	var p = new(DoSelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_doSel
	return p
}

func (*DoSelContext) IsDoSelContext() {}

func NewDoSelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoSelContext {
	var p = new(DoSelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_doSel

	return p
}

func (s *DoSelContext) GetParser() antlr.Parser { return s.parser }

func (s *DoSelContext) StepSel() IStepSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepSelContext)
}

func (s *DoSelContext) PartSel() IPartSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartSelContext)
}

func (s *DoSelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoSelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoSelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDoSel(s)
	}
}

func (s *DoSelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDoSel(s)
	}
}

func (p *jossParser) DoSel() (localctx IDoSelContext) {
	this := p
	_ = this

	localctx = NewDoSelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, jossParserRULE_doSel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.StepSel()
		}

	case jossParserT__29:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.PartSel()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeCMDContext is an interface to support dynamic dispatch.
type ITypeCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeCMDContext differentiates from other interfaces.
	IsTypeCMDContext()
}

type TypeCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeCMDContext() *TypeCMDContext {
	var p = new(TypeCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_typeCMD
	return p
}

func (*TypeCMDContext) IsTypeCMDContext() {}

func NewTypeCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeCMDContext {
	var p = new(TypeCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_typeCMD

	return p
}

func (s *TypeCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeCMDContext) TypeContent() ITypeContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeContentContext)
}

func (s *TypeCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTypeCMD(s)
	}
}

func (s *TypeCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTypeCMD(s)
	}
}

func (p *jossParser) TypeCMD() (localctx ITypeCMDContext) {
	this := p
	_ = this

	localctx = NewTypeCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, jossParserRULE_typeCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__14)
	}
	{
		p.SetState(233)
		p.TypeContent()
	}

	return localctx
}

// ITypeContentContext is an interface to support dynamic dispatch.
type ITypeContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContentContext differentiates from other interfaces.
	IsTypeContentContext()
}

type TypeContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContentContext() *TypeContentContext {
	var p = new(TypeContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_typeContent
	return p
}

func (*TypeContentContext) IsTypeContentContext() {}

func NewTypeContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContentContext {
	var p = new(TypeContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_typeContent

	return p
}

func (s *TypeContentContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContentContext) TypeForm() ITypeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormContext)
}

func (s *TypeContentContext) TypeData() ITypeDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDataContext)
}

func (s *TypeContentContext) TypeSel() ITypeSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSelContext)
}

func (s *TypeContentContext) TypeSys() ITypeSysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSysContext)
}

func (s *TypeContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTypeContent(s)
	}
}

func (s *TypeContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTypeContent(s)
	}
}

func (p *jossParser) TypeContent() (localctx ITypeContentContext) {
	this := p
	_ = this

	localctx = NewTypeContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, jossParserRULE_typeContent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.TypeForm()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.TypeData()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(237)
			p.TypeSel()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(238)
			p.TypeSys()
		}

	}

	return localctx
}

// ITypeSysContext is an interface to support dynamic dispatch.
type ITypeSysContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSysContext differentiates from other interfaces.
	IsTypeSysContext()
}

type TypeSysContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSysContext() *TypeSysContext {
	var p = new(TypeSysContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_typeSys
	return p
}

func (*TypeSysContext) IsTypeSysContext() {}

func NewTypeSysContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSysContext {
	var p = new(TypeSysContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_typeSys

	return p
}

func (s *TypeSysContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeSysContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSysContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSysContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTypeSys(s)
	}
}

func (s *TypeSysContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTypeSys(s)
	}
}

func (p *jossParser) TypeSys() (localctx ITypeSysContext) {
	this := p
	_ = this

	localctx = NewTypeSysContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, jossParserRULE_typeSys)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<jossParserT__15)|(1<<jossParserT__16)|(1<<jossParserT__17))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeFormContext is an interface to support dynamic dispatch.
type ITypeFormContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeFormContext differentiates from other interfaces.
	IsTypeFormContext()
}

type TypeFormContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFormContext() *TypeFormContext {
	var p = new(TypeFormContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_typeForm
	return p
}

func (*TypeFormContext) IsTypeFormContext() {}

func NewTypeFormContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFormContext {
	var p = new(TypeFormContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_typeForm

	return p
}

func (s *TypeFormContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFormContext) VarList() IVarListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarListContext)
}

func (s *TypeFormContext) FormSel() IFormSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormSelContext)
}

func (s *TypeFormContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFormContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFormContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTypeForm(s)
	}
}

func (s *TypeFormContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTypeForm(s)
	}
}

func (p *jossParser) TypeForm() (localctx ITypeFormContext) {
	this := p
	_ = this

	localctx = NewTypeFormContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, jossParserRULE_typeForm)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(243)
		p.VarList()
	}
	{
		p.SetState(244)
		p.Match(jossParserT__18)
	}
	{
		p.SetState(245)
		p.FormSel()
	}

	return localctx
}

// ITypeDataContext is an interface to support dynamic dispatch.
type ITypeDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDataContext differentiates from other interfaces.
	IsTypeDataContext()
}

type TypeDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDataContext() *TypeDataContext {
	var p = new(TypeDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_typeData
	return p
}

func (*TypeDataContext) IsTypeDataContext() {}

func NewTypeDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDataContext {
	var p = new(TypeDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_typeData

	return p
}

func (s *TypeDataContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDataContext) TypeElem() ITypeElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeElemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeElemContext)
}

func (s *TypeDataContext) TypeData() ITypeDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDataContext)
}

func (s *TypeDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTypeData(s)
	}
}

func (s *TypeDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTypeData(s)
	}
}

func (p *jossParser) TypeData() (localctx ITypeDataContext) {
	this := p
	_ = this

	localctx = NewTypeDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, jossParserRULE_typeData)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.TypeElem()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(248)
			p.TypeElem()
		}
		{
			p.SetState(249)
			p.Match(jossParserT__19)
		}
		{
			p.SetState(250)
			p.TypeData()
		}

	}

	return localctx
}

// ITypeElemContext is an interface to support dynamic dispatch.
type ITypeElemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeElemContext differentiates from other interfaces.
	IsTypeElemContext()
}

type TypeElemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeElemContext() *TypeElemContext {
	var p = new(TypeElemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_typeElem
	return p
}

func (*TypeElemContext) IsTypeElemContext() {}

func NewTypeElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeElemContext {
	var p = new(TypeElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_typeElem

	return p
}

func (s *TypeElemContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeElemContext) STRING() antlr.TerminalNode {
	return s.GetToken(jossParserSTRING, 0)
}

func (s *TypeElemContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *TypeElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeElemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeElemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTypeElem(s)
	}
}

func (s *TypeElemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTypeElem(s)
	}
}

func (p *jossParser) TypeElem() (localctx ITypeElemContext) {
	this := p
	_ = this

	localctx = NewTypeElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, jossParserRULE_typeElem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(258)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.Match(jossParserT__20)
		}
		{
			p.SetState(255)
			p.Match(jossParserSTRING)
		}
		{
			p.SetState(256)
			p.Match(jossParserT__20)
		}

	case jossParserALPHA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.Variable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeSelContext is an interface to support dynamic dispatch.
type ITypeSelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSelContext differentiates from other interfaces.
	IsTypeSelContext()
}

type TypeSelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSelContext() *TypeSelContext {
	var p = new(TypeSelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_typeSel
	return p
}

func (*TypeSelContext) IsTypeSelContext() {}

func NewTypeSelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSelContext {
	var p = new(TypeSelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_typeSel

	return p
}

func (s *TypeSelContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSelContext) StepSel() IStepSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepSelContext)
}

func (s *TypeSelContext) PartSel() IPartSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartSelContext)
}

func (s *TypeSelContext) FormSel() IFormSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormSelContext)
}

func (s *TypeSelContext) TypeAllSel() ITypeAllSelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAllSelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAllSelContext)
}

func (s *TypeSelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTypeSel(s)
	}
}

func (s *TypeSelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTypeSel(s)
	}
}

func (p *jossParser) TypeSel() (localctx ITypeSelContext) {
	this := p
	_ = this

	localctx = NewTypeSelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, jossParserRULE_typeSel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.StepSel()
		}

	case jossParserT__29:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.PartSel()
		}

	case jossParserT__30:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(262)
			p.FormSel()
		}

	case jossParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(263)
			p.TypeAllSel()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeAllSelContext is an interface to support dynamic dispatch.
type ITypeAllSelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeAllSelContext differentiates from other interfaces.
	IsTypeAllSelContext()
}

type TypeAllSelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAllSelContext() *TypeAllSelContext {
	var p = new(TypeAllSelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_typeAllSel
	return p
}

func (*TypeAllSelContext) IsTypeAllSelContext() {}

func NewTypeAllSelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAllSelContext {
	var p = new(TypeAllSelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_typeAllSel

	return p
}

func (s *TypeAllSelContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAllSelContext) TypeAllType() ITypeAllTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAllTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAllTypeContext)
}

func (s *TypeAllSelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAllSelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAllSelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTypeAllSel(s)
	}
}

func (s *TypeAllSelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTypeAllSel(s)
	}
}

func (p *jossParser) TypeAllSel() (localctx ITypeAllSelContext) {
	this := p
	_ = this

	localctx = NewTypeAllSelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, jossParserRULE_typeAllSel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.Match(jossParserT__4)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Match(jossParserT__4)
		}
		{
			p.SetState(268)
			p.TypeAllType()
		}

	}

	return localctx
}

// ITypeAllTypeContext is an interface to support dynamic dispatch.
type ITypeAllTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeAllTypeContext differentiates from other interfaces.
	IsTypeAllTypeContext()
}

type TypeAllTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAllTypeContext() *TypeAllTypeContext {
	var p = new(TypeAllTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_typeAllType
	return p
}

func (*TypeAllTypeContext) IsTypeAllTypeContext() {}

func NewTypeAllTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAllTypeContext {
	var p = new(TypeAllTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_typeAllType

	return p
}

func (s *TypeAllTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeAllTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAllTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAllTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTypeAllType(s)
	}
}

func (s *TypeAllTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTypeAllType(s)
	}
}

func (p *jossParser) TypeAllType() (localctx ITypeAllTypeContext) {
	this := p
	_ = this

	localctx = NewTypeAllTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, jossParserRULE_typeAllType)
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
		p.SetState(271)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<jossParserT__5)|(1<<jossParserT__6)|(1<<jossParserT__7)|(1<<jossParserT__21))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDeleteValCMDContext is an interface to support dynamic dispatch.
type IDeleteValCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteValCMDContext differentiates from other interfaces.
	IsDeleteValCMDContext()
}

type DeleteValCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteValCMDContext() *DeleteValCMDContext {
	var p = new(DeleteValCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_deleteValCMD
	return p
}

func (*DeleteValCMDContext) IsDeleteValCMDContext() {}

func NewDeleteValCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteValCMDContext {
	var p = new(DeleteValCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_deleteValCMD

	return p
}

func (s *DeleteValCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteValCMDContext) VarList() IVarListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarListContext)
}

func (s *DeleteValCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteValCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteValCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDeleteValCMD(s)
	}
}

func (s *DeleteValCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDeleteValCMD(s)
	}
}

func (p *jossParser) DeleteValCMD() (localctx IDeleteValCMDContext) {
	this := p
	_ = this

	localctx = NewDeleteValCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, jossParserRULE_deleteValCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(276)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserALPHA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.VarList()
		}

	case jossParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(jossParserT__4)
		}
		{
			p.SetState(275)
			p.Match(jossParserT__21)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFormCMDContext is an interface to support dynamic dispatch.
type IFormCMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormCMDContext differentiates from other interfaces.
	IsFormCMDContext()
}

type FormCMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormCMDContext() *FormCMDContext {
	var p = new(FormCMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_formCMD
	return p
}

func (*FormCMDContext) IsFormCMDContext() {}

func NewFormCMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormCMDContext {
	var p = new(FormCMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_formCMD

	return p
}

func (s *FormCMDContext) GetParser() antlr.Parser { return s.parser }

func (s *FormCMDContext) PARTNUMBER() antlr.TerminalNode {
	return s.GetToken(jossParserPARTNUMBER, 0)
}

func (s *FormCMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormCMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormCMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFormCMD(s)
	}
}

func (s *FormCMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFormCMD(s)
	}
}

func (p *jossParser) FormCMD() (localctx IFormCMDContext) {
	this := p
	_ = this

	localctx = NewFormCMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, jossParserRULE_formCMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(278)
		p.Match(jossParserT__22)
	}
	{
		p.SetState(279)
		p.Match(jossParserPARTNUMBER)
	}
	{
		p.SetState(280)
		p.Match(jossParserT__23)
	}

	return localctx
}

// IFormContentContext is an interface to support dynamic dispatch.
type IFormContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormContentContext differentiates from other interfaces.
	IsFormContentContext()
}

type FormContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormContentContext() *FormContentContext {
	var p = new(FormContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_formContent
	return p
}

func (*FormContentContext) IsFormContentContext() {}

func NewFormContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormContentContext {
	var p = new(FormContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_formContent

	return p
}

func (s *FormContentContext) GetParser() antlr.Parser { return s.parser }

func (s *FormContentContext) FormObject() IFormObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormObjectContext)
}

func (s *FormContentContext) FormContent() IFormContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormContentContext)
}

func (s *FormContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFormContent(s)
	}
}

func (s *FormContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFormContent(s)
	}
}

func (p *jossParser) FormContent() (localctx IFormContentContext) {
	this := p
	_ = this

	localctx = NewFormContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, jossParserRULE_formContent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.FormObject()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.FormObject()
		}
		{
			p.SetState(284)
			p.FormContent()
		}

	}

	return localctx
}

// IFormObjectContext is an interface to support dynamic dispatch.
type IFormObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormObjectContext differentiates from other interfaces.
	IsFormObjectContext()
}

type FormObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormObjectContext() *FormObjectContext {
	var p = new(FormObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_formObject
	return p
}

func (*FormObjectContext) IsFormObjectContext() {}

func NewFormObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormObjectContext {
	var p = new(FormObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_formObject

	return p
}

func (s *FormObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *FormObjectContext) FormPH() IFormPHContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormPHContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormPHContext)
}

func (s *FormObjectContext) STRING() antlr.TerminalNode {
	return s.GetToken(jossParserSTRING, 0)
}

func (s *FormObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFormObject(s)
	}
}

func (s *FormObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFormObject(s)
	}
}

func (p *jossParser) FormObject() (localctx IFormObjectContext) {
	this := p
	_ = this

	localctx = NewFormObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, jossParserRULE_formObject)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__24, jossParserT__25:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.FormPH()
		}

	case jossParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.Match(jossParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFormPHContext is an interface to support dynamic dispatch.
type IFormPHContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormPHContext differentiates from other interfaces.
	IsFormPHContext()
}

type FormPHContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormPHContext() *FormPHContext {
	var p = new(FormPHContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_formPH
	return p
}

func (*FormPHContext) IsFormPHContext() {}

func NewFormPHContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormPHContext {
	var p = new(FormPHContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_formPH

	return p
}

func (s *FormPHContext) GetParser() antlr.Parser { return s.parser }

func (s *FormPHContext) SciNotation() ISciNotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISciNotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISciNotationContext)
}

func (s *FormPHContext) FixedNotation() IFixedNotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFixedNotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFixedNotationContext)
}

func (s *FormPHContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormPHContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormPHContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFormPH(s)
	}
}

func (s *FormPHContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFormPH(s)
	}
}

func (p *jossParser) FormPH() (localctx IFormPHContext) {
	this := p
	_ = this

	localctx = NewFormPHContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, jossParserRULE_formPH)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(294)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__24:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(292)
			p.SciNotation()
		}

	case jossParserT__25:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			p.FixedNotation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISciNotationContext is an interface to support dynamic dispatch.
type ISciNotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSciNotationContext differentiates from other interfaces.
	IsSciNotationContext()
}

type SciNotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySciNotationContext() *SciNotationContext {
	var p = new(SciNotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_sciNotation
	return p
}

func (*SciNotationContext) IsSciNotationContext() {}

func NewSciNotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SciNotationContext {
	var p = new(SciNotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_sciNotation

	return p
}

func (s *SciNotationContext) GetParser() antlr.Parser { return s.parser }

func (s *SciNotationContext) Dot() IDotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDotContext)
}

func (s *SciNotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SciNotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SciNotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterSciNotation(s)
	}
}

func (s *SciNotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitSciNotation(s)
	}
}

func (p *jossParser) SciNotation() (localctx ISciNotationContext) {
	this := p
	_ = this

	localctx = NewSciNotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, jossParserRULE_sciNotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(296)
		p.Match(jossParserT__24)
	}
	{
		p.SetState(297)
		p.Dot()
	}

	return localctx
}

// IDotContext is an interface to support dynamic dispatch.
type IDotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDotContext differentiates from other interfaces.
	IsDotContext()
}

type DotContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotContext() *DotContext {
	var p = new(DotContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_dot
	return p
}

func (*DotContext) IsDotContext() {}

func NewDotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotContext {
	var p = new(DotContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_dot

	return p
}

func (s *DotContext) GetParser() antlr.Parser { return s.parser }

func (s *DotContext) Dot() IDotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDotContext)
}

func (s *DotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterDot(s)
	}
}

func (s *DotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitDot(s)
	}
}

func (p *jossParser) Dot() (localctx IDotContext) {
	this := p
	_ = this

	localctx = NewDotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, jossParserRULE_dot)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(299)
			p.Match(jossParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.Match(jossParserT__0)
		}
		{
			p.SetState(301)
			p.Dot()
		}

	}

	return localctx
}

// IFixedNotationContext is an interface to support dynamic dispatch.
type IFixedNotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFixedNotationContext differentiates from other interfaces.
	IsFixedNotationContext()
}

type FixedNotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFixedNotationContext() *FixedNotationContext {
	var p = new(FixedNotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_fixedNotation
	return p
}

func (*FixedNotationContext) IsFixedNotationContext() {}

func NewFixedNotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixedNotationContext {
	var p = new(FixedNotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_fixedNotation

	return p
}

func (s *FixedNotationContext) GetParser() antlr.Parser { return s.parser }

func (s *FixedNotationContext) AllUScore() []IUScoreContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUScoreContext)(nil)).Elem())
	var tst = make([]IUScoreContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUScoreContext)
		}
	}

	return tst
}

func (s *FixedNotationContext) UScore(i int) IUScoreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUScoreContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUScoreContext)
}

func (s *FixedNotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixedNotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixedNotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFixedNotation(s)
	}
}

func (s *FixedNotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFixedNotation(s)
	}
}

func (p *jossParser) FixedNotation() (localctx IFixedNotationContext) {
	this := p
	_ = this

	localctx = NewFixedNotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, jossParserRULE_fixedNotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(304)
		p.Match(jossParserT__25)
	}
	{
		p.SetState(305)
		p.UScore()
	}
	{
		p.SetState(306)
		p.Match(jossParserT__26)
	}
	{
		p.SetState(307)
		p.UScore()
	}

	return localctx
}

// IUScoreContext is an interface to support dynamic dispatch.
type IUScoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUScoreContext differentiates from other interfaces.
	IsUScoreContext()
}

type UScoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUScoreContext() *UScoreContext {
	var p = new(UScoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_uScore
	return p
}

func (*UScoreContext) IsUScoreContext() {}

func NewUScoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UScoreContext {
	var p = new(UScoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_uScore

	return p
}

func (s *UScoreContext) GetParser() antlr.Parser { return s.parser }
func (s *UScoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UScoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UScoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterUScore(s)
	}
}

func (s *UScoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitUScore(s)
	}
}

func (p *jossParser) UScore() (localctx IUScoreContext) {
	this := p
	_ = this

	localctx = NewUScoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, jossParserRULE_uScore)
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
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == jossParserT__27 {
		{
			p.SetState(309)
			p.Match(jossParserT__27)
		}

		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStepSelContext is an interface to support dynamic dispatch.
type IStepSelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepSelContext differentiates from other interfaces.
	IsStepSelContext()
}

type StepSelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepSelContext() *StepSelContext {
	var p = new(StepSelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_stepSel
	return p
}

func (*StepSelContext) IsStepSelContext() {}

func NewStepSelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepSelContext {
	var p = new(StepSelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_stepSel

	return p
}

func (s *StepSelContext) GetParser() antlr.Parser { return s.parser }

func (s *StepSelContext) STEPNUMBER() antlr.TerminalNode {
	return s.GetToken(jossParserSTEPNUMBER, 0)
}

func (s *StepSelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepSelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepSelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterStepSel(s)
	}
}

func (s *StepSelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitStepSel(s)
	}
}

func (p *jossParser) StepSel() (localctx IStepSelContext) {
	this := p
	_ = this

	localctx = NewStepSelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, jossParserRULE_stepSel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(314)
		p.Match(jossParserT__28)
	}
	{
		p.SetState(315)
		p.Match(jossParserSTEPNUMBER)
	}

	return localctx
}

// IPartSelContext is an interface to support dynamic dispatch.
type IPartSelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPartSelContext differentiates from other interfaces.
	IsPartSelContext()
}

type PartSelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartSelContext() *PartSelContext {
	var p = new(PartSelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_partSel
	return p
}

func (*PartSelContext) IsPartSelContext() {}

func NewPartSelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartSelContext {
	var p = new(PartSelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_partSel

	return p
}

func (s *PartSelContext) GetParser() antlr.Parser { return s.parser }

func (s *PartSelContext) PARTNUMBER() antlr.TerminalNode {
	return s.GetToken(jossParserPARTNUMBER, 0)
}

func (s *PartSelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartSelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartSelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterPartSel(s)
	}
}

func (s *PartSelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitPartSel(s)
	}
}

func (p *jossParser) PartSel() (localctx IPartSelContext) {
	this := p
	_ = this

	localctx = NewPartSelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, jossParserRULE_partSel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__29)
	}
	{
		p.SetState(318)
		p.Match(jossParserPARTNUMBER)
	}

	return localctx
}

// IFormSelContext is an interface to support dynamic dispatch.
type IFormSelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormSelContext differentiates from other interfaces.
	IsFormSelContext()
}

type FormSelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormSelContext() *FormSelContext {
	var p = new(FormSelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_formSel
	return p
}

func (*FormSelContext) IsFormSelContext() {}

func NewFormSelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormSelContext {
	var p = new(FormSelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_formSel

	return p
}

func (s *FormSelContext) GetParser() antlr.Parser { return s.parser }

func (s *FormSelContext) PARTNUMBER() antlr.TerminalNode {
	return s.GetToken(jossParserPARTNUMBER, 0)
}

func (s *FormSelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormSelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormSelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFormSel(s)
	}
}

func (s *FormSelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFormSel(s)
	}
}

func (p *jossParser) FormSel() (localctx IFormSelContext) {
	this := p
	_ = this

	localctx = NewFormSelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, jossParserRULE_formSel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(320)
		p.Match(jossParserT__30)
	}
	{
		p.SetState(321)
		p.Match(jossParserPARTNUMBER)
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) If_() IIf_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_Context)
}

func (s *ConditionalContext) For_() IFor_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_Context)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *jossParser) Conditional() (localctx IConditionalContext) {
	this := p
	_ = this

	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, jossParserRULE_conditional)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(325)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__31:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)
			p.If_()
		}

	case jossParserT__44:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.For_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = jossParserRULE_if_
	return p
}

func (*If_Context) IsIf_Context() {}

func NewIf_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_Context {
	var p = new(If_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_if_

	return p
}

func (s *If_Context) GetParser() antlr.Parser { return s.parser }

func (s *If_Context) BoolExp() IBoolExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExpContext)
}

func (s *If_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterIf_(s)
	}
}

func (s *If_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitIf_(s)
	}
}

func (p *jossParser) If_() (localctx IIf_Context) {
	this := p
	_ = this

	localctx = NewIf_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, jossParserRULE_if_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(327)
		p.Match(jossParserT__31)
	}
	{
		p.SetState(328)
		p.BoolExp()
	}

	return localctx
}

// IBoolExpContext is an interface to support dynamic dispatch.
type IBoolExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExpContext differentiates from other interfaces.
	IsBoolExpContext()
}

type BoolExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExpContext() *BoolExpContext {
	var p = new(BoolExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_boolExp
	return p
}

func (*BoolExpContext) IsBoolExpContext() {}

func NewBoolExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExpContext {
	var p = new(BoolExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_boolExp

	return p
}

func (s *BoolExpContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExpContext) BoolExp() IBoolExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExpContext)
}

func (s *BoolExpContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *BoolExpContext) BoolOp() IBoolOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolOpContext)
}

func (s *BoolExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterBoolExp(s)
	}
}

func (s *BoolExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitBoolExp(s)
	}
}

func (p *jossParser) BoolExp() (localctx IBoolExpContext) {
	this := p
	_ = this

	localctx = NewBoolExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, jossParserRULE_boolExp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Match(jossParserT__32)
		}
		{
			p.SetState(331)
			p.BoolExp()
		}
		{
			p.SetState(332)
			p.Match(jossParserT__33)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.Match(jossParserT__34)
		}
		{
			p.SetState(335)
			p.BoolExp()
		}
		{
			p.SetState(336)
			p.Match(jossParserT__35)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
			p.Comparison()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(339)
			p.Comparison()
		}
		{
			p.SetState(340)
			p.BoolOp()
		}
		{
			p.SetState(341)
			p.BoolExp()
		}

	}

	return localctx
}

// IBoolOpContext is an interface to support dynamic dispatch.
type IBoolOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolOpContext differentiates from other interfaces.
	IsBoolOpContext()
}

type BoolOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolOpContext() *BoolOpContext {
	var p = new(BoolOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_boolOp
	return p
}

func (*BoolOpContext) IsBoolOpContext() {}

func NewBoolOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolOpContext {
	var p = new(BoolOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_boolOp

	return p
}

func (s *BoolOpContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterBoolOp(s)
	}
}

func (s *BoolOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitBoolOp(s)
	}
}

func (p *jossParser) BoolOp() (localctx IBoolOpContext) {
	this := p
	_ = this

	localctx = NewBoolOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, jossParserRULE_boolOp)
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
		p.SetState(345)
		_la = p.GetTokenStream().LA(1)

		if !(_la == jossParserT__36 || _la == jossParserT__37) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) AllMathExp() []IMathExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathExpContext)(nil)).Elem())
	var tst = make([]IMathExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathExpContext)
		}
	}

	return tst
}

func (s *ComparisonContext) MathExp(i int) IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *ComparisonContext) BoolComp() IBoolCompContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolCompContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolCompContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (p *jossParser) Comparison() (localctx IComparisonContext) {
	this := p
	_ = this

	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, jossParserRULE_comparison)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(347)
		p.MathExp()
	}
	{
		p.SetState(348)
		p.BoolComp()
	}
	{
		p.SetState(349)
		p.MathExp()
	}

	return localctx
}

// IBoolCompContext is an interface to support dynamic dispatch.
type IBoolCompContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolCompContext differentiates from other interfaces.
	IsBoolCompContext()
}

type BoolCompContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolCompContext() *BoolCompContext {
	var p = new(BoolCompContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_boolComp
	return p
}

func (*BoolCompContext) IsBoolCompContext() {}

func NewBoolCompContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolCompContext {
	var p = new(BoolCompContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_boolComp

	return p
}

func (s *BoolCompContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolCompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolCompContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolCompContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterBoolComp(s)
	}
}

func (s *BoolCompContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitBoolComp(s)
	}
}

func (p *jossParser) BoolComp() (localctx IBoolCompContext) {
	this := p
	_ = this

	localctx = NewBoolCompContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, jossParserRULE_boolComp)
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
		p.SetState(351)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(jossParserT__38-39))|(1<<(jossParserT__39-39))|(1<<(jossParserT__40-39))|(1<<(jossParserT__41-39))|(1<<(jossParserT__42-39))|(1<<(jossParserT__43-39)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = jossParserRULE_for_
	return p
}

func (*For_Context) IsFor_Context() {}

func NewFor_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_Context {
	var p = new(For_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_for_

	return p
}

func (s *For_Context) GetParser() antlr.Parser { return s.parser }

func (s *For_Context) Range_() IRange_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_Context)
}

func (s *For_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFor_(s)
	}
}

func (s *For_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFor_(s)
	}
}

func (p *jossParser) For_() (localctx IFor_Context) {
	this := p
	_ = this

	localctx = NewFor_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, jossParserRULE_for_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__44)
	}
	{
		p.SetState(354)
		p.Range_()
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
	p.RuleIndex = jossParserRULE_range_
	return p
}

func (*Range_Context) IsRange_Context() {}

func NewRange_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_Context {
	var p = new(Range_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_range_

	return p
}

func (s *Range_Context) GetParser() antlr.Parser { return s.parser }

func (s *Range_Context) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Range_Context) RangeExp() IRangeExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeExpContext)
}

func (s *Range_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterRange_(s)
	}
}

func (s *Range_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitRange_(s)
	}
}

func (p *jossParser) Range_() (localctx IRange_Context) {
	this := p
	_ = this

	localctx = NewRange_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, jossParserRULE_range_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Variable()
	}
	{
		p.SetState(357)
		p.Match(jossParserT__38)
	}
	{
		p.SetState(358)
		p.RangeExp()
	}

	return localctx
}

// IRangeExpContext is an interface to support dynamic dispatch.
type IRangeExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeExpContext differentiates from other interfaces.
	IsRangeExpContext()
}

type RangeExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeExpContext() *RangeExpContext {
	var p = new(RangeExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_rangeExp
	return p
}

func (*RangeExpContext) IsRangeExpContext() {}

func NewRangeExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeExpContext {
	var p = new(RangeExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_rangeExp

	return p
}

func (s *RangeExpContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeExpContext) RangeVal() IRangeValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeValContext)
}

func (s *RangeExpContext) RangeExp() IRangeExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeExpContext)
}

func (s *RangeExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterRangeExp(s)
	}
}

func (s *RangeExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitRangeExp(s)
	}
}

func (p *jossParser) RangeExp() (localctx IRangeExpContext) {
	this := p
	_ = this

	localctx = NewRangeExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, jossParserRULE_rangeExp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)
			p.RangeVal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(361)
			p.RangeVal()
		}
		{
			p.SetState(362)
			p.Match(jossParserT__19)
		}
		{
			p.SetState(363)
			p.RangeExp()
		}

	}

	return localctx
}

// IRangeValContext is an interface to support dynamic dispatch.
type IRangeValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeValContext differentiates from other interfaces.
	IsRangeValContext()
}

type RangeValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeValContext() *RangeValContext {
	var p = new(RangeValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_rangeVal
	return p
}

func (*RangeValContext) IsRangeValContext() {}

func NewRangeValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeValContext {
	var p = new(RangeValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_rangeVal

	return p
}

func (s *RangeValContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeValContext) AllMathExp() []IMathExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathExpContext)(nil)).Elem())
	var tst = make([]IMathExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathExpContext)
		}
	}

	return tst
}

func (s *RangeValContext) MathExp(i int) IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *RangeValContext) RangeVal() IRangeValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeValContext)
}

func (s *RangeValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterRangeVal(s)
	}
}

func (s *RangeValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitRangeVal(s)
	}
}

func (p *jossParser) RangeVal() (localctx IRangeValContext) {
	this := p
	_ = this

	localctx = NewRangeValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, jossParserRULE_rangeVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.MathExp()
		}
		{
			p.SetState(368)
			p.Match(jossParserT__32)
		}
		{
			p.SetState(369)
			p.MathExp()
		}
		{
			p.SetState(370)
			p.Match(jossParserT__33)
		}
		{
			p.SetState(371)
			p.RangeVal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.MathExp()
		}

	}

	return localctx
}

// IMathExpContext is an interface to support dynamic dispatch.
type IMathExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathExpContext differentiates from other interfaces.
	IsMathExpContext()
}

type MathExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathExpContext() *MathExpContext {
	var p = new(MathExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_mathExp
	return p
}

func (*MathExpContext) IsMathExpContext() {}

func NewMathExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathExpContext {
	var p = new(MathExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_mathExp

	return p
}

func (s *MathExpContext) GetParser() antlr.Parser { return s.parser }

func (s *MathExpContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *MathExpContext) ADDOP() antlr.TerminalNode {
	return s.GetToken(jossParserADDOP, 0)
}

func (s *MathExpContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *MathExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterMathExp(s)
	}
}

func (s *MathExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitMathExp(s)
	}
}

func (p *jossParser) MathExp() (localctx IMathExpContext) {
	this := p
	_ = this

	localctx = NewMathExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, jossParserRULE_mathExp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(376)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(377)
			p.Term()
		}
		{
			p.SetState(378)
			p.Match(jossParserADDOP)
		}
		{
			p.SetState(379)
			p.MathExp()
		}

	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) MULOP() antlr.TerminalNode {
	return s.GetToken(jossParserMULOP, 0)
}

func (s *TermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *jossParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, jossParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.Factor()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)
			p.Factor()
		}
		{
			p.SetState(385)
			p.Match(jossParserMULOP)
		}
		{
			p.SetState(386)
			p.Term()
		}

	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FactorContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *jossParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, jossParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	switch p.GetTokenStream().LA(1) {
	case jossParserT__32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(390)
			p.Match(jossParserT__32)
		}
		{
			p.SetState(391)
			p.MathExp()
		}
		{
			p.SetState(392)
			p.Match(jossParserT__33)
		}

	case jossParserT__34:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(394)
			p.Match(jossParserT__34)
		}
		{
			p.SetState(395)
			p.MathExp()
		}
		{
			p.SetState(396)
			p.Match(jossParserT__35)
		}

	case jossParserT__45, jossParserT__46, jossParserT__47, jossParserT__48, jossParserT__49, jossParserT__50, jossParserT__51, jossParserT__52, jossParserT__53, jossParserT__54, jossParserT__55, jossParserT__56, jossParserALPHA, jossParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(398)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *jossParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, jossParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(401)
		p.Variable()
	}
	{
		p.SetState(402)
		p.Match(jossParserT__38)
	}
	{
		p.SetState(403)
		p.MathExp()
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
	p.RuleIndex = jossParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(jossParserNUMBER, 0)
}

func (s *ValueContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ValueContext) Function_() IFunction_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_Context)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *jossParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, jossParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(408)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.Match(jossParserNUMBER)
		}

	case jossParserALPHA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
			p.Variable()
		}

	case jossParserT__45, jossParserT__46, jossParserT__47, jossParserT__48, jossParserT__49, jossParserT__50, jossParserT__51, jossParserT__52, jossParserT__53, jossParserT__54, jossParserT__55, jossParserT__56:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(407)
			p.Function_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunction_Context is an interface to support dynamic dispatch.
type IFunction_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_Context differentiates from other interfaces.
	IsFunction_Context()
}

type Function_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_Context() *Function_Context {
	var p = new(Function_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_function_
	return p
}

func (*Function_Context) IsFunction_Context() {}

func NewFunction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_Context {
	var p = new(Function_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_function_

	return p
}

func (s *Function_Context) GetParser() antlr.Parser { return s.parser }

func (s *Function_Context) FuncSqrt() IFuncSqrtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSqrtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSqrtContext)
}

func (s *Function_Context) FuncLog() IFuncLogContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncLogContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncLogContext)
}

func (s *Function_Context) FuncExp() IFuncExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExpContext)
}

func (s *Function_Context) FuncSin() IFuncSinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSinContext)
}

func (s *Function_Context) FuncCos() IFuncCosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCosContext)
}

func (s *Function_Context) FuncIp() IFuncIpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncIpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncIpContext)
}

func (s *Function_Context) FuncFp() IFuncFpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncFpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncFpContext)
}

func (s *Function_Context) FuncDp() IFuncDpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDpContext)
}

func (s *Function_Context) FuncXp() IFuncXpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncXpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncXpContext)
}

func (s *Function_Context) FuncSgn() IFuncSgnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSgnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSgnContext)
}

func (s *Function_Context) FuncMax() IFuncMaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncMaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncMaxContext)
}

func (s *Function_Context) FuncMin() IFuncMinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncMinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncMinContext)
}

func (s *Function_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFunction_(s)
	}
}

func (s *Function_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFunction_(s)
	}
}

func (p *jossParser) Function_() (localctx IFunction_Context) {
	this := p
	_ = this

	localctx = NewFunction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, jossParserRULE_function_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(422)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jossParserT__45:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(410)
			p.FuncSqrt()
		}

	case jossParserT__46:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(411)
			p.FuncLog()
		}

	case jossParserT__47:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(412)
			p.FuncExp()
		}

	case jossParserT__48:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(413)
			p.FuncSin()
		}

	case jossParserT__49:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(414)
			p.FuncCos()
		}

	case jossParserT__50:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(415)
			p.FuncIp()
		}

	case jossParserT__51:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(416)
			p.FuncFp()
		}

	case jossParserT__52:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(417)
			p.FuncDp()
		}

	case jossParserT__53:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(418)
			p.FuncXp()
		}

	case jossParserT__54:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(419)
			p.FuncSgn()
		}

	case jossParserT__55:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(420)
			p.FuncMax()
		}

	case jossParserT__56:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(421)
			p.FuncMin()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncSqrtContext is an interface to support dynamic dispatch.
type IFuncSqrtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSqrtContext differentiates from other interfaces.
	IsFuncSqrtContext()
}

type FuncSqrtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSqrtContext() *FuncSqrtContext {
	var p = new(FuncSqrtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcSqrt
	return p
}

func (*FuncSqrtContext) IsFuncSqrtContext() {}

func NewFuncSqrtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSqrtContext {
	var p = new(FuncSqrtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcSqrt

	return p
}

func (s *FuncSqrtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSqrtContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncSqrtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSqrtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSqrtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncSqrt(s)
	}
}

func (s *FuncSqrtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncSqrt(s)
	}
}

func (p *jossParser) FuncSqrt() (localctx IFuncSqrtContext) {
	this := p
	_ = this

	localctx = NewFuncSqrtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, jossParserRULE_funcSqrt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(424)
		p.Match(jossParserT__45)
	}
	{
		p.SetState(425)
		p.MathExp()
	}
	{
		p.SetState(426)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncLogContext is an interface to support dynamic dispatch.
type IFuncLogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncLogContext differentiates from other interfaces.
	IsFuncLogContext()
}

type FuncLogContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncLogContext() *FuncLogContext {
	var p = new(FuncLogContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcLog
	return p
}

func (*FuncLogContext) IsFuncLogContext() {}

func NewFuncLogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncLogContext {
	var p = new(FuncLogContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcLog

	return p
}

func (s *FuncLogContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncLogContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncLogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncLogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncLogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncLog(s)
	}
}

func (s *FuncLogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncLog(s)
	}
}

func (p *jossParser) FuncLog() (localctx IFuncLogContext) {
	this := p
	_ = this

	localctx = NewFuncLogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, jossParserRULE_funcLog)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__46)
	}
	{
		p.SetState(429)
		p.MathExp()
	}
	{
		p.SetState(430)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncExpContext is an interface to support dynamic dispatch.
type IFuncExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExpContext differentiates from other interfaces.
	IsFuncExpContext()
}

type FuncExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExpContext() *FuncExpContext {
	var p = new(FuncExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcExp
	return p
}

func (*FuncExpContext) IsFuncExpContext() {}

func NewFuncExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExpContext {
	var p = new(FuncExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcExp

	return p
}

func (s *FuncExpContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExpContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncExp(s)
	}
}

func (s *FuncExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncExp(s)
	}
}

func (p *jossParser) FuncExp() (localctx IFuncExpContext) {
	this := p
	_ = this

	localctx = NewFuncExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, jossParserRULE_funcExp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(432)
		p.Match(jossParserT__47)
	}
	{
		p.SetState(433)
		p.MathExp()
	}
	{
		p.SetState(434)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncSinContext is an interface to support dynamic dispatch.
type IFuncSinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSinContext differentiates from other interfaces.
	IsFuncSinContext()
}

type FuncSinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSinContext() *FuncSinContext {
	var p = new(FuncSinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcSin
	return p
}

func (*FuncSinContext) IsFuncSinContext() {}

func NewFuncSinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSinContext {
	var p = new(FuncSinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcSin

	return p
}

func (s *FuncSinContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSinContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncSinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncSin(s)
	}
}

func (s *FuncSinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncSin(s)
	}
}

func (p *jossParser) FuncSin() (localctx IFuncSinContext) {
	this := p
	_ = this

	localctx = NewFuncSinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, jossParserRULE_funcSin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(436)
		p.Match(jossParserT__48)
	}
	{
		p.SetState(437)
		p.MathExp()
	}
	{
		p.SetState(438)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncCosContext is an interface to support dynamic dispatch.
type IFuncCosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCosContext differentiates from other interfaces.
	IsFuncCosContext()
}

type FuncCosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCosContext() *FuncCosContext {
	var p = new(FuncCosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcCos
	return p
}

func (*FuncCosContext) IsFuncCosContext() {}

func NewFuncCosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCosContext {
	var p = new(FuncCosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcCos

	return p
}

func (s *FuncCosContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCosContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncCosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncCos(s)
	}
}

func (s *FuncCosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncCos(s)
	}
}

func (p *jossParser) FuncCos() (localctx IFuncCosContext) {
	this := p
	_ = this

	localctx = NewFuncCosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, jossParserRULE_funcCos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(440)
		p.Match(jossParserT__49)
	}
	{
		p.SetState(441)
		p.MathExp()
	}
	{
		p.SetState(442)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncIpContext is an interface to support dynamic dispatch.
type IFuncIpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncIpContext differentiates from other interfaces.
	IsFuncIpContext()
}

type FuncIpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncIpContext() *FuncIpContext {
	var p = new(FuncIpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcIp
	return p
}

func (*FuncIpContext) IsFuncIpContext() {}

func NewFuncIpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncIpContext {
	var p = new(FuncIpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcIp

	return p
}

func (s *FuncIpContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncIpContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncIpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncIpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncIpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncIp(s)
	}
}

func (s *FuncIpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncIp(s)
	}
}

func (p *jossParser) FuncIp() (localctx IFuncIpContext) {
	this := p
	_ = this

	localctx = NewFuncIpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, jossParserRULE_funcIp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__50)
	}
	{
		p.SetState(445)
		p.MathExp()
	}
	{
		p.SetState(446)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncFpContext is an interface to support dynamic dispatch.
type IFuncFpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncFpContext differentiates from other interfaces.
	IsFuncFpContext()
}

type FuncFpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncFpContext() *FuncFpContext {
	var p = new(FuncFpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcFp
	return p
}

func (*FuncFpContext) IsFuncFpContext() {}

func NewFuncFpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFpContext {
	var p = new(FuncFpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcFp

	return p
}

func (s *FuncFpContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFpContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncFpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncFpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncFp(s)
	}
}

func (s *FuncFpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncFp(s)
	}
}

func (p *jossParser) FuncFp() (localctx IFuncFpContext) {
	this := p
	_ = this

	localctx = NewFuncFpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, jossParserRULE_funcFp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(448)
		p.Match(jossParserT__51)
	}
	{
		p.SetState(449)
		p.MathExp()
	}
	{
		p.SetState(450)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncDpContext is an interface to support dynamic dispatch.
type IFuncDpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDpContext differentiates from other interfaces.
	IsFuncDpContext()
}

type FuncDpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDpContext() *FuncDpContext {
	var p = new(FuncDpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcDp
	return p
}

func (*FuncDpContext) IsFuncDpContext() {}

func NewFuncDpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDpContext {
	var p = new(FuncDpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcDp

	return p
}

func (s *FuncDpContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDpContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncDpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncDp(s)
	}
}

func (s *FuncDpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncDp(s)
	}
}

func (p *jossParser) FuncDp() (localctx IFuncDpContext) {
	this := p
	_ = this

	localctx = NewFuncDpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, jossParserRULE_funcDp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(jossParserT__52)
	}
	{
		p.SetState(453)
		p.MathExp()
	}
	{
		p.SetState(454)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncXpContext is an interface to support dynamic dispatch.
type IFuncXpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncXpContext differentiates from other interfaces.
	IsFuncXpContext()
}

type FuncXpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncXpContext() *FuncXpContext {
	var p = new(FuncXpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcXp
	return p
}

func (*FuncXpContext) IsFuncXpContext() {}

func NewFuncXpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncXpContext {
	var p = new(FuncXpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcXp

	return p
}

func (s *FuncXpContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncXpContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncXpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncXpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncXpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncXp(s)
	}
}

func (s *FuncXpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncXp(s)
	}
}

func (p *jossParser) FuncXp() (localctx IFuncXpContext) {
	this := p
	_ = this

	localctx = NewFuncXpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, jossParserRULE_funcXp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(456)
		p.Match(jossParserT__53)
	}
	{
		p.SetState(457)
		p.MathExp()
	}
	{
		p.SetState(458)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncSgnContext is an interface to support dynamic dispatch.
type IFuncSgnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSgnContext differentiates from other interfaces.
	IsFuncSgnContext()
}

type FuncSgnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSgnContext() *FuncSgnContext {
	var p = new(FuncSgnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcSgn
	return p
}

func (*FuncSgnContext) IsFuncSgnContext() {}

func NewFuncSgnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSgnContext {
	var p = new(FuncSgnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcSgn

	return p
}

func (s *FuncSgnContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSgnContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncSgnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSgnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSgnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncSgn(s)
	}
}

func (s *FuncSgnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncSgn(s)
	}
}

func (p *jossParser) FuncSgn() (localctx IFuncSgnContext) {
	this := p
	_ = this

	localctx = NewFuncSgnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, jossParserRULE_funcSgn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(460)
		p.Match(jossParserT__54)
	}
	{
		p.SetState(461)
		p.MathExp()
	}
	{
		p.SetState(462)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncMaxContext is an interface to support dynamic dispatch.
type IFuncMaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncMaxContext differentiates from other interfaces.
	IsFuncMaxContext()
}

type FuncMaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncMaxContext() *FuncMaxContext {
	var p = new(FuncMaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcMax
	return p
}

func (*FuncMaxContext) IsFuncMaxContext() {}

func NewFuncMaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncMaxContext {
	var p = new(FuncMaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcMax

	return p
}

func (s *FuncMaxContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncMaxContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncMaxContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FuncMaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncMaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncMaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncMax(s)
	}
}

func (s *FuncMaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncMax(s)
	}
}

func (p *jossParser) FuncMax() (localctx IFuncMaxContext) {
	this := p
	_ = this

	localctx = NewFuncMaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, jossParserRULE_funcMax)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(464)
		p.Match(jossParserT__55)
	}
	{
		p.SetState(465)
		p.MathExp()
	}
	{
		p.SetState(466)
		p.Match(jossParserT__19)
	}
	{
		p.SetState(467)
		p.ArgList()
	}
	{
		p.SetState(468)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IFuncMinContext is an interface to support dynamic dispatch.
type IFuncMinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncMinContext differentiates from other interfaces.
	IsFuncMinContext()
}

type FuncMinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncMinContext() *FuncMinContext {
	var p = new(FuncMinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_funcMin
	return p
}

func (*FuncMinContext) IsFuncMinContext() {}

func NewFuncMinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncMinContext {
	var p = new(FuncMinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_funcMin

	return p
}

func (s *FuncMinContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncMinContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *FuncMinContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FuncMinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncMinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncMinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterFuncMin(s)
	}
}

func (s *FuncMinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitFuncMin(s)
	}
}

func (p *jossParser) FuncMin() (localctx IFuncMinContext) {
	this := p
	_ = this

	localctx = NewFuncMinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, jossParserRULE_funcMin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(470)
		p.Match(jossParserT__56)
	}
	{
		p.SetState(471)
		p.MathExp()
	}
	{
		p.SetState(472)
		p.Match(jossParserT__19)
	}
	{
		p.SetState(473)
		p.ArgList()
	}
	{
		p.SetState(474)
		p.Match(jossParserT__33)
	}

	return localctx
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_argList
	return p
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *ArgListContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (p *jossParser) ArgList() (localctx IArgListContext) {
	this := p
	_ = this

	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, jossParserRULE_argList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.MathExp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(477)
			p.MathExp()
		}
		{
			p.SetState(478)
			p.Match(jossParserT__19)
		}
		{
			p.SetState(479)
			p.ArgList()
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
	p.RuleIndex = jossParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(jossParserALPHA, 0)
}

func (s *VariableContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *jossParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, jossParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)
			p.Match(jossParserALPHA)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(484)
			p.Match(jossParserALPHA)
		}
		{
			p.SetState(485)
			p.Match(jossParserT__32)
		}
		{
			p.SetState(486)
			p.MathExp()
		}
		{
			p.SetState(487)
			p.Match(jossParserT__33)
		}

	}

	return localctx
}

// IVarListContext is an interface to support dynamic dispatch.
type IVarListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarListContext differentiates from other interfaces.
	IsVarListContext()
}

type VarListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarListContext() *VarListContext {
	var p = new(VarListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jossParserRULE_varList
	return p
}

func (*VarListContext) IsVarListContext() {}

func NewVarListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarListContext {
	var p = new(VarListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jossParserRULE_varList

	return p
}

func (s *VarListContext) GetParser() antlr.Parser { return s.parser }

func (s *VarListContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VarListContext) VarList() IVarListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarListContext)
}

func (s *VarListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.EnterVarList(s)
	}
}

func (s *VarListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jossListener); ok {
		listenerT.ExitVarList(s)
	}
}

func (p *jossParser) VarList() (localctx IVarListContext) {
	this := p
	_ = this

	localctx = NewVarListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, jossParserRULE_varList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(491)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(492)
			p.Variable()
		}
		{
			p.SetState(493)
			p.Match(jossParserT__19)
		}
		{
			p.SetState(494)
			p.VarList()
		}

	}

	return localctx
}
