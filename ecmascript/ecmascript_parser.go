// Code generated from ECMAScript.g4 by ANTLR 4.9.3. DO NOT EDIT.

package ecmascript // ECMAScript
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 105, 623,
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
	9, 55, 4, 56, 9, 56, 3, 2, 5, 2, 114, 10, 2, 3, 2, 3, 2, 3, 3, 6, 3, 119,
	10, 3, 13, 3, 14, 3, 120, 3, 4, 3, 4, 3, 4, 5, 4, 126, 10, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 5, 5, 144, 10, 5, 3, 6, 3, 6, 5, 6, 148, 10, 6, 3, 6, 3,
	6, 3, 7, 6, 7, 153, 10, 7, 13, 7, 14, 7, 154, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 7, 9, 164, 10, 9, 12, 9, 14, 9, 167, 11, 9, 3, 10, 3, 10,
	5, 10, 171, 10, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 188, 10, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 207, 10, 15, 3, 15,
	3, 15, 5, 15, 211, 10, 15, 3, 15, 3, 15, 5, 15, 215, 10, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 225, 10, 15, 3, 15,
	3, 15, 5, 15, 229, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 251, 10, 15, 3, 16, 3, 16, 5, 16, 255, 10,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 261, 10, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 5, 18, 267, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21,
	285, 10, 21, 3, 21, 3, 21, 5, 21, 289, 10, 21, 5, 21, 291, 10, 21, 3, 21,
	3, 21, 3, 22, 6, 22, 296, 10, 22, 13, 22, 14, 22, 297, 3, 23, 3, 23, 3,
	23, 3, 23, 5, 23, 304, 10, 23, 3, 24, 3, 24, 3, 24, 5, 24, 309, 10, 24,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	5, 27, 332, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 350,
	10, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 7, 32,
	360, 10, 32, 12, 32, 14, 32, 363, 11, 32, 3, 33, 5, 33, 366, 10, 33, 3,
	34, 3, 34, 5, 34, 370, 10, 34, 3, 34, 5, 34, 373, 10, 34, 3, 34, 5, 34,
	376, 10, 34, 3, 34, 3, 34, 3, 35, 5, 35, 381, 10, 35, 3, 35, 3, 35, 3,
	35, 5, 35, 386, 10, 35, 3, 35, 7, 35, 389, 10, 35, 12, 35, 14, 35, 392,
	11, 35, 3, 36, 6, 36, 395, 10, 36, 13, 36, 14, 36, 396, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 5, 37, 404, 10, 37, 3, 37, 3, 37, 5, 37, 408, 10, 37,
	3, 38, 3, 38, 3, 38, 7, 38, 413, 10, 38, 12, 38, 14, 38, 416, 11, 38, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 437, 10,
	39, 3, 40, 3, 40, 3, 40, 5, 40, 442, 10, 40, 3, 41, 3, 41, 3, 42, 3, 42,
	5, 42, 448, 10, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 7, 43, 455, 10,
	43, 12, 43, 14, 43, 458, 11, 43, 3, 44, 3, 44, 3, 44, 7, 44, 463, 10, 44,
	12, 44, 14, 44, 466, 11, 44, 3, 45, 3, 45, 3, 45, 5, 45, 471, 10, 45, 3,
	45, 3, 45, 5, 45, 475, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 5, 45, 485, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 5, 45, 514, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 7, 45, 581, 10, 45, 12, 45, 14,
	45, 584, 11, 45, 3, 46, 3, 46, 3, 47, 3, 47, 5, 47, 590, 10, 47, 3, 48,
	3, 48, 3, 49, 3, 49, 5, 49, 596, 10, 49, 3, 50, 3, 50, 3, 50, 5, 50, 601,
	10, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54,
	3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 619, 10, 55, 3,
	56, 3, 56, 3, 56, 2, 3, 88, 57, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
	96, 98, 100, 102, 104, 106, 108, 110, 2, 13, 3, 2, 23, 25, 3, 2, 19, 20,
	3, 2, 26, 28, 3, 2, 29, 32, 3, 2, 33, 36, 3, 2, 42, 52, 5, 2, 3, 3, 53,
	54, 101, 101, 3, 2, 55, 57, 3, 2, 53, 54, 3, 2, 58, 83, 3, 2, 84, 99, 2,
	676, 2, 113, 3, 2, 2, 2, 4, 118, 3, 2, 2, 2, 6, 125, 3, 2, 2, 2, 8, 143,
	3, 2, 2, 2, 10, 145, 3, 2, 2, 2, 12, 152, 3, 2, 2, 2, 14, 156, 3, 2, 2,
	2, 16, 160, 3, 2, 2, 2, 18, 168, 3, 2, 2, 2, 20, 172, 3, 2, 2, 2, 22, 175,
	3, 2, 2, 2, 24, 177, 3, 2, 2, 2, 26, 180, 3, 2, 2, 2, 28, 250, 3, 2, 2,
	2, 30, 252, 3, 2, 2, 2, 32, 258, 3, 2, 2, 2, 34, 264, 3, 2, 2, 2, 36, 270,
	3, 2, 2, 2, 38, 276, 3, 2, 2, 2, 40, 282, 3, 2, 2, 2, 42, 295, 3, 2, 2,
	2, 44, 299, 3, 2, 2, 2, 46, 305, 3, 2, 2, 2, 48, 310, 3, 2, 2, 2, 50, 314,
	3, 2, 2, 2, 52, 331, 3, 2, 2, 2, 54, 333, 3, 2, 2, 2, 56, 339, 3, 2, 2,
	2, 58, 342, 3, 2, 2, 2, 60, 345, 3, 2, 2, 2, 62, 356, 3, 2, 2, 2, 64, 365,
	3, 2, 2, 2, 66, 367, 3, 2, 2, 2, 68, 380, 3, 2, 2, 2, 70, 394, 3, 2, 2,
	2, 72, 407, 3, 2, 2, 2, 74, 409, 3, 2, 2, 2, 76, 436, 3, 2, 2, 2, 78, 441,
	3, 2, 2, 2, 80, 443, 3, 2, 2, 2, 82, 445, 3, 2, 2, 2, 84, 451, 3, 2, 2,
	2, 86, 459, 3, 2, 2, 2, 88, 513, 3, 2, 2, 2, 90, 585, 3, 2, 2, 2, 92, 589,
	3, 2, 2, 2, 94, 591, 3, 2, 2, 2, 96, 595, 3, 2, 2, 2, 98, 600, 3, 2, 2,
	2, 100, 602, 3, 2, 2, 2, 102, 604, 3, 2, 2, 2, 104, 606, 3, 2, 2, 2, 106,
	610, 3, 2, 2, 2, 108, 618, 3, 2, 2, 2, 110, 620, 3, 2, 2, 2, 112, 114,
	5, 4, 3, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 3, 2,
	2, 2, 115, 116, 7, 2, 2, 3, 116, 3, 3, 2, 2, 2, 117, 119, 5, 6, 4, 2, 118,
	117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121,
	3, 2, 2, 2, 121, 5, 3, 2, 2, 2, 122, 123, 6, 4, 2, 2, 123, 126, 5, 8, 5,
	2, 124, 126, 5, 60, 31, 2, 125, 122, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2,
	126, 7, 3, 2, 2, 2, 127, 144, 5, 10, 6, 2, 128, 144, 5, 14, 8, 2, 129,
	144, 5, 22, 12, 2, 130, 131, 6, 5, 3, 2, 131, 144, 5, 24, 13, 2, 132, 144,
	5, 26, 14, 2, 133, 144, 5, 28, 15, 2, 134, 144, 5, 30, 16, 2, 135, 144,
	5, 32, 17, 2, 136, 144, 5, 34, 18, 2, 137, 144, 5, 36, 19, 2, 138, 144,
	5, 48, 25, 2, 139, 144, 5, 38, 20, 2, 140, 144, 5, 50, 26, 2, 141, 144,
	5, 52, 27, 2, 142, 144, 5, 58, 30, 2, 143, 127, 3, 2, 2, 2, 143, 128, 3,
	2, 2, 2, 143, 129, 3, 2, 2, 2, 143, 130, 3, 2, 2, 2, 143, 132, 3, 2, 2,
	2, 143, 133, 3, 2, 2, 2, 143, 134, 3, 2, 2, 2, 143, 135, 3, 2, 2, 2, 143,
	136, 3, 2, 2, 2, 143, 137, 3, 2, 2, 2, 143, 138, 3, 2, 2, 2, 143, 139,
	3, 2, 2, 2, 143, 140, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 142, 3, 2,
	2, 2, 144, 9, 3, 2, 2, 2, 145, 147, 7, 9, 2, 2, 146, 148, 5, 12, 7, 2,
	147, 146, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	150, 7, 10, 2, 2, 150, 11, 3, 2, 2, 2, 151, 153, 5, 8, 5, 2, 152, 151,
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154, 155, 3, 2,
	2, 2, 155, 13, 3, 2, 2, 2, 156, 157, 7, 65, 2, 2, 157, 158, 5, 16, 9, 2,
	158, 159, 5, 108, 55, 2, 159, 15, 3, 2, 2, 2, 160, 165, 5, 18, 10, 2, 161,
	162, 7, 12, 2, 2, 162, 164, 5, 18, 10, 2, 163, 161, 3, 2, 2, 2, 164, 167,
	3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 17, 3, 2,
	2, 2, 167, 165, 3, 2, 2, 2, 168, 170, 7, 100, 2, 2, 169, 171, 5, 20, 11,
	2, 170, 169, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 19, 3, 2, 2, 2, 172,
	173, 7, 13, 2, 2, 173, 174, 5, 88, 45, 2, 174, 21, 3, 2, 2, 2, 175, 176,
	7, 11, 2, 2, 176, 23, 3, 2, 2, 2, 177, 178, 5, 86, 44, 2, 178, 179, 5,
	108, 55, 2, 179, 25, 3, 2, 2, 2, 180, 181, 7, 79, 2, 2, 181, 182, 7, 7,
	2, 2, 182, 183, 5, 86, 44, 2, 183, 184, 7, 8, 2, 2, 184, 187, 5, 8, 5,
	2, 185, 186, 7, 63, 2, 2, 186, 188, 5, 8, 5, 2, 187, 185, 3, 2, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 27, 3, 2, 2, 2, 189, 190, 7, 59, 2, 2, 190, 191,
	5, 8, 5, 2, 191, 192, 7, 73, 2, 2, 192, 193, 7, 7, 2, 2, 193, 194, 5, 86,
	44, 2, 194, 195, 7, 8, 2, 2, 195, 196, 5, 108, 55, 2, 196, 251, 3, 2, 2,
	2, 197, 198, 7, 73, 2, 2, 198, 199, 7, 7, 2, 2, 199, 200, 5, 86, 44, 2,
	200, 201, 7, 8, 2, 2, 201, 202, 5, 8, 5, 2, 202, 251, 3, 2, 2, 2, 203,
	204, 7, 71, 2, 2, 204, 206, 7, 7, 2, 2, 205, 207, 5, 86, 44, 2, 206, 205,
	3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 7, 11,
	2, 2, 209, 211, 5, 86, 44, 2, 210, 209, 3, 2, 2, 2, 210, 211, 3, 2, 2,
	2, 211, 212, 3, 2, 2, 2, 212, 214, 7, 11, 2, 2, 213, 215, 5, 86, 44, 2,
	214, 213, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216,
	217, 7, 8, 2, 2, 217, 251, 5, 8, 5, 2, 218, 219, 7, 71, 2, 2, 219, 220,
	7, 7, 2, 2, 220, 221, 7, 65, 2, 2, 221, 222, 5, 16, 9, 2, 222, 224, 7,
	11, 2, 2, 223, 225, 5, 86, 44, 2, 224, 223, 3, 2, 2, 2, 224, 225, 3, 2,
	2, 2, 225, 226, 3, 2, 2, 2, 226, 228, 7, 11, 2, 2, 227, 229, 5, 86, 44,
	2, 228, 227, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230,
	231, 7, 8, 2, 2, 231, 232, 5, 8, 5, 2, 232, 251, 3, 2, 2, 2, 233, 234,
	7, 71, 2, 2, 234, 235, 7, 7, 2, 2, 235, 236, 5, 88, 45, 2, 236, 237, 7,
	82, 2, 2, 237, 238, 5, 86, 44, 2, 238, 239, 7, 8, 2, 2, 239, 240, 5, 8,
	5, 2, 240, 251, 3, 2, 2, 2, 241, 242, 7, 71, 2, 2, 242, 243, 7, 7, 2, 2,
	243, 244, 7, 65, 2, 2, 244, 245, 5, 18, 10, 2, 245, 246, 7, 82, 2, 2, 246,
	247, 5, 86, 44, 2, 247, 248, 7, 8, 2, 2, 248, 249, 5, 8, 5, 2, 249, 251,
	3, 2, 2, 2, 250, 189, 3, 2, 2, 2, 250, 197, 3, 2, 2, 2, 250, 203, 3, 2,
	2, 2, 250, 218, 3, 2, 2, 2, 250, 233, 3, 2, 2, 2, 250, 241, 3, 2, 2, 2,
	251, 29, 3, 2, 2, 2, 252, 254, 7, 70, 2, 2, 253, 255, 7, 100, 2, 2, 254,
	253, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 257,
	5, 108, 55, 2, 257, 31, 3, 2, 2, 2, 258, 260, 7, 58, 2, 2, 259, 261, 7,
	100, 2, 2, 260, 259, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 262, 3, 2,
	2, 2, 262, 263, 5, 108, 55, 2, 263, 33, 3, 2, 2, 2, 264, 266, 7, 68, 2,
	2, 265, 267, 5, 86, 44, 2, 266, 265, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2,
	267, 268, 3, 2, 2, 2, 268, 269, 5, 108, 55, 2, 269, 35, 3, 2, 2, 2, 270,
	271, 7, 77, 2, 2, 271, 272, 7, 7, 2, 2, 272, 273, 5, 86, 44, 2, 273, 274,
	7, 8, 2, 2, 274, 275, 5, 8, 5, 2, 275, 37, 3, 2, 2, 2, 276, 277, 7, 72,
	2, 2, 277, 278, 7, 7, 2, 2, 278, 279, 5, 86, 44, 2, 279, 280, 7, 8, 2,
	2, 280, 281, 5, 40, 21, 2, 281, 39, 3, 2, 2, 2, 282, 284, 7, 9, 2, 2, 283,
	285, 5, 42, 22, 2, 284, 283, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 290,
	3, 2, 2, 2, 286, 288, 5, 46, 24, 2, 287, 289, 5, 42, 22, 2, 288, 287, 3,
	2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 291, 3, 2, 2, 2, 290, 286, 3, 2, 2,
	2, 290, 291, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 293, 7, 10, 2, 2, 293,
	41, 3, 2, 2, 2, 294, 296, 5, 44, 23, 2, 295, 294, 3, 2, 2, 2, 296, 297,
	3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 43, 3, 2,
	2, 2, 299, 300, 7, 62, 2, 2, 300, 301, 5, 86, 44, 2, 301, 303, 7, 15, 2,
	2, 302, 304, 5, 12, 7, 2, 303, 302, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304,
	45, 3, 2, 2, 2, 305, 306, 7, 78, 2, 2, 306, 308, 7, 15, 2, 2, 307, 309,
	5, 12, 7, 2, 308, 307, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 47, 3, 2,
	2, 2, 310, 311, 7, 100, 2, 2, 311, 312, 7, 15, 2, 2, 312, 313, 5, 8, 5,
	2, 313, 49, 3, 2, 2, 2, 314, 315, 7, 80, 2, 2, 315, 316, 5, 86, 44, 2,
	316, 317, 5, 108, 55, 2, 317, 51, 3, 2, 2, 2, 318, 319, 7, 83, 2, 2, 319,
	320, 5, 10, 6, 2, 320, 321, 5, 54, 28, 2, 321, 332, 3, 2, 2, 2, 322, 323,
	7, 83, 2, 2, 323, 324, 5, 10, 6, 2, 324, 325, 5, 56, 29, 2, 325, 332, 3,
	2, 2, 2, 326, 327, 7, 83, 2, 2, 327, 328, 5, 10, 6, 2, 328, 329, 5, 54,
	28, 2, 329, 330, 5, 56, 29, 2, 330, 332, 3, 2, 2, 2, 331, 318, 3, 2, 2,
	2, 331, 322, 3, 2, 2, 2, 331, 326, 3, 2, 2, 2, 332, 53, 3, 2, 2, 2, 333,
	334, 7, 66, 2, 2, 334, 335, 7, 7, 2, 2, 335, 336, 7, 100, 2, 2, 336, 337,
	7, 8, 2, 2, 337, 338, 5, 10, 6, 2, 338, 55, 3, 2, 2, 2, 339, 340, 7, 67,
	2, 2, 340, 341, 5, 10, 6, 2, 341, 57, 3, 2, 2, 2, 342, 343, 7, 74, 2, 2,
	343, 344, 5, 108, 55, 2, 344, 59, 3, 2, 2, 2, 345, 346, 7, 75, 2, 2, 346,
	347, 7, 100, 2, 2, 347, 349, 7, 7, 2, 2, 348, 350, 5, 62, 32, 2, 349, 348,
	3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 352, 7, 8,
	2, 2, 352, 353, 7, 9, 2, 2, 353, 354, 5, 64, 33, 2, 354, 355, 7, 10, 2,
	2, 355, 61, 3, 2, 2, 2, 356, 361, 7, 100, 2, 2, 357, 358, 7, 12, 2, 2,
	358, 360, 7, 100, 2, 2, 359, 357, 3, 2, 2, 2, 360, 363, 3, 2, 2, 2, 361,
	359, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 63, 3, 2, 2, 2, 363, 361, 3,
	2, 2, 2, 364, 366, 5, 4, 3, 2, 365, 364, 3, 2, 2, 2, 365, 366, 3, 2, 2,
	2, 366, 65, 3, 2, 2, 2, 367, 369, 7, 5, 2, 2, 368, 370, 5, 68, 35, 2, 369,
	368, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 372, 3, 2, 2, 2, 371, 373,
	7, 12, 2, 2, 372, 371, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373, 375, 3, 2,
	2, 2, 374, 376, 5, 70, 36, 2, 375, 374, 3, 2, 2, 2, 375, 376, 3, 2, 2,
	2, 376, 377, 3, 2, 2, 2, 377, 378, 7, 6, 2, 2, 378, 67, 3, 2, 2, 2, 379,
	381, 5, 70, 36, 2, 380, 379, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 382,
	3, 2, 2, 2, 382, 390, 5, 88, 45, 2, 383, 385, 7, 12, 2, 2, 384, 386, 5,
	70, 36, 2, 385, 384, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 387, 3, 2,
	2, 2, 387, 389, 5, 88, 45, 2, 388, 383, 3, 2, 2, 2, 389, 392, 3, 2, 2,
	2, 390, 388, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 69, 3, 2, 2, 2, 392,
	390, 3, 2, 2, 2, 393, 395, 7, 12, 2, 2, 394, 393, 3, 2, 2, 2, 395, 396,
	3, 2, 2, 2, 396, 394, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397, 71, 3, 2,
	2, 2, 398, 399, 7, 9, 2, 2, 399, 408, 7, 10, 2, 2, 400, 401, 7, 9, 2, 2,
	401, 403, 5, 74, 38, 2, 402, 404, 7, 12, 2, 2, 403, 402, 3, 2, 2, 2, 403,
	404, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 406, 7, 10, 2, 2, 406, 408,
	3, 2, 2, 2, 407, 398, 3, 2, 2, 2, 407, 400, 3, 2, 2, 2, 408, 73, 3, 2,
	2, 2, 409, 414, 5, 76, 39, 2, 410, 411, 7, 12, 2, 2, 411, 413, 5, 76, 39,
	2, 412, 410, 3, 2, 2, 2, 413, 416, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2, 414,
	415, 3, 2, 2, 2, 415, 75, 3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 417, 418, 5,
	78, 40, 2, 418, 419, 7, 15, 2, 2, 419, 420, 5, 88, 45, 2, 420, 437, 3,
	2, 2, 2, 421, 422, 5, 104, 53, 2, 422, 423, 7, 7, 2, 2, 423, 424, 7, 8,
	2, 2, 424, 425, 7, 9, 2, 2, 425, 426, 5, 64, 33, 2, 426, 427, 7, 10, 2,
	2, 427, 437, 3, 2, 2, 2, 428, 429, 5, 106, 54, 2, 429, 430, 7, 7, 2, 2,
	430, 431, 5, 80, 41, 2, 431, 432, 7, 8, 2, 2, 432, 433, 7, 9, 2, 2, 433,
	434, 5, 64, 33, 2, 434, 435, 7, 10, 2, 2, 435, 437, 3, 2, 2, 2, 436, 417,
	3, 2, 2, 2, 436, 421, 3, 2, 2, 2, 436, 428, 3, 2, 2, 2, 437, 77, 3, 2,
	2, 2, 438, 442, 5, 96, 49, 2, 439, 442, 7, 101, 2, 2, 440, 442, 5, 94,
	48, 2, 441, 438, 3, 2, 2, 2, 441, 439, 3, 2, 2, 2, 441, 440, 3, 2, 2, 2,
	442, 79, 3, 2, 2, 2, 443, 444, 7, 100, 2, 2, 444, 81, 3, 2, 2, 2, 445,
	447, 7, 7, 2, 2, 446, 448, 5, 84, 43, 2, 447, 446, 3, 2, 2, 2, 447, 448,
	3, 2, 2, 2, 448, 449, 3, 2, 2, 2, 449, 450, 7, 8, 2, 2, 450, 83, 3, 2,
	2, 2, 451, 456, 5, 88, 45, 2, 452, 453, 7, 12, 2, 2, 453, 455, 5, 88, 45,
	2, 454, 452, 3, 2, 2, 2, 455, 458, 3, 2, 2, 2, 456, 454, 3, 2, 2, 2, 456,
	457, 3, 2, 2, 2, 457, 85, 3, 2, 2, 2, 458, 456, 3, 2, 2, 2, 459, 464, 5,
	88, 45, 2, 460, 461, 7, 12, 2, 2, 461, 463, 5, 88, 45, 2, 462, 460, 3,
	2, 2, 2, 463, 466, 3, 2, 2, 2, 464, 462, 3, 2, 2, 2, 464, 465, 3, 2, 2,
	2, 465, 87, 3, 2, 2, 2, 466, 464, 3, 2, 2, 2, 467, 468, 8, 45, 1, 2, 468,
	470, 7, 75, 2, 2, 469, 471, 7, 100, 2, 2, 470, 469, 3, 2, 2, 2, 470, 471,
	3, 2, 2, 2, 471, 472, 3, 2, 2, 2, 472, 474, 7, 7, 2, 2, 473, 475, 5, 62,
	32, 2, 474, 473, 3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475, 476, 3, 2, 2, 2,
	476, 477, 7, 8, 2, 2, 477, 478, 7, 9, 2, 2, 478, 479, 5, 64, 33, 2, 479,
	480, 7, 10, 2, 2, 480, 514, 3, 2, 2, 2, 481, 482, 7, 64, 2, 2, 482, 484,
	5, 88, 45, 2, 483, 485, 5, 82, 42, 2, 484, 483, 3, 2, 2, 2, 484, 485, 3,
	2, 2, 2, 485, 514, 3, 2, 2, 2, 486, 487, 7, 81, 2, 2, 487, 514, 5, 88,
	45, 32, 488, 489, 7, 69, 2, 2, 489, 514, 5, 88, 45, 31, 490, 491, 7, 61,
	2, 2, 491, 514, 5, 88, 45, 30, 492, 493, 7, 17, 2, 2, 493, 514, 5, 88,
	45, 29, 494, 495, 7, 18, 2, 2, 495, 514, 5, 88, 45, 28, 496, 497, 7, 19,
	2, 2, 497, 514, 5, 88, 45, 27, 498, 499, 7, 20, 2, 2, 499, 514, 5, 88,
	45, 26, 500, 501, 7, 21, 2, 2, 501, 514, 5, 88, 45, 25, 502, 503, 7, 22,
	2, 2, 503, 514, 5, 88, 45, 24, 504, 514, 7, 76, 2, 2, 505, 514, 7, 100,
	2, 2, 506, 514, 5, 92, 47, 2, 507, 514, 5, 66, 34, 2, 508, 514, 5, 72,
	37, 2, 509, 510, 7, 7, 2, 2, 510, 511, 5, 86, 44, 2, 511, 512, 7, 8, 2,
	2, 512, 514, 3, 2, 2, 2, 513, 467, 3, 2, 2, 2, 513, 481, 3, 2, 2, 2, 513,
	486, 3, 2, 2, 2, 513, 488, 3, 2, 2, 2, 513, 490, 3, 2, 2, 2, 513, 492,
	3, 2, 2, 2, 513, 494, 3, 2, 2, 2, 513, 496, 3, 2, 2, 2, 513, 498, 3, 2,
	2, 2, 513, 500, 3, 2, 2, 2, 513, 502, 3, 2, 2, 2, 513, 504, 3, 2, 2, 2,
	513, 505, 3, 2, 2, 2, 513, 506, 3, 2, 2, 2, 513, 507, 3, 2, 2, 2, 513,
	508, 3, 2, 2, 2, 513, 509, 3, 2, 2, 2, 514, 582, 3, 2, 2, 2, 515, 516,
	12, 23, 2, 2, 516, 517, 9, 2, 2, 2, 517, 581, 5, 88, 45, 24, 518, 519,
	12, 22, 2, 2, 519, 520, 9, 3, 2, 2, 520, 581, 5, 88, 45, 23, 521, 522,
	12, 21, 2, 2, 522, 523, 9, 4, 2, 2, 523, 581, 5, 88, 45, 22, 524, 525,
	12, 20, 2, 2, 525, 526, 9, 5, 2, 2, 526, 581, 5, 88, 45, 21, 527, 528,
	12, 19, 2, 2, 528, 529, 7, 60, 2, 2, 529, 581, 5, 88, 45, 20, 530, 531,
	12, 18, 2, 2, 531, 532, 7, 82, 2, 2, 532, 581, 5, 88, 45, 19, 533, 534,
	12, 17, 2, 2, 534, 535, 9, 6, 2, 2, 535, 581, 5, 88, 45, 18, 536, 537,
	12, 16, 2, 2, 537, 538, 7, 37, 2, 2, 538, 581, 5, 88, 45, 17, 539, 540,
	12, 15, 2, 2, 540, 541, 7, 38, 2, 2, 541, 581, 5, 88, 45, 16, 542, 543,
	12, 14, 2, 2, 543, 544, 7, 39, 2, 2, 544, 581, 5, 88, 45, 15, 545, 546,
	12, 13, 2, 2, 546, 547, 7, 40, 2, 2, 547, 581, 5, 88, 45, 14, 548, 549,
	12, 12, 2, 2, 549, 550, 7, 41, 2, 2, 550, 581, 5, 88, 45, 13, 551, 552,
	12, 11, 2, 2, 552, 553, 7, 14, 2, 2, 553, 554, 5, 88, 45, 2, 554, 555,
	7, 15, 2, 2, 555, 556, 5, 88, 45, 12, 556, 581, 3, 2, 2, 2, 557, 558, 12,
	10, 2, 2, 558, 559, 7, 13, 2, 2, 559, 581, 5, 88, 45, 11, 560, 561, 12,
	9, 2, 2, 561, 562, 5, 90, 46, 2, 562, 563, 5, 88, 45, 10, 563, 581, 3,
	2, 2, 2, 564, 565, 12, 38, 2, 2, 565, 566, 7, 5, 2, 2, 566, 567, 5, 86,
	44, 2, 567, 568, 7, 6, 2, 2, 568, 581, 3, 2, 2, 2, 569, 570, 12, 37, 2,
	2, 570, 571, 7, 16, 2, 2, 571, 581, 5, 96, 49, 2, 572, 573, 12, 36, 2,
	2, 573, 581, 5, 82, 42, 2, 574, 575, 12, 34, 2, 2, 575, 576, 6, 45, 23,
	2, 576, 581, 7, 17, 2, 2, 577, 578, 12, 33, 2, 2, 578, 579, 6, 45, 25,
	2, 579, 581, 7, 18, 2, 2, 580, 515, 3, 2, 2, 2, 580, 518, 3, 2, 2, 2, 580,
	521, 3, 2, 2, 2, 580, 524, 3, 2, 2, 2, 580, 527, 3, 2, 2, 2, 580, 530,
	3, 2, 2, 2, 580, 533, 3, 2, 2, 2, 580, 536, 3, 2, 2, 2, 580, 539, 3, 2,
	2, 2, 580, 542, 3, 2, 2, 2, 580, 545, 3, 2, 2, 2, 580, 548, 3, 2, 2, 2,
	580, 551, 3, 2, 2, 2, 580, 557, 3, 2, 2, 2, 580, 560, 3, 2, 2, 2, 580,
	564, 3, 2, 2, 2, 580, 569, 3, 2, 2, 2, 580, 572, 3, 2, 2, 2, 580, 574,
	3, 2, 2, 2, 580, 577, 3, 2, 2, 2, 581, 584, 3, 2, 2, 2, 582, 580, 3, 2,
	2, 2, 582, 583, 3, 2, 2, 2, 583, 89, 3, 2, 2, 2, 584, 582, 3, 2, 2, 2,
	585, 586, 9, 7, 2, 2, 586, 91, 3, 2, 2, 2, 587, 590, 9, 8, 2, 2, 588, 590,
	5, 94, 48, 2, 589, 587, 3, 2, 2, 2, 589, 588, 3, 2, 2, 2, 590, 93, 3, 2,
	2, 2, 591, 592, 9, 9, 2, 2, 592, 95, 3, 2, 2, 2, 593, 596, 7, 100, 2, 2,
	594, 596, 5, 98, 50, 2, 595, 593, 3, 2, 2, 2, 595, 594, 3, 2, 2, 2, 596,
	97, 3, 2, 2, 2, 597, 601, 5, 100, 51, 2, 598, 601, 5, 102, 52, 2, 599,
	601, 9, 10, 2, 2, 600, 597, 3, 2, 2, 2, 600, 598, 3, 2, 2, 2, 600, 599,
	3, 2, 2, 2, 601, 99, 3, 2, 2, 2, 602, 603, 9, 11, 2, 2, 603, 101, 3, 2,
	2, 2, 604, 605, 9, 12, 2, 2, 605, 103, 3, 2, 2, 2, 606, 607, 6, 53, 26,
	2, 607, 608, 7, 100, 2, 2, 608, 609, 5, 78, 40, 2, 609, 105, 3, 2, 2, 2,
	610, 611, 6, 54, 27, 2, 611, 612, 7, 100, 2, 2, 612, 613, 5, 78, 40, 2,
	613, 107, 3, 2, 2, 2, 614, 619, 7, 11, 2, 2, 615, 619, 7, 2, 2, 3, 616,
	619, 6, 55, 28, 2, 617, 619, 6, 55, 29, 2, 618, 614, 3, 2, 2, 2, 618, 615,
	3, 2, 2, 2, 618, 616, 3, 2, 2, 2, 618, 617, 3, 2, 2, 2, 619, 109, 3, 2,
	2, 2, 620, 621, 7, 2, 2, 3, 621, 111, 3, 2, 2, 2, 55, 113, 120, 125, 143,
	147, 154, 165, 170, 187, 206, 210, 214, 224, 228, 250, 254, 260, 266, 284,
	288, 290, 297, 303, 308, 331, 349, 361, 365, 369, 372, 375, 380, 385, 390,
	396, 403, 407, 414, 436, 441, 447, 456, 464, 470, 474, 484, 513, 580, 582,
	589, 595, 600, 618,
}
var literalNames = []string{
	"", "", "", "'['", "']'", "'('", "')'", "'{'", "'}'", "';'", "','", "'='",
	"'?'", "':'", "'.'", "'++'", "'--'", "'+'", "'-'", "'~'", "'!'", "'*'",
	"'/'", "'%'", "'>>'", "'<<'", "'>>>'", "'<'", "'>'", "'<='", "'>='", "'=='",
	"'!='", "'==='", "'!=='", "'&'", "'^'", "'|'", "'&&'", "'||'", "'*='",
	"'/='", "'%='", "'+='", "'-='", "'<<='", "'>>='", "'>>>='", "'&='", "'^='",
	"'|='", "'null'", "", "", "", "", "'break'", "'do'", "'instanceof'", "'typeof'",
	"'case'", "'else'", "'new'", "'var'", "'catch'", "'finally'", "'return'",
	"'void'", "'continue'", "'for'", "'switch'", "'while'", "'debugger'", "'function'",
	"'this'", "'with'", "'default'", "'if'", "'throw'", "'delete'", "'in'",
	"'try'", "'class'", "'enum'", "'extends'", "'super'", "'const'", "'export'",
	"'import'",
}
var symbolicNames = []string{
	"", "RegularExpressionLiteral", "LineTerminator", "OpenBracket", "CloseBracket",
	"OpenParen", "CloseParen", "OpenBrace", "CloseBrace", "SemiColon", "Comma",
	"Assign", "QuestionMark", "Colon", "Dot", "PlusPlus", "MinusMinus", "Plus",
	"Minus", "BitNot", "Not", "Multiply", "Divide", "Modulus", "RightShiftArithmetic",
	"LeftShiftArithmetic", "RightShiftLogical", "LessThan", "MoreThan", "LessThanEquals",
	"GreaterThanEquals", "Equals", "NotEquals", "IdentityEquals", "IdentityNotEquals",
	"BitAnd", "BitXOr", "BitOr", "And", "Or", "MultiplyAssign", "DivideAssign",
	"ModulusAssign", "PlusAssign", "MinusAssign", "LeftShiftArithmeticAssign",
	"RightShiftArithmeticAssign", "RightShiftLogicalAssign", "BitAndAssign",
	"BitXorAssign", "BitOrAssign", "NullLiteral", "BooleanLiteral", "DecimalLiteral",
	"HexIntegerLiteral", "OctalIntegerLiteral", "Break", "Do", "Instanceof",
	"Typeof", "Case", "Else", "New", "Var", "Catch", "Finally", "Return", "Void",
	"Continue", "For", "Switch", "While", "Debugger", "Function", "This", "With",
	"Default", "If", "Throw", "Delete", "In", "Try", "Class", "Enum", "Extends",
	"Super", "Const", "Export", "Import", "Implements", "Let", "Private", "Public",
	"Interface", "Package", "Protected", "Static", "Yield", "Identifier", "StringLiteral",
	"WhiteSpaces", "MultiLineComment", "SingleLineComment", "UnexpectedCharacter",
}

var ruleNames = []string{
	"program", "sourceElements", "sourceElement", "statement", "block", "statementList",
	"variableStatement", "variableDeclarationList", "variableDeclaration",
	"initialiser", "voidStatement", "expressionStatement", "ifStatement", "iterationStatement",
	"continueStatement", "breakStatement", "returnStatement", "withStatement",
	"switchStatement", "caseBlock", "caseClauses", "caseClause", "defaultClause",
	"labelledStatement", "throwStatement", "tryStatement", "catchProduction",
	"finallyProduction", "debuggerStatement", "functionDeclaration", "formalParameterList",
	"functionBody", "arrayLiteral", "elementList", "elision", "objectLiteral",
	"propertyNameAndValueList", "propertyAssignment", "propertyName", "propertySetParameterList",
	"arguments", "argumentList", "expressionSequence", "singleExpression",
	"assignmentOperator", "literal", "numericLiteral", "identifierName", "reservedWord",
	"keyword", "futureReservedWord", "getter", "setter", "eos", "eof",
}

type ECMAScriptParser struct {
	*antlr.BaseParser
}

// NewECMAScriptParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ECMAScriptParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewECMAScriptParser(input antlr.TokenStream) *ECMAScriptParser {
	this := new(ECMAScriptParser)
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
	this.GrammarFileName = "ECMAScript.g4"

	return this
}

// here returns `true` iff on the current index of the parser's
// token stream a token of the given `type` exists on the
// `HIDDEN` channel.
//
// Args:
//  type (int): the type of the token on the `HIDDEN` channel
//      to check.
//
//  Returns:
//      `True` iff on the current index of the parser's
//      token stream a token of the given `type` exists on the
//      `HIDDEN` channel.
func (p *ECMAScriptParser) here(tokenType int) bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	// Check if the token resides on the HIDDEN channel and if it is of the
	// provided tokenType.
	return (ahead.GetChannel() == antlr.LexerHidden) && (ahead.GetTokenType() == tokenType)
}

// lineTerminatorAhead returns `true` iff on the current index of the parser's
// token stream a token exists on the `HIDDEN` channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func (p *ECMAScriptParser) lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return false
	}

	if ahead.GetTokenType() == ECMAScriptParserLineTerminator {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == ECMAScriptParserWhiteSpaces {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	tokenType := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	if tokenType == ECMAScriptParserMultiLineComment && strings.ContainsAny(text, "\r\n") {
		return true
	}

	return tokenType == ECMAScriptParserLineTerminator
}

// ECMAScriptParser tokens.
const (
	ECMAScriptParserEOF                        = antlr.TokenEOF
	ECMAScriptParserRegularExpressionLiteral   = 1
	ECMAScriptParserLineTerminator             = 2
	ECMAScriptParserOpenBracket                = 3
	ECMAScriptParserCloseBracket               = 4
	ECMAScriptParserOpenParen                  = 5
	ECMAScriptParserCloseParen                 = 6
	ECMAScriptParserOpenBrace                  = 7
	ECMAScriptParserCloseBrace                 = 8
	ECMAScriptParserSemiColon                  = 9
	ECMAScriptParserComma                      = 10
	ECMAScriptParserAssign                     = 11
	ECMAScriptParserQuestionMark               = 12
	ECMAScriptParserColon                      = 13
	ECMAScriptParserDot                        = 14
	ECMAScriptParserPlusPlus                   = 15
	ECMAScriptParserMinusMinus                 = 16
	ECMAScriptParserPlus                       = 17
	ECMAScriptParserMinus                      = 18
	ECMAScriptParserBitNot                     = 19
	ECMAScriptParserNot                        = 20
	ECMAScriptParserMultiply                   = 21
	ECMAScriptParserDivide                     = 22
	ECMAScriptParserModulus                    = 23
	ECMAScriptParserRightShiftArithmetic       = 24
	ECMAScriptParserLeftShiftArithmetic        = 25
	ECMAScriptParserRightShiftLogical          = 26
	ECMAScriptParserLessThan                   = 27
	ECMAScriptParserMoreThan                   = 28
	ECMAScriptParserLessThanEquals             = 29
	ECMAScriptParserGreaterThanEquals          = 30
	ECMAScriptParserEquals                     = 31
	ECMAScriptParserNotEquals                  = 32
	ECMAScriptParserIdentityEquals             = 33
	ECMAScriptParserIdentityNotEquals          = 34
	ECMAScriptParserBitAnd                     = 35
	ECMAScriptParserBitXOr                     = 36
	ECMAScriptParserBitOr                      = 37
	ECMAScriptParserAnd                        = 38
	ECMAScriptParserOr                         = 39
	ECMAScriptParserMultiplyAssign             = 40
	ECMAScriptParserDivideAssign               = 41
	ECMAScriptParserModulusAssign              = 42
	ECMAScriptParserPlusAssign                 = 43
	ECMAScriptParserMinusAssign                = 44
	ECMAScriptParserLeftShiftArithmeticAssign  = 45
	ECMAScriptParserRightShiftArithmeticAssign = 46
	ECMAScriptParserRightShiftLogicalAssign    = 47
	ECMAScriptParserBitAndAssign               = 48
	ECMAScriptParserBitXorAssign               = 49
	ECMAScriptParserBitOrAssign                = 50
	ECMAScriptParserNullLiteral                = 51
	ECMAScriptParserBooleanLiteral             = 52
	ECMAScriptParserDecimalLiteral             = 53
	ECMAScriptParserHexIntegerLiteral          = 54
	ECMAScriptParserOctalIntegerLiteral        = 55
	ECMAScriptParserBreak                      = 56
	ECMAScriptParserDo                         = 57
	ECMAScriptParserInstanceof                 = 58
	ECMAScriptParserTypeof                     = 59
	ECMAScriptParserCase                       = 60
	ECMAScriptParserElse                       = 61
	ECMAScriptParserNew                        = 62
	ECMAScriptParserVar                        = 63
	ECMAScriptParserCatch                      = 64
	ECMAScriptParserFinally                    = 65
	ECMAScriptParserReturn                     = 66
	ECMAScriptParserVoid                       = 67
	ECMAScriptParserContinue                   = 68
	ECMAScriptParserFor                        = 69
	ECMAScriptParserSwitch                     = 70
	ECMAScriptParserWhile                      = 71
	ECMAScriptParserDebugger                   = 72
	ECMAScriptParserFunction                   = 73
	ECMAScriptParserThis                       = 74
	ECMAScriptParserWith                       = 75
	ECMAScriptParserDefault                    = 76
	ECMAScriptParserIf                         = 77
	ECMAScriptParserThrow                      = 78
	ECMAScriptParserDelete                     = 79
	ECMAScriptParserIn                         = 80
	ECMAScriptParserTry                        = 81
	ECMAScriptParserClass                      = 82
	ECMAScriptParserEnum                       = 83
	ECMAScriptParserExtends                    = 84
	ECMAScriptParserSuper                      = 85
	ECMAScriptParserConst                      = 86
	ECMAScriptParserExport                     = 87
	ECMAScriptParserImport                     = 88
	ECMAScriptParserImplements                 = 89
	ECMAScriptParserLet                        = 90
	ECMAScriptParserPrivate                    = 91
	ECMAScriptParserPublic                     = 92
	ECMAScriptParserInterface                  = 93
	ECMAScriptParserPackage                    = 94
	ECMAScriptParserProtected                  = 95
	ECMAScriptParserStatic                     = 96
	ECMAScriptParserYield                      = 97
	ECMAScriptParserIdentifier                 = 98
	ECMAScriptParserStringLiteral              = 99
	ECMAScriptParserWhiteSpaces                = 100
	ECMAScriptParserMultiLineComment           = 101
	ECMAScriptParserSingleLineComment          = 102
	ECMAScriptParserUnexpectedCharacter        = 103
)

// ECMAScriptParser rules.
const (
	ECMAScriptParserRULE_program                  = 0
	ECMAScriptParserRULE_sourceElements           = 1
	ECMAScriptParserRULE_sourceElement            = 2
	ECMAScriptParserRULE_statement                = 3
	ECMAScriptParserRULE_block                    = 4
	ECMAScriptParserRULE_statementList            = 5
	ECMAScriptParserRULE_variableStatement        = 6
	ECMAScriptParserRULE_variableDeclarationList  = 7
	ECMAScriptParserRULE_variableDeclaration      = 8
	ECMAScriptParserRULE_initialiser              = 9
	ECMAScriptParserRULE_voidStatement            = 10
	ECMAScriptParserRULE_expressionStatement      = 11
	ECMAScriptParserRULE_ifStatement              = 12
	ECMAScriptParserRULE_iterationStatement       = 13
	ECMAScriptParserRULE_continueStatement        = 14
	ECMAScriptParserRULE_breakStatement           = 15
	ECMAScriptParserRULE_returnStatement          = 16
	ECMAScriptParserRULE_withStatement            = 17
	ECMAScriptParserRULE_switchStatement          = 18
	ECMAScriptParserRULE_caseBlock                = 19
	ECMAScriptParserRULE_caseClauses              = 20
	ECMAScriptParserRULE_caseClause               = 21
	ECMAScriptParserRULE_defaultClause            = 22
	ECMAScriptParserRULE_labelledStatement        = 23
	ECMAScriptParserRULE_throwStatement           = 24
	ECMAScriptParserRULE_tryStatement             = 25
	ECMAScriptParserRULE_catchProduction          = 26
	ECMAScriptParserRULE_finallyProduction        = 27
	ECMAScriptParserRULE_debuggerStatement        = 28
	ECMAScriptParserRULE_functionDeclaration      = 29
	ECMAScriptParserRULE_formalParameterList      = 30
	ECMAScriptParserRULE_functionBody             = 31
	ECMAScriptParserRULE_arrayLiteral             = 32
	ECMAScriptParserRULE_elementList              = 33
	ECMAScriptParserRULE_elision                  = 34
	ECMAScriptParserRULE_objectLiteral            = 35
	ECMAScriptParserRULE_propertyNameAndValueList = 36
	ECMAScriptParserRULE_propertyAssignment       = 37
	ECMAScriptParserRULE_propertyName             = 38
	ECMAScriptParserRULE_propertySetParameterList = 39
	ECMAScriptParserRULE_arguments                = 40
	ECMAScriptParserRULE_argumentList             = 41
	ECMAScriptParserRULE_expressionSequence       = 42
	ECMAScriptParserRULE_singleExpression         = 43
	ECMAScriptParserRULE_assignmentOperator       = 44
	ECMAScriptParserRULE_literal                  = 45
	ECMAScriptParserRULE_numericLiteral           = 46
	ECMAScriptParserRULE_identifierName           = 47
	ECMAScriptParserRULE_reservedWord             = 48
	ECMAScriptParserRULE_keyword                  = 49
	ECMAScriptParserRULE_futureReservedWord       = 50
	ECMAScriptParserRULE_getter                   = 51
	ECMAScriptParserRULE_setter                   = 52
	ECMAScriptParserRULE_eos                      = 53
	ECMAScriptParserRULE_eof                      = 54
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
	p.RuleIndex = ECMAScriptParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserEOF, 0)
}

func (s *ProgramContext) SourceElements() ISourceElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceElementsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *ECMAScriptParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ECMAScriptParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(110)
			p.SourceElements()
		}

	}
	{
		p.SetState(113)
		p.Match(ECMAScriptParserEOF)
	}

	return localctx
}

// ISourceElementsContext is an interface to support dynamic dispatch.
type ISourceElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceElementsContext differentiates from other interfaces.
	IsSourceElementsContext()
}

type SourceElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceElementsContext() *SourceElementsContext {
	var p = new(SourceElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_sourceElements
	return p
}

func (*SourceElementsContext) IsSourceElementsContext() {}

func NewSourceElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceElementsContext {
	var p = new(SourceElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_sourceElements

	return p
}

func (s *SourceElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceElementsContext) AllSourceElement() []ISourceElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceElementContext)(nil)).Elem())
	var tst = make([]ISourceElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceElementContext)
		}
	}

	return tst
}

func (s *SourceElementsContext) SourceElement(i int) ISourceElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceElementContext)
}

func (s *SourceElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterSourceElements(s)
	}
}

func (s *SourceElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitSourceElements(s)
	}
}

func (p *ECMAScriptParser) SourceElements() (localctx ISourceElementsContext) {
	this := p
	_ = this

	localctx = NewSourceElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ECMAScriptParserRULE_sourceElements)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(115)
				p.SourceElement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// ISourceElementContext is an interface to support dynamic dispatch.
type ISourceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceElementContext differentiates from other interfaces.
	IsSourceElementContext()
}

type SourceElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceElementContext() *SourceElementContext {
	var p = new(SourceElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_sourceElement
	return p
}

func (*SourceElementContext) IsSourceElementContext() {}

func NewSourceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceElementContext {
	var p = new(SourceElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_sourceElement

	return p
}

func (s *SourceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceElementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SourceElementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *SourceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterSourceElement(s)
	}
}

func (s *SourceElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitSourceElement(s)
	}
}

func (p *ECMAScriptParser) SourceElement() (localctx ISourceElementContext) {
	this := p
	_ = this

	localctx = NewSourceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ECMAScriptParserRULE_sourceElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(120)

		if !(p.GetInputStream().LA(1) != ECMAScriptParserFunction) {
			panic(antlr.NewFailedPredicateException(p, "p.GetInputStream().LA(1) != ECMAScriptParserFunction", ""))
		}
		{
			p.SetState(121)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.FunctionDeclaration()
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
	p.RuleIndex = ECMAScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) VariableStatement() IVariableStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableStatementContext)
}

func (s *StatementContext) VoidStatement() IVoidStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVoidStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVoidStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) IterationStatement() IIterationStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterationStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinueStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) WithStatement() IWithStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithStatementContext)
}

func (s *StatementContext) LabelledStatement() ILabelledStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelledStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelledStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) ThrowStatement() IThrowStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThrowStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThrowStatementContext)
}

func (s *StatementContext) TryStatement() ITryStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITryStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITryStatementContext)
}

func (s *StatementContext) DebuggerStatement() IDebuggerStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDebuggerStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDebuggerStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *ECMAScriptParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ECMAScriptParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.VariableStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)
			p.VoidStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(128)

		if !(p.GetInputStream().LA(1) != ECMAScriptParserOpenBrace) {
			panic(antlr.NewFailedPredicateException(p, "p.GetInputStream().LA(1) != ECMAScriptParserOpenBrace", ""))
		}
		{
			p.SetState(129)
			p.ExpressionStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(130)
			p.IfStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(131)
			p.IterationStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(132)
			p.ContinueStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(133)
			p.BreakStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(134)
			p.ReturnStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(135)
			p.WithStatement()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(136)
			p.LabelledStatement()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(137)
			p.SwitchStatement()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(138)
			p.ThrowStatement()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(139)
			p.TryStatement()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(140)
			p.DebuggerStatement()
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenBrace, 0)
}

func (s *BlockContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseBrace, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *ECMAScriptParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ECMAScriptParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(ECMAScriptParserOpenBrace)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(144)
			p.StatementList()
		}

	}
	{
		p.SetState(147)
		p.Match(ECMAScriptParserCloseBrace)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *ECMAScriptParser) StatementList() (localctx IStatementListContext) {
	this := p
	_ = this

	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ECMAScriptParserRULE_statementList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(149)
				p.Statement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IVariableStatementContext is an interface to support dynamic dispatch.
type IVariableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableStatementContext differentiates from other interfaces.
	IsVariableStatementContext()
}

type VariableStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableStatementContext() *VariableStatementContext {
	var p = new(VariableStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_variableStatement
	return p
}

func (*VariableStatementContext) IsVariableStatementContext() {}

func NewVariableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableStatementContext {
	var p = new(VariableStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_variableStatement

	return p
}

func (s *VariableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVar, 0)
}

func (s *VariableStatementContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *VariableStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *VariableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVariableStatement(s)
	}
}

func (s *VariableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVariableStatement(s)
	}
}

func (p *ECMAScriptParser) VariableStatement() (localctx IVariableStatementContext) {
	this := p
	_ = this

	localctx = NewVariableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ECMAScriptParserRULE_variableStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(ECMAScriptParserVar)
	}
	{
		p.SetState(155)
		p.VariableDeclarationList()
	}
	{
		p.SetState(156)
		p.Eos()
	}

	return localctx
}

// IVariableDeclarationListContext is an interface to support dynamic dispatch.
type IVariableDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationListContext differentiates from other interfaces.
	IsVariableDeclarationListContext()
}

type VariableDeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationListContext() *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_variableDeclarationList
	return p
}

func (*VariableDeclarationListContext) IsVariableDeclarationListContext() {}

func NewVariableDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_variableDeclarationList

	return p
}

func (s *VariableDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationListContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *VariableDeclarationListContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *VariableDeclarationListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserComma)
}

func (s *VariableDeclarationListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserComma, i)
}

func (s *VariableDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVariableDeclarationList(s)
	}
}

func (s *VariableDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVariableDeclarationList(s)
	}
}

func (p *ECMAScriptParser) VariableDeclarationList() (localctx IVariableDeclarationListContext) {
	this := p
	_ = this

	localctx = NewVariableDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ECMAScriptParserRULE_variableDeclarationList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.VariableDeclaration()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(159)
				p.Match(ECMAScriptParserComma)
			}
			{
				p.SetState(160)
				p.VariableDeclaration()
			}

		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *VariableDeclarationContext) Initialiser() IInitialiserContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitialiserContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitialiserContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *ECMAScriptParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	this := p
	_ = this

	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ECMAScriptParserRULE_variableDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(ECMAScriptParserIdentifier)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(167)
			p.Initialiser()
		}

	}

	return localctx
}

// IInitialiserContext is an interface to support dynamic dispatch.
type IInitialiserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitialiserContext differentiates from other interfaces.
	IsInitialiserContext()
}

type InitialiserContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitialiserContext() *InitialiserContext {
	var p = new(InitialiserContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_initialiser
	return p
}

func (*InitialiserContext) IsInitialiserContext() {}

func NewInitialiserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitialiserContext {
	var p = new(InitialiserContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_initialiser

	return p
}

func (s *InitialiserContext) GetParser() antlr.Parser { return s.parser }

func (s *InitialiserContext) Assign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserAssign, 0)
}

func (s *InitialiserContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *InitialiserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitialiserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitialiserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterInitialiser(s)
	}
}

func (s *InitialiserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitInitialiser(s)
	}
}

func (p *ECMAScriptParser) Initialiser() (localctx IInitialiserContext) {
	this := p
	_ = this

	localctx = NewInitialiserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ECMAScriptParserRULE_initialiser)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(ECMAScriptParserAssign)
	}
	{
		p.SetState(171)
		p.singleExpression(0)
	}

	return localctx
}

// IVoidStatementContext is an interface to support dynamic dispatch.
type IVoidStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVoidStatementContext differentiates from other interfaces.
	IsVoidStatementContext()
}

type VoidStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVoidStatementContext() *VoidStatementContext {
	var p = new(VoidStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_voidStatement
	return p
}

func (*VoidStatementContext) IsVoidStatementContext() {}

func NewVoidStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VoidStatementContext {
	var p = new(VoidStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_voidStatement

	return p
}

func (s *VoidStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VoidStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSemiColon, 0)
}

func (s *VoidStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VoidStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVoidStatement(s)
	}
}

func (s *VoidStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVoidStatement(s)
	}
}

func (p *ECMAScriptParser) VoidStatement() (localctx IVoidStatementContext) {
	this := p
	_ = this

	localctx = NewVoidStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ECMAScriptParserRULE_voidStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(ECMAScriptParserSemiColon)
	}

	return localctx
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_expressionStatement
	return p
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ExpressionStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (p *ECMAScriptParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	this := p
	_ = this

	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ECMAScriptParserRULE_expressionStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.ExpressionSequence()
	}
	{
		p.SetState(176)
		p.Eos()
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIf, 0)
}

func (s *IfStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *IfStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *IfStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserElse, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *ECMAScriptParser) IfStatement() (localctx IIfStatementContext) {
	this := p
	_ = this

	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ECMAScriptParserRULE_ifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(ECMAScriptParserIf)
	}
	{
		p.SetState(179)
		p.Match(ECMAScriptParserOpenParen)
	}
	{
		p.SetState(180)
		p.ExpressionSequence()
	}
	{
		p.SetState(181)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(182)
		p.Statement()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(183)
			p.Match(ECMAScriptParserElse)
		}
		{
			p.SetState(184)
			p.Statement()
		}

	}

	return localctx
}

// IIterationStatementContext is an interface to support dynamic dispatch.
type IIterationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationStatementContext differentiates from other interfaces.
	IsIterationStatementContext()
}

type IterationStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationStatementContext() *IterationStatementContext {
	var p = new(IterationStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_iterationStatement
	return p
}

func (*IterationStatementContext) IsIterationStatementContext() {}

func NewIterationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationStatementContext {
	var p = new(IterationStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_iterationStatement

	return p
}

func (s *IterationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationStatementContext) CopyFrom(ctx *IterationStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IterationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DoStatementContext struct {
	*IterationStatementContext
}

func NewDoStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoStatementContext {
	var p = new(DoStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *DoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDo, 0)
}

func (s *DoStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DoStatementContext) While() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWhile, 0)
}

func (s *DoStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *DoStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *DoStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *DoStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *DoStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterDoStatement(s)
	}
}

func (s *DoStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitDoStatement(s)
	}
}

type ForVarStatementContext struct {
	*IterationStatementContext
}

func NewForVarStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForVarStatementContext {
	var p = new(ForVarStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForVarStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForVarStatementContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *ForVarStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *ForVarStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVar, 0)
}

func (s *ForVarStatementContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *ForVarStatementContext) AllSemiColon() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserSemiColon)
}

func (s *ForVarStatementContext) SemiColon(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSemiColon, i)
}

func (s *ForVarStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *ForVarStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForVarStatementContext) AllExpressionSequence() []IExpressionSequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem())
	var tst = make([]IExpressionSequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionSequenceContext)
		}
	}

	return tst
}

func (s *ForVarStatementContext) ExpressionSequence(i int) IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForVarStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterForVarStatement(s)
	}
}

func (s *ForVarStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitForVarStatement(s)
	}
}

type ForVarInStatementContext struct {
	*IterationStatementContext
}

func NewForVarInStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForVarInStatementContext {
	var p = new(ForVarInStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForVarInStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForVarInStatementContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *ForVarInStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *ForVarInStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVar, 0)
}

func (s *ForVarInStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ForVarInStatementContext) In() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIn, 0)
}

func (s *ForVarInStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForVarInStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *ForVarInStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForVarInStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterForVarInStatement(s)
	}
}

func (s *ForVarInStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitForVarInStatement(s)
	}
}

type WhileStatementContext struct {
	*IterationStatementContext
}

func NewWhileStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWhile, 0)
}

func (s *WhileStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *WhileStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *WhileStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *WhileStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

type ForStatementContext struct {
	*IterationStatementContext
}

func NewForStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStatementContext {
	var p = new(ForStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *ForStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *ForStatementContext) AllSemiColon() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserSemiColon)
}

func (s *ForStatementContext) SemiColon(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSemiColon, i)
}

func (s *ForStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *ForStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) AllExpressionSequence() []IExpressionSequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem())
	var tst = make([]IExpressionSequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionSequenceContext)
		}
	}

	return tst
}

func (s *ForStatementContext) ExpressionSequence(i int) IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitForStatement(s)
	}
}

type ForInStatementContext struct {
	*IterationStatementContext
}

func NewForInStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInStatementContext {
	var p = new(ForInStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForInStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInStatementContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *ForInStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *ForInStatementContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ForInStatementContext) In() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIn, 0)
}

func (s *ForInStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForInStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *ForInStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForInStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterForInStatement(s)
	}
}

func (s *ForInStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitForInStatement(s)
	}
}

func (p *ECMAScriptParser) IterationStatement() (localctx IIterationStatementContext) {
	this := p
	_ = this

	localctx = NewIterationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ECMAScriptParserRULE_iterationStatement)
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

	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDoStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Match(ECMAScriptParserDo)
		}
		{
			p.SetState(188)
			p.Statement()
		}
		{
			p.SetState(189)
			p.Match(ECMAScriptParserWhile)
		}
		{
			p.SetState(190)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(191)
			p.ExpressionSequence()
		}
		{
			p.SetState(192)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(193)
			p.Eos()
		}

	case 2:
		localctx = NewWhileStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Match(ECMAScriptParserWhile)
		}
		{
			p.SetState(196)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(197)
			p.ExpressionSequence()
		}
		{
			p.SetState(198)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(199)
			p.Statement()
		}

	case 3:
		localctx = NewForStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(201)
			p.Match(ECMAScriptParserFor)
		}
		{
			p.SetState(202)
			p.Match(ECMAScriptParserOpenParen)
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(203)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(206)
			p.Match(ECMAScriptParserSemiColon)
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(207)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(210)
			p.Match(ECMAScriptParserSemiColon)
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(211)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(214)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(215)
			p.Statement()
		}

	case 4:
		localctx = NewForVarStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(216)
			p.Match(ECMAScriptParserFor)
		}
		{
			p.SetState(217)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(218)
			p.Match(ECMAScriptParserVar)
		}
		{
			p.SetState(219)
			p.VariableDeclarationList()
		}
		{
			p.SetState(220)
			p.Match(ECMAScriptParserSemiColon)
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(221)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(224)
			p.Match(ECMAScriptParserSemiColon)
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(225)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(228)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(229)
			p.Statement()
		}

	case 5:
		localctx = NewForInStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(231)
			p.Match(ECMAScriptParserFor)
		}
		{
			p.SetState(232)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(233)
			p.singleExpression(0)
		}
		{
			p.SetState(234)
			p.Match(ECMAScriptParserIn)
		}
		{
			p.SetState(235)
			p.ExpressionSequence()
		}
		{
			p.SetState(236)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(237)
			p.Statement()
		}

	case 6:
		localctx = NewForVarInStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(239)
			p.Match(ECMAScriptParserFor)
		}
		{
			p.SetState(240)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(241)
			p.Match(ECMAScriptParserVar)
		}
		{
			p.SetState(242)
			p.VariableDeclaration()
		}
		{
			p.SetState(243)
			p.Match(ECMAScriptParserIn)
		}
		{
			p.SetState(244)
			p.ExpressionSequence()
		}
		{
			p.SetState(245)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(246)
			p.Statement()
		}

	}

	return localctx
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_continueStatement
	return p
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserContinue, 0)
}

func (s *ContinueStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ContinueStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (p *ECMAScriptParser) ContinueStatement() (localctx IContinueStatementContext) {
	this := p
	_ = this

	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ECMAScriptParserRULE_continueStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(ECMAScriptParserContinue)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(251)
			p.Match(ECMAScriptParserIdentifier)
		}

	}
	{
		p.SetState(254)
		p.Eos()
	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBreak, 0)
}

func (s *BreakStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *BreakStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (p *ECMAScriptParser) BreakStatement() (localctx IBreakStatementContext) {
	this := p
	_ = this

	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ECMAScriptParserRULE_breakStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(ECMAScriptParserBreak)
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(257)
			p.Match(ECMAScriptParserIdentifier)
		}

	}
	{
		p.SetState(260)
		p.Eos()
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserReturn, 0)
}

func (s *ReturnStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ReturnStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *ECMAScriptParser) ReturnStatement() (localctx IReturnStatementContext) {
	this := p
	_ = this

	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ECMAScriptParserRULE_returnStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(ECMAScriptParserReturn)
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(263)
			p.ExpressionSequence()
		}

	}
	{
		p.SetState(266)
		p.Eos()
	}

	return localctx
}

// IWithStatementContext is an interface to support dynamic dispatch.
type IWithStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithStatementContext differentiates from other interfaces.
	IsWithStatementContext()
}

type WithStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithStatementContext() *WithStatementContext {
	var p = new(WithStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_withStatement
	return p
}

func (*WithStatementContext) IsWithStatementContext() {}

func NewWithStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithStatementContext {
	var p = new(WithStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_withStatement

	return p
}

func (s *WithStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WithStatementContext) With() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWith, 0)
}

func (s *WithStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *WithStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *WithStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *WithStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WithStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterWithStatement(s)
	}
}

func (s *WithStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitWithStatement(s)
	}
}

func (p *ECMAScriptParser) WithStatement() (localctx IWithStatementContext) {
	this := p
	_ = this

	localctx = NewWithStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ECMAScriptParserRULE_withStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(ECMAScriptParserWith)
	}
	{
		p.SetState(269)
		p.Match(ECMAScriptParserOpenParen)
	}
	{
		p.SetState(270)
		p.ExpressionSequence()
	}
	{
		p.SetState(271)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(272)
		p.Statement()
	}

	return localctx
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_switchStatement
	return p
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) Switch() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSwitch, 0)
}

func (s *SwitchStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *SwitchStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *SwitchStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *SwitchStatementContext) CaseBlock() ICaseBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseBlockContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (p *ECMAScriptParser) SwitchStatement() (localctx ISwitchStatementContext) {
	this := p
	_ = this

	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ECMAScriptParserRULE_switchStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(ECMAScriptParserSwitch)
	}
	{
		p.SetState(275)
		p.Match(ECMAScriptParserOpenParen)
	}
	{
		p.SetState(276)
		p.ExpressionSequence()
	}
	{
		p.SetState(277)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(278)
		p.CaseBlock()
	}

	return localctx
}

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_caseBlock
	return p
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenBrace, 0)
}

func (s *CaseBlockContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseBrace, 0)
}

func (s *CaseBlockContext) AllCaseClauses() []ICaseClausesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseClausesContext)(nil)).Elem())
	var tst = make([]ICaseClausesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseClausesContext)
		}
	}

	return tst
}

func (s *CaseBlockContext) CaseClauses(i int) ICaseClausesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseClausesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseClausesContext)
}

func (s *CaseBlockContext) DefaultClause() IDefaultClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultClauseContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (p *ECMAScriptParser) CaseBlock() (localctx ICaseBlockContext) {
	this := p
	_ = this

	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ECMAScriptParserRULE_caseBlock)
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
		p.SetState(280)
		p.Match(ECMAScriptParserOpenBrace)
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserCase {
		{
			p.SetState(281)
			p.CaseClauses()
		}

	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserDefault {
		{
			p.SetState(284)
			p.DefaultClause()
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ECMAScriptParserCase {
			{
				p.SetState(285)
				p.CaseClauses()
			}

		}

	}
	{
		p.SetState(290)
		p.Match(ECMAScriptParserCloseBrace)
	}

	return localctx
}

// ICaseClausesContext is an interface to support dynamic dispatch.
type ICaseClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseClausesContext differentiates from other interfaces.
	IsCaseClausesContext()
}

type CaseClausesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseClausesContext() *CaseClausesContext {
	var p = new(CaseClausesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_caseClauses
	return p
}

func (*CaseClausesContext) IsCaseClausesContext() {}

func NewCaseClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseClausesContext {
	var p = new(CaseClausesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_caseClauses

	return p
}

func (s *CaseClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseClausesContext) AllCaseClause() []ICaseClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseClauseContext)(nil)).Elem())
	var tst = make([]ICaseClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseClauseContext)
		}
	}

	return tst
}

func (s *CaseClausesContext) CaseClause(i int) ICaseClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseClauseContext)
}

func (s *CaseClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterCaseClauses(s)
	}
}

func (s *CaseClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitCaseClauses(s)
	}
}

func (p *ECMAScriptParser) CaseClauses() (localctx ICaseClausesContext) {
	this := p
	_ = this

	localctx = NewCaseClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ECMAScriptParserRULE_caseClauses)
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
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ECMAScriptParserCase {
		{
			p.SetState(292)
			p.CaseClause()
		}

		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICaseClauseContext is an interface to support dynamic dispatch.
type ICaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseClauseContext differentiates from other interfaces.
	IsCaseClauseContext()
}

type CaseClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseClauseContext() *CaseClauseContext {
	var p = new(CaseClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_caseClause
	return p
}

func (*CaseClauseContext) IsCaseClauseContext() {}

func NewCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseClauseContext {
	var p = new(CaseClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_caseClause

	return p
}

func (s *CaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseClauseContext) Case() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCase, 0)
}

func (s *CaseClauseContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *CaseClauseContext) Colon() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserColon, 0)
}

func (s *CaseClauseContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *CaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterCaseClause(s)
	}
}

func (s *CaseClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitCaseClause(s)
	}
}

func (p *ECMAScriptParser) CaseClause() (localctx ICaseClauseContext) {
	this := p
	_ = this

	localctx = NewCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ECMAScriptParserRULE_caseClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(ECMAScriptParserCase)
	}
	{
		p.SetState(298)
		p.ExpressionSequence()
	}
	{
		p.SetState(299)
		p.Match(ECMAScriptParserColon)
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(300)
			p.StatementList()
		}

	}

	return localctx
}

// IDefaultClauseContext is an interface to support dynamic dispatch.
type IDefaultClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultClauseContext differentiates from other interfaces.
	IsDefaultClauseContext()
}

type DefaultClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultClauseContext() *DefaultClauseContext {
	var p = new(DefaultClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_defaultClause
	return p
}

func (*DefaultClauseContext) IsDefaultClauseContext() {}

func NewDefaultClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultClauseContext {
	var p = new(DefaultClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_defaultClause

	return p
}

func (s *DefaultClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultClauseContext) Default() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDefault, 0)
}

func (s *DefaultClauseContext) Colon() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserColon, 0)
}

func (s *DefaultClauseContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *DefaultClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterDefaultClause(s)
	}
}

func (s *DefaultClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitDefaultClause(s)
	}
}

func (p *ECMAScriptParser) DefaultClause() (localctx IDefaultClauseContext) {
	this := p
	_ = this

	localctx = NewDefaultClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ECMAScriptParserRULE_defaultClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(ECMAScriptParserDefault)
	}
	{
		p.SetState(304)
		p.Match(ECMAScriptParserColon)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(305)
			p.StatementList()
		}

	}

	return localctx
}

// ILabelledStatementContext is an interface to support dynamic dispatch.
type ILabelledStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelledStatementContext differentiates from other interfaces.
	IsLabelledStatementContext()
}

type LabelledStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelledStatementContext() *LabelledStatementContext {
	var p = new(LabelledStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_labelledStatement
	return p
}

func (*LabelledStatementContext) IsLabelledStatementContext() {}

func NewLabelledStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelledStatementContext {
	var p = new(LabelledStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_labelledStatement

	return p
}

func (s *LabelledStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelledStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *LabelledStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserColon, 0)
}

func (s *LabelledStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LabelledStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelledStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelledStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLabelledStatement(s)
	}
}

func (s *LabelledStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLabelledStatement(s)
	}
}

func (p *ECMAScriptParser) LabelledStatement() (localctx ILabelledStatementContext) {
	this := p
	_ = this

	localctx = NewLabelledStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ECMAScriptParserRULE_labelledStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(309)
		p.Match(ECMAScriptParserColon)
	}
	{
		p.SetState(310)
		p.Statement()
	}

	return localctx
}

// IThrowStatementContext is an interface to support dynamic dispatch.
type IThrowStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThrowStatementContext differentiates from other interfaces.
	IsThrowStatementContext()
}

type ThrowStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrowStatementContext() *ThrowStatementContext {
	var p = new(ThrowStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_throwStatement
	return p
}

func (*ThrowStatementContext) IsThrowStatementContext() {}

func NewThrowStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThrowStatementContext {
	var p = new(ThrowStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_throwStatement

	return p
}

func (s *ThrowStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ThrowStatementContext) Throw() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserThrow, 0)
}

func (s *ThrowStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ThrowStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ThrowStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThrowStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThrowStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterThrowStatement(s)
	}
}

func (s *ThrowStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitThrowStatement(s)
	}
}

func (p *ECMAScriptParser) ThrowStatement() (localctx IThrowStatementContext) {
	this := p
	_ = this

	localctx = NewThrowStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ECMAScriptParserRULE_throwStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(ECMAScriptParserThrow)
	}
	{
		p.SetState(313)
		p.ExpressionSequence()
	}
	{
		p.SetState(314)
		p.Eos()
	}

	return localctx
}

// ITryStatementContext is an interface to support dynamic dispatch.
type ITryStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTryStatementContext differentiates from other interfaces.
	IsTryStatementContext()
}

type TryStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTryStatementContext() *TryStatementContext {
	var p = new(TryStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_tryStatement
	return p
}

func (*TryStatementContext) IsTryStatementContext() {}

func NewTryStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryStatementContext {
	var p = new(TryStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_tryStatement

	return p
}

func (s *TryStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TryStatementContext) Try() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserTry, 0)
}

func (s *TryStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TryStatementContext) CatchProduction() ICatchProductionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICatchProductionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICatchProductionContext)
}

func (s *TryStatementContext) FinallyProduction() IFinallyProductionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFinallyProductionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFinallyProductionContext)
}

func (s *TryStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterTryStatement(s)
	}
}

func (s *TryStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitTryStatement(s)
	}
}

func (p *ECMAScriptParser) TryStatement() (localctx ITryStatementContext) {
	this := p
	_ = this

	localctx = NewTryStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ECMAScriptParserRULE_tryStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.Match(ECMAScriptParserTry)
		}
		{
			p.SetState(317)
			p.Block()
		}
		{
			p.SetState(318)
			p.CatchProduction()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.Match(ECMAScriptParserTry)
		}
		{
			p.SetState(321)
			p.Block()
		}
		{
			p.SetState(322)
			p.FinallyProduction()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.Match(ECMAScriptParserTry)
		}
		{
			p.SetState(325)
			p.Block()
		}
		{
			p.SetState(326)
			p.CatchProduction()
		}
		{
			p.SetState(327)
			p.FinallyProduction()
		}

	}

	return localctx
}

// ICatchProductionContext is an interface to support dynamic dispatch.
type ICatchProductionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCatchProductionContext differentiates from other interfaces.
	IsCatchProductionContext()
}

type CatchProductionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatchProductionContext() *CatchProductionContext {
	var p = new(CatchProductionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_catchProduction
	return p
}

func (*CatchProductionContext) IsCatchProductionContext() {}

func NewCatchProductionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatchProductionContext {
	var p = new(CatchProductionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_catchProduction

	return p
}

func (s *CatchProductionContext) GetParser() antlr.Parser { return s.parser }

func (s *CatchProductionContext) Catch() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCatch, 0)
}

func (s *CatchProductionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *CatchProductionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *CatchProductionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *CatchProductionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CatchProductionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatchProductionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatchProductionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterCatchProduction(s)
	}
}

func (s *CatchProductionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitCatchProduction(s)
	}
}

func (p *ECMAScriptParser) CatchProduction() (localctx ICatchProductionContext) {
	this := p
	_ = this

	localctx = NewCatchProductionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ECMAScriptParserRULE_catchProduction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(331)
		p.Match(ECMAScriptParserCatch)
	}
	{
		p.SetState(332)
		p.Match(ECMAScriptParserOpenParen)
	}
	{
		p.SetState(333)
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(334)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(335)
		p.Block()
	}

	return localctx
}

// IFinallyProductionContext is an interface to support dynamic dispatch.
type IFinallyProductionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFinallyProductionContext differentiates from other interfaces.
	IsFinallyProductionContext()
}

type FinallyProductionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFinallyProductionContext() *FinallyProductionContext {
	var p = new(FinallyProductionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_finallyProduction
	return p
}

func (*FinallyProductionContext) IsFinallyProductionContext() {}

func NewFinallyProductionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FinallyProductionContext {
	var p = new(FinallyProductionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_finallyProduction

	return p
}

func (s *FinallyProductionContext) GetParser() antlr.Parser { return s.parser }

func (s *FinallyProductionContext) Finally() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFinally, 0)
}

func (s *FinallyProductionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FinallyProductionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FinallyProductionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FinallyProductionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFinallyProduction(s)
	}
}

func (s *FinallyProductionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFinallyProduction(s)
	}
}

func (p *ECMAScriptParser) FinallyProduction() (localctx IFinallyProductionContext) {
	this := p
	_ = this

	localctx = NewFinallyProductionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ECMAScriptParserRULE_finallyProduction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(ECMAScriptParserFinally)
	}
	{
		p.SetState(338)
		p.Block()
	}

	return localctx
}

// IDebuggerStatementContext is an interface to support dynamic dispatch.
type IDebuggerStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDebuggerStatementContext differentiates from other interfaces.
	IsDebuggerStatementContext()
}

type DebuggerStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDebuggerStatementContext() *DebuggerStatementContext {
	var p = new(DebuggerStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_debuggerStatement
	return p
}

func (*DebuggerStatementContext) IsDebuggerStatementContext() {}

func NewDebuggerStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DebuggerStatementContext {
	var p = new(DebuggerStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_debuggerStatement

	return p
}

func (s *DebuggerStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DebuggerStatementContext) Debugger() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDebugger, 0)
}

func (s *DebuggerStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *DebuggerStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DebuggerStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DebuggerStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterDebuggerStatement(s)
	}
}

func (s *DebuggerStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitDebuggerStatement(s)
	}
}

func (p *ECMAScriptParser) DebuggerStatement() (localctx IDebuggerStatementContext) {
	this := p
	_ = this

	localctx = NewDebuggerStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ECMAScriptParserRULE_debuggerStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(ECMAScriptParserDebugger)
	}
	{
		p.SetState(341)
		p.Eos()
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) Function() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFunction, 0)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *FunctionDeclarationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *FunctionDeclarationContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenBrace, 0)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseBrace, 0)
}

func (s *FunctionDeclarationContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *ECMAScriptParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	this := p
	_ = this

	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ECMAScriptParserRULE_functionDeclaration)
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
		p.SetState(343)
		p.Match(ECMAScriptParserFunction)
	}
	{
		p.SetState(344)
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(345)
		p.Match(ECMAScriptParserOpenParen)
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserIdentifier {
		{
			p.SetState(346)
			p.FormalParameterList()
		}

	}
	{
		p.SetState(349)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(350)
		p.Match(ECMAScriptParserOpenBrace)
	}
	{
		p.SetState(351)
		p.FunctionBody()
	}
	{
		p.SetState(352)
		p.Match(ECMAScriptParserCloseBrace)
	}

	return localctx
}

// IFormalParameterListContext is an interface to support dynamic dispatch.
type IFormalParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterListContext differentiates from other interfaces.
	IsFormalParameterListContext()
}

type FormalParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterListContext() *FormalParameterListContext {
	var p = new(FormalParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_formalParameterList
	return p
}

func (*FormalParameterListContext) IsFormalParameterListContext() {}

func NewFormalParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterListContext {
	var p = new(FormalParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_formalParameterList

	return p
}

func (s *FormalParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserIdentifier)
}

func (s *FormalParameterListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, i)
}

func (s *FormalParameterListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserComma)
}

func (s *FormalParameterListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserComma, i)
}

func (s *FormalParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFormalParameterList(s)
	}
}

func (p *ECMAScriptParser) FormalParameterList() (localctx IFormalParameterListContext) {
	this := p
	_ = this

	localctx = NewFormalParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ECMAScriptParserRULE_formalParameterList)
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
		p.SetState(354)
		p.Match(ECMAScriptParserIdentifier)
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ECMAScriptParserComma {
		{
			p.SetState(355)
			p.Match(ECMAScriptParserComma)
		}
		{
			p.SetState(356)
			p.Match(ECMAScriptParserIdentifier)
		}

		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) SourceElements() ISourceElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceElementsContext)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (p *ECMAScriptParser) FunctionBody() (localctx IFunctionBodyContext) {
	this := p
	_ = this

	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ECMAScriptParserRULE_functionBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(363)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(362)
			p.SourceElements()
		}

	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenBracket, 0)
}

func (s *ArrayLiteralContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseBracket, 0)
}

func (s *ArrayLiteralContext) ElementList() IElementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementListContext)
}

func (s *ArrayLiteralContext) Comma() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserComma, 0)
}

func (s *ArrayLiteralContext) Elision() IElisionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElisionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElisionContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *ECMAScriptParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	this := p
	_ = this

	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ECMAScriptParserRULE_arrayLiteral)
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
		p.SetState(365)
		p.Match(ECMAScriptParserOpenBracket)
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(366)
			p.ElementList()
		}

	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(369)
			p.Match(ECMAScriptParserComma)
		}

	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserComma {
		{
			p.SetState(372)
			p.Elision()
		}

	}
	{
		p.SetState(375)
		p.Match(ECMAScriptParserCloseBracket)
	}

	return localctx
}

// IElementListContext is an interface to support dynamic dispatch.
type IElementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementListContext differentiates from other interfaces.
	IsElementListContext()
}

type ElementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementListContext() *ElementListContext {
	var p = new(ElementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_elementList
	return p
}

func (*ElementListContext) IsElementListContext() {}

func NewElementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementListContext {
	var p = new(ElementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_elementList

	return p
}

func (s *ElementListContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementListContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ElementListContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ElementListContext) AllElision() []IElisionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElisionContext)(nil)).Elem())
	var tst = make([]IElisionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElisionContext)
		}
	}

	return tst
}

func (s *ElementListContext) Elision(i int) IElisionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElisionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElisionContext)
}

func (s *ElementListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserComma)
}

func (s *ElementListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserComma, i)
}

func (s *ElementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterElementList(s)
	}
}

func (s *ElementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitElementList(s)
	}
}

func (p *ECMAScriptParser) ElementList() (localctx IElementListContext) {
	this := p
	_ = this

	localctx = NewElementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ECMAScriptParserRULE_elementList)
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
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserComma {
		{
			p.SetState(377)
			p.Elision()
		}

	}
	{
		p.SetState(380)
		p.singleExpression(0)
	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(381)
				p.Match(ECMAScriptParserComma)
			}
			p.SetState(383)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ECMAScriptParserComma {
				{
					p.SetState(382)
					p.Elision()
				}

			}
			{
				p.SetState(385)
				p.singleExpression(0)
			}

		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}

	return localctx
}

// IElisionContext is an interface to support dynamic dispatch.
type IElisionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElisionContext differentiates from other interfaces.
	IsElisionContext()
}

type ElisionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElisionContext() *ElisionContext {
	var p = new(ElisionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_elision
	return p
}

func (*ElisionContext) IsElisionContext() {}

func NewElisionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElisionContext {
	var p = new(ElisionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_elision

	return p
}

func (s *ElisionContext) GetParser() antlr.Parser { return s.parser }

func (s *ElisionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserComma)
}

func (s *ElisionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserComma, i)
}

func (s *ElisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElisionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterElision(s)
	}
}

func (s *ElisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitElision(s)
	}
}

func (p *ECMAScriptParser) Elision() (localctx IElisionContext) {
	this := p
	_ = this

	localctx = NewElisionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ECMAScriptParserRULE_elision)
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
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ECMAScriptParserComma {
		{
			p.SetState(391)
			p.Match(ECMAScriptParserComma)
		}

		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObjectLiteralContext is an interface to support dynamic dispatch.
type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectLiteralContext differentiates from other interfaces.
	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_objectLiteral
	return p
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenBrace, 0)
}

func (s *ObjectLiteralContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseBrace, 0)
}

func (s *ObjectLiteralContext) PropertyNameAndValueList() IPropertyNameAndValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameAndValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameAndValueListContext)
}

func (s *ObjectLiteralContext) Comma() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserComma, 0)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (p *ECMAScriptParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	this := p
	_ = this

	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ECMAScriptParserRULE_objectLiteral)
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

	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(397)
			p.Match(ECMAScriptParserCloseBrace)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(398)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(399)
			p.PropertyNameAndValueList()
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ECMAScriptParserComma {
			{
				p.SetState(400)
				p.Match(ECMAScriptParserComma)
			}

		}
		{
			p.SetState(403)
			p.Match(ECMAScriptParserCloseBrace)
		}

	}

	return localctx
}

// IPropertyNameAndValueListContext is an interface to support dynamic dispatch.
type IPropertyNameAndValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameAndValueListContext differentiates from other interfaces.
	IsPropertyNameAndValueListContext()
}

type PropertyNameAndValueListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameAndValueListContext() *PropertyNameAndValueListContext {
	var p = new(PropertyNameAndValueListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_propertyNameAndValueList
	return p
}

func (*PropertyNameAndValueListContext) IsPropertyNameAndValueListContext() {}

func NewPropertyNameAndValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameAndValueListContext {
	var p = new(PropertyNameAndValueListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_propertyNameAndValueList

	return p
}

func (s *PropertyNameAndValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameAndValueListContext) AllPropertyAssignment() []IPropertyAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem())
	var tst = make([]IPropertyAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyAssignmentContext)
		}
	}

	return tst
}

func (s *PropertyNameAndValueListContext) PropertyAssignment(i int) IPropertyAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyAssignmentContext)
}

func (s *PropertyNameAndValueListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserComma)
}

func (s *PropertyNameAndValueListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserComma, i)
}

func (s *PropertyNameAndValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameAndValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameAndValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertyNameAndValueList(s)
	}
}

func (s *PropertyNameAndValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertyNameAndValueList(s)
	}
}

func (p *ECMAScriptParser) PropertyNameAndValueList() (localctx IPropertyNameAndValueListContext) {
	this := p
	_ = this

	localctx = NewPropertyNameAndValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ECMAScriptParserRULE_propertyNameAndValueList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(407)
		p.PropertyAssignment()
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(408)
				p.Match(ECMAScriptParserComma)
			}
			{
				p.SetState(409)
				p.PropertyAssignment()
			}

		}
		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// IPropertyAssignmentContext is an interface to support dynamic dispatch.
type IPropertyAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyAssignmentContext differentiates from other interfaces.
	IsPropertyAssignmentContext()
}

type PropertyAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyAssignmentContext() *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_propertyAssignment
	return p
}

func (*PropertyAssignmentContext) IsPropertyAssignmentContext() {}

func NewPropertyAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_propertyAssignment

	return p
}

func (s *PropertyAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyAssignmentContext) CopyFrom(ctx *PropertyAssignmentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PropertyAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropertyExpressionAssignmentContext struct {
	*PropertyAssignmentContext
}

func NewPropertyExpressionAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyExpressionAssignmentContext {
	var p = new(PropertyExpressionAssignmentContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyExpressionAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyExpressionAssignmentContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *PropertyExpressionAssignmentContext) Colon() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserColon, 0)
}

func (s *PropertyExpressionAssignmentContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PropertyExpressionAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertyExpressionAssignment(s)
	}
}

func (s *PropertyExpressionAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertyExpressionAssignment(s)
	}
}

type PropertySetterContext struct {
	*PropertyAssignmentContext
}

func NewPropertySetterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertySetterContext {
	var p = new(PropertySetterContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertySetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertySetterContext) Setter() ISetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetterContext)
}

func (s *PropertySetterContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *PropertySetterContext) PropertySetParameterList() IPropertySetParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertySetParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertySetParameterListContext)
}

func (s *PropertySetterContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *PropertySetterContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenBrace, 0)
}

func (s *PropertySetterContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *PropertySetterContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseBrace, 0)
}

func (s *PropertySetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertySetter(s)
	}
}

func (s *PropertySetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertySetter(s)
	}
}

type PropertyGetterContext struct {
	*PropertyAssignmentContext
}

func NewPropertyGetterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyGetterContext {
	var p = new(PropertyGetterContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyGetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyGetterContext) Getter() IGetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetterContext)
}

func (s *PropertyGetterContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *PropertyGetterContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *PropertyGetterContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenBrace, 0)
}

func (s *PropertyGetterContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *PropertyGetterContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseBrace, 0)
}

func (s *PropertyGetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertyGetter(s)
	}
}

func (s *PropertyGetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertyGetter(s)
	}
}

func (p *ECMAScriptParser) PropertyAssignment() (localctx IPropertyAssignmentContext) {
	this := p
	_ = this

	localctx = NewPropertyAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ECMAScriptParserRULE_propertyAssignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPropertyExpressionAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(415)
			p.PropertyName()
		}
		{
			p.SetState(416)
			p.Match(ECMAScriptParserColon)
		}
		{
			p.SetState(417)
			p.singleExpression(0)
		}

	case 2:
		localctx = NewPropertyGetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(419)
			p.Getter()
		}
		{
			p.SetState(420)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(421)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(422)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(423)
			p.FunctionBody()
		}
		{
			p.SetState(424)
			p.Match(ECMAScriptParserCloseBrace)
		}

	case 3:
		localctx = NewPropertySetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(426)
			p.Setter()
		}
		{
			p.SetState(427)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(428)
			p.PropertySetParameterList()
		}
		{
			p.SetState(429)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(430)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(431)
			p.FunctionBody()
		}
		{
			p.SetState(432)
			p.Match(ECMAScriptParserCloseBrace)
		}

	}

	return localctx
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_propertyName
	return p
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) IdentifierName() IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *PropertyNameContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserStringLiteral, 0)
}

func (s *PropertyNameContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (p *ECMAScriptParser) PropertyName() (localctx IPropertyNameContext) {
	this := p
	_ = this

	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ECMAScriptParserRULE_propertyName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(439)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral, ECMAScriptParserBreak, ECMAScriptParserDo, ECMAScriptParserInstanceof, ECMAScriptParserTypeof, ECMAScriptParserCase, ECMAScriptParserElse, ECMAScriptParserNew, ECMAScriptParserVar, ECMAScriptParserCatch, ECMAScriptParserFinally, ECMAScriptParserReturn, ECMAScriptParserVoid, ECMAScriptParserContinue, ECMAScriptParserFor, ECMAScriptParserSwitch, ECMAScriptParserWhile, ECMAScriptParserDebugger, ECMAScriptParserFunction, ECMAScriptParserThis, ECMAScriptParserWith, ECMAScriptParserDefault, ECMAScriptParserIf, ECMAScriptParserThrow, ECMAScriptParserDelete, ECMAScriptParserIn, ECMAScriptParserTry, ECMAScriptParserClass, ECMAScriptParserEnum, ECMAScriptParserExtends, ECMAScriptParserSuper, ECMAScriptParserConst, ECMAScriptParserExport, ECMAScriptParserImport, ECMAScriptParserImplements, ECMAScriptParserLet, ECMAScriptParserPrivate, ECMAScriptParserPublic, ECMAScriptParserInterface, ECMAScriptParserPackage, ECMAScriptParserProtected, ECMAScriptParserStatic, ECMAScriptParserYield, ECMAScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(436)
			p.IdentifierName()
		}

	case ECMAScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(437)
			p.Match(ECMAScriptParserStringLiteral)
		}

	case ECMAScriptParserDecimalLiteral, ECMAScriptParserHexIntegerLiteral, ECMAScriptParserOctalIntegerLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(438)
			p.NumericLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPropertySetParameterListContext is an interface to support dynamic dispatch.
type IPropertySetParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertySetParameterListContext differentiates from other interfaces.
	IsPropertySetParameterListContext()
}

type PropertySetParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertySetParameterListContext() *PropertySetParameterListContext {
	var p = new(PropertySetParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_propertySetParameterList
	return p
}

func (*PropertySetParameterListContext) IsPropertySetParameterListContext() {}

func NewPropertySetParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertySetParameterListContext {
	var p = new(PropertySetParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_propertySetParameterList

	return p
}

func (s *PropertySetParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertySetParameterListContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *PropertySetParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertySetParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertySetParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertySetParameterList(s)
	}
}

func (s *PropertySetParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertySetParameterList(s)
	}
}

func (p *ECMAScriptParser) PropertySetParameterList() (localctx IPropertySetParameterListContext) {
	this := p
	_ = this

	localctx = NewPropertySetParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ECMAScriptParserRULE_propertySetParameterList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.Match(ECMAScriptParserIdentifier)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *ArgumentsContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *ArgumentsContext) ArgumentList() IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *ECMAScriptParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ECMAScriptParserRULE_arguments)
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
		p.SetState(443)
		p.Match(ECMAScriptParserOpenParen)
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
		{
			p.SetState(444)
			p.ArgumentList()
		}

	}
	{
		p.SetState(447)
		p.Match(ECMAScriptParserCloseParen)
	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_argumentList
	return p
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ArgumentListContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArgumentListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserComma)
}

func (s *ArgumentListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserComma, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (p *ECMAScriptParser) ArgumentList() (localctx IArgumentListContext) {
	this := p
	_ = this

	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ECMAScriptParserRULE_argumentList)
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
		p.SetState(449)
		p.singleExpression(0)
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ECMAScriptParserComma {
		{
			p.SetState(450)
			p.Match(ECMAScriptParserComma)
		}
		{
			p.SetState(451)
			p.singleExpression(0)
		}

		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionSequenceContext is an interface to support dynamic dispatch.
type IExpressionSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionSequenceContext differentiates from other interfaces.
	IsExpressionSequenceContext()
}

type ExpressionSequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSequenceContext() *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_expressionSequence
	return p
}

func (*ExpressionSequenceContext) IsExpressionSequenceContext() {}

func NewExpressionSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_expressionSequence

	return p
}

func (s *ExpressionSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSequenceContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionSequenceContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ExpressionSequenceContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserComma)
}

func (s *ExpressionSequenceContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserComma, i)
}

func (s *ExpressionSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterExpressionSequence(s)
	}
}

func (s *ExpressionSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitExpressionSequence(s)
	}
}

func (p *ECMAScriptParser) ExpressionSequence() (localctx IExpressionSequenceContext) {
	this := p
	_ = this

	localctx = NewExpressionSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ECMAScriptParserRULE_expressionSequence)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(457)
		p.singleExpression(0)
	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(458)
				p.Match(ECMAScriptParserComma)
			}
			{
				p.SetState(459)
				p.singleExpression(0)
			}

		}
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
	}

	return localctx
}

// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) CopyFrom(ctx *SingleExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TernaryExpressionContext struct {
	*SingleExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *TernaryExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserQuestionMark, 0)
}

func (s *TernaryExpressionContext) Colon() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserColon, 0)
}

func (s *TernaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterTernaryExpression(s)
	}
}

func (s *TernaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitTernaryExpression(s)
	}
}

type LogicalAndExpressionContext struct {
	*SingleExpressionContext
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *LogicalAndExpressionContext) And() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserAnd, 0)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

type PreIncrementExpressionContext struct {
	*SingleExpressionContext
}

func NewPreIncrementExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreIncrementExpressionContext {
	var p = new(PreIncrementExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PreIncrementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreIncrementExpressionContext) PlusPlus() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPlusPlus, 0)
}

func (s *PreIncrementExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PreIncrementExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPreIncrementExpression(s)
	}
}

func (s *PreIncrementExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPreIncrementExpression(s)
	}
}

type ObjectLiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewObjectLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectLiteralExpressionContext {
	var p = new(ObjectLiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ObjectLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralExpressionContext) ObjectLiteral() IObjectLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *ObjectLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterObjectLiteralExpression(s)
	}
}

func (s *ObjectLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitObjectLiteralExpression(s)
	}
}

type InExpressionContext struct {
	*SingleExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *InExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *InExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIn, 0)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitInExpression(s)
	}
}

type LogicalOrExpressionContext struct {
	*SingleExpressionContext
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *LogicalOrExpressionContext) Or() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOr, 0)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

type NotExpressionContext struct {
	*SingleExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) Not() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNot, 0)
}

func (s *NotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

type PreDecreaseExpressionContext struct {
	*SingleExpressionContext
}

func NewPreDecreaseExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreDecreaseExpressionContext {
	var p = new(PreDecreaseExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PreDecreaseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreDecreaseExpressionContext) MinusMinus() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserMinusMinus, 0)
}

func (s *PreDecreaseExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PreDecreaseExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPreDecreaseExpression(s)
	}
}

func (s *PreDecreaseExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPreDecreaseExpression(s)
	}
}

type ArgumentsExpressionContext struct {
	*SingleExpressionContext
}

func NewArgumentsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsExpressionContext {
	var p = new(ArgumentsExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ArgumentsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArgumentsExpressionContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *ArgumentsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArgumentsExpression(s)
	}
}

func (s *ArgumentsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArgumentsExpression(s)
	}
}

type ThisExpressionContext struct {
	*SingleExpressionContext
}

func NewThisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThisExpressionContext {
	var p = new(ThisExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ThisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisExpressionContext) This() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserThis, 0)
}

func (s *ThisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterThisExpression(s)
	}
}

func (s *ThisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitThisExpression(s)
	}
}

type FunctionExpressionContext struct {
	*SingleExpressionContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) Function() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFunction, 0)
}

func (s *FunctionExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *FunctionExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *FunctionExpressionContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenBrace, 0)
}

func (s *FunctionExpressionContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionExpressionContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseBrace, 0)
}

func (s *FunctionExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *FunctionExpressionContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

type UnaryMinusExpressionContext struct {
	*SingleExpressionContext
}

func NewUnaryMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExpressionContext {
	var p = new(UnaryMinusExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *UnaryMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserMinus, 0)
}

func (s *UnaryMinusExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *UnaryMinusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterUnaryMinusExpression(s)
	}
}

func (s *UnaryMinusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitUnaryMinusExpression(s)
	}
}

type AssignmentExpressionContext struct {
	*SingleExpressionContext
}

func NewAssignmentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AssignmentExpressionContext) Assign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserAssign, 0)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

type PostDecreaseExpressionContext struct {
	*SingleExpressionContext
}

func NewPostDecreaseExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostDecreaseExpressionContext {
	var p = new(PostDecreaseExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PostDecreaseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostDecreaseExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PostDecreaseExpressionContext) MinusMinus() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserMinusMinus, 0)
}

func (s *PostDecreaseExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPostDecreaseExpression(s)
	}
}

func (s *PostDecreaseExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPostDecreaseExpression(s)
	}
}

type TypeofExpressionContext struct {
	*SingleExpressionContext
}

func NewTypeofExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeofExpressionContext {
	var p = new(TypeofExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *TypeofExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeofExpressionContext) Typeof() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserTypeof, 0)
}

func (s *TypeofExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *TypeofExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterTypeofExpression(s)
	}
}

func (s *TypeofExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitTypeofExpression(s)
	}
}

type InstanceofExpressionContext struct {
	*SingleExpressionContext
}

func NewInstanceofExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstanceofExpressionContext {
	var p = new(InstanceofExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *InstanceofExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceofExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *InstanceofExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *InstanceofExpressionContext) Instanceof() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserInstanceof, 0)
}

func (s *InstanceofExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterInstanceofExpression(s)
	}
}

func (s *InstanceofExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitInstanceofExpression(s)
	}
}

type UnaryPlusExpressionContext struct {
	*SingleExpressionContext
}

func NewUnaryPlusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryPlusExpressionContext {
	var p = new(UnaryPlusExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *UnaryPlusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryPlusExpressionContext) Plus() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPlus, 0)
}

func (s *UnaryPlusExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *UnaryPlusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterUnaryPlusExpression(s)
	}
}

func (s *UnaryPlusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitUnaryPlusExpression(s)
	}
}

type DeleteExpressionContext struct {
	*SingleExpressionContext
}

func NewDeleteExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteExpressionContext {
	var p = new(DeleteExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *DeleteExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteExpressionContext) Delete() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDelete, 0)
}

func (s *DeleteExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *DeleteExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterDeleteExpression(s)
	}
}

func (s *DeleteExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitDeleteExpression(s)
	}
}

type EqualityExpressionContext struct {
	*SingleExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *EqualityExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *EqualityExpressionContext) Equals() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserEquals, 0)
}

func (s *EqualityExpressionContext) NotEquals() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNotEquals, 0)
}

func (s *EqualityExpressionContext) IdentityEquals() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentityEquals, 0)
}

func (s *EqualityExpressionContext) IdentityNotEquals() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentityNotEquals, 0)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

type BitXOrExpressionContext struct {
	*SingleExpressionContext
}

func NewBitXOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXOrExpressionContext {
	var p = new(BitXOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitXOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitXOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitXOrExpressionContext) BitXOr() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBitXOr, 0)
}

func (s *BitXOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitXOrExpression(s)
	}
}

func (s *BitXOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitXOrExpression(s)
	}
}

type MultiplicativeExpressionContext struct {
	*SingleExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MultiplicativeExpressionContext) Multiply() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserMultiply, 0)
}

func (s *MultiplicativeExpressionContext) Divide() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDivide, 0)
}

func (s *MultiplicativeExpressionContext) Modulus() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserModulus, 0)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

type BitShiftExpressionContext struct {
	*SingleExpressionContext
}

func NewBitShiftExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitShiftExpressionContext {
	var p = new(BitShiftExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitShiftExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitShiftExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitShiftExpressionContext) LeftShiftArithmetic() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserLeftShiftArithmetic, 0)
}

func (s *BitShiftExpressionContext) RightShiftArithmetic() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserRightShiftArithmetic, 0)
}

func (s *BitShiftExpressionContext) RightShiftLogical() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserRightShiftLogical, 0)
}

func (s *BitShiftExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitShiftExpression(s)
	}
}

func (s *BitShiftExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitShiftExpression(s)
	}
}

type ParenthesizedExpressionContext struct {
	*SingleExpressionContext
}

func NewParenthesizedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenParen, 0)
}

func (s *ParenthesizedExpressionContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ParenthesizedExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseParen, 0)
}

func (s *ParenthesizedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitParenthesizedExpression(s)
	}
}

type AdditiveExpressionContext struct {
	*SingleExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AdditiveExpressionContext) Plus() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPlus, 0)
}

func (s *AdditiveExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserMinus, 0)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

type RelationalExpressionContext struct {
	*SingleExpressionContext
}

func NewRelationalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *RelationalExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *RelationalExpressionContext) LessThan() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserLessThan, 0)
}

func (s *RelationalExpressionContext) MoreThan() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserMoreThan, 0)
}

func (s *RelationalExpressionContext) LessThanEquals() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserLessThanEquals, 0)
}

func (s *RelationalExpressionContext) GreaterThanEquals() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserGreaterThanEquals, 0)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

type PostIncrementExpressionContext struct {
	*SingleExpressionContext
}

func NewPostIncrementExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostIncrementExpressionContext {
	var p = new(PostIncrementExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PostIncrementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostIncrementExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PostIncrementExpressionContext) PlusPlus() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPlusPlus, 0)
}

func (s *PostIncrementExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPostIncrementExpression(s)
	}
}

func (s *PostIncrementExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPostIncrementExpression(s)
	}
}

type BitNotExpressionContext struct {
	*SingleExpressionContext
}

func NewBitNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitNotExpressionContext {
	var p = new(BitNotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitNotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitNotExpressionContext) BitNot() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBitNot, 0)
}

func (s *BitNotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitNotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitNotExpression(s)
	}
}

func (s *BitNotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitNotExpression(s)
	}
}

type NewExpressionContext struct {
	*SingleExpressionContext
}

func NewNewExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewExpressionContext {
	var p = new(NewExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *NewExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewExpressionContext) New() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNew, 0)
}

func (s *NewExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *NewExpressionContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *NewExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterNewExpression(s)
	}
}

func (s *NewExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitNewExpression(s)
	}
}

type LiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

type ArrayLiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewArrayLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralExpressionContext {
	var p = new(ArrayLiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ArrayLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ArrayLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArrayLiteralExpression(s)
	}
}

func (s *ArrayLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArrayLiteralExpression(s)
	}
}

type MemberDotExpressionContext struct {
	*SingleExpressionContext
}

func NewMemberDotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberDotExpressionContext {
	var p = new(MemberDotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MemberDotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MemberDotExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDot, 0)
}

func (s *MemberDotExpressionContext) IdentifierName() IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *MemberDotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterMemberDotExpression(s)
	}
}

func (s *MemberDotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitMemberDotExpression(s)
	}
}

type MemberIndexExpressionContext struct {
	*SingleExpressionContext
}

func NewMemberIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberIndexExpressionContext {
	var p = new(MemberIndexExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MemberIndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberIndexExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MemberIndexExpressionContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOpenBracket, 0)
}

func (s *MemberIndexExpressionContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *MemberIndexExpressionContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCloseBracket, 0)
}

func (s *MemberIndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterMemberIndexExpression(s)
	}
}

func (s *MemberIndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitMemberIndexExpression(s)
	}
}

type IdentifierExpressionContext struct {
	*SingleExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

type BitAndExpressionContext struct {
	*SingleExpressionContext
}

func NewBitAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndExpressionContext {
	var p = new(BitAndExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitAndExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitAndExpressionContext) BitAnd() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBitAnd, 0)
}

func (s *BitAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitAndExpression(s)
	}
}

func (s *BitAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitAndExpression(s)
	}
}

type BitOrExpressionContext struct {
	*SingleExpressionContext
}

func NewBitOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrExpressionContext {
	var p = new(BitOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitOrExpressionContext) BitOr() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBitOr, 0)
}

func (s *BitOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitOrExpression(s)
	}
}

func (s *BitOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitOrExpression(s)
	}
}

type AssignmentOperatorExpressionContext struct {
	*SingleExpressionContext
}

func NewAssignmentOperatorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentOperatorExpressionContext {
	var p = new(AssignmentOperatorExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AssignmentOperatorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentOperatorExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AssignmentOperatorExpressionContext) AssignmentOperator() IAssignmentOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentOperatorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterAssignmentOperatorExpression(s)
	}
}

func (s *AssignmentOperatorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitAssignmentOperatorExpression(s)
	}
}

type VoidExpressionContext struct {
	*SingleExpressionContext
}

func NewVoidExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidExpressionContext {
	var p = new(VoidExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *VoidExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidExpressionContext) Void() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVoid, 0)
}

func (s *VoidExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *VoidExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVoidExpression(s)
	}
}

func (s *VoidExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVoidExpression(s)
	}
}

func (p *ECMAScriptParser) SingleExpression() (localctx ISingleExpressionContext) {
	return p.singleExpression(0)
}

func (p *ECMAScriptParser) singleExpression(_p int) (localctx ISingleExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISingleExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 86
	p.EnterRecursionRule(localctx, 86, ECMAScriptParserRULE_singleExpression, _p)
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
	p.SetState(511)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserFunction:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(466)
			p.Match(ECMAScriptParserFunction)
		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ECMAScriptParserIdentifier {
			{
				p.SetState(467)
				p.Match(ECMAScriptParserIdentifier)
			}

		}
		{
			p.SetState(470)
			p.Match(ECMAScriptParserOpenParen)
		}
		p.SetState(472)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ECMAScriptParserIdentifier {
			{
				p.SetState(471)
				p.FormalParameterList()
			}

		}
		{
			p.SetState(474)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(475)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(476)
			p.FunctionBody()
		}
		{
			p.SetState(477)
			p.Match(ECMAScriptParserCloseBrace)
		}

	case ECMAScriptParserNew:
		localctx = NewNewExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(479)
			p.Match(ECMAScriptParserNew)
		}
		{
			p.SetState(480)
			p.singleExpression(0)
		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(481)
				p.Arguments()
			}

		}

	case ECMAScriptParserDelete:
		localctx = NewDeleteExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(484)
			p.Match(ECMAScriptParserDelete)
		}
		{
			p.SetState(485)
			p.singleExpression(30)
		}

	case ECMAScriptParserVoid:
		localctx = NewVoidExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(486)
			p.Match(ECMAScriptParserVoid)
		}
		{
			p.SetState(487)
			p.singleExpression(29)
		}

	case ECMAScriptParserTypeof:
		localctx = NewTypeofExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(488)
			p.Match(ECMAScriptParserTypeof)
		}
		{
			p.SetState(489)
			p.singleExpression(28)
		}

	case ECMAScriptParserPlusPlus:
		localctx = NewPreIncrementExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(490)
			p.Match(ECMAScriptParserPlusPlus)
		}
		{
			p.SetState(491)
			p.singleExpression(27)
		}

	case ECMAScriptParserMinusMinus:
		localctx = NewPreDecreaseExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(492)
			p.Match(ECMAScriptParserMinusMinus)
		}
		{
			p.SetState(493)
			p.singleExpression(26)
		}

	case ECMAScriptParserPlus:
		localctx = NewUnaryPlusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(494)
			p.Match(ECMAScriptParserPlus)
		}
		{
			p.SetState(495)
			p.singleExpression(25)
		}

	case ECMAScriptParserMinus:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(496)
			p.Match(ECMAScriptParserMinus)
		}
		{
			p.SetState(497)
			p.singleExpression(24)
		}

	case ECMAScriptParserBitNot:
		localctx = NewBitNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(498)
			p.Match(ECMAScriptParserBitNot)
		}
		{
			p.SetState(499)
			p.singleExpression(23)
		}

	case ECMAScriptParserNot:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(500)
			p.Match(ECMAScriptParserNot)
		}
		{
			p.SetState(501)
			p.singleExpression(22)
		}

	case ECMAScriptParserThis:
		localctx = NewThisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(502)
			p.Match(ECMAScriptParserThis)
		}

	case ECMAScriptParserIdentifier:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(503)
			p.Match(ECMAScriptParserIdentifier)
		}

	case ECMAScriptParserRegularExpressionLiteral, ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral, ECMAScriptParserDecimalLiteral, ECMAScriptParserHexIntegerLiteral, ECMAScriptParserOctalIntegerLiteral, ECMAScriptParserStringLiteral:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(504)
			p.Literal()
		}

	case ECMAScriptParserOpenBracket:
		localctx = NewArrayLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(505)
			p.ArrayLiteral()
		}

	case ECMAScriptParserOpenBrace:
		localctx = NewObjectLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(506)
			p.ObjectLiteral()
		}

	case ECMAScriptParserOpenParen:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(507)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(508)
			p.ExpressionSequence()
		}
		{
			p.SetState(509)
			p.Match(ECMAScriptParserCloseParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(578)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(513)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(514)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserMultiply)|(1<<ECMAScriptParserDivide)|(1<<ECMAScriptParserModulus))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(515)
					p.singleExpression(22)
				}

			case 2:
				localctx = NewAdditiveExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(516)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(517)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ECMAScriptParserPlus || _la == ECMAScriptParserMinus) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(518)
					p.singleExpression(21)
				}

			case 3:
				localctx = NewBitShiftExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(519)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(520)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRightShiftArithmetic)|(1<<ECMAScriptParserLeftShiftArithmetic)|(1<<ECMAScriptParserRightShiftLogical))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(521)
					p.singleExpression(20)
				}

			case 4:
				localctx = NewRelationalExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(522)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(523)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserLessThan)|(1<<ECMAScriptParserMoreThan)|(1<<ECMAScriptParserLessThanEquals)|(1<<ECMAScriptParserGreaterThanEquals))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(524)
					p.singleExpression(19)
				}

			case 5:
				localctx = NewInstanceofExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(525)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(526)
					p.Match(ECMAScriptParserInstanceof)
				}
				{
					p.SetState(527)
					p.singleExpression(18)
				}

			case 6:
				localctx = NewInExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(528)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(529)
					p.Match(ECMAScriptParserIn)
				}
				{
					p.SetState(530)
					p.singleExpression(17)
				}

			case 7:
				localctx = NewEqualityExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(531)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(532)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(ECMAScriptParserEquals-31))|(1<<(ECMAScriptParserNotEquals-31))|(1<<(ECMAScriptParserIdentityEquals-31))|(1<<(ECMAScriptParserIdentityNotEquals-31)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(533)
					p.singleExpression(16)
				}

			case 8:
				localctx = NewBitAndExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(534)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(535)
					p.Match(ECMAScriptParserBitAnd)
				}
				{
					p.SetState(536)
					p.singleExpression(15)
				}

			case 9:
				localctx = NewBitXOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(537)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(538)
					p.Match(ECMAScriptParserBitXOr)
				}
				{
					p.SetState(539)
					p.singleExpression(14)
				}

			case 10:
				localctx = NewBitOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(540)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(541)
					p.Match(ECMAScriptParserBitOr)
				}
				{
					p.SetState(542)
					p.singleExpression(13)
				}

			case 11:
				localctx = NewLogicalAndExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(543)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(544)
					p.Match(ECMAScriptParserAnd)
				}
				{
					p.SetState(545)
					p.singleExpression(12)
				}

			case 12:
				localctx = NewLogicalOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(546)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(547)
					p.Match(ECMAScriptParserOr)
				}
				{
					p.SetState(548)
					p.singleExpression(11)
				}

			case 13:
				localctx = NewTernaryExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(549)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(550)
					p.Match(ECMAScriptParserQuestionMark)
				}
				{
					p.SetState(551)
					p.singleExpression(0)
				}
				{
					p.SetState(552)
					p.Match(ECMAScriptParserColon)
				}
				{
					p.SetState(553)
					p.singleExpression(10)
				}

			case 14:
				localctx = NewAssignmentExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(555)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(556)
					p.Match(ECMAScriptParserAssign)
				}
				{
					p.SetState(557)
					p.singleExpression(9)
				}

			case 15:
				localctx = NewAssignmentOperatorExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(558)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(559)
					p.AssignmentOperator()
				}
				{
					p.SetState(560)
					p.singleExpression(8)
				}

			case 16:
				localctx = NewMemberIndexExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(562)

				if !(p.Precpred(p.GetParserRuleContext(), 36)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 36)", ""))
				}
				{
					p.SetState(563)
					p.Match(ECMAScriptParserOpenBracket)
				}
				{
					p.SetState(564)
					p.ExpressionSequence()
				}
				{
					p.SetState(565)
					p.Match(ECMAScriptParserCloseBracket)
				}

			case 17:
				localctx = NewMemberDotExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(567)

				if !(p.Precpred(p.GetParserRuleContext(), 35)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 35)", ""))
				}
				{
					p.SetState(568)
					p.Match(ECMAScriptParserDot)
				}
				{
					p.SetState(569)
					p.IdentifierName()
				}

			case 18:
				localctx = NewArgumentsExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(570)

				if !(p.Precpred(p.GetParserRuleContext(), 34)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 34)", ""))
				}
				{
					p.SetState(571)
					p.Arguments()
				}

			case 19:
				localctx = NewPostIncrementExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(572)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
				}
				p.SetState(573)

				if !(!p.here(ECMAScriptParserLineTerminator)) {
					panic(antlr.NewFailedPredicateException(p, "!p.here(ECMAScriptParserLineTerminator)", ""))
				}
				{
					p.SetState(574)
					p.Match(ECMAScriptParserPlusPlus)
				}

			case 20:
				localctx = NewPostDecreaseExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(575)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
				}
				p.SetState(576)

				if !(!p.here(ECMAScriptParserLineTerminator)) {
					panic(antlr.NewFailedPredicateException(p, "!p.here(ECMAScriptParserLineTerminator)", ""))
				}
				{
					p.SetState(577)
					p.Match(ECMAScriptParserMinusMinus)
				}

			}

		}
		p.SetState(582)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) MultiplyAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserMultiplyAssign, 0)
}

func (s *AssignmentOperatorContext) DivideAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDivideAssign, 0)
}

func (s *AssignmentOperatorContext) ModulusAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserModulusAssign, 0)
}

func (s *AssignmentOperatorContext) PlusAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPlusAssign, 0)
}

func (s *AssignmentOperatorContext) MinusAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserMinusAssign, 0)
}

func (s *AssignmentOperatorContext) LeftShiftArithmeticAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserLeftShiftArithmeticAssign, 0)
}

func (s *AssignmentOperatorContext) RightShiftArithmeticAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserRightShiftArithmeticAssign, 0)
}

func (s *AssignmentOperatorContext) RightShiftLogicalAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserRightShiftLogicalAssign, 0)
}

func (s *AssignmentOperatorContext) BitAndAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBitAndAssign, 0)
}

func (s *AssignmentOperatorContext) BitXorAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBitXorAssign, 0)
}

func (s *AssignmentOperatorContext) BitOrAssign() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBitOrAssign, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (p *ECMAScriptParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	this := p
	_ = this

	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ECMAScriptParserRULE_assignmentOperator)
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
		p.SetState(583)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ECMAScriptParserMultiplyAssign-40))|(1<<(ECMAScriptParserDivideAssign-40))|(1<<(ECMAScriptParserModulusAssign-40))|(1<<(ECMAScriptParserPlusAssign-40))|(1<<(ECMAScriptParserMinusAssign-40))|(1<<(ECMAScriptParserLeftShiftArithmeticAssign-40))|(1<<(ECMAScriptParserRightShiftArithmeticAssign-40))|(1<<(ECMAScriptParserRightShiftLogicalAssign-40))|(1<<(ECMAScriptParserBitAndAssign-40))|(1<<(ECMAScriptParserBitXorAssign-40))|(1<<(ECMAScriptParserBitOrAssign-40)))) != 0) {
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
	p.RuleIndex = ECMAScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNullLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBooleanLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserStringLiteral, 0)
}

func (s *LiteralContext) RegularExpressionLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserRegularExpressionLiteral, 0)
}

func (s *LiteralContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *ECMAScriptParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ECMAScriptParserRULE_literal)
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

	p.SetState(587)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserRegularExpressionLiteral, ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral, ECMAScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(585)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ECMAScriptParserRegularExpressionLiteral || _la == ECMAScriptParserNullLiteral || _la == ECMAScriptParserBooleanLiteral || _la == ECMAScriptParserStringLiteral) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case ECMAScriptParserDecimalLiteral, ECMAScriptParserHexIntegerLiteral, ECMAScriptParserOctalIntegerLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(586)
			p.NumericLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDecimalLiteral, 0)
}

func (s *NumericLiteralContext) HexIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserHexIntegerLiteral, 0)
}

func (s *NumericLiteralContext) OctalIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOctalIntegerLiteral, 0)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *ECMAScriptParser) NumericLiteral() (localctx INumericLiteralContext) {
	this := p
	_ = this

	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ECMAScriptParserRULE_numericLiteral)
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
		p.SetState(589)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(ECMAScriptParserDecimalLiteral-53))|(1<<(ECMAScriptParserHexIntegerLiteral-53))|(1<<(ECMAScriptParserOctalIntegerLiteral-53)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifierNameContext is an interface to support dynamic dispatch.
type IIdentifierNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierNameContext differentiates from other interfaces.
	IsIdentifierNameContext()
}

type IdentifierNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierNameContext() *IdentifierNameContext {
	var p = new(IdentifierNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_identifierName
	return p
}

func (*IdentifierNameContext) IsIdentifierNameContext() {}

func NewIdentifierNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierNameContext {
	var p = new(IdentifierNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_identifierName

	return p
}

func (s *IdentifierNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *IdentifierNameContext) ReservedWord() IReservedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedWordContext)
}

func (s *IdentifierNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterIdentifierName(s)
	}
}

func (s *IdentifierNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitIdentifierName(s)
	}
}

func (p *ECMAScriptParser) IdentifierName() (localctx IIdentifierNameContext) {
	this := p
	_ = this

	localctx = NewIdentifierNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ECMAScriptParserRULE_identifierName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(593)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(591)
			p.Match(ECMAScriptParserIdentifier)
		}

	case ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral, ECMAScriptParserBreak, ECMAScriptParserDo, ECMAScriptParserInstanceof, ECMAScriptParserTypeof, ECMAScriptParserCase, ECMAScriptParserElse, ECMAScriptParserNew, ECMAScriptParserVar, ECMAScriptParserCatch, ECMAScriptParserFinally, ECMAScriptParserReturn, ECMAScriptParserVoid, ECMAScriptParserContinue, ECMAScriptParserFor, ECMAScriptParserSwitch, ECMAScriptParserWhile, ECMAScriptParserDebugger, ECMAScriptParserFunction, ECMAScriptParserThis, ECMAScriptParserWith, ECMAScriptParserDefault, ECMAScriptParserIf, ECMAScriptParserThrow, ECMAScriptParserDelete, ECMAScriptParserIn, ECMAScriptParserTry, ECMAScriptParserClass, ECMAScriptParserEnum, ECMAScriptParserExtends, ECMAScriptParserSuper, ECMAScriptParserConst, ECMAScriptParserExport, ECMAScriptParserImport, ECMAScriptParserImplements, ECMAScriptParserLet, ECMAScriptParserPrivate, ECMAScriptParserPublic, ECMAScriptParserInterface, ECMAScriptParserPackage, ECMAScriptParserProtected, ECMAScriptParserStatic, ECMAScriptParserYield:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(592)
			p.ReservedWord()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReservedWordContext is an interface to support dynamic dispatch.
type IReservedWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedWordContext differentiates from other interfaces.
	IsReservedWordContext()
}

type ReservedWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedWordContext() *ReservedWordContext {
	var p = new(ReservedWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_reservedWord
	return p
}

func (*ReservedWordContext) IsReservedWordContext() {}

func NewReservedWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedWordContext {
	var p = new(ReservedWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_reservedWord

	return p
}

func (s *ReservedWordContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedWordContext) Keyword() IKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *ReservedWordContext) FutureReservedWord() IFutureReservedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFutureReservedWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFutureReservedWordContext)
}

func (s *ReservedWordContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNullLiteral, 0)
}

func (s *ReservedWordContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBooleanLiteral, 0)
}

func (s *ReservedWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterReservedWord(s)
	}
}

func (s *ReservedWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitReservedWord(s)
	}
}

func (p *ECMAScriptParser) ReservedWord() (localctx IReservedWordContext) {
	this := p
	_ = this

	localctx = NewReservedWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ECMAScriptParserRULE_reservedWord)
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

	p.SetState(598)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserBreak, ECMAScriptParserDo, ECMAScriptParserInstanceof, ECMAScriptParserTypeof, ECMAScriptParserCase, ECMAScriptParserElse, ECMAScriptParserNew, ECMAScriptParserVar, ECMAScriptParserCatch, ECMAScriptParserFinally, ECMAScriptParserReturn, ECMAScriptParserVoid, ECMAScriptParserContinue, ECMAScriptParserFor, ECMAScriptParserSwitch, ECMAScriptParserWhile, ECMAScriptParserDebugger, ECMAScriptParserFunction, ECMAScriptParserThis, ECMAScriptParserWith, ECMAScriptParserDefault, ECMAScriptParserIf, ECMAScriptParserThrow, ECMAScriptParserDelete, ECMAScriptParserIn, ECMAScriptParserTry:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(595)
			p.Keyword()
		}

	case ECMAScriptParserClass, ECMAScriptParserEnum, ECMAScriptParserExtends, ECMAScriptParserSuper, ECMAScriptParserConst, ECMAScriptParserExport, ECMAScriptParserImport, ECMAScriptParserImplements, ECMAScriptParserLet, ECMAScriptParserPrivate, ECMAScriptParserPublic, ECMAScriptParserInterface, ECMAScriptParserPackage, ECMAScriptParserProtected, ECMAScriptParserStatic, ECMAScriptParserYield:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(596)
			p.FutureReservedWord()
		}

	case ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(597)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ECMAScriptParserNullLiteral || _la == ECMAScriptParserBooleanLiteral) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) Break() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBreak, 0)
}

func (s *KeywordContext) Do() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDo, 0)
}

func (s *KeywordContext) Instanceof() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserInstanceof, 0)
}

func (s *KeywordContext) Typeof() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserTypeof, 0)
}

func (s *KeywordContext) Case() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCase, 0)
}

func (s *KeywordContext) Else() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserElse, 0)
}

func (s *KeywordContext) New() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNew, 0)
}

func (s *KeywordContext) Var() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVar, 0)
}

func (s *KeywordContext) Catch() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCatch, 0)
}

func (s *KeywordContext) Finally() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFinally, 0)
}

func (s *KeywordContext) Return() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserReturn, 0)
}

func (s *KeywordContext) Void() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVoid, 0)
}

func (s *KeywordContext) Continue() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserContinue, 0)
}

func (s *KeywordContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *KeywordContext) Switch() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSwitch, 0)
}

func (s *KeywordContext) While() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWhile, 0)
}

func (s *KeywordContext) Debugger() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDebugger, 0)
}

func (s *KeywordContext) Function() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFunction, 0)
}

func (s *KeywordContext) This() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserThis, 0)
}

func (s *KeywordContext) With() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWith, 0)
}

func (s *KeywordContext) Default() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDefault, 0)
}

func (s *KeywordContext) If() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIf, 0)
}

func (s *KeywordContext) Throw() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserThrow, 0)
}

func (s *KeywordContext) Delete() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDelete, 0)
}

func (s *KeywordContext) In() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIn, 0)
}

func (s *KeywordContext) Try() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserTry, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *ECMAScriptParser) Keyword() (localctx IKeywordContext) {
	this := p
	_ = this

	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ECMAScriptParserRULE_keyword)
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
		p.SetState(600)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-56)&-(0x1f+1)) == 0 && ((1<<uint((_la-56)))&((1<<(ECMAScriptParserBreak-56))|(1<<(ECMAScriptParserDo-56))|(1<<(ECMAScriptParserInstanceof-56))|(1<<(ECMAScriptParserTypeof-56))|(1<<(ECMAScriptParserCase-56))|(1<<(ECMAScriptParserElse-56))|(1<<(ECMAScriptParserNew-56))|(1<<(ECMAScriptParserVar-56))|(1<<(ECMAScriptParserCatch-56))|(1<<(ECMAScriptParserFinally-56))|(1<<(ECMAScriptParserReturn-56))|(1<<(ECMAScriptParserVoid-56))|(1<<(ECMAScriptParserContinue-56))|(1<<(ECMAScriptParserFor-56))|(1<<(ECMAScriptParserSwitch-56))|(1<<(ECMAScriptParserWhile-56))|(1<<(ECMAScriptParserDebugger-56))|(1<<(ECMAScriptParserFunction-56))|(1<<(ECMAScriptParserThis-56))|(1<<(ECMAScriptParserWith-56))|(1<<(ECMAScriptParserDefault-56))|(1<<(ECMAScriptParserIf-56))|(1<<(ECMAScriptParserThrow-56))|(1<<(ECMAScriptParserDelete-56))|(1<<(ECMAScriptParserIn-56))|(1<<(ECMAScriptParserTry-56)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFutureReservedWordContext is an interface to support dynamic dispatch.
type IFutureReservedWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFutureReservedWordContext differentiates from other interfaces.
	IsFutureReservedWordContext()
}

type FutureReservedWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFutureReservedWordContext() *FutureReservedWordContext {
	var p = new(FutureReservedWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_futureReservedWord
	return p
}

func (*FutureReservedWordContext) IsFutureReservedWordContext() {}

func NewFutureReservedWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FutureReservedWordContext {
	var p = new(FutureReservedWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_futureReservedWord

	return p
}

func (s *FutureReservedWordContext) GetParser() antlr.Parser { return s.parser }

func (s *FutureReservedWordContext) Class() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserClass, 0)
}

func (s *FutureReservedWordContext) Enum() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserEnum, 0)
}

func (s *FutureReservedWordContext) Extends() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserExtends, 0)
}

func (s *FutureReservedWordContext) Super() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSuper, 0)
}

func (s *FutureReservedWordContext) Const() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserConst, 0)
}

func (s *FutureReservedWordContext) Export() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserExport, 0)
}

func (s *FutureReservedWordContext) Import() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserImport, 0)
}

func (s *FutureReservedWordContext) Implements() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserImplements, 0)
}

func (s *FutureReservedWordContext) Let() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserLet, 0)
}

func (s *FutureReservedWordContext) Private() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPrivate, 0)
}

func (s *FutureReservedWordContext) Public() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPublic, 0)
}

func (s *FutureReservedWordContext) Interface() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserInterface, 0)
}

func (s *FutureReservedWordContext) Package() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPackage, 0)
}

func (s *FutureReservedWordContext) Protected() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserProtected, 0)
}

func (s *FutureReservedWordContext) Static() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserStatic, 0)
}

func (s *FutureReservedWordContext) Yield() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserYield, 0)
}

func (s *FutureReservedWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FutureReservedWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FutureReservedWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFutureReservedWord(s)
	}
}

func (s *FutureReservedWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFutureReservedWord(s)
	}
}

func (p *ECMAScriptParser) FutureReservedWord() (localctx IFutureReservedWordContext) {
	this := p
	_ = this

	localctx = NewFutureReservedWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ECMAScriptParserRULE_futureReservedWord)
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
		p.SetState(602)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-82)&-(0x1f+1)) == 0 && ((1<<uint((_la-82)))&((1<<(ECMAScriptParserClass-82))|(1<<(ECMAScriptParserEnum-82))|(1<<(ECMAScriptParserExtends-82))|(1<<(ECMAScriptParserSuper-82))|(1<<(ECMAScriptParserConst-82))|(1<<(ECMAScriptParserExport-82))|(1<<(ECMAScriptParserImport-82))|(1<<(ECMAScriptParserImplements-82))|(1<<(ECMAScriptParserLet-82))|(1<<(ECMAScriptParserPrivate-82))|(1<<(ECMAScriptParserPublic-82))|(1<<(ECMAScriptParserInterface-82))|(1<<(ECMAScriptParserPackage-82))|(1<<(ECMAScriptParserProtected-82))|(1<<(ECMAScriptParserStatic-82))|(1<<(ECMAScriptParserYield-82)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGetterContext is an interface to support dynamic dispatch.
type IGetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetterContext differentiates from other interfaces.
	IsGetterContext()
}

type GetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetterContext() *GetterContext {
	var p = new(GetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_getter
	return p
}

func (*GetterContext) IsGetterContext() {}

func NewGetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetterContext {
	var p = new(GetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_getter

	return p
}

func (s *GetterContext) GetParser() antlr.Parser { return s.parser }

func (s *GetterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *GetterContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *GetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterGetter(s)
	}
}

func (s *GetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitGetter(s)
	}
}

func (p *ECMAScriptParser) Getter() (localctx IGetterContext) {
	this := p
	_ = this

	localctx = NewGetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ECMAScriptParserRULE_getter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(604)

	if !(p.GetTokenStream().LT(1).GetText() == "get") {
		panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetText() == \"get\"", ""))
	}
	{
		p.SetState(605)
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(606)
		p.PropertyName()
	}

	return localctx
}

// ISetterContext is an interface to support dynamic dispatch.
type ISetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetterContext differentiates from other interfaces.
	IsSetterContext()
}

type SetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetterContext() *SetterContext {
	var p = new(SetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_setter
	return p
}

func (*SetterContext) IsSetterContext() {}

func NewSetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetterContext {
	var p = new(SetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_setter

	return p
}

func (s *SetterContext) GetParser() antlr.Parser { return s.parser }

func (s *SetterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *SetterContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *SetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterSetter(s)
	}
}

func (s *SetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitSetter(s)
	}
}

func (p *ECMAScriptParser) Setter() (localctx ISetterContext) {
	this := p
	_ = this

	localctx = NewSetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, ECMAScriptParserRULE_setter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(608)

	if !(p.GetTokenStream().LT(1).GetText() == "set") {
		panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetText() == \"set\"", ""))
	}
	{
		p.SetState(609)
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(610)
		p.PropertyName()
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSemiColon, 0)
}

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitEos(s)
	}
}

func (p *ECMAScriptParser) Eos() (localctx IEosContext) {
	this := p
	_ = this

	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, ECMAScriptParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(612)
			p.Match(ECMAScriptParserSemiColon)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(613)
			p.Match(ECMAScriptParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(614)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(615)

		if !(p.GetTokenStream().LT(1).GetTokenType() == ECMAScriptParserCloseBrace) {
			panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetTokenType() == ECMAScriptParserCloseBrace", ""))
		}

	}

	return localctx
}

// IEofContext is an interface to support dynamic dispatch.
type IEofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEofContext differentiates from other interfaces.
	IsEofContext()
}

type EofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEofContext() *EofContext {
	var p = new(EofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_eof
	return p
}

func (*EofContext) IsEofContext() {}

func NewEofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EofContext {
	var p = new(EofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_eof

	return p
}

func (s *EofContext) GetParser() antlr.Parser { return s.parser }

func (s *EofContext) EOF() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserEOF, 0)
}

func (s *EofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterEof(s)
	}
}

func (s *EofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitEof(s)
	}
}

func (p *ECMAScriptParser) Eof() (localctx IEofContext) {
	this := p
	_ = this

	localctx = NewEofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, ECMAScriptParserRULE_eof)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(618)
		p.Match(ECMAScriptParserEOF)
	}

	return localctx
}

func (p *ECMAScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *SourceElementContext = nil
		if localctx != nil {
			t = localctx.(*SourceElementContext)
		}
		return p.SourceElement_Sempred(t, predIndex)

	case 3:
		var t *StatementContext = nil
		if localctx != nil {
			t = localctx.(*StatementContext)
		}
		return p.Statement_Sempred(t, predIndex)

	case 43:
		var t *SingleExpressionContext = nil
		if localctx != nil {
			t = localctx.(*SingleExpressionContext)
		}
		return p.SingleExpression_Sempred(t, predIndex)

	case 51:
		var t *GetterContext = nil
		if localctx != nil {
			t = localctx.(*GetterContext)
		}
		return p.Getter_Sempred(t, predIndex)

	case 52:
		var t *SetterContext = nil
		if localctx != nil {
			t = localctx.(*SetterContext)
		}
		return p.Setter_Sempred(t, predIndex)

	case 53:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ECMAScriptParser) SourceElement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.GetInputStream().LA(1) != ECMAScriptParserFunction

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ECMAScriptParser) Statement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.GetInputStream().LA(1) != ECMAScriptParserOpenBrace

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ECMAScriptParser) SingleExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 36)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 35)

	case 19:
		return p.Precpred(p.GetParserRuleContext(), 34)

	case 20:
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 21:
		return !p.here(ECMAScriptParserLineTerminator)

	case 22:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 23:
		return !p.here(ECMAScriptParserLineTerminator)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ECMAScriptParser) Getter_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 24:
		return p.GetTokenStream().LT(1).GetText() == "get"

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ECMAScriptParser) Setter_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 25:
		return p.GetTokenStream().LT(1).GetText() == "set"

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ECMAScriptParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 26:
		return p.lineTerminatorAhead()

	case 27:
		return p.GetTokenStream().LT(1).GetTokenType() == ECMAScriptParserCloseBrace

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
