// Code generated from parkingsign.g4 by ANTLR 4.9.3. DO NOT EDIT.

package file.
import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 55, 472, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 
	3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 
	3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 
	3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 
	3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 
	3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 
	34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 
	38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 
	40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 
	3, 41, 5, 41, 346, 10, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 
	42, 3, 42, 3, 42, 3, 42, 5, 42, 358, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 372, 10, 
	43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 
	3, 44, 5, 44, 385, 10, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 
	45, 3, 45, 3, 45, 5, 45, 396, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 409, 10, 46, 3, 47, 3, 
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 420, 10, 47, 
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 
	49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 
	3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 
	52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 6, 53, 462, 
	10, 53, 13, 53, 14, 53, 463, 3, 54, 6, 54, 467, 10, 54, 13, 54, 14, 54, 
	468, 3, 54, 3, 54, 2, 2, 55, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 
	18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 
	27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 
	36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 
	45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 
	54, 107, 55, 3, 2, 4, 3, 2, 50, 59, 6, 2, 11, 12, 15, 15, 34, 34, 48, 48, 
	2, 480, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 
	2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 
	3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 
	71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 
	2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 
	2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 
	2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 
	3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 
	3, 109, 3, 2, 2, 2, 5, 111, 3, 2, 2, 2, 7, 114, 3, 2, 2, 2, 9, 117, 3, 
	2, 2, 2, 11, 122, 3, 2, 2, 2, 13, 125, 3, 2, 2, 2, 15, 130, 3, 2, 2, 2, 
	17, 133, 3, 2, 2, 2, 19, 141, 3, 2, 2, 2, 21, 144, 3, 2, 2, 2, 23, 149, 
	3, 2, 2, 2, 25, 151, 3, 2, 2, 2, 27, 159, 3, 2, 2, 2, 29, 163, 3, 2, 2, 
	2, 31, 168, 3, 2, 2, 2, 33, 175, 3, 2, 2, 2, 35, 181, 3, 2, 2, 2, 37, 189, 
	3, 2, 2, 2, 39, 196, 3, 2, 2, 2, 41, 201, 3, 2, 2, 2, 43, 210, 3, 2, 2, 
	2, 45, 214, 3, 2, 2, 2, 47, 216, 3, 2, 2, 2, 49, 220, 3, 2, 2, 2, 51, 225, 
	3, 2, 2, 2, 53, 234, 3, 2, 2, 2, 55, 240, 3, 2, 2, 2, 57, 245, 3, 2, 2, 
	2, 59, 254, 3, 2, 2, 2, 61, 259, 3, 2, 2, 2, 63, 268, 3, 2, 2, 2, 65, 276, 
	3, 2, 2, 2, 67, 285, 3, 2, 2, 2, 69, 292, 3, 2, 2, 2, 71, 297, 3, 2, 2, 
	2, 73, 304, 3, 2, 2, 2, 75, 308, 3, 2, 2, 2, 77, 318, 3, 2, 2, 2, 79, 328, 
	3, 2, 2, 2, 81, 345, 3, 2, 2, 2, 83, 357, 3, 2, 2, 2, 85, 371, 3, 2, 2, 
	2, 87, 384, 3, 2, 2, 2, 89, 395, 3, 2, 2, 2, 91, 408, 3, 2, 2, 2, 93, 419, 
	3, 2, 2, 2, 95, 421, 3, 2, 2, 2, 97, 428, 3, 2, 2, 2, 99, 437, 3, 2, 2, 
	2, 101, 446, 3, 2, 2, 2, 103, 451, 3, 2, 2, 2, 105, 461, 3, 2, 2, 2, 107, 
	466, 3, 2, 2, 2, 109, 110, 7, 60, 2, 2, 110, 4, 3, 2, 2, 2, 111, 112, 7, 
	51, 2, 2, 112, 113, 7, 52, 2, 2, 113, 6, 3, 2, 2, 2, 114, 115, 7, 67, 2, 
	2, 115, 116, 7, 79, 2, 2, 116, 8, 3, 2, 2, 2, 117, 118, 7, 67, 2, 2, 118, 
	119, 7, 48, 2, 2, 119, 120, 7, 79, 2, 2, 120, 121, 7, 48, 2, 2, 121, 10, 
	3, 2, 2, 2, 122, 123, 7, 82, 2, 2, 123, 124, 7, 79, 2, 2, 124, 12, 3, 2, 
	2, 2, 125, 126, 7, 82, 2, 2, 126, 127, 7, 48, 2, 2, 127, 128, 7, 79, 2, 
	2, 128, 129, 7, 48, 2, 2, 129, 14, 3, 2, 2, 2, 130, 131, 7, 80, 2, 2, 131, 
	132, 7, 81, 2, 2, 132, 16, 3, 2, 2, 2, 133, 134, 7, 82, 2, 2, 134, 135, 
	7, 67, 2, 2, 135, 136, 7, 84, 2, 2, 136, 137, 7, 77, 2, 2, 137, 138, 7, 
	75, 2, 2, 138, 139, 7, 80, 2, 2, 139, 140, 7, 73, 2, 2, 140, 18, 3, 2, 
	2, 2, 141, 142, 7, 86, 2, 2, 142, 143, 7, 81, 2, 2, 143, 20, 3, 2, 2, 2, 
	144, 145, 7, 86, 2, 2, 145, 146, 7, 74, 2, 2, 146, 147, 7, 84, 2, 2, 147, 
	148, 7, 87, 2, 2, 148, 22, 3, 2, 2, 2, 149, 150, 7, 47, 2, 2, 150, 24, 
	3, 2, 2, 2, 151, 152, 7, 67, 2, 2, 152, 153, 7, 80, 2, 2, 153, 154, 7, 
	91, 2, 2, 154, 155, 7, 86, 2, 2, 155, 156, 7, 75, 2, 2, 156, 157, 7, 79, 
	2, 2, 157, 158, 7, 71, 2, 2, 158, 26, 3, 2, 2, 2, 159, 160, 7, 67, 2, 2, 
	160, 161, 7, 80, 2, 2, 161, 162, 7, 91, 2, 2, 162, 28, 3, 2, 2, 2, 163, 
	164, 7, 86, 2, 2, 164, 165, 7, 75, 2, 2, 165, 166, 7, 79, 2, 2, 166, 167, 
	7, 71, 2, 2, 167, 30, 3, 2, 2, 2, 168, 169, 7, 71, 2, 2, 169, 170, 7, 90, 
	2, 2, 170, 171, 7, 69, 2, 2, 171, 172, 7, 71, 2, 2, 172, 173, 7, 82, 2, 
	2, 173, 174, 7, 86, 2, 2, 174, 32, 3, 2, 2, 2, 175, 176, 7, 70, 2, 2, 176, 
	177, 7, 67, 2, 2, 177, 178, 7, 75, 2, 2, 178, 179, 7, 78, 2, 2, 179, 180, 
	7, 91, 2, 2, 180, 34, 3, 2, 2, 2, 181, 182, 7, 80, 2, 2, 182, 183, 7, 75, 
	2, 2, 183, 184, 7, 73, 2, 2, 184, 185, 7, 74, 2, 2, 185, 186, 7, 86, 2, 
	2, 186, 187, 7, 78, 2, 2, 187, 188, 7, 91, 2, 2, 188, 36, 3, 2, 2, 2, 189, 
	190, 7, 85, 2, 2, 190, 191, 7, 69, 2, 2, 191, 192, 7, 74, 2, 2, 192, 193, 
	7, 81, 2, 2, 193, 194, 7, 81, 2, 2, 194, 195, 7, 78, 2, 2, 195, 38, 3, 
	2, 2, 2, 196, 197, 7, 70, 2, 2, 197, 198, 7, 67, 2, 2, 198, 199, 7, 91, 
	2, 2, 199, 200, 7, 85, 2, 2, 200, 40, 3, 2, 2, 2, 201, 202, 7, 74, 2, 2, 
	202, 203, 7, 81, 2, 2, 203, 204, 7, 78, 2, 2, 204, 205, 7, 75, 2, 2, 205, 
	206, 7, 70, 2, 2, 206, 207, 7, 67, 2, 2, 207, 208, 7, 91, 2, 2, 208, 209, 
	7, 85, 2, 2, 209, 42, 3, 2, 2, 2, 210, 211, 7, 67, 2, 2, 211, 212, 7, 80, 
	2, 2, 212, 213, 7, 70, 2, 2, 213, 44, 3, 2, 2, 2, 214, 215, 7, 40, 2, 2, 
	215, 46, 3, 2, 2, 2, 216, 217, 7, 86, 2, 2, 217, 218, 7, 81, 2, 2, 218, 
	219, 7, 89, 2, 2, 219, 48, 3, 2, 2, 2, 220, 221, 7, 67, 2, 2, 221, 222, 
	7, 89, 2, 2, 222, 223, 7, 67, 2, 2, 223, 224, 7, 91, 2, 2, 224, 50, 3, 
	2, 2, 2, 225, 226, 7, 85, 2, 2, 226, 227, 7, 86, 2, 2, 227, 228, 7, 81, 
	2, 2, 228, 229, 7, 82, 2, 2, 229, 230, 7, 82, 2, 2, 230, 231, 7, 75, 2, 
	2, 231, 232, 7, 80, 2, 2, 232, 233, 7, 73, 2, 2, 233, 52, 3, 2, 2, 2, 234, 
	235, 7, 88, 2, 2, 235, 236, 7, 67, 2, 2, 236, 237, 7, 78, 2, 2, 237, 238, 
	7, 71, 2, 2, 238, 239, 7, 86, 2, 2, 239, 54, 3, 2, 2, 2, 240, 241, 7, 81, 
	2, 2, 241, 242, 7, 80, 2, 2, 242, 243, 7, 78, 2, 2, 243, 244, 7, 91, 2, 
	2, 244, 56, 3, 2, 2, 2, 245, 246, 7, 88, 2, 2, 246, 247, 7, 71, 2, 2, 247, 
	248, 7, 74, 2, 2, 248, 249, 7, 75, 2, 2, 249, 250, 7, 69, 2, 2, 250, 251, 
	7, 78, 2, 2, 251, 252, 7, 71, 2, 2, 252, 253, 7, 85, 2, 2, 253, 58, 3, 
	2, 2, 2, 254, 255, 7, 89, 2, 2, 255, 256, 7, 75, 2, 2, 256, 257, 7, 86, 
	2, 2, 257, 258, 7, 74, 2, 2, 258, 60, 3, 2, 2, 2, 259, 260, 7, 70, 2, 2, 
	260, 261, 7, 75, 2, 2, 261, 262, 7, 85, 2, 2, 262, 263, 7, 86, 2, 2, 263, 
	264, 7, 84, 2, 2, 264, 265, 7, 75, 2, 2, 265, 266, 7, 69, 2, 2, 266, 267, 
	7, 86, 2, 2, 267, 62, 3, 2, 2, 2, 268, 269, 7, 82, 2, 2, 269, 270, 7, 71, 
	2, 2, 270, 271, 7, 84, 2, 2, 271, 272, 7, 79, 2, 2, 272, 273, 7, 75, 2, 
	2, 273, 274, 7, 86, 2, 2, 274, 275, 7, 85, 2, 2, 275, 64, 3, 2, 2, 2, 276, 
	277, 7, 71, 2, 2, 277, 278, 7, 90, 2, 2, 278, 279, 7, 71, 2, 2, 279, 280, 
	7, 79, 2, 2, 280, 281, 7, 82, 2, 2, 281, 282, 7, 86, 2, 2, 282, 283, 7, 
	71, 2, 2, 283, 284, 7, 70, 2, 2, 284, 66, 3, 2, 2, 2, 285, 286, 7, 71, 
	2, 2, 286, 287, 7, 90, 2, 2, 287, 288, 7, 71, 2, 2, 288, 289, 7, 79, 2, 
	2, 289, 290, 7, 82, 2, 2, 290, 291, 7, 86, 2, 2, 291, 68, 3, 2, 2, 2, 292, 
	293, 7, 74, 2, 2, 293, 294, 7, 81, 2, 2, 294, 295, 7, 87, 2, 2, 295, 296, 
	7, 84, 2, 2, 296, 70, 3, 2, 2, 2, 297, 298, 7, 79, 2, 2, 298, 299, 7, 75, 
	2, 2, 299, 300, 7, 80, 2, 2, 300, 301, 7, 87, 2, 2, 301, 302, 7, 86, 2, 
	2, 302, 303, 7, 71, 2, 2, 303, 72, 3, 2, 2, 2, 304, 305, 7, 79, 2, 2, 305, 
	306, 7, 75, 2, 2, 306, 307, 7, 80, 2, 2, 307, 74, 3, 2, 2, 2, 308, 309, 
	7, 86, 2, 2, 309, 310, 7, 71, 2, 2, 310, 311, 7, 79, 2, 2, 311, 312, 7, 
	82, 2, 2, 312, 313, 7, 81, 2, 2, 313, 314, 7, 84, 2, 2, 314, 315, 7, 67, 
	2, 2, 315, 316, 7, 84, 2, 2, 316, 317, 7, 91, 2, 2, 317, 76, 3, 2, 2, 2, 
	318, 319, 7, 82, 2, 2, 319, 320, 7, 67, 2, 2, 320, 321, 7, 85, 2, 2, 321, 
	322, 7, 85, 2, 2, 322, 323, 7, 71, 2, 2, 323, 324, 7, 80, 2, 2, 324, 325, 
	7, 73, 2, 2, 325, 326, 7, 71, 2, 2, 326, 327, 7, 84, 2, 2, 327, 78, 3, 
	2, 2, 2, 328, 329, 7, 78, 2, 2, 329, 330, 7, 81, 2, 2, 330, 331, 7, 67, 
	2, 2, 331, 332, 7, 70, 2, 2, 332, 333, 7, 75, 2, 2, 333, 334, 7, 80, 2, 
	2, 334, 335, 7, 73, 2, 2, 335, 80, 3, 2, 2, 2, 336, 337, 7, 79, 2, 2, 337, 
	338, 7, 81, 2, 2, 338, 339, 7, 80, 2, 2, 339, 340, 7, 70, 2, 2, 340, 341, 
	7, 67, 2, 2, 341, 346, 7, 91, 2, 2, 342, 343, 7, 79, 2, 2, 343, 344, 7, 
	81, 2, 2, 344, 346, 7, 80, 2, 2, 345, 336, 3, 2, 2, 2, 345, 342, 3, 2, 
	2, 2, 346, 82, 3, 2, 2, 2, 347, 348, 7, 86, 2, 2, 348, 349, 7, 87, 2, 2, 
	349, 350, 7, 71, 2, 2, 350, 351, 7, 85, 2, 2, 351, 352, 7, 70, 2, 2, 352, 
	353, 7, 67, 2, 2, 353, 358, 7, 91, 2, 2, 354, 355, 7, 86, 2, 2, 355, 356, 
	7, 87, 2, 2, 356, 358, 7, 71, 2, 2, 357, 347, 3, 2, 2, 2, 357, 354, 3, 
	2, 2, 2, 358, 84, 3, 2, 2, 2, 359, 360, 7, 89, 2, 2, 360, 361, 7, 71, 2, 
	2, 361, 362, 7, 70, 2, 2, 362, 363, 7, 80, 2, 2, 363, 364, 7, 71, 2, 2, 
	364, 365, 7, 85, 2, 2, 365, 366, 7, 70, 2, 2, 366, 367, 7, 67, 2, 2, 367, 
	372, 7, 91, 2, 2, 368, 369, 7, 89, 2, 2, 369, 370, 7, 71, 2, 2, 370, 372, 
	7, 70, 2, 2, 371, 359, 3, 2, 2, 2, 371, 368, 3, 2, 2, 2, 372, 86, 3, 2, 
	2, 2, 373, 374, 7, 86, 2, 2, 374, 375, 7, 74, 2, 2, 375, 376, 7, 87, 2, 
	2, 376, 377, 7, 84, 2, 2, 377, 378, 7, 85, 2, 2, 378, 379, 7, 70, 2, 2, 
	379, 380, 7, 67, 2, 2, 380, 385, 7, 91, 2, 2, 381, 382, 7, 86, 2, 2, 382, 
	383, 7, 74, 2, 2, 383, 385, 7, 87, 2, 2, 384, 373, 3, 2, 2, 2, 384, 381, 
	3, 2, 2, 2, 385, 88, 3, 2, 2, 2, 386, 387, 7, 72, 2, 2, 387, 388, 7, 84, 
	2, 2, 388, 389, 7, 75, 2, 2, 389, 390, 7, 70, 2, 2, 390, 391, 7, 67, 2, 
	2, 391, 396, 7, 91, 2, 2, 392, 393, 7, 72, 2, 2, 393, 394, 7, 84, 2, 2, 
	394, 396, 7, 75, 2, 2, 395, 386, 3, 2, 2, 2, 395, 392, 3, 2, 2, 2, 396, 
	90, 3, 2, 2, 2, 397, 398, 7, 85, 2, 2, 398, 399, 7, 67, 2, 2, 399, 400, 
	7, 86, 2, 2, 400, 401, 7, 87, 2, 2, 401, 402, 7, 84, 2, 2, 402, 403, 7, 
	70, 2, 2, 403, 404, 7, 67, 2, 2, 404, 409, 7, 91, 2, 2, 405, 406, 7, 85, 
	2, 2, 406, 407, 7, 67, 2, 2, 407, 409, 7, 86, 2, 2, 408, 397, 3, 2, 2, 
	2, 408, 405, 3, 2, 2, 2, 409, 92, 3, 2, 2, 2, 410, 411, 7, 85, 2, 2, 411, 
	412, 7, 87, 2, 2, 412, 413, 7, 80, 2, 2, 413, 414, 7, 70, 2, 2, 414, 415, 
	7, 67, 2, 2, 415, 420, 7, 91, 2, 2, 416, 417, 7, 85, 2, 2, 417, 418, 7, 
	87, 2, 2, 418, 420, 7, 80, 2, 2, 419, 410, 3, 2, 2, 2, 419, 416, 3, 2, 
	2, 2, 420, 94, 3, 2, 2, 2, 421, 422, 7, 85, 2, 2, 422, 423, 7, 86, 2, 2, 
	423, 424, 7, 84, 2, 2, 424, 425, 7, 71, 2, 2, 425, 426, 7, 71, 2, 2, 426, 
	427, 7, 86, 2, 2, 427, 96, 3, 2, 2, 2, 428, 429, 7, 85, 2, 2, 429, 430, 
	7, 89, 2, 2, 430, 431, 7, 71, 2, 2, 431, 432, 7, 71, 2, 2, 432, 433, 7, 
	82, 2, 2, 433, 434, 7, 75, 2, 2, 434, 435, 7, 80, 2, 2, 435, 436, 7, 73, 
	2, 2, 436, 98, 3, 2, 2, 2, 437, 438, 7, 69, 2, 2, 438, 439, 7, 78, 2, 2, 
	439, 440, 7, 71, 2, 2, 440, 441, 7, 67, 2, 2, 441, 442, 7, 80, 2, 2, 442, 
	443, 7, 75, 2, 2, 443, 444, 7, 80, 2, 2, 444, 445, 7, 73, 2, 2, 445, 100, 
	3, 2, 2, 2, 446, 447, 7, 80, 2, 2, 447, 448, 7, 81, 2, 2, 448, 449, 7, 
	81, 2, 2, 449, 450, 7, 80, 2, 2, 450, 102, 3, 2, 2, 2, 451, 452, 7, 79, 
	2, 2, 452, 453, 7, 75, 2, 2, 453, 454, 7, 70, 2, 2, 454, 455, 7, 80, 2, 
	2, 455, 456, 7, 75, 2, 2, 456, 457, 7, 73, 2, 2, 457, 458, 7, 74, 2, 2, 
	458, 459, 7, 86, 2, 2, 459, 104, 3, 2, 2, 2, 460, 462, 9, 2, 2, 2, 461, 
	460, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463, 461, 3, 2, 2, 2, 463, 464, 
	3, 2, 2, 2, 464, 106, 3, 2, 2, 2, 465, 467, 9, 3, 2, 2, 466, 465, 3, 2, 
	2, 2, 467, 468, 3, 2, 2, 2, 468, 466, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 
	469, 470, 3, 2, 2, 2, 470, 471, 8, 54, 2, 2, 471, 108, 3, 2, 2, 2, 12, 
	2, 345, 357, 371, 384, 395, 408, 419, 463, 468, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':'", "'12'", "'AM'", "'A.M.'", "'PM'", "'P.M.'", "'NO'", "'PARKING'", 
	"'TO'", "'THRU'", "'-'", "'ANYTIME'", "'ANY'", "'TIME'", "'EXCEPT'", "'DAILY'", 
	"'NIGHTLY'", "'SCHOOL'", "'DAYS'", "'HOLIDAYS'", "'AND'", "'&'", "'TOW'", 
	"'AWAY'", "'STOPPING'", "'VALET'", "'ONLY'", "'VEHICLES'", "'WITH'", "'DISTRICT'", 
	"'PERMITS'", "'EXEMPTED'", "'EXEMPT'", "'HOUR'", "'MINUTE'", "'MIN'", "'TEMPORARY'", 
	"'PASSENGER'", "'LOADING'", "", "", "", "", "", "", "", "'STREET'", "'SWEEPING'", 
	"'CLEANING'", "'NOON'", "'MIDNIGHT'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "NO", "PARKING", "TO", "THRU", "DASH", "ANYTIME", 
	"ANY", "TIME", "EXCEPT", "DAILY", "NIGHTLY", "SCHOOL", "DAYS", "HOLIDAYS", 
	"AND", "AMPERSAND", "TOW", "AWAY", "STOPPING", "VALET", "ONLY", "VEHICLES", 
	"WITH", "DISTRICT", "PERMITS", "EXEMPTED", "EXEMPT", "HOUR", "MINUTE", 
	"MIN", "TEMPORARY", "PASSENGER", "LOADING", "MON", "TUE", "WED", "THU", 
	"FRI", "SAT", "SUN", "STREET", "SWEEPING", "CLEANING", "NOON", "MIDNIGHT", 
	"INT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "NO", "PARKING", "TO", 
	"THRU", "DASH", "ANYTIME", "ANY", "TIME", "EXCEPT", "DAILY", "NIGHTLY", 
	"SCHOOL", "DAYS", "HOLIDAYS", "AND", "AMPERSAND", "TOW", "AWAY", "STOPPING", 
	"VALET", "ONLY", "VEHICLES", "WITH", "DISTRICT", "PERMITS", "EXEMPTED", 
	"EXEMPT", "HOUR", "MINUTE", "MIN", "TEMPORARY", "PASSENGER", "LOADING", 
	"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN", "STREET", "SWEEPING", 
	"CLEANING", "NOON", "MIDNIGHT", "INT", "WS",
}
type parkingsignLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

// NewparkingsignLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *parkingsignLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewparkingsignLexer(input antlr.CharStream) *parkingsignLexer {
	l := new(parkingsignLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "parkingsign.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// parkingsignLexer tokens.
const (
	parkingsignLexerT__0 = 1
	parkingsignLexerT__1 = 2
	parkingsignLexerT__2 = 3
	parkingsignLexerT__3 = 4
	parkingsignLexerT__4 = 5
	parkingsignLexerT__5 = 6
	parkingsignLexerNO = 7
	parkingsignLexerPARKING = 8
	parkingsignLexerTO = 9
	parkingsignLexerTHRU = 10
	parkingsignLexerDASH = 11
	parkingsignLexerANYTIME = 12
	parkingsignLexerANY = 13
	parkingsignLexerTIME = 14
	parkingsignLexerEXCEPT = 15
	parkingsignLexerDAILY = 16
	parkingsignLexerNIGHTLY = 17
	parkingsignLexerSCHOOL = 18
	parkingsignLexerDAYS = 19
	parkingsignLexerHOLIDAYS = 20
	parkingsignLexerAND = 21
	parkingsignLexerAMPERSAND = 22
	parkingsignLexerTOW = 23
	parkingsignLexerAWAY = 24
	parkingsignLexerSTOPPING = 25
	parkingsignLexerVALET = 26
	parkingsignLexerONLY = 27
	parkingsignLexerVEHICLES = 28
	parkingsignLexerWITH = 29
	parkingsignLexerDISTRICT = 30
	parkingsignLexerPERMITS = 31
	parkingsignLexerEXEMPTED = 32
	parkingsignLexerEXEMPT = 33
	parkingsignLexerHOUR = 34
	parkingsignLexerMINUTE = 35
	parkingsignLexerMIN = 36
	parkingsignLexerTEMPORARY = 37
	parkingsignLexerPASSENGER = 38
	parkingsignLexerLOADING = 39
	parkingsignLexerMON = 40
	parkingsignLexerTUE = 41
	parkingsignLexerWED = 42
	parkingsignLexerTHU = 43
	parkingsignLexerFRI = 44
	parkingsignLexerSAT = 45
	parkingsignLexerSUN = 46
	parkingsignLexerSTREET = 47
	parkingsignLexerSWEEPING = 48
	parkingsignLexerCLEANING = 49
	parkingsignLexerNOON = 50
	parkingsignLexerMIDNIGHT = 51
	parkingsignLexerINT = 52
	parkingsignLexerWS = 53
)

