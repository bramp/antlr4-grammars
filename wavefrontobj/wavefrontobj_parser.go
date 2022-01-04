// Code generated from WavefrontOBJ.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wavefrontobj // WavefrontOBJ
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 421,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	3, 2, 7, 2, 90, 10, 2, 12, 2, 14, 2, 93, 11, 2, 3, 2, 3, 2, 3, 2, 6, 2,
	98, 10, 2, 13, 2, 14, 2, 99, 3, 2, 7, 2, 103, 10, 2, 12, 2, 14, 2, 106,
	11, 2, 3, 2, 7, 2, 109, 10, 2, 12, 2, 14, 2, 112, 11, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 145, 10, 3, 3, 4, 3, 4, 7, 4, 149, 10,
	4, 12, 4, 14, 4, 152, 11, 4, 3, 5, 3, 5, 5, 5, 156, 10, 5, 3, 5, 7, 5,
	159, 10, 5, 12, 5, 14, 5, 162, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 169, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 174, 10, 7, 3, 7, 5, 7, 177, 10,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 187, 10, 9, 3,
	9, 5, 9, 190, 10, 9, 3, 10, 3, 10, 5, 10, 194, 10, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 5, 11, 201, 10, 11, 3, 12, 3, 12, 3, 12, 6, 12, 206,
	10, 12, 13, 12, 14, 12, 207, 3, 13, 3, 13, 5, 13, 212, 10, 13, 3, 14, 3,
	14, 6, 14, 216, 10, 14, 13, 14, 14, 14, 217, 3, 15, 3, 15, 6, 15, 222,
	10, 15, 13, 15, 14, 15, 223, 3, 16, 3, 16, 6, 16, 228, 10, 16, 13, 16,
	14, 16, 229, 3, 17, 3, 17, 3, 17, 5, 17, 235, 10, 17, 3, 17, 6, 17, 238,
	10, 17, 13, 17, 14, 17, 239, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17,
	247, 10, 17, 3, 17, 6, 17, 250, 10, 17, 13, 17, 14, 17, 251, 7, 17, 254,
	10, 17, 12, 17, 14, 17, 257, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 6, 18, 265, 10, 18, 13, 18, 14, 18, 266, 3, 19, 3, 19, 3, 19, 6,
	19, 272, 10, 19, 13, 19, 14, 19, 273, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 6, 20, 282, 10, 20, 13, 20, 14, 20, 283, 3, 21, 3, 21, 3, 21, 6,
	21, 289, 10, 21, 13, 21, 14, 21, 290, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	6, 22, 298, 10, 22, 13, 22, 14, 22, 299, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 6, 23, 307, 10, 23, 13, 23, 14, 23, 308, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 6, 24, 316, 10, 24, 13, 24, 14, 24, 317, 3, 25, 3, 25, 6, 25, 322,
	10, 25, 13, 25, 14, 25, 323, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 6, 28, 340, 10,
	28, 13, 28, 14, 28, 341, 3, 29, 3, 29, 3, 29, 5, 29, 347, 10, 29, 3, 30,
	3, 30, 3, 30, 3, 30, 5, 30, 353, 10, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35,
	3, 36, 3, 36, 6, 36, 372, 10, 36, 13, 36, 14, 36, 373, 3, 37, 3, 37, 3,
	37, 5, 37, 379, 10, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 40,
	3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 5, 42, 402, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 417,
	10, 43, 3, 44, 3, 44, 3, 44, 2, 2, 45, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 2, 8, 3,
	2, 10, 14, 4, 2, 4, 4, 17, 17, 4, 2, 53, 53, 55, 55, 3, 2, 53, 55, 4, 2,
	33, 33, 37, 37, 3, 2, 55, 56, 2, 453, 2, 91, 3, 2, 2, 2, 4, 144, 3, 2,
	2, 2, 6, 146, 3, 2, 2, 2, 8, 153, 3, 2, 2, 2, 10, 163, 3, 2, 2, 2, 12,
	170, 3, 2, 2, 2, 14, 178, 3, 2, 2, 2, 16, 183, 3, 2, 2, 2, 18, 191, 3,
	2, 2, 2, 20, 197, 3, 2, 2, 2, 22, 202, 3, 2, 2, 2, 24, 209, 3, 2, 2, 2,
	26, 213, 3, 2, 2, 2, 28, 219, 3, 2, 2, 2, 30, 225, 3, 2, 2, 2, 32, 234,
	3, 2, 2, 2, 34, 260, 3, 2, 2, 2, 36, 268, 3, 2, 2, 2, 38, 275, 3, 2, 2,
	2, 40, 285, 3, 2, 2, 2, 42, 292, 3, 2, 2, 2, 44, 301, 3, 2, 2, 2, 46, 310,
	3, 2, 2, 2, 48, 319, 3, 2, 2, 2, 50, 325, 3, 2, 2, 2, 52, 327, 3, 2, 2,
	2, 54, 337, 3, 2, 2, 2, 56, 343, 3, 2, 2, 2, 58, 348, 3, 2, 2, 2, 60, 354,
	3, 2, 2, 2, 62, 357, 3, 2, 2, 2, 64, 360, 3, 2, 2, 2, 66, 363, 3, 2, 2,
	2, 68, 366, 3, 2, 2, 2, 70, 369, 3, 2, 2, 2, 72, 375, 3, 2, 2, 2, 74, 380,
	3, 2, 2, 2, 76, 383, 3, 2, 2, 2, 78, 386, 3, 2, 2, 2, 80, 389, 3, 2, 2,
	2, 82, 392, 3, 2, 2, 2, 84, 403, 3, 2, 2, 2, 86, 418, 3, 2, 2, 2, 88, 90,
	7, 61, 2, 2, 89, 88, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2,
	91, 92, 3, 2, 2, 2, 92, 94, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 95, 5,
	4, 3, 2, 95, 104, 3, 2, 2, 2, 96, 98, 7, 61, 2, 2, 97, 96, 3, 2, 2, 2,
	98, 99, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 101,
	3, 2, 2, 2, 101, 103, 5, 4, 3, 2, 102, 97, 3, 2, 2, 2, 103, 106, 3, 2,
	2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 110, 3, 2, 2, 2,
	106, 104, 3, 2, 2, 2, 107, 109, 7, 61, 2, 2, 108, 107, 3, 2, 2, 2, 109,
	112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 3, 3,
	2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 145, 5, 6, 4, 2, 114, 145, 5, 8, 5,
	2, 115, 145, 5, 10, 6, 2, 116, 145, 5, 14, 8, 2, 117, 145, 5, 16, 9, 2,
	118, 145, 5, 12, 7, 2, 119, 145, 5, 26, 14, 2, 120, 145, 5, 28, 15, 2,
	121, 145, 5, 30, 16, 2, 122, 145, 5, 18, 10, 2, 123, 145, 5, 20, 11, 2,
	124, 145, 5, 22, 12, 2, 125, 145, 5, 24, 13, 2, 126, 145, 5, 32, 17, 2,
	127, 145, 5, 52, 27, 2, 128, 145, 5, 54, 28, 2, 129, 145, 5, 56, 29, 2,
	130, 145, 5, 58, 30, 2, 131, 145, 5, 60, 31, 2, 132, 145, 5, 62, 32, 2,
	133, 145, 5, 64, 33, 2, 134, 145, 5, 66, 34, 2, 135, 145, 5, 68, 35, 2,
	136, 145, 5, 70, 36, 2, 137, 145, 5, 72, 37, 2, 138, 145, 5, 74, 38, 2,
	139, 145, 5, 76, 39, 2, 140, 145, 5, 78, 40, 2, 141, 145, 5, 80, 41, 2,
	142, 145, 5, 82, 42, 2, 143, 145, 5, 84, 43, 2, 144, 113, 3, 2, 2, 2, 144,
	114, 3, 2, 2, 2, 144, 115, 3, 2, 2, 2, 144, 116, 3, 2, 2, 2, 144, 117,
	3, 2, 2, 2, 144, 118, 3, 2, 2, 2, 144, 119, 3, 2, 2, 2, 144, 120, 3, 2,
	2, 2, 144, 121, 3, 2, 2, 2, 144, 122, 3, 2, 2, 2, 144, 123, 3, 2, 2, 2,
	144, 124, 3, 2, 2, 2, 144, 125, 3, 2, 2, 2, 144, 126, 3, 2, 2, 2, 144,
	127, 3, 2, 2, 2, 144, 128, 3, 2, 2, 2, 144, 129, 3, 2, 2, 2, 144, 130,
	3, 2, 2, 2, 144, 131, 3, 2, 2, 2, 144, 132, 3, 2, 2, 2, 144, 133, 3, 2,
	2, 2, 144, 134, 3, 2, 2, 2, 144, 135, 3, 2, 2, 2, 144, 136, 3, 2, 2, 2,
	144, 137, 3, 2, 2, 2, 144, 138, 3, 2, 2, 2, 144, 139, 3, 2, 2, 2, 144,
	140, 3, 2, 2, 2, 144, 141, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 143,
	3, 2, 2, 2, 145, 5, 3, 2, 2, 2, 146, 150, 7, 59, 2, 2, 147, 149, 7, 63,
	2, 2, 148, 147, 3, 2, 2, 2, 149, 152, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2,
	150, 151, 3, 2, 2, 2, 151, 7, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 153, 155,
	7, 59, 2, 2, 154, 156, 7, 3, 2, 2, 155, 154, 3, 2, 2, 2, 155, 156, 3, 2,
	2, 2, 156, 160, 3, 2, 2, 2, 157, 159, 7, 63, 2, 2, 158, 157, 3, 2, 2, 2,
	159, 162, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161,
	9, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 163, 164, 7, 4, 2, 2, 164, 165, 5,
	86, 44, 2, 165, 166, 5, 86, 44, 2, 166, 168, 5, 86, 44, 2, 167, 169, 5,
	86, 44, 2, 168, 167, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 11, 3, 2, 2,
	2, 170, 171, 7, 5, 2, 2, 171, 173, 5, 86, 44, 2, 172, 174, 5, 86, 44, 2,
	173, 172, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 176, 3, 2, 2, 2, 175,
	177, 5, 86, 44, 2, 176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 13,
	3, 2, 2, 2, 178, 179, 7, 6, 2, 2, 179, 180, 5, 86, 44, 2, 180, 181, 5,
	86, 44, 2, 181, 182, 5, 86, 44, 2, 182, 15, 3, 2, 2, 2, 183, 184, 7, 7,
	2, 2, 184, 186, 5, 86, 44, 2, 185, 187, 5, 86, 44, 2, 186, 185, 3, 2, 2,
	2, 186, 187, 3, 2, 2, 2, 187, 189, 3, 2, 2, 2, 188, 190, 5, 86, 44, 2,
	189, 188, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 17, 3, 2, 2, 2, 191, 193,
	7, 8, 2, 2, 192, 194, 7, 9, 2, 2, 193, 192, 3, 2, 2, 2, 193, 194, 3, 2,
	2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 9, 2, 2, 2, 196, 19, 3, 2, 2, 2,
	197, 198, 7, 15, 2, 2, 198, 200, 7, 55, 2, 2, 199, 201, 7, 55, 2, 2, 200,
	199, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 21, 3, 2, 2, 2, 202, 203, 7,
	16, 2, 2, 203, 205, 9, 3, 2, 2, 204, 206, 5, 86, 44, 2, 205, 204, 3, 2,
	2, 2, 206, 207, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2,
	208, 23, 3, 2, 2, 2, 209, 211, 7, 55, 2, 2, 210, 212, 7, 55, 2, 2, 211,
	210, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 25, 3, 2, 2, 2, 213, 215, 7,
	18, 2, 2, 214, 216, 7, 55, 2, 2, 215, 214, 3, 2, 2, 2, 216, 217, 3, 2,
	2, 2, 217, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 27, 3, 2, 2, 2,
	219, 221, 7, 19, 2, 2, 220, 222, 9, 4, 2, 2, 221, 220, 3, 2, 2, 2, 222,
	223, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 29, 3,
	2, 2, 2, 225, 227, 7, 20, 2, 2, 226, 228, 9, 5, 2, 2, 227, 226, 3, 2, 2,
	2, 228, 229, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230,
	31, 3, 2, 2, 2, 231, 235, 5, 34, 18, 2, 232, 235, 5, 36, 19, 2, 233, 235,
	5, 38, 20, 2, 234, 231, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 233, 3,
	2, 2, 2, 235, 237, 3, 2, 2, 2, 236, 238, 7, 61, 2, 2, 237, 236, 3, 2, 2,
	2, 238, 239, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240,
	255, 3, 2, 2, 2, 241, 247, 5, 40, 21, 2, 242, 247, 5, 42, 22, 2, 243, 247,
	5, 44, 23, 2, 244, 247, 5, 46, 24, 2, 245, 247, 5, 48, 25, 2, 246, 241,
	3, 2, 2, 2, 246, 242, 3, 2, 2, 2, 246, 243, 3, 2, 2, 2, 246, 244, 3, 2,
	2, 2, 246, 245, 3, 2, 2, 2, 247, 249, 3, 2, 2, 2, 248, 250, 7, 61, 2, 2,
	249, 248, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 251,
	252, 3, 2, 2, 2, 252, 254, 3, 2, 2, 2, 253, 246, 3, 2, 2, 2, 254, 257,
	3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 258, 3, 2,
	2, 2, 257, 255, 3, 2, 2, 2, 258, 259, 5, 50, 26, 2, 259, 33, 3, 2, 2, 2,
	260, 261, 7, 21, 2, 2, 261, 262, 5, 86, 44, 2, 262, 264, 5, 86, 44, 2,
	263, 265, 7, 55, 2, 2, 264, 263, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266,
	264, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 35, 3, 2, 2, 2, 268, 269, 7,
	22, 2, 2, 269, 271, 7, 55, 2, 2, 270, 272, 7, 55, 2, 2, 271, 270, 3, 2,
	2, 2, 272, 273, 3, 2, 2, 2, 273, 271, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2,
	274, 37, 3, 2, 2, 2, 275, 276, 7, 23, 2, 2, 276, 277, 5, 86, 44, 2, 277,
	278, 5, 86, 44, 2, 278, 279, 5, 86, 44, 2, 279, 281, 5, 86, 44, 2, 280,
	282, 9, 5, 2, 2, 281, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 281,
	3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 39, 3, 2, 2, 2, 285, 286, 7, 24,
	2, 2, 286, 288, 9, 3, 2, 2, 287, 289, 5, 86, 44, 2, 288, 287, 3, 2, 2,
	2, 289, 290, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291,
	41, 3, 2, 2, 2, 292, 297, 7, 25, 2, 2, 293, 294, 5, 86, 44, 2, 294, 295,
	5, 86, 44, 2, 295, 296, 7, 55, 2, 2, 296, 298, 3, 2, 2, 2, 297, 293, 3,
	2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2,
	2, 300, 43, 3, 2, 2, 2, 301, 306, 7, 26, 2, 2, 302, 303, 5, 86, 44, 2,
	303, 304, 5, 86, 44, 2, 304, 305, 7, 55, 2, 2, 305, 307, 3, 2, 2, 2, 306,
	302, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 309,
	3, 2, 2, 2, 309, 45, 3, 2, 2, 2, 310, 315, 7, 27, 2, 2, 311, 312, 5, 86,
	44, 2, 312, 313, 5, 86, 44, 2, 313, 314, 7, 55, 2, 2, 314, 316, 3, 2, 2,
	2, 315, 311, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 317,
	318, 3, 2, 2, 2, 318, 47, 3, 2, 2, 2, 319, 321, 7, 28, 2, 2, 320, 322,
	7, 55, 2, 2, 321, 320, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 321, 3, 2,
	2, 2, 323, 324, 3, 2, 2, 2, 324, 49, 3, 2, 2, 2, 325, 326, 7, 29, 2, 2,
	326, 51, 3, 2, 2, 2, 327, 328, 7, 30, 2, 2, 328, 329, 7, 55, 2, 2, 329,
	330, 5, 86, 44, 2, 330, 331, 5, 86, 44, 2, 331, 332, 7, 55, 2, 2, 332,
	333, 7, 55, 2, 2, 333, 334, 5, 86, 44, 2, 334, 335, 5, 86, 44, 2, 335,
	336, 7, 55, 2, 2, 336, 53, 3, 2, 2, 2, 337, 339, 7, 31, 2, 2, 338, 340,
	7, 58, 2, 2, 339, 338, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 339, 3, 2,
	2, 2, 341, 342, 3, 2, 2, 2, 342, 55, 3, 2, 2, 2, 343, 346, 7, 32, 2, 2,
	344, 347, 7, 55, 2, 2, 345, 347, 7, 33, 2, 2, 346, 344, 3, 2, 2, 2, 346,
	345, 3, 2, 2, 2, 347, 57, 3, 2, 2, 2, 348, 352, 7, 34, 2, 2, 349, 350,
	7, 55, 2, 2, 350, 353, 5, 86, 44, 2, 351, 353, 7, 33, 2, 2, 352, 349, 3,
	2, 2, 2, 352, 351, 3, 2, 2, 2, 353, 59, 3, 2, 2, 2, 354, 355, 7, 35, 2,
	2, 355, 356, 7, 58, 2, 2, 356, 61, 3, 2, 2, 2, 357, 358, 7, 36, 2, 2, 358,
	359, 9, 6, 2, 2, 359, 63, 3, 2, 2, 2, 360, 361, 7, 38, 2, 2, 361, 362,
	9, 6, 2, 2, 362, 65, 3, 2, 2, 2, 363, 364, 7, 39, 2, 2, 364, 365, 9, 6,
	2, 2, 365, 67, 3, 2, 2, 2, 366, 367, 7, 40, 2, 2, 367, 368, 7, 55, 2, 2,
	368, 69, 3, 2, 2, 2, 369, 371, 7, 41, 2, 2, 370, 372, 7, 59, 2, 2, 371,
	370, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 373, 374,
	3, 2, 2, 2, 374, 71, 3, 2, 2, 2, 375, 378, 7, 42, 2, 2, 376, 379, 7, 58,
	2, 2, 377, 379, 7, 33, 2, 2, 378, 376, 3, 2, 2, 2, 378, 377, 3, 2, 2, 2,
	379, 73, 3, 2, 2, 2, 380, 381, 7, 43, 2, 2, 381, 382, 7, 59, 2, 2, 382,
	75, 3, 2, 2, 2, 383, 384, 7, 44, 2, 2, 384, 385, 7, 58, 2, 2, 385, 77,
	3, 2, 2, 2, 386, 387, 7, 45, 2, 2, 387, 388, 7, 59, 2, 2, 388, 79, 3, 2,
	2, 2, 389, 390, 7, 46, 2, 2, 390, 391, 7, 59, 2, 2, 391, 81, 3, 2, 2, 2,
	392, 401, 7, 47, 2, 2, 393, 394, 7, 48, 2, 2, 394, 402, 5, 86, 44, 2, 395,
	396, 7, 49, 2, 2, 396, 402, 5, 86, 44, 2, 397, 398, 7, 21, 2, 2, 398, 399,
	5, 86, 44, 2, 399, 400, 5, 86, 44, 2, 400, 402, 3, 2, 2, 2, 401, 393, 3,
	2, 2, 2, 401, 395, 3, 2, 2, 2, 401, 397, 3, 2, 2, 2, 402, 83, 3, 2, 2,
	2, 403, 416, 7, 50, 2, 2, 404, 405, 7, 51, 2, 2, 405, 406, 5, 86, 44, 2,
	406, 407, 5, 86, 44, 2, 407, 417, 3, 2, 2, 2, 408, 409, 7, 52, 2, 2, 409,
	417, 5, 86, 44, 2, 410, 411, 7, 49, 2, 2, 411, 417, 5, 86, 44, 2, 412,
	413, 7, 21, 2, 2, 413, 414, 5, 86, 44, 2, 414, 415, 5, 86, 44, 2, 415,
	417, 3, 2, 2, 2, 416, 404, 3, 2, 2, 2, 416, 408, 3, 2, 2, 2, 416, 410,
	3, 2, 2, 2, 416, 412, 3, 2, 2, 2, 417, 85, 3, 2, 2, 2, 418, 419, 9, 7,
	2, 2, 419, 87, 3, 2, 2, 2, 42, 91, 99, 104, 110, 144, 150, 155, 160, 168,
	173, 176, 186, 189, 193, 200, 207, 211, 217, 223, 229, 234, 239, 246, 251,
	255, 266, 273, 283, 290, 299, 308, 317, 323, 341, 346, 352, 373, 378, 401,
	416,
}
var literalNames = []string{
	"", "'-'", "'v'", "'vp'", "'vn'", "'vt'", "'cstype'", "'rat'", "'bmatrix'",
	"'bezier'", "'bspline'", "'cardinal'", "'taylor'", "'deg'", "'bmat'", "'u'",
	"'p'", "'l'", "'f'", "'curv'", "'curv2'", "'surf'", "'parm'", "'trim'",
	"'hole'", "'scrv'", "'sp'", "'end'", "'con'", "'g'", "'s'", "'off'", "'mg'",
	"'o'", "'bevel'", "'on'", "'c_interp'", "'d_interp'", "'lod'", "'maplib'",
	"'usemap'", "'mtllib'", "'usemtl'", "'shadow_obj'", "'trace_obj'", "'ctech'",
	"'cparm'", "'cspace'", "'stech'", "'cparma'", "'cparmb'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "INTEGER_PAIR",
	"INTEGER_TRIPLET", "INTEGER", "DECIMAL", "COMMENT", "NAME", "FILENAME",
	"WS", "NL", "NON_NL", "NON_WS",
}

var ruleNames = []string{
	"start", "statement", "call", "csh", "vertex", "vertex_parameter", "vertex_normal",
	"vertex_texture", "curve_surface_type", "degree", "basis_matrix", "step",
	"points", "lines", "faces", "free_form_surface", "curve", "curve2d", "surface",
	"parameter", "outer_trimming_loop", "inner_trimming_loop", "special_curve",
	"special_point", "end", "connectivity", "group", "smoothing_group", "merging_group",
	"object_name", "bevel", "color_interpolation", "dissolve_interpolation",
	"level_of_detail", "map_library", "use_map", "material_library", "use_material",
	"shadow_object", "trace_object", "curve_approximation_technique", "surface_approximation_technique",
	"decimal",
}

type WavefrontOBJParser struct {
	*antlr.BaseParser
}

// NewWavefrontOBJParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *WavefrontOBJParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewWavefrontOBJParser(input antlr.TokenStream) *WavefrontOBJParser {
	this := new(WavefrontOBJParser)
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
	this.GrammarFileName = "WavefrontOBJ.g4"

	return this
}

// WavefrontOBJParser tokens.
const (
	WavefrontOBJParserEOF             = antlr.TokenEOF
	WavefrontOBJParserT__0            = 1
	WavefrontOBJParserT__1            = 2
	WavefrontOBJParserT__2            = 3
	WavefrontOBJParserT__3            = 4
	WavefrontOBJParserT__4            = 5
	WavefrontOBJParserT__5            = 6
	WavefrontOBJParserT__6            = 7
	WavefrontOBJParserT__7            = 8
	WavefrontOBJParserT__8            = 9
	WavefrontOBJParserT__9            = 10
	WavefrontOBJParserT__10           = 11
	WavefrontOBJParserT__11           = 12
	WavefrontOBJParserT__12           = 13
	WavefrontOBJParserT__13           = 14
	WavefrontOBJParserT__14           = 15
	WavefrontOBJParserT__15           = 16
	WavefrontOBJParserT__16           = 17
	WavefrontOBJParserT__17           = 18
	WavefrontOBJParserT__18           = 19
	WavefrontOBJParserT__19           = 20
	WavefrontOBJParserT__20           = 21
	WavefrontOBJParserT__21           = 22
	WavefrontOBJParserT__22           = 23
	WavefrontOBJParserT__23           = 24
	WavefrontOBJParserT__24           = 25
	WavefrontOBJParserT__25           = 26
	WavefrontOBJParserT__26           = 27
	WavefrontOBJParserT__27           = 28
	WavefrontOBJParserT__28           = 29
	WavefrontOBJParserT__29           = 30
	WavefrontOBJParserT__30           = 31
	WavefrontOBJParserT__31           = 32
	WavefrontOBJParserT__32           = 33
	WavefrontOBJParserT__33           = 34
	WavefrontOBJParserT__34           = 35
	WavefrontOBJParserT__35           = 36
	WavefrontOBJParserT__36           = 37
	WavefrontOBJParserT__37           = 38
	WavefrontOBJParserT__38           = 39
	WavefrontOBJParserT__39           = 40
	WavefrontOBJParserT__40           = 41
	WavefrontOBJParserT__41           = 42
	WavefrontOBJParserT__42           = 43
	WavefrontOBJParserT__43           = 44
	WavefrontOBJParserT__44           = 45
	WavefrontOBJParserT__45           = 46
	WavefrontOBJParserT__46           = 47
	WavefrontOBJParserT__47           = 48
	WavefrontOBJParserT__48           = 49
	WavefrontOBJParserT__49           = 50
	WavefrontOBJParserINTEGER_PAIR    = 51
	WavefrontOBJParserINTEGER_TRIPLET = 52
	WavefrontOBJParserINTEGER         = 53
	WavefrontOBJParserDECIMAL         = 54
	WavefrontOBJParserCOMMENT         = 55
	WavefrontOBJParserNAME            = 56
	WavefrontOBJParserFILENAME        = 57
	WavefrontOBJParserWS              = 58
	WavefrontOBJParserNL              = 59
	WavefrontOBJParserNON_NL          = 60
	WavefrontOBJParserNON_WS          = 61
)

// WavefrontOBJParser rules.
const (
	WavefrontOBJParserRULE_start                           = 0
	WavefrontOBJParserRULE_statement                       = 1
	WavefrontOBJParserRULE_call                            = 2
	WavefrontOBJParserRULE_csh                             = 3
	WavefrontOBJParserRULE_vertex                          = 4
	WavefrontOBJParserRULE_vertex_parameter                = 5
	WavefrontOBJParserRULE_vertex_normal                   = 6
	WavefrontOBJParserRULE_vertex_texture                  = 7
	WavefrontOBJParserRULE_curve_surface_type              = 8
	WavefrontOBJParserRULE_degree                          = 9
	WavefrontOBJParserRULE_basis_matrix                    = 10
	WavefrontOBJParserRULE_step                            = 11
	WavefrontOBJParserRULE_points                          = 12
	WavefrontOBJParserRULE_lines                           = 13
	WavefrontOBJParserRULE_faces                           = 14
	WavefrontOBJParserRULE_free_form_surface               = 15
	WavefrontOBJParserRULE_curve                           = 16
	WavefrontOBJParserRULE_curve2d                         = 17
	WavefrontOBJParserRULE_surface                         = 18
	WavefrontOBJParserRULE_parameter                       = 19
	WavefrontOBJParserRULE_outer_trimming_loop             = 20
	WavefrontOBJParserRULE_inner_trimming_loop             = 21
	WavefrontOBJParserRULE_special_curve                   = 22
	WavefrontOBJParserRULE_special_point                   = 23
	WavefrontOBJParserRULE_end                             = 24
	WavefrontOBJParserRULE_connectivity                    = 25
	WavefrontOBJParserRULE_group                           = 26
	WavefrontOBJParserRULE_smoothing_group                 = 27
	WavefrontOBJParserRULE_merging_group                   = 28
	WavefrontOBJParserRULE_object_name                     = 29
	WavefrontOBJParserRULE_bevel                           = 30
	WavefrontOBJParserRULE_color_interpolation             = 31
	WavefrontOBJParserRULE_dissolve_interpolation          = 32
	WavefrontOBJParserRULE_level_of_detail                 = 33
	WavefrontOBJParserRULE_map_library                     = 34
	WavefrontOBJParserRULE_use_map                         = 35
	WavefrontOBJParserRULE_material_library                = 36
	WavefrontOBJParserRULE_use_material                    = 37
	WavefrontOBJParserRULE_shadow_object                   = 38
	WavefrontOBJParserRULE_trace_object                    = 39
	WavefrontOBJParserRULE_curve_approximation_technique   = 40
	WavefrontOBJParserRULE_surface_approximation_technique = 41
	WavefrontOBJParserRULE_decimal                         = 42
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StartContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StartContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserNL)
}

func (s *StartContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNL, i)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *WavefrontOBJParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WavefrontOBJParserRULE_start)
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == WavefrontOBJParserNL {
		{
			p.SetState(86)
			p.Match(WavefrontOBJParserNL)
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Statement()
	}

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == WavefrontOBJParserNL {
				{
					p.SetState(94)
					p.Match(WavefrontOBJParserNL)
				}

				p.SetState(97)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(99)
				p.Statement()
			}

		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == WavefrontOBJParserNL {
		{
			p.SetState(105)
			p.Match(WavefrontOBJParserNL)
		}

		p.SetState(110)
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
	p.RuleIndex = WavefrontOBJParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Call() ICallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *StatementContext) Csh() ICshContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICshContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICshContext)
}

func (s *StatementContext) Vertex() IVertexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertexContext)
}

func (s *StatementContext) Vertex_normal() IVertex_normalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertex_normalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertex_normalContext)
}

func (s *StatementContext) Vertex_texture() IVertex_textureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertex_textureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertex_textureContext)
}

func (s *StatementContext) Vertex_parameter() IVertex_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertex_parameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertex_parameterContext)
}

func (s *StatementContext) Points() IPointsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointsContext)
}

func (s *StatementContext) Lines() ILinesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinesContext)
}

func (s *StatementContext) Faces() IFacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFacesContext)
}

func (s *StatementContext) Curve_surface_type() ICurve_surface_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurve_surface_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurve_surface_typeContext)
}

func (s *StatementContext) Degree() IDegreeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDegreeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDegreeContext)
}

func (s *StatementContext) Basis_matrix() IBasis_matrixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasis_matrixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasis_matrixContext)
}

func (s *StatementContext) Step() IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *StatementContext) Free_form_surface() IFree_form_surfaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFree_form_surfaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFree_form_surfaceContext)
}

func (s *StatementContext) Connectivity() IConnectivityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectivityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectivityContext)
}

func (s *StatementContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *StatementContext) Smoothing_group() ISmoothing_groupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISmoothing_groupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISmoothing_groupContext)
}

func (s *StatementContext) Merging_group() IMerging_groupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMerging_groupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMerging_groupContext)
}

func (s *StatementContext) Object_name() IObject_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_nameContext)
}

func (s *StatementContext) Bevel() IBevelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBevelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBevelContext)
}

func (s *StatementContext) Color_interpolation() IColor_interpolationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColor_interpolationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColor_interpolationContext)
}

func (s *StatementContext) Dissolve_interpolation() IDissolve_interpolationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDissolve_interpolationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDissolve_interpolationContext)
}

func (s *StatementContext) Level_of_detail() ILevel_of_detailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILevel_of_detailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILevel_of_detailContext)
}

func (s *StatementContext) Map_library() IMap_libraryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMap_libraryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMap_libraryContext)
}

func (s *StatementContext) Use_map() IUse_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUse_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUse_mapContext)
}

func (s *StatementContext) Material_library() IMaterial_libraryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaterial_libraryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMaterial_libraryContext)
}

func (s *StatementContext) Use_material() IUse_materialContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUse_materialContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUse_materialContext)
}

func (s *StatementContext) Shadow_object() IShadow_objectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShadow_objectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShadow_objectContext)
}

func (s *StatementContext) Trace_object() ITrace_objectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrace_objectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrace_objectContext)
}

func (s *StatementContext) Curve_approximation_technique() ICurve_approximation_techniqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurve_approximation_techniqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurve_approximation_techniqueContext)
}

func (s *StatementContext) Surface_approximation_technique() ISurface_approximation_techniqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISurface_approximation_techniqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISurface_approximation_techniqueContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *WavefrontOBJParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WavefrontOBJParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Csh()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.Vertex()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.Vertex_normal()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(115)
			p.Vertex_texture()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(116)
			p.Vertex_parameter()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(117)
			p.Points()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(118)
			p.Lines()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(119)
			p.Faces()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(120)
			p.Curve_surface_type()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(121)
			p.Degree()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(122)
			p.Basis_matrix()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(123)
			p.Step()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(124)
			p.Free_form_surface()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(125)
			p.Connectivity()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(126)
			p.Group()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(127)
			p.Smoothing_group()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(128)
			p.Merging_group()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(129)
			p.Object_name()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(130)
			p.Bevel()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(131)
			p.Color_interpolation()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(132)
			p.Dissolve_interpolation()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(133)
			p.Level_of_detail()
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(134)
			p.Map_library()
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(135)
			p.Use_map()
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(136)
			p.Material_library()
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(137)
			p.Use_material()
		}

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(138)
			p.Shadow_object()
		}

	case 29:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(139)
			p.Trace_object()
		}

	case 30:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(140)
			p.Curve_approximation_technique()
		}

	case 31:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(141)
			p.Surface_approximation_technique()
		}

	}

	return localctx
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFilename returns the filename token.
	GetFilename() antlr.Token

	// GetArgs returns the args token.
	GetArgs() antlr.Token

	// SetFilename sets the filename token.
	SetFilename(antlr.Token)

	// SetArgs sets the args token.
	SetArgs(antlr.Token)

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	filename antlr.Token
	args     antlr.Token
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) GetFilename() antlr.Token { return s.filename }

func (s *CallContext) GetArgs() antlr.Token { return s.args }

func (s *CallContext) SetFilename(v antlr.Token) { s.filename = v }

func (s *CallContext) SetArgs(v antlr.Token) { s.args = v }

func (s *CallContext) FILENAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserFILENAME, 0)
}

func (s *CallContext) AllNON_WS() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserNON_WS)
}

func (s *CallContext) NON_WS(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNON_WS, i)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitCall(s)
	}
}

func (p *WavefrontOBJParser) Call() (localctx ICallContext) {
	this := p
	_ = this

	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WavefrontOBJParserRULE_call)
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
		p.SetState(144)

		var _m = p.Match(WavefrontOBJParserFILENAME)

		localctx.(*CallContext).filename = _m
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == WavefrontOBJParserNON_WS {
		{
			p.SetState(145)

			var _m = p.Match(WavefrontOBJParserNON_WS)

			localctx.(*CallContext).args = _m
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICshContext is an interface to support dynamic dispatch.
type ICshContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCommand returns the command token.
	GetCommand() antlr.Token

	// GetArgs returns the args token.
	GetArgs() antlr.Token

	// SetCommand sets the command token.
	SetCommand(antlr.Token)

	// SetArgs sets the args token.
	SetArgs(antlr.Token)

	// IsCshContext differentiates from other interfaces.
	IsCshContext()
}

type CshContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	command antlr.Token
	args    antlr.Token
}

func NewEmptyCshContext() *CshContext {
	var p = new(CshContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_csh
	return p
}

func (*CshContext) IsCshContext() {}

func NewCshContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CshContext {
	var p = new(CshContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_csh

	return p
}

func (s *CshContext) GetParser() antlr.Parser { return s.parser }

func (s *CshContext) GetCommand() antlr.Token { return s.command }

func (s *CshContext) GetArgs() antlr.Token { return s.args }

func (s *CshContext) SetCommand(v antlr.Token) { s.command = v }

func (s *CshContext) SetArgs(v antlr.Token) { s.args = v }

func (s *CshContext) FILENAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserFILENAME, 0)
}

func (s *CshContext) AllNON_WS() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserNON_WS)
}

func (s *CshContext) NON_WS(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNON_WS, i)
}

func (s *CshContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CshContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CshContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterCsh(s)
	}
}

func (s *CshContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitCsh(s)
	}
}

func (p *WavefrontOBJParser) Csh() (localctx ICshContext) {
	this := p
	_ = this

	localctx = NewCshContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, WavefrontOBJParserRULE_csh)
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
		p.SetState(151)

		var _m = p.Match(WavefrontOBJParserFILENAME)

		localctx.(*CshContext).command = _m
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WavefrontOBJParserT__0 {
		{
			p.SetState(152)
			p.Match(WavefrontOBJParserT__0)
		}

	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == WavefrontOBJParserNON_WS {
		{
			p.SetState(155)

			var _m = p.Match(WavefrontOBJParserNON_WS)

			localctx.(*CshContext).args = _m
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVertexContext is an interface to support dynamic dispatch.
type IVertexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetX returns the x rule contexts.
	GetX() IDecimalContext

	// GetY returns the y rule contexts.
	GetY() IDecimalContext

	// GetZ returns the z rule contexts.
	GetZ() IDecimalContext

	// GetW returns the w rule contexts.
	GetW() IDecimalContext

	// SetX sets the x rule contexts.
	SetX(IDecimalContext)

	// SetY sets the y rule contexts.
	SetY(IDecimalContext)

	// SetZ sets the z rule contexts.
	SetZ(IDecimalContext)

	// SetW sets the w rule contexts.
	SetW(IDecimalContext)

	// IsVertexContext differentiates from other interfaces.
	IsVertexContext()
}

type VertexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	x      IDecimalContext
	y      IDecimalContext
	z      IDecimalContext
	w      IDecimalContext
}

func NewEmptyVertexContext() *VertexContext {
	var p = new(VertexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_vertex
	return p
}

func (*VertexContext) IsVertexContext() {}

func NewVertexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VertexContext {
	var p = new(VertexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_vertex

	return p
}

func (s *VertexContext) GetParser() antlr.Parser { return s.parser }

func (s *VertexContext) GetX() IDecimalContext { return s.x }

func (s *VertexContext) GetY() IDecimalContext { return s.y }

func (s *VertexContext) GetZ() IDecimalContext { return s.z }

func (s *VertexContext) GetW() IDecimalContext { return s.w }

func (s *VertexContext) SetX(v IDecimalContext) { s.x = v }

func (s *VertexContext) SetY(v IDecimalContext) { s.y = v }

func (s *VertexContext) SetZ(v IDecimalContext) { s.z = v }

func (s *VertexContext) SetW(v IDecimalContext) { s.w = v }

func (s *VertexContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *VertexContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *VertexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VertexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VertexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterVertex(s)
	}
}

func (s *VertexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitVertex(s)
	}
}

func (p *WavefrontOBJParser) Vertex() (localctx IVertexContext) {
	this := p
	_ = this

	localctx = NewVertexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, WavefrontOBJParserRULE_vertex)
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
		p.SetState(161)
		p.Match(WavefrontOBJParserT__1)
	}
	{
		p.SetState(162)

		var _x = p.Decimal()

		localctx.(*VertexContext).x = _x
	}
	{
		p.SetState(163)

		var _x = p.Decimal()

		localctx.(*VertexContext).y = _x
	}
	{
		p.SetState(164)

		var _x = p.Decimal()

		localctx.(*VertexContext).z = _x
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WavefrontOBJParserINTEGER || _la == WavefrontOBJParserDECIMAL {
		{
			p.SetState(165)

			var _x = p.Decimal()

			localctx.(*VertexContext).w = _x
		}

	}

	return localctx
}

// IVertex_parameterContext is an interface to support dynamic dispatch.
type IVertex_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetU returns the u rule contexts.
	GetU() IDecimalContext

	// GetV returns the v rule contexts.
	GetV() IDecimalContext

	// GetW returns the w rule contexts.
	GetW() IDecimalContext

	// SetU sets the u rule contexts.
	SetU(IDecimalContext)

	// SetV sets the v rule contexts.
	SetV(IDecimalContext)

	// SetW sets the w rule contexts.
	SetW(IDecimalContext)

	// IsVertex_parameterContext differentiates from other interfaces.
	IsVertex_parameterContext()
}

type Vertex_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	u      IDecimalContext
	v      IDecimalContext
	w      IDecimalContext
}

func NewEmptyVertex_parameterContext() *Vertex_parameterContext {
	var p = new(Vertex_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_vertex_parameter
	return p
}

func (*Vertex_parameterContext) IsVertex_parameterContext() {}

func NewVertex_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vertex_parameterContext {
	var p = new(Vertex_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_vertex_parameter

	return p
}

func (s *Vertex_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Vertex_parameterContext) GetU() IDecimalContext { return s.u }

func (s *Vertex_parameterContext) GetV() IDecimalContext { return s.v }

func (s *Vertex_parameterContext) GetW() IDecimalContext { return s.w }

func (s *Vertex_parameterContext) SetU(v IDecimalContext) { s.u = v }

func (s *Vertex_parameterContext) SetV(v IDecimalContext) { s.v = v }

func (s *Vertex_parameterContext) SetW(v IDecimalContext) { s.w = v }

func (s *Vertex_parameterContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *Vertex_parameterContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Vertex_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vertex_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vertex_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterVertex_parameter(s)
	}
}

func (s *Vertex_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitVertex_parameter(s)
	}
}

func (p *WavefrontOBJParser) Vertex_parameter() (localctx IVertex_parameterContext) {
	this := p
	_ = this

	localctx = NewVertex_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, WavefrontOBJParserRULE_vertex_parameter)
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
		p.SetState(168)
		p.Match(WavefrontOBJParserT__2)
	}
	{
		p.SetState(169)

		var _x = p.Decimal()

		localctx.(*Vertex_parameterContext).u = _x
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(170)

			var _x = p.Decimal()

			localctx.(*Vertex_parameterContext).v = _x
		}

	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WavefrontOBJParserINTEGER || _la == WavefrontOBJParserDECIMAL {
		{
			p.SetState(173)

			var _x = p.Decimal()

			localctx.(*Vertex_parameterContext).w = _x
		}

	}

	return localctx
}

// IVertex_normalContext is an interface to support dynamic dispatch.
type IVertex_normalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetX returns the x rule contexts.
	GetX() IDecimalContext

	// GetY returns the y rule contexts.
	GetY() IDecimalContext

	// GetZ returns the z rule contexts.
	GetZ() IDecimalContext

	// SetX sets the x rule contexts.
	SetX(IDecimalContext)

	// SetY sets the y rule contexts.
	SetY(IDecimalContext)

	// SetZ sets the z rule contexts.
	SetZ(IDecimalContext)

	// IsVertex_normalContext differentiates from other interfaces.
	IsVertex_normalContext()
}

type Vertex_normalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	x      IDecimalContext
	y      IDecimalContext
	z      IDecimalContext
}

func NewEmptyVertex_normalContext() *Vertex_normalContext {
	var p = new(Vertex_normalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_vertex_normal
	return p
}

func (*Vertex_normalContext) IsVertex_normalContext() {}

func NewVertex_normalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vertex_normalContext {
	var p = new(Vertex_normalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_vertex_normal

	return p
}

func (s *Vertex_normalContext) GetParser() antlr.Parser { return s.parser }

func (s *Vertex_normalContext) GetX() IDecimalContext { return s.x }

func (s *Vertex_normalContext) GetY() IDecimalContext { return s.y }

func (s *Vertex_normalContext) GetZ() IDecimalContext { return s.z }

func (s *Vertex_normalContext) SetX(v IDecimalContext) { s.x = v }

func (s *Vertex_normalContext) SetY(v IDecimalContext) { s.y = v }

func (s *Vertex_normalContext) SetZ(v IDecimalContext) { s.z = v }

func (s *Vertex_normalContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *Vertex_normalContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Vertex_normalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vertex_normalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vertex_normalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterVertex_normal(s)
	}
}

func (s *Vertex_normalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitVertex_normal(s)
	}
}

func (p *WavefrontOBJParser) Vertex_normal() (localctx IVertex_normalContext) {
	this := p
	_ = this

	localctx = NewVertex_normalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, WavefrontOBJParserRULE_vertex_normal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(WavefrontOBJParserT__3)
	}
	{
		p.SetState(177)

		var _x = p.Decimal()

		localctx.(*Vertex_normalContext).x = _x
	}
	{
		p.SetState(178)

		var _x = p.Decimal()

		localctx.(*Vertex_normalContext).y = _x
	}
	{
		p.SetState(179)

		var _x = p.Decimal()

		localctx.(*Vertex_normalContext).z = _x
	}

	return localctx
}

// IVertex_textureContext is an interface to support dynamic dispatch.
type IVertex_textureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetU returns the u rule contexts.
	GetU() IDecimalContext

	// GetV returns the v rule contexts.
	GetV() IDecimalContext

	// GetW returns the w rule contexts.
	GetW() IDecimalContext

	// SetU sets the u rule contexts.
	SetU(IDecimalContext)

	// SetV sets the v rule contexts.
	SetV(IDecimalContext)

	// SetW sets the w rule contexts.
	SetW(IDecimalContext)

	// IsVertex_textureContext differentiates from other interfaces.
	IsVertex_textureContext()
}

type Vertex_textureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	u      IDecimalContext
	v      IDecimalContext
	w      IDecimalContext
}

func NewEmptyVertex_textureContext() *Vertex_textureContext {
	var p = new(Vertex_textureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_vertex_texture
	return p
}

func (*Vertex_textureContext) IsVertex_textureContext() {}

func NewVertex_textureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vertex_textureContext {
	var p = new(Vertex_textureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_vertex_texture

	return p
}

func (s *Vertex_textureContext) GetParser() antlr.Parser { return s.parser }

func (s *Vertex_textureContext) GetU() IDecimalContext { return s.u }

func (s *Vertex_textureContext) GetV() IDecimalContext { return s.v }

func (s *Vertex_textureContext) GetW() IDecimalContext { return s.w }

func (s *Vertex_textureContext) SetU(v IDecimalContext) { s.u = v }

func (s *Vertex_textureContext) SetV(v IDecimalContext) { s.v = v }

func (s *Vertex_textureContext) SetW(v IDecimalContext) { s.w = v }

func (s *Vertex_textureContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *Vertex_textureContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Vertex_textureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vertex_textureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vertex_textureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterVertex_texture(s)
	}
}

func (s *Vertex_textureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitVertex_texture(s)
	}
}

func (p *WavefrontOBJParser) Vertex_texture() (localctx IVertex_textureContext) {
	this := p
	_ = this

	localctx = NewVertex_textureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, WavefrontOBJParserRULE_vertex_texture)
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
		p.SetState(181)
		p.Match(WavefrontOBJParserT__4)
	}
	{
		p.SetState(182)

		var _x = p.Decimal()

		localctx.(*Vertex_textureContext).u = _x
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(183)

			var _x = p.Decimal()

			localctx.(*Vertex_textureContext).v = _x
		}

	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WavefrontOBJParserINTEGER || _la == WavefrontOBJParserDECIMAL {
		{
			p.SetState(186)

			var _x = p.Decimal()

			localctx.(*Vertex_textureContext).w = _x
		}

	}

	return localctx
}

// ICurve_surface_typeContext is an interface to support dynamic dispatch.
type ICurve_surface_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRational returns the rational token.
	GetRational() antlr.Token

	// GetCstype returns the cstype token.
	GetCstype() antlr.Token

	// SetRational sets the rational token.
	SetRational(antlr.Token)

	// SetCstype sets the cstype token.
	SetCstype(antlr.Token)

	// IsCurve_surface_typeContext differentiates from other interfaces.
	IsCurve_surface_typeContext()
}

type Curve_surface_typeContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	rational antlr.Token
	cstype   antlr.Token
}

func NewEmptyCurve_surface_typeContext() *Curve_surface_typeContext {
	var p = new(Curve_surface_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_curve_surface_type
	return p
}

func (*Curve_surface_typeContext) IsCurve_surface_typeContext() {}

func NewCurve_surface_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Curve_surface_typeContext {
	var p = new(Curve_surface_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_curve_surface_type

	return p
}

func (s *Curve_surface_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Curve_surface_typeContext) GetRational() antlr.Token { return s.rational }

func (s *Curve_surface_typeContext) GetCstype() antlr.Token { return s.cstype }

func (s *Curve_surface_typeContext) SetRational(v antlr.Token) { s.rational = v }

func (s *Curve_surface_typeContext) SetCstype(v antlr.Token) { s.cstype = v }

func (s *Curve_surface_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Curve_surface_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Curve_surface_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterCurve_surface_type(s)
	}
}

func (s *Curve_surface_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitCurve_surface_type(s)
	}
}

func (p *WavefrontOBJParser) Curve_surface_type() (localctx ICurve_surface_typeContext) {
	this := p
	_ = this

	localctx = NewCurve_surface_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, WavefrontOBJParserRULE_curve_surface_type)
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
		p.SetState(189)
		p.Match(WavefrontOBJParserT__5)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WavefrontOBJParserT__6 {
		{
			p.SetState(190)

			var _m = p.Match(WavefrontOBJParserT__6)

			localctx.(*Curve_surface_typeContext).rational = _m
		}

	}
	{
		p.SetState(193)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Curve_surface_typeContext).cstype = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<WavefrontOBJParserT__7)|(1<<WavefrontOBJParserT__8)|(1<<WavefrontOBJParserT__9)|(1<<WavefrontOBJParserT__10)|(1<<WavefrontOBJParserT__11))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Curve_surface_typeContext).cstype = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDegreeContext is an interface to support dynamic dispatch.
type IDegreeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetU returns the u token.
	GetU() antlr.Token

	// GetV returns the v token.
	GetV() antlr.Token

	// SetU sets the u token.
	SetU(antlr.Token)

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsDegreeContext differentiates from other interfaces.
	IsDegreeContext()
}

type DegreeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	u      antlr.Token
	v      antlr.Token
}

func NewEmptyDegreeContext() *DegreeContext {
	var p = new(DegreeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_degree
	return p
}

func (*DegreeContext) IsDegreeContext() {}

func NewDegreeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DegreeContext {
	var p = new(DegreeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_degree

	return p
}

func (s *DegreeContext) GetParser() antlr.Parser { return s.parser }

func (s *DegreeContext) GetU() antlr.Token { return s.u }

func (s *DegreeContext) GetV() antlr.Token { return s.v }

func (s *DegreeContext) SetU(v antlr.Token) { s.u = v }

func (s *DegreeContext) SetV(v antlr.Token) { s.v = v }

func (s *DegreeContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *DegreeContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *DegreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DegreeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DegreeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterDegree(s)
	}
}

func (s *DegreeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitDegree(s)
	}
}

func (p *WavefrontOBJParser) Degree() (localctx IDegreeContext) {
	this := p
	_ = this

	localctx = NewDegreeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, WavefrontOBJParserRULE_degree)
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
		p.Match(WavefrontOBJParserT__12)
	}
	{
		p.SetState(196)

		var _m = p.Match(WavefrontOBJParserINTEGER)

		localctx.(*DegreeContext).u = _m
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WavefrontOBJParserINTEGER {
		{
			p.SetState(197)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*DegreeContext).v = _m
		}

	}

	return localctx
}

// IBasis_matrixContext is an interface to support dynamic dispatch.
type IBasis_matrixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasis_matrixContext differentiates from other interfaces.
	IsBasis_matrixContext()
}

type Basis_matrixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasis_matrixContext() *Basis_matrixContext {
	var p = new(Basis_matrixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_basis_matrix
	return p
}

func (*Basis_matrixContext) IsBasis_matrixContext() {}

func NewBasis_matrixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Basis_matrixContext {
	var p = new(Basis_matrixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_basis_matrix

	return p
}

func (s *Basis_matrixContext) GetParser() antlr.Parser { return s.parser }

func (s *Basis_matrixContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *Basis_matrixContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Basis_matrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Basis_matrixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Basis_matrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterBasis_matrix(s)
	}
}

func (s *Basis_matrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitBasis_matrix(s)
	}
}

func (p *WavefrontOBJParser) Basis_matrix() (localctx IBasis_matrixContext) {
	this := p
	_ = this

	localctx = NewBasis_matrixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, WavefrontOBJParserRULE_basis_matrix)
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
		p.SetState(200)
		p.Match(WavefrontOBJParserT__13)
	}
	{
		p.SetState(201)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WavefrontOBJParserT__1 || _la == WavefrontOBJParserT__14) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER || _la == WavefrontOBJParserDECIMAL {
		{
			p.SetState(202)
			p.Decimal()
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStepContext is an interface to support dynamic dispatch.
type IStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetU returns the u token.
	GetU() antlr.Token

	// GetV returns the v token.
	GetV() antlr.Token

	// SetU sets the u token.
	SetU(antlr.Token)

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsStepContext differentiates from other interfaces.
	IsStepContext()
}

type StepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	u      antlr.Token
	v      antlr.Token
}

func NewEmptyStepContext() *StepContext {
	var p = new(StepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_step
	return p
}

func (*StepContext) IsStepContext() {}

func NewStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepContext {
	var p = new(StepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_step

	return p
}

func (s *StepContext) GetParser() antlr.Parser { return s.parser }

func (s *StepContext) GetU() antlr.Token { return s.u }

func (s *StepContext) GetV() antlr.Token { return s.v }

func (s *StepContext) SetU(v antlr.Token) { s.u = v }

func (s *StepContext) SetV(v antlr.Token) { s.v = v }

func (s *StepContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *StepContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *StepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterStep(s)
	}
}

func (s *StepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitStep(s)
	}
}

func (p *WavefrontOBJParser) Step() (localctx IStepContext) {
	this := p
	_ = this

	localctx = NewStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, WavefrontOBJParserRULE_step)
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
		p.SetState(207)

		var _m = p.Match(WavefrontOBJParserINTEGER)

		localctx.(*StepContext).u = _m
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WavefrontOBJParserINTEGER {
		{
			p.SetState(208)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*StepContext).v = _m
		}

	}

	return localctx
}

// IPointsContext is an interface to support dynamic dispatch.
type IPointsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsPointsContext differentiates from other interfaces.
	IsPointsContext()
}

type PointsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyPointsContext() *PointsContext {
	var p = new(PointsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_points
	return p
}

func (*PointsContext) IsPointsContext() {}

func NewPointsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointsContext {
	var p = new(PointsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_points

	return p
}

func (s *PointsContext) GetParser() antlr.Parser { return s.parser }

func (s *PointsContext) GetV() antlr.Token { return s.v }

func (s *PointsContext) SetV(v antlr.Token) { s.v = v }

func (s *PointsContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *PointsContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *PointsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterPoints(s)
	}
}

func (s *PointsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitPoints(s)
	}
}

func (p *WavefrontOBJParser) Points() (localctx IPointsContext) {
	this := p
	_ = this

	localctx = NewPointsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, WavefrontOBJParserRULE_points)
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
		p.SetState(211)
		p.Match(WavefrontOBJParserT__15)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER {
		{
			p.SetState(212)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*PointsContext).v = _m
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILinesContext is an interface to support dynamic dispatch.
type ILinesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsLinesContext differentiates from other interfaces.
	IsLinesContext()
}

type LinesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyLinesContext() *LinesContext {
	var p = new(LinesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_lines
	return p
}

func (*LinesContext) IsLinesContext() {}

func NewLinesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinesContext {
	var p = new(LinesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_lines

	return p
}

func (s *LinesContext) GetParser() antlr.Parser { return s.parser }

func (s *LinesContext) GetV() antlr.Token { return s.v }

func (s *LinesContext) SetV(v antlr.Token) { s.v = v }

func (s *LinesContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *LinesContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *LinesContext) AllINTEGER_PAIR() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER_PAIR)
}

func (s *LinesContext) INTEGER_PAIR(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER_PAIR, i)
}

func (s *LinesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterLines(s)
	}
}

func (s *LinesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitLines(s)
	}
}

func (p *WavefrontOBJParser) Lines() (localctx ILinesContext) {
	this := p
	_ = this

	localctx = NewLinesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, WavefrontOBJParserRULE_lines)
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
		p.SetState(217)
		p.Match(WavefrontOBJParserT__16)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER_PAIR || _la == WavefrontOBJParserINTEGER {
		{
			p.SetState(218)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LinesContext).v = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == WavefrontOBJParserINTEGER_PAIR || _la == WavefrontOBJParserINTEGER) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LinesContext).v = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFacesContext is an interface to support dynamic dispatch.
type IFacesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// IsFacesContext differentiates from other interfaces.
	IsFacesContext()
}

type FacesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      antlr.Token
}

func NewEmptyFacesContext() *FacesContext {
	var p = new(FacesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_faces
	return p
}

func (*FacesContext) IsFacesContext() {}

func NewFacesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FacesContext {
	var p = new(FacesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_faces

	return p
}

func (s *FacesContext) GetParser() antlr.Parser { return s.parser }

func (s *FacesContext) GetV() antlr.Token { return s.v }

func (s *FacesContext) SetV(v antlr.Token) { s.v = v }

func (s *FacesContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *FacesContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *FacesContext) AllINTEGER_PAIR() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER_PAIR)
}

func (s *FacesContext) INTEGER_PAIR(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER_PAIR, i)
}

func (s *FacesContext) AllINTEGER_TRIPLET() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER_TRIPLET)
}

func (s *FacesContext) INTEGER_TRIPLET(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER_TRIPLET, i)
}

func (s *FacesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FacesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FacesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterFaces(s)
	}
}

func (s *FacesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitFaces(s)
	}
}

func (p *WavefrontOBJParser) Faces() (localctx IFacesContext) {
	this := p
	_ = this

	localctx = NewFacesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, WavefrontOBJParserRULE_faces)
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
		p.SetState(223)
		p.Match(WavefrontOBJParserT__17)
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(WavefrontOBJParserINTEGER_PAIR-51))|(1<<(WavefrontOBJParserINTEGER_TRIPLET-51))|(1<<(WavefrontOBJParserINTEGER-51)))) != 0) {
		{
			p.SetState(224)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*FacesContext).v = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(WavefrontOBJParserINTEGER_PAIR-51))|(1<<(WavefrontOBJParserINTEGER_TRIPLET-51))|(1<<(WavefrontOBJParserINTEGER-51)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*FacesContext).v = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFree_form_surfaceContext is an interface to support dynamic dispatch.
type IFree_form_surfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFree_form_surfaceContext differentiates from other interfaces.
	IsFree_form_surfaceContext()
}

type Free_form_surfaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFree_form_surfaceContext() *Free_form_surfaceContext {
	var p = new(Free_form_surfaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_free_form_surface
	return p
}

func (*Free_form_surfaceContext) IsFree_form_surfaceContext() {}

func NewFree_form_surfaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Free_form_surfaceContext {
	var p = new(Free_form_surfaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_free_form_surface

	return p
}

func (s *Free_form_surfaceContext) GetParser() antlr.Parser { return s.parser }

func (s *Free_form_surfaceContext) End() IEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndContext)
}

func (s *Free_form_surfaceContext) Curve() ICurveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurveContext)
}

func (s *Free_form_surfaceContext) Curve2d() ICurve2dContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurve2dContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurve2dContext)
}

func (s *Free_form_surfaceContext) Surface() ISurfaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISurfaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISurfaceContext)
}

func (s *Free_form_surfaceContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserNL)
}

func (s *Free_form_surfaceContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNL, i)
}

func (s *Free_form_surfaceContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *Free_form_surfaceContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *Free_form_surfaceContext) AllOuter_trimming_loop() []IOuter_trimming_loopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOuter_trimming_loopContext)(nil)).Elem())
	var tst = make([]IOuter_trimming_loopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOuter_trimming_loopContext)
		}
	}

	return tst
}

func (s *Free_form_surfaceContext) Outer_trimming_loop(i int) IOuter_trimming_loopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOuter_trimming_loopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOuter_trimming_loopContext)
}

func (s *Free_form_surfaceContext) AllInner_trimming_loop() []IInner_trimming_loopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInner_trimming_loopContext)(nil)).Elem())
	var tst = make([]IInner_trimming_loopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInner_trimming_loopContext)
		}
	}

	return tst
}

func (s *Free_form_surfaceContext) Inner_trimming_loop(i int) IInner_trimming_loopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInner_trimming_loopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInner_trimming_loopContext)
}

func (s *Free_form_surfaceContext) AllSpecial_curve() []ISpecial_curveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpecial_curveContext)(nil)).Elem())
	var tst = make([]ISpecial_curveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpecial_curveContext)
		}
	}

	return tst
}

func (s *Free_form_surfaceContext) Special_curve(i int) ISpecial_curveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecial_curveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpecial_curveContext)
}

func (s *Free_form_surfaceContext) AllSpecial_point() []ISpecial_pointContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpecial_pointContext)(nil)).Elem())
	var tst = make([]ISpecial_pointContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpecial_pointContext)
		}
	}

	return tst
}

func (s *Free_form_surfaceContext) Special_point(i int) ISpecial_pointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecial_pointContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpecial_pointContext)
}

func (s *Free_form_surfaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Free_form_surfaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Free_form_surfaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterFree_form_surface(s)
	}
}

func (s *Free_form_surfaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitFree_form_surface(s)
	}
}

func (p *WavefrontOBJParser) Free_form_surface() (localctx IFree_form_surfaceContext) {
	this := p
	_ = this

	localctx = NewFree_form_surfaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, WavefrontOBJParserRULE_free_form_surface)
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

	switch p.GetTokenStream().LA(1) {
	case WavefrontOBJParserT__18:
		{
			p.SetState(229)
			p.Curve()
		}

	case WavefrontOBJParserT__19:
		{
			p.SetState(230)
			p.Curve2d()
		}

	case WavefrontOBJParserT__20:
		{
			p.SetState(231)
			p.Surface()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserNL {
		{
			p.SetState(234)
			p.Match(WavefrontOBJParserNL)
		}

		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<WavefrontOBJParserT__21)|(1<<WavefrontOBJParserT__22)|(1<<WavefrontOBJParserT__23)|(1<<WavefrontOBJParserT__24)|(1<<WavefrontOBJParserT__25))) != 0 {
		p.SetState(244)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case WavefrontOBJParserT__21:
			{
				p.SetState(239)
				p.Parameter()
			}

		case WavefrontOBJParserT__22:
			{
				p.SetState(240)
				p.Outer_trimming_loop()
			}

		case WavefrontOBJParserT__23:
			{
				p.SetState(241)
				p.Inner_trimming_loop()
			}

		case WavefrontOBJParserT__24:
			{
				p.SetState(242)
				p.Special_curve()
			}

		case WavefrontOBJParserT__25:
			{
				p.SetState(243)
				p.Special_point()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == WavefrontOBJParserNL {
			{
				p.SetState(246)
				p.Match(WavefrontOBJParserNL)
			}

			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(256)
		p.End()
	}

	return localctx
}

// ICurveContext is an interface to support dynamic dispatch.
type ICurveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// GetU0 returns the u0 rule contexts.
	GetU0() IDecimalContext

	// GetU1 returns the u1 rule contexts.
	GetU1() IDecimalContext

	// SetU0 sets the u0 rule contexts.
	SetU0(IDecimalContext)

	// SetU1 sets the u1 rule contexts.
	SetU1(IDecimalContext)

	// IsCurveContext differentiates from other interfaces.
	IsCurveContext()
}

type CurveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	u0     IDecimalContext
	u1     IDecimalContext
	v      antlr.Token
}

func NewEmptyCurveContext() *CurveContext {
	var p = new(CurveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_curve
	return p
}

func (*CurveContext) IsCurveContext() {}

func NewCurveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurveContext {
	var p = new(CurveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_curve

	return p
}

func (s *CurveContext) GetParser() antlr.Parser { return s.parser }

func (s *CurveContext) GetV() antlr.Token { return s.v }

func (s *CurveContext) SetV(v antlr.Token) { s.v = v }

func (s *CurveContext) GetU0() IDecimalContext { return s.u0 }

func (s *CurveContext) GetU1() IDecimalContext { return s.u1 }

func (s *CurveContext) SetU0(v IDecimalContext) { s.u0 = v }

func (s *CurveContext) SetU1(v IDecimalContext) { s.u1 = v }

func (s *CurveContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *CurveContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *CurveContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *CurveContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *CurveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterCurve(s)
	}
}

func (s *CurveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitCurve(s)
	}
}

func (p *WavefrontOBJParser) Curve() (localctx ICurveContext) {
	this := p
	_ = this

	localctx = NewCurveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, WavefrontOBJParserRULE_curve)
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
		p.Match(WavefrontOBJParserT__18)
	}
	{
		p.SetState(259)

		var _x = p.Decimal()

		localctx.(*CurveContext).u0 = _x
	}
	{
		p.SetState(260)

		var _x = p.Decimal()

		localctx.(*CurveContext).u1 = _x
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER {
		{
			p.SetState(261)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*CurveContext).v = _m
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICurve2dContext is an interface to support dynamic dispatch.
type ICurve2dContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVp1 returns the vp1 token.
	GetVp1() antlr.Token

	// GetVp2 returns the vp2 token.
	GetVp2() antlr.Token

	// SetVp1 sets the vp1 token.
	SetVp1(antlr.Token)

	// SetVp2 sets the vp2 token.
	SetVp2(antlr.Token)

	// IsCurve2dContext differentiates from other interfaces.
	IsCurve2dContext()
}

type Curve2dContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vp1    antlr.Token
	vp2    antlr.Token
}

func NewEmptyCurve2dContext() *Curve2dContext {
	var p = new(Curve2dContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_curve2d
	return p
}

func (*Curve2dContext) IsCurve2dContext() {}

func NewCurve2dContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Curve2dContext {
	var p = new(Curve2dContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_curve2d

	return p
}

func (s *Curve2dContext) GetParser() antlr.Parser { return s.parser }

func (s *Curve2dContext) GetVp1() antlr.Token { return s.vp1 }

func (s *Curve2dContext) GetVp2() antlr.Token { return s.vp2 }

func (s *Curve2dContext) SetVp1(v antlr.Token) { s.vp1 = v }

func (s *Curve2dContext) SetVp2(v antlr.Token) { s.vp2 = v }

func (s *Curve2dContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *Curve2dContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *Curve2dContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Curve2dContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Curve2dContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterCurve2d(s)
	}
}

func (s *Curve2dContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitCurve2d(s)
	}
}

func (p *WavefrontOBJParser) Curve2d() (localctx ICurve2dContext) {
	this := p
	_ = this

	localctx = NewCurve2dContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, WavefrontOBJParserRULE_curve2d)
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
		p.SetState(266)
		p.Match(WavefrontOBJParserT__19)
	}
	{
		p.SetState(267)

		var _m = p.Match(WavefrontOBJParserINTEGER)

		localctx.(*Curve2dContext).vp1 = _m
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER {
		{
			p.SetState(268)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*Curve2dContext).vp2 = _m
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISurfaceContext is an interface to support dynamic dispatch.
type ISurfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v token.
	GetV() antlr.Token

	// SetV sets the v token.
	SetV(antlr.Token)

	// GetS0 returns the s0 rule contexts.
	GetS0() IDecimalContext

	// GetS1 returns the s1 rule contexts.
	GetS1() IDecimalContext

	// GetT0 returns the t0 rule contexts.
	GetT0() IDecimalContext

	// GetT1 returns the t1 rule contexts.
	GetT1() IDecimalContext

	// SetS0 sets the s0 rule contexts.
	SetS0(IDecimalContext)

	// SetS1 sets the s1 rule contexts.
	SetS1(IDecimalContext)

	// SetT0 sets the t0 rule contexts.
	SetT0(IDecimalContext)

	// SetT1 sets the t1 rule contexts.
	SetT1(IDecimalContext)

	// IsSurfaceContext differentiates from other interfaces.
	IsSurfaceContext()
}

type SurfaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	s0     IDecimalContext
	s1     IDecimalContext
	t0     IDecimalContext
	t1     IDecimalContext
	v      antlr.Token
}

func NewEmptySurfaceContext() *SurfaceContext {
	var p = new(SurfaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_surface
	return p
}

func (*SurfaceContext) IsSurfaceContext() {}

func NewSurfaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SurfaceContext {
	var p = new(SurfaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_surface

	return p
}

func (s *SurfaceContext) GetParser() antlr.Parser { return s.parser }

func (s *SurfaceContext) GetV() antlr.Token { return s.v }

func (s *SurfaceContext) SetV(v antlr.Token) { s.v = v }

func (s *SurfaceContext) GetS0() IDecimalContext { return s.s0 }

func (s *SurfaceContext) GetS1() IDecimalContext { return s.s1 }

func (s *SurfaceContext) GetT0() IDecimalContext { return s.t0 }

func (s *SurfaceContext) GetT1() IDecimalContext { return s.t1 }

func (s *SurfaceContext) SetS0(v IDecimalContext) { s.s0 = v }

func (s *SurfaceContext) SetS1(v IDecimalContext) { s.s1 = v }

func (s *SurfaceContext) SetT0(v IDecimalContext) { s.t0 = v }

func (s *SurfaceContext) SetT1(v IDecimalContext) { s.t1 = v }

func (s *SurfaceContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *SurfaceContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *SurfaceContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *SurfaceContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *SurfaceContext) AllINTEGER_PAIR() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER_PAIR)
}

func (s *SurfaceContext) INTEGER_PAIR(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER_PAIR, i)
}

func (s *SurfaceContext) AllINTEGER_TRIPLET() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER_TRIPLET)
}

func (s *SurfaceContext) INTEGER_TRIPLET(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER_TRIPLET, i)
}

func (s *SurfaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SurfaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SurfaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterSurface(s)
	}
}

func (s *SurfaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitSurface(s)
	}
}

func (p *WavefrontOBJParser) Surface() (localctx ISurfaceContext) {
	this := p
	_ = this

	localctx = NewSurfaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, WavefrontOBJParserRULE_surface)
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
		p.SetState(273)
		p.Match(WavefrontOBJParserT__20)
	}
	{
		p.SetState(274)

		var _x = p.Decimal()

		localctx.(*SurfaceContext).s0 = _x
	}
	{
		p.SetState(275)

		var _x = p.Decimal()

		localctx.(*SurfaceContext).s1 = _x
	}
	{
		p.SetState(276)

		var _x = p.Decimal()

		localctx.(*SurfaceContext).t0 = _x
	}
	{
		p.SetState(277)

		var _x = p.Decimal()

		localctx.(*SurfaceContext).t1 = _x
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(WavefrontOBJParserINTEGER_PAIR-51))|(1<<(WavefrontOBJParserINTEGER_TRIPLET-51))|(1<<(WavefrontOBJParserINTEGER-51)))) != 0) {
		{
			p.SetState(278)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SurfaceContext).v = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(WavefrontOBJParserINTEGER_PAIR-51))|(1<<(WavefrontOBJParserINTEGER_TRIPLET-51))|(1<<(WavefrontOBJParserINTEGER-51)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SurfaceContext).v = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP returns the p rule contexts.
	GetP() IDecimalContext

	// SetP sets the p rule contexts.
	SetP(IDecimalContext)

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	p      IDecimalContext
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) GetP() IDecimalContext { return s.p }

func (s *ParameterContext) SetP(v IDecimalContext) { s.p = v }

func (s *ParameterContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *ParameterContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *WavefrontOBJParser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, WavefrontOBJParserRULE_parameter)
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
		p.SetState(283)
		p.Match(WavefrontOBJParserT__21)
	}
	{
		p.SetState(284)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WavefrontOBJParserT__1 || _la == WavefrontOBJParserT__14) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER || _la == WavefrontOBJParserDECIMAL {
		{
			p.SetState(285)

			var _x = p.Decimal()

			localctx.(*ParameterContext).p = _x
		}

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOuter_trimming_loopContext is an interface to support dynamic dispatch.
type IOuter_trimming_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCurv2d returns the curv2d token.
	GetCurv2d() antlr.Token

	// SetCurv2d sets the curv2d token.
	SetCurv2d(antlr.Token)

	// GetU0 returns the u0 rule contexts.
	GetU0() IDecimalContext

	// GetU1 returns the u1 rule contexts.
	GetU1() IDecimalContext

	// SetU0 sets the u0 rule contexts.
	SetU0(IDecimalContext)

	// SetU1 sets the u1 rule contexts.
	SetU1(IDecimalContext)

	// IsOuter_trimming_loopContext differentiates from other interfaces.
	IsOuter_trimming_loopContext()
}

type Outer_trimming_loopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	u0     IDecimalContext
	u1     IDecimalContext
	curv2d antlr.Token
}

func NewEmptyOuter_trimming_loopContext() *Outer_trimming_loopContext {
	var p = new(Outer_trimming_loopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_outer_trimming_loop
	return p
}

func (*Outer_trimming_loopContext) IsOuter_trimming_loopContext() {}

func NewOuter_trimming_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Outer_trimming_loopContext {
	var p = new(Outer_trimming_loopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_outer_trimming_loop

	return p
}

func (s *Outer_trimming_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Outer_trimming_loopContext) GetCurv2d() antlr.Token { return s.curv2d }

func (s *Outer_trimming_loopContext) SetCurv2d(v antlr.Token) { s.curv2d = v }

func (s *Outer_trimming_loopContext) GetU0() IDecimalContext { return s.u0 }

func (s *Outer_trimming_loopContext) GetU1() IDecimalContext { return s.u1 }

func (s *Outer_trimming_loopContext) SetU0(v IDecimalContext) { s.u0 = v }

func (s *Outer_trimming_loopContext) SetU1(v IDecimalContext) { s.u1 = v }

func (s *Outer_trimming_loopContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *Outer_trimming_loopContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Outer_trimming_loopContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *Outer_trimming_loopContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *Outer_trimming_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Outer_trimming_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Outer_trimming_loopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterOuter_trimming_loop(s)
	}
}

func (s *Outer_trimming_loopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitOuter_trimming_loop(s)
	}
}

func (p *WavefrontOBJParser) Outer_trimming_loop() (localctx IOuter_trimming_loopContext) {
	this := p
	_ = this

	localctx = NewOuter_trimming_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, WavefrontOBJParserRULE_outer_trimming_loop)
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
		p.Match(WavefrontOBJParserT__22)
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER || _la == WavefrontOBJParserDECIMAL {
		{
			p.SetState(291)

			var _x = p.Decimal()

			localctx.(*Outer_trimming_loopContext).u0 = _x
		}
		{
			p.SetState(292)

			var _x = p.Decimal()

			localctx.(*Outer_trimming_loopContext).u1 = _x
		}
		{
			p.SetState(293)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*Outer_trimming_loopContext).curv2d = _m
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInner_trimming_loopContext is an interface to support dynamic dispatch.
type IInner_trimming_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCurv2d returns the curv2d token.
	GetCurv2d() antlr.Token

	// SetCurv2d sets the curv2d token.
	SetCurv2d(antlr.Token)

	// GetU0 returns the u0 rule contexts.
	GetU0() IDecimalContext

	// GetU1 returns the u1 rule contexts.
	GetU1() IDecimalContext

	// SetU0 sets the u0 rule contexts.
	SetU0(IDecimalContext)

	// SetU1 sets the u1 rule contexts.
	SetU1(IDecimalContext)

	// IsInner_trimming_loopContext differentiates from other interfaces.
	IsInner_trimming_loopContext()
}

type Inner_trimming_loopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	u0     IDecimalContext
	u1     IDecimalContext
	curv2d antlr.Token
}

func NewEmptyInner_trimming_loopContext() *Inner_trimming_loopContext {
	var p = new(Inner_trimming_loopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_inner_trimming_loop
	return p
}

func (*Inner_trimming_loopContext) IsInner_trimming_loopContext() {}

func NewInner_trimming_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inner_trimming_loopContext {
	var p = new(Inner_trimming_loopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_inner_trimming_loop

	return p
}

func (s *Inner_trimming_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Inner_trimming_loopContext) GetCurv2d() antlr.Token { return s.curv2d }

func (s *Inner_trimming_loopContext) SetCurv2d(v antlr.Token) { s.curv2d = v }

func (s *Inner_trimming_loopContext) GetU0() IDecimalContext { return s.u0 }

func (s *Inner_trimming_loopContext) GetU1() IDecimalContext { return s.u1 }

func (s *Inner_trimming_loopContext) SetU0(v IDecimalContext) { s.u0 = v }

func (s *Inner_trimming_loopContext) SetU1(v IDecimalContext) { s.u1 = v }

func (s *Inner_trimming_loopContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *Inner_trimming_loopContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Inner_trimming_loopContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *Inner_trimming_loopContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *Inner_trimming_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inner_trimming_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inner_trimming_loopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterInner_trimming_loop(s)
	}
}

func (s *Inner_trimming_loopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitInner_trimming_loop(s)
	}
}

func (p *WavefrontOBJParser) Inner_trimming_loop() (localctx IInner_trimming_loopContext) {
	this := p
	_ = this

	localctx = NewInner_trimming_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, WavefrontOBJParserRULE_inner_trimming_loop)
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
		p.Match(WavefrontOBJParserT__23)
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER || _la == WavefrontOBJParserDECIMAL {
		{
			p.SetState(300)

			var _x = p.Decimal()

			localctx.(*Inner_trimming_loopContext).u0 = _x
		}
		{
			p.SetState(301)

			var _x = p.Decimal()

			localctx.(*Inner_trimming_loopContext).u1 = _x
		}
		{
			p.SetState(302)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*Inner_trimming_loopContext).curv2d = _m
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISpecial_curveContext is an interface to support dynamic dispatch.
type ISpecial_curveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCurv2d returns the curv2d token.
	GetCurv2d() antlr.Token

	// SetCurv2d sets the curv2d token.
	SetCurv2d(antlr.Token)

	// GetU0 returns the u0 rule contexts.
	GetU0() IDecimalContext

	// GetU1 returns the u1 rule contexts.
	GetU1() IDecimalContext

	// SetU0 sets the u0 rule contexts.
	SetU0(IDecimalContext)

	// SetU1 sets the u1 rule contexts.
	SetU1(IDecimalContext)

	// IsSpecial_curveContext differentiates from other interfaces.
	IsSpecial_curveContext()
}

type Special_curveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	u0     IDecimalContext
	u1     IDecimalContext
	curv2d antlr.Token
}

func NewEmptySpecial_curveContext() *Special_curveContext {
	var p = new(Special_curveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_special_curve
	return p
}

func (*Special_curveContext) IsSpecial_curveContext() {}

func NewSpecial_curveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Special_curveContext {
	var p = new(Special_curveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_special_curve

	return p
}

func (s *Special_curveContext) GetParser() antlr.Parser { return s.parser }

func (s *Special_curveContext) GetCurv2d() antlr.Token { return s.curv2d }

func (s *Special_curveContext) SetCurv2d(v antlr.Token) { s.curv2d = v }

func (s *Special_curveContext) GetU0() IDecimalContext { return s.u0 }

func (s *Special_curveContext) GetU1() IDecimalContext { return s.u1 }

func (s *Special_curveContext) SetU0(v IDecimalContext) { s.u0 = v }

func (s *Special_curveContext) SetU1(v IDecimalContext) { s.u1 = v }

func (s *Special_curveContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *Special_curveContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Special_curveContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *Special_curveContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *Special_curveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Special_curveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Special_curveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterSpecial_curve(s)
	}
}

func (s *Special_curveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitSpecial_curve(s)
	}
}

func (p *WavefrontOBJParser) Special_curve() (localctx ISpecial_curveContext) {
	this := p
	_ = this

	localctx = NewSpecial_curveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, WavefrontOBJParserRULE_special_curve)
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
		p.SetState(308)
		p.Match(WavefrontOBJParserT__24)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER || _la == WavefrontOBJParserDECIMAL {
		{
			p.SetState(309)

			var _x = p.Decimal()

			localctx.(*Special_curveContext).u0 = _x
		}
		{
			p.SetState(310)

			var _x = p.Decimal()

			localctx.(*Special_curveContext).u1 = _x
		}
		{
			p.SetState(311)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*Special_curveContext).curv2d = _m
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISpecial_pointContext is an interface to support dynamic dispatch.
type ISpecial_pointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVp returns the vp token.
	GetVp() antlr.Token

	// SetVp sets the vp token.
	SetVp(antlr.Token)

	// IsSpecial_pointContext differentiates from other interfaces.
	IsSpecial_pointContext()
}

type Special_pointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vp     antlr.Token
}

func NewEmptySpecial_pointContext() *Special_pointContext {
	var p = new(Special_pointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_special_point
	return p
}

func (*Special_pointContext) IsSpecial_pointContext() {}

func NewSpecial_pointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Special_pointContext {
	var p = new(Special_pointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_special_point

	return p
}

func (s *Special_pointContext) GetParser() antlr.Parser { return s.parser }

func (s *Special_pointContext) GetVp() antlr.Token { return s.vp }

func (s *Special_pointContext) SetVp(v antlr.Token) { s.vp = v }

func (s *Special_pointContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *Special_pointContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *Special_pointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Special_pointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Special_pointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterSpecial_point(s)
	}
}

func (s *Special_pointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitSpecial_point(s)
	}
}

func (p *WavefrontOBJParser) Special_point() (localctx ISpecial_pointContext) {
	this := p
	_ = this

	localctx = NewSpecial_pointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, WavefrontOBJParserRULE_special_point)
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
		p.SetState(317)
		p.Match(WavefrontOBJParserT__25)
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER {
		{
			p.SetState(318)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*Special_pointContext).vp = _m
		}

		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEndContext is an interface to support dynamic dispatch.
type IEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndContext differentiates from other interfaces.
	IsEndContext()
}

type EndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndContext() *EndContext {
	var p = new(EndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_end
	return p
}

func (*EndContext) IsEndContext() {}

func NewEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndContext {
	var p = new(EndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_end

	return p
}

func (s *EndContext) GetParser() antlr.Parser { return s.parser }
func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitEnd(s)
	}
}

func (p *WavefrontOBJParser) End() (localctx IEndContext) {
	this := p
	_ = this

	localctx = NewEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, WavefrontOBJParserRULE_end)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(WavefrontOBJParserT__26)
	}

	return localctx
}

// IConnectivityContext is an interface to support dynamic dispatch.
type IConnectivityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSurf_1 returns the surf_1 token.
	GetSurf_1() antlr.Token

	// GetCurv2d_1 returns the curv2d_1 token.
	GetCurv2d_1() antlr.Token

	// GetSurf_2 returns the surf_2 token.
	GetSurf_2() antlr.Token

	// GetCurv2d_2 returns the curv2d_2 token.
	GetCurv2d_2() antlr.Token

	// SetSurf_1 sets the surf_1 token.
	SetSurf_1(antlr.Token)

	// SetCurv2d_1 sets the curv2d_1 token.
	SetCurv2d_1(antlr.Token)

	// SetSurf_2 sets the surf_2 token.
	SetSurf_2(antlr.Token)

	// SetCurv2d_2 sets the curv2d_2 token.
	SetCurv2d_2(antlr.Token)

	// GetQ0_1 returns the q0_1 rule contexts.
	GetQ0_1() IDecimalContext

	// GetQ1_1 returns the q1_1 rule contexts.
	GetQ1_1() IDecimalContext

	// GetQ0_2 returns the q0_2 rule contexts.
	GetQ0_2() IDecimalContext

	// GetQ1_2 returns the q1_2 rule contexts.
	GetQ1_2() IDecimalContext

	// SetQ0_1 sets the q0_1 rule contexts.
	SetQ0_1(IDecimalContext)

	// SetQ1_1 sets the q1_1 rule contexts.
	SetQ1_1(IDecimalContext)

	// SetQ0_2 sets the q0_2 rule contexts.
	SetQ0_2(IDecimalContext)

	// SetQ1_2 sets the q1_2 rule contexts.
	SetQ1_2(IDecimalContext)

	// IsConnectivityContext differentiates from other interfaces.
	IsConnectivityContext()
}

type ConnectivityContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	surf_1   antlr.Token
	q0_1     IDecimalContext
	q1_1     IDecimalContext
	curv2d_1 antlr.Token
	surf_2   antlr.Token
	q0_2     IDecimalContext
	q1_2     IDecimalContext
	curv2d_2 antlr.Token
}

func NewEmptyConnectivityContext() *ConnectivityContext {
	var p = new(ConnectivityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_connectivity
	return p
}

func (*ConnectivityContext) IsConnectivityContext() {}

func NewConnectivityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectivityContext {
	var p = new(ConnectivityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_connectivity

	return p
}

func (s *ConnectivityContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectivityContext) GetSurf_1() antlr.Token { return s.surf_1 }

func (s *ConnectivityContext) GetCurv2d_1() antlr.Token { return s.curv2d_1 }

func (s *ConnectivityContext) GetSurf_2() antlr.Token { return s.surf_2 }

func (s *ConnectivityContext) GetCurv2d_2() antlr.Token { return s.curv2d_2 }

func (s *ConnectivityContext) SetSurf_1(v antlr.Token) { s.surf_1 = v }

func (s *ConnectivityContext) SetCurv2d_1(v antlr.Token) { s.curv2d_1 = v }

func (s *ConnectivityContext) SetSurf_2(v antlr.Token) { s.surf_2 = v }

func (s *ConnectivityContext) SetCurv2d_2(v antlr.Token) { s.curv2d_2 = v }

func (s *ConnectivityContext) GetQ0_1() IDecimalContext { return s.q0_1 }

func (s *ConnectivityContext) GetQ1_1() IDecimalContext { return s.q1_1 }

func (s *ConnectivityContext) GetQ0_2() IDecimalContext { return s.q0_2 }

func (s *ConnectivityContext) GetQ1_2() IDecimalContext { return s.q1_2 }

func (s *ConnectivityContext) SetQ0_1(v IDecimalContext) { s.q0_1 = v }

func (s *ConnectivityContext) SetQ1_1(v IDecimalContext) { s.q1_1 = v }

func (s *ConnectivityContext) SetQ0_2(v IDecimalContext) { s.q0_2 = v }

func (s *ConnectivityContext) SetQ1_2(v IDecimalContext) { s.q1_2 = v }

func (s *ConnectivityContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *ConnectivityContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *ConnectivityContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *ConnectivityContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *ConnectivityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectivityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectivityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterConnectivity(s)
	}
}

func (s *ConnectivityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitConnectivity(s)
	}
}

func (p *WavefrontOBJParser) Connectivity() (localctx IConnectivityContext) {
	this := p
	_ = this

	localctx = NewConnectivityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, WavefrontOBJParserRULE_connectivity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(WavefrontOBJParserT__27)
	}
	{
		p.SetState(326)

		var _m = p.Match(WavefrontOBJParserINTEGER)

		localctx.(*ConnectivityContext).surf_1 = _m
	}
	{
		p.SetState(327)

		var _x = p.Decimal()

		localctx.(*ConnectivityContext).q0_1 = _x
	}
	{
		p.SetState(328)

		var _x = p.Decimal()

		localctx.(*ConnectivityContext).q1_1 = _x
	}
	{
		p.SetState(329)

		var _m = p.Match(WavefrontOBJParserINTEGER)

		localctx.(*ConnectivityContext).curv2d_1 = _m
	}
	{
		p.SetState(330)

		var _m = p.Match(WavefrontOBJParserINTEGER)

		localctx.(*ConnectivityContext).surf_2 = _m
	}
	{
		p.SetState(331)

		var _x = p.Decimal()

		localctx.(*ConnectivityContext).q0_2 = _x
	}
	{
		p.SetState(332)

		var _x = p.Decimal()

		localctx.(*ConnectivityContext).q1_2 = _x
	}
	{
		p.SetState(333)

		var _m = p.Match(WavefrontOBJParserINTEGER)

		localctx.(*ConnectivityContext).curv2d_2 = _m
	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetGroup_name returns the group_name token.
	GetGroup_name() antlr.Token

	// SetGroup_name sets the group_name token.
	SetGroup_name(antlr.Token)

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	group_name antlr.Token
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) GetGroup_name() antlr.Token { return s.group_name }

func (s *GroupContext) SetGroup_name(v antlr.Token) { s.group_name = v }

func (s *GroupContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserNAME)
}

func (s *GroupContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNAME, i)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *WavefrontOBJParser) Group() (localctx IGroupContext) {
	this := p
	_ = this

	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, WavefrontOBJParserRULE_group)
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
		p.SetState(335)
		p.Match(WavefrontOBJParserT__28)
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserNAME {
		{
			p.SetState(336)

			var _m = p.Match(WavefrontOBJParserNAME)

			localctx.(*GroupContext).group_name = _m
		}

		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISmoothing_groupContext is an interface to support dynamic dispatch.
type ISmoothing_groupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetGroup_number returns the group_number token.
	GetGroup_number() antlr.Token

	// SetGroup_number sets the group_number token.
	SetGroup_number(antlr.Token)

	// IsSmoothing_groupContext differentiates from other interfaces.
	IsSmoothing_groupContext()
}

type Smoothing_groupContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	group_number antlr.Token
}

func NewEmptySmoothing_groupContext() *Smoothing_groupContext {
	var p = new(Smoothing_groupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_smoothing_group
	return p
}

func (*Smoothing_groupContext) IsSmoothing_groupContext() {}

func NewSmoothing_groupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Smoothing_groupContext {
	var p = new(Smoothing_groupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_smoothing_group

	return p
}

func (s *Smoothing_groupContext) GetParser() antlr.Parser { return s.parser }

func (s *Smoothing_groupContext) GetGroup_number() antlr.Token { return s.group_number }

func (s *Smoothing_groupContext) SetGroup_number(v antlr.Token) { s.group_number = v }

func (s *Smoothing_groupContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, 0)
}

func (s *Smoothing_groupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Smoothing_groupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Smoothing_groupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterSmoothing_group(s)
	}
}

func (s *Smoothing_groupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitSmoothing_group(s)
	}
}

func (p *WavefrontOBJParser) Smoothing_group() (localctx ISmoothing_groupContext) {
	this := p
	_ = this

	localctx = NewSmoothing_groupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, WavefrontOBJParserRULE_smoothing_group)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(WavefrontOBJParserT__29)
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WavefrontOBJParserINTEGER:
		{
			p.SetState(342)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*Smoothing_groupContext).group_number = _m
		}

	case WavefrontOBJParserT__30:
		{
			p.SetState(343)
			p.Match(WavefrontOBJParserT__30)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMerging_groupContext is an interface to support dynamic dispatch.
type IMerging_groupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetGroup_number returns the group_number token.
	GetGroup_number() antlr.Token

	// SetGroup_number sets the group_number token.
	SetGroup_number(antlr.Token)

	// GetRes returns the res rule contexts.
	GetRes() IDecimalContext

	// SetRes sets the res rule contexts.
	SetRes(IDecimalContext)

	// IsMerging_groupContext differentiates from other interfaces.
	IsMerging_groupContext()
}

type Merging_groupContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	group_number antlr.Token
	res          IDecimalContext
}

func NewEmptyMerging_groupContext() *Merging_groupContext {
	var p = new(Merging_groupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_merging_group
	return p
}

func (*Merging_groupContext) IsMerging_groupContext() {}

func NewMerging_groupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Merging_groupContext {
	var p = new(Merging_groupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_merging_group

	return p
}

func (s *Merging_groupContext) GetParser() antlr.Parser { return s.parser }

func (s *Merging_groupContext) GetGroup_number() antlr.Token { return s.group_number }

func (s *Merging_groupContext) SetGroup_number(v antlr.Token) { s.group_number = v }

func (s *Merging_groupContext) GetRes() IDecimalContext { return s.res }

func (s *Merging_groupContext) SetRes(v IDecimalContext) { s.res = v }

func (s *Merging_groupContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, 0)
}

func (s *Merging_groupContext) Decimal() IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Merging_groupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Merging_groupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Merging_groupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterMerging_group(s)
	}
}

func (s *Merging_groupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitMerging_group(s)
	}
}

func (p *WavefrontOBJParser) Merging_group() (localctx IMerging_groupContext) {
	this := p
	_ = this

	localctx = NewMerging_groupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, WavefrontOBJParserRULE_merging_group)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(WavefrontOBJParserT__31)
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WavefrontOBJParserINTEGER:
		{
			p.SetState(347)

			var _m = p.Match(WavefrontOBJParserINTEGER)

			localctx.(*Merging_groupContext).group_number = _m
		}
		{
			p.SetState(348)

			var _x = p.Decimal()

			localctx.(*Merging_groupContext).res = _x
		}

	case WavefrontOBJParserT__30:
		{
			p.SetState(349)
			p.Match(WavefrontOBJParserT__30)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObject_nameContext is an interface to support dynamic dispatch.
type IObject_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsObject_nameContext differentiates from other interfaces.
	IsObject_nameContext()
}

type Object_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyObject_nameContext() *Object_nameContext {
	var p = new(Object_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_object_name
	return p
}

func (*Object_nameContext) IsObject_nameContext() {}

func NewObject_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_nameContext {
	var p = new(Object_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_object_name

	return p
}

func (s *Object_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_nameContext) GetName() antlr.Token { return s.name }

func (s *Object_nameContext) SetName(v antlr.Token) { s.name = v }

func (s *Object_nameContext) NAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNAME, 0)
}

func (s *Object_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterObject_name(s)
	}
}

func (s *Object_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitObject_name(s)
	}
}

func (p *WavefrontOBJParser) Object_name() (localctx IObject_nameContext) {
	this := p
	_ = this

	localctx = NewObject_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, WavefrontOBJParserRULE_object_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(WavefrontOBJParserT__32)
	}
	{
		p.SetState(353)

		var _m = p.Match(WavefrontOBJParserNAME)

		localctx.(*Object_nameContext).name = _m
	}

	return localctx
}

// IBevelContext is an interface to support dynamic dispatch.
type IBevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBevelContext differentiates from other interfaces.
	IsBevelContext()
}

type BevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBevelContext() *BevelContext {
	var p = new(BevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_bevel
	return p
}

func (*BevelContext) IsBevelContext() {}

func NewBevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BevelContext {
	var p = new(BevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_bevel

	return p
}

func (s *BevelContext) GetParser() antlr.Parser { return s.parser }
func (s *BevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterBevel(s)
	}
}

func (s *BevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitBevel(s)
	}
}

func (p *WavefrontOBJParser) Bevel() (localctx IBevelContext) {
	this := p
	_ = this

	localctx = NewBevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, WavefrontOBJParserRULE_bevel)
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
		p.SetState(355)
		p.Match(WavefrontOBJParserT__33)
	}
	{
		p.SetState(356)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WavefrontOBJParserT__30 || _la == WavefrontOBJParserT__34) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IColor_interpolationContext is an interface to support dynamic dispatch.
type IColor_interpolationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColor_interpolationContext differentiates from other interfaces.
	IsColor_interpolationContext()
}

type Color_interpolationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColor_interpolationContext() *Color_interpolationContext {
	var p = new(Color_interpolationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_color_interpolation
	return p
}

func (*Color_interpolationContext) IsColor_interpolationContext() {}

func NewColor_interpolationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Color_interpolationContext {
	var p = new(Color_interpolationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_color_interpolation

	return p
}

func (s *Color_interpolationContext) GetParser() antlr.Parser { return s.parser }
func (s *Color_interpolationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Color_interpolationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Color_interpolationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterColor_interpolation(s)
	}
}

func (s *Color_interpolationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitColor_interpolation(s)
	}
}

func (p *WavefrontOBJParser) Color_interpolation() (localctx IColor_interpolationContext) {
	this := p
	_ = this

	localctx = NewColor_interpolationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, WavefrontOBJParserRULE_color_interpolation)
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
		p.SetState(358)
		p.Match(WavefrontOBJParserT__35)
	}
	{
		p.SetState(359)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WavefrontOBJParserT__30 || _la == WavefrontOBJParserT__34) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDissolve_interpolationContext is an interface to support dynamic dispatch.
type IDissolve_interpolationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDissolve_interpolationContext differentiates from other interfaces.
	IsDissolve_interpolationContext()
}

type Dissolve_interpolationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDissolve_interpolationContext() *Dissolve_interpolationContext {
	var p = new(Dissolve_interpolationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_dissolve_interpolation
	return p
}

func (*Dissolve_interpolationContext) IsDissolve_interpolationContext() {}

func NewDissolve_interpolationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dissolve_interpolationContext {
	var p = new(Dissolve_interpolationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_dissolve_interpolation

	return p
}

func (s *Dissolve_interpolationContext) GetParser() antlr.Parser { return s.parser }
func (s *Dissolve_interpolationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dissolve_interpolationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dissolve_interpolationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterDissolve_interpolation(s)
	}
}

func (s *Dissolve_interpolationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitDissolve_interpolation(s)
	}
}

func (p *WavefrontOBJParser) Dissolve_interpolation() (localctx IDissolve_interpolationContext) {
	this := p
	_ = this

	localctx = NewDissolve_interpolationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, WavefrontOBJParserRULE_dissolve_interpolation)
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
		p.SetState(361)
		p.Match(WavefrontOBJParserT__36)
	}
	{
		p.SetState(362)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WavefrontOBJParserT__30 || _la == WavefrontOBJParserT__34) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILevel_of_detailContext is an interface to support dynamic dispatch.
type ILevel_of_detailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLevel returns the level token.
	GetLevel() antlr.Token

	// SetLevel sets the level token.
	SetLevel(antlr.Token)

	// IsLevel_of_detailContext differentiates from other interfaces.
	IsLevel_of_detailContext()
}

type Level_of_detailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	level  antlr.Token
}

func NewEmptyLevel_of_detailContext() *Level_of_detailContext {
	var p = new(Level_of_detailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_level_of_detail
	return p
}

func (*Level_of_detailContext) IsLevel_of_detailContext() {}

func NewLevel_of_detailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Level_of_detailContext {
	var p = new(Level_of_detailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_level_of_detail

	return p
}

func (s *Level_of_detailContext) GetParser() antlr.Parser { return s.parser }

func (s *Level_of_detailContext) GetLevel() antlr.Token { return s.level }

func (s *Level_of_detailContext) SetLevel(v antlr.Token) { s.level = v }

func (s *Level_of_detailContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, 0)
}

func (s *Level_of_detailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Level_of_detailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Level_of_detailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterLevel_of_detail(s)
	}
}

func (s *Level_of_detailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitLevel_of_detail(s)
	}
}

func (p *WavefrontOBJParser) Level_of_detail() (localctx ILevel_of_detailContext) {
	this := p
	_ = this

	localctx = NewLevel_of_detailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, WavefrontOBJParserRULE_level_of_detail)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(WavefrontOBJParserT__37)
	}
	{
		p.SetState(365)

		var _m = p.Match(WavefrontOBJParserINTEGER)

		localctx.(*Level_of_detailContext).level = _m
	}

	return localctx
}

// IMap_libraryContext is an interface to support dynamic dispatch.
type IMap_libraryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFilename returns the filename token.
	GetFilename() antlr.Token

	// SetFilename sets the filename token.
	SetFilename(antlr.Token)

	// IsMap_libraryContext differentiates from other interfaces.
	IsMap_libraryContext()
}

type Map_libraryContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	filename antlr.Token
}

func NewEmptyMap_libraryContext() *Map_libraryContext {
	var p = new(Map_libraryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_map_library
	return p
}

func (*Map_libraryContext) IsMap_libraryContext() {}

func NewMap_libraryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_libraryContext {
	var p = new(Map_libraryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_map_library

	return p
}

func (s *Map_libraryContext) GetParser() antlr.Parser { return s.parser }

func (s *Map_libraryContext) GetFilename() antlr.Token { return s.filename }

func (s *Map_libraryContext) SetFilename(v antlr.Token) { s.filename = v }

func (s *Map_libraryContext) AllFILENAME() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserFILENAME)
}

func (s *Map_libraryContext) FILENAME(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserFILENAME, i)
}

func (s *Map_libraryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_libraryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_libraryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterMap_library(s)
	}
}

func (s *Map_libraryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitMap_library(s)
	}
}

func (p *WavefrontOBJParser) Map_library() (localctx IMap_libraryContext) {
	this := p
	_ = this

	localctx = NewMap_libraryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, WavefrontOBJParserRULE_map_library)
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
		p.SetState(367)
		p.Match(WavefrontOBJParserT__38)
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserFILENAME {
		{
			p.SetState(368)

			var _m = p.Match(WavefrontOBJParserFILENAME)

			localctx.(*Map_libraryContext).filename = _m
		}

		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUse_mapContext is an interface to support dynamic dispatch.
type IUse_mapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMap_name returns the map_name token.
	GetMap_name() antlr.Token

	// SetMap_name sets the map_name token.
	SetMap_name(antlr.Token)

	// IsUse_mapContext differentiates from other interfaces.
	IsUse_mapContext()
}

type Use_mapContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	map_name antlr.Token
}

func NewEmptyUse_mapContext() *Use_mapContext {
	var p = new(Use_mapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_use_map
	return p
}

func (*Use_mapContext) IsUse_mapContext() {}

func NewUse_mapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Use_mapContext {
	var p = new(Use_mapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_use_map

	return p
}

func (s *Use_mapContext) GetParser() antlr.Parser { return s.parser }

func (s *Use_mapContext) GetMap_name() antlr.Token { return s.map_name }

func (s *Use_mapContext) SetMap_name(v antlr.Token) { s.map_name = v }

func (s *Use_mapContext) NAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNAME, 0)
}

func (s *Use_mapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Use_mapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Use_mapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterUse_map(s)
	}
}

func (s *Use_mapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitUse_map(s)
	}
}

func (p *WavefrontOBJParser) Use_map() (localctx IUse_mapContext) {
	this := p
	_ = this

	localctx = NewUse_mapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, WavefrontOBJParserRULE_use_map)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Match(WavefrontOBJParserT__39)
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WavefrontOBJParserNAME:
		{
			p.SetState(374)

			var _m = p.Match(WavefrontOBJParserNAME)

			localctx.(*Use_mapContext).map_name = _m
		}

	case WavefrontOBJParserT__30:
		{
			p.SetState(375)
			p.Match(WavefrontOBJParserT__30)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMaterial_libraryContext is an interface to support dynamic dispatch.
type IMaterial_libraryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFilename returns the filename token.
	GetFilename() antlr.Token

	// SetFilename sets the filename token.
	SetFilename(antlr.Token)

	// IsMaterial_libraryContext differentiates from other interfaces.
	IsMaterial_libraryContext()
}

type Material_libraryContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	filename antlr.Token
}

func NewEmptyMaterial_libraryContext() *Material_libraryContext {
	var p = new(Material_libraryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_material_library
	return p
}

func (*Material_libraryContext) IsMaterial_libraryContext() {}

func NewMaterial_libraryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Material_libraryContext {
	var p = new(Material_libraryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_material_library

	return p
}

func (s *Material_libraryContext) GetParser() antlr.Parser { return s.parser }

func (s *Material_libraryContext) GetFilename() antlr.Token { return s.filename }

func (s *Material_libraryContext) SetFilename(v antlr.Token) { s.filename = v }

func (s *Material_libraryContext) FILENAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserFILENAME, 0)
}

func (s *Material_libraryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Material_libraryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Material_libraryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterMaterial_library(s)
	}
}

func (s *Material_libraryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitMaterial_library(s)
	}
}

func (p *WavefrontOBJParser) Material_library() (localctx IMaterial_libraryContext) {
	this := p
	_ = this

	localctx = NewMaterial_libraryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, WavefrontOBJParserRULE_material_library)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(WavefrontOBJParserT__40)
	}
	{
		p.SetState(379)

		var _m = p.Match(WavefrontOBJParserFILENAME)

		localctx.(*Material_libraryContext).filename = _m
	}

	return localctx
}

// IUse_materialContext is an interface to support dynamic dispatch.
type IUse_materialContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUse_materialContext differentiates from other interfaces.
	IsUse_materialContext()
}

type Use_materialContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUse_materialContext() *Use_materialContext {
	var p = new(Use_materialContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_use_material
	return p
}

func (*Use_materialContext) IsUse_materialContext() {}

func NewUse_materialContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Use_materialContext {
	var p = new(Use_materialContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_use_material

	return p
}

func (s *Use_materialContext) GetParser() antlr.Parser { return s.parser }

func (s *Use_materialContext) NAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNAME, 0)
}

func (s *Use_materialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Use_materialContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Use_materialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterUse_material(s)
	}
}

func (s *Use_materialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitUse_material(s)
	}
}

func (p *WavefrontOBJParser) Use_material() (localctx IUse_materialContext) {
	this := p
	_ = this

	localctx = NewUse_materialContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, WavefrontOBJParserRULE_use_material)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(WavefrontOBJParserT__41)
	}
	{
		p.SetState(382)
		p.Match(WavefrontOBJParserNAME)
	}

	return localctx
}

// IShadow_objectContext is an interface to support dynamic dispatch.
type IShadow_objectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFilename returns the filename token.
	GetFilename() antlr.Token

	// SetFilename sets the filename token.
	SetFilename(antlr.Token)

	// IsShadow_objectContext differentiates from other interfaces.
	IsShadow_objectContext()
}

type Shadow_objectContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	filename antlr.Token
}

func NewEmptyShadow_objectContext() *Shadow_objectContext {
	var p = new(Shadow_objectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_shadow_object
	return p
}

func (*Shadow_objectContext) IsShadow_objectContext() {}

func NewShadow_objectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Shadow_objectContext {
	var p = new(Shadow_objectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_shadow_object

	return p
}

func (s *Shadow_objectContext) GetParser() antlr.Parser { return s.parser }

func (s *Shadow_objectContext) GetFilename() antlr.Token { return s.filename }

func (s *Shadow_objectContext) SetFilename(v antlr.Token) { s.filename = v }

func (s *Shadow_objectContext) FILENAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserFILENAME, 0)
}

func (s *Shadow_objectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Shadow_objectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Shadow_objectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterShadow_object(s)
	}
}

func (s *Shadow_objectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitShadow_object(s)
	}
}

func (p *WavefrontOBJParser) Shadow_object() (localctx IShadow_objectContext) {
	this := p
	_ = this

	localctx = NewShadow_objectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, WavefrontOBJParserRULE_shadow_object)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		p.Match(WavefrontOBJParserT__42)
	}
	{
		p.SetState(385)

		var _m = p.Match(WavefrontOBJParserFILENAME)

		localctx.(*Shadow_objectContext).filename = _m
	}

	return localctx
}

// ITrace_objectContext is an interface to support dynamic dispatch.
type ITrace_objectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFilename returns the filename token.
	GetFilename() antlr.Token

	// SetFilename sets the filename token.
	SetFilename(antlr.Token)

	// IsTrace_objectContext differentiates from other interfaces.
	IsTrace_objectContext()
}

type Trace_objectContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	filename antlr.Token
}

func NewEmptyTrace_objectContext() *Trace_objectContext {
	var p = new(Trace_objectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_trace_object
	return p
}

func (*Trace_objectContext) IsTrace_objectContext() {}

func NewTrace_objectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Trace_objectContext {
	var p = new(Trace_objectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_trace_object

	return p
}

func (s *Trace_objectContext) GetParser() antlr.Parser { return s.parser }

func (s *Trace_objectContext) GetFilename() antlr.Token { return s.filename }

func (s *Trace_objectContext) SetFilename(v antlr.Token) { s.filename = v }

func (s *Trace_objectContext) FILENAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserFILENAME, 0)
}

func (s *Trace_objectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Trace_objectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Trace_objectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterTrace_object(s)
	}
}

func (s *Trace_objectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitTrace_object(s)
	}
}

func (p *WavefrontOBJParser) Trace_object() (localctx ITrace_objectContext) {
	this := p
	_ = this

	localctx = NewTrace_objectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, WavefrontOBJParserRULE_trace_object)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(WavefrontOBJParserT__43)
	}
	{
		p.SetState(388)

		var _m = p.Match(WavefrontOBJParserFILENAME)

		localctx.(*Trace_objectContext).filename = _m
	}

	return localctx
}

// ICurve_approximation_techniqueContext is an interface to support dynamic dispatch.
type ICurve_approximation_techniqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRes returns the res rule contexts.
	GetRes() IDecimalContext

	// GetMaxlength returns the maxlength rule contexts.
	GetMaxlength() IDecimalContext

	// GetMaxdist returns the maxdist rule contexts.
	GetMaxdist() IDecimalContext

	// GetMaxangle returns the maxangle rule contexts.
	GetMaxangle() IDecimalContext

	// SetRes sets the res rule contexts.
	SetRes(IDecimalContext)

	// SetMaxlength sets the maxlength rule contexts.
	SetMaxlength(IDecimalContext)

	// SetMaxdist sets the maxdist rule contexts.
	SetMaxdist(IDecimalContext)

	// SetMaxangle sets the maxangle rule contexts.
	SetMaxangle(IDecimalContext)

	// IsCurve_approximation_techniqueContext differentiates from other interfaces.
	IsCurve_approximation_techniqueContext()
}

type Curve_approximation_techniqueContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	res       IDecimalContext
	maxlength IDecimalContext
	maxdist   IDecimalContext
	maxangle  IDecimalContext
}

func NewEmptyCurve_approximation_techniqueContext() *Curve_approximation_techniqueContext {
	var p = new(Curve_approximation_techniqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_curve_approximation_technique
	return p
}

func (*Curve_approximation_techniqueContext) IsCurve_approximation_techniqueContext() {}

func NewCurve_approximation_techniqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Curve_approximation_techniqueContext {
	var p = new(Curve_approximation_techniqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_curve_approximation_technique

	return p
}

func (s *Curve_approximation_techniqueContext) GetParser() antlr.Parser { return s.parser }

func (s *Curve_approximation_techniqueContext) GetRes() IDecimalContext { return s.res }

func (s *Curve_approximation_techniqueContext) GetMaxlength() IDecimalContext { return s.maxlength }

func (s *Curve_approximation_techniqueContext) GetMaxdist() IDecimalContext { return s.maxdist }

func (s *Curve_approximation_techniqueContext) GetMaxangle() IDecimalContext { return s.maxangle }

func (s *Curve_approximation_techniqueContext) SetRes(v IDecimalContext) { s.res = v }

func (s *Curve_approximation_techniqueContext) SetMaxlength(v IDecimalContext) { s.maxlength = v }

func (s *Curve_approximation_techniqueContext) SetMaxdist(v IDecimalContext) { s.maxdist = v }

func (s *Curve_approximation_techniqueContext) SetMaxangle(v IDecimalContext) { s.maxangle = v }

func (s *Curve_approximation_techniqueContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *Curve_approximation_techniqueContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Curve_approximation_techniqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Curve_approximation_techniqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Curve_approximation_techniqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterCurve_approximation_technique(s)
	}
}

func (s *Curve_approximation_techniqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitCurve_approximation_technique(s)
	}
}

func (p *WavefrontOBJParser) Curve_approximation_technique() (localctx ICurve_approximation_techniqueContext) {
	this := p
	_ = this

	localctx = NewCurve_approximation_techniqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, WavefrontOBJParserRULE_curve_approximation_technique)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(WavefrontOBJParserT__44)
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WavefrontOBJParserT__45:
		{
			p.SetState(391)
			p.Match(WavefrontOBJParserT__45)
		}
		{
			p.SetState(392)

			var _x = p.Decimal()

			localctx.(*Curve_approximation_techniqueContext).res = _x
		}

	case WavefrontOBJParserT__46:
		{
			p.SetState(393)
			p.Match(WavefrontOBJParserT__46)
		}
		{
			p.SetState(394)

			var _x = p.Decimal()

			localctx.(*Curve_approximation_techniqueContext).maxlength = _x
		}

	case WavefrontOBJParserT__18:
		{
			p.SetState(395)
			p.Match(WavefrontOBJParserT__18)
		}
		{
			p.SetState(396)

			var _x = p.Decimal()

			localctx.(*Curve_approximation_techniqueContext).maxdist = _x
		}
		{
			p.SetState(397)

			var _x = p.Decimal()

			localctx.(*Curve_approximation_techniqueContext).maxangle = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISurface_approximation_techniqueContext is an interface to support dynamic dispatch.
type ISurface_approximation_techniqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUres returns the ures rule contexts.
	GetUres() IDecimalContext

	// GetVres returns the vres rule contexts.
	GetVres() IDecimalContext

	// GetUvres returns the uvres rule contexts.
	GetUvres() IDecimalContext

	// GetMaxlength returns the maxlength rule contexts.
	GetMaxlength() IDecimalContext

	// GetMaxdist returns the maxdist rule contexts.
	GetMaxdist() IDecimalContext

	// GetMaxangle returns the maxangle rule contexts.
	GetMaxangle() IDecimalContext

	// SetUres sets the ures rule contexts.
	SetUres(IDecimalContext)

	// SetVres sets the vres rule contexts.
	SetVres(IDecimalContext)

	// SetUvres sets the uvres rule contexts.
	SetUvres(IDecimalContext)

	// SetMaxlength sets the maxlength rule contexts.
	SetMaxlength(IDecimalContext)

	// SetMaxdist sets the maxdist rule contexts.
	SetMaxdist(IDecimalContext)

	// SetMaxangle sets the maxangle rule contexts.
	SetMaxangle(IDecimalContext)

	// IsSurface_approximation_techniqueContext differentiates from other interfaces.
	IsSurface_approximation_techniqueContext()
}

type Surface_approximation_techniqueContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	ures      IDecimalContext
	vres      IDecimalContext
	uvres     IDecimalContext
	maxlength IDecimalContext
	maxdist   IDecimalContext
	maxangle  IDecimalContext
}

func NewEmptySurface_approximation_techniqueContext() *Surface_approximation_techniqueContext {
	var p = new(Surface_approximation_techniqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_surface_approximation_technique
	return p
}

func (*Surface_approximation_techniqueContext) IsSurface_approximation_techniqueContext() {}

func NewSurface_approximation_techniqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Surface_approximation_techniqueContext {
	var p = new(Surface_approximation_techniqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_surface_approximation_technique

	return p
}

func (s *Surface_approximation_techniqueContext) GetParser() antlr.Parser { return s.parser }

func (s *Surface_approximation_techniqueContext) GetUres() IDecimalContext { return s.ures }

func (s *Surface_approximation_techniqueContext) GetVres() IDecimalContext { return s.vres }

func (s *Surface_approximation_techniqueContext) GetUvres() IDecimalContext { return s.uvres }

func (s *Surface_approximation_techniqueContext) GetMaxlength() IDecimalContext { return s.maxlength }

func (s *Surface_approximation_techniqueContext) GetMaxdist() IDecimalContext { return s.maxdist }

func (s *Surface_approximation_techniqueContext) GetMaxangle() IDecimalContext { return s.maxangle }

func (s *Surface_approximation_techniqueContext) SetUres(v IDecimalContext) { s.ures = v }

func (s *Surface_approximation_techniqueContext) SetVres(v IDecimalContext) { s.vres = v }

func (s *Surface_approximation_techniqueContext) SetUvres(v IDecimalContext) { s.uvres = v }

func (s *Surface_approximation_techniqueContext) SetMaxlength(v IDecimalContext) { s.maxlength = v }

func (s *Surface_approximation_techniqueContext) SetMaxdist(v IDecimalContext) { s.maxdist = v }

func (s *Surface_approximation_techniqueContext) SetMaxangle(v IDecimalContext) { s.maxangle = v }

func (s *Surface_approximation_techniqueContext) AllDecimal() []IDecimalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalContext)(nil)).Elem())
	var tst = make([]IDecimalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalContext)
		}
	}

	return tst
}

func (s *Surface_approximation_techniqueContext) Decimal(i int) IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *Surface_approximation_techniqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Surface_approximation_techniqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Surface_approximation_techniqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterSurface_approximation_technique(s)
	}
}

func (s *Surface_approximation_techniqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitSurface_approximation_technique(s)
	}
}

func (p *WavefrontOBJParser) Surface_approximation_technique() (localctx ISurface_approximation_techniqueContext) {
	this := p
	_ = this

	localctx = NewSurface_approximation_techniqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, WavefrontOBJParserRULE_surface_approximation_technique)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(WavefrontOBJParserT__47)
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WavefrontOBJParserT__48:
		{
			p.SetState(402)
			p.Match(WavefrontOBJParserT__48)
		}
		{
			p.SetState(403)

			var _x = p.Decimal()

			localctx.(*Surface_approximation_techniqueContext).ures = _x
		}
		{
			p.SetState(404)

			var _x = p.Decimal()

			localctx.(*Surface_approximation_techniqueContext).vres = _x
		}

	case WavefrontOBJParserT__49:
		{
			p.SetState(406)
			p.Match(WavefrontOBJParserT__49)
		}
		{
			p.SetState(407)

			var _x = p.Decimal()

			localctx.(*Surface_approximation_techniqueContext).uvres = _x
		}

	case WavefrontOBJParserT__46:
		{
			p.SetState(408)
			p.Match(WavefrontOBJParserT__46)
		}
		{
			p.SetState(409)

			var _x = p.Decimal()

			localctx.(*Surface_approximation_techniqueContext).maxlength = _x
		}

	case WavefrontOBJParserT__18:
		{
			p.SetState(410)
			p.Match(WavefrontOBJParserT__18)
		}
		{
			p.SetState(411)

			var _x = p.Decimal()

			localctx.(*Surface_approximation_techniqueContext).maxdist = _x
		}
		{
			p.SetState(412)

			var _x = p.Decimal()

			localctx.(*Surface_approximation_techniqueContext).maxangle = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDecimalContext is an interface to support dynamic dispatch.
type IDecimalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalContext differentiates from other interfaces.
	IsDecimalContext()
}

type DecimalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalContext() *DecimalContext {
	var p = new(DecimalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_decimal
	return p
}

func (*DecimalContext) IsDecimalContext() {}

func NewDecimalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalContext {
	var p = new(DecimalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_decimal

	return p
}

func (s *DecimalContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserDECIMAL, 0)
}

func (s *DecimalContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, 0)
}

func (s *DecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterDecimal(s)
	}
}

func (s *DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitDecimal(s)
	}
}

func (p *WavefrontOBJParser) Decimal() (localctx IDecimalContext) {
	this := p
	_ = this

	localctx = NewDecimalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, WavefrontOBJParserRULE_decimal)
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
		p.SetState(416)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WavefrontOBJParserINTEGER || _la == WavefrontOBJParserDECIMAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
