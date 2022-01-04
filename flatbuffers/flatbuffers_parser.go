// Code generated from FlatBuffers.g4 by ANTLR 4.9.3. DO NOT EDIT.

package flatbuffers // FlatBuffers
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 309,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 3, 2,
	7, 2, 68, 10, 2, 12, 2, 14, 2, 71, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 83, 10, 2, 12, 2, 14, 2, 86, 11, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 96, 10, 4, 12, 4,
	14, 4, 99, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 7, 6, 112, 10, 6, 12, 6, 14, 6, 115, 11, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 5, 7, 123, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 5, 10, 146, 10, 10, 3, 10, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 6, 11, 155, 10, 11, 13, 11, 14, 11, 156, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 5, 13, 174, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5,
	13, 180, 10, 13, 3, 14, 3, 14, 3, 14, 5, 14, 185, 10, 14, 3, 15, 3, 15,
	3, 15, 7, 15, 190, 10, 15, 12, 15, 14, 15, 193, 11, 15, 3, 15, 5, 15, 196,
	10, 15, 3, 16, 3, 16, 3, 16, 5, 16, 201, 10, 16, 3, 16, 3, 16, 5, 16, 205,
	10, 16, 3, 17, 3, 17, 3, 17, 7, 17, 210, 10, 17, 12, 17, 14, 17, 213, 11,
	17, 3, 17, 5, 17, 216, 10, 17, 3, 18, 3, 18, 3, 18, 5, 18, 221, 10, 18,
	3, 19, 3, 19, 3, 19, 7, 19, 226, 10, 19, 12, 19, 14, 19, 229, 11, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 5, 20, 235, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21,
	5, 21, 241, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 7, 24, 254, 10, 24, 12, 24, 14, 24, 257, 11, 24,
	3, 24, 5, 24, 260, 10, 24, 3, 25, 3, 25, 5, 25, 264, 10, 25, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 272, 10, 26, 3, 27, 3, 27, 3, 27,
	7, 27, 277, 10, 27, 12, 27, 14, 27, 280, 11, 27, 3, 27, 5, 27, 283, 10,
	27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 30, 7, 30, 296, 10, 30, 12, 30, 14, 30, 299, 11, 30, 3, 31, 3, 31, 3,
	32, 3, 32, 5, 32, 305, 10, 32, 3, 33, 3, 33, 3, 33, 2, 2, 34, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 2, 6, 3, 2, 7, 8, 3, 2, 12,
	13, 4, 2, 30, 30, 32, 32, 3, 2, 3, 14, 2, 317, 2, 69, 3, 2, 2, 2, 4, 87,
	3, 2, 2, 2, 6, 91, 3, 2, 2, 2, 8, 102, 3, 2, 2, 2, 10, 106, 3, 2, 2, 2,
	12, 118, 3, 2, 2, 2, 14, 129, 3, 2, 2, 2, 16, 136, 3, 2, 2, 2, 18, 140,
	3, 2, 2, 2, 20, 150, 3, 2, 2, 2, 22, 160, 3, 2, 2, 2, 24, 179, 3, 2, 2,
	2, 26, 181, 3, 2, 2, 2, 28, 186, 3, 2, 2, 2, 30, 197, 3, 2, 2, 2, 32, 206,
	3, 2, 2, 2, 34, 217, 3, 2, 2, 2, 36, 222, 3, 2, 2, 2, 38, 234, 3, 2, 2,
	2, 40, 240, 3, 2, 2, 2, 42, 242, 3, 2, 2, 2, 44, 246, 3, 2, 2, 2, 46, 250,
	3, 2, 2, 2, 48, 263, 3, 2, 2, 2, 50, 271, 3, 2, 2, 2, 52, 273, 3, 2, 2,
	2, 54, 284, 3, 2, 2, 2, 56, 288, 3, 2, 2, 2, 58, 292, 3, 2, 2, 2, 60, 300,
	3, 2, 2, 2, 62, 304, 3, 2, 2, 2, 64, 306, 3, 2, 2, 2, 66, 68, 5, 4, 3,
	2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70,
	3, 2, 2, 2, 70, 84, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 83, 5, 6, 4, 2,
	73, 83, 5, 10, 6, 2, 74, 83, 5, 12, 7, 2, 75, 83, 5, 14, 8, 2, 76, 83,
	5, 16, 9, 2, 77, 83, 5, 54, 28, 2, 78, 83, 5, 56, 29, 2, 79, 83, 5, 8,
	5, 2, 80, 83, 5, 20, 11, 2, 81, 83, 5, 42, 22, 2, 82, 72, 3, 2, 2, 2, 82,
	73, 3, 2, 2, 2, 82, 74, 3, 2, 2, 2, 82, 75, 3, 2, 2, 2, 82, 76, 3, 2, 2,
	2, 82, 77, 3, 2, 2, 2, 82, 78, 3, 2, 2, 2, 82, 79, 3, 2, 2, 2, 82, 80,
	3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2,
	84, 85, 3, 2, 2, 2, 85, 3, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 87, 88, 9, 2,
	2, 2, 88, 89, 7, 28, 2, 2, 89, 90, 7, 15, 2, 2, 90, 5, 3, 2, 2, 2, 91,
	92, 7, 9, 2, 2, 92, 97, 5, 62, 32, 2, 93, 94, 7, 23, 2, 2, 94, 96, 5, 62,
	32, 2, 95, 93, 3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97,
	98, 3, 2, 2, 2, 98, 100, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 101, 7, 15,
	2, 2, 101, 7, 3, 2, 2, 2, 102, 103, 7, 3, 2, 2, 103, 104, 7, 28, 2, 2,
	104, 105, 7, 15, 2, 2, 105, 9, 3, 2, 2, 2, 106, 107, 9, 3, 2, 2, 107, 108,
	5, 62, 32, 2, 108, 109, 5, 38, 20, 2, 109, 113, 7, 21, 2, 2, 110, 112,
	5, 18, 10, 2, 111, 110, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3,
	2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2, 115, 113, 3, 2, 2,
	2, 116, 117, 7, 22, 2, 2, 117, 11, 3, 2, 2, 2, 118, 119, 7, 4, 2, 2, 119,
	122, 5, 62, 32, 2, 120, 121, 7, 25, 2, 2, 121, 123, 5, 24, 13, 2, 122,
	120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125,
	5, 38, 20, 2, 125, 126, 7, 21, 2, 2, 126, 127, 5, 28, 15, 2, 127, 128,
	7, 22, 2, 2, 128, 13, 3, 2, 2, 2, 129, 130, 7, 14, 2, 2, 130, 131, 5, 62,
	32, 2, 131, 132, 5, 38, 20, 2, 132, 133, 7, 21, 2, 2, 133, 134, 5, 32,
	17, 2, 134, 135, 7, 22, 2, 2, 135, 15, 3, 2, 2, 2, 136, 137, 7, 10, 2,
	2, 137, 138, 5, 62, 32, 2, 138, 139, 7, 15, 2, 2, 139, 17, 3, 2, 2, 2,
	140, 141, 5, 62, 32, 2, 141, 142, 7, 25, 2, 2, 142, 145, 5, 24, 13, 2,
	143, 144, 7, 16, 2, 2, 144, 146, 5, 40, 21, 2, 145, 143, 3, 2, 2, 2, 145,
	146, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 148, 5, 38, 20, 2, 148, 149,
	7, 15, 2, 2, 149, 19, 3, 2, 2, 2, 150, 151, 7, 11, 2, 2, 151, 152, 5, 62,
	32, 2, 152, 154, 7, 21, 2, 2, 153, 155, 5, 22, 12, 2, 154, 153, 3, 2, 2,
	2, 155, 156, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157,
	158, 3, 2, 2, 2, 158, 159, 7, 22, 2, 2, 159, 21, 3, 2, 2, 2, 160, 161,
	5, 62, 32, 2, 161, 162, 7, 17, 2, 2, 162, 163, 5, 62, 32, 2, 163, 164,
	7, 18, 2, 2, 164, 165, 7, 25, 2, 2, 165, 166, 5, 62, 32, 2, 166, 167, 5,
	38, 20, 2, 167, 168, 7, 15, 2, 2, 168, 23, 3, 2, 2, 2, 169, 170, 7, 19,
	2, 2, 170, 173, 5, 24, 13, 2, 171, 172, 7, 25, 2, 2, 172, 174, 5, 60, 31,
	2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175,
	176, 7, 20, 2, 2, 176, 180, 3, 2, 2, 2, 177, 180, 7, 29, 2, 2, 178, 180,
	5, 58, 30, 2, 179, 169, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 178, 3,
	2, 2, 2, 180, 25, 3, 2, 2, 2, 181, 184, 5, 58, 30, 2, 182, 183, 7, 16,
	2, 2, 183, 185, 5, 60, 31, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2,
	2, 185, 27, 3, 2, 2, 2, 186, 191, 5, 26, 14, 2, 187, 188, 7, 24, 2, 2,
	188, 190, 5, 26, 14, 2, 189, 187, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191,
	189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191,
	3, 2, 2, 2, 194, 196, 7, 24, 2, 2, 195, 194, 3, 2, 2, 2, 195, 196, 3, 2,
	2, 2, 196, 29, 3, 2, 2, 2, 197, 200, 5, 58, 30, 2, 198, 199, 7, 25, 2,
	2, 199, 201, 5, 58, 30, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2,
	201, 204, 3, 2, 2, 2, 202, 203, 7, 16, 2, 2, 203, 205, 5, 60, 31, 2, 204,
	202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 31, 3, 2, 2, 2, 206, 211, 5,
	30, 16, 2, 207, 208, 7, 24, 2, 2, 208, 210, 5, 30, 16, 2, 209, 207, 3,
	2, 2, 2, 210, 213, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2,
	2, 212, 215, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 214, 216, 7, 24, 2, 2, 215,
	214, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 33, 3, 2, 2, 2, 217, 220, 5,
	62, 32, 2, 218, 219, 7, 25, 2, 2, 219, 221, 5, 48, 25, 2, 220, 218, 3,
	2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 35, 3, 2, 2, 2, 222, 227, 5, 34, 18,
	2, 223, 224, 7, 24, 2, 2, 224, 226, 5, 34, 18, 2, 225, 223, 3, 2, 2, 2,
	226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228,
	37, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 231, 7, 17, 2, 2, 231, 232,
	5, 36, 19, 2, 232, 233, 7, 18, 2, 2, 233, 235, 3, 2, 2, 2, 234, 230, 3,
	2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 39, 3, 2, 2, 2, 236, 241, 7, 30, 2,
	2, 237, 241, 7, 32, 2, 2, 238, 241, 7, 33, 2, 2, 239, 241, 5, 62, 32, 2,
	240, 236, 3, 2, 2, 2, 240, 237, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240,
	239, 3, 2, 2, 2, 241, 41, 3, 2, 2, 2, 242, 243, 7, 21, 2, 2, 243, 244,
	5, 46, 24, 2, 244, 245, 7, 22, 2, 2, 245, 43, 3, 2, 2, 2, 246, 247, 5,
	62, 32, 2, 247, 248, 7, 25, 2, 2, 248, 249, 5, 50, 26, 2, 249, 45, 3, 2,
	2, 2, 250, 255, 5, 44, 23, 2, 251, 252, 7, 24, 2, 2, 252, 254, 5, 44, 23,
	2, 253, 251, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255,
	256, 3, 2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 258, 260,
	7, 24, 2, 2, 259, 258, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 47, 3, 2,
	2, 2, 261, 264, 5, 40, 21, 2, 262, 264, 7, 28, 2, 2, 263, 261, 3, 2, 2,
	2, 263, 262, 3, 2, 2, 2, 264, 49, 3, 2, 2, 2, 265, 272, 5, 48, 25, 2, 266,
	272, 5, 42, 22, 2, 267, 268, 7, 19, 2, 2, 268, 269, 5, 52, 27, 2, 269,
	270, 7, 20, 2, 2, 270, 272, 3, 2, 2, 2, 271, 265, 3, 2, 2, 2, 271, 266,
	3, 2, 2, 2, 271, 267, 3, 2, 2, 2, 272, 51, 3, 2, 2, 2, 273, 278, 5, 50,
	26, 2, 274, 275, 7, 24, 2, 2, 275, 277, 5, 50, 26, 2, 276, 274, 3, 2, 2,
	2, 277, 280, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279,
	282, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 281, 283, 7, 24, 2, 2, 282, 281,
	3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 53, 3, 2, 2, 2, 284, 285, 7, 5,
	2, 2, 285, 286, 7, 28, 2, 2, 286, 287, 7, 15, 2, 2, 287, 55, 3, 2, 2, 2,
	288, 289, 7, 6, 2, 2, 289, 290, 7, 28, 2, 2, 290, 291, 7, 15, 2, 2, 291,
	57, 3, 2, 2, 2, 292, 297, 5, 62, 32, 2, 293, 294, 7, 23, 2, 2, 294, 296,
	5, 62, 32, 2, 295, 293, 3, 2, 2, 2, 296, 299, 3, 2, 2, 2, 297, 295, 3,
	2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 59, 3, 2, 2, 2, 299, 297, 3, 2, 2,
	2, 300, 301, 9, 4, 2, 2, 301, 61, 3, 2, 2, 2, 302, 305, 7, 31, 2, 2, 303,
	305, 5, 64, 33, 2, 304, 302, 3, 2, 2, 2, 304, 303, 3, 2, 2, 2, 305, 63,
	3, 2, 2, 2, 306, 307, 9, 5, 2, 2, 307, 65, 3, 2, 2, 2, 31, 69, 82, 84,
	97, 113, 122, 145, 156, 173, 179, 184, 191, 195, 200, 204, 211, 215, 220,
	227, 234, 240, 255, 259, 263, 271, 278, 282, 297, 304,
}
var literalNames = []string{
	"", "'attribute'", "'enum'", "'file_extension'", "'file_identifier'", "'include'",
	"'native_include'", "'namespace'", "'root_type'", "'rpc_service'", "'struct'",
	"'table'", "'union'", "';'", "'='", "'('", "')'", "'['", "']'", "'{'",
	"'}'", "'.'", "','", "':'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "ATTRIBUTE", "ENUM", "FILE_EXTENSION", "FILE_IDENTIFIER", "INCLUDE",
	"NATIVE_INCLUDE", "NAMESPACE", "ROOT_TYPE", "RPC_SERVICE", "STRUCT", "TABLE",
	"UNION", "SEMI", "EQ", "LP", "RP", "LB", "RB", "LC", "RC", "DOT", "COMMA",
	"COLON", "PLUS", "MINUS", "STRING_CONSTANT", "BASE_TYPE_NAME", "INTEGER_CONSTANT",
	"IDENT", "HEX_INTEGER_CONSTANT", "FLOAT_CONSTANT", "BLOCK_COMMENT", "COMMENT",
	"WS",
}

var ruleNames = []string{
	"schema", "include_", "namespace_decl", "attribute_decl", "type_decl",
	"enum_decl", "union_decl", "root_decl", "field_decl", "rpc_decl", "rpc_method",
	"type_", "enumval_decl", "commasep_enumval_decl", "unionval_with_opt_alias",
	"commasep_unionval_with_opt_alias", "ident_with_opt_single_value", "commasep_ident_with_opt_single_value",
	"metadata", "scalar", "object_", "ident_with_value", "commasep_ident_with_value",
	"single_value", "value", "commasep_value", "file_extension_decl", "file_identifier_decl",
	"ns_ident", "integer_const", "identifier", "keywords",
}

type FlatBuffersParser struct {
	*antlr.BaseParser
}

// NewFlatBuffersParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *FlatBuffersParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewFlatBuffersParser(input antlr.TokenStream) *FlatBuffersParser {
	this := new(FlatBuffersParser)
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
	this.GrammarFileName = "FlatBuffers.g4"

	return this
}

// FlatBuffersParser tokens.
const (
	FlatBuffersParserEOF                  = antlr.TokenEOF
	FlatBuffersParserATTRIBUTE            = 1
	FlatBuffersParserENUM                 = 2
	FlatBuffersParserFILE_EXTENSION       = 3
	FlatBuffersParserFILE_IDENTIFIER      = 4
	FlatBuffersParserINCLUDE              = 5
	FlatBuffersParserNATIVE_INCLUDE       = 6
	FlatBuffersParserNAMESPACE            = 7
	FlatBuffersParserROOT_TYPE            = 8
	FlatBuffersParserRPC_SERVICE          = 9
	FlatBuffersParserSTRUCT               = 10
	FlatBuffersParserTABLE                = 11
	FlatBuffersParserUNION                = 12
	FlatBuffersParserSEMI                 = 13
	FlatBuffersParserEQ                   = 14
	FlatBuffersParserLP                   = 15
	FlatBuffersParserRP                   = 16
	FlatBuffersParserLB                   = 17
	FlatBuffersParserRB                   = 18
	FlatBuffersParserLC                   = 19
	FlatBuffersParserRC                   = 20
	FlatBuffersParserDOT                  = 21
	FlatBuffersParserCOMMA                = 22
	FlatBuffersParserCOLON                = 23
	FlatBuffersParserPLUS                 = 24
	FlatBuffersParserMINUS                = 25
	FlatBuffersParserSTRING_CONSTANT      = 26
	FlatBuffersParserBASE_TYPE_NAME       = 27
	FlatBuffersParserINTEGER_CONSTANT     = 28
	FlatBuffersParserIDENT                = 29
	FlatBuffersParserHEX_INTEGER_CONSTANT = 30
	FlatBuffersParserFLOAT_CONSTANT       = 31
	FlatBuffersParserBLOCK_COMMENT        = 32
	FlatBuffersParserCOMMENT              = 33
	FlatBuffersParserWS                   = 34
)

// FlatBuffersParser rules.
const (
	FlatBuffersParserRULE_schema                               = 0
	FlatBuffersParserRULE_include_                             = 1
	FlatBuffersParserRULE_namespace_decl                       = 2
	FlatBuffersParserRULE_attribute_decl                       = 3
	FlatBuffersParserRULE_type_decl                            = 4
	FlatBuffersParserRULE_enum_decl                            = 5
	FlatBuffersParserRULE_union_decl                           = 6
	FlatBuffersParserRULE_root_decl                            = 7
	FlatBuffersParserRULE_field_decl                           = 8
	FlatBuffersParserRULE_rpc_decl                             = 9
	FlatBuffersParserRULE_rpc_method                           = 10
	FlatBuffersParserRULE_type_                                = 11
	FlatBuffersParserRULE_enumval_decl                         = 12
	FlatBuffersParserRULE_commasep_enumval_decl                = 13
	FlatBuffersParserRULE_unionval_with_opt_alias              = 14
	FlatBuffersParserRULE_commasep_unionval_with_opt_alias     = 15
	FlatBuffersParserRULE_ident_with_opt_single_value          = 16
	FlatBuffersParserRULE_commasep_ident_with_opt_single_value = 17
	FlatBuffersParserRULE_metadata                             = 18
	FlatBuffersParserRULE_scalar                               = 19
	FlatBuffersParserRULE_object_                              = 20
	FlatBuffersParserRULE_ident_with_value                     = 21
	FlatBuffersParserRULE_commasep_ident_with_value            = 22
	FlatBuffersParserRULE_single_value                         = 23
	FlatBuffersParserRULE_value                                = 24
	FlatBuffersParserRULE_commasep_value                       = 25
	FlatBuffersParserRULE_file_extension_decl                  = 26
	FlatBuffersParserRULE_file_identifier_decl                 = 27
	FlatBuffersParserRULE_ns_ident                             = 28
	FlatBuffersParserRULE_integer_const                        = 29
	FlatBuffersParserRULE_identifier                           = 30
	FlatBuffersParserRULE_keywords                             = 31
)

// ISchemaContext is an interface to support dynamic dispatch.
type ISchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaContext differentiates from other interfaces.
	IsSchemaContext()
}

type SchemaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaContext() *SchemaContext {
	var p = new(SchemaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_schema
	return p
}

func (*SchemaContext) IsSchemaContext() {}

func NewSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaContext {
	var p = new(SchemaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_schema

	return p
}

func (s *SchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaContext) AllInclude_() []IInclude_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInclude_Context)(nil)).Elem())
	var tst = make([]IInclude_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInclude_Context)
		}
	}

	return tst
}

func (s *SchemaContext) Include_(i int) IInclude_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInclude_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInclude_Context)
}

func (s *SchemaContext) AllNamespace_decl() []INamespace_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INamespace_declContext)(nil)).Elem())
	var tst = make([]INamespace_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INamespace_declContext)
		}
	}

	return tst
}

func (s *SchemaContext) Namespace_decl(i int) INamespace_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespace_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INamespace_declContext)
}

func (s *SchemaContext) AllType_decl() []IType_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_declContext)(nil)).Elem())
	var tst = make([]IType_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_declContext)
		}
	}

	return tst
}

func (s *SchemaContext) Type_decl(i int) IType_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_declContext)
}

func (s *SchemaContext) AllEnum_decl() []IEnum_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnum_declContext)(nil)).Elem())
	var tst = make([]IEnum_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnum_declContext)
		}
	}

	return tst
}

func (s *SchemaContext) Enum_decl(i int) IEnum_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnum_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnum_declContext)
}

func (s *SchemaContext) AllUnion_decl() []IUnion_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnion_declContext)(nil)).Elem())
	var tst = make([]IUnion_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnion_declContext)
		}
	}

	return tst
}

func (s *SchemaContext) Union_decl(i int) IUnion_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnion_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnion_declContext)
}

func (s *SchemaContext) AllRoot_decl() []IRoot_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRoot_declContext)(nil)).Elem())
	var tst = make([]IRoot_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRoot_declContext)
		}
	}

	return tst
}

func (s *SchemaContext) Root_decl(i int) IRoot_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRoot_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRoot_declContext)
}

func (s *SchemaContext) AllFile_extension_decl() []IFile_extension_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFile_extension_declContext)(nil)).Elem())
	var tst = make([]IFile_extension_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFile_extension_declContext)
		}
	}

	return tst
}

func (s *SchemaContext) File_extension_decl(i int) IFile_extension_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFile_extension_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFile_extension_declContext)
}

func (s *SchemaContext) AllFile_identifier_decl() []IFile_identifier_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFile_identifier_declContext)(nil)).Elem())
	var tst = make([]IFile_identifier_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFile_identifier_declContext)
		}
	}

	return tst
}

func (s *SchemaContext) File_identifier_decl(i int) IFile_identifier_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFile_identifier_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFile_identifier_declContext)
}

func (s *SchemaContext) AllAttribute_decl() []IAttribute_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttribute_declContext)(nil)).Elem())
	var tst = make([]IAttribute_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttribute_declContext)
		}
	}

	return tst
}

func (s *SchemaContext) Attribute_decl(i int) IAttribute_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribute_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttribute_declContext)
}

func (s *SchemaContext) AllRpc_decl() []IRpc_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRpc_declContext)(nil)).Elem())
	var tst = make([]IRpc_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRpc_declContext)
		}
	}

	return tst
}

func (s *SchemaContext) Rpc_decl(i int) IRpc_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpc_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRpc_declContext)
}

func (s *SchemaContext) AllObject_() []IObject_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObject_Context)(nil)).Elem())
	var tst = make([]IObject_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObject_Context)
		}
	}

	return tst
}

func (s *SchemaContext) Object_(i int) IObject_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObject_Context)
}

func (s *SchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterSchema(s)
	}
}

func (s *SchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitSchema(s)
	}
}

func (p *FlatBuffersParser) Schema() (localctx ISchemaContext) {
	this := p
	_ = this

	localctx = NewSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FlatBuffersParserRULE_schema)
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FlatBuffersParserINCLUDE || _la == FlatBuffersParserNATIVE_INCLUDE {
		{
			p.SetState(64)
			p.Include_()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FlatBuffersParserATTRIBUTE)|(1<<FlatBuffersParserENUM)|(1<<FlatBuffersParserFILE_EXTENSION)|(1<<FlatBuffersParserFILE_IDENTIFIER)|(1<<FlatBuffersParserNAMESPACE)|(1<<FlatBuffersParserROOT_TYPE)|(1<<FlatBuffersParserRPC_SERVICE)|(1<<FlatBuffersParserSTRUCT)|(1<<FlatBuffersParserTABLE)|(1<<FlatBuffersParserUNION)|(1<<FlatBuffersParserLC))) != 0 {
		p.SetState(80)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case FlatBuffersParserNAMESPACE:
			{
				p.SetState(70)
				p.Namespace_decl()
			}

		case FlatBuffersParserSTRUCT, FlatBuffersParserTABLE:
			{
				p.SetState(71)
				p.Type_decl()
			}

		case FlatBuffersParserENUM:
			{
				p.SetState(72)
				p.Enum_decl()
			}

		case FlatBuffersParserUNION:
			{
				p.SetState(73)
				p.Union_decl()
			}

		case FlatBuffersParserROOT_TYPE:
			{
				p.SetState(74)
				p.Root_decl()
			}

		case FlatBuffersParserFILE_EXTENSION:
			{
				p.SetState(75)
				p.File_extension_decl()
			}

		case FlatBuffersParserFILE_IDENTIFIER:
			{
				p.SetState(76)
				p.File_identifier_decl()
			}

		case FlatBuffersParserATTRIBUTE:
			{
				p.SetState(77)
				p.Attribute_decl()
			}

		case FlatBuffersParserRPC_SERVICE:
			{
				p.SetState(78)
				p.Rpc_decl()
			}

		case FlatBuffersParserLC:
			{
				p.SetState(79)
				p.Object_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInclude_Context is an interface to support dynamic dispatch.
type IInclude_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInclude_Context differentiates from other interfaces.
	IsInclude_Context()
}

type Include_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInclude_Context() *Include_Context {
	var p = new(Include_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_include_
	return p
}

func (*Include_Context) IsInclude_Context() {}

func NewInclude_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Include_Context {
	var p = new(Include_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_include_

	return p
}

func (s *Include_Context) GetParser() antlr.Parser { return s.parser }

func (s *Include_Context) STRING_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSTRING_CONSTANT, 0)
}

func (s *Include_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSEMI, 0)
}

func (s *Include_Context) INCLUDE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserINCLUDE, 0)
}

func (s *Include_Context) NATIVE_INCLUDE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserNATIVE_INCLUDE, 0)
}

func (s *Include_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Include_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Include_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterInclude_(s)
	}
}

func (s *Include_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitInclude_(s)
	}
}

func (p *FlatBuffersParser) Include_() (localctx IInclude_Context) {
	this := p
	_ = this

	localctx = NewInclude_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FlatBuffersParserRULE_include_)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == FlatBuffersParserINCLUDE || _la == FlatBuffersParserNATIVE_INCLUDE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(86)
		p.Match(FlatBuffersParserSTRING_CONSTANT)
	}
	{
		p.SetState(87)
		p.Match(FlatBuffersParserSEMI)
	}

	return localctx
}

// INamespace_declContext is an interface to support dynamic dispatch.
type INamespace_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespace_declContext differentiates from other interfaces.
	IsNamespace_declContext()
}

type Namespace_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespace_declContext() *Namespace_declContext {
	var p = new(Namespace_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_namespace_decl
	return p
}

func (*Namespace_declContext) IsNamespace_declContext() {}

func NewNamespace_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Namespace_declContext {
	var p = new(Namespace_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_namespace_decl

	return p
}

func (s *Namespace_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Namespace_declContext) NAMESPACE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserNAMESPACE, 0)
}

func (s *Namespace_declContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *Namespace_declContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Namespace_declContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSEMI, 0)
}

func (s *Namespace_declContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(FlatBuffersParserDOT)
}

func (s *Namespace_declContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserDOT, i)
}

func (s *Namespace_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Namespace_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Namespace_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterNamespace_decl(s)
	}
}

func (s *Namespace_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitNamespace_decl(s)
	}
}

func (p *FlatBuffersParser) Namespace_decl() (localctx INamespace_declContext) {
	this := p
	_ = this

	localctx = NewNamespace_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FlatBuffersParserRULE_namespace_decl)
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
		p.SetState(89)
		p.Match(FlatBuffersParserNAMESPACE)
	}
	{
		p.SetState(90)
		p.Identifier()
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FlatBuffersParserDOT {
		{
			p.SetState(91)
			p.Match(FlatBuffersParserDOT)
		}
		{
			p.SetState(92)
			p.Identifier()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.Match(FlatBuffersParserSEMI)
	}

	return localctx
}

// IAttribute_declContext is an interface to support dynamic dispatch.
type IAttribute_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribute_declContext differentiates from other interfaces.
	IsAttribute_declContext()
}

type Attribute_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribute_declContext() *Attribute_declContext {
	var p = new(Attribute_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_attribute_decl
	return p
}

func (*Attribute_declContext) IsAttribute_declContext() {}

func NewAttribute_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Attribute_declContext {
	var p = new(Attribute_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_attribute_decl

	return p
}

func (s *Attribute_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Attribute_declContext) ATTRIBUTE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserATTRIBUTE, 0)
}

func (s *Attribute_declContext) STRING_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSTRING_CONSTANT, 0)
}

func (s *Attribute_declContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSEMI, 0)
}

func (s *Attribute_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Attribute_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Attribute_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterAttribute_decl(s)
	}
}

func (s *Attribute_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitAttribute_decl(s)
	}
}

func (p *FlatBuffersParser) Attribute_decl() (localctx IAttribute_declContext) {
	this := p
	_ = this

	localctx = NewAttribute_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FlatBuffersParserRULE_attribute_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(FlatBuffersParserATTRIBUTE)
	}
	{
		p.SetState(101)
		p.Match(FlatBuffersParserSTRING_CONSTANT)
	}
	{
		p.SetState(102)
		p.Match(FlatBuffersParserSEMI)
	}

	return localctx
}

// IType_declContext is an interface to support dynamic dispatch.
type IType_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_declContext differentiates from other interfaces.
	IsType_declContext()
}

type Type_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_declContext() *Type_declContext {
	var p = new(Type_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_type_decl
	return p
}

func (*Type_declContext) IsType_declContext() {}

func NewType_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_declContext {
	var p = new(Type_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_type_decl

	return p
}

func (s *Type_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_declContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Type_declContext) Metadata() IMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataContext)
}

func (s *Type_declContext) LC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserLC, 0)
}

func (s *Type_declContext) RC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRC, 0)
}

func (s *Type_declContext) TABLE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserTABLE, 0)
}

func (s *Type_declContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSTRUCT, 0)
}

func (s *Type_declContext) AllField_decl() []IField_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IField_declContext)(nil)).Elem())
	var tst = make([]IField_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IField_declContext)
		}
	}

	return tst
}

func (s *Type_declContext) Field_decl(i int) IField_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IField_declContext)
}

func (s *Type_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterType_decl(s)
	}
}

func (s *Type_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitType_decl(s)
	}
}

func (p *FlatBuffersParser) Type_decl() (localctx IType_declContext) {
	this := p
	_ = this

	localctx = NewType_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FlatBuffersParserRULE_type_decl)
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
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FlatBuffersParserSTRUCT || _la == FlatBuffersParserTABLE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(105)
		p.Identifier()
	}
	{
		p.SetState(106)
		p.Metadata()
	}
	{
		p.SetState(107)
		p.Match(FlatBuffersParserLC)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FlatBuffersParserATTRIBUTE)|(1<<FlatBuffersParserENUM)|(1<<FlatBuffersParserFILE_EXTENSION)|(1<<FlatBuffersParserFILE_IDENTIFIER)|(1<<FlatBuffersParserINCLUDE)|(1<<FlatBuffersParserNATIVE_INCLUDE)|(1<<FlatBuffersParserNAMESPACE)|(1<<FlatBuffersParserROOT_TYPE)|(1<<FlatBuffersParserRPC_SERVICE)|(1<<FlatBuffersParserSTRUCT)|(1<<FlatBuffersParserTABLE)|(1<<FlatBuffersParserUNION)|(1<<FlatBuffersParserIDENT))) != 0 {
		{
			p.SetState(108)
			p.Field_decl()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(114)
		p.Match(FlatBuffersParserRC)
	}

	return localctx
}

// IEnum_declContext is an interface to support dynamic dispatch.
type IEnum_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnum_declContext differentiates from other interfaces.
	IsEnum_declContext()
}

type Enum_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnum_declContext() *Enum_declContext {
	var p = new(Enum_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_enum_decl
	return p
}

func (*Enum_declContext) IsEnum_declContext() {}

func NewEnum_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_declContext {
	var p = new(Enum_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_enum_decl

	return p
}

func (s *Enum_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_declContext) ENUM() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserENUM, 0)
}

func (s *Enum_declContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Enum_declContext) Metadata() IMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataContext)
}

func (s *Enum_declContext) LC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserLC, 0)
}

func (s *Enum_declContext) Commasep_enumval_decl() ICommasep_enumval_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommasep_enumval_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommasep_enumval_declContext)
}

func (s *Enum_declContext) RC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRC, 0)
}

func (s *Enum_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOLON, 0)
}

func (s *Enum_declContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Enum_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterEnum_decl(s)
	}
}

func (s *Enum_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitEnum_decl(s)
	}
}

func (p *FlatBuffersParser) Enum_decl() (localctx IEnum_declContext) {
	this := p
	_ = this

	localctx = NewEnum_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FlatBuffersParserRULE_enum_decl)
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
		p.SetState(116)
		p.Match(FlatBuffersParserENUM)
	}
	{
		p.SetState(117)
		p.Identifier()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserCOLON {
		{
			p.SetState(118)
			p.Match(FlatBuffersParserCOLON)
		}
		{
			p.SetState(119)
			p.Type_()
		}

	}
	{
		p.SetState(122)
		p.Metadata()
	}
	{
		p.SetState(123)
		p.Match(FlatBuffersParserLC)
	}
	{
		p.SetState(124)
		p.Commasep_enumval_decl()
	}
	{
		p.SetState(125)
		p.Match(FlatBuffersParserRC)
	}

	return localctx
}

// IUnion_declContext is an interface to support dynamic dispatch.
type IUnion_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnion_declContext differentiates from other interfaces.
	IsUnion_declContext()
}

type Union_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnion_declContext() *Union_declContext {
	var p = new(Union_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_union_decl
	return p
}

func (*Union_declContext) IsUnion_declContext() {}

func NewUnion_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Union_declContext {
	var p = new(Union_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_union_decl

	return p
}

func (s *Union_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Union_declContext) UNION() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserUNION, 0)
}

func (s *Union_declContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Union_declContext) Metadata() IMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataContext)
}

func (s *Union_declContext) LC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserLC, 0)
}

func (s *Union_declContext) Commasep_unionval_with_opt_alias() ICommasep_unionval_with_opt_aliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommasep_unionval_with_opt_aliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommasep_unionval_with_opt_aliasContext)
}

func (s *Union_declContext) RC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRC, 0)
}

func (s *Union_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Union_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Union_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterUnion_decl(s)
	}
}

func (s *Union_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitUnion_decl(s)
	}
}

func (p *FlatBuffersParser) Union_decl() (localctx IUnion_declContext) {
	this := p
	_ = this

	localctx = NewUnion_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FlatBuffersParserRULE_union_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(FlatBuffersParserUNION)
	}
	{
		p.SetState(128)
		p.Identifier()
	}
	{
		p.SetState(129)
		p.Metadata()
	}
	{
		p.SetState(130)
		p.Match(FlatBuffersParserLC)
	}
	{
		p.SetState(131)
		p.Commasep_unionval_with_opt_alias()
	}
	{
		p.SetState(132)
		p.Match(FlatBuffersParserRC)
	}

	return localctx
}

// IRoot_declContext is an interface to support dynamic dispatch.
type IRoot_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoot_declContext differentiates from other interfaces.
	IsRoot_declContext()
}

type Root_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoot_declContext() *Root_declContext {
	var p = new(Root_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_root_decl
	return p
}

func (*Root_declContext) IsRoot_declContext() {}

func NewRoot_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Root_declContext {
	var p = new(Root_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_root_decl

	return p
}

func (s *Root_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Root_declContext) ROOT_TYPE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserROOT_TYPE, 0)
}

func (s *Root_declContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Root_declContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSEMI, 0)
}

func (s *Root_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Root_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Root_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterRoot_decl(s)
	}
}

func (s *Root_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitRoot_decl(s)
	}
}

func (p *FlatBuffersParser) Root_decl() (localctx IRoot_declContext) {
	this := p
	_ = this

	localctx = NewRoot_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FlatBuffersParserRULE_root_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(FlatBuffersParserROOT_TYPE)
	}
	{
		p.SetState(135)
		p.Identifier()
	}
	{
		p.SetState(136)
		p.Match(FlatBuffersParserSEMI)
	}

	return localctx
}

// IField_declContext is an interface to support dynamic dispatch.
type IField_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_declContext differentiates from other interfaces.
	IsField_declContext()
}

type Field_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_declContext() *Field_declContext {
	var p = new(Field_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_field_decl
	return p
}

func (*Field_declContext) IsField_declContext() {}

func NewField_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_declContext {
	var p = new(Field_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_field_decl

	return p
}

func (s *Field_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_declContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Field_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOLON, 0)
}

func (s *Field_declContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Field_declContext) Metadata() IMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataContext)
}

func (s *Field_declContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSEMI, 0)
}

func (s *Field_declContext) EQ() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserEQ, 0)
}

func (s *Field_declContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *Field_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterField_decl(s)
	}
}

func (s *Field_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitField_decl(s)
	}
}

func (p *FlatBuffersParser) Field_decl() (localctx IField_declContext) {
	this := p
	_ = this

	localctx = NewField_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FlatBuffersParserRULE_field_decl)
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
		p.SetState(138)
		p.Identifier()
	}
	{
		p.SetState(139)
		p.Match(FlatBuffersParserCOLON)
	}
	{
		p.SetState(140)
		p.Type_()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserEQ {
		{
			p.SetState(141)
			p.Match(FlatBuffersParserEQ)
		}
		{
			p.SetState(142)
			p.Scalar()
		}

	}
	{
		p.SetState(145)
		p.Metadata()
	}
	{
		p.SetState(146)
		p.Match(FlatBuffersParserSEMI)
	}

	return localctx
}

// IRpc_declContext is an interface to support dynamic dispatch.
type IRpc_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpc_declContext differentiates from other interfaces.
	IsRpc_declContext()
}

type Rpc_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpc_declContext() *Rpc_declContext {
	var p = new(Rpc_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_rpc_decl
	return p
}

func (*Rpc_declContext) IsRpc_declContext() {}

func NewRpc_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rpc_declContext {
	var p = new(Rpc_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_rpc_decl

	return p
}

func (s *Rpc_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Rpc_declContext) RPC_SERVICE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRPC_SERVICE, 0)
}

func (s *Rpc_declContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Rpc_declContext) LC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserLC, 0)
}

func (s *Rpc_declContext) RC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRC, 0)
}

func (s *Rpc_declContext) AllRpc_method() []IRpc_methodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRpc_methodContext)(nil)).Elem())
	var tst = make([]IRpc_methodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRpc_methodContext)
		}
	}

	return tst
}

func (s *Rpc_declContext) Rpc_method(i int) IRpc_methodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpc_methodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRpc_methodContext)
}

func (s *Rpc_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rpc_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rpc_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterRpc_decl(s)
	}
}

func (s *Rpc_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitRpc_decl(s)
	}
}

func (p *FlatBuffersParser) Rpc_decl() (localctx IRpc_declContext) {
	this := p
	_ = this

	localctx = NewRpc_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FlatBuffersParserRULE_rpc_decl)
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
		p.SetState(148)
		p.Match(FlatBuffersParserRPC_SERVICE)
	}
	{
		p.SetState(149)
		p.Identifier()
	}
	{
		p.SetState(150)
		p.Match(FlatBuffersParserLC)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FlatBuffersParserATTRIBUTE)|(1<<FlatBuffersParserENUM)|(1<<FlatBuffersParserFILE_EXTENSION)|(1<<FlatBuffersParserFILE_IDENTIFIER)|(1<<FlatBuffersParserINCLUDE)|(1<<FlatBuffersParserNATIVE_INCLUDE)|(1<<FlatBuffersParserNAMESPACE)|(1<<FlatBuffersParserROOT_TYPE)|(1<<FlatBuffersParserRPC_SERVICE)|(1<<FlatBuffersParserSTRUCT)|(1<<FlatBuffersParserTABLE)|(1<<FlatBuffersParserUNION)|(1<<FlatBuffersParserIDENT))) != 0) {
		{
			p.SetState(151)
			p.Rpc_method()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(FlatBuffersParserRC)
	}

	return localctx
}

// IRpc_methodContext is an interface to support dynamic dispatch.
type IRpc_methodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpc_methodContext differentiates from other interfaces.
	IsRpc_methodContext()
}

type Rpc_methodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpc_methodContext() *Rpc_methodContext {
	var p = new(Rpc_methodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_rpc_method
	return p
}

func (*Rpc_methodContext) IsRpc_methodContext() {}

func NewRpc_methodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rpc_methodContext {
	var p = new(Rpc_methodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_rpc_method

	return p
}

func (s *Rpc_methodContext) GetParser() antlr.Parser { return s.parser }

func (s *Rpc_methodContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *Rpc_methodContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Rpc_methodContext) LP() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserLP, 0)
}

func (s *Rpc_methodContext) RP() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRP, 0)
}

func (s *Rpc_methodContext) COLON() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOLON, 0)
}

func (s *Rpc_methodContext) Metadata() IMetadataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataContext)
}

func (s *Rpc_methodContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSEMI, 0)
}

func (s *Rpc_methodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rpc_methodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rpc_methodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterRpc_method(s)
	}
}

func (s *Rpc_methodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitRpc_method(s)
	}
}

func (p *FlatBuffersParser) Rpc_method() (localctx IRpc_methodContext) {
	this := p
	_ = this

	localctx = NewRpc_methodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FlatBuffersParserRULE_rpc_method)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Identifier()
	}
	{
		p.SetState(159)
		p.Match(FlatBuffersParserLP)
	}
	{
		p.SetState(160)
		p.Identifier()
	}
	{
		p.SetState(161)
		p.Match(FlatBuffersParserRP)
	}
	{
		p.SetState(162)
		p.Match(FlatBuffersParserCOLON)
	}
	{
		p.SetState(163)
		p.Identifier()
	}
	{
		p.SetState(164)
		p.Metadata()
	}
	{
		p.SetState(165)
		p.Match(FlatBuffersParserSEMI)
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
	p.RuleIndex = FlatBuffersParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) LB() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserLB, 0)
}

func (s *Type_Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Type_Context) RB() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRB, 0)
}

func (s *Type_Context) COLON() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOLON, 0)
}

func (s *Type_Context) Integer_const() IInteger_constContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_constContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_constContext)
}

func (s *Type_Context) BASE_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserBASE_TYPE_NAME, 0)
}

func (s *Type_Context) Ns_ident() INs_identContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INs_identContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INs_identContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *FlatBuffersParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FlatBuffersParserRULE_type_)
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

	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FlatBuffersParserLB:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(FlatBuffersParserLB)
		}
		{
			p.SetState(168)
			p.Type_()
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FlatBuffersParserCOLON {
			{
				p.SetState(169)
				p.Match(FlatBuffersParserCOLON)
			}
			{
				p.SetState(170)
				p.Integer_const()
			}

		}
		{
			p.SetState(173)
			p.Match(FlatBuffersParserRB)
		}

	case FlatBuffersParserBASE_TYPE_NAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(FlatBuffersParserBASE_TYPE_NAME)
		}

	case FlatBuffersParserATTRIBUTE, FlatBuffersParserENUM, FlatBuffersParserFILE_EXTENSION, FlatBuffersParserFILE_IDENTIFIER, FlatBuffersParserINCLUDE, FlatBuffersParserNATIVE_INCLUDE, FlatBuffersParserNAMESPACE, FlatBuffersParserROOT_TYPE, FlatBuffersParserRPC_SERVICE, FlatBuffersParserSTRUCT, FlatBuffersParserTABLE, FlatBuffersParserUNION, FlatBuffersParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(176)
			p.Ns_ident()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumval_declContext is an interface to support dynamic dispatch.
type IEnumval_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumval_declContext differentiates from other interfaces.
	IsEnumval_declContext()
}

type Enumval_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumval_declContext() *Enumval_declContext {
	var p = new(Enumval_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_enumval_decl
	return p
}

func (*Enumval_declContext) IsEnumval_declContext() {}

func NewEnumval_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enumval_declContext {
	var p = new(Enumval_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_enumval_decl

	return p
}

func (s *Enumval_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Enumval_declContext) Ns_ident() INs_identContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INs_identContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INs_identContext)
}

func (s *Enumval_declContext) EQ() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserEQ, 0)
}

func (s *Enumval_declContext) Integer_const() IInteger_constContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_constContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_constContext)
}

func (s *Enumval_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enumval_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enumval_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterEnumval_decl(s)
	}
}

func (s *Enumval_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitEnumval_decl(s)
	}
}

func (p *FlatBuffersParser) Enumval_decl() (localctx IEnumval_declContext) {
	this := p
	_ = this

	localctx = NewEnumval_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FlatBuffersParserRULE_enumval_decl)
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
		p.SetState(179)
		p.Ns_ident()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserEQ {
		{
			p.SetState(180)
			p.Match(FlatBuffersParserEQ)
		}
		{
			p.SetState(181)
			p.Integer_const()
		}

	}

	return localctx
}

// ICommasep_enumval_declContext is an interface to support dynamic dispatch.
type ICommasep_enumval_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommasep_enumval_declContext differentiates from other interfaces.
	IsCommasep_enumval_declContext()
}

type Commasep_enumval_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommasep_enumval_declContext() *Commasep_enumval_declContext {
	var p = new(Commasep_enumval_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_commasep_enumval_decl
	return p
}

func (*Commasep_enumval_declContext) IsCommasep_enumval_declContext() {}

func NewCommasep_enumval_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Commasep_enumval_declContext {
	var p = new(Commasep_enumval_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_commasep_enumval_decl

	return p
}

func (s *Commasep_enumval_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Commasep_enumval_declContext) AllEnumval_decl() []IEnumval_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumval_declContext)(nil)).Elem())
	var tst = make([]IEnumval_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumval_declContext)
		}
	}

	return tst
}

func (s *Commasep_enumval_declContext) Enumval_decl(i int) IEnumval_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumval_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumval_declContext)
}

func (s *Commasep_enumval_declContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FlatBuffersParserCOMMA)
}

func (s *Commasep_enumval_declContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOMMA, i)
}

func (s *Commasep_enumval_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Commasep_enumval_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Commasep_enumval_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterCommasep_enumval_decl(s)
	}
}

func (s *Commasep_enumval_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitCommasep_enumval_decl(s)
	}
}

func (p *FlatBuffersParser) Commasep_enumval_decl() (localctx ICommasep_enumval_declContext) {
	this := p
	_ = this

	localctx = NewCommasep_enumval_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FlatBuffersParserRULE_commasep_enumval_decl)
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
		p.SetState(184)
		p.Enumval_decl()
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(185)
				p.Match(FlatBuffersParserCOMMA)
			}
			{
				p.SetState(186)
				p.Enumval_decl()
			}

		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserCOMMA {
		{
			p.SetState(192)
			p.Match(FlatBuffersParserCOMMA)
		}

	}

	return localctx
}

// IUnionval_with_opt_aliasContext is an interface to support dynamic dispatch.
type IUnionval_with_opt_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionval_with_opt_aliasContext differentiates from other interfaces.
	IsUnionval_with_opt_aliasContext()
}

type Unionval_with_opt_aliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionval_with_opt_aliasContext() *Unionval_with_opt_aliasContext {
	var p = new(Unionval_with_opt_aliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_unionval_with_opt_alias
	return p
}

func (*Unionval_with_opt_aliasContext) IsUnionval_with_opt_aliasContext() {}

func NewUnionval_with_opt_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unionval_with_opt_aliasContext {
	var p = new(Unionval_with_opt_aliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_unionval_with_opt_alias

	return p
}

func (s *Unionval_with_opt_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Unionval_with_opt_aliasContext) AllNs_ident() []INs_identContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INs_identContext)(nil)).Elem())
	var tst = make([]INs_identContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INs_identContext)
		}
	}

	return tst
}

func (s *Unionval_with_opt_aliasContext) Ns_ident(i int) INs_identContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INs_identContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INs_identContext)
}

func (s *Unionval_with_opt_aliasContext) COLON() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOLON, 0)
}

func (s *Unionval_with_opt_aliasContext) EQ() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserEQ, 0)
}

func (s *Unionval_with_opt_aliasContext) Integer_const() IInteger_constContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_constContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_constContext)
}

func (s *Unionval_with_opt_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unionval_with_opt_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unionval_with_opt_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterUnionval_with_opt_alias(s)
	}
}

func (s *Unionval_with_opt_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitUnionval_with_opt_alias(s)
	}
}

func (p *FlatBuffersParser) Unionval_with_opt_alias() (localctx IUnionval_with_opt_aliasContext) {
	this := p
	_ = this

	localctx = NewUnionval_with_opt_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FlatBuffersParserRULE_unionval_with_opt_alias)
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
		p.SetState(195)
		p.Ns_ident()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserCOLON {
		{
			p.SetState(196)
			p.Match(FlatBuffersParserCOLON)
		}
		{
			p.SetState(197)
			p.Ns_ident()
		}

	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserEQ {
		{
			p.SetState(200)
			p.Match(FlatBuffersParserEQ)
		}
		{
			p.SetState(201)
			p.Integer_const()
		}

	}

	return localctx
}

// ICommasep_unionval_with_opt_aliasContext is an interface to support dynamic dispatch.
type ICommasep_unionval_with_opt_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommasep_unionval_with_opt_aliasContext differentiates from other interfaces.
	IsCommasep_unionval_with_opt_aliasContext()
}

type Commasep_unionval_with_opt_aliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommasep_unionval_with_opt_aliasContext() *Commasep_unionval_with_opt_aliasContext {
	var p = new(Commasep_unionval_with_opt_aliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_commasep_unionval_with_opt_alias
	return p
}

func (*Commasep_unionval_with_opt_aliasContext) IsCommasep_unionval_with_opt_aliasContext() {}

func NewCommasep_unionval_with_opt_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Commasep_unionval_with_opt_aliasContext {
	var p = new(Commasep_unionval_with_opt_aliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_commasep_unionval_with_opt_alias

	return p
}

func (s *Commasep_unionval_with_opt_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Commasep_unionval_with_opt_aliasContext) AllUnionval_with_opt_alias() []IUnionval_with_opt_aliasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnionval_with_opt_aliasContext)(nil)).Elem())
	var tst = make([]IUnionval_with_opt_aliasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnionval_with_opt_aliasContext)
		}
	}

	return tst
}

func (s *Commasep_unionval_with_opt_aliasContext) Unionval_with_opt_alias(i int) IUnionval_with_opt_aliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionval_with_opt_aliasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnionval_with_opt_aliasContext)
}

func (s *Commasep_unionval_with_opt_aliasContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FlatBuffersParserCOMMA)
}

func (s *Commasep_unionval_with_opt_aliasContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOMMA, i)
}

func (s *Commasep_unionval_with_opt_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Commasep_unionval_with_opt_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Commasep_unionval_with_opt_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterCommasep_unionval_with_opt_alias(s)
	}
}

func (s *Commasep_unionval_with_opt_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitCommasep_unionval_with_opt_alias(s)
	}
}

func (p *FlatBuffersParser) Commasep_unionval_with_opt_alias() (localctx ICommasep_unionval_with_opt_aliasContext) {
	this := p
	_ = this

	localctx = NewCommasep_unionval_with_opt_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FlatBuffersParserRULE_commasep_unionval_with_opt_alias)
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
		p.SetState(204)
		p.Unionval_with_opt_alias()
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(205)
				p.Match(FlatBuffersParserCOMMA)
			}
			{
				p.SetState(206)
				p.Unionval_with_opt_alias()
			}

		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserCOMMA {
		{
			p.SetState(212)
			p.Match(FlatBuffersParserCOMMA)
		}

	}

	return localctx
}

// IIdent_with_opt_single_valueContext is an interface to support dynamic dispatch.
type IIdent_with_opt_single_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdent_with_opt_single_valueContext differentiates from other interfaces.
	IsIdent_with_opt_single_valueContext()
}

type Ident_with_opt_single_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdent_with_opt_single_valueContext() *Ident_with_opt_single_valueContext {
	var p = new(Ident_with_opt_single_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_ident_with_opt_single_value
	return p
}

func (*Ident_with_opt_single_valueContext) IsIdent_with_opt_single_valueContext() {}

func NewIdent_with_opt_single_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ident_with_opt_single_valueContext {
	var p = new(Ident_with_opt_single_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_ident_with_opt_single_value

	return p
}

func (s *Ident_with_opt_single_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Ident_with_opt_single_valueContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Ident_with_opt_single_valueContext) COLON() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOLON, 0)
}

func (s *Ident_with_opt_single_valueContext) Single_value() ISingle_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingle_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingle_valueContext)
}

func (s *Ident_with_opt_single_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ident_with_opt_single_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ident_with_opt_single_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterIdent_with_opt_single_value(s)
	}
}

func (s *Ident_with_opt_single_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitIdent_with_opt_single_value(s)
	}
}

func (p *FlatBuffersParser) Ident_with_opt_single_value() (localctx IIdent_with_opt_single_valueContext) {
	this := p
	_ = this

	localctx = NewIdent_with_opt_single_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FlatBuffersParserRULE_ident_with_opt_single_value)
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
		p.Identifier()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserCOLON {
		{
			p.SetState(216)
			p.Match(FlatBuffersParserCOLON)
		}
		{
			p.SetState(217)
			p.Single_value()
		}

	}

	return localctx
}

// ICommasep_ident_with_opt_single_valueContext is an interface to support dynamic dispatch.
type ICommasep_ident_with_opt_single_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommasep_ident_with_opt_single_valueContext differentiates from other interfaces.
	IsCommasep_ident_with_opt_single_valueContext()
}

type Commasep_ident_with_opt_single_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommasep_ident_with_opt_single_valueContext() *Commasep_ident_with_opt_single_valueContext {
	var p = new(Commasep_ident_with_opt_single_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_commasep_ident_with_opt_single_value
	return p
}

func (*Commasep_ident_with_opt_single_valueContext) IsCommasep_ident_with_opt_single_valueContext() {}

func NewCommasep_ident_with_opt_single_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Commasep_ident_with_opt_single_valueContext {
	var p = new(Commasep_ident_with_opt_single_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_commasep_ident_with_opt_single_value

	return p
}

func (s *Commasep_ident_with_opt_single_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Commasep_ident_with_opt_single_valueContext) AllIdent_with_opt_single_value() []IIdent_with_opt_single_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdent_with_opt_single_valueContext)(nil)).Elem())
	var tst = make([]IIdent_with_opt_single_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdent_with_opt_single_valueContext)
		}
	}

	return tst
}

func (s *Commasep_ident_with_opt_single_valueContext) Ident_with_opt_single_value(i int) IIdent_with_opt_single_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdent_with_opt_single_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdent_with_opt_single_valueContext)
}

func (s *Commasep_ident_with_opt_single_valueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FlatBuffersParserCOMMA)
}

func (s *Commasep_ident_with_opt_single_valueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOMMA, i)
}

func (s *Commasep_ident_with_opt_single_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Commasep_ident_with_opt_single_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Commasep_ident_with_opt_single_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterCommasep_ident_with_opt_single_value(s)
	}
}

func (s *Commasep_ident_with_opt_single_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitCommasep_ident_with_opt_single_value(s)
	}
}

func (p *FlatBuffersParser) Commasep_ident_with_opt_single_value() (localctx ICommasep_ident_with_opt_single_valueContext) {
	this := p
	_ = this

	localctx = NewCommasep_ident_with_opt_single_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FlatBuffersParserRULE_commasep_ident_with_opt_single_value)
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
		p.SetState(220)
		p.Ident_with_opt_single_value()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FlatBuffersParserCOMMA {
		{
			p.SetState(221)
			p.Match(FlatBuffersParserCOMMA)
		}
		{
			p.SetState(222)
			p.Ident_with_opt_single_value()
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMetadataContext is an interface to support dynamic dispatch.
type IMetadataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetadataContext differentiates from other interfaces.
	IsMetadataContext()
}

type MetadataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetadataContext() *MetadataContext {
	var p = new(MetadataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_metadata
	return p
}

func (*MetadataContext) IsMetadataContext() {}

func NewMetadataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetadataContext {
	var p = new(MetadataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_metadata

	return p
}

func (s *MetadataContext) GetParser() antlr.Parser { return s.parser }

func (s *MetadataContext) LP() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserLP, 0)
}

func (s *MetadataContext) Commasep_ident_with_opt_single_value() ICommasep_ident_with_opt_single_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommasep_ident_with_opt_single_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommasep_ident_with_opt_single_valueContext)
}

func (s *MetadataContext) RP() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRP, 0)
}

func (s *MetadataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetadataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetadataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterMetadata(s)
	}
}

func (s *MetadataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitMetadata(s)
	}
}

func (p *FlatBuffersParser) Metadata() (localctx IMetadataContext) {
	this := p
	_ = this

	localctx = NewMetadataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FlatBuffersParserRULE_metadata)
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
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserLP {
		{
			p.SetState(228)
			p.Match(FlatBuffersParserLP)
		}
		{
			p.SetState(229)
			p.Commasep_ident_with_opt_single_value()
		}
		{
			p.SetState(230)
			p.Match(FlatBuffersParserRP)
		}

	}

	return localctx
}

// IScalarContext is an interface to support dynamic dispatch.
type IScalarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarContext differentiates from other interfaces.
	IsScalarContext()
}

type ScalarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarContext() *ScalarContext {
	var p = new(ScalarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_scalar
	return p
}

func (*ScalarContext) IsScalarContext() {}

func NewScalarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarContext {
	var p = new(ScalarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_scalar

	return p
}

func (s *ScalarContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarContext) INTEGER_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserINTEGER_CONSTANT, 0)
}

func (s *ScalarContext) HEX_INTEGER_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserHEX_INTEGER_CONSTANT, 0)
}

func (s *ScalarContext) FLOAT_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserFLOAT_CONSTANT, 0)
}

func (s *ScalarContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ScalarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterScalar(s)
	}
}

func (s *ScalarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitScalar(s)
	}
}

func (p *FlatBuffersParser) Scalar() (localctx IScalarContext) {
	this := p
	_ = this

	localctx = NewScalarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FlatBuffersParserRULE_scalar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(238)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FlatBuffersParserINTEGER_CONSTANT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Match(FlatBuffersParserINTEGER_CONSTANT)
		}

	case FlatBuffersParserHEX_INTEGER_CONSTANT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(FlatBuffersParserHEX_INTEGER_CONSTANT)
		}

	case FlatBuffersParserFLOAT_CONSTANT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(236)
			p.Match(FlatBuffersParserFLOAT_CONSTANT)
		}

	case FlatBuffersParserATTRIBUTE, FlatBuffersParserENUM, FlatBuffersParserFILE_EXTENSION, FlatBuffersParserFILE_IDENTIFIER, FlatBuffersParserINCLUDE, FlatBuffersParserNATIVE_INCLUDE, FlatBuffersParserNAMESPACE, FlatBuffersParserROOT_TYPE, FlatBuffersParserRPC_SERVICE, FlatBuffersParserSTRUCT, FlatBuffersParserTABLE, FlatBuffersParserUNION, FlatBuffersParserIDENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(237)
			p.Identifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObject_Context is an interface to support dynamic dispatch.
type IObject_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_Context differentiates from other interfaces.
	IsObject_Context()
}

type Object_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_Context() *Object_Context {
	var p = new(Object_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_object_
	return p
}

func (*Object_Context) IsObject_Context() {}

func NewObject_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_Context {
	var p = new(Object_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_object_

	return p
}

func (s *Object_Context) GetParser() antlr.Parser { return s.parser }

func (s *Object_Context) LC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserLC, 0)
}

func (s *Object_Context) Commasep_ident_with_value() ICommasep_ident_with_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommasep_ident_with_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommasep_ident_with_valueContext)
}

func (s *Object_Context) RC() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRC, 0)
}

func (s *Object_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterObject_(s)
	}
}

func (s *Object_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitObject_(s)
	}
}

func (p *FlatBuffersParser) Object_() (localctx IObject_Context) {
	this := p
	_ = this

	localctx = NewObject_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FlatBuffersParserRULE_object_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(FlatBuffersParserLC)
	}
	{
		p.SetState(241)
		p.Commasep_ident_with_value()
	}
	{
		p.SetState(242)
		p.Match(FlatBuffersParserRC)
	}

	return localctx
}

// IIdent_with_valueContext is an interface to support dynamic dispatch.
type IIdent_with_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdent_with_valueContext differentiates from other interfaces.
	IsIdent_with_valueContext()
}

type Ident_with_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdent_with_valueContext() *Ident_with_valueContext {
	var p = new(Ident_with_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_ident_with_value
	return p
}

func (*Ident_with_valueContext) IsIdent_with_valueContext() {}

func NewIdent_with_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ident_with_valueContext {
	var p = new(Ident_with_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_ident_with_value

	return p
}

func (s *Ident_with_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Ident_with_valueContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Ident_with_valueContext) COLON() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOLON, 0)
}

func (s *Ident_with_valueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Ident_with_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ident_with_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ident_with_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterIdent_with_value(s)
	}
}

func (s *Ident_with_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitIdent_with_value(s)
	}
}

func (p *FlatBuffersParser) Ident_with_value() (localctx IIdent_with_valueContext) {
	this := p
	_ = this

	localctx = NewIdent_with_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FlatBuffersParserRULE_ident_with_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Identifier()
	}
	{
		p.SetState(245)
		p.Match(FlatBuffersParserCOLON)
	}
	{
		p.SetState(246)
		p.Value()
	}

	return localctx
}

// ICommasep_ident_with_valueContext is an interface to support dynamic dispatch.
type ICommasep_ident_with_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommasep_ident_with_valueContext differentiates from other interfaces.
	IsCommasep_ident_with_valueContext()
}

type Commasep_ident_with_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommasep_ident_with_valueContext() *Commasep_ident_with_valueContext {
	var p = new(Commasep_ident_with_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_commasep_ident_with_value
	return p
}

func (*Commasep_ident_with_valueContext) IsCommasep_ident_with_valueContext() {}

func NewCommasep_ident_with_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Commasep_ident_with_valueContext {
	var p = new(Commasep_ident_with_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_commasep_ident_with_value

	return p
}

func (s *Commasep_ident_with_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Commasep_ident_with_valueContext) AllIdent_with_value() []IIdent_with_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdent_with_valueContext)(nil)).Elem())
	var tst = make([]IIdent_with_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdent_with_valueContext)
		}
	}

	return tst
}

func (s *Commasep_ident_with_valueContext) Ident_with_value(i int) IIdent_with_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdent_with_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdent_with_valueContext)
}

func (s *Commasep_ident_with_valueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FlatBuffersParserCOMMA)
}

func (s *Commasep_ident_with_valueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOMMA, i)
}

func (s *Commasep_ident_with_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Commasep_ident_with_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Commasep_ident_with_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterCommasep_ident_with_value(s)
	}
}

func (s *Commasep_ident_with_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitCommasep_ident_with_value(s)
	}
}

func (p *FlatBuffersParser) Commasep_ident_with_value() (localctx ICommasep_ident_with_valueContext) {
	this := p
	_ = this

	localctx = NewCommasep_ident_with_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FlatBuffersParserRULE_commasep_ident_with_value)
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
		p.SetState(248)
		p.Ident_with_value()
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(249)
				p.Match(FlatBuffersParserCOMMA)
			}
			{
				p.SetState(250)
				p.Ident_with_value()
			}

		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserCOMMA {
		{
			p.SetState(256)
			p.Match(FlatBuffersParserCOMMA)
		}

	}

	return localctx
}

// ISingle_valueContext is an interface to support dynamic dispatch.
type ISingle_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingle_valueContext differentiates from other interfaces.
	IsSingle_valueContext()
}

type Single_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_valueContext() *Single_valueContext {
	var p = new(Single_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_single_value
	return p
}

func (*Single_valueContext) IsSingle_valueContext() {}

func NewSingle_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_valueContext {
	var p = new(Single_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_single_value

	return p
}

func (s *Single_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_valueContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *Single_valueContext) STRING_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSTRING_CONSTANT, 0)
}

func (s *Single_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterSingle_value(s)
	}
}

func (s *Single_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitSingle_value(s)
	}
}

func (p *FlatBuffersParser) Single_value() (localctx ISingle_valueContext) {
	this := p
	_ = this

	localctx = NewSingle_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FlatBuffersParserRULE_single_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FlatBuffersParserATTRIBUTE, FlatBuffersParserENUM, FlatBuffersParserFILE_EXTENSION, FlatBuffersParserFILE_IDENTIFIER, FlatBuffersParserINCLUDE, FlatBuffersParserNATIVE_INCLUDE, FlatBuffersParserNAMESPACE, FlatBuffersParserROOT_TYPE, FlatBuffersParserRPC_SERVICE, FlatBuffersParserSTRUCT, FlatBuffersParserTABLE, FlatBuffersParserUNION, FlatBuffersParserINTEGER_CONSTANT, FlatBuffersParserIDENT, FlatBuffersParserHEX_INTEGER_CONSTANT, FlatBuffersParserFLOAT_CONSTANT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Scalar()
		}

	case FlatBuffersParserSTRING_CONSTANT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Match(FlatBuffersParserSTRING_CONSTANT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = FlatBuffersParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Single_value() ISingle_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingle_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingle_valueContext)
}

func (s *ValueContext) Object_() IObject_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_Context)
}

func (s *ValueContext) LB() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserLB, 0)
}

func (s *ValueContext) Commasep_value() ICommasep_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommasep_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommasep_valueContext)
}

func (s *ValueContext) RB() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRB, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *FlatBuffersParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FlatBuffersParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	switch p.GetTokenStream().LA(1) {
	case FlatBuffersParserATTRIBUTE, FlatBuffersParserENUM, FlatBuffersParserFILE_EXTENSION, FlatBuffersParserFILE_IDENTIFIER, FlatBuffersParserINCLUDE, FlatBuffersParserNATIVE_INCLUDE, FlatBuffersParserNAMESPACE, FlatBuffersParserROOT_TYPE, FlatBuffersParserRPC_SERVICE, FlatBuffersParserSTRUCT, FlatBuffersParserTABLE, FlatBuffersParserUNION, FlatBuffersParserSTRING_CONSTANT, FlatBuffersParserINTEGER_CONSTANT, FlatBuffersParserIDENT, FlatBuffersParserHEX_INTEGER_CONSTANT, FlatBuffersParserFLOAT_CONSTANT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Single_value()
		}

	case FlatBuffersParserLC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.Object_()
		}

	case FlatBuffersParserLB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(265)
			p.Match(FlatBuffersParserLB)
		}
		{
			p.SetState(266)
			p.Commasep_value()
		}
		{
			p.SetState(267)
			p.Match(FlatBuffersParserRB)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommasep_valueContext is an interface to support dynamic dispatch.
type ICommasep_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommasep_valueContext differentiates from other interfaces.
	IsCommasep_valueContext()
}

type Commasep_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommasep_valueContext() *Commasep_valueContext {
	var p = new(Commasep_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_commasep_value
	return p
}

func (*Commasep_valueContext) IsCommasep_valueContext() {}

func NewCommasep_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Commasep_valueContext {
	var p = new(Commasep_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_commasep_value

	return p
}

func (s *Commasep_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Commasep_valueContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Commasep_valueContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Commasep_valueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FlatBuffersParserCOMMA)
}

func (s *Commasep_valueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserCOMMA, i)
}

func (s *Commasep_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Commasep_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Commasep_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterCommasep_value(s)
	}
}

func (s *Commasep_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitCommasep_value(s)
	}
}

func (p *FlatBuffersParser) Commasep_value() (localctx ICommasep_valueContext) {
	this := p
	_ = this

	localctx = NewCommasep_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FlatBuffersParserRULE_commasep_value)
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
		p.SetState(271)
		p.Value()
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(272)
				p.Match(FlatBuffersParserCOMMA)
			}
			{
				p.SetState(273)
				p.Value()
			}

		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FlatBuffersParserCOMMA {
		{
			p.SetState(279)
			p.Match(FlatBuffersParserCOMMA)
		}

	}

	return localctx
}

// IFile_extension_declContext is an interface to support dynamic dispatch.
type IFile_extension_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_extension_declContext differentiates from other interfaces.
	IsFile_extension_declContext()
}

type File_extension_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_extension_declContext() *File_extension_declContext {
	var p = new(File_extension_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_file_extension_decl
	return p
}

func (*File_extension_declContext) IsFile_extension_declContext() {}

func NewFile_extension_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_extension_declContext {
	var p = new(File_extension_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_file_extension_decl

	return p
}

func (s *File_extension_declContext) GetParser() antlr.Parser { return s.parser }

func (s *File_extension_declContext) FILE_EXTENSION() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserFILE_EXTENSION, 0)
}

func (s *File_extension_declContext) STRING_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSTRING_CONSTANT, 0)
}

func (s *File_extension_declContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSEMI, 0)
}

func (s *File_extension_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_extension_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_extension_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterFile_extension_decl(s)
	}
}

func (s *File_extension_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitFile_extension_decl(s)
	}
}

func (p *FlatBuffersParser) File_extension_decl() (localctx IFile_extension_declContext) {
	this := p
	_ = this

	localctx = NewFile_extension_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FlatBuffersParserRULE_file_extension_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(FlatBuffersParserFILE_EXTENSION)
	}
	{
		p.SetState(283)
		p.Match(FlatBuffersParserSTRING_CONSTANT)
	}
	{
		p.SetState(284)
		p.Match(FlatBuffersParserSEMI)
	}

	return localctx
}

// IFile_identifier_declContext is an interface to support dynamic dispatch.
type IFile_identifier_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_identifier_declContext differentiates from other interfaces.
	IsFile_identifier_declContext()
}

type File_identifier_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_identifier_declContext() *File_identifier_declContext {
	var p = new(File_identifier_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_file_identifier_decl
	return p
}

func (*File_identifier_declContext) IsFile_identifier_declContext() {}

func NewFile_identifier_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_identifier_declContext {
	var p = new(File_identifier_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_file_identifier_decl

	return p
}

func (s *File_identifier_declContext) GetParser() antlr.Parser { return s.parser }

func (s *File_identifier_declContext) FILE_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserFILE_IDENTIFIER, 0)
}

func (s *File_identifier_declContext) STRING_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSTRING_CONSTANT, 0)
}

func (s *File_identifier_declContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSEMI, 0)
}

func (s *File_identifier_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_identifier_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_identifier_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterFile_identifier_decl(s)
	}
}

func (s *File_identifier_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitFile_identifier_decl(s)
	}
}

func (p *FlatBuffersParser) File_identifier_decl() (localctx IFile_identifier_declContext) {
	this := p
	_ = this

	localctx = NewFile_identifier_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FlatBuffersParserRULE_file_identifier_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(FlatBuffersParserFILE_IDENTIFIER)
	}
	{
		p.SetState(287)
		p.Match(FlatBuffersParserSTRING_CONSTANT)
	}
	{
		p.SetState(288)
		p.Match(FlatBuffersParserSEMI)
	}

	return localctx
}

// INs_identContext is an interface to support dynamic dispatch.
type INs_identContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNs_identContext differentiates from other interfaces.
	IsNs_identContext()
}

type Ns_identContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNs_identContext() *Ns_identContext {
	var p = new(Ns_identContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_ns_ident
	return p
}

func (*Ns_identContext) IsNs_identContext() {}

func NewNs_identContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ns_identContext {
	var p = new(Ns_identContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_ns_ident

	return p
}

func (s *Ns_identContext) GetParser() antlr.Parser { return s.parser }

func (s *Ns_identContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *Ns_identContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Ns_identContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(FlatBuffersParserDOT)
}

func (s *Ns_identContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserDOT, i)
}

func (s *Ns_identContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ns_identContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ns_identContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterNs_ident(s)
	}
}

func (s *Ns_identContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitNs_ident(s)
	}
}

func (p *FlatBuffersParser) Ns_ident() (localctx INs_identContext) {
	this := p
	_ = this

	localctx = NewNs_identContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FlatBuffersParserRULE_ns_ident)
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
		p.SetState(290)
		p.Identifier()
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FlatBuffersParserDOT {
		{
			p.SetState(291)
			p.Match(FlatBuffersParserDOT)
		}
		{
			p.SetState(292)
			p.Identifier()
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInteger_constContext is an interface to support dynamic dispatch.
type IInteger_constContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_constContext differentiates from other interfaces.
	IsInteger_constContext()
}

type Integer_constContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_constContext() *Integer_constContext {
	var p = new(Integer_constContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FlatBuffersParserRULE_integer_const
	return p
}

func (*Integer_constContext) IsInteger_constContext() {}

func NewInteger_constContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_constContext {
	var p = new(Integer_constContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_integer_const

	return p
}

func (s *Integer_constContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_constContext) INTEGER_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserINTEGER_CONSTANT, 0)
}

func (s *Integer_constContext) HEX_INTEGER_CONSTANT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserHEX_INTEGER_CONSTANT, 0)
}

func (s *Integer_constContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_constContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_constContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterInteger_const(s)
	}
}

func (s *Integer_constContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitInteger_const(s)
	}
}

func (p *FlatBuffersParser) Integer_const() (localctx IInteger_constContext) {
	this := p
	_ = this

	localctx = NewInteger_constContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FlatBuffersParserRULE_integer_const)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == FlatBuffersParserINTEGER_CONSTANT || _la == FlatBuffersParserHEX_INTEGER_CONSTANT) {
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
	p.RuleIndex = FlatBuffersParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserIDENT, 0)
}

func (s *IdentifierContext) Keywords() IKeywordsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordsContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *FlatBuffersParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FlatBuffersParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	switch p.GetTokenStream().LA(1) {
	case FlatBuffersParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(300)
			p.Match(FlatBuffersParserIDENT)
		}

	case FlatBuffersParserATTRIBUTE, FlatBuffersParserENUM, FlatBuffersParserFILE_EXTENSION, FlatBuffersParserFILE_IDENTIFIER, FlatBuffersParserINCLUDE, FlatBuffersParserNATIVE_INCLUDE, FlatBuffersParserNAMESPACE, FlatBuffersParserROOT_TYPE, FlatBuffersParserRPC_SERVICE, FlatBuffersParserSTRUCT, FlatBuffersParserTABLE, FlatBuffersParserUNION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(301)
			p.Keywords()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = FlatBuffersParserRULE_keywords
	return p
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlatBuffersParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) ATTRIBUTE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserATTRIBUTE, 0)
}

func (s *KeywordsContext) ENUM() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserENUM, 0)
}

func (s *KeywordsContext) FILE_EXTENSION() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserFILE_EXTENSION, 0)
}

func (s *KeywordsContext) FILE_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserFILE_IDENTIFIER, 0)
}

func (s *KeywordsContext) INCLUDE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserINCLUDE, 0)
}

func (s *KeywordsContext) NATIVE_INCLUDE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserNATIVE_INCLUDE, 0)
}

func (s *KeywordsContext) NAMESPACE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserNAMESPACE, 0)
}

func (s *KeywordsContext) ROOT_TYPE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserROOT_TYPE, 0)
}

func (s *KeywordsContext) RPC_SERVICE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserRPC_SERVICE, 0)
}

func (s *KeywordsContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserSTRUCT, 0)
}

func (s *KeywordsContext) TABLE() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserTABLE, 0)
}

func (s *KeywordsContext) UNION() antlr.TerminalNode {
	return s.GetToken(FlatBuffersParserUNION, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.EnterKeywords(s)
	}
}

func (s *KeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlatBuffersListener); ok {
		listenerT.ExitKeywords(s)
	}
}

func (p *FlatBuffersParser) Keywords() (localctx IKeywordsContext) {
	this := p
	_ = this

	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FlatBuffersParserRULE_keywords)
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
		p.SetState(304)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FlatBuffersParserATTRIBUTE)|(1<<FlatBuffersParserENUM)|(1<<FlatBuffersParserFILE_EXTENSION)|(1<<FlatBuffersParserFILE_IDENTIFIER)|(1<<FlatBuffersParserINCLUDE)|(1<<FlatBuffersParserNATIVE_INCLUDE)|(1<<FlatBuffersParserNAMESPACE)|(1<<FlatBuffersParserROOT_TYPE)|(1<<FlatBuffersParserRPC_SERVICE)|(1<<FlatBuffersParserSTRUCT)|(1<<FlatBuffersParserTABLE)|(1<<FlatBuffersParserUNION))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
