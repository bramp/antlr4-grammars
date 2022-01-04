// Code generated from CapnProto.g4 by ANTLR 4.9.3. DO NOT EDIT.

package capnproto // CapnProto
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 433,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 3, 2,
	3, 2, 7, 2, 69, 10, 2, 12, 2, 14, 2, 72, 11, 2, 3, 2, 5, 2, 75, 10, 2,
	3, 2, 7, 2, 78, 10, 2, 12, 2, 14, 2, 81, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 91, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	97, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 115, 10, 6, 3, 7, 3, 7, 3, 7,
	5, 7, 120, 10, 7, 3, 7, 3, 7, 7, 7, 124, 10, 7, 12, 7, 14, 7, 127, 11,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 5, 8, 141, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 150,
	10, 9, 3, 9, 3, 9, 7, 9, 154, 10, 9, 12, 9, 14, 9, 157, 11, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 168, 10, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 176, 10, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 5, 12, 182, 10, 12, 3, 12, 3, 12, 5, 12, 186, 10, 12,
	3, 13, 3, 13, 3, 13, 5, 13, 191, 10, 13, 3, 13, 3, 13, 3, 13, 5, 13, 196,
	10, 13, 7, 13, 198, 10, 13, 12, 13, 14, 13, 201, 11, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 5, 14, 208, 10, 14, 3, 14, 3, 14, 7, 14, 212, 10,
	14, 12, 14, 14, 14, 215, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 5,
	15, 222, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 5, 16,
	231, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 237, 10, 17, 3, 17, 3,
	17, 3, 17, 7, 17, 242, 10, 17, 12, 17, 14, 17, 245, 11, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 7, 18, 252, 10, 18, 12, 18, 14, 18, 255, 11, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 263, 10, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 7, 20, 269, 10, 20, 12, 20, 14, 20, 272, 11, 20, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 21, 5, 21, 279, 10, 21, 3, 22, 3, 22, 5, 22, 283,
	10, 22, 3, 22, 5, 22, 286, 10, 22, 3, 22, 3, 22, 5, 22, 290, 10, 22, 3,
	22, 3, 22, 3, 22, 5, 22, 295, 10, 22, 5, 22, 297, 10, 22, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 305, 10, 23, 12, 23, 14, 23, 308, 11,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 318,
	10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 326, 10, 24, 7,
	24, 328, 10, 24, 12, 24, 14, 24, 331, 11, 24, 5, 24, 333, 10, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 25, 5, 25, 340, 10, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 28, 5, 28, 359, 10, 28, 3, 28, 5, 28, 362, 10, 28, 3,
	28, 3, 28, 3, 28, 5, 28, 367, 10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 378, 10, 28, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 388, 10, 29, 12, 29, 14, 29, 391,
	11, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 399, 10, 30, 12,
	30, 14, 30, 402, 11, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 416, 10, 32, 3, 33, 3, 33, 3,
	33, 3, 33, 7, 33, 422, 10, 33, 12, 33, 14, 33, 425, 11, 33, 3, 33, 3, 33,
	5, 33, 429, 10, 33, 3, 33, 3, 33, 3, 33, 2, 2, 34, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 54, 56, 58, 60, 62, 64, 2, 2, 2, 476, 2, 66, 3, 2, 2, 2, 4, 84,
	3, 2, 2, 2, 6, 87, 3, 2, 2, 2, 8, 100, 3, 2, 2, 2, 10, 114, 3, 2, 2, 2,
	12, 116, 3, 2, 2, 2, 14, 140, 3, 2, 2, 2, 16, 142, 3, 2, 2, 2, 18, 167,
	3, 2, 2, 2, 20, 169, 3, 2, 2, 2, 22, 179, 3, 2, 2, 2, 24, 187, 3, 2, 2,
	2, 26, 204, 3, 2, 2, 2, 28, 218, 3, 2, 2, 2, 30, 227, 3, 2, 2, 2, 32, 234,
	3, 2, 2, 2, 34, 248, 3, 2, 2, 2, 36, 262, 3, 2, 2, 2, 38, 264, 3, 2, 2,
	2, 40, 278, 3, 2, 2, 2, 42, 280, 3, 2, 2, 2, 44, 300, 3, 2, 2, 2, 46, 311,
	3, 2, 2, 2, 48, 336, 3, 2, 2, 2, 50, 345, 3, 2, 2, 2, 52, 349, 3, 2, 2,
	2, 54, 377, 3, 2, 2, 2, 56, 379, 3, 2, 2, 2, 58, 394, 3, 2, 2, 2, 60, 405,
	3, 2, 2, 2, 62, 415, 3, 2, 2, 2, 64, 417, 3, 2, 2, 2, 66, 70, 5, 4, 3,
	2, 67, 69, 5, 6, 4, 2, 68, 67, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68,
	3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2,
	73, 75, 5, 8, 5, 2, 74, 73, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 79, 3,
	2, 2, 2, 76, 78, 5, 10, 6, 2, 77, 76, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79,
	77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81, 79, 3, 2, 2,
	2, 82, 83, 7, 2, 2, 3, 83, 3, 3, 2, 2, 2, 84, 85, 7, 36, 2, 2, 85, 86,
	7, 3, 2, 2, 86, 5, 3, 2, 2, 2, 87, 90, 7, 4, 2, 2, 88, 89, 7, 39, 2, 2,
	89, 91, 7, 5, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 3,
	2, 2, 2, 92, 93, 7, 6, 2, 2, 93, 96, 7, 32, 2, 2, 94, 95, 7, 7, 2, 2, 95,
	97, 7, 39, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 3, 2,
	2, 2, 98, 99, 7, 3, 2, 2, 99, 7, 3, 2, 2, 2, 100, 101, 7, 8, 2, 2, 101,
	102, 7, 39, 2, 2, 102, 103, 7, 9, 2, 2, 103, 104, 7, 10, 2, 2, 104, 105,
	7, 32, 2, 2, 105, 106, 7, 11, 2, 2, 106, 107, 7, 3, 2, 2, 107, 9, 3, 2,
	2, 2, 108, 115, 5, 12, 7, 2, 109, 115, 5, 16, 9, 2, 110, 115, 5, 42, 22,
	2, 111, 115, 5, 48, 25, 2, 112, 115, 5, 52, 27, 2, 113, 115, 5, 26, 14,
	2, 114, 108, 3, 2, 2, 2, 114, 109, 3, 2, 2, 2, 114, 110, 3, 2, 2, 2, 114,
	111, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 113, 3, 2, 2, 2, 115, 11, 3,
	2, 2, 2, 116, 117, 7, 12, 2, 2, 117, 119, 5, 22, 12, 2, 118, 120, 5, 28,
	15, 2, 119, 118, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2,
	121, 125, 7, 13, 2, 2, 122, 124, 5, 14, 8, 2, 123, 122, 3, 2, 2, 2, 124,
	127, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 128,
	3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 128, 129, 7, 14, 2, 2, 129, 13, 3, 2,
	2, 2, 130, 141, 5, 20, 11, 2, 131, 141, 5, 26, 14, 2, 132, 141, 5, 32,
	17, 2, 133, 141, 5, 34, 18, 2, 134, 141, 5, 16, 9, 2, 135, 141, 5, 48,
	25, 2, 136, 141, 5, 12, 7, 2, 137, 141, 5, 38, 20, 2, 138, 141, 5, 52,
	27, 2, 139, 141, 5, 64, 33, 2, 140, 130, 3, 2, 2, 2, 140, 131, 3, 2, 2,
	2, 140, 132, 3, 2, 2, 2, 140, 133, 3, 2, 2, 2, 140, 134, 3, 2, 2, 2, 140,
	135, 3, 2, 2, 2, 140, 136, 3, 2, 2, 2, 140, 137, 3, 2, 2, 2, 140, 138,
	3, 2, 2, 2, 140, 139, 3, 2, 2, 2, 141, 15, 3, 2, 2, 2, 142, 143, 7, 15,
	2, 2, 143, 149, 5, 22, 12, 2, 144, 145, 7, 16, 2, 2, 145, 146, 7, 10, 2,
	2, 146, 147, 5, 22, 12, 2, 147, 148, 7, 11, 2, 2, 148, 150, 3, 2, 2, 2,
	149, 144, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151,
	155, 7, 13, 2, 2, 152, 154, 5, 18, 10, 2, 153, 152, 3, 2, 2, 2, 154, 157,
	3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 158, 3, 2,
	2, 2, 157, 155, 3, 2, 2, 2, 158, 159, 7, 14, 2, 2, 159, 17, 3, 2, 2, 2,
	160, 168, 5, 20, 11, 2, 161, 168, 5, 26, 14, 2, 162, 168, 5, 32, 17, 2,
	163, 168, 5, 34, 18, 2, 164, 168, 5, 16, 9, 2, 165, 168, 5, 12, 7, 2, 166,
	168, 5, 42, 22, 2, 167, 160, 3, 2, 2, 2, 167, 161, 3, 2, 2, 2, 167, 162,
	3, 2, 2, 2, 167, 163, 3, 2, 2, 2, 167, 164, 3, 2, 2, 2, 167, 165, 3, 2,
	2, 2, 167, 166, 3, 2, 2, 2, 168, 19, 3, 2, 2, 2, 169, 170, 7, 39, 2, 2,
	170, 171, 7, 31, 2, 2, 171, 172, 7, 17, 2, 2, 172, 175, 5, 22, 12, 2, 173,
	174, 7, 5, 2, 2, 174, 176, 5, 54, 28, 2, 175, 173, 3, 2, 2, 2, 175, 176,
	3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 178, 7, 3, 2, 2, 178, 21, 3, 2,
	2, 2, 179, 181, 7, 39, 2, 2, 180, 182, 5, 24, 13, 2, 181, 180, 3, 2, 2,
	2, 181, 182, 3, 2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 184, 7, 7, 2, 2, 184,
	186, 5, 22, 12, 2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 23,
	3, 2, 2, 2, 187, 188, 7, 10, 2, 2, 188, 190, 5, 22, 12, 2, 189, 191, 5,
	24, 13, 2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 199, 3, 2,
	2, 2, 192, 193, 7, 18, 2, 2, 193, 195, 5, 22, 12, 2, 194, 196, 5, 24, 13,
	2, 195, 194, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 198, 3, 2, 2, 2, 197,
	192, 3, 2, 2, 2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 200,
	3, 2, 2, 2, 200, 202, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202, 203, 7, 11,
	2, 2, 203, 25, 3, 2, 2, 2, 204, 205, 7, 19, 2, 2, 205, 207, 7, 39, 2, 2,
	206, 208, 5, 28, 15, 2, 207, 206, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208,
	209, 3, 2, 2, 2, 209, 213, 7, 13, 2, 2, 210, 212, 5, 30, 16, 2, 211, 210,
	3, 2, 2, 2, 212, 215, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2,
	2, 2, 214, 216, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 216, 217, 7, 14, 2, 2,
	217, 27, 3, 2, 2, 2, 218, 219, 7, 8, 2, 2, 219, 221, 5, 22, 12, 2, 220,
	222, 7, 20, 2, 2, 221, 220, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 223,
	3, 2, 2, 2, 223, 224, 7, 10, 2, 2, 224, 225, 7, 32, 2, 2, 225, 226, 7,
	11, 2, 2, 226, 29, 3, 2, 2, 2, 227, 228, 7, 39, 2, 2, 228, 230, 7, 31,
	2, 2, 229, 231, 5, 28, 15, 2, 230, 229, 3, 2, 2, 2, 230, 231, 3, 2, 2,
	2, 231, 232, 3, 2, 2, 2, 232, 233, 7, 3, 2, 2, 233, 31, 3, 2, 2, 2, 234,
	236, 7, 39, 2, 2, 235, 237, 7, 31, 2, 2, 236, 235, 3, 2, 2, 2, 236, 237,
	3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 239, 7, 21, 2, 2, 239, 243, 7, 13,
	2, 2, 240, 242, 5, 36, 19, 2, 241, 240, 3, 2, 2, 2, 242, 245, 3, 2, 2,
	2, 243, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 246, 3, 2, 2, 2, 245,
	243, 3, 2, 2, 2, 246, 247, 7, 14, 2, 2, 247, 33, 3, 2, 2, 2, 248, 249,
	7, 22, 2, 2, 249, 253, 7, 13, 2, 2, 250, 252, 5, 36, 19, 2, 251, 250, 3,
	2, 2, 2, 252, 255, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2,
	2, 254, 256, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 256, 257, 7, 14, 2, 2, 257,
	35, 3, 2, 2, 2, 258, 263, 5, 20, 11, 2, 259, 263, 5, 38, 20, 2, 260, 263,
	5, 34, 18, 2, 261, 263, 5, 32, 17, 2, 262, 258, 3, 2, 2, 2, 262, 259, 3,
	2, 2, 2, 262, 260, 3, 2, 2, 2, 262, 261, 3, 2, 2, 2, 263, 37, 3, 2, 2,
	2, 264, 265, 7, 39, 2, 2, 265, 266, 7, 23, 2, 2, 266, 270, 7, 13, 2, 2,
	267, 269, 5, 40, 21, 2, 268, 267, 3, 2, 2, 2, 269, 272, 3, 2, 2, 2, 270,
	268, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 273, 3, 2, 2, 2, 272, 270,
	3, 2, 2, 2, 273, 274, 7, 14, 2, 2, 274, 39, 3, 2, 2, 2, 275, 279, 5, 20,
	11, 2, 276, 279, 5, 34, 18, 2, 277, 279, 5, 32, 17, 2, 278, 275, 3, 2,
	2, 2, 278, 276, 3, 2, 2, 2, 278, 277, 3, 2, 2, 2, 279, 41, 3, 2, 2, 2,
	280, 282, 7, 39, 2, 2, 281, 283, 7, 31, 2, 2, 282, 281, 3, 2, 2, 2, 282,
	283, 3, 2, 2, 2, 283, 285, 3, 2, 2, 2, 284, 286, 5, 44, 23, 2, 285, 284,
	3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 289, 3, 2, 2, 2, 287, 290, 5, 46,
	24, 2, 288, 290, 5, 22, 12, 2, 289, 287, 3, 2, 2, 2, 289, 288, 3, 2, 2,
	2, 290, 296, 3, 2, 2, 2, 291, 294, 7, 24, 2, 2, 292, 295, 5, 46, 24, 2,
	293, 295, 5, 22, 12, 2, 294, 292, 3, 2, 2, 2, 294, 293, 3, 2, 2, 2, 295,
	297, 3, 2, 2, 2, 296, 291, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298,
	3, 2, 2, 2, 298, 299, 7, 3, 2, 2, 299, 43, 3, 2, 2, 2, 300, 301, 7, 25,
	2, 2, 301, 306, 7, 39, 2, 2, 302, 303, 7, 18, 2, 2, 303, 305, 7, 39, 2,
	2, 304, 302, 3, 2, 2, 2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306,
	307, 3, 2, 2, 2, 307, 309, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 310,
	7, 26, 2, 2, 310, 45, 3, 2, 2, 2, 311, 332, 7, 10, 2, 2, 312, 313, 7, 39,
	2, 2, 313, 314, 7, 17, 2, 2, 314, 317, 5, 22, 12, 2, 315, 316, 7, 5, 2,
	2, 316, 318, 5, 54, 28, 2, 317, 315, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2,
	318, 329, 3, 2, 2, 2, 319, 320, 7, 18, 2, 2, 320, 321, 7, 39, 2, 2, 321,
	322, 7, 17, 2, 2, 322, 325, 5, 22, 12, 2, 323, 324, 7, 5, 2, 2, 324, 326,
	5, 54, 28, 2, 325, 323, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 328, 3,
	2, 2, 2, 327, 319, 3, 2, 2, 2, 328, 331, 3, 2, 2, 2, 329, 327, 3, 2, 2,
	2, 329, 330, 3, 2, 2, 2, 330, 333, 3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 332,
	312, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 335,
	7, 11, 2, 2, 335, 47, 3, 2, 2, 2, 336, 337, 7, 27, 2, 2, 337, 339, 5, 22,
	12, 2, 338, 340, 5, 50, 26, 2, 339, 338, 3, 2, 2, 2, 339, 340, 3, 2, 2,
	2, 340, 341, 3, 2, 2, 2, 341, 342, 7, 17, 2, 2, 342, 343, 5, 22, 12, 2,
	343, 344, 7, 3, 2, 2, 344, 49, 3, 2, 2, 2, 345, 346, 7, 10, 2, 2, 346,
	347, 7, 12, 2, 2, 347, 348, 7, 11, 2, 2, 348, 51, 3, 2, 2, 2, 349, 350,
	7, 28, 2, 2, 350, 351, 7, 39, 2, 2, 351, 352, 7, 17, 2, 2, 352, 353, 5,
	22, 12, 2, 353, 354, 7, 5, 2, 2, 354, 355, 5, 54, 28, 2, 355, 356, 7, 3,
	2, 2, 356, 53, 3, 2, 2, 2, 357, 359, 7, 29, 2, 2, 358, 357, 3, 2, 2, 2,
	358, 359, 3, 2, 2, 2, 359, 361, 3, 2, 2, 2, 360, 362, 7, 7, 2, 2, 361,
	360, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 366,
	7, 39, 2, 2, 364, 365, 7, 7, 2, 2, 365, 367, 7, 39, 2, 2, 366, 364, 3,
	2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 378, 3, 2, 2, 2, 368, 378, 7, 33, 2,
	2, 369, 378, 7, 34, 2, 2, 370, 378, 7, 32, 2, 2, 371, 378, 7, 37, 2, 2,
	372, 378, 7, 35, 2, 2, 373, 378, 7, 38, 2, 2, 374, 378, 5, 58, 30, 2, 375,
	378, 5, 56, 29, 2, 376, 378, 5, 60, 31, 2, 377, 358, 3, 2, 2, 2, 377, 368,
	3, 2, 2, 2, 377, 369, 3, 2, 2, 2, 377, 370, 3, 2, 2, 2, 377, 371, 3, 2,
	2, 2, 377, 372, 3, 2, 2, 2, 377, 373, 3, 2, 2, 2, 377, 374, 3, 2, 2, 2,
	377, 375, 3, 2, 2, 2, 377, 376, 3, 2, 2, 2, 378, 55, 3, 2, 2, 2, 379, 380,
	7, 10, 2, 2, 380, 381, 7, 39, 2, 2, 381, 382, 7, 5, 2, 2, 382, 389, 5,
	62, 32, 2, 383, 384, 7, 18, 2, 2, 384, 385, 7, 39, 2, 2, 385, 386, 7, 5,
	2, 2, 386, 388, 5, 62, 32, 2, 387, 383, 3, 2, 2, 2, 388, 391, 3, 2, 2,
	2, 389, 387, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 392, 3, 2, 2, 2, 391,
	389, 3, 2, 2, 2, 392, 393, 7, 11, 2, 2, 393, 57, 3, 2, 2, 2, 394, 395,
	7, 25, 2, 2, 395, 400, 5, 54, 28, 2, 396, 397, 7, 18, 2, 2, 397, 399, 5,
	54, 28, 2, 398, 396, 3, 2, 2, 2, 399, 402, 3, 2, 2, 2, 400, 398, 3, 2,
	2, 2, 400, 401, 3, 2, 2, 2, 401, 403, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2,
	403, 404, 7, 26, 2, 2, 404, 59, 3, 2, 2, 2, 405, 406, 7, 30, 2, 2, 406,
	407, 7, 32, 2, 2, 407, 61, 3, 2, 2, 2, 408, 409, 7, 10, 2, 2, 409, 410,
	7, 39, 2, 2, 410, 411, 7, 5, 2, 2, 411, 412, 5, 54, 28, 2, 412, 413, 7,
	11, 2, 2, 413, 416, 3, 2, 2, 2, 414, 416, 5, 54, 28, 2, 415, 408, 3, 2,
	2, 2, 415, 414, 3, 2, 2, 2, 416, 63, 3, 2, 2, 2, 417, 418, 7, 4, 2, 2,
	418, 423, 7, 39, 2, 2, 419, 420, 7, 7, 2, 2, 420, 422, 7, 39, 2, 2, 421,
	419, 3, 2, 2, 2, 422, 425, 3, 2, 2, 2, 423, 421, 3, 2, 2, 2, 423, 424,
	3, 2, 2, 2, 424, 428, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 426, 427, 7, 5,
	2, 2, 427, 429, 5, 22, 12, 2, 428, 426, 3, 2, 2, 2, 428, 429, 3, 2, 2,
	2, 429, 430, 3, 2, 2, 2, 430, 431, 7, 3, 2, 2, 431, 65, 3, 2, 2, 2, 50,
	70, 74, 79, 90, 96, 114, 119, 125, 140, 149, 155, 167, 175, 181, 185, 190,
	195, 199, 207, 213, 221, 230, 236, 243, 253, 262, 270, 278, 282, 285, 289,
	294, 296, 306, 317, 325, 329, 332, 339, 358, 361, 366, 377, 389, 400, 415,
	423, 428,
}
var literalNames = []string{
	"", "';'", "'using'", "'='", "'import'", "'.'", "'$'", "'.namespace'",
	"'('", "')'", "'struct'", "'{'", "'}'", "'interface'", "'extends'", "':'",
	"','", "'enum'", "'.ann'", "':union'", "'union'", "':group'", "'->'", "'['",
	"']'", "'annotation'", "'const'", "'-'", "'0x'", "", "", "", "", "", "",
	"", "'void'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "LOCATOR", "TEXT", "INTEGER",
	"FLOAT", "HEXADECIMAL", "FILE_ID", "BOOLEAN", "VOID", "NAME", "COMMENT",
	"WHITESPACE",
}

var ruleNames = []string{
	"document", "file_identifier", "using_import", "namespace_", "document_content",
	"struct_def", "struct_content", "interface_def", "interface_content", "field_def",
	"type_", "inner_type", "enum_def", "annotation_reference", "enum_content",
	"named_union_def", "unnamed_union_def", "union_content", "group_def", "group_content",
	"function_def", "generic_type_parameters", "function_parameters", "annotation_def",
	"annotation_parameters", "const_def", "const_value", "literal_union", "literal_list",
	"literal_bytes", "union_mapping", "inner_using",
}

type CapnProtoParser struct {
	*antlr.BaseParser
}

// NewCapnProtoParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *CapnProtoParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCapnProtoParser(input antlr.TokenStream) *CapnProtoParser {
	this := new(CapnProtoParser)
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
	this.GrammarFileName = "CapnProto.g4"

	return this
}

// CapnProtoParser tokens.
const (
	CapnProtoParserEOF         = antlr.TokenEOF
	CapnProtoParserT__0        = 1
	CapnProtoParserT__1        = 2
	CapnProtoParserT__2        = 3
	CapnProtoParserT__3        = 4
	CapnProtoParserT__4        = 5
	CapnProtoParserT__5        = 6
	CapnProtoParserT__6        = 7
	CapnProtoParserT__7        = 8
	CapnProtoParserT__8        = 9
	CapnProtoParserT__9        = 10
	CapnProtoParserT__10       = 11
	CapnProtoParserT__11       = 12
	CapnProtoParserT__12       = 13
	CapnProtoParserT__13       = 14
	CapnProtoParserT__14       = 15
	CapnProtoParserT__15       = 16
	CapnProtoParserT__16       = 17
	CapnProtoParserT__17       = 18
	CapnProtoParserT__18       = 19
	CapnProtoParserT__19       = 20
	CapnProtoParserT__20       = 21
	CapnProtoParserT__21       = 22
	CapnProtoParserT__22       = 23
	CapnProtoParserT__23       = 24
	CapnProtoParserT__24       = 25
	CapnProtoParserT__25       = 26
	CapnProtoParserT__26       = 27
	CapnProtoParserT__27       = 28
	CapnProtoParserLOCATOR     = 29
	CapnProtoParserTEXT        = 30
	CapnProtoParserINTEGER     = 31
	CapnProtoParserFLOAT       = 32
	CapnProtoParserHEXADECIMAL = 33
	CapnProtoParserFILE_ID     = 34
	CapnProtoParserBOOLEAN     = 35
	CapnProtoParserVOID        = 36
	CapnProtoParserNAME        = 37
	CapnProtoParserCOMMENT     = 38
	CapnProtoParserWHITESPACE  = 39
)

// CapnProtoParser rules.
const (
	CapnProtoParserRULE_document                = 0
	CapnProtoParserRULE_file_identifier         = 1
	CapnProtoParserRULE_using_import            = 2
	CapnProtoParserRULE_namespace_              = 3
	CapnProtoParserRULE_document_content        = 4
	CapnProtoParserRULE_struct_def              = 5
	CapnProtoParserRULE_struct_content          = 6
	CapnProtoParserRULE_interface_def           = 7
	CapnProtoParserRULE_interface_content       = 8
	CapnProtoParserRULE_field_def               = 9
	CapnProtoParserRULE_type_                   = 10
	CapnProtoParserRULE_inner_type              = 11
	CapnProtoParserRULE_enum_def                = 12
	CapnProtoParserRULE_annotation_reference    = 13
	CapnProtoParserRULE_enum_content            = 14
	CapnProtoParserRULE_named_union_def         = 15
	CapnProtoParserRULE_unnamed_union_def       = 16
	CapnProtoParserRULE_union_content           = 17
	CapnProtoParserRULE_group_def               = 18
	CapnProtoParserRULE_group_content           = 19
	CapnProtoParserRULE_function_def            = 20
	CapnProtoParserRULE_generic_type_parameters = 21
	CapnProtoParserRULE_function_parameters     = 22
	CapnProtoParserRULE_annotation_def          = 23
	CapnProtoParserRULE_annotation_parameters   = 24
	CapnProtoParserRULE_const_def               = 25
	CapnProtoParserRULE_const_value             = 26
	CapnProtoParserRULE_literal_union           = 27
	CapnProtoParserRULE_literal_list            = 28
	CapnProtoParserRULE_literal_bytes           = 29
	CapnProtoParserRULE_union_mapping           = 30
	CapnProtoParserRULE_inner_using             = 31
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
	p.RuleIndex = CapnProtoParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) File_identifier() IFile_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFile_identifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFile_identifierContext)
}

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserEOF, 0)
}

func (s *DocumentContext) AllUsing_import() []IUsing_importContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUsing_importContext)(nil)).Elem())
	var tst = make([]IUsing_importContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUsing_importContext)
		}
	}

	return tst
}

func (s *DocumentContext) Using_import(i int) IUsing_importContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUsing_importContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUsing_importContext)
}

func (s *DocumentContext) Namespace_() INamespace_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespace_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespace_Context)
}

func (s *DocumentContext) AllDocument_content() []IDocument_contentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDocument_contentContext)(nil)).Elem())
	var tst = make([]IDocument_contentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDocument_contentContext)
		}
	}

	return tst
}

func (s *DocumentContext) Document_content(i int) IDocument_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocument_contentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDocument_contentContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *CapnProtoParser) Document() (localctx IDocumentContext) {
	this := p
	_ = this

	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CapnProtoParserRULE_document)
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
		p.SetState(64)
		p.File_identifier()
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserT__1 {
		{
			p.SetState(65)
			p.Using_import()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__5 {
		{
			p.SetState(71)
			p.Namespace_()
		}

	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(CapnProtoParserT__9-10))|(1<<(CapnProtoParserT__12-10))|(1<<(CapnProtoParserT__16-10))|(1<<(CapnProtoParserT__24-10))|(1<<(CapnProtoParserT__25-10))|(1<<(CapnProtoParserNAME-10)))) != 0 {
		{
			p.SetState(74)
			p.Document_content()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(CapnProtoParserEOF)
	}

	return localctx
}

// IFile_identifierContext is an interface to support dynamic dispatch.
type IFile_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_identifierContext differentiates from other interfaces.
	IsFile_identifierContext()
}

type File_identifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_identifierContext() *File_identifierContext {
	var p = new(File_identifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_file_identifier
	return p
}

func (*File_identifierContext) IsFile_identifierContext() {}

func NewFile_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_identifierContext {
	var p = new(File_identifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_file_identifier

	return p
}

func (s *File_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *File_identifierContext) FILE_ID() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserFILE_ID, 0)
}

func (s *File_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterFile_identifier(s)
	}
}

func (s *File_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitFile_identifier(s)
	}
}

func (p *CapnProtoParser) File_identifier() (localctx IFile_identifierContext) {
	this := p
	_ = this

	localctx = NewFile_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CapnProtoParserRULE_file_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(CapnProtoParserFILE_ID)
	}
	{
		p.SetState(83)
		p.Match(CapnProtoParserT__0)
	}

	return localctx
}

// IUsing_importContext is an interface to support dynamic dispatch.
type IUsing_importContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUsing_importContext differentiates from other interfaces.
	IsUsing_importContext()
}

type Using_importContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUsing_importContext() *Using_importContext {
	var p = new(Using_importContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_using_import
	return p
}

func (*Using_importContext) IsUsing_importContext() {}

func NewUsing_importContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Using_importContext {
	var p = new(Using_importContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_using_import

	return p
}

func (s *Using_importContext) GetParser() antlr.Parser { return s.parser }

func (s *Using_importContext) TEXT() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserTEXT, 0)
}

func (s *Using_importContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(CapnProtoParserNAME)
}

func (s *Using_importContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, i)
}

func (s *Using_importContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Using_importContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Using_importContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterUsing_import(s)
	}
}

func (s *Using_importContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitUsing_import(s)
	}
}

func (p *CapnProtoParser) Using_import() (localctx IUsing_importContext) {
	this := p
	_ = this

	localctx = NewUsing_importContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CapnProtoParserRULE_using_import)
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
		p.SetState(85)
		p.Match(CapnProtoParserT__1)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserNAME {
		{
			p.SetState(86)
			p.Match(CapnProtoParserNAME)
		}
		{
			p.SetState(87)
			p.Match(CapnProtoParserT__2)
		}

	}
	{
		p.SetState(90)
		p.Match(CapnProtoParserT__3)
	}
	{
		p.SetState(91)
		p.Match(CapnProtoParserTEXT)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__4 {
		{
			p.SetState(92)
			p.Match(CapnProtoParserT__4)
		}
		{
			p.SetState(93)
			p.Match(CapnProtoParserNAME)
		}

	}
	{
		p.SetState(96)
		p.Match(CapnProtoParserT__0)
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
	p.RuleIndex = CapnProtoParserRULE_namespace_
	return p
}

func (*Namespace_Context) IsNamespace_Context() {}

func NewNamespace_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Namespace_Context {
	var p = new(Namespace_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_namespace_

	return p
}

func (s *Namespace_Context) GetParser() antlr.Parser { return s.parser }

func (s *Namespace_Context) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Namespace_Context) TEXT() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserTEXT, 0)
}

func (s *Namespace_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Namespace_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Namespace_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterNamespace_(s)
	}
}

func (s *Namespace_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitNamespace_(s)
	}
}

func (p *CapnProtoParser) Namespace_() (localctx INamespace_Context) {
	this := p
	_ = this

	localctx = NewNamespace_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CapnProtoParserRULE_namespace_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(CapnProtoParserT__5)
	}
	{
		p.SetState(99)
		p.Match(CapnProtoParserNAME)
	}
	{
		p.SetState(100)
		p.Match(CapnProtoParserT__6)
	}
	{
		p.SetState(101)
		p.Match(CapnProtoParserT__7)
	}
	{
		p.SetState(102)
		p.Match(CapnProtoParserTEXT)
	}
	{
		p.SetState(103)
		p.Match(CapnProtoParserT__8)
	}
	{
		p.SetState(104)
		p.Match(CapnProtoParserT__0)
	}

	return localctx
}

// IDocument_contentContext is an interface to support dynamic dispatch.
type IDocument_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocument_contentContext differentiates from other interfaces.
	IsDocument_contentContext()
}

type Document_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocument_contentContext() *Document_contentContext {
	var p = new(Document_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_document_content
	return p
}

func (*Document_contentContext) IsDocument_contentContext() {}

func NewDocument_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Document_contentContext {
	var p = new(Document_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_document_content

	return p
}

func (s *Document_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Document_contentContext) Struct_def() IStruct_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStruct_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStruct_defContext)
}

func (s *Document_contentContext) Interface_def() IInterface_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterface_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterface_defContext)
}

func (s *Document_contentContext) Function_def() IFunction_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_defContext)
}

func (s *Document_contentContext) Annotation_def() IAnnotation_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotation_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotation_defContext)
}

func (s *Document_contentContext) Const_def() IConst_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_defContext)
}

func (s *Document_contentContext) Enum_def() IEnum_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnum_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnum_defContext)
}

func (s *Document_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Document_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Document_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterDocument_content(s)
	}
}

func (s *Document_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitDocument_content(s)
	}
}

func (p *CapnProtoParser) Document_content() (localctx IDocument_contentContext) {
	this := p
	_ = this

	localctx = NewDocument_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CapnProtoParserRULE_document_content)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CapnProtoParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Struct_def()
		}

	case CapnProtoParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Interface_def()
		}

	case CapnProtoParserNAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Function_def()
		}

	case CapnProtoParserT__24:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(109)
			p.Annotation_def()
		}

	case CapnProtoParserT__25:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(110)
			p.Const_def()
		}

	case CapnProtoParserT__16:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(111)
			p.Enum_def()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStruct_defContext is an interface to support dynamic dispatch.
type IStruct_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStruct_defContext differentiates from other interfaces.
	IsStruct_defContext()
}

type Struct_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_defContext() *Struct_defContext {
	var p = new(Struct_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_struct_def
	return p
}

func (*Struct_defContext) IsStruct_defContext() {}

func NewStruct_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_defContext {
	var p = new(Struct_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_struct_def

	return p
}

func (s *Struct_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_defContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Struct_defContext) Annotation_reference() IAnnotation_referenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotation_referenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotation_referenceContext)
}

func (s *Struct_defContext) AllStruct_content() []IStruct_contentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStruct_contentContext)(nil)).Elem())
	var tst = make([]IStruct_contentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStruct_contentContext)
		}
	}

	return tst
}

func (s *Struct_defContext) Struct_content(i int) IStruct_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStruct_contentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStruct_contentContext)
}

func (s *Struct_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterStruct_def(s)
	}
}

func (s *Struct_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitStruct_def(s)
	}
}

func (p *CapnProtoParser) Struct_def() (localctx IStruct_defContext) {
	this := p
	_ = this

	localctx = NewStruct_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CapnProtoParserRULE_struct_def)
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
		p.SetState(114)
		p.Match(CapnProtoParserT__9)
	}
	{
		p.SetState(115)
		p.Type_()
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__5 {
		{
			p.SetState(116)
			p.Annotation_reference()
		}

	}
	{
		p.SetState(119)
		p.Match(CapnProtoParserT__10)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CapnProtoParserT__1)|(1<<CapnProtoParserT__9)|(1<<CapnProtoParserT__12)|(1<<CapnProtoParserT__16)|(1<<CapnProtoParserT__19)|(1<<CapnProtoParserT__24)|(1<<CapnProtoParserT__25))) != 0) || _la == CapnProtoParserNAME {
		{
			p.SetState(120)
			p.Struct_content()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(CapnProtoParserT__11)
	}

	return localctx
}

// IStruct_contentContext is an interface to support dynamic dispatch.
type IStruct_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStruct_contentContext differentiates from other interfaces.
	IsStruct_contentContext()
}

type Struct_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_contentContext() *Struct_contentContext {
	var p = new(Struct_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_struct_content
	return p
}

func (*Struct_contentContext) IsStruct_contentContext() {}

func NewStruct_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_contentContext {
	var p = new(Struct_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_struct_content

	return p
}

func (s *Struct_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_contentContext) Field_def() IField_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_defContext)
}

func (s *Struct_contentContext) Enum_def() IEnum_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnum_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnum_defContext)
}

func (s *Struct_contentContext) Named_union_def() INamed_union_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamed_union_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamed_union_defContext)
}

func (s *Struct_contentContext) Unnamed_union_def() IUnnamed_union_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnnamed_union_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnnamed_union_defContext)
}

func (s *Struct_contentContext) Interface_def() IInterface_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterface_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterface_defContext)
}

func (s *Struct_contentContext) Annotation_def() IAnnotation_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotation_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotation_defContext)
}

func (s *Struct_contentContext) Struct_def() IStruct_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStruct_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStruct_defContext)
}

func (s *Struct_contentContext) Group_def() IGroup_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_defContext)
}

func (s *Struct_contentContext) Const_def() IConst_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_defContext)
}

func (s *Struct_contentContext) Inner_using() IInner_usingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInner_usingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInner_usingContext)
}

func (s *Struct_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterStruct_content(s)
	}
}

func (s *Struct_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitStruct_content(s)
	}
}

func (p *CapnProtoParser) Struct_content() (localctx IStruct_contentContext) {
	this := p
	_ = this

	localctx = NewStruct_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CapnProtoParserRULE_struct_content)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Field_def()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Enum_def()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Named_union_def()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.Unnamed_union_def()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.Interface_def()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(133)
			p.Annotation_def()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(134)
			p.Struct_def()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(135)
			p.Group_def()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(136)
			p.Const_def()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(137)
			p.Inner_using()
		}

	}

	return localctx
}

// IInterface_defContext is an interface to support dynamic dispatch.
type IInterface_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterface_defContext differentiates from other interfaces.
	IsInterface_defContext()
}

type Interface_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterface_defContext() *Interface_defContext {
	var p = new(Interface_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_interface_def
	return p
}

func (*Interface_defContext) IsInterface_defContext() {}

func NewInterface_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Interface_defContext {
	var p = new(Interface_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_interface_def

	return p
}

func (s *Interface_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Interface_defContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *Interface_defContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Interface_defContext) AllInterface_content() []IInterface_contentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInterface_contentContext)(nil)).Elem())
	var tst = make([]IInterface_contentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInterface_contentContext)
		}
	}

	return tst
}

func (s *Interface_defContext) Interface_content(i int) IInterface_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterface_contentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInterface_contentContext)
}

func (s *Interface_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Interface_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Interface_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterInterface_def(s)
	}
}

func (s *Interface_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitInterface_def(s)
	}
}

func (p *CapnProtoParser) Interface_def() (localctx IInterface_defContext) {
	this := p
	_ = this

	localctx = NewInterface_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CapnProtoParserRULE_interface_def)
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
		p.SetState(140)
		p.Match(CapnProtoParserT__12)
	}
	{
		p.SetState(141)
		p.Type_()
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__13 {
		{
			p.SetState(142)
			p.Match(CapnProtoParserT__13)
		}
		{
			p.SetState(143)
			p.Match(CapnProtoParserT__7)
		}
		{
			p.SetState(144)
			p.Type_()
		}
		{
			p.SetState(145)
			p.Match(CapnProtoParserT__8)
		}

	}
	{
		p.SetState(149)
		p.Match(CapnProtoParserT__10)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(CapnProtoParserT__9-10))|(1<<(CapnProtoParserT__12-10))|(1<<(CapnProtoParserT__16-10))|(1<<(CapnProtoParserT__19-10))|(1<<(CapnProtoParserNAME-10)))) != 0 {
		{
			p.SetState(150)
			p.Interface_content()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(CapnProtoParserT__11)
	}

	return localctx
}

// IInterface_contentContext is an interface to support dynamic dispatch.
type IInterface_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterface_contentContext differentiates from other interfaces.
	IsInterface_contentContext()
}

type Interface_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterface_contentContext() *Interface_contentContext {
	var p = new(Interface_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_interface_content
	return p
}

func (*Interface_contentContext) IsInterface_contentContext() {}

func NewInterface_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Interface_contentContext {
	var p = new(Interface_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_interface_content

	return p
}

func (s *Interface_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Interface_contentContext) Field_def() IField_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_defContext)
}

func (s *Interface_contentContext) Enum_def() IEnum_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnum_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnum_defContext)
}

func (s *Interface_contentContext) Named_union_def() INamed_union_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamed_union_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamed_union_defContext)
}

func (s *Interface_contentContext) Unnamed_union_def() IUnnamed_union_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnnamed_union_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnnamed_union_defContext)
}

func (s *Interface_contentContext) Interface_def() IInterface_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterface_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterface_defContext)
}

func (s *Interface_contentContext) Struct_def() IStruct_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStruct_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStruct_defContext)
}

func (s *Interface_contentContext) Function_def() IFunction_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_defContext)
}

func (s *Interface_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Interface_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Interface_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterInterface_content(s)
	}
}

func (s *Interface_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitInterface_content(s)
	}
}

func (p *CapnProtoParser) Interface_content() (localctx IInterface_contentContext) {
	this := p
	_ = this

	localctx = NewInterface_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CapnProtoParserRULE_interface_content)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Field_def()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Enum_def()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Named_union_def()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(161)
			p.Unnamed_union_def()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(162)
			p.Interface_def()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(163)
			p.Struct_def()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(164)
			p.Function_def()
		}

	}

	return localctx
}

// IField_defContext is an interface to support dynamic dispatch.
type IField_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_defContext differentiates from other interfaces.
	IsField_defContext()
}

type Field_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_defContext() *Field_defContext {
	var p = new(Field_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_field_def
	return p
}

func (*Field_defContext) IsField_defContext() {}

func NewField_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_defContext {
	var p = new(Field_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_field_def

	return p
}

func (s *Field_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_defContext) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Field_defContext) LOCATOR() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserLOCATOR, 0)
}

func (s *Field_defContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Field_defContext) Const_value() IConst_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Field_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterField_def(s)
	}
}

func (s *Field_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitField_def(s)
	}
}

func (p *CapnProtoParser) Field_def() (localctx IField_defContext) {
	this := p
	_ = this

	localctx = NewField_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CapnProtoParserRULE_field_def)
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
		p.SetState(167)
		p.Match(CapnProtoParserNAME)
	}
	{
		p.SetState(168)
		p.Match(CapnProtoParserLOCATOR)
	}
	{
		p.SetState(169)
		p.Match(CapnProtoParserT__14)
	}
	{
		p.SetState(170)
		p.Type_()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__2 {
		{
			p.SetState(171)
			p.Match(CapnProtoParserT__2)
		}
		{
			p.SetState(172)
			p.Const_value()
		}

	}
	{
		p.SetState(175)
		p.Match(CapnProtoParserT__0)
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
	p.RuleIndex = CapnProtoParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Type_Context) Inner_type() IInner_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInner_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInner_typeContext)
}

func (s *Type_Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

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
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *CapnProtoParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CapnProtoParserRULE_type_)
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
		p.SetState(177)
		p.Match(CapnProtoParserNAME)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(178)
			p.Inner_type()
		}

	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__4 {
		{
			p.SetState(181)
			p.Match(CapnProtoParserT__4)
		}
		{
			p.SetState(182)
			p.Type_()
		}

	}

	return localctx
}

// IInner_typeContext is an interface to support dynamic dispatch.
type IInner_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInner_typeContext differentiates from other interfaces.
	IsInner_typeContext()
}

type Inner_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInner_typeContext() *Inner_typeContext {
	var p = new(Inner_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_inner_type
	return p
}

func (*Inner_typeContext) IsInner_typeContext() {}

func NewInner_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inner_typeContext {
	var p = new(Inner_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_inner_type

	return p
}

func (s *Inner_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Inner_typeContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *Inner_typeContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Inner_typeContext) AllInner_type() []IInner_typeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInner_typeContext)(nil)).Elem())
	var tst = make([]IInner_typeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInner_typeContext)
		}
	}

	return tst
}

func (s *Inner_typeContext) Inner_type(i int) IInner_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInner_typeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInner_typeContext)
}

func (s *Inner_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inner_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inner_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterInner_type(s)
	}
}

func (s *Inner_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitInner_type(s)
	}
}

func (p *CapnProtoParser) Inner_type() (localctx IInner_typeContext) {
	this := p
	_ = this

	localctx = NewInner_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CapnProtoParserRULE_inner_type)
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
		p.SetState(185)
		p.Match(CapnProtoParserT__7)
	}
	{
		p.SetState(186)
		p.Type_()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__7 {
		{
			p.SetState(187)
			p.Inner_type()
		}

	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserT__15 {
		{
			p.SetState(190)
			p.Match(CapnProtoParserT__15)
		}
		{
			p.SetState(191)
			p.Type_()
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CapnProtoParserT__7 {
			{
				p.SetState(192)
				p.Inner_type()
			}

		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(200)
		p.Match(CapnProtoParserT__8)
	}

	return localctx
}

// IEnum_defContext is an interface to support dynamic dispatch.
type IEnum_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnum_defContext differentiates from other interfaces.
	IsEnum_defContext()
}

type Enum_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnum_defContext() *Enum_defContext {
	var p = new(Enum_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_enum_def
	return p
}

func (*Enum_defContext) IsEnum_defContext() {}

func NewEnum_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_defContext {
	var p = new(Enum_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_enum_def

	return p
}

func (s *Enum_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_defContext) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Enum_defContext) Annotation_reference() IAnnotation_referenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotation_referenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotation_referenceContext)
}

func (s *Enum_defContext) AllEnum_content() []IEnum_contentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnum_contentContext)(nil)).Elem())
	var tst = make([]IEnum_contentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnum_contentContext)
		}
	}

	return tst
}

func (s *Enum_defContext) Enum_content(i int) IEnum_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnum_contentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnum_contentContext)
}

func (s *Enum_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterEnum_def(s)
	}
}

func (s *Enum_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitEnum_def(s)
	}
}

func (p *CapnProtoParser) Enum_def() (localctx IEnum_defContext) {
	this := p
	_ = this

	localctx = NewEnum_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CapnProtoParserRULE_enum_def)
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
		p.SetState(202)
		p.Match(CapnProtoParserT__16)
	}
	{
		p.SetState(203)
		p.Match(CapnProtoParserNAME)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__5 {
		{
			p.SetState(204)
			p.Annotation_reference()
		}

	}
	{
		p.SetState(207)
		p.Match(CapnProtoParserT__10)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserNAME {
		{
			p.SetState(208)
			p.Enum_content()
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(214)
		p.Match(CapnProtoParserT__11)
	}

	return localctx
}

// IAnnotation_referenceContext is an interface to support dynamic dispatch.
type IAnnotation_referenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotation_referenceContext differentiates from other interfaces.
	IsAnnotation_referenceContext()
}

type Annotation_referenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_referenceContext() *Annotation_referenceContext {
	var p = new(Annotation_referenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_annotation_reference
	return p
}

func (*Annotation_referenceContext) IsAnnotation_referenceContext() {}

func NewAnnotation_referenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_referenceContext {
	var p = new(Annotation_referenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_annotation_reference

	return p
}

func (s *Annotation_referenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_referenceContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Annotation_referenceContext) TEXT() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserTEXT, 0)
}

func (s *Annotation_referenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_referenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_referenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterAnnotation_reference(s)
	}
}

func (s *Annotation_referenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitAnnotation_reference(s)
	}
}

func (p *CapnProtoParser) Annotation_reference() (localctx IAnnotation_referenceContext) {
	this := p
	_ = this

	localctx = NewAnnotation_referenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CapnProtoParserRULE_annotation_reference)
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
		p.Match(CapnProtoParserT__5)
	}
	{
		p.SetState(217)
		p.Type_()
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__17 {
		{
			p.SetState(218)
			p.Match(CapnProtoParserT__17)
		}

	}
	{
		p.SetState(221)
		p.Match(CapnProtoParserT__7)
	}
	{
		p.SetState(222)
		p.Match(CapnProtoParserTEXT)
	}
	{
		p.SetState(223)
		p.Match(CapnProtoParserT__8)
	}

	return localctx
}

// IEnum_contentContext is an interface to support dynamic dispatch.
type IEnum_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnum_contentContext differentiates from other interfaces.
	IsEnum_contentContext()
}

type Enum_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnum_contentContext() *Enum_contentContext {
	var p = new(Enum_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_enum_content
	return p
}

func (*Enum_contentContext) IsEnum_contentContext() {}

func NewEnum_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_contentContext {
	var p = new(Enum_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_enum_content

	return p
}

func (s *Enum_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_contentContext) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Enum_contentContext) LOCATOR() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserLOCATOR, 0)
}

func (s *Enum_contentContext) Annotation_reference() IAnnotation_referenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotation_referenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotation_referenceContext)
}

func (s *Enum_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterEnum_content(s)
	}
}

func (s *Enum_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitEnum_content(s)
	}
}

func (p *CapnProtoParser) Enum_content() (localctx IEnum_contentContext) {
	this := p
	_ = this

	localctx = NewEnum_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CapnProtoParserRULE_enum_content)
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
		p.SetState(225)
		p.Match(CapnProtoParserNAME)
	}
	{
		p.SetState(226)
		p.Match(CapnProtoParserLOCATOR)
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__5 {
		{
			p.SetState(227)
			p.Annotation_reference()
		}

	}
	{
		p.SetState(230)
		p.Match(CapnProtoParserT__0)
	}

	return localctx
}

// INamed_union_defContext is an interface to support dynamic dispatch.
type INamed_union_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamed_union_defContext differentiates from other interfaces.
	IsNamed_union_defContext()
}

type Named_union_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamed_union_defContext() *Named_union_defContext {
	var p = new(Named_union_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_named_union_def
	return p
}

func (*Named_union_defContext) IsNamed_union_defContext() {}

func NewNamed_union_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Named_union_defContext {
	var p = new(Named_union_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_named_union_def

	return p
}

func (s *Named_union_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Named_union_defContext) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Named_union_defContext) LOCATOR() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserLOCATOR, 0)
}

func (s *Named_union_defContext) AllUnion_content() []IUnion_contentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnion_contentContext)(nil)).Elem())
	var tst = make([]IUnion_contentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnion_contentContext)
		}
	}

	return tst
}

func (s *Named_union_defContext) Union_content(i int) IUnion_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnion_contentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnion_contentContext)
}

func (s *Named_union_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Named_union_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Named_union_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterNamed_union_def(s)
	}
}

func (s *Named_union_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitNamed_union_def(s)
	}
}

func (p *CapnProtoParser) Named_union_def() (localctx INamed_union_defContext) {
	this := p
	_ = this

	localctx = NewNamed_union_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CapnProtoParserRULE_named_union_def)
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
		p.Match(CapnProtoParserNAME)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserLOCATOR {
		{
			p.SetState(233)
			p.Match(CapnProtoParserLOCATOR)
		}

	}
	{
		p.SetState(236)
		p.Match(CapnProtoParserT__18)
	}
	{
		p.SetState(237)
		p.Match(CapnProtoParserT__10)
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserT__19 || _la == CapnProtoParserNAME {
		{
			p.SetState(238)
			p.Union_content()
		}

		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(244)
		p.Match(CapnProtoParserT__11)
	}

	return localctx
}

// IUnnamed_union_defContext is an interface to support dynamic dispatch.
type IUnnamed_union_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnnamed_union_defContext differentiates from other interfaces.
	IsUnnamed_union_defContext()
}

type Unnamed_union_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnnamed_union_defContext() *Unnamed_union_defContext {
	var p = new(Unnamed_union_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_unnamed_union_def
	return p
}

func (*Unnamed_union_defContext) IsUnnamed_union_defContext() {}

func NewUnnamed_union_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unnamed_union_defContext {
	var p = new(Unnamed_union_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_unnamed_union_def

	return p
}

func (s *Unnamed_union_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Unnamed_union_defContext) AllUnion_content() []IUnion_contentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnion_contentContext)(nil)).Elem())
	var tst = make([]IUnion_contentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnion_contentContext)
		}
	}

	return tst
}

func (s *Unnamed_union_defContext) Union_content(i int) IUnion_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnion_contentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnion_contentContext)
}

func (s *Unnamed_union_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unnamed_union_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unnamed_union_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterUnnamed_union_def(s)
	}
}

func (s *Unnamed_union_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitUnnamed_union_def(s)
	}
}

func (p *CapnProtoParser) Unnamed_union_def() (localctx IUnnamed_union_defContext) {
	this := p
	_ = this

	localctx = NewUnnamed_union_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CapnProtoParserRULE_unnamed_union_def)
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
		p.SetState(246)
		p.Match(CapnProtoParserT__19)
	}
	{
		p.SetState(247)
		p.Match(CapnProtoParserT__10)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserT__19 || _la == CapnProtoParserNAME {
		{
			p.SetState(248)
			p.Union_content()
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(254)
		p.Match(CapnProtoParserT__11)
	}

	return localctx
}

// IUnion_contentContext is an interface to support dynamic dispatch.
type IUnion_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnion_contentContext differentiates from other interfaces.
	IsUnion_contentContext()
}

type Union_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnion_contentContext() *Union_contentContext {
	var p = new(Union_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_union_content
	return p
}

func (*Union_contentContext) IsUnion_contentContext() {}

func NewUnion_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Union_contentContext {
	var p = new(Union_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_union_content

	return p
}

func (s *Union_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Union_contentContext) Field_def() IField_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_defContext)
}

func (s *Union_contentContext) Group_def() IGroup_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_defContext)
}

func (s *Union_contentContext) Unnamed_union_def() IUnnamed_union_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnnamed_union_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnnamed_union_defContext)
}

func (s *Union_contentContext) Named_union_def() INamed_union_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamed_union_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamed_union_defContext)
}

func (s *Union_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Union_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Union_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterUnion_content(s)
	}
}

func (s *Union_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitUnion_content(s)
	}
}

func (p *CapnProtoParser) Union_content() (localctx IUnion_contentContext) {
	this := p
	_ = this

	localctx = NewUnion_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CapnProtoParserRULE_union_content)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Field_def()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.Group_def()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(258)
			p.Unnamed_union_def()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(259)
			p.Named_union_def()
		}

	}

	return localctx
}

// IGroup_defContext is an interface to support dynamic dispatch.
type IGroup_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_defContext differentiates from other interfaces.
	IsGroup_defContext()
}

type Group_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_defContext() *Group_defContext {
	var p = new(Group_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_group_def
	return p
}

func (*Group_defContext) IsGroup_defContext() {}

func NewGroup_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_defContext {
	var p = new(Group_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_group_def

	return p
}

func (s *Group_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_defContext) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Group_defContext) AllGroup_content() []IGroup_contentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroup_contentContext)(nil)).Elem())
	var tst = make([]IGroup_contentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroup_contentContext)
		}
	}

	return tst
}

func (s *Group_defContext) Group_content(i int) IGroup_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_contentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroup_contentContext)
}

func (s *Group_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterGroup_def(s)
	}
}

func (s *Group_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitGroup_def(s)
	}
}

func (p *CapnProtoParser) Group_def() (localctx IGroup_defContext) {
	this := p
	_ = this

	localctx = NewGroup_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CapnProtoParserRULE_group_def)
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
		p.SetState(262)
		p.Match(CapnProtoParserNAME)
	}
	{
		p.SetState(263)
		p.Match(CapnProtoParserT__20)
	}
	{
		p.SetState(264)
		p.Match(CapnProtoParserT__10)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserT__19 || _la == CapnProtoParserNAME {
		{
			p.SetState(265)
			p.Group_content()
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(271)
		p.Match(CapnProtoParserT__11)
	}

	return localctx
}

// IGroup_contentContext is an interface to support dynamic dispatch.
type IGroup_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_contentContext differentiates from other interfaces.
	IsGroup_contentContext()
}

type Group_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_contentContext() *Group_contentContext {
	var p = new(Group_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_group_content
	return p
}

func (*Group_contentContext) IsGroup_contentContext() {}

func NewGroup_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_contentContext {
	var p = new(Group_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_group_content

	return p
}

func (s *Group_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_contentContext) Field_def() IField_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_defContext)
}

func (s *Group_contentContext) Unnamed_union_def() IUnnamed_union_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnnamed_union_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnnamed_union_defContext)
}

func (s *Group_contentContext) Named_union_def() INamed_union_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamed_union_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamed_union_defContext)
}

func (s *Group_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterGroup_content(s)
	}
}

func (s *Group_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitGroup_content(s)
	}
}

func (p *CapnProtoParser) Group_content() (localctx IGroup_contentContext) {
	this := p
	_ = this

	localctx = NewGroup_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CapnProtoParserRULE_group_content)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.Field_def()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Unnamed_union_def()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(275)
			p.Named_union_def()
		}

	}

	return localctx
}

// IFunction_defContext is an interface to support dynamic dispatch.
type IFunction_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_defContext differentiates from other interfaces.
	IsFunction_defContext()
}

type Function_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_defContext() *Function_defContext {
	var p = new(Function_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_function_def
	return p
}

func (*Function_defContext) IsFunction_defContext() {}

func NewFunction_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_defContext {
	var p = new(Function_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_function_def

	return p
}

func (s *Function_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_defContext) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Function_defContext) AllFunction_parameters() []IFunction_parametersContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunction_parametersContext)(nil)).Elem())
	var tst = make([]IFunction_parametersContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunction_parametersContext)
		}
	}

	return tst
}

func (s *Function_defContext) Function_parameters(i int) IFunction_parametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_parametersContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunction_parametersContext)
}

func (s *Function_defContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *Function_defContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Function_defContext) LOCATOR() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserLOCATOR, 0)
}

func (s *Function_defContext) Generic_type_parameters() IGeneric_type_parametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeneric_type_parametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeneric_type_parametersContext)
}

func (s *Function_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterFunction_def(s)
	}
}

func (s *Function_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitFunction_def(s)
	}
}

func (p *CapnProtoParser) Function_def() (localctx IFunction_defContext) {
	this := p
	_ = this

	localctx = NewFunction_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CapnProtoParserRULE_function_def)
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
		p.SetState(278)
		p.Match(CapnProtoParserNAME)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserLOCATOR {
		{
			p.SetState(279)
			p.Match(CapnProtoParserLOCATOR)
		}

	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__22 {
		{
			p.SetState(282)
			p.Generic_type_parameters()
		}

	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CapnProtoParserT__7:
		{
			p.SetState(285)
			p.Function_parameters()
		}

	case CapnProtoParserNAME:
		{
			p.SetState(286)
			p.Type_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__21 {
		{
			p.SetState(289)
			p.Match(CapnProtoParserT__21)
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CapnProtoParserT__7:
			{
				p.SetState(290)
				p.Function_parameters()
			}

		case CapnProtoParserNAME:
			{
				p.SetState(291)
				p.Type_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}
	{
		p.SetState(296)
		p.Match(CapnProtoParserT__0)
	}

	return localctx
}

// IGeneric_type_parametersContext is an interface to support dynamic dispatch.
type IGeneric_type_parametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeneric_type_parametersContext differentiates from other interfaces.
	IsGeneric_type_parametersContext()
}

type Generic_type_parametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeneric_type_parametersContext() *Generic_type_parametersContext {
	var p = new(Generic_type_parametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_generic_type_parameters
	return p
}

func (*Generic_type_parametersContext) IsGeneric_type_parametersContext() {}

func NewGeneric_type_parametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Generic_type_parametersContext {
	var p = new(Generic_type_parametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_generic_type_parameters

	return p
}

func (s *Generic_type_parametersContext) GetParser() antlr.Parser { return s.parser }

func (s *Generic_type_parametersContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(CapnProtoParserNAME)
}

func (s *Generic_type_parametersContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, i)
}

func (s *Generic_type_parametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Generic_type_parametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Generic_type_parametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterGeneric_type_parameters(s)
	}
}

func (s *Generic_type_parametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitGeneric_type_parameters(s)
	}
}

func (p *CapnProtoParser) Generic_type_parameters() (localctx IGeneric_type_parametersContext) {
	this := p
	_ = this

	localctx = NewGeneric_type_parametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CapnProtoParserRULE_generic_type_parameters)
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
		p.SetState(298)
		p.Match(CapnProtoParserT__22)
	}
	{
		p.SetState(299)
		p.Match(CapnProtoParserNAME)
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserT__15 {
		{
			p.SetState(300)
			p.Match(CapnProtoParserT__15)
		}
		{
			p.SetState(301)
			p.Match(CapnProtoParserNAME)
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(307)
		p.Match(CapnProtoParserT__23)
	}

	return localctx
}

// IFunction_parametersContext is an interface to support dynamic dispatch.
type IFunction_parametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_parametersContext differentiates from other interfaces.
	IsFunction_parametersContext()
}

type Function_parametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_parametersContext() *Function_parametersContext {
	var p = new(Function_parametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_function_parameters
	return p
}

func (*Function_parametersContext) IsFunction_parametersContext() {}

func NewFunction_parametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_parametersContext {
	var p = new(Function_parametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_function_parameters

	return p
}

func (s *Function_parametersContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_parametersContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(CapnProtoParserNAME)
}

func (s *Function_parametersContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, i)
}

func (s *Function_parametersContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *Function_parametersContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Function_parametersContext) AllConst_value() []IConst_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConst_valueContext)(nil)).Elem())
	var tst = make([]IConst_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConst_valueContext)
		}
	}

	return tst
}

func (s *Function_parametersContext) Const_value(i int) IConst_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Function_parametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_parametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_parametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterFunction_parameters(s)
	}
}

func (s *Function_parametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitFunction_parameters(s)
	}
}

func (p *CapnProtoParser) Function_parameters() (localctx IFunction_parametersContext) {
	this := p
	_ = this

	localctx = NewFunction_parametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CapnProtoParserRULE_function_parameters)
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
		p.SetState(309)
		p.Match(CapnProtoParserT__7)
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserNAME {
		{
			p.SetState(310)
			p.Match(CapnProtoParserNAME)
		}
		{
			p.SetState(311)
			p.Match(CapnProtoParserT__14)
		}
		{
			p.SetState(312)
			p.Type_()
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CapnProtoParserT__2 {
			{
				p.SetState(313)
				p.Match(CapnProtoParserT__2)
			}
			{
				p.SetState(314)
				p.Const_value()
			}

		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CapnProtoParserT__15 {
			{
				p.SetState(317)
				p.Match(CapnProtoParserT__15)
			}
			{
				p.SetState(318)
				p.Match(CapnProtoParserNAME)
			}
			{
				p.SetState(319)
				p.Match(CapnProtoParserT__14)
			}
			{
				p.SetState(320)
				p.Type_()
			}
			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CapnProtoParserT__2 {
				{
					p.SetState(321)
					p.Match(CapnProtoParserT__2)
				}
				{
					p.SetState(322)
					p.Const_value()
				}

			}

			p.SetState(329)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(332)
		p.Match(CapnProtoParserT__8)
	}

	return localctx
}

// IAnnotation_defContext is an interface to support dynamic dispatch.
type IAnnotation_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotation_defContext differentiates from other interfaces.
	IsAnnotation_defContext()
}

type Annotation_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_defContext() *Annotation_defContext {
	var p = new(Annotation_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_annotation_def
	return p
}

func (*Annotation_defContext) IsAnnotation_defContext() {}

func NewAnnotation_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_defContext {
	var p = new(Annotation_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_annotation_def

	return p
}

func (s *Annotation_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_defContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *Annotation_defContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Annotation_defContext) Annotation_parameters() IAnnotation_parametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotation_parametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotation_parametersContext)
}

func (s *Annotation_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterAnnotation_def(s)
	}
}

func (s *Annotation_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitAnnotation_def(s)
	}
}

func (p *CapnProtoParser) Annotation_def() (localctx IAnnotation_defContext) {
	this := p
	_ = this

	localctx = NewAnnotation_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CapnProtoParserRULE_annotation_def)
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
		p.SetState(334)
		p.Match(CapnProtoParserT__24)
	}
	{
		p.SetState(335)
		p.Type_()
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__7 {
		{
			p.SetState(336)
			p.Annotation_parameters()
		}

	}
	{
		p.SetState(339)
		p.Match(CapnProtoParserT__14)
	}
	{
		p.SetState(340)
		p.Type_()
	}
	{
		p.SetState(341)
		p.Match(CapnProtoParserT__0)
	}

	return localctx
}

// IAnnotation_parametersContext is an interface to support dynamic dispatch.
type IAnnotation_parametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotation_parametersContext differentiates from other interfaces.
	IsAnnotation_parametersContext()
}

type Annotation_parametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_parametersContext() *Annotation_parametersContext {
	var p = new(Annotation_parametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_annotation_parameters
	return p
}

func (*Annotation_parametersContext) IsAnnotation_parametersContext() {}

func NewAnnotation_parametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_parametersContext {
	var p = new(Annotation_parametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_annotation_parameters

	return p
}

func (s *Annotation_parametersContext) GetParser() antlr.Parser { return s.parser }
func (s *Annotation_parametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_parametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_parametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterAnnotation_parameters(s)
	}
}

func (s *Annotation_parametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitAnnotation_parameters(s)
	}
}

func (p *CapnProtoParser) Annotation_parameters() (localctx IAnnotation_parametersContext) {
	this := p
	_ = this

	localctx = NewAnnotation_parametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CapnProtoParserRULE_annotation_parameters)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(CapnProtoParserT__7)
	}
	{
		p.SetState(344)
		p.Match(CapnProtoParserT__9)
	}
	{
		p.SetState(345)
		p.Match(CapnProtoParserT__8)
	}

	return localctx
}

// IConst_defContext is an interface to support dynamic dispatch.
type IConst_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConst_defContext differentiates from other interfaces.
	IsConst_defContext()
}

type Const_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_defContext() *Const_defContext {
	var p = new(Const_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_const_def
	return p
}

func (*Const_defContext) IsConst_defContext() {}

func NewConst_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_defContext {
	var p = new(Const_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_const_def

	return p
}

func (s *Const_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_defContext) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Const_defContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Const_defContext) Const_value() IConst_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Const_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterConst_def(s)
	}
}

func (s *Const_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitConst_def(s)
	}
}

func (p *CapnProtoParser) Const_def() (localctx IConst_defContext) {
	this := p
	_ = this

	localctx = NewConst_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CapnProtoParserRULE_const_def)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(CapnProtoParserT__25)
	}
	{
		p.SetState(348)
		p.Match(CapnProtoParserNAME)
	}
	{
		p.SetState(349)
		p.Match(CapnProtoParserT__14)
	}
	{
		p.SetState(350)
		p.Type_()
	}
	{
		p.SetState(351)
		p.Match(CapnProtoParserT__2)
	}
	{
		p.SetState(352)
		p.Const_value()
	}
	{
		p.SetState(353)
		p.Match(CapnProtoParserT__0)
	}

	return localctx
}

// IConst_valueContext is an interface to support dynamic dispatch.
type IConst_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConst_valueContext differentiates from other interfaces.
	IsConst_valueContext()
}

type Const_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_valueContext() *Const_valueContext {
	var p = new(Const_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_const_value
	return p
}

func (*Const_valueContext) IsConst_valueContext() {}

func NewConst_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_valueContext {
	var p = new(Const_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_const_value

	return p
}

func (s *Const_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_valueContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(CapnProtoParserNAME)
}

func (s *Const_valueContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, i)
}

func (s *Const_valueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserINTEGER, 0)
}

func (s *Const_valueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserFLOAT, 0)
}

func (s *Const_valueContext) TEXT() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserTEXT, 0)
}

func (s *Const_valueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserBOOLEAN, 0)
}

func (s *Const_valueContext) HEXADECIMAL() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserHEXADECIMAL, 0)
}

func (s *Const_valueContext) VOID() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserVOID, 0)
}

func (s *Const_valueContext) Literal_list() ILiteral_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_listContext)
}

func (s *Const_valueContext) Literal_union() ILiteral_unionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_unionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_unionContext)
}

func (s *Const_valueContext) Literal_bytes() ILiteral_bytesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_bytesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_bytesContext)
}

func (s *Const_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterConst_value(s)
	}
}

func (s *Const_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitConst_value(s)
	}
}

func (p *CapnProtoParser) Const_value() (localctx IConst_valueContext) {
	this := p
	_ = this

	localctx = NewConst_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CapnProtoParserRULE_const_value)
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

	p.SetState(375)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CapnProtoParserT__4, CapnProtoParserT__26, CapnProtoParserNAME:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CapnProtoParserT__26 {
			{
				p.SetState(355)
				p.Match(CapnProtoParserT__26)
			}

		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CapnProtoParserT__4 {
			{
				p.SetState(358)
				p.Match(CapnProtoParserT__4)
			}

		}
		{
			p.SetState(361)
			p.Match(CapnProtoParserNAME)
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CapnProtoParserT__4 {
			{
				p.SetState(362)
				p.Match(CapnProtoParserT__4)
			}
			{
				p.SetState(363)
				p.Match(CapnProtoParserNAME)
			}

		}

	case CapnProtoParserINTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(366)
			p.Match(CapnProtoParserINTEGER)
		}

	case CapnProtoParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(367)
			p.Match(CapnProtoParserFLOAT)
		}

	case CapnProtoParserTEXT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(368)
			p.Match(CapnProtoParserTEXT)
		}

	case CapnProtoParserBOOLEAN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(369)
			p.Match(CapnProtoParserBOOLEAN)
		}

	case CapnProtoParserHEXADECIMAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(370)
			p.Match(CapnProtoParserHEXADECIMAL)
		}

	case CapnProtoParserVOID:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(371)
			p.Match(CapnProtoParserVOID)
		}

	case CapnProtoParserT__22:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(372)
			p.Literal_list()
		}

	case CapnProtoParserT__7:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(373)
			p.Literal_union()
		}

	case CapnProtoParserT__27:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(374)
			p.Literal_bytes()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteral_unionContext is an interface to support dynamic dispatch.
type ILiteral_unionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteral_unionContext differentiates from other interfaces.
	IsLiteral_unionContext()
}

type Literal_unionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_unionContext() *Literal_unionContext {
	var p = new(Literal_unionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_literal_union
	return p
}

func (*Literal_unionContext) IsLiteral_unionContext() {}

func NewLiteral_unionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_unionContext {
	var p = new(Literal_unionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_literal_union

	return p
}

func (s *Literal_unionContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_unionContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(CapnProtoParserNAME)
}

func (s *Literal_unionContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, i)
}

func (s *Literal_unionContext) AllUnion_mapping() []IUnion_mappingContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnion_mappingContext)(nil)).Elem())
	var tst = make([]IUnion_mappingContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnion_mappingContext)
		}
	}

	return tst
}

func (s *Literal_unionContext) Union_mapping(i int) IUnion_mappingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnion_mappingContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnion_mappingContext)
}

func (s *Literal_unionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_unionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_unionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterLiteral_union(s)
	}
}

func (s *Literal_unionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitLiteral_union(s)
	}
}

func (p *CapnProtoParser) Literal_union() (localctx ILiteral_unionContext) {
	this := p
	_ = this

	localctx = NewLiteral_unionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CapnProtoParserRULE_literal_union)
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
		p.SetState(377)
		p.Match(CapnProtoParserT__7)
	}
	{
		p.SetState(378)
		p.Match(CapnProtoParserNAME)
	}
	{
		p.SetState(379)
		p.Match(CapnProtoParserT__2)
	}
	{
		p.SetState(380)
		p.Union_mapping()
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserT__15 {
		{
			p.SetState(381)
			p.Match(CapnProtoParserT__15)
		}
		{
			p.SetState(382)
			p.Match(CapnProtoParserNAME)
		}
		{
			p.SetState(383)
			p.Match(CapnProtoParserT__2)
		}
		{
			p.SetState(384)
			p.Union_mapping()
		}

		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(390)
		p.Match(CapnProtoParserT__8)
	}

	return localctx
}

// ILiteral_listContext is an interface to support dynamic dispatch.
type ILiteral_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteral_listContext differentiates from other interfaces.
	IsLiteral_listContext()
}

type Literal_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_listContext() *Literal_listContext {
	var p = new(Literal_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_literal_list
	return p
}

func (*Literal_listContext) IsLiteral_listContext() {}

func NewLiteral_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_listContext {
	var p = new(Literal_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_literal_list

	return p
}

func (s *Literal_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_listContext) AllConst_value() []IConst_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConst_valueContext)(nil)).Elem())
	var tst = make([]IConst_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConst_valueContext)
		}
	}

	return tst
}

func (s *Literal_listContext) Const_value(i int) IConst_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Literal_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterLiteral_list(s)
	}
}

func (s *Literal_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitLiteral_list(s)
	}
}

func (p *CapnProtoParser) Literal_list() (localctx ILiteral_listContext) {
	this := p
	_ = this

	localctx = NewLiteral_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CapnProtoParserRULE_literal_list)
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
		p.Match(CapnProtoParserT__22)
	}
	{
		p.SetState(393)
		p.Const_value()
	}
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserT__15 {
		{
			p.SetState(394)
			p.Match(CapnProtoParserT__15)
		}
		{
			p.SetState(395)
			p.Const_value()
		}

		p.SetState(400)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(401)
		p.Match(CapnProtoParserT__23)
	}

	return localctx
}

// ILiteral_bytesContext is an interface to support dynamic dispatch.
type ILiteral_bytesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteral_bytesContext differentiates from other interfaces.
	IsLiteral_bytesContext()
}

type Literal_bytesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_bytesContext() *Literal_bytesContext {
	var p = new(Literal_bytesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_literal_bytes
	return p
}

func (*Literal_bytesContext) IsLiteral_bytesContext() {}

func NewLiteral_bytesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_bytesContext {
	var p = new(Literal_bytesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_literal_bytes

	return p
}

func (s *Literal_bytesContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_bytesContext) TEXT() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserTEXT, 0)
}

func (s *Literal_bytesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_bytesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_bytesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterLiteral_bytes(s)
	}
}

func (s *Literal_bytesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitLiteral_bytes(s)
	}
}

func (p *CapnProtoParser) Literal_bytes() (localctx ILiteral_bytesContext) {
	this := p
	_ = this

	localctx = NewLiteral_bytesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CapnProtoParserRULE_literal_bytes)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.Match(CapnProtoParserT__27)
	}
	{
		p.SetState(404)
		p.Match(CapnProtoParserTEXT)
	}

	return localctx
}

// IUnion_mappingContext is an interface to support dynamic dispatch.
type IUnion_mappingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnion_mappingContext differentiates from other interfaces.
	IsUnion_mappingContext()
}

type Union_mappingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnion_mappingContext() *Union_mappingContext {
	var p = new(Union_mappingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_union_mapping
	return p
}

func (*Union_mappingContext) IsUnion_mappingContext() {}

func NewUnion_mappingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Union_mappingContext {
	var p = new(Union_mappingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_union_mapping

	return p
}

func (s *Union_mappingContext) GetParser() antlr.Parser { return s.parser }

func (s *Union_mappingContext) NAME() antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, 0)
}

func (s *Union_mappingContext) Const_value() IConst_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Union_mappingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Union_mappingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Union_mappingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterUnion_mapping(s)
	}
}

func (s *Union_mappingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitUnion_mapping(s)
	}
}

func (p *CapnProtoParser) Union_mapping() (localctx IUnion_mappingContext) {
	this := p
	_ = this

	localctx = NewUnion_mappingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CapnProtoParserRULE_union_mapping)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(406)
			p.Match(CapnProtoParserT__7)
		}
		{
			p.SetState(407)
			p.Match(CapnProtoParserNAME)
		}
		{
			p.SetState(408)
			p.Match(CapnProtoParserT__2)
		}
		{
			p.SetState(409)
			p.Const_value()
		}
		{
			p.SetState(410)
			p.Match(CapnProtoParserT__8)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(412)
			p.Const_value()
		}

	}

	return localctx
}

// IInner_usingContext is an interface to support dynamic dispatch.
type IInner_usingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInner_usingContext differentiates from other interfaces.
	IsInner_usingContext()
}

type Inner_usingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInner_usingContext() *Inner_usingContext {
	var p = new(Inner_usingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CapnProtoParserRULE_inner_using
	return p
}

func (*Inner_usingContext) IsInner_usingContext() {}

func NewInner_usingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inner_usingContext {
	var p = new(Inner_usingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CapnProtoParserRULE_inner_using

	return p
}

func (s *Inner_usingContext) GetParser() antlr.Parser { return s.parser }

func (s *Inner_usingContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(CapnProtoParserNAME)
}

func (s *Inner_usingContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(CapnProtoParserNAME, i)
}

func (s *Inner_usingContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Inner_usingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inner_usingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inner_usingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.EnterInner_using(s)
	}
}

func (s *Inner_usingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CapnProtoListener); ok {
		listenerT.ExitInner_using(s)
	}
}

func (p *CapnProtoParser) Inner_using() (localctx IInner_usingContext) {
	this := p
	_ = this

	localctx = NewInner_usingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CapnProtoParserRULE_inner_using)
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
		p.Match(CapnProtoParserT__1)
	}
	{
		p.SetState(416)
		p.Match(CapnProtoParserNAME)
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CapnProtoParserT__4 {
		{
			p.SetState(417)
			p.Match(CapnProtoParserT__4)
		}
		{
			p.SetState(418)
			p.Match(CapnProtoParserNAME)
		}

		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CapnProtoParserT__2 {
		{
			p.SetState(424)
			p.Match(CapnProtoParserT__2)
		}
		{
			p.SetState(425)
			p.Type_()
		}

	}
	{
		p.SetState(428)
		p.Match(CapnProtoParserT__0)
	}

	return localctx
}
