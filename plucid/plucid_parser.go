// Code generated from plucid.g4 by ANTLR 4.9.3. DO NOT EDIT.

package plucid // plucid
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 57, 322,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 102,
	10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 108, 10, 3, 12, 3, 14, 3, 111, 11,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 117, 10, 4, 3, 5, 3, 5, 5, 5, 121, 10,
	5, 3, 6, 6, 6, 124, 10, 6, 13, 6, 14, 6, 125, 3, 6, 3, 6, 5, 6, 130, 10,
	6, 3, 7, 3, 7, 3, 7, 7, 7, 135, 10, 7, 12, 7, 14, 7, 138, 11, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 7, 9, 146, 10, 9, 12, 9, 14, 9, 149, 11,
	9, 3, 9, 6, 9, 152, 10, 9, 13, 9, 14, 9, 153, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 160, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 166, 10, 10, 12, 10,
	14, 10, 169, 11, 10, 3, 10, 5, 10, 172, 10, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 5, 11, 178, 10, 11, 3, 12, 3, 12, 7, 12, 182, 10, 12, 12, 12, 14, 12,
	185, 11, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 7, 15, 194,
	10, 15, 12, 15, 14, 15, 197, 11, 15, 3, 15, 5, 15, 200, 10, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 7, 16, 206, 10, 16, 12, 16, 14, 16, 209, 11, 16, 5,
	16, 211, 10, 16, 3, 17, 3, 17, 5, 17, 215, 10, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 5, 19, 233, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 7,
	22, 250, 10, 22, 12, 22, 14, 22, 253, 11, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 5, 25, 271, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 7, 28, 284, 10, 28, 12, 28, 14, 28,
	287, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 7,
	30, 297, 10, 30, 12, 30, 14, 30, 300, 11, 30, 3, 31, 3, 31, 5, 31, 304,
	10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 320, 10, 34, 3, 34, 2, 3, 4,
	35, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 2, 4, 3, 2,
	43, 48, 3, 2, 49, 53, 2, 335, 2, 68, 3, 2, 2, 2, 4, 101, 3, 2, 2, 2, 6,
	116, 3, 2, 2, 2, 8, 120, 3, 2, 2, 2, 10, 129, 3, 2, 2, 2, 12, 131, 3, 2,
	2, 2, 14, 139, 3, 2, 2, 2, 16, 159, 3, 2, 2, 2, 18, 171, 3, 2, 2, 2, 20,
	177, 3, 2, 2, 2, 22, 179, 3, 2, 2, 2, 24, 186, 3, 2, 2, 2, 26, 188, 3,
	2, 2, 2, 28, 199, 3, 2, 2, 2, 30, 210, 3, 2, 2, 2, 32, 214, 3, 2, 2, 2,
	34, 216, 3, 2, 2, 2, 36, 232, 3, 2, 2, 2, 38, 234, 3, 2, 2, 2, 40, 240,
	3, 2, 2, 2, 42, 251, 3, 2, 2, 2, 44, 256, 3, 2, 2, 2, 46, 260, 3, 2, 2,
	2, 48, 270, 3, 2, 2, 2, 50, 272, 3, 2, 2, 2, 52, 277, 3, 2, 2, 2, 54, 285,
	3, 2, 2, 2, 56, 288, 3, 2, 2, 2, 58, 298, 3, 2, 2, 2, 60, 303, 3, 2, 2,
	2, 62, 305, 3, 2, 2, 2, 64, 309, 3, 2, 2, 2, 66, 319, 3, 2, 2, 2, 68, 69,
	5, 4, 3, 2, 69, 3, 3, 2, 2, 2, 70, 71, 8, 3, 1, 2, 71, 102, 5, 6, 4, 2,
	72, 102, 5, 22, 12, 2, 73, 102, 7, 3, 2, 2, 74, 102, 7, 4, 2, 2, 75, 76,
	5, 24, 13, 2, 76, 77, 5, 4, 3, 13, 77, 102, 3, 2, 2, 2, 78, 79, 7, 5, 2,
	2, 79, 80, 5, 4, 3, 2, 80, 81, 7, 6, 2, 2, 81, 82, 5, 4, 3, 2, 82, 83,
	7, 6, 2, 2, 83, 84, 5, 4, 3, 2, 84, 102, 3, 2, 2, 2, 85, 86, 7, 7, 2, 2,
	86, 87, 5, 4, 3, 2, 87, 88, 7, 6, 2, 2, 88, 89, 5, 4, 3, 2, 89, 90, 7,
	6, 2, 2, 90, 91, 5, 4, 3, 2, 91, 102, 3, 2, 2, 2, 92, 93, 7, 8, 2, 2, 93,
	102, 5, 4, 3, 9, 94, 95, 7, 9, 2, 2, 95, 102, 5, 4, 3, 8, 96, 102, 5, 28,
	15, 2, 97, 102, 5, 34, 18, 2, 98, 102, 5, 38, 20, 2, 99, 102, 5, 40, 21,
	2, 100, 102, 5, 46, 24, 2, 101, 70, 3, 2, 2, 2, 101, 72, 3, 2, 2, 2, 101,
	73, 3, 2, 2, 2, 101, 74, 3, 2, 2, 2, 101, 75, 3, 2, 2, 2, 101, 78, 3, 2,
	2, 2, 101, 85, 3, 2, 2, 2, 101, 92, 3, 2, 2, 2, 101, 94, 3, 2, 2, 2, 101,
	96, 3, 2, 2, 2, 101, 97, 3, 2, 2, 2, 101, 98, 3, 2, 2, 2, 101, 99, 3, 2,
	2, 2, 101, 100, 3, 2, 2, 2, 102, 109, 3, 2, 2, 2, 103, 104, 12, 12, 2,
	2, 104, 105, 5, 26, 14, 2, 105, 106, 5, 4, 3, 13, 106, 108, 3, 2, 2, 2,
	107, 103, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109,
	110, 3, 2, 2, 2, 110, 5, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 117, 5,
	8, 5, 2, 113, 117, 5, 14, 8, 2, 114, 117, 7, 42, 2, 2, 115, 117, 5, 18,
	10, 2, 116, 112, 3, 2, 2, 2, 116, 113, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2,
	116, 115, 3, 2, 2, 2, 117, 7, 3, 2, 2, 2, 118, 121, 5, 10, 6, 2, 119, 121,
	5, 12, 7, 2, 120, 118, 3, 2, 2, 2, 120, 119, 3, 2, 2, 2, 121, 9, 3, 2,
	2, 2, 122, 124, 7, 55, 2, 2, 123, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2,
	125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 130, 3, 2, 2, 2, 127,
	128, 7, 36, 2, 2, 128, 130, 5, 10, 6, 2, 129, 123, 3, 2, 2, 2, 129, 127,
	3, 2, 2, 2, 130, 11, 3, 2, 2, 2, 131, 132, 5, 10, 6, 2, 132, 136, 7, 40,
	2, 2, 133, 135, 7, 55, 2, 2, 134, 133, 3, 2, 2, 2, 135, 138, 3, 2, 2, 2,
	136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 13, 3, 2, 2, 2, 138, 136,
	3, 2, 2, 2, 139, 140, 7, 38, 2, 2, 140, 141, 5, 16, 9, 2, 141, 142, 7,
	38, 2, 2, 142, 15, 3, 2, 2, 2, 143, 147, 7, 56, 2, 2, 144, 146, 7, 54,
	2, 2, 145, 144, 3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2,
	147, 148, 3, 2, 2, 2, 148, 160, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150,
	152, 7, 37, 2, 2, 151, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 151,
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 160, 3, 2, 2, 2, 155, 160, 7, 39,
	2, 2, 156, 160, 7, 40, 2, 2, 157, 160, 7, 41, 2, 2, 158, 160, 7, 38, 2,
	2, 159, 143, 3, 2, 2, 2, 159, 151, 3, 2, 2, 2, 159, 155, 3, 2, 2, 2, 159,
	156, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 158, 3, 2, 2, 2, 160, 17, 3,
	2, 2, 2, 161, 172, 7, 10, 2, 2, 162, 172, 7, 11, 2, 2, 163, 167, 7, 12,
	2, 2, 164, 166, 5, 20, 11, 2, 165, 164, 3, 2, 2, 2, 166, 169, 3, 2, 2,
	2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 170, 3, 2, 2, 2, 169,
	167, 3, 2, 2, 2, 170, 172, 7, 13, 2, 2, 171, 161, 3, 2, 2, 2, 171, 162,
	3, 2, 2, 2, 171, 163, 3, 2, 2, 2, 172, 19, 3, 2, 2, 2, 173, 178, 5, 8,
	5, 2, 174, 178, 5, 16, 9, 2, 175, 178, 7, 42, 2, 2, 176, 178, 5, 18, 10,
	2, 177, 173, 3, 2, 2, 2, 177, 174, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177,
	176, 3, 2, 2, 2, 178, 21, 3, 2, 2, 2, 179, 183, 7, 56, 2, 2, 180, 182,
	7, 54, 2, 2, 181, 180, 3, 2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 181, 3, 2,
	2, 2, 183, 184, 3, 2, 2, 2, 184, 23, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2,
	186, 187, 9, 2, 2, 2, 187, 25, 3, 2, 2, 2, 188, 189, 9, 3, 2, 2, 189, 27,
	3, 2, 2, 2, 190, 200, 7, 14, 2, 2, 191, 195, 7, 15, 2, 2, 192, 194, 5,
	30, 16, 2, 193, 192, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193, 3, 2,
	2, 2, 195, 196, 3, 2, 2, 2, 196, 198, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2,
	198, 200, 7, 16, 2, 2, 199, 190, 3, 2, 2, 2, 199, 191, 3, 2, 2, 2, 200,
	29, 3, 2, 2, 2, 201, 211, 5, 32, 17, 2, 202, 203, 5, 32, 17, 2, 203, 207,
	7, 6, 2, 2, 204, 206, 5, 30, 16, 2, 205, 204, 3, 2, 2, 2, 206, 209, 3,
	2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 211, 3, 2, 2,
	2, 209, 207, 3, 2, 2, 2, 210, 201, 3, 2, 2, 2, 210, 202, 3, 2, 2, 2, 211,
	31, 3, 2, 2, 2, 212, 215, 5, 4, 3, 2, 213, 215, 5, 28, 15, 2, 214, 212,
	3, 2, 2, 2, 214, 213, 3, 2, 2, 2, 215, 33, 3, 2, 2, 2, 216, 217, 7, 17,
	2, 2, 217, 218, 5, 4, 3, 2, 218, 219, 7, 18, 2, 2, 219, 220, 5, 4, 3, 2,
	220, 221, 5, 36, 19, 2, 221, 35, 3, 2, 2, 2, 222, 223, 7, 19, 2, 2, 223,
	224, 5, 4, 3, 2, 224, 225, 7, 20, 2, 2, 225, 233, 3, 2, 2, 2, 226, 227,
	7, 21, 2, 2, 227, 228, 5, 4, 3, 2, 228, 229, 7, 18, 2, 2, 229, 230, 5,
	4, 3, 2, 230, 231, 7, 22, 2, 2, 231, 233, 3, 2, 2, 2, 232, 222, 3, 2, 2,
	2, 232, 226, 3, 2, 2, 2, 233, 37, 3, 2, 2, 2, 234, 235, 7, 23, 2, 2, 235,
	236, 5, 4, 3, 2, 236, 237, 7, 24, 2, 2, 237, 238, 5, 42, 22, 2, 238, 239,
	7, 25, 2, 2, 239, 39, 3, 2, 2, 2, 240, 241, 7, 26, 2, 2, 241, 242, 5, 42,
	22, 2, 242, 243, 7, 25, 2, 2, 243, 41, 3, 2, 2, 2, 244, 245, 5, 4, 3, 2,
	245, 246, 7, 27, 2, 2, 246, 247, 5, 4, 3, 2, 247, 248, 7, 28, 2, 2, 248,
	250, 3, 2, 2, 2, 249, 244, 3, 2, 2, 2, 250, 253, 3, 2, 2, 2, 251, 249,
	3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 254, 3, 2, 2, 2, 253, 251, 3, 2,
	2, 2, 254, 255, 5, 44, 23, 2, 255, 43, 3, 2, 2, 2, 256, 257, 7, 29, 2,
	2, 257, 258, 7, 27, 2, 2, 258, 259, 5, 4, 3, 2, 259, 45, 3, 2, 2, 2, 260,
	261, 5, 22, 12, 2, 261, 262, 7, 30, 2, 2, 262, 263, 5, 48, 25, 2, 263,
	264, 7, 31, 2, 2, 264, 47, 3, 2, 2, 2, 265, 271, 5, 4, 3, 2, 266, 267,
	5, 4, 3, 2, 267, 268, 7, 6, 2, 2, 268, 269, 5, 48, 25, 2, 269, 271, 3,
	2, 2, 2, 270, 265, 3, 2, 2, 2, 270, 266, 3, 2, 2, 2, 271, 49, 3, 2, 2,
	2, 272, 273, 5, 4, 3, 2, 273, 274, 7, 32, 2, 2, 274, 275, 5, 52, 27, 2,
	275, 276, 7, 25, 2, 2, 276, 51, 3, 2, 2, 2, 277, 278, 5, 54, 28, 2, 278,
	279, 5, 58, 30, 2, 279, 53, 3, 2, 2, 2, 280, 281, 5, 56, 29, 2, 281, 282,
	7, 28, 2, 2, 282, 284, 3, 2, 2, 2, 283, 280, 3, 2, 2, 2, 284, 287, 3, 2,
	2, 2, 285, 283, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 55, 3, 2, 2, 2,
	287, 285, 3, 2, 2, 2, 288, 289, 5, 22, 12, 2, 289, 290, 7, 33, 2, 2, 290,
	291, 7, 34, 2, 2, 291, 292, 5, 4, 3, 2, 292, 57, 3, 2, 2, 2, 293, 294,
	5, 60, 31, 2, 294, 295, 7, 28, 2, 2, 295, 297, 3, 2, 2, 2, 296, 293, 3,
	2, 2, 2, 297, 300, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2,
	2, 299, 59, 3, 2, 2, 2, 300, 298, 3, 2, 2, 2, 301, 304, 5, 62, 32, 2, 302,
	304, 5, 64, 33, 2, 303, 301, 3, 2, 2, 2, 303, 302, 3, 2, 2, 2, 304, 61,
	3, 2, 2, 2, 305, 306, 5, 22, 12, 2, 306, 307, 7, 35, 2, 2, 307, 308, 5,
	4, 3, 2, 308, 63, 3, 2, 2, 2, 309, 310, 5, 22, 12, 2, 310, 311, 5, 66,
	34, 2, 311, 312, 7, 35, 2, 2, 312, 313, 5, 4, 3, 2, 313, 65, 3, 2, 2, 2,
	314, 320, 5, 22, 12, 2, 315, 316, 5, 22, 12, 2, 316, 317, 7, 6, 2, 2, 317,
	318, 5, 66, 34, 2, 318, 320, 3, 2, 2, 2, 319, 314, 3, 2, 2, 2, 319, 315,
	3, 2, 2, 2, 320, 67, 3, 2, 2, 2, 28, 101, 109, 116, 120, 125, 129, 136,
	147, 153, 159, 167, 171, 177, 183, 195, 199, 207, 210, 214, 232, 251, 270,
	285, 298, 303, 319,
}
var literalNames = []string{
	"", "'error'", "'eod'", "'filter'", "','", "'substr'", "'length'", "'arg'",
	"'nil'", "'[]'", "'['", "']'", "'[%%]'", "'[%'", "'%]'", "'if'", "'then'",
	"'else'", "'fi'", "'elseif'", "'endif'", "'case'", "'of'", "'end'", "'cond'",
	"':'", "';'", "'default'", "'('", "')'", "'where'", "'is'", "'current'",
	"'='", "'\u02DC'", "", "'\"'", "", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "N_SIGN",
	"SIGN", "QUOTE", "BRACKET", "PERIOD", "SEPARATOR", "STRING_CONSTANT", "P_NUMERIC_OPERATOR",
	"P_WORD_OPERATOR", "P_STRING_OPERATOR", "P_LIST_OPERATOR", "P_LUCID_OPERATOR",
	"P_SPECIAL_OPERATOR", "I_NUMERIC_OPERATOR", "I_WORD_OPERATOR", "I_STRING_OPERATOR",
	"I_LIST_OPERATOR", "I_LUCID_OPERATOR", "ALPHANUMERIC", "DIGIT", "LETTER",
	"WS",
}

var ruleNames = []string{
	"program", "expression", "constant", "numeric_constant", "integer_constant",
	"real_constant", "word_constant", "word_constant_less_the_quotes", "list_constant",
	"list_constant_element", "identifier", "prefix_operator", "infix_operator",
	"list_expression", "expressions_list", "expression_item", "if_expression",
	"endif", "case_expression", "cond_expression", "cbody", "defaultcase",
	"function_call", "actuals_list", "where_clause", "body", "declarations_list",
	"current_declaration", "definitions_list", "definition", "simple_definition",
	"function_definition", "formals_list",
}

type plucidParser struct {
	*antlr.BaseParser
}

// NewplucidParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *plucidParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewplucidParser(input antlr.TokenStream) *plucidParser {
	this := new(plucidParser)
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
	this.GrammarFileName = "plucid.g4"

	return this
}

// plucidParser tokens.
const (
	plucidParserEOF                = antlr.TokenEOF
	plucidParserT__0               = 1
	plucidParserT__1               = 2
	plucidParserT__2               = 3
	plucidParserT__3               = 4
	plucidParserT__4               = 5
	plucidParserT__5               = 6
	plucidParserT__6               = 7
	plucidParserT__7               = 8
	plucidParserT__8               = 9
	plucidParserT__9               = 10
	plucidParserT__10              = 11
	plucidParserT__11              = 12
	plucidParserT__12              = 13
	plucidParserT__13              = 14
	plucidParserT__14              = 15
	plucidParserT__15              = 16
	plucidParserT__16              = 17
	plucidParserT__17              = 18
	plucidParserT__18              = 19
	plucidParserT__19              = 20
	plucidParserT__20              = 21
	plucidParserT__21              = 22
	plucidParserT__22              = 23
	plucidParserT__23              = 24
	plucidParserT__24              = 25
	plucidParserT__25              = 26
	plucidParserT__26              = 27
	plucidParserT__27              = 28
	plucidParserT__28              = 29
	plucidParserT__29              = 30
	plucidParserT__30              = 31
	plucidParserT__31              = 32
	plucidParserT__32              = 33
	plucidParserN_SIGN             = 34
	plucidParserSIGN               = 35
	plucidParserQUOTE              = 36
	plucidParserBRACKET            = 37
	plucidParserPERIOD             = 38
	plucidParserSEPARATOR          = 39
	plucidParserSTRING_CONSTANT    = 40
	plucidParserP_NUMERIC_OPERATOR = 41
	plucidParserP_WORD_OPERATOR    = 42
	plucidParserP_STRING_OPERATOR  = 43
	plucidParserP_LIST_OPERATOR    = 44
	plucidParserP_LUCID_OPERATOR   = 45
	plucidParserP_SPECIAL_OPERATOR = 46
	plucidParserI_NUMERIC_OPERATOR = 47
	plucidParserI_WORD_OPERATOR    = 48
	plucidParserI_STRING_OPERATOR  = 49
	plucidParserI_LIST_OPERATOR    = 50
	plucidParserI_LUCID_OPERATOR   = 51
	plucidParserALPHANUMERIC       = 52
	plucidParserDIGIT              = 53
	plucidParserLETTER             = 54
	plucidParserWS                 = 55
)

// plucidParser rules.
const (
	plucidParserRULE_program                       = 0
	plucidParserRULE_expression                    = 1
	plucidParserRULE_constant                      = 2
	plucidParserRULE_numeric_constant              = 3
	plucidParserRULE_integer_constant              = 4
	plucidParserRULE_real_constant                 = 5
	plucidParserRULE_word_constant                 = 6
	plucidParserRULE_word_constant_less_the_quotes = 7
	plucidParserRULE_list_constant                 = 8
	plucidParserRULE_list_constant_element         = 9
	plucidParserRULE_identifier                    = 10
	plucidParserRULE_prefix_operator               = 11
	plucidParserRULE_infix_operator                = 12
	plucidParserRULE_list_expression               = 13
	plucidParserRULE_expressions_list              = 14
	plucidParserRULE_expression_item               = 15
	plucidParserRULE_if_expression                 = 16
	plucidParserRULE_endif                         = 17
	plucidParserRULE_case_expression               = 18
	plucidParserRULE_cond_expression               = 19
	plucidParserRULE_cbody                         = 20
	plucidParserRULE_defaultcase                   = 21
	plucidParserRULE_function_call                 = 22
	plucidParserRULE_actuals_list                  = 23
	plucidParserRULE_where_clause                  = 24
	plucidParserRULE_body                          = 25
	plucidParserRULE_declarations_list             = 26
	plucidParserRULE_current_declaration           = 27
	plucidParserRULE_definitions_list              = 28
	plucidParserRULE_definition                    = 29
	plucidParserRULE_simple_definition             = 30
	plucidParserRULE_function_definition           = 31
	plucidParserRULE_formals_list                  = 32
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *plucidParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, plucidParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.expression(0)
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
	p.RuleIndex = plucidParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpressionContext) Prefix_operator() IPrefix_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefix_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefix_operatorContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *ExpressionContext) If_expression() IIf_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_expressionContext)
}

func (s *ExpressionContext) Case_expression() ICase_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICase_expressionContext)
}

func (s *ExpressionContext) Cond_expression() ICond_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICond_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICond_expressionContext)
}

func (s *ExpressionContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *ExpressionContext) Infix_operator() IInfix_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfix_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfix_operatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *plucidParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *plucidParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, plucidParserRULE_expression, _p)

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
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(69)
			p.Constant()
		}

	case 2:
		{
			p.SetState(70)
			p.Identifier()
		}

	case 3:
		{
			p.SetState(71)
			p.Match(plucidParserT__0)
		}

	case 4:
		{
			p.SetState(72)
			p.Match(plucidParserT__1)
		}

	case 5:
		{
			p.SetState(73)
			p.Prefix_operator()
		}
		{
			p.SetState(74)
			p.expression(11)
		}

	case 6:
		{
			p.SetState(76)
			p.Match(plucidParserT__2)
		}

		{
			p.SetState(77)
			p.expression(0)
		}
		{
			p.SetState(78)
			p.Match(plucidParserT__3)
		}
		{
			p.SetState(79)
			p.expression(0)
		}
		{
			p.SetState(80)
			p.Match(plucidParserT__3)
		}
		{
			p.SetState(81)
			p.expression(0)
		}

	case 7:
		{
			p.SetState(83)
			p.Match(plucidParserT__4)
		}

		{
			p.SetState(84)
			p.expression(0)
		}
		{
			p.SetState(85)
			p.Match(plucidParserT__3)
		}
		{
			p.SetState(86)
			p.expression(0)
		}
		{
			p.SetState(87)
			p.Match(plucidParserT__3)
		}
		{
			p.SetState(88)
			p.expression(0)
		}

	case 8:
		{
			p.SetState(90)
			p.Match(plucidParserT__5)
		}
		{
			p.SetState(91)
			p.expression(7)
		}

	case 9:
		{
			p.SetState(92)
			p.Match(plucidParserT__6)
		}
		{
			p.SetState(93)
			p.expression(6)
		}

	case 10:
		{
			p.SetState(94)
			p.List_expression()
		}

	case 11:
		{
			p.SetState(95)
			p.If_expression()
		}

	case 12:
		{
			p.SetState(96)
			p.Case_expression()
		}

	case 13:
		{
			p.SetState(97)
			p.Cond_expression()
		}

	case 14:
		{
			p.SetState(98)
			p.Function_call()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, plucidParserRULE_expression)
			p.SetState(101)

			if !(p.Precpred(p.GetParserRuleContext(), 10)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
			}
			{
				p.SetState(102)
				p.Infix_operator()
			}
			{
				p.SetState(103)
				p.expression(11)
			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
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
	p.RuleIndex = plucidParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) Numeric_constant() INumeric_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_constantContext)
}

func (s *ConstantContext) Word_constant() IWord_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWord_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWord_constantContext)
}

func (s *ConstantContext) STRING_CONSTANT() antlr.TerminalNode {
	return s.GetToken(plucidParserSTRING_CONSTANT, 0)
}

func (s *ConstantContext) List_constant() IList_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_constantContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *plucidParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, plucidParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(114)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case plucidParserN_SIGN, plucidParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Numeric_constant()
		}

	case plucidParserQUOTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Word_constant()
		}

	case plucidParserSTRING_CONSTANT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Match(plucidParserSTRING_CONSTANT)
		}

	case plucidParserT__7, plucidParserT__8, plucidParserT__9:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.List_constant()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumeric_constantContext is an interface to support dynamic dispatch.
type INumeric_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumeric_constantContext differentiates from other interfaces.
	IsNumeric_constantContext()
}

type Numeric_constantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumeric_constantContext() *Numeric_constantContext {
	var p = new(Numeric_constantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_numeric_constant
	return p
}

func (*Numeric_constantContext) IsNumeric_constantContext() {}

func NewNumeric_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numeric_constantContext {
	var p = new(Numeric_constantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_numeric_constant

	return p
}

func (s *Numeric_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Numeric_constantContext) Integer_constant() IInteger_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_constantContext)
}

func (s *Numeric_constantContext) Real_constant() IReal_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReal_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReal_constantContext)
}

func (s *Numeric_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numeric_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numeric_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterNumeric_constant(s)
	}
}

func (s *Numeric_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitNumeric_constant(s)
	}
}

func (p *plucidParser) Numeric_constant() (localctx INumeric_constantContext) {
	this := p
	_ = this

	localctx = NewNumeric_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, plucidParserRULE_numeric_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Integer_constant()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Real_constant()
		}

	}

	return localctx
}

// IInteger_constantContext is an interface to support dynamic dispatch.
type IInteger_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_constantContext differentiates from other interfaces.
	IsInteger_constantContext()
}

type Integer_constantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_constantContext() *Integer_constantContext {
	var p = new(Integer_constantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_integer_constant
	return p
}

func (*Integer_constantContext) IsInteger_constantContext() {}

func NewInteger_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_constantContext {
	var p = new(Integer_constantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_integer_constant

	return p
}

func (s *Integer_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_constantContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(plucidParserDIGIT)
}

func (s *Integer_constantContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(plucidParserDIGIT, i)
}

func (s *Integer_constantContext) N_SIGN() antlr.TerminalNode {
	return s.GetToken(plucidParserN_SIGN, 0)
}

func (s *Integer_constantContext) Integer_constant() IInteger_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_constantContext)
}

func (s *Integer_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterInteger_constant(s)
	}
}

func (s *Integer_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitInteger_constant(s)
	}
}

func (p *plucidParser) Integer_constant() (localctx IInteger_constantContext) {
	this := p
	_ = this

	localctx = NewInteger_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, plucidParserRULE_integer_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case plucidParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(120)
					p.Match(plucidParserDIGIT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(123)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}

	case plucidParserN_SIGN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(plucidParserN_SIGN)
		}
		{
			p.SetState(126)
			p.Integer_constant()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReal_constantContext is an interface to support dynamic dispatch.
type IReal_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReal_constantContext differentiates from other interfaces.
	IsReal_constantContext()
}

type Real_constantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_constantContext() *Real_constantContext {
	var p = new(Real_constantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_real_constant
	return p
}

func (*Real_constantContext) IsReal_constantContext() {}

func NewReal_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_constantContext {
	var p = new(Real_constantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_real_constant

	return p
}

func (s *Real_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Real_constantContext) Integer_constant() IInteger_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_constantContext)
}

func (s *Real_constantContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(plucidParserPERIOD, 0)
}

func (s *Real_constantContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(plucidParserDIGIT)
}

func (s *Real_constantContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(plucidParserDIGIT, i)
}

func (s *Real_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterReal_constant(s)
	}
}

func (s *Real_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitReal_constant(s)
	}
}

func (p *plucidParser) Real_constant() (localctx IReal_constantContext) {
	this := p
	_ = this

	localctx = NewReal_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, plucidParserRULE_real_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(129)
		p.Integer_constant()
	}
	{
		p.SetState(130)
		p.Match(plucidParserPERIOD)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(131)
				p.Match(plucidParserDIGIT)
			}

		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IWord_constantContext is an interface to support dynamic dispatch.
type IWord_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWord_constantContext differentiates from other interfaces.
	IsWord_constantContext()
}

type Word_constantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWord_constantContext() *Word_constantContext {
	var p = new(Word_constantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_word_constant
	return p
}

func (*Word_constantContext) IsWord_constantContext() {}

func NewWord_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Word_constantContext {
	var p = new(Word_constantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_word_constant

	return p
}

func (s *Word_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Word_constantContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(plucidParserQUOTE)
}

func (s *Word_constantContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(plucidParserQUOTE, i)
}

func (s *Word_constantContext) Word_constant_less_the_quotes() IWord_constant_less_the_quotesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWord_constant_less_the_quotesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWord_constant_less_the_quotesContext)
}

func (s *Word_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Word_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Word_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterWord_constant(s)
	}
}

func (s *Word_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitWord_constant(s)
	}
}

func (p *plucidParser) Word_constant() (localctx IWord_constantContext) {
	this := p
	_ = this

	localctx = NewWord_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, plucidParserRULE_word_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(plucidParserQUOTE)
	}
	{
		p.SetState(138)
		p.Word_constant_less_the_quotes()
	}
	{
		p.SetState(139)
		p.Match(plucidParserQUOTE)
	}

	return localctx
}

// IWord_constant_less_the_quotesContext is an interface to support dynamic dispatch.
type IWord_constant_less_the_quotesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWord_constant_less_the_quotesContext differentiates from other interfaces.
	IsWord_constant_less_the_quotesContext()
}

type Word_constant_less_the_quotesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWord_constant_less_the_quotesContext() *Word_constant_less_the_quotesContext {
	var p = new(Word_constant_less_the_quotesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_word_constant_less_the_quotes
	return p
}

func (*Word_constant_less_the_quotesContext) IsWord_constant_less_the_quotesContext() {}

func NewWord_constant_less_the_quotesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Word_constant_less_the_quotesContext {
	var p = new(Word_constant_less_the_quotesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_word_constant_less_the_quotes

	return p
}

func (s *Word_constant_less_the_quotesContext) GetParser() antlr.Parser { return s.parser }

func (s *Word_constant_less_the_quotesContext) LETTER() antlr.TerminalNode {
	return s.GetToken(plucidParserLETTER, 0)
}

func (s *Word_constant_less_the_quotesContext) AllALPHANUMERIC() []antlr.TerminalNode {
	return s.GetTokens(plucidParserALPHANUMERIC)
}

func (s *Word_constant_less_the_quotesContext) ALPHANUMERIC(i int) antlr.TerminalNode {
	return s.GetToken(plucidParserALPHANUMERIC, i)
}

func (s *Word_constant_less_the_quotesContext) AllSIGN() []antlr.TerminalNode {
	return s.GetTokens(plucidParserSIGN)
}

func (s *Word_constant_less_the_quotesContext) SIGN(i int) antlr.TerminalNode {
	return s.GetToken(plucidParserSIGN, i)
}

func (s *Word_constant_less_the_quotesContext) BRACKET() antlr.TerminalNode {
	return s.GetToken(plucidParserBRACKET, 0)
}

func (s *Word_constant_less_the_quotesContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(plucidParserPERIOD, 0)
}

func (s *Word_constant_less_the_quotesContext) SEPARATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserSEPARATOR, 0)
}

func (s *Word_constant_less_the_quotesContext) QUOTE() antlr.TerminalNode {
	return s.GetToken(plucidParserQUOTE, 0)
}

func (s *Word_constant_less_the_quotesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Word_constant_less_the_quotesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Word_constant_less_the_quotesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterWord_constant_less_the_quotes(s)
	}
}

func (s *Word_constant_less_the_quotesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitWord_constant_less_the_quotes(s)
	}
}

func (p *plucidParser) Word_constant_less_the_quotes() (localctx IWord_constant_less_the_quotesContext) {
	this := p
	_ = this

	localctx = NewWord_constant_less_the_quotesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, plucidParserRULE_word_constant_less_the_quotes)
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

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case plucidParserLETTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(plucidParserLETTER)
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == plucidParserALPHANUMERIC {
			{
				p.SetState(142)
				p.Match(plucidParserALPHANUMERIC)
			}

			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case plucidParserSIGN:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(148)
					p.Match(plucidParserSIGN)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(151)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
		}

	case plucidParserBRACKET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
			p.Match(plucidParserBRACKET)
		}

	case plucidParserPERIOD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(154)
			p.Match(plucidParserPERIOD)
		}

	case plucidParserSEPARATOR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(155)
			p.Match(plucidParserSEPARATOR)
		}

	case plucidParserQUOTE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(156)
			p.Match(plucidParserQUOTE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IList_constantContext is an interface to support dynamic dispatch.
type IList_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_constantContext differentiates from other interfaces.
	IsList_constantContext()
}

type List_constantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_constantContext() *List_constantContext {
	var p = new(List_constantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_list_constant
	return p
}

func (*List_constantContext) IsList_constantContext() {}

func NewList_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_constantContext {
	var p = new(List_constantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_list_constant

	return p
}

func (s *List_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *List_constantContext) AllList_constant_element() []IList_constant_elementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IList_constant_elementContext)(nil)).Elem())
	var tst = make([]IList_constant_elementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IList_constant_elementContext)
		}
	}

	return tst
}

func (s *List_constantContext) List_constant_element(i int) IList_constant_elementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_constant_elementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IList_constant_elementContext)
}

func (s *List_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterList_constant(s)
	}
}

func (s *List_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitList_constant(s)
	}
}

func (p *plucidParser) List_constant() (localctx IList_constantContext) {
	this := p
	_ = this

	localctx = NewList_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, plucidParserRULE_list_constant)
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
	case plucidParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Match(plucidParserT__7)
		}

	case plucidParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.Match(plucidParserT__8)
		}

	case plucidParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.Match(plucidParserT__9)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<plucidParserT__7)|(1<<plucidParserT__8)|(1<<plucidParserT__9))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(plucidParserN_SIGN-34))|(1<<(plucidParserSIGN-34))|(1<<(plucidParserQUOTE-34))|(1<<(plucidParserBRACKET-34))|(1<<(plucidParserPERIOD-34))|(1<<(plucidParserSEPARATOR-34))|(1<<(plucidParserSTRING_CONSTANT-34))|(1<<(plucidParserDIGIT-34))|(1<<(plucidParserLETTER-34)))) != 0) {
			{
				p.SetState(162)
				p.List_constant_element()
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(168)
			p.Match(plucidParserT__10)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IList_constant_elementContext is an interface to support dynamic dispatch.
type IList_constant_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_constant_elementContext differentiates from other interfaces.
	IsList_constant_elementContext()
}

type List_constant_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_constant_elementContext() *List_constant_elementContext {
	var p = new(List_constant_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_list_constant_element
	return p
}

func (*List_constant_elementContext) IsList_constant_elementContext() {}

func NewList_constant_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_constant_elementContext {
	var p = new(List_constant_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_list_constant_element

	return p
}

func (s *List_constant_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *List_constant_elementContext) Numeric_constant() INumeric_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_constantContext)
}

func (s *List_constant_elementContext) Word_constant_less_the_quotes() IWord_constant_less_the_quotesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWord_constant_less_the_quotesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWord_constant_less_the_quotesContext)
}

func (s *List_constant_elementContext) STRING_CONSTANT() antlr.TerminalNode {
	return s.GetToken(plucidParserSTRING_CONSTANT, 0)
}

func (s *List_constant_elementContext) List_constant() IList_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_constantContext)
}

func (s *List_constant_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_constant_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_constant_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterList_constant_element(s)
	}
}

func (s *List_constant_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitList_constant_element(s)
	}
}

func (p *plucidParser) List_constant_element() (localctx IList_constant_elementContext) {
	this := p
	_ = this

	localctx = NewList_constant_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, plucidParserRULE_list_constant_element)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case plucidParserN_SIGN, plucidParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Numeric_constant()
		}

	case plucidParserSIGN, plucidParserQUOTE, plucidParserBRACKET, plucidParserPERIOD, plucidParserSEPARATOR, plucidParserLETTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Word_constant_less_the_quotes()
		}

	case plucidParserSTRING_CONSTANT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(173)
			p.Match(plucidParserSTRING_CONSTANT)
		}

	case plucidParserT__7, plucidParserT__8, plucidParserT__9:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(174)
			p.List_constant()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = plucidParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) LETTER() antlr.TerminalNode {
	return s.GetToken(plucidParserLETTER, 0)
}

func (s *IdentifierContext) AllALPHANUMERIC() []antlr.TerminalNode {
	return s.GetTokens(plucidParserALPHANUMERIC)
}

func (s *IdentifierContext) ALPHANUMERIC(i int) antlr.TerminalNode {
	return s.GetToken(plucidParserALPHANUMERIC, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *plucidParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, plucidParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(177)
		p.Match(plucidParserLETTER)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(178)
				p.Match(plucidParserALPHANUMERIC)
			}

		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IPrefix_operatorContext is an interface to support dynamic dispatch.
type IPrefix_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefix_operatorContext differentiates from other interfaces.
	IsPrefix_operatorContext()
}

type Prefix_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefix_operatorContext() *Prefix_operatorContext {
	var p = new(Prefix_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_prefix_operator
	return p
}

func (*Prefix_operatorContext) IsPrefix_operatorContext() {}

func NewPrefix_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prefix_operatorContext {
	var p = new(Prefix_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_prefix_operator

	return p
}

func (s *Prefix_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Prefix_operatorContext) P_NUMERIC_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserP_NUMERIC_OPERATOR, 0)
}

func (s *Prefix_operatorContext) P_WORD_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserP_WORD_OPERATOR, 0)
}

func (s *Prefix_operatorContext) P_STRING_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserP_STRING_OPERATOR, 0)
}

func (s *Prefix_operatorContext) P_LIST_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserP_LIST_OPERATOR, 0)
}

func (s *Prefix_operatorContext) P_LUCID_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserP_LUCID_OPERATOR, 0)
}

func (s *Prefix_operatorContext) P_SPECIAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserP_SPECIAL_OPERATOR, 0)
}

func (s *Prefix_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prefix_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prefix_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterPrefix_operator(s)
	}
}

func (s *Prefix_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitPrefix_operator(s)
	}
}

func (p *plucidParser) Prefix_operator() (localctx IPrefix_operatorContext) {
	this := p
	_ = this

	localctx = NewPrefix_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, plucidParserRULE_prefix_operator)
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
		p.SetState(184)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(plucidParserP_NUMERIC_OPERATOR-41))|(1<<(plucidParserP_WORD_OPERATOR-41))|(1<<(plucidParserP_STRING_OPERATOR-41))|(1<<(plucidParserP_LIST_OPERATOR-41))|(1<<(plucidParserP_LUCID_OPERATOR-41))|(1<<(plucidParserP_SPECIAL_OPERATOR-41)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInfix_operatorContext is an interface to support dynamic dispatch.
type IInfix_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfix_operatorContext differentiates from other interfaces.
	IsInfix_operatorContext()
}

type Infix_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfix_operatorContext() *Infix_operatorContext {
	var p = new(Infix_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_infix_operator
	return p
}

func (*Infix_operatorContext) IsInfix_operatorContext() {}

func NewInfix_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Infix_operatorContext {
	var p = new(Infix_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_infix_operator

	return p
}

func (s *Infix_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Infix_operatorContext) I_NUMERIC_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserI_NUMERIC_OPERATOR, 0)
}

func (s *Infix_operatorContext) I_WORD_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserI_WORD_OPERATOR, 0)
}

func (s *Infix_operatorContext) I_STRING_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserI_STRING_OPERATOR, 0)
}

func (s *Infix_operatorContext) I_LIST_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserI_LIST_OPERATOR, 0)
}

func (s *Infix_operatorContext) I_LUCID_OPERATOR() antlr.TerminalNode {
	return s.GetToken(plucidParserI_LUCID_OPERATOR, 0)
}

func (s *Infix_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Infix_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Infix_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterInfix_operator(s)
	}
}

func (s *Infix_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitInfix_operator(s)
	}
}

func (p *plucidParser) Infix_operator() (localctx IInfix_operatorContext) {
	this := p
	_ = this

	localctx = NewInfix_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, plucidParserRULE_infix_operator)
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
		p.SetState(186)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(plucidParserI_NUMERIC_OPERATOR-47))|(1<<(plucidParserI_WORD_OPERATOR-47))|(1<<(plucidParserI_STRING_OPERATOR-47))|(1<<(plucidParserI_LIST_OPERATOR-47))|(1<<(plucidParserI_LUCID_OPERATOR-47)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IList_expressionContext is an interface to support dynamic dispatch.
type IList_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_expressionContext differentiates from other interfaces.
	IsList_expressionContext()
}

type List_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_expressionContext() *List_expressionContext {
	var p = new(List_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_list_expression
	return p
}

func (*List_expressionContext) IsList_expressionContext() {}

func NewList_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expressionContext {
	var p = new(List_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_list_expression

	return p
}

func (s *List_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expressionContext) AllExpressions_list() []IExpressions_listContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressions_listContext)(nil)).Elem())
	var tst = make([]IExpressions_listContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressions_listContext)
		}
	}

	return tst
}

func (s *List_expressionContext) Expressions_list(i int) IExpressions_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressions_listContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressions_listContext)
}

func (s *List_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterList_expression(s)
	}
}

func (s *List_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitList_expression(s)
	}
}

func (p *plucidParser) List_expression() (localctx IList_expressionContext) {
	this := p
	_ = this

	localctx = NewList_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, plucidParserRULE_list_expression)
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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case plucidParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(plucidParserT__11)
		}

	case plucidParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Match(plucidParserT__12)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<plucidParserT__0)|(1<<plucidParserT__1)|(1<<plucidParserT__2)|(1<<plucidParserT__4)|(1<<plucidParserT__5)|(1<<plucidParserT__6)|(1<<plucidParserT__7)|(1<<plucidParserT__8)|(1<<plucidParserT__9)|(1<<plucidParserT__11)|(1<<plucidParserT__12)|(1<<plucidParserT__14)|(1<<plucidParserT__20)|(1<<plucidParserT__23))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(plucidParserN_SIGN-34))|(1<<(plucidParserQUOTE-34))|(1<<(plucidParserSTRING_CONSTANT-34))|(1<<(plucidParserP_NUMERIC_OPERATOR-34))|(1<<(plucidParserP_WORD_OPERATOR-34))|(1<<(plucidParserP_STRING_OPERATOR-34))|(1<<(plucidParserP_LIST_OPERATOR-34))|(1<<(plucidParserP_LUCID_OPERATOR-34))|(1<<(plucidParserP_SPECIAL_OPERATOR-34))|(1<<(plucidParserDIGIT-34))|(1<<(plucidParserLETTER-34)))) != 0) {
			{
				p.SetState(190)
				p.Expressions_list()
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(196)
			p.Match(plucidParserT__13)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressions_listContext is an interface to support dynamic dispatch.
type IExpressions_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressions_listContext differentiates from other interfaces.
	IsExpressions_listContext()
}

type Expressions_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressions_listContext() *Expressions_listContext {
	var p = new(Expressions_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_expressions_list
	return p
}

func (*Expressions_listContext) IsExpressions_listContext() {}

func NewExpressions_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expressions_listContext {
	var p = new(Expressions_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_expressions_list

	return p
}

func (s *Expressions_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expressions_listContext) Expression_item() IExpression_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_itemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_itemContext)
}

func (s *Expressions_listContext) AllExpressions_list() []IExpressions_listContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressions_listContext)(nil)).Elem())
	var tst = make([]IExpressions_listContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressions_listContext)
		}
	}

	return tst
}

func (s *Expressions_listContext) Expressions_list(i int) IExpressions_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressions_listContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressions_listContext)
}

func (s *Expressions_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expressions_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expressions_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterExpressions_list(s)
	}
}

func (s *Expressions_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitExpressions_list(s)
	}
}

func (p *plucidParser) Expressions_list() (localctx IExpressions_listContext) {
	this := p
	_ = this

	localctx = NewExpressions_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, plucidParserRULE_expressions_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Expression_item()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
			p.Expression_item()
		}
		{
			p.SetState(201)
			p.Match(plucidParserT__3)
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(202)
					p.Expressions_list()
				}

			}
			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IExpression_itemContext is an interface to support dynamic dispatch.
type IExpression_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_itemContext differentiates from other interfaces.
	IsExpression_itemContext()
}

type Expression_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_itemContext() *Expression_itemContext {
	var p = new(Expression_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_expression_item
	return p
}

func (*Expression_itemContext) IsExpression_itemContext() {}

func NewExpression_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_itemContext {
	var p = new(Expression_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_expression_item

	return p
}

func (s *Expression_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_itemContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expression_itemContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Expression_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterExpression_item(s)
	}
}

func (s *Expression_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitExpression_item(s)
	}
}

func (p *plucidParser) Expression_item() (localctx IExpression_itemContext) {
	this := p
	_ = this

	localctx = NewExpression_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, plucidParserRULE_expression_item)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.List_expression()
		}

	}

	return localctx
}

// IIf_expressionContext is an interface to support dynamic dispatch.
type IIf_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_expressionContext differentiates from other interfaces.
	IsIf_expressionContext()
}

type If_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_expressionContext() *If_expressionContext {
	var p = new(If_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_if_expression
	return p
}

func (*If_expressionContext) IsIf_expressionContext() {}

func NewIf_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_expressionContext {
	var p = new(If_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_if_expression

	return p
}

func (s *If_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *If_expressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *If_expressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_expressionContext) Endif() IEndifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndifContext)
}

func (s *If_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterIf_expression(s)
	}
}

func (s *If_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitIf_expression(s)
	}
}

func (p *plucidParser) If_expression() (localctx IIf_expressionContext) {
	this := p
	_ = this

	localctx = NewIf_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, plucidParserRULE_if_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(plucidParserT__14)
	}
	{
		p.SetState(215)
		p.expression(0)
	}
	{
		p.SetState(216)
		p.Match(plucidParserT__15)
	}
	{
		p.SetState(217)
		p.expression(0)
	}
	{
		p.SetState(218)
		p.Endif()
	}

	return localctx
}

// IEndifContext is an interface to support dynamic dispatch.
type IEndifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndifContext differentiates from other interfaces.
	IsEndifContext()
}

type EndifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndifContext() *EndifContext {
	var p = new(EndifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_endif
	return p
}

func (*EndifContext) IsEndifContext() {}

func NewEndifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndifContext {
	var p = new(EndifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_endif

	return p
}

func (s *EndifContext) GetParser() antlr.Parser { return s.parser }

func (s *EndifContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EndifContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EndifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterEndif(s)
	}
}

func (s *EndifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitEndif(s)
	}
}

func (p *plucidParser) Endif() (localctx IEndifContext) {
	this := p
	_ = this

	localctx = NewEndifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, plucidParserRULE_endif)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case plucidParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Match(plucidParserT__16)
		}
		{
			p.SetState(221)
			p.expression(0)
		}
		{
			p.SetState(222)
			p.Match(plucidParserT__17)
		}

	case plucidParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(plucidParserT__18)
		}
		{
			p.SetState(225)
			p.expression(0)
		}
		{
			p.SetState(226)
			p.Match(plucidParserT__15)
		}
		{
			p.SetState(227)
			p.expression(0)
		}
		{
			p.SetState(228)
			p.Match(plucidParserT__19)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICase_expressionContext is an interface to support dynamic dispatch.
type ICase_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCase_expressionContext differentiates from other interfaces.
	IsCase_expressionContext()
}

type Case_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_expressionContext() *Case_expressionContext {
	var p = new(Case_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_case_expression
	return p
}

func (*Case_expressionContext) IsCase_expressionContext() {}

func NewCase_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_expressionContext {
	var p = new(Case_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_case_expression

	return p
}

func (s *Case_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Case_expressionContext) Cbody() ICbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICbodyContext)
}

func (s *Case_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterCase_expression(s)
	}
}

func (s *Case_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitCase_expression(s)
	}
}

func (p *plucidParser) Case_expression() (localctx ICase_expressionContext) {
	this := p
	_ = this

	localctx = NewCase_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, plucidParserRULE_case_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(plucidParserT__20)
	}
	{
		p.SetState(233)
		p.expression(0)
	}
	{
		p.SetState(234)
		p.Match(plucidParserT__21)
	}
	{
		p.SetState(235)
		p.Cbody()
	}
	{
		p.SetState(236)
		p.Match(plucidParserT__22)
	}

	return localctx
}

// ICond_expressionContext is an interface to support dynamic dispatch.
type ICond_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCond_expressionContext differentiates from other interfaces.
	IsCond_expressionContext()
}

type Cond_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCond_expressionContext() *Cond_expressionContext {
	var p = new(Cond_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_cond_expression
	return p
}

func (*Cond_expressionContext) IsCond_expressionContext() {}

func NewCond_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cond_expressionContext {
	var p = new(Cond_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_cond_expression

	return p
}

func (s *Cond_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Cond_expressionContext) Cbody() ICbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICbodyContext)
}

func (s *Cond_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cond_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cond_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterCond_expression(s)
	}
}

func (s *Cond_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitCond_expression(s)
	}
}

func (p *plucidParser) Cond_expression() (localctx ICond_expressionContext) {
	this := p
	_ = this

	localctx = NewCond_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, plucidParserRULE_cond_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(plucidParserT__23)
	}
	{
		p.SetState(239)
		p.Cbody()
	}
	{
		p.SetState(240)
		p.Match(plucidParserT__22)
	}

	return localctx
}

// ICbodyContext is an interface to support dynamic dispatch.
type ICbodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCbodyContext differentiates from other interfaces.
	IsCbodyContext()
}

type CbodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCbodyContext() *CbodyContext {
	var p = new(CbodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_cbody
	return p
}

func (*CbodyContext) IsCbodyContext() {}

func NewCbodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CbodyContext {
	var p = new(CbodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_cbody

	return p
}

func (s *CbodyContext) GetParser() antlr.Parser { return s.parser }

func (s *CbodyContext) Defaultcase() IDefaultcaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultcaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultcaseContext)
}

func (s *CbodyContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CbodyContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CbodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CbodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CbodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterCbody(s)
	}
}

func (s *CbodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitCbody(s)
	}
}

func (p *plucidParser) Cbody() (localctx ICbodyContext) {
	this := p
	_ = this

	localctx = NewCbodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, plucidParserRULE_cbody)
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
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<plucidParserT__0)|(1<<plucidParserT__1)|(1<<plucidParserT__2)|(1<<plucidParserT__4)|(1<<plucidParserT__5)|(1<<plucidParserT__6)|(1<<plucidParserT__7)|(1<<plucidParserT__8)|(1<<plucidParserT__9)|(1<<plucidParserT__11)|(1<<plucidParserT__12)|(1<<plucidParserT__14)|(1<<plucidParserT__20)|(1<<plucidParserT__23))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(plucidParserN_SIGN-34))|(1<<(plucidParserQUOTE-34))|(1<<(plucidParserSTRING_CONSTANT-34))|(1<<(plucidParserP_NUMERIC_OPERATOR-34))|(1<<(plucidParserP_WORD_OPERATOR-34))|(1<<(plucidParserP_STRING_OPERATOR-34))|(1<<(plucidParserP_LIST_OPERATOR-34))|(1<<(plucidParserP_LUCID_OPERATOR-34))|(1<<(plucidParserP_SPECIAL_OPERATOR-34))|(1<<(plucidParserDIGIT-34))|(1<<(plucidParserLETTER-34)))) != 0) {
		{
			p.SetState(242)
			p.expression(0)
		}
		{
			p.SetState(243)
			p.Match(plucidParserT__24)
		}
		{
			p.SetState(244)
			p.expression(0)
		}
		{
			p.SetState(245)
			p.Match(plucidParserT__25)
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(252)
		p.Defaultcase()
	}

	return localctx
}

// IDefaultcaseContext is an interface to support dynamic dispatch.
type IDefaultcaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultcaseContext differentiates from other interfaces.
	IsDefaultcaseContext()
}

type DefaultcaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultcaseContext() *DefaultcaseContext {
	var p = new(DefaultcaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_defaultcase
	return p
}

func (*DefaultcaseContext) IsDefaultcaseContext() {}

func NewDefaultcaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultcaseContext {
	var p = new(DefaultcaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_defaultcase

	return p
}

func (s *DefaultcaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultcaseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DefaultcaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultcaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultcaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterDefaultcase(s)
	}
}

func (s *DefaultcaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitDefaultcase(s)
	}
}

func (p *plucidParser) Defaultcase() (localctx IDefaultcaseContext) {
	this := p
	_ = this

	localctx = NewDefaultcaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, plucidParserRULE_defaultcase)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(plucidParserT__26)
	}
	{
		p.SetState(255)
		p.Match(plucidParserT__24)
	}
	{
		p.SetState(256)
		p.expression(0)
	}

	return localctx
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_function_call
	return p
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Function_callContext) Actuals_list() IActuals_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActuals_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActuals_listContext)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *plucidParser) Function_call() (localctx IFunction_callContext) {
	this := p
	_ = this

	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, plucidParserRULE_function_call)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Identifier()
	}
	{
		p.SetState(259)
		p.Match(plucidParserT__27)
	}
	{
		p.SetState(260)
		p.Actuals_list()
	}
	{
		p.SetState(261)
		p.Match(plucidParserT__28)
	}

	return localctx
}

// IActuals_listContext is an interface to support dynamic dispatch.
type IActuals_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActuals_listContext differentiates from other interfaces.
	IsActuals_listContext()
}

type Actuals_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActuals_listContext() *Actuals_listContext {
	var p = new(Actuals_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_actuals_list
	return p
}

func (*Actuals_listContext) IsActuals_listContext() {}

func NewActuals_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Actuals_listContext {
	var p = new(Actuals_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_actuals_list

	return p
}

func (s *Actuals_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Actuals_listContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Actuals_listContext) Actuals_list() IActuals_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActuals_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActuals_listContext)
}

func (s *Actuals_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Actuals_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Actuals_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterActuals_list(s)
	}
}

func (s *Actuals_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitActuals_list(s)
	}
}

func (p *plucidParser) Actuals_list() (localctx IActuals_listContext) {
	this := p
	_ = this

	localctx = NewActuals_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, plucidParserRULE_actuals_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.expression(0)
		}
		{
			p.SetState(265)
			p.Match(plucidParserT__3)
		}
		{
			p.SetState(266)
			p.Actuals_list()
		}

	}

	return localctx
}

// IWhere_clauseContext is an interface to support dynamic dispatch.
type IWhere_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhere_clauseContext differentiates from other interfaces.
	IsWhere_clauseContext()
}

type Where_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhere_clauseContext() *Where_clauseContext {
	var p = new(Where_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_where_clause
	return p
}

func (*Where_clauseContext) IsWhere_clauseContext() {}

func NewWhere_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Where_clauseContext {
	var p = new(Where_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_where_clause

	return p
}

func (s *Where_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Where_clauseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Where_clauseContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *Where_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Where_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Where_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterWhere_clause(s)
	}
}

func (s *Where_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitWhere_clause(s)
	}
}

func (p *plucidParser) Where_clause() (localctx IWhere_clauseContext) {
	this := p
	_ = this

	localctx = NewWhere_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, plucidParserRULE_where_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.expression(0)
	}
	{
		p.SetState(271)
		p.Match(plucidParserT__29)
	}
	{
		p.SetState(272)
		p.Body()
	}
	{
		p.SetState(273)
		p.Match(plucidParserT__22)
	}

	return localctx
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) Declarations_list() IDeclarations_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarations_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarations_listContext)
}

func (s *BodyContext) Definitions_list() IDefinitions_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitions_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitions_listContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitBody(s)
	}
}

func (p *plucidParser) Body() (localctx IBodyContext) {
	this := p
	_ = this

	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, plucidParserRULE_body)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Declarations_list()
	}
	{
		p.SetState(276)
		p.Definitions_list()
	}

	return localctx
}

// IDeclarations_listContext is an interface to support dynamic dispatch.
type IDeclarations_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarations_listContext differentiates from other interfaces.
	IsDeclarations_listContext()
}

type Declarations_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarations_listContext() *Declarations_listContext {
	var p = new(Declarations_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_declarations_list
	return p
}

func (*Declarations_listContext) IsDeclarations_listContext() {}

func NewDeclarations_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declarations_listContext {
	var p = new(Declarations_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_declarations_list

	return p
}

func (s *Declarations_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Declarations_listContext) AllCurrent_declaration() []ICurrent_declarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICurrent_declarationContext)(nil)).Elem())
	var tst = make([]ICurrent_declarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICurrent_declarationContext)
		}
	}

	return tst
}

func (s *Declarations_listContext) Current_declaration(i int) ICurrent_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurrent_declarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICurrent_declarationContext)
}

func (s *Declarations_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declarations_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declarations_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterDeclarations_list(s)
	}
}

func (s *Declarations_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitDeclarations_list(s)
	}
}

func (p *plucidParser) Declarations_list() (localctx IDeclarations_listContext) {
	this := p
	_ = this

	localctx = NewDeclarations_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, plucidParserRULE_declarations_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(278)
				p.Current_declaration()
			}
			{
				p.SetState(279)
				p.Match(plucidParserT__25)
			}

		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// ICurrent_declarationContext is an interface to support dynamic dispatch.
type ICurrent_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurrent_declarationContext differentiates from other interfaces.
	IsCurrent_declarationContext()
}

type Current_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrent_declarationContext() *Current_declarationContext {
	var p = new(Current_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_current_declaration
	return p
}

func (*Current_declarationContext) IsCurrent_declarationContext() {}

func NewCurrent_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Current_declarationContext {
	var p = new(Current_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_current_declaration

	return p
}

func (s *Current_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Current_declarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Current_declarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Current_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Current_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Current_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterCurrent_declaration(s)
	}
}

func (s *Current_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitCurrent_declaration(s)
	}
}

func (p *plucidParser) Current_declaration() (localctx ICurrent_declarationContext) {
	this := p
	_ = this

	localctx = NewCurrent_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, plucidParserRULE_current_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Identifier()
	}
	{
		p.SetState(287)
		p.Match(plucidParserT__30)
	}
	{
		p.SetState(288)
		p.Match(plucidParserT__31)
	}
	{
		p.SetState(289)
		p.expression(0)
	}

	return localctx
}

// IDefinitions_listContext is an interface to support dynamic dispatch.
type IDefinitions_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitions_listContext differentiates from other interfaces.
	IsDefinitions_listContext()
}

type Definitions_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitions_listContext() *Definitions_listContext {
	var p = new(Definitions_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_definitions_list
	return p
}

func (*Definitions_listContext) IsDefinitions_listContext() {}

func NewDefinitions_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Definitions_listContext {
	var p = new(Definitions_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_definitions_list

	return p
}

func (s *Definitions_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Definitions_listContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *Definitions_listContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *Definitions_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Definitions_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Definitions_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterDefinitions_list(s)
	}
}

func (s *Definitions_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitDefinitions_list(s)
	}
}

func (p *plucidParser) Definitions_list() (localctx IDefinitions_listContext) {
	this := p
	_ = this

	localctx = NewDefinitions_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, plucidParserRULE_definitions_list)
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
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == plucidParserLETTER {
		{
			p.SetState(291)
			p.Definition()
		}
		{
			p.SetState(292)
			p.Match(plucidParserT__25)
		}

		p.SetState(298)
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
	p.RuleIndex = plucidParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) Simple_definition() ISimple_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_definitionContext)
}

func (s *DefinitionContext) Function_definition() IFunction_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_definitionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *plucidParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, plucidParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(299)
			p.Simple_definition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.Function_definition()
		}

	}

	return localctx
}

// ISimple_definitionContext is an interface to support dynamic dispatch.
type ISimple_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_definitionContext differentiates from other interfaces.
	IsSimple_definitionContext()
}

type Simple_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_definitionContext() *Simple_definitionContext {
	var p = new(Simple_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_simple_definition
	return p
}

func (*Simple_definitionContext) IsSimple_definitionContext() {}

func NewSimple_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_definitionContext {
	var p = new(Simple_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_simple_definition

	return p
}

func (s *Simple_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_definitionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Simple_definitionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Simple_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterSimple_definition(s)
	}
}

func (s *Simple_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitSimple_definition(s)
	}
}

func (p *plucidParser) Simple_definition() (localctx ISimple_definitionContext) {
	this := p
	_ = this

	localctx = NewSimple_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, plucidParserRULE_simple_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Identifier()
	}
	{
		p.SetState(304)
		p.Match(plucidParserT__32)
	}
	{
		p.SetState(305)
		p.expression(0)
	}

	return localctx
}

// IFunction_definitionContext is an interface to support dynamic dispatch.
type IFunction_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_definitionContext differentiates from other interfaces.
	IsFunction_definitionContext()
}

type Function_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_definitionContext() *Function_definitionContext {
	var p = new(Function_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_function_definition
	return p
}

func (*Function_definitionContext) IsFunction_definitionContext() {}

func NewFunction_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_definitionContext {
	var p = new(Function_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_function_definition

	return p
}

func (s *Function_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_definitionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Function_definitionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Function_definitionContext) Formals_list() IFormals_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormals_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormals_listContext)
}

func (s *Function_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterFunction_definition(s)
	}
}

func (s *Function_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitFunction_definition(s)
	}
}

func (p *plucidParser) Function_definition() (localctx IFunction_definitionContext) {
	this := p
	_ = this

	localctx = NewFunction_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, plucidParserRULE_function_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Identifier()
	}

	{
		p.SetState(308)
		p.Formals_list()
	}

	{
		p.SetState(309)
		p.Match(plucidParserT__32)
	}
	{
		p.SetState(310)
		p.expression(0)
	}

	return localctx
}

// IFormals_listContext is an interface to support dynamic dispatch.
type IFormals_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormals_listContext differentiates from other interfaces.
	IsFormals_listContext()
}

type Formals_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormals_listContext() *Formals_listContext {
	var p = new(Formals_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = plucidParserRULE_formals_list
	return p
}

func (*Formals_listContext) IsFormals_listContext() {}

func NewFormals_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Formals_listContext {
	var p = new(Formals_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = plucidParserRULE_formals_list

	return p
}

func (s *Formals_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Formals_listContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Formals_listContext) Formals_list() IFormals_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormals_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormals_listContext)
}

func (s *Formals_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Formals_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Formals_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.EnterFormals_list(s)
	}
}

func (s *Formals_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(plucidListener); ok {
		listenerT.ExitFormals_list(s)
	}
}

func (p *plucidParser) Formals_list() (localctx IFormals_listContext) {
	this := p
	_ = this

	localctx = NewFormals_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, plucidParserRULE_formals_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(312)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.Identifier()
		}
		{
			p.SetState(314)
			p.Match(plucidParserT__3)
		}
		{
			p.SetState(315)
			p.Formals_list()
		}

	}

	return localctx
}

func (p *plucidParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *plucidParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
