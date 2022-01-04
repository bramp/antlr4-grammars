// Code generated from PeopleCode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package peoplecode // PeopleCode
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 85, 468,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 3, 2, 3, 3, 3, 3, 6, 3, 87, 10, 3,
	13, 3, 14, 3, 88, 7, 3, 91, 10, 3, 12, 3, 14, 3, 94, 11, 3, 3, 3, 5, 3,
	97, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	120, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 129, 10, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 145, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 150, 10, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 156, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	5, 5, 164, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 170, 10, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 176, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 193, 10, 5, 3, 5,
	7, 5, 196, 10, 5, 12, 5, 14, 5, 199, 11, 5, 3, 6, 3, 6, 3, 6, 7, 6, 204,
	10, 6, 12, 6, 14, 6, 207, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 214,
	10, 7, 12, 7, 14, 7, 217, 11, 7, 3, 8, 3, 8, 3, 8, 5, 8, 222, 10, 8, 3,
	9, 3, 9, 3, 9, 5, 9, 227, 10, 9, 3, 9, 5, 9, 230, 10, 9, 3, 10, 3, 10,
	3, 10, 5, 10, 235, 10, 10, 3, 11, 3, 11, 3, 11, 7, 11, 240, 10, 11, 12,
	11, 14, 11, 243, 11, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 6, 12,
	251, 10, 12, 13, 12, 14, 12, 252, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	7, 16, 271, 10, 16, 12, 16, 14, 16, 274, 11, 16, 3, 16, 3, 16, 3, 17, 5,
	17, 279, 10, 17, 3, 17, 3, 17, 7, 17, 283, 10, 17, 12, 17, 14, 17, 286,
	11, 17, 6, 17, 288, 10, 17, 13, 17, 14, 17, 289, 3, 18, 3, 18, 3, 18, 3,
	18, 5, 18, 296, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 302, 10, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 313,
	10, 21, 3, 21, 5, 21, 316, 10, 21, 3, 21, 5, 21, 319, 10, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 7, 22, 326, 10, 22, 12, 22, 14, 22, 329, 11, 22,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27,
	3, 27, 5, 27, 353, 10, 27, 3, 27, 5, 27, 356, 10, 27, 3, 27, 5, 27, 359,
	10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 365, 10, 28, 12, 28, 14, 28,
	368, 11, 28, 5, 28, 370, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5,
	29, 377, 10, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31,
	386, 10, 31, 3, 31, 3, 31, 3, 31, 5, 31, 391, 10, 31, 3, 31, 5, 31, 394,
	10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 5, 32, 407, 10, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3,
	33, 5, 33, 415, 10, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 6, 34,
	423, 10, 34, 13, 34, 14, 34, 424, 3, 34, 5, 34, 428, 10, 34, 3, 34, 3,
	34, 3, 35, 3, 35, 3, 35, 5, 35, 435, 10, 35, 3, 35, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 39, 3, 39, 3, 39, 5, 39, 456, 10, 39, 3, 39, 3, 39, 5, 39,
	460, 10, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 2, 3, 8,
	42, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
	74, 76, 78, 80, 2, 5, 3, 2, 37, 38, 3, 2, 74, 77, 3, 2, 78, 80, 2, 512,
	2, 82, 3, 2, 2, 2, 4, 92, 3, 2, 2, 2, 6, 128, 3, 2, 2, 2, 8, 144, 3, 2,
	2, 2, 10, 200, 3, 2, 2, 2, 12, 208, 3, 2, 2, 2, 14, 218, 3, 2, 2, 2, 16,
	229, 3, 2, 2, 2, 18, 231, 3, 2, 2, 2, 20, 236, 3, 2, 2, 2, 22, 247, 3,
	2, 2, 2, 24, 254, 3, 2, 2, 2, 26, 261, 3, 2, 2, 2, 28, 265, 3, 2, 2, 2,
	30, 267, 3, 2, 2, 2, 32, 278, 3, 2, 2, 2, 34, 295, 3, 2, 2, 2, 36, 297,
	3, 2, 2, 2, 38, 303, 3, 2, 2, 2, 40, 308, 3, 2, 2, 2, 42, 320, 3, 2, 2,
	2, 44, 330, 3, 2, 2, 2, 46, 335, 3, 2, 2, 2, 48, 340, 3, 2, 2, 2, 50, 345,
	3, 2, 2, 2, 52, 349, 3, 2, 2, 2, 54, 360, 3, 2, 2, 2, 56, 373, 3, 2, 2,
	2, 58, 378, 3, 2, 2, 2, 60, 381, 3, 2, 2, 2, 62, 397, 3, 2, 2, 2, 64, 411,
	3, 2, 2, 2, 66, 419, 3, 2, 2, 2, 68, 431, 3, 2, 2, 2, 70, 439, 3, 2, 2,
	2, 72, 442, 3, 2, 2, 2, 74, 448, 3, 2, 2, 2, 76, 452, 3, 2, 2, 2, 78, 463,
	3, 2, 2, 2, 80, 465, 3, 2, 2, 2, 82, 83, 5, 4, 3, 2, 83, 3, 3, 2, 2, 2,
	84, 86, 5, 6, 4, 2, 85, 87, 7, 3, 2, 2, 86, 85, 3, 2, 2, 2, 87, 88, 3,
	2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 91, 3, 2, 2, 2, 90,
	84, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2,
	2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 97, 5, 6, 4, 2, 96, 95,
	3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 5, 3, 2, 2, 2, 98, 129, 5, 18, 10,
	2, 99, 129, 5, 24, 13, 2, 100, 129, 5, 30, 16, 2, 101, 129, 5, 44, 23,
	2, 102, 129, 5, 46, 24, 2, 103, 129, 5, 48, 25, 2, 104, 129, 5, 50, 26,
	2, 105, 129, 5, 12, 7, 2, 106, 129, 5, 60, 31, 2, 107, 129, 5, 62, 32,
	2, 108, 129, 5, 64, 33, 2, 109, 129, 5, 66, 34, 2, 110, 129, 5, 72, 37,
	2, 111, 129, 7, 4, 2, 2, 112, 129, 7, 5, 2, 2, 113, 114, 7, 6, 2, 2, 114,
	129, 5, 8, 5, 2, 115, 116, 7, 7, 2, 2, 116, 129, 5, 8, 5, 2, 117, 119,
	7, 8, 2, 2, 118, 120, 5, 8, 5, 2, 119, 118, 3, 2, 2, 2, 119, 120, 3, 2,
	2, 2, 120, 129, 3, 2, 2, 2, 121, 122, 7, 9, 2, 2, 122, 129, 5, 8, 5, 2,
	123, 124, 5, 8, 5, 2, 124, 125, 7, 10, 2, 2, 125, 126, 5, 8, 5, 2, 126,
	129, 3, 2, 2, 2, 127, 129, 5, 8, 5, 2, 128, 98, 3, 2, 2, 2, 128, 99, 3,
	2, 2, 2, 128, 100, 3, 2, 2, 2, 128, 101, 3, 2, 2, 2, 128, 102, 3, 2, 2,
	2, 128, 103, 3, 2, 2, 2, 128, 104, 3, 2, 2, 2, 128, 105, 3, 2, 2, 2, 128,
	106, 3, 2, 2, 2, 128, 107, 3, 2, 2, 2, 128, 108, 3, 2, 2, 2, 128, 109,
	3, 2, 2, 2, 128, 110, 3, 2, 2, 2, 128, 111, 3, 2, 2, 2, 128, 112, 3, 2,
	2, 2, 128, 113, 3, 2, 2, 2, 128, 115, 3, 2, 2, 2, 128, 117, 3, 2, 2, 2,
	128, 121, 3, 2, 2, 2, 128, 123, 3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129,
	7, 3, 2, 2, 2, 130, 131, 8, 5, 1, 2, 131, 132, 7, 11, 2, 2, 132, 133, 5,
	8, 5, 2, 133, 134, 7, 12, 2, 2, 134, 145, 3, 2, 2, 2, 135, 136, 7, 13,
	2, 2, 136, 145, 5, 8, 5, 17, 137, 145, 5, 78, 40, 2, 138, 145, 5, 80, 41,
	2, 139, 145, 5, 76, 39, 2, 140, 141, 7, 17, 2, 2, 141, 145, 5, 8, 5, 10,
	142, 143, 7, 18, 2, 2, 143, 145, 5, 8, 5, 9, 144, 130, 3, 2, 2, 2, 144,
	135, 3, 2, 2, 2, 144, 137, 3, 2, 2, 2, 144, 138, 3, 2, 2, 2, 144, 139,
	3, 2, 2, 2, 144, 140, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 197, 3, 2,
	2, 2, 146, 149, 12, 8, 2, 2, 147, 150, 7, 19, 2, 2, 148, 150, 7, 20, 2,
	2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151,
	196, 5, 8, 5, 9, 152, 155, 12, 7, 2, 2, 153, 156, 7, 21, 2, 2, 154, 156,
	7, 17, 2, 2, 155, 153, 3, 2, 2, 2, 155, 154, 3, 2, 2, 2, 156, 157, 3, 2,
	2, 2, 157, 196, 5, 8, 5, 8, 158, 163, 12, 6, 2, 2, 159, 164, 7, 22, 2,
	2, 160, 164, 7, 23, 2, 2, 161, 164, 7, 24, 2, 2, 162, 164, 7, 25, 2, 2,
	163, 159, 3, 2, 2, 2, 163, 160, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163,
	162, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 196, 5, 8, 5, 7, 166, 169,
	12, 5, 2, 2, 167, 170, 7, 10, 2, 2, 168, 170, 7, 26, 2, 2, 169, 167, 3,
	2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 196, 5, 8, 5,
	6, 172, 175, 12, 4, 2, 2, 173, 176, 7, 27, 2, 2, 174, 176, 7, 28, 2, 2,
	175, 173, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177,
	196, 5, 8, 5, 5, 178, 179, 12, 3, 2, 2, 179, 180, 7, 29, 2, 2, 180, 196,
	5, 8, 5, 4, 181, 182, 12, 13, 2, 2, 182, 183, 7, 14, 2, 2, 183, 196, 5,
	80, 41, 2, 184, 185, 12, 12, 2, 2, 185, 186, 7, 15, 2, 2, 186, 187, 5,
	10, 6, 2, 187, 188, 7, 16, 2, 2, 188, 196, 3, 2, 2, 2, 189, 190, 12, 11,
	2, 2, 190, 192, 7, 11, 2, 2, 191, 193, 5, 10, 6, 2, 192, 191, 3, 2, 2,
	2, 192, 193, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 196, 7, 12, 2, 2, 195,
	146, 3, 2, 2, 2, 195, 152, 3, 2, 2, 2, 195, 158, 3, 2, 2, 2, 195, 166,
	3, 2, 2, 2, 195, 172, 3, 2, 2, 2, 195, 178, 3, 2, 2, 2, 195, 181, 3, 2,
	2, 2, 195, 184, 3, 2, 2, 2, 195, 189, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2,
	197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 9, 3, 2, 2, 2, 199, 197,
	3, 2, 2, 2, 200, 205, 5, 8, 5, 2, 201, 202, 7, 30, 2, 2, 202, 204, 5, 8,
	5, 2, 203, 201, 3, 2, 2, 2, 204, 207, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2,
	205, 206, 3, 2, 2, 2, 206, 11, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 208, 209,
	7, 80, 2, 2, 209, 210, 5, 16, 9, 2, 210, 215, 5, 14, 8, 2, 211, 212, 7,
	30, 2, 2, 212, 214, 5, 14, 8, 2, 213, 211, 3, 2, 2, 2, 214, 217, 3, 2,
	2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 13, 3, 2, 2, 2,
	217, 215, 3, 2, 2, 2, 218, 221, 7, 78, 2, 2, 219, 220, 7, 10, 2, 2, 220,
	222, 5, 8, 5, 2, 221, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 15, 3,
	2, 2, 2, 223, 226, 7, 80, 2, 2, 224, 225, 7, 31, 2, 2, 225, 227, 5, 16,
	9, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2,
	228, 230, 5, 22, 12, 2, 229, 223, 3, 2, 2, 2, 229, 228, 3, 2, 2, 2, 230,
	17, 3, 2, 2, 2, 231, 234, 7, 32, 2, 2, 232, 235, 5, 20, 11, 2, 233, 235,
	5, 22, 12, 2, 234, 232, 3, 2, 2, 2, 234, 233, 3, 2, 2, 2, 235, 19, 3, 2,
	2, 2, 236, 241, 7, 80, 2, 2, 237, 238, 7, 33, 2, 2, 238, 240, 7, 80, 2,
	2, 239, 237, 3, 2, 2, 2, 240, 243, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 241,
	242, 3, 2, 2, 2, 242, 244, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 244, 245,
	7, 33, 2, 2, 245, 246, 7, 19, 2, 2, 246, 21, 3, 2, 2, 2, 247, 250, 7, 80,
	2, 2, 248, 249, 7, 33, 2, 2, 249, 251, 7, 80, 2, 2, 250, 248, 3, 2, 2,
	2, 251, 252, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253,
	23, 3, 2, 2, 2, 254, 255, 7, 34, 2, 2, 255, 256, 7, 35, 2, 2, 256, 257,
	7, 80, 2, 2, 257, 258, 7, 36, 2, 2, 258, 259, 5, 26, 14, 2, 259, 260, 5,
	28, 15, 2, 260, 25, 3, 2, 2, 2, 261, 262, 7, 80, 2, 2, 262, 263, 7, 14,
	2, 2, 263, 264, 7, 80, 2, 2, 264, 27, 3, 2, 2, 2, 265, 266, 9, 2, 2, 2,
	266, 29, 3, 2, 2, 2, 267, 268, 7, 39, 2, 2, 268, 272, 7, 80, 2, 2, 269,
	271, 5, 32, 17, 2, 270, 269, 3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270,
	3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 275, 3, 2, 2, 2, 274, 272, 3, 2,
	2, 2, 275, 276, 7, 40, 2, 2, 276, 31, 3, 2, 2, 2, 277, 279, 7, 41, 2, 2,
	278, 277, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 287, 3, 2, 2, 2, 280,
	284, 5, 34, 18, 2, 281, 283, 7, 3, 2, 2, 282, 281, 3, 2, 2, 2, 283, 286,
	3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 288, 3, 2,
	2, 2, 286, 284, 3, 2, 2, 2, 287, 280, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2,
	289, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 33, 3, 2, 2, 2, 291, 296,
	5, 36, 19, 2, 292, 296, 5, 38, 20, 2, 293, 296, 5, 40, 21, 2, 294, 296,
	5, 42, 22, 2, 295, 291, 3, 2, 2, 2, 295, 292, 3, 2, 2, 2, 295, 293, 3,
	2, 2, 2, 295, 294, 3, 2, 2, 2, 296, 35, 3, 2, 2, 2, 297, 298, 7, 42, 2,
	2, 298, 299, 7, 80, 2, 2, 299, 301, 5, 54, 28, 2, 300, 302, 5, 58, 30,
	2, 301, 300, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 37, 3, 2, 2, 2, 303,
	304, 7, 43, 2, 2, 304, 305, 7, 78, 2, 2, 305, 306, 7, 10, 2, 2, 306, 307,
	5, 8, 5, 2, 307, 39, 3, 2, 2, 2, 308, 309, 7, 44, 2, 2, 309, 310, 5, 16,
	9, 2, 310, 312, 7, 80, 2, 2, 311, 313, 7, 45, 2, 2, 312, 311, 3, 2, 2,
	2, 312, 313, 3, 2, 2, 2, 313, 315, 3, 2, 2, 2, 314, 316, 7, 46, 2, 2, 315,
	314, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 318, 3, 2, 2, 2, 317, 319,
	7, 47, 2, 2, 318, 317, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 41, 3, 2,
	2, 2, 320, 321, 7, 48, 2, 2, 321, 322, 5, 16, 9, 2, 322, 327, 7, 78, 2,
	2, 323, 324, 7, 30, 2, 2, 324, 326, 7, 78, 2, 2, 325, 323, 3, 2, 2, 2,
	326, 329, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328,
	43, 3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 330, 331, 7, 42, 2, 2, 331, 332,
	7, 80, 2, 2, 332, 333, 5, 4, 3, 2, 333, 334, 7, 49, 2, 2, 334, 45, 3, 2,
	2, 2, 335, 336, 7, 45, 2, 2, 336, 337, 7, 80, 2, 2, 337, 338, 5, 4, 3,
	2, 338, 339, 7, 50, 2, 2, 339, 47, 3, 2, 2, 2, 340, 341, 7, 46, 2, 2, 341,
	342, 7, 80, 2, 2, 342, 343, 5, 4, 3, 2, 343, 344, 7, 51, 2, 2, 344, 49,
	3, 2, 2, 2, 345, 346, 5, 52, 27, 2, 346, 347, 5, 4, 3, 2, 347, 348, 7,
	52, 2, 2, 348, 51, 3, 2, 2, 2, 349, 350, 7, 35, 2, 2, 350, 352, 7, 80,
	2, 2, 351, 353, 5, 54, 28, 2, 352, 351, 3, 2, 2, 2, 352, 353, 3, 2, 2,
	2, 353, 355, 3, 2, 2, 2, 354, 356, 5, 58, 30, 2, 355, 354, 3, 2, 2, 2,
	355, 356, 3, 2, 2, 2, 356, 358, 3, 2, 2, 2, 357, 359, 7, 3, 2, 2, 358,
	357, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 53, 3, 2, 2, 2, 360, 369, 7,
	11, 2, 2, 361, 366, 5, 56, 29, 2, 362, 363, 7, 30, 2, 2, 363, 365, 5, 56,
	29, 2, 364, 362, 3, 2, 2, 2, 365, 368, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2,
	366, 367, 3, 2, 2, 2, 367, 370, 3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 369,
	361, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 372,
	7, 12, 2, 2, 372, 55, 3, 2, 2, 2, 373, 376, 7, 78, 2, 2, 374, 375, 7, 53,
	2, 2, 375, 377, 5, 16, 9, 2, 376, 374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2,
	377, 57, 3, 2, 2, 2, 378, 379, 7, 54, 2, 2, 379, 380, 5, 16, 9, 2, 380,
	59, 3, 2, 2, 2, 381, 382, 7, 55, 2, 2, 382, 383, 5, 8, 5, 2, 383, 385,
	7, 56, 2, 2, 384, 386, 7, 3, 2, 2, 385, 384, 3, 2, 2, 2, 385, 386, 3, 2,
	2, 2, 386, 387, 3, 2, 2, 2, 387, 393, 5, 4, 3, 2, 388, 390, 7, 57, 2, 2,
	389, 391, 7, 3, 2, 2, 390, 389, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391,
	392, 3, 2, 2, 2, 392, 394, 5, 4, 3, 2, 393, 388, 3, 2, 2, 2, 393, 394,
	3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395, 396, 7, 58, 2, 2, 396, 61, 3, 2,
	2, 2, 397, 398, 7, 59, 2, 2, 398, 399, 7, 78, 2, 2, 399, 400, 7, 10, 2,
	2, 400, 401, 5, 8, 5, 2, 401, 402, 7, 60, 2, 2, 402, 406, 5, 8, 5, 2, 403,
	407, 7, 3, 2, 2, 404, 405, 7, 61, 2, 2, 405, 407, 5, 8, 5, 2, 406, 403,
	3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 408, 3, 2,
	2, 2, 408, 409, 5, 4, 3, 2, 409, 410, 7, 62, 2, 2, 410, 63, 3, 2, 2, 2,
	411, 412, 7, 63, 2, 2, 412, 414, 5, 8, 5, 2, 413, 415, 7, 3, 2, 2, 414,
	413, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417,
	5, 4, 3, 2, 417, 418, 7, 64, 2, 2, 418, 65, 3, 2, 2, 2, 419, 420, 7, 65,
	2, 2, 420, 422, 5, 8, 5, 2, 421, 423, 5, 68, 35, 2, 422, 421, 3, 2, 2,
	2, 423, 424, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425,
	427, 3, 2, 2, 2, 426, 428, 5, 70, 36, 2, 427, 426, 3, 2, 2, 2, 427, 428,
	3, 2, 2, 2, 428, 429, 3, 2, 2, 2, 429, 430, 7, 66, 2, 2, 430, 67, 3, 2,
	2, 2, 431, 434, 7, 67, 2, 2, 432, 435, 7, 10, 2, 2, 433, 435, 7, 25, 2,
	2, 434, 432, 3, 2, 2, 2, 434, 433, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435,
	436, 3, 2, 2, 2, 436, 437, 5, 8, 5, 2, 437, 438, 5, 4, 3, 2, 438, 69, 3,
	2, 2, 2, 439, 440, 7, 68, 2, 2, 440, 441, 5, 4, 3, 2, 441, 71, 3, 2, 2,
	2, 442, 443, 7, 69, 2, 2, 443, 444, 5, 4, 3, 2, 444, 445, 5, 74, 38, 2,
	445, 446, 5, 4, 3, 2, 446, 447, 7, 70, 2, 2, 447, 73, 3, 2, 2, 2, 448,
	449, 7, 71, 2, 2, 449, 450, 7, 72, 2, 2, 450, 451, 7, 78, 2, 2, 451, 75,
	3, 2, 2, 2, 452, 455, 7, 73, 2, 2, 453, 456, 5, 22, 12, 2, 454, 456, 7,
	80, 2, 2, 455, 453, 3, 2, 2, 2, 455, 454, 3, 2, 2, 2, 456, 457, 3, 2, 2,
	2, 457, 459, 7, 11, 2, 2, 458, 460, 5, 10, 6, 2, 459, 458, 3, 2, 2, 2,
	459, 460, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461, 462, 7, 12, 2, 2, 462,
	77, 3, 2, 2, 2, 463, 464, 9, 3, 2, 2, 464, 79, 3, 2, 2, 2, 465, 466, 9,
	4, 2, 2, 466, 81, 3, 2, 2, 2, 50, 88, 92, 96, 119, 128, 144, 149, 155,
	163, 169, 175, 192, 195, 197, 205, 215, 221, 226, 229, 234, 241, 252, 272,
	278, 284, 289, 295, 301, 312, 315, 318, 327, 352, 355, 358, 366, 369, 376,
	385, 390, 393, 406, 414, 424, 427, 434, 455, 459,
}
var literalNames = []string{
	"", "';'", "'Exit'", "'Break'", "'Error'", "'Warning'", "'Return'", "'throw'",
	"'='", "'('", "')'", "'@'", "'.'", "'['", "']'", "'-'", "'Not'", "'*'",
	"'/'", "'+'", "'<='", "'>='", "'<'", "'>'", "'<>'", "'And'", "'Or'", "'|'",
	"','", "'of'", "'import'", "':'", "'Declare'", "'Function'", "'PeopleCode'",
	"'FieldFormula'", "'FieldChange'", "'class'", "'end-class'", "'private'",
	"'method'", "'Constant'", "'property'", "'get'", "'set'", "'readonly'",
	"'instance'", "'end-method'", "'end-get'", "'end-set'", "'End-Function'",
	"'As'", "'Returns'", "'If'", "'Then'", "'Else'", "'End-If'", "'For'", "'To'",
	"'Step'", "'End-For'", "'While'", "'End-While'", "'Evaluate'", "'End-Evaluate'",
	"'When'", "'When-Other'", "'try'", "'end-try'", "'catch'", "'Exception'",
	"'create'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"DecimalLiteral", "IntegerLiteral", "StringLiteral", "BoolLiteral", "VAR_ID",
	"SYS_VAR_ID", "GENERIC_ID", "REM", "COMMENT_1", "COMMENT_2", "COMMENT_3",
	"WS",
}

var ruleNames = []string{
	"program", "stmtList", "stmt", "expr", "exprList", "varDeclaration", "varDeclarator",
	"varType", "appClassImport", "appPkgPath", "appClassPath", "extFuncImport",
	"recDefnPath", "event", "classDeclaration", "classBlock", "classBlockStmt",
	"method", "constant", "property_", "instance", "methodImpl", "getImpl",
	"setImpl", "funcImpl", "funcSignature", "formalParamList", "param", "returnType",
	"ifStmt", "forStmt", "whileStmt", "evaluateStmt", "whenBranch", "whenOtherBranch",
	"tryCatchStmt", "catchSignature", "createInvocation", "literal", "id_",
}

type PeopleCodeParser struct {
	*antlr.BaseParser
}

// NewPeopleCodeParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *PeopleCodeParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPeopleCodeParser(input antlr.TokenStream) *PeopleCodeParser {
	this := new(PeopleCodeParser)
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
	this.GrammarFileName = "PeopleCode.g4"

	return this
}

// PeopleCodeParser tokens.
const (
	PeopleCodeParserEOF            = antlr.TokenEOF
	PeopleCodeParserT__0           = 1
	PeopleCodeParserT__1           = 2
	PeopleCodeParserT__2           = 3
	PeopleCodeParserT__3           = 4
	PeopleCodeParserT__4           = 5
	PeopleCodeParserT__5           = 6
	PeopleCodeParserT__6           = 7
	PeopleCodeParserT__7           = 8
	PeopleCodeParserT__8           = 9
	PeopleCodeParserT__9           = 10
	PeopleCodeParserT__10          = 11
	PeopleCodeParserT__11          = 12
	PeopleCodeParserT__12          = 13
	PeopleCodeParserT__13          = 14
	PeopleCodeParserT__14          = 15
	PeopleCodeParserT__15          = 16
	PeopleCodeParserT__16          = 17
	PeopleCodeParserT__17          = 18
	PeopleCodeParserT__18          = 19
	PeopleCodeParserT__19          = 20
	PeopleCodeParserT__20          = 21
	PeopleCodeParserT__21          = 22
	PeopleCodeParserT__22          = 23
	PeopleCodeParserT__23          = 24
	PeopleCodeParserT__24          = 25
	PeopleCodeParserT__25          = 26
	PeopleCodeParserT__26          = 27
	PeopleCodeParserT__27          = 28
	PeopleCodeParserT__28          = 29
	PeopleCodeParserT__29          = 30
	PeopleCodeParserT__30          = 31
	PeopleCodeParserT__31          = 32
	PeopleCodeParserT__32          = 33
	PeopleCodeParserT__33          = 34
	PeopleCodeParserT__34          = 35
	PeopleCodeParserT__35          = 36
	PeopleCodeParserT__36          = 37
	PeopleCodeParserT__37          = 38
	PeopleCodeParserT__38          = 39
	PeopleCodeParserT__39          = 40
	PeopleCodeParserT__40          = 41
	PeopleCodeParserT__41          = 42
	PeopleCodeParserT__42          = 43
	PeopleCodeParserT__43          = 44
	PeopleCodeParserT__44          = 45
	PeopleCodeParserT__45          = 46
	PeopleCodeParserT__46          = 47
	PeopleCodeParserT__47          = 48
	PeopleCodeParserT__48          = 49
	PeopleCodeParserT__49          = 50
	PeopleCodeParserT__50          = 51
	PeopleCodeParserT__51          = 52
	PeopleCodeParserT__52          = 53
	PeopleCodeParserT__53          = 54
	PeopleCodeParserT__54          = 55
	PeopleCodeParserT__55          = 56
	PeopleCodeParserT__56          = 57
	PeopleCodeParserT__57          = 58
	PeopleCodeParserT__58          = 59
	PeopleCodeParserT__59          = 60
	PeopleCodeParserT__60          = 61
	PeopleCodeParserT__61          = 62
	PeopleCodeParserT__62          = 63
	PeopleCodeParserT__63          = 64
	PeopleCodeParserT__64          = 65
	PeopleCodeParserT__65          = 66
	PeopleCodeParserT__66          = 67
	PeopleCodeParserT__67          = 68
	PeopleCodeParserT__68          = 69
	PeopleCodeParserT__69          = 70
	PeopleCodeParserT__70          = 71
	PeopleCodeParserDecimalLiteral = 72
	PeopleCodeParserIntegerLiteral = 73
	PeopleCodeParserStringLiteral  = 74
	PeopleCodeParserBoolLiteral    = 75
	PeopleCodeParserVAR_ID         = 76
	PeopleCodeParserSYS_VAR_ID     = 77
	PeopleCodeParserGENERIC_ID     = 78
	PeopleCodeParserREM            = 79
	PeopleCodeParserCOMMENT_1      = 80
	PeopleCodeParserCOMMENT_2      = 81
	PeopleCodeParserCOMMENT_3      = 82
	PeopleCodeParserWS             = 83
)

// PeopleCodeParser rules.
const (
	PeopleCodeParserRULE_program          = 0
	PeopleCodeParserRULE_stmtList         = 1
	PeopleCodeParserRULE_stmt             = 2
	PeopleCodeParserRULE_expr             = 3
	PeopleCodeParserRULE_exprList         = 4
	PeopleCodeParserRULE_varDeclaration   = 5
	PeopleCodeParserRULE_varDeclarator    = 6
	PeopleCodeParserRULE_varType          = 7
	PeopleCodeParserRULE_appClassImport   = 8
	PeopleCodeParserRULE_appPkgPath       = 9
	PeopleCodeParserRULE_appClassPath     = 10
	PeopleCodeParserRULE_extFuncImport    = 11
	PeopleCodeParserRULE_recDefnPath      = 12
	PeopleCodeParserRULE_event            = 13
	PeopleCodeParserRULE_classDeclaration = 14
	PeopleCodeParserRULE_classBlock       = 15
	PeopleCodeParserRULE_classBlockStmt   = 16
	PeopleCodeParserRULE_method           = 17
	PeopleCodeParserRULE_constant         = 18
	PeopleCodeParserRULE_property_        = 19
	PeopleCodeParserRULE_instance         = 20
	PeopleCodeParserRULE_methodImpl       = 21
	PeopleCodeParserRULE_getImpl          = 22
	PeopleCodeParserRULE_setImpl          = 23
	PeopleCodeParserRULE_funcImpl         = 24
	PeopleCodeParserRULE_funcSignature    = 25
	PeopleCodeParserRULE_formalParamList  = 26
	PeopleCodeParserRULE_param            = 27
	PeopleCodeParserRULE_returnType       = 28
	PeopleCodeParserRULE_ifStmt           = 29
	PeopleCodeParserRULE_forStmt          = 30
	PeopleCodeParserRULE_whileStmt        = 31
	PeopleCodeParserRULE_evaluateStmt     = 32
	PeopleCodeParserRULE_whenBranch       = 33
	PeopleCodeParserRULE_whenOtherBranch  = 34
	PeopleCodeParserRULE_tryCatchStmt     = 35
	PeopleCodeParserRULE_catchSignature   = 36
	PeopleCodeParserRULE_createInvocation = 37
	PeopleCodeParserRULE_literal          = 38
	PeopleCodeParserRULE_id_              = 39
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
	p.RuleIndex = PeopleCodeParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) StmtList() IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *PeopleCodeParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PeopleCodeParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.StmtList()
	}

	return localctx
}

// IStmtListContext is an interface to support dynamic dispatch.
type IStmtListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtListContext differentiates from other interfaces.
	IsStmtListContext()
}

type StmtListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtListContext() *StmtListContext {
	var p = new(StmtListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_stmtList
	return p
}

func (*StmtListContext) IsStmtListContext() {}

func NewStmtListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtListContext {
	var p = new(StmtListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_stmtList

	return p
}

func (s *StmtListContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtListContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *StmtListContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StmtListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtList(s)
	}
}

func (s *StmtListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtList(s)
	}
}

func (p *PeopleCodeParser) StmtList() (localctx IStmtListContext) {
	this := p
	_ = this

	localctx = NewStmtListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PeopleCodeParserRULE_stmtList)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(82)
				p.Stmt()
			}
			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == PeopleCodeParserT__0 {
				{
					p.SetState(83)
					p.Match(PeopleCodeParserT__0)
				}

				p.SetState(86)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PeopleCodeParserT__1)|(1<<PeopleCodeParserT__2)|(1<<PeopleCodeParserT__3)|(1<<PeopleCodeParserT__4)|(1<<PeopleCodeParserT__5)|(1<<PeopleCodeParserT__6)|(1<<PeopleCodeParserT__8)|(1<<PeopleCodeParserT__10)|(1<<PeopleCodeParserT__14)|(1<<PeopleCodeParserT__15)|(1<<PeopleCodeParserT__29))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(PeopleCodeParserT__31-32))|(1<<(PeopleCodeParserT__32-32))|(1<<(PeopleCodeParserT__36-32))|(1<<(PeopleCodeParserT__39-32))|(1<<(PeopleCodeParserT__42-32))|(1<<(PeopleCodeParserT__43-32))|(1<<(PeopleCodeParserT__52-32))|(1<<(PeopleCodeParserT__56-32))|(1<<(PeopleCodeParserT__60-32))|(1<<(PeopleCodeParserT__62-32)))) != 0) || (((_la-67)&-(0x1f+1)) == 0 && ((1<<uint((_la-67)))&((1<<(PeopleCodeParserT__66-67))|(1<<(PeopleCodeParserT__70-67))|(1<<(PeopleCodeParserDecimalLiteral-67))|(1<<(PeopleCodeParserIntegerLiteral-67))|(1<<(PeopleCodeParserStringLiteral-67))|(1<<(PeopleCodeParserBoolLiteral-67))|(1<<(PeopleCodeParserVAR_ID-67))|(1<<(PeopleCodeParserSYS_VAR_ID-67))|(1<<(PeopleCodeParserGENERIC_ID-67)))) != 0) {
		{
			p.SetState(93)
			p.Stmt()
		}

	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyFrom(ctx *StmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StmtReturnContext struct {
	*StmtContext
}

func NewStmtReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtReturnContext {
	var p = new(StmtReturnContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtReturnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtReturn(s)
	}
}

func (s *StmtReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtReturn(s)
	}
}

type StmtEvaluateContext struct {
	*StmtContext
}

func NewStmtEvaluateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtEvaluateContext {
	var p = new(StmtEvaluateContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtEvaluateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtEvaluateContext) EvaluateStmt() IEvaluateStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEvaluateStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEvaluateStmtContext)
}

func (s *StmtEvaluateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtEvaluate(s)
	}
}

func (s *StmtEvaluateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtEvaluate(s)
	}
}

type StmtAppClassImportContext struct {
	*StmtContext
}

func NewStmtAppClassImportContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtAppClassImportContext {
	var p = new(StmtAppClassImportContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtAppClassImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtAppClassImportContext) AppClassImport() IAppClassImportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAppClassImportContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAppClassImportContext)
}

func (s *StmtAppClassImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtAppClassImport(s)
	}
}

func (s *StmtAppClassImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtAppClassImport(s)
	}
}

type StmtWarningContext struct {
	*StmtContext
}

func NewStmtWarningContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtWarningContext {
	var p = new(StmtWarningContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtWarningContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtWarningContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtWarningContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtWarning(s)
	}
}

func (s *StmtWarningContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtWarning(s)
	}
}

type StmtWhileContext struct {
	*StmtContext
}

func NewStmtWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtWhileContext {
	var p = new(StmtWhileContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtWhileContext) WhileStmt() IWhileStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StmtWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtWhile(s)
	}
}

func (s *StmtWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtWhile(s)
	}
}

type StmtExprContext struct {
	*StmtContext
}

func NewStmtExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtExprContext {
	var p = new(StmtExprContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtExpr(s)
	}
}

func (s *StmtExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtExpr(s)
	}
}

type StmtSetImplContext struct {
	*StmtContext
}

func NewStmtSetImplContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtSetImplContext {
	var p = new(StmtSetImplContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtSetImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtSetImplContext) SetImpl() ISetImplContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetImplContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetImplContext)
}

func (s *StmtSetImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtSetImpl(s)
	}
}

func (s *StmtSetImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtSetImpl(s)
	}
}

type StmtMethodImplContext struct {
	*StmtContext
}

func NewStmtMethodImplContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtMethodImplContext {
	var p = new(StmtMethodImplContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtMethodImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtMethodImplContext) MethodImpl() IMethodImplContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodImplContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodImplContext)
}

func (s *StmtMethodImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtMethodImpl(s)
	}
}

func (s *StmtMethodImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtMethodImpl(s)
	}
}

type StmtTryCatchContext struct {
	*StmtContext
}

func NewStmtTryCatchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtTryCatchContext {
	var p = new(StmtTryCatchContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtTryCatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtTryCatchContext) TryCatchStmt() ITryCatchStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITryCatchStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITryCatchStmtContext)
}

func (s *StmtTryCatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtTryCatch(s)
	}
}

func (s *StmtTryCatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtTryCatch(s)
	}
}

type StmtVarDeclarationContext struct {
	*StmtContext
}

func NewStmtVarDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtVarDeclarationContext {
	var p = new(StmtVarDeclarationContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtVarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtVarDeclarationContext) VarDeclaration() IVarDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclarationContext)
}

func (s *StmtVarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtVarDeclaration(s)
	}
}

func (s *StmtVarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtVarDeclaration(s)
	}
}

type StmtForContext struct {
	*StmtContext
}

func NewStmtForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtForContext {
	var p = new(StmtForContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtForContext) ForStmt() IForStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StmtForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtFor(s)
	}
}

func (s *StmtForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtFor(s)
	}
}

type StmtAssignContext struct {
	*StmtContext
}

func NewStmtAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtAssignContext {
	var p = new(StmtAssignContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtAssignContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *StmtAssignContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtAssign(s)
	}
}

func (s *StmtAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtAssign(s)
	}
}

type StmtErrorContext struct {
	*StmtContext
}

func NewStmtErrorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtErrorContext {
	var p = new(StmtErrorContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtErrorContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtError(s)
	}
}

func (s *StmtErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtError(s)
	}
}

type StmtExternalFuncImportContext struct {
	*StmtContext
}

func NewStmtExternalFuncImportContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtExternalFuncImportContext {
	var p = new(StmtExternalFuncImportContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtExternalFuncImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtExternalFuncImportContext) ExtFuncImport() IExtFuncImportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtFuncImportContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtFuncImportContext)
}

func (s *StmtExternalFuncImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtExternalFuncImport(s)
	}
}

func (s *StmtExternalFuncImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtExternalFuncImport(s)
	}
}

type StmtFuncImplContext struct {
	*StmtContext
}

func NewStmtFuncImplContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtFuncImplContext {
	var p = new(StmtFuncImplContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtFuncImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtFuncImplContext) FuncImpl() IFuncImplContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncImplContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncImplContext)
}

func (s *StmtFuncImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtFuncImpl(s)
	}
}

func (s *StmtFuncImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtFuncImpl(s)
	}
}

type StmtGetImplContext struct {
	*StmtContext
}

func NewStmtGetImplContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtGetImplContext {
	var p = new(StmtGetImplContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtGetImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtGetImplContext) GetImpl() IGetImplContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetImplContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetImplContext)
}

func (s *StmtGetImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtGetImpl(s)
	}
}

func (s *StmtGetImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtGetImpl(s)
	}
}

type StmtThrowContext struct {
	*StmtContext
}

func NewStmtThrowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtThrowContext {
	var p = new(StmtThrowContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtThrowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtThrowContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtThrowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtThrow(s)
	}
}

func (s *StmtThrowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtThrow(s)
	}
}

type StmtIfContext struct {
	*StmtContext
}

func NewStmtIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtIfContext {
	var p = new(StmtIfContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtIfContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtIf(s)
	}
}

func (s *StmtIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtIf(s)
	}
}

type StmtClassDeclarationContext struct {
	*StmtContext
}

func NewStmtClassDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtClassDeclarationContext {
	var p = new(StmtClassDeclarationContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtClassDeclarationContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *StmtClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtClassDeclaration(s)
	}
}

func (s *StmtClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtClassDeclaration(s)
	}
}

type StmtExitContext struct {
	*StmtContext
}

func NewStmtExitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtExitContext {
	var p = new(StmtExitContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtExitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtExitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtExit(s)
	}
}

func (s *StmtExitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtExit(s)
	}
}

type StmtBreakContext struct {
	*StmtContext
}

func NewStmtBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtBreakContext {
	var p = new(StmtBreakContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *StmtBreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtBreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterStmtBreak(s)
	}
}

func (s *StmtBreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitStmtBreak(s)
	}
}

func (p *PeopleCodeParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PeopleCodeParserRULE_stmt)
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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStmtAppClassImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.AppClassImport()
		}

	case 2:
		localctx = NewStmtExternalFuncImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.ExtFuncImport()
		}

	case 3:
		localctx = NewStmtClassDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.ClassDeclaration()
		}

	case 4:
		localctx = NewStmtMethodImplContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(99)
			p.MethodImpl()
		}

	case 5:
		localctx = NewStmtGetImplContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(100)
			p.GetImpl()
		}

	case 6:
		localctx = NewStmtSetImplContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(101)
			p.SetImpl()
		}

	case 7:
		localctx = NewStmtFuncImplContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(102)
			p.FuncImpl()
		}

	case 8:
		localctx = NewStmtVarDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(103)
			p.VarDeclaration()
		}

	case 9:
		localctx = NewStmtIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(104)
			p.IfStmt()
		}

	case 10:
		localctx = NewStmtForContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(105)
			p.ForStmt()
		}

	case 11:
		localctx = NewStmtWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(106)
			p.WhileStmt()
		}

	case 12:
		localctx = NewStmtEvaluateContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(107)
			p.EvaluateStmt()
		}

	case 13:
		localctx = NewStmtTryCatchContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(108)
			p.TryCatchStmt()
		}

	case 14:
		localctx = NewStmtExitContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(109)
			p.Match(PeopleCodeParserT__1)
		}

	case 15:
		localctx = NewStmtBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(110)
			p.Match(PeopleCodeParserT__2)
		}

	case 16:
		localctx = NewStmtErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(111)
			p.Match(PeopleCodeParserT__3)
		}
		{
			p.SetState(112)
			p.expr(0)
		}

	case 17:
		localctx = NewStmtWarningContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(113)
			p.Match(PeopleCodeParserT__4)
		}
		{
			p.SetState(114)
			p.expr(0)
		}

	case 18:
		localctx = NewStmtReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(115)
			p.Match(PeopleCodeParserT__5)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PeopleCodeParserT__8)|(1<<PeopleCodeParserT__10)|(1<<PeopleCodeParserT__14)|(1<<PeopleCodeParserT__15))) != 0) || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(PeopleCodeParserT__70-71))|(1<<(PeopleCodeParserDecimalLiteral-71))|(1<<(PeopleCodeParserIntegerLiteral-71))|(1<<(PeopleCodeParserStringLiteral-71))|(1<<(PeopleCodeParserBoolLiteral-71))|(1<<(PeopleCodeParserVAR_ID-71))|(1<<(PeopleCodeParserSYS_VAR_ID-71))|(1<<(PeopleCodeParserGENERIC_ID-71)))) != 0) {
			{
				p.SetState(116)
				p.expr(0)
			}

		}

	case 19:
		localctx = NewStmtThrowContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(119)
			p.Match(PeopleCodeParserT__6)
		}
		{
			p.SetState(120)
			p.expr(0)
		}

	case 20:
		localctx = NewStmtAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(121)
			p.expr(0)
		}
		{
			p.SetState(122)
			p.Match(PeopleCodeParserT__7)
		}
		{
			p.SetState(123)
			p.expr(0)
		}

	case 21:
		localctx = NewStmtExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(125)
			p.expr(0)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprComparisonContext struct {
	*ExprContext
	le antlr.Token
	ge antlr.Token
	l  antlr.Token
	g  antlr.Token
}

func NewExprComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprComparisonContext {
	var p = new(ExprComparisonContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprComparisonContext) GetLe() antlr.Token { return s.le }

func (s *ExprComparisonContext) GetGe() antlr.Token { return s.ge }

func (s *ExprComparisonContext) GetL() antlr.Token { return s.l }

func (s *ExprComparisonContext) GetG() antlr.Token { return s.g }

func (s *ExprComparisonContext) SetLe(v antlr.Token) { s.le = v }

func (s *ExprComparisonContext) SetGe(v antlr.Token) { s.ge = v }

func (s *ExprComparisonContext) SetL(v antlr.Token) { s.l = v }

func (s *ExprComparisonContext) SetG(v antlr.Token) { s.g = v }

func (s *ExprComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprComparisonContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprComparisonContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprComparison(s)
	}
}

func (s *ExprComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprComparison(s)
	}
}

type ExprConcatContext struct {
	*ExprContext
}

func NewExprConcatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprConcatContext {
	var p = new(ExprConcatContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprConcatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprConcatContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprConcatContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprConcatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprConcat(s)
	}
}

func (s *ExprConcatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprConcat(s)
	}
}

type ExprCreateContext struct {
	*ExprContext
}

func NewExprCreateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprCreateContext {
	var p = new(ExprCreateContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprCreateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprCreateContext) CreateInvocation() ICreateInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateInvocationContext)
}

func (s *ExprCreateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprCreate(s)
	}
}

func (s *ExprCreateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprCreate(s)
	}
}

type ExprBooleanContext struct {
	*ExprContext
	op antlr.Token
}

func NewExprBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprBooleanContext {
	var p = new(ExprBooleanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprBooleanContext) GetOp() antlr.Token { return s.op }

func (s *ExprBooleanContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBooleanContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprBooleanContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprBoolean(s)
	}
}

func (s *ExprBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprBoolean(s)
	}
}

type ExprNotContext struct {
	*ExprContext
}

func NewExprNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNotContext {
	var p = new(ExprNotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNotContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprNot(s)
	}
}

func (s *ExprNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprNot(s)
	}
}

type ExprAddSubContext struct {
	*ExprContext
	a antlr.Token
	s antlr.Token
}

func NewExprAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAddSubContext {
	var p = new(ExprAddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprAddSubContext) GetA() antlr.Token { return s.a }

func (s *ExprAddSubContext) GetS() antlr.Token { return s.s }

func (s *ExprAddSubContext) SetA(v antlr.Token) { s.a = v }

func (s *ExprAddSubContext) SetS(v antlr.Token) { s.s = v }

func (s *ExprAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAddSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprAddSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprAddSub(s)
	}
}

func (s *ExprAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprAddSub(s)
	}
}

type ExprDotAccessContext struct {
	*ExprContext
}

func NewExprDotAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprDotAccessContext {
	var p = new(ExprDotAccessContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprDotAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprDotAccessContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprDotAccessContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *ExprDotAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprDotAccess(s)
	}
}

func (s *ExprDotAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprDotAccess(s)
	}
}

type ExprFnOrIdxCallContext struct {
	*ExprContext
}

func NewExprFnOrIdxCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFnOrIdxCallContext {
	var p = new(ExprFnOrIdxCallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprFnOrIdxCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFnOrIdxCallContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprFnOrIdxCallContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ExprFnOrIdxCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprFnOrIdxCall(s)
	}
}

func (s *ExprFnOrIdxCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprFnOrIdxCall(s)
	}
}

type ExprParenthesizedContext struct {
	*ExprContext
}

func NewExprParenthesizedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprParenthesizedContext {
	var p = new(ExprParenthesizedContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprParenthesizedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprParenthesizedContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprParenthesizedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprParenthesized(s)
	}
}

func (s *ExprParenthesizedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprParenthesized(s)
	}
}

type ExprMulDivContext struct {
	*ExprContext
	m antlr.Token
	d antlr.Token
}

func NewExprMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprMulDivContext {
	var p = new(ExprMulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprMulDivContext) GetM() antlr.Token { return s.m }

func (s *ExprMulDivContext) GetD() antlr.Token { return s.d }

func (s *ExprMulDivContext) SetM(v antlr.Token) { s.m = v }

func (s *ExprMulDivContext) SetD(v antlr.Token) { s.d = v }

func (s *ExprMulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMulDivContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprMulDivContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprMulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprMulDiv(s)
	}
}

func (s *ExprMulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprMulDiv(s)
	}
}

type ExprNegateContext struct {
	*ExprContext
}

func NewExprNegateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNegateContext {
	var p = new(ExprNegateContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprNegateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNegateContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprNegateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprNegate(s)
	}
}

func (s *ExprNegateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprNegate(s)
	}
}

type ExprArrayIndexContext struct {
	*ExprContext
}

func NewExprArrayIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprArrayIndexContext {
	var p = new(ExprArrayIndexContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprArrayIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprArrayIndexContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprArrayIndexContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ExprArrayIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprArrayIndex(s)
	}
}

func (s *ExprArrayIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprArrayIndex(s)
	}
}

type ExprLiteralContext struct {
	*ExprContext
}

func NewExprLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLiteralContext {
	var p = new(ExprLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLiteralContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprLiteral(s)
	}
}

func (s *ExprLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprLiteral(s)
	}
}

type ExprEqualityContext struct {
	*ExprContext
	e antlr.Token
	i antlr.Token
}

func NewExprEqualityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprEqualityContext {
	var p = new(ExprEqualityContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprEqualityContext) GetE() antlr.Token { return s.e }

func (s *ExprEqualityContext) GetI() antlr.Token { return s.i }

func (s *ExprEqualityContext) SetE(v antlr.Token) { s.e = v }

func (s *ExprEqualityContext) SetI(v antlr.Token) { s.i = v }

func (s *ExprEqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprEqualityContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprEqualityContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprEqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprEquality(s)
	}
}

func (s *ExprEqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprEquality(s)
	}
}

type ExprDynamicReferenceContext struct {
	*ExprContext
}

func NewExprDynamicReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprDynamicReferenceContext {
	var p = new(ExprDynamicReferenceContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprDynamicReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprDynamicReferenceContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprDynamicReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprDynamicReference(s)
	}
}

func (s *ExprDynamicReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprDynamicReference(s)
	}
}

type ExprIdContext struct {
	*ExprContext
}

func NewExprIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIdContext {
	var p = new(ExprIdContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIdContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *ExprIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprId(s)
	}
}

func (s *ExprIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprId(s)
	}
}

func (p *PeopleCodeParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PeopleCodeParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, PeopleCodeParserRULE_expr, _p)
	var _la int

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
	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PeopleCodeParserT__8:
		localctx = NewExprParenthesizedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(129)
			p.Match(PeopleCodeParserT__8)
		}
		{
			p.SetState(130)
			p.expr(0)
		}
		{
			p.SetState(131)
			p.Match(PeopleCodeParserT__9)
		}

	case PeopleCodeParserT__10:
		localctx = NewExprDynamicReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(133)
			p.Match(PeopleCodeParserT__10)
		}
		{
			p.SetState(134)
			p.expr(15)
		}

	case PeopleCodeParserDecimalLiteral, PeopleCodeParserIntegerLiteral, PeopleCodeParserStringLiteral, PeopleCodeParserBoolLiteral:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(135)
			p.Literal()
		}

	case PeopleCodeParserVAR_ID, PeopleCodeParserSYS_VAR_ID, PeopleCodeParserGENERIC_ID:
		localctx = NewExprIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(136)
			p.Id_()
		}

	case PeopleCodeParserT__70:
		localctx = NewExprCreateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(137)
			p.CreateInvocation()
		}

	case PeopleCodeParserT__14:
		localctx = NewExprNegateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(138)
			p.Match(PeopleCodeParserT__14)
		}
		{
			p.SetState(139)
			p.expr(8)
		}

	case PeopleCodeParserT__15:
		localctx = NewExprNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(140)
			p.Match(PeopleCodeParserT__15)
		}
		{
			p.SetState(141)
			p.expr(7)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PeopleCodeParserRULE_expr)
				p.SetState(144)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				p.SetState(147)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case PeopleCodeParserT__16:
					{
						p.SetState(145)

						var _m = p.Match(PeopleCodeParserT__16)

						localctx.(*ExprMulDivContext).m = _m
					}

				case PeopleCodeParserT__17:
					{
						p.SetState(146)

						var _m = p.Match(PeopleCodeParserT__17)

						localctx.(*ExprMulDivContext).d = _m
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}
				{
					p.SetState(149)
					p.expr(7)
				}

			case 2:
				localctx = NewExprAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PeopleCodeParserRULE_expr)
				p.SetState(150)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(153)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case PeopleCodeParserT__18:
					{
						p.SetState(151)

						var _m = p.Match(PeopleCodeParserT__18)

						localctx.(*ExprAddSubContext).a = _m
					}

				case PeopleCodeParserT__14:
					{
						p.SetState(152)

						var _m = p.Match(PeopleCodeParserT__14)

						localctx.(*ExprAddSubContext).s = _m
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}
				{
					p.SetState(155)
					p.expr(6)
				}

			case 3:
				localctx = NewExprComparisonContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PeopleCodeParserRULE_expr)
				p.SetState(156)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(161)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case PeopleCodeParserT__19:
					{
						p.SetState(157)

						var _m = p.Match(PeopleCodeParserT__19)

						localctx.(*ExprComparisonContext).le = _m
					}

				case PeopleCodeParserT__20:
					{
						p.SetState(158)

						var _m = p.Match(PeopleCodeParserT__20)

						localctx.(*ExprComparisonContext).ge = _m
					}

				case PeopleCodeParserT__21:
					{
						p.SetState(159)

						var _m = p.Match(PeopleCodeParserT__21)

						localctx.(*ExprComparisonContext).l = _m
					}

				case PeopleCodeParserT__22:
					{
						p.SetState(160)

						var _m = p.Match(PeopleCodeParserT__22)

						localctx.(*ExprComparisonContext).g = _m
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}
				{
					p.SetState(163)
					p.expr(5)
				}

			case 4:
				localctx = NewExprEqualityContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PeopleCodeParserRULE_expr)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(167)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case PeopleCodeParserT__7:
					{
						p.SetState(165)

						var _m = p.Match(PeopleCodeParserT__7)

						localctx.(*ExprEqualityContext).e = _m
					}

				case PeopleCodeParserT__23:
					{
						p.SetState(166)

						var _m = p.Match(PeopleCodeParserT__23)

						localctx.(*ExprEqualityContext).i = _m
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}
				{
					p.SetState(169)
					p.expr(4)
				}

			case 5:
				localctx = NewExprBooleanContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PeopleCodeParserRULE_expr)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(173)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case PeopleCodeParserT__24:
					{
						p.SetState(171)

						var _m = p.Match(PeopleCodeParserT__24)

						localctx.(*ExprBooleanContext).op = _m
					}

				case PeopleCodeParserT__25:
					{
						p.SetState(172)

						var _m = p.Match(PeopleCodeParserT__25)

						localctx.(*ExprBooleanContext).op = _m
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}
				{
					p.SetState(175)
					p.expr(3)
				}

			case 6:
				localctx = NewExprConcatContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PeopleCodeParserRULE_expr)
				p.SetState(176)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(177)
					p.Match(PeopleCodeParserT__26)
				}
				{
					p.SetState(178)
					p.expr(2)
				}

			case 7:
				localctx = NewExprDotAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PeopleCodeParserRULE_expr)
				p.SetState(179)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(180)
					p.Match(PeopleCodeParserT__11)
				}
				{
					p.SetState(181)
					p.Id_()
				}

			case 8:
				localctx = NewExprArrayIndexContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PeopleCodeParserRULE_expr)
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(183)
					p.Match(PeopleCodeParserT__12)
				}
				{
					p.SetState(184)
					p.ExprList()
				}
				{
					p.SetState(185)
					p.Match(PeopleCodeParserT__13)
				}

			case 9:
				localctx = NewExprFnOrIdxCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PeopleCodeParserRULE_expr)
				p.SetState(187)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(188)
					p.Match(PeopleCodeParserT__8)
				}
				p.SetState(190)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PeopleCodeParserT__8)|(1<<PeopleCodeParserT__10)|(1<<PeopleCodeParserT__14)|(1<<PeopleCodeParserT__15))) != 0) || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(PeopleCodeParserT__70-71))|(1<<(PeopleCodeParserDecimalLiteral-71))|(1<<(PeopleCodeParserIntegerLiteral-71))|(1<<(PeopleCodeParserStringLiteral-71))|(1<<(PeopleCodeParserBoolLiteral-71))|(1<<(PeopleCodeParserVAR_ID-71))|(1<<(PeopleCodeParserSYS_VAR_ID-71))|(1<<(PeopleCodeParserGENERIC_ID-71)))) != 0) {
					{
						p.SetState(189)
						p.ExprList()
					}

				}
				{
					p.SetState(192)
					p.Match(PeopleCodeParserT__9)
				}

			}

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_exprList
	return p
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (p *PeopleCodeParser) ExprList() (localctx IExprListContext) {
	this := p
	_ = this

	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PeopleCodeParserRULE_exprList)
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
		p.SetState(198)
		p.expr(0)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PeopleCodeParserT__27 {
		{
			p.SetState(199)
			p.Match(PeopleCodeParserT__27)
		}
		{
			p.SetState(200)
			p.expr(0)
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVarScope returns the varScope token.
	GetVarScope() antlr.Token

	// SetVarScope sets the varScope token.
	SetVarScope(antlr.Token)

	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	varScope antlr.Token
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_varDeclaration
	return p
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) GetVarScope() antlr.Token { return s.varScope }

func (s *VarDeclarationContext) SetVarScope(v antlr.Token) { s.varScope = v }

func (s *VarDeclarationContext) VarType() IVarTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *VarDeclarationContext) AllVarDeclarator() []IVarDeclaratorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarDeclaratorContext)(nil)).Elem())
	var tst = make([]IVarDeclaratorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarDeclaratorContext)
		}
	}

	return tst
}

func (s *VarDeclarationContext) VarDeclarator(i int) IVarDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclaratorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarDeclaratorContext)
}

func (s *VarDeclarationContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitVarDeclaration(s)
	}
}

func (p *PeopleCodeParser) VarDeclaration() (localctx IVarDeclarationContext) {
	this := p
	_ = this

	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PeopleCodeParserRULE_varDeclaration)
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
		p.SetState(206)

		var _m = p.Match(PeopleCodeParserGENERIC_ID)

		localctx.(*VarDeclarationContext).varScope = _m
	}
	{
		p.SetState(207)
		p.VarType()
	}
	{
		p.SetState(208)
		p.VarDeclarator()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PeopleCodeParserT__27 {
		{
			p.SetState(209)
			p.Match(PeopleCodeParserT__27)
		}
		{
			p.SetState(210)
			p.VarDeclarator()
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVarDeclaratorContext is an interface to support dynamic dispatch.
type IVarDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclaratorContext differentiates from other interfaces.
	IsVarDeclaratorContext()
}

type VarDeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclaratorContext() *VarDeclaratorContext {
	var p = new(VarDeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_varDeclarator
	return p
}

func (*VarDeclaratorContext) IsVarDeclaratorContext() {}

func NewVarDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclaratorContext {
	var p = new(VarDeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_varDeclarator

	return p
}

func (s *VarDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclaratorContext) VAR_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserVAR_ID, 0)
}

func (s *VarDeclaratorContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterVarDeclarator(s)
	}
}

func (s *VarDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitVarDeclarator(s)
	}
}

func (p *PeopleCodeParser) VarDeclarator() (localctx IVarDeclaratorContext) {
	this := p
	_ = this

	localctx = NewVarDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PeopleCodeParserRULE_varDeclarator)
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
		p.Match(PeopleCodeParserVAR_ID)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__7 {
		{
			p.SetState(217)
			p.Match(PeopleCodeParserT__7)
		}
		{
			p.SetState(218)
			p.expr(0)
		}

	}

	return localctx
}

// IVarTypeContext is an interface to support dynamic dispatch.
type IVarTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarTypeContext differentiates from other interfaces.
	IsVarTypeContext()
}

type VarTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarTypeContext() *VarTypeContext {
	var p = new(VarTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_varType
	return p
}

func (*VarTypeContext) IsVarTypeContext() {}

func NewVarTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarTypeContext {
	var p = new(VarTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_varType

	return p
}

func (s *VarTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VarTypeContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *VarTypeContext) VarType() IVarTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *VarTypeContext) AppClassPath() IAppClassPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAppClassPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAppClassPathContext)
}

func (s *VarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterVarType(s)
	}
}

func (s *VarTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitVarType(s)
	}
}

func (p *PeopleCodeParser) VarType() (localctx IVarTypeContext) {
	this := p
	_ = this

	localctx = NewVarTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PeopleCodeParserRULE_varType)
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

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Match(PeopleCodeParserGENERIC_ID)
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PeopleCodeParserT__28 {
			{
				p.SetState(222)
				p.Match(PeopleCodeParserT__28)
			}
			{
				p.SetState(223)
				p.VarType()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.AppClassPath()
		}

	}

	return localctx
}

// IAppClassImportContext is an interface to support dynamic dispatch.
type IAppClassImportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAppClassImportContext differentiates from other interfaces.
	IsAppClassImportContext()
}

type AppClassImportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppClassImportContext() *AppClassImportContext {
	var p = new(AppClassImportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_appClassImport
	return p
}

func (*AppClassImportContext) IsAppClassImportContext() {}

func NewAppClassImportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppClassImportContext {
	var p = new(AppClassImportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_appClassImport

	return p
}

func (s *AppClassImportContext) GetParser() antlr.Parser { return s.parser }

func (s *AppClassImportContext) AppPkgPath() IAppPkgPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAppPkgPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAppPkgPathContext)
}

func (s *AppClassImportContext) AppClassPath() IAppClassPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAppClassPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAppClassPathContext)
}

func (s *AppClassImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppClassImportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppClassImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterAppClassImport(s)
	}
}

func (s *AppClassImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitAppClassImport(s)
	}
}

func (p *PeopleCodeParser) AppClassImport() (localctx IAppClassImportContext) {
	this := p
	_ = this

	localctx = NewAppClassImportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PeopleCodeParserRULE_appClassImport)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(PeopleCodeParserT__29)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(230)
			p.AppPkgPath()
		}

	case 2:
		{
			p.SetState(231)
			p.AppClassPath()
		}

	}

	return localctx
}

// IAppPkgPathContext is an interface to support dynamic dispatch.
type IAppPkgPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAppPkgPathContext differentiates from other interfaces.
	IsAppPkgPathContext()
}

type AppPkgPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppPkgPathContext() *AppPkgPathContext {
	var p = new(AppPkgPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_appPkgPath
	return p
}

func (*AppPkgPathContext) IsAppPkgPathContext() {}

func NewAppPkgPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppPkgPathContext {
	var p = new(AppPkgPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_appPkgPath

	return p
}

func (s *AppPkgPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AppPkgPathContext) AllGENERIC_ID() []antlr.TerminalNode {
	return s.GetTokens(PeopleCodeParserGENERIC_ID)
}

func (s *AppPkgPathContext) GENERIC_ID(i int) antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, i)
}

func (s *AppPkgPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppPkgPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppPkgPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterAppPkgPath(s)
	}
}

func (s *AppPkgPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitAppPkgPath(s)
	}
}

func (p *PeopleCodeParser) AppPkgPath() (localctx IAppPkgPathContext) {
	this := p
	_ = this

	localctx = NewAppPkgPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PeopleCodeParserRULE_appPkgPath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(234)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(235)
				p.Match(PeopleCodeParserT__30)
			}
			{
				p.SetState(236)
				p.Match(PeopleCodeParserGENERIC_ID)
			}

		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	{
		p.SetState(242)
		p.Match(PeopleCodeParserT__30)
	}
	{
		p.SetState(243)
		p.Match(PeopleCodeParserT__16)
	}

	return localctx
}

// IAppClassPathContext is an interface to support dynamic dispatch.
type IAppClassPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAppClassPathContext differentiates from other interfaces.
	IsAppClassPathContext()
}

type AppClassPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppClassPathContext() *AppClassPathContext {
	var p = new(AppClassPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_appClassPath
	return p
}

func (*AppClassPathContext) IsAppClassPathContext() {}

func NewAppClassPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppClassPathContext {
	var p = new(AppClassPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_appClassPath

	return p
}

func (s *AppClassPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AppClassPathContext) AllGENERIC_ID() []antlr.TerminalNode {
	return s.GetTokens(PeopleCodeParserGENERIC_ID)
}

func (s *AppClassPathContext) GENERIC_ID(i int) antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, i)
}

func (s *AppClassPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppClassPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppClassPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterAppClassPath(s)
	}
}

func (s *AppClassPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitAppClassPath(s)
	}
}

func (p *PeopleCodeParser) AppClassPath() (localctx IAppClassPathContext) {
	this := p
	_ = this

	localctx = NewAppClassPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PeopleCodeParserRULE_appClassPath)
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
		p.SetState(245)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PeopleCodeParserT__30 {
		{
			p.SetState(246)
			p.Match(PeopleCodeParserT__30)
		}
		{
			p.SetState(247)
			p.Match(PeopleCodeParserGENERIC_ID)
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExtFuncImportContext is an interface to support dynamic dispatch.
type IExtFuncImportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtFuncImportContext differentiates from other interfaces.
	IsExtFuncImportContext()
}

type ExtFuncImportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtFuncImportContext() *ExtFuncImportContext {
	var p = new(ExtFuncImportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_extFuncImport
	return p
}

func (*ExtFuncImportContext) IsExtFuncImportContext() {}

func NewExtFuncImportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtFuncImportContext {
	var p = new(ExtFuncImportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_extFuncImport

	return p
}

func (s *ExtFuncImportContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtFuncImportContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *ExtFuncImportContext) RecDefnPath() IRecDefnPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecDefnPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecDefnPathContext)
}

func (s *ExtFuncImportContext) Event() IEventContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEventContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEventContext)
}

func (s *ExtFuncImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtFuncImportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtFuncImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterExtFuncImport(s)
	}
}

func (s *ExtFuncImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitExtFuncImport(s)
	}
}

func (p *PeopleCodeParser) ExtFuncImport() (localctx IExtFuncImportContext) {
	this := p
	_ = this

	localctx = NewExtFuncImportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PeopleCodeParserRULE_extFuncImport)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(PeopleCodeParserT__31)
	}
	{
		p.SetState(253)
		p.Match(PeopleCodeParserT__32)
	}
	{
		p.SetState(254)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	{
		p.SetState(255)
		p.Match(PeopleCodeParserT__33)
	}
	{
		p.SetState(256)
		p.RecDefnPath()
	}
	{
		p.SetState(257)
		p.Event()
	}

	return localctx
}

// IRecDefnPathContext is an interface to support dynamic dispatch.
type IRecDefnPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecDefnPathContext differentiates from other interfaces.
	IsRecDefnPathContext()
}

type RecDefnPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecDefnPathContext() *RecDefnPathContext {
	var p = new(RecDefnPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_recDefnPath
	return p
}

func (*RecDefnPathContext) IsRecDefnPathContext() {}

func NewRecDefnPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecDefnPathContext {
	var p = new(RecDefnPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_recDefnPath

	return p
}

func (s *RecDefnPathContext) GetParser() antlr.Parser { return s.parser }

func (s *RecDefnPathContext) AllGENERIC_ID() []antlr.TerminalNode {
	return s.GetTokens(PeopleCodeParserGENERIC_ID)
}

func (s *RecDefnPathContext) GENERIC_ID(i int) antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, i)
}

func (s *RecDefnPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecDefnPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecDefnPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterRecDefnPath(s)
	}
}

func (s *RecDefnPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitRecDefnPath(s)
	}
}

func (p *PeopleCodeParser) RecDefnPath() (localctx IRecDefnPathContext) {
	this := p
	_ = this

	localctx = NewRecDefnPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PeopleCodeParserRULE_recDefnPath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	{
		p.SetState(260)
		p.Match(PeopleCodeParserT__11)
	}
	{
		p.SetState(261)
		p.Match(PeopleCodeParserGENERIC_ID)
	}

	return localctx
}

// IEventContext is an interface to support dynamic dispatch.
type IEventContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEventContext differentiates from other interfaces.
	IsEventContext()
}

type EventContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventContext() *EventContext {
	var p = new(EventContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_event
	return p
}

func (*EventContext) IsEventContext() {}

func NewEventContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventContext {
	var p = new(EventContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_event

	return p
}

func (s *EventContext) GetParser() antlr.Parser { return s.parser }
func (s *EventContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterEvent(s)
	}
}

func (s *EventContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitEvent(s)
	}
}

func (p *PeopleCodeParser) Event() (localctx IEventContext) {
	this := p
	_ = this

	localctx = NewEventContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PeopleCodeParserRULE_event)
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
		p.SetState(263)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PeopleCodeParserT__34 || _la == PeopleCodeParserT__35) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_classDeclaration
	return p
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *ClassDeclarationContext) AllClassBlock() []IClassBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassBlockContext)(nil)).Elem())
	var tst = make([]IClassBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassBlockContext)
		}
	}

	return tst
}

func (s *ClassDeclarationContext) ClassBlock(i int) IClassBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassBlockContext)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (p *PeopleCodeParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	this := p
	_ = this

	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PeopleCodeParserRULE_classDeclaration)
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
		p.SetState(265)
		p.Match(PeopleCodeParserT__36)
	}
	{
		p.SetState(266)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(PeopleCodeParserT__38-39))|(1<<(PeopleCodeParserT__39-39))|(1<<(PeopleCodeParserT__40-39))|(1<<(PeopleCodeParserT__41-39))|(1<<(PeopleCodeParserT__45-39)))) != 0 {
		{
			p.SetState(267)
			p.ClassBlock()
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(273)
		p.Match(PeopleCodeParserT__37)
	}

	return localctx
}

// IClassBlockContext is an interface to support dynamic dispatch.
type IClassBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetALvl returns the aLvl token.
	GetALvl() antlr.Token

	// SetALvl sets the aLvl token.
	SetALvl(antlr.Token)

	// IsClassBlockContext differentiates from other interfaces.
	IsClassBlockContext()
}

type ClassBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	aLvl   antlr.Token
}

func NewEmptyClassBlockContext() *ClassBlockContext {
	var p = new(ClassBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_classBlock
	return p
}

func (*ClassBlockContext) IsClassBlockContext() {}

func NewClassBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBlockContext {
	var p = new(ClassBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_classBlock

	return p
}

func (s *ClassBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBlockContext) GetALvl() antlr.Token { return s.aLvl }

func (s *ClassBlockContext) SetALvl(v antlr.Token) { s.aLvl = v }

func (s *ClassBlockContext) AllClassBlockStmt() []IClassBlockStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassBlockStmtContext)(nil)).Elem())
	var tst = make([]IClassBlockStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassBlockStmtContext)
		}
	}

	return tst
}

func (s *ClassBlockContext) ClassBlockStmt(i int) IClassBlockStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBlockStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassBlockStmtContext)
}

func (s *ClassBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterClassBlock(s)
	}
}

func (s *ClassBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitClassBlock(s)
	}
}

func (p *PeopleCodeParser) ClassBlock() (localctx IClassBlockContext) {
	this := p
	_ = this

	localctx = NewClassBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PeopleCodeParserRULE_classBlock)
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
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__38 {
		{
			p.SetState(275)

			var _m = p.Match(PeopleCodeParserT__38)

			localctx.(*ClassBlockContext).aLvl = _m
		}

	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(278)
				p.ClassBlockStmt()
			}
			p.SetState(282)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == PeopleCodeParserT__0 {
				{
					p.SetState(279)
					p.Match(PeopleCodeParserT__0)
				}

				p.SetState(284)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IClassBlockStmtContext is an interface to support dynamic dispatch.
type IClassBlockStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBlockStmtContext differentiates from other interfaces.
	IsClassBlockStmtContext()
}

type ClassBlockStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBlockStmtContext() *ClassBlockStmtContext {
	var p = new(ClassBlockStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_classBlockStmt
	return p
}

func (*ClassBlockStmtContext) IsClassBlockStmtContext() {}

func NewClassBlockStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBlockStmtContext {
	var p = new(ClassBlockStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_classBlockStmt

	return p
}

func (s *ClassBlockStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBlockStmtContext) Method() IMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *ClassBlockStmtContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ClassBlockStmtContext) Property_() IProperty_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProperty_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProperty_Context)
}

func (s *ClassBlockStmtContext) Instance() IInstanceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstanceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstanceContext)
}

func (s *ClassBlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBlockStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterClassBlockStmt(s)
	}
}

func (s *ClassBlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitClassBlockStmt(s)
	}
}

func (p *PeopleCodeParser) ClassBlockStmt() (localctx IClassBlockStmtContext) {
	this := p
	_ = this

	localctx = NewClassBlockStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PeopleCodeParserRULE_classBlockStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(293)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PeopleCodeParserT__39:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(289)
			p.Method()
		}

	case PeopleCodeParserT__40:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(290)
			p.Constant()
		}

	case PeopleCodeParserT__41:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(291)
			p.Property_()
		}

	case PeopleCodeParserT__45:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(292)
			p.Instance()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *MethodContext) FormalParamList() IFormalParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParamListContext)
}

func (s *MethodContext) ReturnType() IReturnTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *PeopleCodeParser) Method() (localctx IMethodContext) {
	this := p
	_ = this

	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PeopleCodeParserRULE_method)
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
		p.SetState(295)
		p.Match(PeopleCodeParserT__39)
	}
	{
		p.SetState(296)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	{
		p.SetState(297)
		p.FormalParamList()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__51 {
		{
			p.SetState(298)
			p.ReturnType()
		}

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
	p.RuleIndex = PeopleCodeParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) VAR_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserVAR_ID, 0)
}

func (s *ConstantContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *PeopleCodeParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PeopleCodeParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(PeopleCodeParserT__40)
	}
	{
		p.SetState(302)
		p.Match(PeopleCodeParserVAR_ID)
	}
	{
		p.SetState(303)
		p.Match(PeopleCodeParserT__7)
	}
	{
		p.SetState(304)
		p.expr(0)
	}

	return localctx
}

// IProperty_Context is an interface to support dynamic dispatch.
type IProperty_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetG returns the g token.
	GetG() antlr.Token

	// GetS returns the s token.
	GetS() antlr.Token

	// GetR returns the r token.
	GetR() antlr.Token

	// SetG sets the g token.
	SetG(antlr.Token)

	// SetS sets the s token.
	SetS(antlr.Token)

	// SetR sets the r token.
	SetR(antlr.Token)

	// IsProperty_Context differentiates from other interfaces.
	IsProperty_Context()
}

type Property_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	g      antlr.Token
	s      antlr.Token
	r      antlr.Token
}

func NewEmptyProperty_Context() *Property_Context {
	var p = new(Property_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_property_
	return p
}

func (*Property_Context) IsProperty_Context() {}

func NewProperty_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_Context {
	var p = new(Property_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_property_

	return p
}

func (s *Property_Context) GetParser() antlr.Parser { return s.parser }

func (s *Property_Context) GetG() antlr.Token { return s.g }

func (s *Property_Context) GetS() antlr.Token { return s.s }

func (s *Property_Context) GetR() antlr.Token { return s.r }

func (s *Property_Context) SetG(v antlr.Token) { s.g = v }

func (s *Property_Context) SetS(v antlr.Token) { s.s = v }

func (s *Property_Context) SetR(v antlr.Token) { s.r = v }

func (s *Property_Context) VarType() IVarTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *Property_Context) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *Property_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterProperty_(s)
	}
}

func (s *Property_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitProperty_(s)
	}
}

func (p *PeopleCodeParser) Property_() (localctx IProperty_Context) {
	this := p
	_ = this

	localctx = NewProperty_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PeopleCodeParserRULE_property_)
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
		p.SetState(306)
		p.Match(PeopleCodeParserT__41)
	}
	{
		p.SetState(307)
		p.VarType()
	}
	{
		p.SetState(308)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__42 {
		{
			p.SetState(309)

			var _m = p.Match(PeopleCodeParserT__42)

			localctx.(*Property_Context).g = _m
		}

	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__43 {
		{
			p.SetState(312)

			var _m = p.Match(PeopleCodeParserT__43)

			localctx.(*Property_Context).s = _m
		}

	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__44 {
		{
			p.SetState(315)

			var _m = p.Match(PeopleCodeParserT__44)

			localctx.(*Property_Context).r = _m
		}

	}

	return localctx
}

// IInstanceContext is an interface to support dynamic dispatch.
type IInstanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstanceContext differentiates from other interfaces.
	IsInstanceContext()
}

type InstanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstanceContext() *InstanceContext {
	var p = new(InstanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_instance
	return p
}

func (*InstanceContext) IsInstanceContext() {}

func NewInstanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceContext {
	var p = new(InstanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_instance

	return p
}

func (s *InstanceContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceContext) VarType() IVarTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *InstanceContext) AllVAR_ID() []antlr.TerminalNode {
	return s.GetTokens(PeopleCodeParserVAR_ID)
}

func (s *InstanceContext) VAR_ID(i int) antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserVAR_ID, i)
}

func (s *InstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterInstance(s)
	}
}

func (s *InstanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitInstance(s)
	}
}

func (p *PeopleCodeParser) Instance() (localctx IInstanceContext) {
	this := p
	_ = this

	localctx = NewInstanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PeopleCodeParserRULE_instance)
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
		p.SetState(318)
		p.Match(PeopleCodeParserT__45)
	}
	{
		p.SetState(319)
		p.VarType()
	}
	{
		p.SetState(320)
		p.Match(PeopleCodeParserVAR_ID)
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PeopleCodeParserT__27 {
		{
			p.SetState(321)
			p.Match(PeopleCodeParserT__27)
		}
		{
			p.SetState(322)
			p.Match(PeopleCodeParserVAR_ID)
		}

		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodImplContext is an interface to support dynamic dispatch.
type IMethodImplContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEndmethod returns the endmethod token.
	GetEndmethod() antlr.Token

	// SetEndmethod sets the endmethod token.
	SetEndmethod(antlr.Token)

	// IsMethodImplContext differentiates from other interfaces.
	IsMethodImplContext()
}

type MethodImplContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	endmethod antlr.Token
}

func NewEmptyMethodImplContext() *MethodImplContext {
	var p = new(MethodImplContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_methodImpl
	return p
}

func (*MethodImplContext) IsMethodImplContext() {}

func NewMethodImplContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodImplContext {
	var p = new(MethodImplContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_methodImpl

	return p
}

func (s *MethodImplContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodImplContext) GetEndmethod() antlr.Token { return s.endmethod }

func (s *MethodImplContext) SetEndmethod(v antlr.Token) { s.endmethod = v }

func (s *MethodImplContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *MethodImplContext) StmtList() IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *MethodImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodImplContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterMethodImpl(s)
	}
}

func (s *MethodImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitMethodImpl(s)
	}
}

func (p *PeopleCodeParser) MethodImpl() (localctx IMethodImplContext) {
	this := p
	_ = this

	localctx = NewMethodImplContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PeopleCodeParserRULE_methodImpl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(PeopleCodeParserT__39)
	}
	{
		p.SetState(329)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	{
		p.SetState(330)
		p.StmtList()
	}
	{
		p.SetState(331)

		var _m = p.Match(PeopleCodeParserT__46)

		localctx.(*MethodImplContext).endmethod = _m
	}

	return localctx
}

// IGetImplContext is an interface to support dynamic dispatch.
type IGetImplContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEndget returns the endget token.
	GetEndget() antlr.Token

	// SetEndget sets the endget token.
	SetEndget(antlr.Token)

	// IsGetImplContext differentiates from other interfaces.
	IsGetImplContext()
}

type GetImplContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	endget antlr.Token
}

func NewEmptyGetImplContext() *GetImplContext {
	var p = new(GetImplContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_getImpl
	return p
}

func (*GetImplContext) IsGetImplContext() {}

func NewGetImplContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetImplContext {
	var p = new(GetImplContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_getImpl

	return p
}

func (s *GetImplContext) GetParser() antlr.Parser { return s.parser }

func (s *GetImplContext) GetEndget() antlr.Token { return s.endget }

func (s *GetImplContext) SetEndget(v antlr.Token) { s.endget = v }

func (s *GetImplContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *GetImplContext) StmtList() IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *GetImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetImplContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterGetImpl(s)
	}
}

func (s *GetImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitGetImpl(s)
	}
}

func (p *PeopleCodeParser) GetImpl() (localctx IGetImplContext) {
	this := p
	_ = this

	localctx = NewGetImplContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PeopleCodeParserRULE_getImpl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(PeopleCodeParserT__42)
	}
	{
		p.SetState(334)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	{
		p.SetState(335)
		p.StmtList()
	}
	{
		p.SetState(336)

		var _m = p.Match(PeopleCodeParserT__47)

		localctx.(*GetImplContext).endget = _m
	}

	return localctx
}

// ISetImplContext is an interface to support dynamic dispatch.
type ISetImplContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEndset returns the endset token.
	GetEndset() antlr.Token

	// SetEndset sets the endset token.
	SetEndset(antlr.Token)

	// IsSetImplContext differentiates from other interfaces.
	IsSetImplContext()
}

type SetImplContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	endset antlr.Token
}

func NewEmptySetImplContext() *SetImplContext {
	var p = new(SetImplContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_setImpl
	return p
}

func (*SetImplContext) IsSetImplContext() {}

func NewSetImplContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetImplContext {
	var p = new(SetImplContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_setImpl

	return p
}

func (s *SetImplContext) GetParser() antlr.Parser { return s.parser }

func (s *SetImplContext) GetEndset() antlr.Token { return s.endset }

func (s *SetImplContext) SetEndset(v antlr.Token) { s.endset = v }

func (s *SetImplContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *SetImplContext) StmtList() IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *SetImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetImplContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterSetImpl(s)
	}
}

func (s *SetImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitSetImpl(s)
	}
}

func (p *PeopleCodeParser) SetImpl() (localctx ISetImplContext) {
	this := p
	_ = this

	localctx = NewSetImplContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PeopleCodeParserRULE_setImpl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Match(PeopleCodeParserT__43)
	}
	{
		p.SetState(339)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	{
		p.SetState(340)
		p.StmtList()
	}
	{
		p.SetState(341)

		var _m = p.Match(PeopleCodeParserT__48)

		localctx.(*SetImplContext).endset = _m
	}

	return localctx
}

// IFuncImplContext is an interface to support dynamic dispatch.
type IFuncImplContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEndfunction returns the endfunction token.
	GetEndfunction() antlr.Token

	// SetEndfunction sets the endfunction token.
	SetEndfunction(antlr.Token)

	// IsFuncImplContext differentiates from other interfaces.
	IsFuncImplContext()
}

type FuncImplContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	endfunction antlr.Token
}

func NewEmptyFuncImplContext() *FuncImplContext {
	var p = new(FuncImplContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_funcImpl
	return p
}

func (*FuncImplContext) IsFuncImplContext() {}

func NewFuncImplContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncImplContext {
	var p = new(FuncImplContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_funcImpl

	return p
}

func (s *FuncImplContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncImplContext) GetEndfunction() antlr.Token { return s.endfunction }

func (s *FuncImplContext) SetEndfunction(v antlr.Token) { s.endfunction = v }

func (s *FuncImplContext) FuncSignature() IFuncSignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSignatureContext)
}

func (s *FuncImplContext) StmtList() IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *FuncImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncImplContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterFuncImpl(s)
	}
}

func (s *FuncImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitFuncImpl(s)
	}
}

func (p *PeopleCodeParser) FuncImpl() (localctx IFuncImplContext) {
	this := p
	_ = this

	localctx = NewFuncImplContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PeopleCodeParserRULE_funcImpl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.FuncSignature()
	}
	{
		p.SetState(344)
		p.StmtList()
	}
	{
		p.SetState(345)

		var _m = p.Match(PeopleCodeParserT__49)

		localctx.(*FuncImplContext).endfunction = _m
	}

	return localctx
}

// IFuncSignatureContext is an interface to support dynamic dispatch.
type IFuncSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSignatureContext differentiates from other interfaces.
	IsFuncSignatureContext()
}

type FuncSignatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSignatureContext() *FuncSignatureContext {
	var p = new(FuncSignatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_funcSignature
	return p
}

func (*FuncSignatureContext) IsFuncSignatureContext() {}

func NewFuncSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSignatureContext {
	var p = new(FuncSignatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_funcSignature

	return p
}

func (s *FuncSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSignatureContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *FuncSignatureContext) FormalParamList() IFormalParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParamListContext)
}

func (s *FuncSignatureContext) ReturnType() IReturnTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *FuncSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterFuncSignature(s)
	}
}

func (s *FuncSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitFuncSignature(s)
	}
}

func (p *PeopleCodeParser) FuncSignature() (localctx IFuncSignatureContext) {
	this := p
	_ = this

	localctx = NewFuncSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PeopleCodeParserRULE_funcSignature)
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
		p.SetState(347)
		p.Match(PeopleCodeParserT__32)
	}
	{
		p.SetState(348)
		p.Match(PeopleCodeParserGENERIC_ID)
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(349)
			p.FormalParamList()
		}

	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__51 {
		{
			p.SetState(352)
			p.ReturnType()
		}

	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__0 {
		{
			p.SetState(355)
			p.Match(PeopleCodeParserT__0)
		}

	}

	return localctx
}

// IFormalParamListContext is an interface to support dynamic dispatch.
type IFormalParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParamListContext differentiates from other interfaces.
	IsFormalParamListContext()
}

type FormalParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParamListContext() *FormalParamListContext {
	var p = new(FormalParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_formalParamList
	return p
}

func (*FormalParamListContext) IsFormalParamListContext() {}

func NewFormalParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParamListContext {
	var p = new(FormalParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_formalParamList

	return p
}

func (s *FormalParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParamListContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *FormalParamListContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *FormalParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterFormalParamList(s)
	}
}

func (s *FormalParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitFormalParamList(s)
	}
}

func (p *PeopleCodeParser) FormalParamList() (localctx IFormalParamListContext) {
	this := p
	_ = this

	localctx = NewFormalParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PeopleCodeParserRULE_formalParamList)
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
		p.Match(PeopleCodeParserT__8)
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserVAR_ID {
		{
			p.SetState(359)
			p.Param()
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PeopleCodeParserT__27 {
			{
				p.SetState(360)
				p.Match(PeopleCodeParserT__27)
			}
			{
				p.SetState(361)
				p.Param()
			}

			p.SetState(366)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(369)
		p.Match(PeopleCodeParserT__9)
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
	p.RuleIndex = PeopleCodeParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) VAR_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserVAR_ID, 0)
}

func (s *ParamContext) VarType() IVarTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *PeopleCodeParser) Param() (localctx IParamContext) {
	this := p
	_ = this

	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PeopleCodeParserRULE_param)
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
		p.SetState(371)
		p.Match(PeopleCodeParserVAR_ID)
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__50 {
		{
			p.SetState(372)
			p.Match(PeopleCodeParserT__50)
		}
		{
			p.SetState(373)
			p.VarType()
		}

	}

	return localctx
}

// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_returnType
	return p
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) VarType() IVarTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterReturnType(s)
	}
}

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitReturnType(s)
	}
}

func (p *PeopleCodeParser) ReturnType() (localctx IReturnTypeContext) {
	this := p
	_ = this

	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PeopleCodeParserRULE_returnType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(PeopleCodeParserT__51)
	}
	{
		p.SetState(377)
		p.VarType()
	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetElsetok returns the elsetok token.
	GetElsetok() antlr.Token

	// GetEndif returns the endif token.
	GetEndif() antlr.Token

	// SetElsetok sets the elsetok token.
	SetElsetok(antlr.Token)

	// SetEndif sets the endif token.
	SetEndif(antlr.Token)

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	elsetok antlr.Token
	endif   antlr.Token
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) GetElsetok() antlr.Token { return s.elsetok }

func (s *IfStmtContext) GetEndif() antlr.Token { return s.endif }

func (s *IfStmtContext) SetElsetok(v antlr.Token) { s.elsetok = v }

func (s *IfStmtContext) SetEndif(v antlr.Token) { s.endif = v }

func (s *IfStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) AllStmtList() []IStmtListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtListContext)(nil)).Elem())
	var tst = make([]IStmtListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtListContext)
		}
	}

	return tst
}

func (s *IfStmtContext) StmtList(i int) IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (p *PeopleCodeParser) IfStmt() (localctx IIfStmtContext) {
	this := p
	_ = this

	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PeopleCodeParserRULE_ifStmt)
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
		p.SetState(379)
		p.Match(PeopleCodeParserT__52)
	}
	{
		p.SetState(380)
		p.expr(0)
	}
	{
		p.SetState(381)
		p.Match(PeopleCodeParserT__53)
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__0 {
		{
			p.SetState(382)
			p.Match(PeopleCodeParserT__0)
		}

	}
	{
		p.SetState(385)
		p.StmtList()
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__54 {
		{
			p.SetState(386)

			var _m = p.Match(PeopleCodeParserT__54)

			localctx.(*IfStmtContext).elsetok = _m
		}
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PeopleCodeParserT__0 {
			{
				p.SetState(387)
				p.Match(PeopleCodeParserT__0)
			}

		}
		{
			p.SetState(390)
			p.StmtList()
		}

	}
	{
		p.SetState(393)

		var _m = p.Match(PeopleCodeParserT__55)

		localctx.(*IfStmtContext).endif = _m
	}

	return localctx
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEndfor returns the endfor token.
	GetEndfor() antlr.Token

	// SetEndfor sets the endfor token.
	SetEndfor(antlr.Token)

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	endfor antlr.Token
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_forStmt
	return p
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) GetEndfor() antlr.Token { return s.endfor }

func (s *ForStmtContext) SetEndfor(v antlr.Token) { s.endfor = v }

func (s *ForStmtContext) VAR_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserVAR_ID, 0)
}

func (s *ForStmtContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ForStmtContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) StmtList() IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (p *PeopleCodeParser) ForStmt() (localctx IForStmtContext) {
	this := p
	_ = this

	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PeopleCodeParserRULE_forStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(PeopleCodeParserT__56)
	}
	{
		p.SetState(396)
		p.Match(PeopleCodeParserVAR_ID)
	}
	{
		p.SetState(397)
		p.Match(PeopleCodeParserT__7)
	}
	{
		p.SetState(398)
		p.expr(0)
	}
	{
		p.SetState(399)
		p.Match(PeopleCodeParserT__57)
	}
	{
		p.SetState(400)
		p.expr(0)
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PeopleCodeParserT__0:
		{
			p.SetState(401)
			p.Match(PeopleCodeParserT__0)
		}

	case PeopleCodeParserT__58:
		{
			p.SetState(402)
			p.Match(PeopleCodeParserT__58)
		}
		{
			p.SetState(403)
			p.expr(0)
		}

	case PeopleCodeParserT__1, PeopleCodeParserT__2, PeopleCodeParserT__3, PeopleCodeParserT__4, PeopleCodeParserT__5, PeopleCodeParserT__6, PeopleCodeParserT__8, PeopleCodeParserT__10, PeopleCodeParserT__14, PeopleCodeParserT__15, PeopleCodeParserT__29, PeopleCodeParserT__31, PeopleCodeParserT__32, PeopleCodeParserT__36, PeopleCodeParserT__39, PeopleCodeParserT__42, PeopleCodeParserT__43, PeopleCodeParserT__52, PeopleCodeParserT__56, PeopleCodeParserT__59, PeopleCodeParserT__60, PeopleCodeParserT__62, PeopleCodeParserT__66, PeopleCodeParserT__70, PeopleCodeParserDecimalLiteral, PeopleCodeParserIntegerLiteral, PeopleCodeParserStringLiteral, PeopleCodeParserBoolLiteral, PeopleCodeParserVAR_ID, PeopleCodeParserSYS_VAR_ID, PeopleCodeParserGENERIC_ID:

	default:
	}
	{
		p.SetState(406)
		p.StmtList()
	}
	{
		p.SetState(407)

		var _m = p.Match(PeopleCodeParserT__59)

		localctx.(*ForStmtContext).endfor = _m
	}

	return localctx
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_whileStmt
	return p
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) StmtList() IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (p *PeopleCodeParser) WhileStmt() (localctx IWhileStmtContext) {
	this := p
	_ = this

	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PeopleCodeParserRULE_whileStmt)
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
		p.SetState(409)
		p.Match(PeopleCodeParserT__60)
	}
	{
		p.SetState(410)
		p.expr(0)
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__0 {
		{
			p.SetState(411)
			p.Match(PeopleCodeParserT__0)
		}

	}
	{
		p.SetState(414)
		p.StmtList()
	}
	{
		p.SetState(415)
		p.Match(PeopleCodeParserT__61)
	}

	return localctx
}

// IEvaluateStmtContext is an interface to support dynamic dispatch.
type IEvaluateStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEndevaluate returns the endevaluate token.
	GetEndevaluate() antlr.Token

	// SetEndevaluate sets the endevaluate token.
	SetEndevaluate(antlr.Token)

	// IsEvaluateStmtContext differentiates from other interfaces.
	IsEvaluateStmtContext()
}

type EvaluateStmtContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	endevaluate antlr.Token
}

func NewEmptyEvaluateStmtContext() *EvaluateStmtContext {
	var p = new(EvaluateStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_evaluateStmt
	return p
}

func (*EvaluateStmtContext) IsEvaluateStmtContext() {}

func NewEvaluateStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvaluateStmtContext {
	var p = new(EvaluateStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_evaluateStmt

	return p
}

func (s *EvaluateStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *EvaluateStmtContext) GetEndevaluate() antlr.Token { return s.endevaluate }

func (s *EvaluateStmtContext) SetEndevaluate(v antlr.Token) { s.endevaluate = v }

func (s *EvaluateStmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EvaluateStmtContext) AllWhenBranch() []IWhenBranchContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhenBranchContext)(nil)).Elem())
	var tst = make([]IWhenBranchContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhenBranchContext)
		}
	}

	return tst
}

func (s *EvaluateStmtContext) WhenBranch(i int) IWhenBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhenBranchContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhenBranchContext)
}

func (s *EvaluateStmtContext) WhenOtherBranch() IWhenOtherBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhenOtherBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhenOtherBranchContext)
}

func (s *EvaluateStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvaluateStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvaluateStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterEvaluateStmt(s)
	}
}

func (s *EvaluateStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitEvaluateStmt(s)
	}
}

func (p *PeopleCodeParser) EvaluateStmt() (localctx IEvaluateStmtContext) {
	this := p
	_ = this

	localctx = NewEvaluateStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PeopleCodeParserRULE_evaluateStmt)
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
		p.SetState(417)
		p.Match(PeopleCodeParserT__62)
	}
	{
		p.SetState(418)
		p.expr(0)
	}
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PeopleCodeParserT__64 {
		{
			p.SetState(419)
			p.WhenBranch()
		}

		p.SetState(422)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PeopleCodeParserT__65 {
		{
			p.SetState(424)
			p.WhenOtherBranch()
		}

	}
	{
		p.SetState(427)

		var _m = p.Match(PeopleCodeParserT__63)

		localctx.(*EvaluateStmtContext).endevaluate = _m
	}

	return localctx
}

// IWhenBranchContext is an interface to support dynamic dispatch.
type IWhenBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsWhenBranchContext differentiates from other interfaces.
	IsWhenBranchContext()
}

type WhenBranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyWhenBranchContext() *WhenBranchContext {
	var p = new(WhenBranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_whenBranch
	return p
}

func (*WhenBranchContext) IsWhenBranchContext() {}

func NewWhenBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenBranchContext {
	var p = new(WhenBranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_whenBranch

	return p
}

func (s *WhenBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenBranchContext) GetOp() antlr.Token { return s.op }

func (s *WhenBranchContext) SetOp(v antlr.Token) { s.op = v }

func (s *WhenBranchContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhenBranchContext) StmtList() IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *WhenBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterWhenBranch(s)
	}
}

func (s *WhenBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitWhenBranch(s)
	}
}

func (p *PeopleCodeParser) WhenBranch() (localctx IWhenBranchContext) {
	this := p
	_ = this

	localctx = NewWhenBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PeopleCodeParserRULE_whenBranch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(PeopleCodeParserT__64)
	}
	p.SetState(432)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PeopleCodeParserT__7:
		{
			p.SetState(430)

			var _m = p.Match(PeopleCodeParserT__7)

			localctx.(*WhenBranchContext).op = _m
		}

	case PeopleCodeParserT__22:
		{
			p.SetState(431)

			var _m = p.Match(PeopleCodeParserT__22)

			localctx.(*WhenBranchContext).op = _m
		}

	case PeopleCodeParserT__8, PeopleCodeParserT__10, PeopleCodeParserT__14, PeopleCodeParserT__15, PeopleCodeParserT__70, PeopleCodeParserDecimalLiteral, PeopleCodeParserIntegerLiteral, PeopleCodeParserStringLiteral, PeopleCodeParserBoolLiteral, PeopleCodeParserVAR_ID, PeopleCodeParserSYS_VAR_ID, PeopleCodeParserGENERIC_ID:

	default:
	}
	{
		p.SetState(434)
		p.expr(0)
	}
	{
		p.SetState(435)
		p.StmtList()
	}

	return localctx
}

// IWhenOtherBranchContext is an interface to support dynamic dispatch.
type IWhenOtherBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhenOtherBranchContext differentiates from other interfaces.
	IsWhenOtherBranchContext()
}

type WhenOtherBranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhenOtherBranchContext() *WhenOtherBranchContext {
	var p = new(WhenOtherBranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_whenOtherBranch
	return p
}

func (*WhenOtherBranchContext) IsWhenOtherBranchContext() {}

func NewWhenOtherBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenOtherBranchContext {
	var p = new(WhenOtherBranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_whenOtherBranch

	return p
}

func (s *WhenOtherBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenOtherBranchContext) StmtList() IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *WhenOtherBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenOtherBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenOtherBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterWhenOtherBranch(s)
	}
}

func (s *WhenOtherBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitWhenOtherBranch(s)
	}
}

func (p *PeopleCodeParser) WhenOtherBranch() (localctx IWhenOtherBranchContext) {
	this := p
	_ = this

	localctx = NewWhenOtherBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PeopleCodeParserRULE_whenOtherBranch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(PeopleCodeParserT__65)
	}
	{
		p.SetState(438)
		p.StmtList()
	}

	return localctx
}

// ITryCatchStmtContext is an interface to support dynamic dispatch.
type ITryCatchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTrytok returns the trytok token.
	GetTrytok() antlr.Token

	// GetEndtry returns the endtry token.
	GetEndtry() antlr.Token

	// SetTrytok sets the trytok token.
	SetTrytok(antlr.Token)

	// SetEndtry sets the endtry token.
	SetEndtry(antlr.Token)

	// IsTryCatchStmtContext differentiates from other interfaces.
	IsTryCatchStmtContext()
}

type TryCatchStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	trytok antlr.Token
	endtry antlr.Token
}

func NewEmptyTryCatchStmtContext() *TryCatchStmtContext {
	var p = new(TryCatchStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_tryCatchStmt
	return p
}

func (*TryCatchStmtContext) IsTryCatchStmtContext() {}

func NewTryCatchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryCatchStmtContext {
	var p = new(TryCatchStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_tryCatchStmt

	return p
}

func (s *TryCatchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TryCatchStmtContext) GetTrytok() antlr.Token { return s.trytok }

func (s *TryCatchStmtContext) GetEndtry() antlr.Token { return s.endtry }

func (s *TryCatchStmtContext) SetTrytok(v antlr.Token) { s.trytok = v }

func (s *TryCatchStmtContext) SetEndtry(v antlr.Token) { s.endtry = v }

func (s *TryCatchStmtContext) AllStmtList() []IStmtListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtListContext)(nil)).Elem())
	var tst = make([]IStmtListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtListContext)
		}
	}

	return tst
}

func (s *TryCatchStmtContext) StmtList(i int) IStmtListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *TryCatchStmtContext) CatchSignature() ICatchSignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICatchSignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICatchSignatureContext)
}

func (s *TryCatchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryCatchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryCatchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterTryCatchStmt(s)
	}
}

func (s *TryCatchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitTryCatchStmt(s)
	}
}

func (p *PeopleCodeParser) TryCatchStmt() (localctx ITryCatchStmtContext) {
	this := p
	_ = this

	localctx = NewTryCatchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PeopleCodeParserRULE_tryCatchStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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

		var _m = p.Match(PeopleCodeParserT__66)

		localctx.(*TryCatchStmtContext).trytok = _m
	}
	{
		p.SetState(441)
		p.StmtList()
	}
	{
		p.SetState(442)
		p.CatchSignature()
	}
	{
		p.SetState(443)
		p.StmtList()
	}
	{
		p.SetState(444)

		var _m = p.Match(PeopleCodeParserT__67)

		localctx.(*TryCatchStmtContext).endtry = _m
	}

	return localctx
}

// ICatchSignatureContext is an interface to support dynamic dispatch.
type ICatchSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExClass returns the exClass token.
	GetExClass() antlr.Token

	// SetExClass sets the exClass token.
	SetExClass(antlr.Token)

	// IsCatchSignatureContext differentiates from other interfaces.
	IsCatchSignatureContext()
}

type CatchSignatureContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	exClass antlr.Token
}

func NewEmptyCatchSignatureContext() *CatchSignatureContext {
	var p = new(CatchSignatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_catchSignature
	return p
}

func (*CatchSignatureContext) IsCatchSignatureContext() {}

func NewCatchSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatchSignatureContext {
	var p = new(CatchSignatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_catchSignature

	return p
}

func (s *CatchSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *CatchSignatureContext) GetExClass() antlr.Token { return s.exClass }

func (s *CatchSignatureContext) SetExClass(v antlr.Token) { s.exClass = v }

func (s *CatchSignatureContext) VAR_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserVAR_ID, 0)
}

func (s *CatchSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatchSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatchSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterCatchSignature(s)
	}
}

func (s *CatchSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitCatchSignature(s)
	}
}

func (p *PeopleCodeParser) CatchSignature() (localctx ICatchSignatureContext) {
	this := p
	_ = this

	localctx = NewCatchSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PeopleCodeParserRULE_catchSignature)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(PeopleCodeParserT__68)
	}
	{
		p.SetState(447)

		var _m = p.Match(PeopleCodeParserT__69)

		localctx.(*CatchSignatureContext).exClass = _m
	}
	{
		p.SetState(448)
		p.Match(PeopleCodeParserVAR_ID)
	}

	return localctx
}

// ICreateInvocationContext is an interface to support dynamic dispatch.
type ICreateInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateInvocationContext differentiates from other interfaces.
	IsCreateInvocationContext()
}

type CreateInvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateInvocationContext() *CreateInvocationContext {
	var p = new(CreateInvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_createInvocation
	return p
}

func (*CreateInvocationContext) IsCreateInvocationContext() {}

func NewCreateInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateInvocationContext {
	var p = new(CreateInvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_createInvocation

	return p
}

func (s *CreateInvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateInvocationContext) AppClassPath() IAppClassPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAppClassPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAppClassPathContext)
}

func (s *CreateInvocationContext) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *CreateInvocationContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *CreateInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateInvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateInvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterCreateInvocation(s)
	}
}

func (s *CreateInvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitCreateInvocation(s)
	}
}

func (p *PeopleCodeParser) CreateInvocation() (localctx ICreateInvocationContext) {
	this := p
	_ = this

	localctx = NewCreateInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PeopleCodeParserRULE_createInvocation)
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
		p.SetState(450)
		p.Match(PeopleCodeParserT__70)
	}
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(451)
			p.AppClassPath()
		}

	case 2:
		{
			p.SetState(452)
			p.Match(PeopleCodeParserGENERIC_ID)
		}

	}
	{
		p.SetState(455)
		p.Match(PeopleCodeParserT__8)
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PeopleCodeParserT__8)|(1<<PeopleCodeParserT__10)|(1<<PeopleCodeParserT__14)|(1<<PeopleCodeParserT__15))) != 0) || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(PeopleCodeParserT__70-71))|(1<<(PeopleCodeParserDecimalLiteral-71))|(1<<(PeopleCodeParserIntegerLiteral-71))|(1<<(PeopleCodeParserStringLiteral-71))|(1<<(PeopleCodeParserBoolLiteral-71))|(1<<(PeopleCodeParserVAR_ID-71))|(1<<(PeopleCodeParserSYS_VAR_ID-71))|(1<<(PeopleCodeParserGENERIC_ID-71)))) != 0) {
		{
			p.SetState(456)
			p.ExprList()
		}

	}
	{
		p.SetState(459)
		p.Match(PeopleCodeParserT__9)
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
	p.RuleIndex = PeopleCodeParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserDecimalLiteral, 0)
}

func (s *LiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserIntegerLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserStringLiteral, 0)
}

func (s *LiteralContext) BoolLiteral() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserBoolLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *PeopleCodeParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PeopleCodeParserRULE_literal)
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

		if !(((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(PeopleCodeParserDecimalLiteral-72))|(1<<(PeopleCodeParserIntegerLiteral-72))|(1<<(PeopleCodeParserStringLiteral-72))|(1<<(PeopleCodeParserBoolLiteral-72)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IId_Context is an interface to support dynamic dispatch.
type IId_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsId_Context differentiates from other interfaces.
	IsId_Context()
}

type Id_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_Context() *Id_Context {
	var p = new(Id_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PeopleCodeParserRULE_id_
	return p
}

func (*Id_Context) IsId_Context() {}

func NewId_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_Context {
	var p = new(Id_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PeopleCodeParserRULE_id_

	return p
}

func (s *Id_Context) GetParser() antlr.Parser { return s.parser }

func (s *Id_Context) SYS_VAR_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserSYS_VAR_ID, 0)
}

func (s *Id_Context) VAR_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserVAR_ID, 0)
}

func (s *Id_Context) GENERIC_ID() antlr.TerminalNode {
	return s.GetToken(PeopleCodeParserGENERIC_ID, 0)
}

func (s *Id_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.EnterId_(s)
	}
}

func (s *Id_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PeopleCodeListener); ok {
		listenerT.ExitId_(s)
	}
}

func (p *PeopleCodeParser) Id_() (localctx IId_Context) {
	this := p
	_ = this

	localctx = NewId_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PeopleCodeParserRULE_id_)
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
		p.SetState(463)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(PeopleCodeParserVAR_ID-76))|(1<<(PeopleCodeParserSYS_VAR_ID-76))|(1<<(PeopleCodeParserGENERIC_ID-76)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *PeopleCodeParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PeopleCodeParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
