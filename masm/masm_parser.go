// Generated from MASM.g4 by ANTLR 4.7.

package masm // MASM
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)


 
 	 package com.Ostermiller.Syntax;
 	 

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 276, 377, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9, 
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44, 
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4, 
	50, 9, 50, 3, 2, 3, 2, 7, 2, 103, 10, 2, 12, 2, 14, 2, 106, 11, 2, 3, 2, 
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 117, 10, 3, 12, 3, 
	14, 3, 120, 11, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 128, 10, 4, 
	12, 4, 14, 4, 131, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 157, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 
	6, 3, 6, 5, 6, 165, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 172, 10, 
	6, 5, 6, 174, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 180, 10, 7, 3, 8, 3, 
	8, 3, 8, 5, 8, 185, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 192, 10, 
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 199, 10, 9, 3, 10, 3, 10, 3, 11, 
	3, 11, 3, 11, 5, 11, 206, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 212, 
	10, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 
	222, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 5, 
	15, 232, 10, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 5, 16, 240, 
	10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 245, 10, 16, 3, 17, 3, 17, 3, 17, 5, 
	17, 250, 10, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 257, 10, 17, 
	3, 18, 3, 18, 3, 18, 5, 18, 262, 10, 18, 3, 18, 3, 18, 3, 18, 5, 18, 267, 
	10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 273, 10, 19, 3, 20, 3, 20, 3, 
	20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 283, 10, 21, 3, 22, 3, 22, 
	3, 22, 5, 22, 288, 10, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 
	24, 3, 24, 3, 24, 3, 24, 5, 24, 300, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 
	3, 25, 5, 25, 307, 10, 25, 3, 26, 3, 26, 3, 26, 5, 26, 312, 10, 26, 3, 
	26, 3, 26, 3, 26, 3, 26, 5, 26, 318, 10, 26, 3, 26, 3, 26, 3, 26, 5, 26, 
	323, 10, 26, 5, 26, 325, 10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 
	28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 
	3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 
	39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 
	3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 
	49, 3, 50, 3, 50, 3, 50, 2, 2, 51, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 
	58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 
	94, 96, 98, 2, 20, 3, 2, 268, 270, 3, 2, 13, 18, 3, 2, 19, 42, 3, 2, 43, 
	45, 3, 2, 47, 85, 3, 2, 87, 140, 3, 2, 141, 146, 3, 2, 147, 183, 3, 2, 
	184, 186, 3, 2, 187, 191, 3, 2, 192, 193, 3, 2, 194, 201, 3, 2, 202, 203, 
	3, 2, 204, 206, 3, 2, 208, 211, 3, 2, 214, 218, 3, 2, 219, 244, 3, 2, 245, 
	265, 2, 384, 2, 104, 3, 2, 2, 2, 4, 110, 3, 2, 2, 2, 6, 124, 3, 2, 2, 2, 
	8, 156, 3, 2, 2, 2, 10, 173, 3, 2, 2, 2, 12, 175, 3, 2, 2, 2, 14, 181, 
	3, 2, 2, 2, 16, 198, 3, 2, 2, 2, 18, 200, 3, 2, 2, 2, 20, 202, 3, 2, 2, 
	2, 22, 213, 3, 2, 2, 2, 24, 216, 3, 2, 2, 2, 26, 223, 3, 2, 2, 2, 28, 228, 
	3, 2, 2, 2, 30, 236, 3, 2, 2, 2, 32, 246, 3, 2, 2, 2, 34, 258, 3, 2, 2, 
	2, 36, 268, 3, 2, 2, 2, 38, 274, 3, 2, 2, 2, 40, 277, 3, 2, 2, 2, 42, 284, 
	3, 2, 2, 2, 44, 292, 3, 2, 2, 2, 46, 299, 3, 2, 2, 2, 48, 301, 3, 2, 2, 
	2, 50, 308, 3, 2, 2, 2, 52, 328, 3, 2, 2, 2, 54, 330, 3, 2, 2, 2, 56, 332, 
	3, 2, 2, 2, 58, 334, 3, 2, 2, 2, 60, 336, 3, 2, 2, 2, 62, 338, 3, 2, 2, 
	2, 64, 340, 3, 2, 2, 2, 66, 342, 3, 2, 2, 2, 68, 344, 3, 2, 2, 2, 70, 346, 
	3, 2, 2, 2, 72, 348, 3, 2, 2, 2, 74, 350, 3, 2, 2, 2, 76, 352, 3, 2, 2, 
	2, 78, 354, 3, 2, 2, 2, 80, 356, 3, 2, 2, 2, 82, 358, 3, 2, 2, 2, 84, 360, 
	3, 2, 2, 2, 86, 362, 3, 2, 2, 2, 88, 364, 3, 2, 2, 2, 90, 366, 3, 2, 2, 
	2, 92, 368, 3, 2, 2, 2, 94, 370, 3, 2, 2, 2, 96, 372, 3, 2, 2, 2, 98, 374, 
	3, 2, 2, 2, 100, 103, 5, 4, 3, 2, 101, 103, 5, 46, 24, 2, 102, 100, 3, 
	2, 2, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2, 2, 
	2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 
	108, 7, 3, 2, 2, 108, 109, 7, 12, 2, 2, 109, 3, 3, 2, 2, 2, 110, 111, 7, 
	12, 2, 2, 111, 112, 7, 4, 2, 2, 112, 113, 7, 5, 2, 2, 113, 118, 7, 6, 2, 
	2, 114, 117, 5, 8, 5, 2, 115, 117, 5, 6, 4, 2, 116, 114, 3, 2, 2, 2, 116, 
	115, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 
	3, 2, 2, 2, 119, 121, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 7, 12, 
	2, 2, 122, 123, 7, 7, 2, 2, 123, 5, 3, 2, 2, 2, 124, 125, 7, 12, 2, 2, 
	125, 129, 7, 265, 2, 2, 126, 128, 5, 8, 5, 2, 127, 126, 3, 2, 2, 2, 128, 
	131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 132, 
	3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 133, 7, 210, 2, 2, 133, 134, 7, 
	12, 2, 2, 134, 135, 7, 8, 2, 2, 135, 7, 3, 2, 2, 2, 136, 157, 5, 10, 6, 
	2, 137, 157, 5, 40, 21, 2, 138, 157, 5, 42, 22, 2, 139, 157, 5, 44, 23, 
	2, 140, 157, 5, 16, 9, 2, 141, 157, 5, 20, 11, 2, 142, 157, 5, 24, 13, 
	2, 143, 157, 5, 26, 14, 2, 144, 157, 5, 28, 15, 2, 145, 157, 5, 30, 16, 
	2, 146, 157, 5, 32, 17, 2, 147, 157, 5, 34, 18, 2, 148, 157, 5, 12, 7, 
	2, 149, 157, 5, 14, 8, 2, 150, 157, 5, 22, 12, 2, 151, 157, 5, 36, 19, 
	2, 152, 157, 5, 38, 20, 2, 153, 157, 5, 18, 10, 2, 154, 157, 5, 48, 25, 
	2, 155, 157, 5, 46, 24, 2, 156, 136, 3, 2, 2, 2, 156, 137, 3, 2, 2, 2, 
	156, 138, 3, 2, 2, 2, 156, 139, 3, 2, 2, 2, 156, 140, 3, 2, 2, 2, 156, 
	141, 3, 2, 2, 2, 156, 142, 3, 2, 2, 2, 156, 143, 3, 2, 2, 2, 156, 144, 
	3, 2, 2, 2, 156, 145, 3, 2, 2, 2, 156, 146, 3, 2, 2, 2, 156, 147, 3, 2, 
	2, 2, 156, 148, 3, 2, 2, 2, 156, 149, 3, 2, 2, 2, 156, 150, 3, 2, 2, 2, 
	156, 151, 3, 2, 2, 2, 156, 152, 3, 2, 2, 2, 156, 153, 3, 2, 2, 2, 156, 
	154, 3, 2, 2, 2, 156, 155, 3, 2, 2, 2, 157, 9, 3, 2, 2, 2, 158, 159, 5, 
	56, 29, 2, 159, 160, 5, 54, 28, 2, 160, 164, 7, 274, 2, 2, 161, 165, 5, 
	54, 28, 2, 162, 165, 7, 269, 2, 2, 163, 165, 5, 50, 26, 2, 164, 161, 3, 
	2, 2, 2, 164, 162, 3, 2, 2, 2, 164, 163, 3, 2, 2, 2, 165, 174, 3, 2, 2, 
	2, 166, 167, 5, 56, 29, 2, 167, 168, 5, 50, 26, 2, 168, 171, 7, 274, 2, 
	2, 169, 172, 5, 54, 28, 2, 170, 172, 7, 269, 2, 2, 171, 169, 3, 2, 2, 2, 
	171, 170, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2, 173, 158, 3, 2, 2, 2, 173, 
	166, 3, 2, 2, 2, 174, 11, 3, 2, 2, 2, 175, 179, 5, 58, 30, 2, 176, 180, 
	7, 269, 2, 2, 177, 180, 5, 54, 28, 2, 178, 180, 5, 50, 26, 2, 179, 176, 
	3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 178, 3, 2, 2, 2, 180, 13, 3, 2, 
	2, 2, 181, 184, 5, 60, 31, 2, 182, 185, 5, 54, 28, 2, 183, 185, 5, 50, 
	26, 2, 184, 182, 3, 2, 2, 2, 184, 183, 3, 2, 2, 2, 185, 15, 3, 2, 2, 2, 
	186, 187, 5, 62, 32, 2, 187, 188, 5, 54, 28, 2, 188, 191, 7, 274, 2, 2, 
	189, 192, 5, 54, 28, 2, 190, 192, 5, 50, 26, 2, 191, 189, 3, 2, 2, 2, 191, 
	190, 3, 2, 2, 2, 192, 199, 3, 2, 2, 2, 193, 194, 5, 62, 32, 2, 194, 195, 
	5, 50, 26, 2, 195, 196, 7, 274, 2, 2, 196, 197, 5, 54, 28, 2, 197, 199, 
	3, 2, 2, 2, 198, 186, 3, 2, 2, 2, 198, 193, 3, 2, 2, 2, 199, 17, 3, 2, 
	2, 2, 200, 201, 5, 64, 33, 2, 201, 19, 3, 2, 2, 2, 202, 205, 5, 66, 34, 
	2, 203, 206, 5, 54, 28, 2, 204, 206, 5, 50, 26, 2, 205, 203, 3, 2, 2, 2, 
	205, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 211, 7, 274, 2, 2, 208, 
	212, 5, 54, 28, 2, 209, 212, 7, 269, 2, 2, 210, 212, 5, 50, 26, 2, 211, 
	208, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 211, 210, 3, 2, 2, 2, 212, 21, 3, 
	2, 2, 2, 213, 214, 5, 68, 35, 2, 214, 215, 7, 12, 2, 2, 215, 23, 3, 2, 
	2, 2, 216, 217, 5, 70, 36, 2, 217, 218, 5, 54, 28, 2, 218, 221, 7, 274, 
	2, 2, 219, 222, 5, 54, 28, 2, 220, 222, 5, 50, 26, 2, 221, 219, 3, 2, 2, 
	2, 221, 220, 3, 2, 2, 2, 222, 25, 3, 2, 2, 2, 223, 224, 5, 72, 37, 2, 224, 
	225, 5, 54, 28, 2, 225, 226, 7, 274, 2, 2, 226, 227, 5, 50, 26, 2, 227, 
	27, 3, 2, 2, 2, 228, 231, 5, 74, 38, 2, 229, 232, 5, 54, 28, 2, 230, 232, 
	5, 50, 26, 2, 231, 229, 3, 2, 2, 2, 231, 230, 3, 2, 2, 2, 232, 233, 3, 
	2, 2, 2, 233, 234, 7, 274, 2, 2, 234, 235, 5, 54, 28, 2, 235, 29, 3, 2, 
	2, 2, 236, 239, 5, 76, 39, 2, 237, 240, 5, 54, 28, 2, 238, 240, 5, 50, 
	26, 2, 239, 237, 3, 2, 2, 2, 239, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 
	241, 244, 7, 274, 2, 2, 242, 245, 7, 269, 2, 2, 243, 245, 5, 54, 28, 2, 
	244, 242, 3, 2, 2, 2, 244, 243, 3, 2, 2, 2, 245, 31, 3, 2, 2, 2, 246, 249, 
	5, 78, 40, 2, 247, 250, 5, 54, 28, 2, 248, 250, 5, 50, 26, 2, 249, 247, 
	3, 2, 2, 2, 249, 248, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 7, 274, 
	2, 2, 252, 253, 5, 54, 28, 2, 253, 256, 7, 274, 2, 2, 254, 257, 5, 54, 
	28, 2, 255, 257, 7, 269, 2, 2, 256, 254, 3, 2, 2, 2, 256, 255, 3, 2, 2, 
	2, 257, 33, 3, 2, 2, 2, 258, 261, 5, 80, 41, 2, 259, 262, 5, 54, 28, 2, 
	260, 262, 5, 50, 26, 2, 261, 259, 3, 2, 2, 2, 261, 260, 3, 2, 2, 2, 262, 
	263, 3, 2, 2, 2, 263, 266, 7, 274, 2, 2, 264, 267, 5, 54, 28, 2, 265, 267, 
	5, 50, 26, 2, 266, 264, 3, 2, 2, 2, 266, 265, 3, 2, 2, 2, 267, 35, 3, 2, 
	2, 2, 268, 272, 5, 82, 42, 2, 269, 273, 5, 54, 28, 2, 270, 273, 5, 50, 
	26, 2, 271, 273, 7, 269, 2, 2, 272, 269, 3, 2, 2, 2, 272, 270, 3, 2, 2, 
	2, 272, 271, 3, 2, 2, 2, 273, 37, 3, 2, 2, 2, 274, 275, 5, 84, 43, 2, 275, 
	276, 7, 269, 2, 2, 276, 39, 3, 2, 2, 2, 277, 278, 5, 86, 44, 2, 278, 279, 
	5, 54, 28, 2, 279, 282, 7, 274, 2, 2, 280, 283, 5, 54, 28, 2, 281, 283, 
	7, 269, 2, 2, 282, 280, 3, 2, 2, 2, 282, 281, 3, 2, 2, 2, 283, 41, 3, 2, 
	2, 2, 284, 287, 5, 88, 45, 2, 285, 288, 5, 54, 28, 2, 286, 288, 7, 269, 
	2, 2, 287, 285, 3, 2, 2, 2, 287, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 
	289, 290, 7, 274, 2, 2, 290, 291, 5, 54, 28, 2, 291, 43, 3, 2, 2, 2, 292, 
	293, 5, 90, 46, 2, 293, 294, 5, 64, 33, 2, 294, 45, 3, 2, 2, 2, 295, 296, 
	5, 92, 47, 2, 296, 297, 7, 12, 2, 2, 297, 300, 3, 2, 2, 2, 298, 300, 5, 
	92, 47, 2, 299, 295, 3, 2, 2, 2, 299, 298, 3, 2, 2, 2, 300, 47, 3, 2, 2, 
	2, 301, 302, 7, 12, 2, 2, 302, 306, 5, 94, 48, 2, 303, 307, 5, 96, 49, 
	2, 304, 307, 7, 272, 2, 2, 305, 307, 7, 269, 2, 2, 306, 303, 3, 2, 2, 2, 
	306, 304, 3, 2, 2, 2, 306, 305, 3, 2, 2, 2, 307, 49, 3, 2, 2, 2, 308, 311, 
	7, 9, 2, 2, 309, 312, 5, 54, 28, 2, 310, 312, 7, 12, 2, 2, 311, 309, 3, 
	2, 2, 2, 311, 310, 3, 2, 2, 2, 312, 324, 3, 2, 2, 2, 313, 322, 7, 10, 2, 
	2, 314, 317, 5, 54, 28, 2, 315, 316, 7, 10, 2, 2, 316, 318, 9, 2, 2, 2, 
	317, 315, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 323, 3, 2, 2, 2, 319, 
	323, 7, 269, 2, 2, 320, 323, 7, 268, 2, 2, 321, 323, 7, 270, 2, 2, 322, 
	314, 3, 2, 2, 2, 322, 319, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 321, 
	3, 2, 2, 2, 323, 325, 3, 2, 2, 2, 324, 313, 3, 2, 2, 2, 324, 325, 3, 2, 
	2, 2, 325, 326, 3, 2, 2, 2, 326, 327, 7, 11, 2, 2, 327, 51, 3, 2, 2, 2, 
	328, 329, 9, 3, 2, 2, 329, 53, 3, 2, 2, 2, 330, 331, 9, 4, 2, 2, 331, 55, 
	3, 2, 2, 2, 332, 333, 9, 5, 2, 2, 333, 57, 3, 2, 2, 2, 334, 335, 7, 46, 
	2, 2, 335, 59, 3, 2, 2, 2, 336, 337, 9, 6, 2, 2, 337, 61, 3, 2, 2, 2, 338, 
	339, 7, 86, 2, 2, 339, 63, 3, 2, 2, 2, 340, 341, 9, 7, 2, 2, 341, 65, 3, 
	2, 2, 2, 342, 343, 9, 8, 2, 2, 343, 67, 3, 2, 2, 2, 344, 345, 9, 9, 2, 
	2, 345, 69, 3, 2, 2, 2, 346, 347, 9, 10, 2, 2, 347, 71, 3, 2, 2, 2, 348, 
	349, 9, 11, 2, 2, 349, 73, 3, 2, 2, 2, 350, 351, 9, 12, 2, 2, 351, 75, 
	3, 2, 2, 2, 352, 353, 9, 13, 2, 2, 353, 77, 3, 2, 2, 2, 354, 355, 9, 14, 
	2, 2, 355, 79, 3, 2, 2, 2, 356, 357, 9, 15, 2, 2, 357, 81, 3, 2, 2, 2, 
	358, 359, 7, 207, 2, 2, 359, 83, 3, 2, 2, 2, 360, 361, 9, 16, 2, 2, 361, 
	85, 3, 2, 2, 2, 362, 363, 7, 212, 2, 2, 363, 87, 3, 2, 2, 2, 364, 365, 
	7, 213, 2, 2, 365, 89, 3, 2, 2, 2, 366, 367, 9, 17, 2, 2, 367, 91, 3, 2, 
	2, 2, 368, 369, 9, 18, 2, 2, 369, 93, 3, 2, 2, 2, 370, 371, 9, 19, 2, 2, 
	371, 95, 3, 2, 2, 2, 372, 373, 7, 266, 2, 2, 373, 97, 3, 2, 2, 2, 374, 
	375, 7, 267, 2, 2, 375, 99, 3, 2, 2, 2, 34, 102, 104, 116, 118, 129, 156, 
	164, 171, 173, 179, 184, 191, 198, 205, 211, 221, 231, 239, 244, 249, 256, 
	261, 266, 272, 282, 287, 299, 306, 311, 317, 322, 324,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'end'", "'segments'", "'para'", "'public'", "'ends'", "'endp'", "'['", 
	"'+'", "']'", "", "'ds'", "'es'", "'cs'", "'ss'", "'gs'", "'fs'", "'ah'", 
	"'al'", "'ax'", "'bh'", "'bl'", "'bx'", "'ch'", "'cl'", "'cx'", "'dh'", 
	"'dl'", "'dx'", "'si'", "'di'", "'sp'", "'bp'", "'eax'", "'ebx'", "'ecx'", 
	"'edx'", "'esi'", "'edi'", "'esp'", "'ebp'", "'mov'", "'cmp'", "'test'", 
	"'push'", "'pop'", "'idiv'", "'inc'", "'dec'", "'neg'", "'mul'", "'div'", 
	"'imul'", "'not'", "'setpo'", "'setae'", "'setnle'", "'setc'", "'setno'", 
	"'setnb'", "'setp'", "'setnge'", "'setl'", "'setge'", "'setpe'", "'setnl'", 
	"'setnz'", "'setne'", "'setnc'", "'setbe'", "'setnp'", "'setns'", "'seto'", 
	"'setna'", "'setnae'", "'setz'", "'setle'", "'setnbe'", "'sets'", "'sete'", 
	"'setb'", "'seta'", "'setg'", "'setng'", "'xchg'", "'popad'", "'aaa'", 
	"'popa'", "'popfd'", "'cwd'", "'lahf'", "'pushad'", "'pushf'", "'aas'", 
	"'bswap'", "'pushfd'", "'cbw'", "'cwde'", "'xlat'", "'aam'", "'aad'", "'cdq'", 
	"'daa'", "'sahf'", "'das'", "'into'", "'iret'", "'clc'", "'stc'", "'cmc'", 
	"'cld'", "'std'", "'cli'", "'sti'", "'movsb'", "'movsw'", "'movsd'", "'lods'", 
	"'lodsb'", "'lodsw'", "'lodsd'", "'stos'", "'stosb'", "'stosw'", "'sotsd'", 
	"'scas'", "'scasb'", "'scasw'", "'scasd'", "'cmps'", "'cmpsb'", "'cmpsw'", 
	"'cmpsd'", "'insb'", "'insw'", "'insd'", "'outsb'", "'outsw'", "'outsd'", 
	"'adc'", "'add'", "'sub'", "'cbb'", "'xor'", "'or'", "'jnbe'", "'jnz'", 
	"'jpo'", "'jz'", "'js'", "'loopnz'", "'jge'", "'jb'", "'jnb'", "'jo'", 
	"'jp'", "'jno'", "'jnl'", "'jnae'", "'loopz'", "'jmp'", "'jnp'", "'loop'", 
	"'jl'", "'jcxz'", "'jae'", "'jnge'", "'ja'", "'loopne'", "'loope'", "'jg'", 
	"'jnle'", "'je'", "'jnc'", "'jc'", "'jna'", "'jbe'", "'jle'", "'jpe'", 
	"'jns'", "'jecxz'", "'jng'", "'movzx'", "'bsf'", "'bsr'", "'les'", "'lea'", 
	"'lds'", "'ins'", "'outs'", "'xadd'", "'cmpxchg'", "'shl'", "'shr'", "'ror'", 
	"'rol'", "'rcl'", "'sal'", "'rcr'", "'sar'", "'shrd'", "'shld'", "'btr'", 
	"'bt'", "'btc'", "'call'", "'int'", "'retn'", "'ret'", "'retf'", "'in'", 
	"'out'", "'rep'", "'repe'", "'repz'", "'repne'", "'repnz'", "'.alpha'", 
	"'.const'", "'.cref'", "'.xcref'", "'data'", "'data?'", "'dosseg'", "'.err'", 
	"'.err1'", "'.err2'", "'.erre'", "'.errnz'", "'.errdef'", "'.errndef'", 
	"'.errb'", "'.errnb'", "'.erridn[i]'", "'.errdif[i]'", "'even'", "'.list'", 
	"'width'", "'mask'", "'.seq'", "'.xlist'", "'.exit'", "'.model'", "'byte'", 
	"'sbyte'", "'db'", "'word'", "'sword'", "'dw'", "'dword'", "'sdword'", 
	"'dd'", "'fword'", "'df'", "'qword'", "'dq'", "'tbyte'", "'dt'", "'real4'", 
	"'real8'", "'real'", "'far'", "'near'", "'proc'", "'?'", "'times'", "", 
	"", "", "", "", "", "','",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "Identifier", "DS", "ES", "CS", 
	"SS", "GS", "FS", "AH", "AL", "AX", "BH", "BL", "BX", "CH", "CL", "CX", 
	"DH", "DL", "DX", "SI", "DI", "SP", "BP", "EAX", "EBX", "ECX", "EDX", "ESI", 
	"EDI", "ESP", "EBP", "MOV", "CMP", "TEST", "PUSH", "POP", "IDIV", "INC", 
	"DEC", "NEG", "MUL", "DIV", "IMUL", "NOT", "SETPO", "SETAE", "SETNLE", 
	"SETC", "SETNO", "SETNB", "SETP", "SETNGE", "SETL", "SETGE", "SETPE", "SETNL", 
	"SETNZ", "SETNE", "SETNC", "SETBE", "SETNP", "SETNS", "SETO", "SETNA", 
	"SETNAE", "SETZ", "SETLE", "SETNBE", "SETS", "SETE", "SETB", "SETA", "SETG", 
	"SETNG", "XCHG", "POPAD", "AAA", "POPA", "POPFD", "CWD", "LAHF", "PUSHAD", 
	"PUSHF", "AAS", "BSWAP", "PUSHFD", "CBW", "CWDE", "XLAT", "AAM", "AAD", 
	"CDQ", "DAA", "SAHF", "DAS", "INTO", "IRET", "CLC", "STC", "CMC", "CLD", 
	"STD", "CLI", "STI", "MOVSB", "MOVSW", "MOVSD", "LODS", "LODSB", "LODSW", 
	"LODSD", "STOS", "STOSB", "STOSW", "SOTSD", "SCAS", "SCASB", "SCASW", "SCASD", 
	"CMPS", "CMPSB", "CMPSW", "CMPSD", "INSB", "INSW", "INSD", "OUTSB", "OUTSW", 
	"OUTSD", "ADC", "ADD", "SUB", "CBB", "XOR", "OR", "JNBE", "JNZ", "JPO", 
	"JZ", "JS", "LOOPNZ", "JGE", "JB", "JNB", "JO", "JP", "JNO", "JNL", "JNAE", 
	"LOOPZ", "JMP", "JNP", "LOOP", "JL", "JCXZ", "JAE", "JNGE", "JA", "LOOPNE", 
	"LOOPE", "JG", "JNLE", "JE", "JNC", "JC", "JNA", "JBE", "JLE", "JPE", "JNS", 
	"JECXZ", "JNG", "MOVZX", "BSF", "BSR", "LES", "LEA", "LDS", "INS", "OUTS", 
	"XADD", "CMPXCHG", "SHL", "SHR", "ROR", "ROL", "RCL", "SAL", "RCR", "SAR", 
	"SHRD", "SHLD", "BTR", "BT", "BTC", "CALL", "INT", "RETN", "RET", "RETF", 
	"IN", "OUT", "REP", "REPE", "REPZ", "REPNE", "REPNZ", "ALPHA", "CONST", 
	"CREF", "XCREF", "DATA", "DATA2", "DOSSEG", "ERR", "ERR1", "ERR2", "ERRE", 
	"ERRNZ", "ERRDEF", "ERRNDEF", "ERRB", "ERRNB", "ERRIDN", "ERRDIF", "EVEN", 
	"LIST", "WIDTH", "MASK", "SEQ", "XLIST", "EXIT", "MODEL", "BYTE", "SBYTE", 
	"DB", "WORD", "SWORD", "DW", "DWORD", "SDWORD", "DD", "FWORD", "DF", "QWORD", 
	"DQ", "TBYTE", "DT", "REAL4", "REAL8", "REAL", "FAR", "NEAR", "PROC", "QUESTION", 
	"TIMES", "Hexnum", "Integer", "Octalnum", "FloatingPointLiteral", "String", 
	"Etiqueta", "Separator", "WS", "LINE_COMMENT",
}

var ruleNames = []string{
	"compilationUnit", "segments", "proc", "code", "binary_exp1", "unuary_exp1", 
	"unuary_exp2", "binary_exp2", "notarguments", "binary_exp3", "unuary_exp3", 
	"binary_exp4", "binary_exp5", "binary_exp6", "binary_exp7", "binary_exp8", 
	"binary_exp9", "unuary_exp4", "unuary_exp5", "binary_exp10", "binary_exp11", 
	"binary_exp12", "directive_exp1", "variabledeclaration", "memory", "segmentos", 
	"register", "o", "op", "ope", "oper", "opera", "operat", "operato", "operator", 
	"l", "x", "s", "sh", "b", "call", "interruption", "in", "out", "re", "directives", 
	"ty", "question", "time",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MASMParser struct {
	*antlr.BaseParser
}

func NewMASMParser(input antlr.TokenStream) *MASMParser {
	this := new(MASMParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MASM.g4"

	return this
}

// MASMParser tokens.
const (
	MASMParserEOF = antlr.TokenEOF
	MASMParserT__0 = 1
	MASMParserT__1 = 2
	MASMParserT__2 = 3
	MASMParserT__3 = 4
	MASMParserT__4 = 5
	MASMParserT__5 = 6
	MASMParserT__6 = 7
	MASMParserT__7 = 8
	MASMParserT__8 = 9
	MASMParserIdentifier = 10
	MASMParserDS = 11
	MASMParserES = 12
	MASMParserCS = 13
	MASMParserSS = 14
	MASMParserGS = 15
	MASMParserFS = 16
	MASMParserAH = 17
	MASMParserAL = 18
	MASMParserAX = 19
	MASMParserBH = 20
	MASMParserBL = 21
	MASMParserBX = 22
	MASMParserCH = 23
	MASMParserCL = 24
	MASMParserCX = 25
	MASMParserDH = 26
	MASMParserDL = 27
	MASMParserDX = 28
	MASMParserSI = 29
	MASMParserDI = 30
	MASMParserSP = 31
	MASMParserBP = 32
	MASMParserEAX = 33
	MASMParserEBX = 34
	MASMParserECX = 35
	MASMParserEDX = 36
	MASMParserESI = 37
	MASMParserEDI = 38
	MASMParserESP = 39
	MASMParserEBP = 40
	MASMParserMOV = 41
	MASMParserCMP = 42
	MASMParserTEST = 43
	MASMParserPUSH = 44
	MASMParserPOP = 45
	MASMParserIDIV = 46
	MASMParserINC = 47
	MASMParserDEC = 48
	MASMParserNEG = 49
	MASMParserMUL = 50
	MASMParserDIV = 51
	MASMParserIMUL = 52
	MASMParserNOT = 53
	MASMParserSETPO = 54
	MASMParserSETAE = 55
	MASMParserSETNLE = 56
	MASMParserSETC = 57
	MASMParserSETNO = 58
	MASMParserSETNB = 59
	MASMParserSETP = 60
	MASMParserSETNGE = 61
	MASMParserSETL = 62
	MASMParserSETGE = 63
	MASMParserSETPE = 64
	MASMParserSETNL = 65
	MASMParserSETNZ = 66
	MASMParserSETNE = 67
	MASMParserSETNC = 68
	MASMParserSETBE = 69
	MASMParserSETNP = 70
	MASMParserSETNS = 71
	MASMParserSETO = 72
	MASMParserSETNA = 73
	MASMParserSETNAE = 74
	MASMParserSETZ = 75
	MASMParserSETLE = 76
	MASMParserSETNBE = 77
	MASMParserSETS = 78
	MASMParserSETE = 79
	MASMParserSETB = 80
	MASMParserSETA = 81
	MASMParserSETG = 82
	MASMParserSETNG = 83
	MASMParserXCHG = 84
	MASMParserPOPAD = 85
	MASMParserAAA = 86
	MASMParserPOPA = 87
	MASMParserPOPFD = 88
	MASMParserCWD = 89
	MASMParserLAHF = 90
	MASMParserPUSHAD = 91
	MASMParserPUSHF = 92
	MASMParserAAS = 93
	MASMParserBSWAP = 94
	MASMParserPUSHFD = 95
	MASMParserCBW = 96
	MASMParserCWDE = 97
	MASMParserXLAT = 98
	MASMParserAAM = 99
	MASMParserAAD = 100
	MASMParserCDQ = 101
	MASMParserDAA = 102
	MASMParserSAHF = 103
	MASMParserDAS = 104
	MASMParserINTO = 105
	MASMParserIRET = 106
	MASMParserCLC = 107
	MASMParserSTC = 108
	MASMParserCMC = 109
	MASMParserCLD = 110
	MASMParserSTD = 111
	MASMParserCLI = 112
	MASMParserSTI = 113
	MASMParserMOVSB = 114
	MASMParserMOVSW = 115
	MASMParserMOVSD = 116
	MASMParserLODS = 117
	MASMParserLODSB = 118
	MASMParserLODSW = 119
	MASMParserLODSD = 120
	MASMParserSTOS = 121
	MASMParserSTOSB = 122
	MASMParserSTOSW = 123
	MASMParserSOTSD = 124
	MASMParserSCAS = 125
	MASMParserSCASB = 126
	MASMParserSCASW = 127
	MASMParserSCASD = 128
	MASMParserCMPS = 129
	MASMParserCMPSB = 130
	MASMParserCMPSW = 131
	MASMParserCMPSD = 132
	MASMParserINSB = 133
	MASMParserINSW = 134
	MASMParserINSD = 135
	MASMParserOUTSB = 136
	MASMParserOUTSW = 137
	MASMParserOUTSD = 138
	MASMParserADC = 139
	MASMParserADD = 140
	MASMParserSUB = 141
	MASMParserCBB = 142
	MASMParserXOR = 143
	MASMParserOR = 144
	MASMParserJNBE = 145
	MASMParserJNZ = 146
	MASMParserJPO = 147
	MASMParserJZ = 148
	MASMParserJS = 149
	MASMParserLOOPNZ = 150
	MASMParserJGE = 151
	MASMParserJB = 152
	MASMParserJNB = 153
	MASMParserJO = 154
	MASMParserJP = 155
	MASMParserJNO = 156
	MASMParserJNL = 157
	MASMParserJNAE = 158
	MASMParserLOOPZ = 159
	MASMParserJMP = 160
	MASMParserJNP = 161
	MASMParserLOOP = 162
	MASMParserJL = 163
	MASMParserJCXZ = 164
	MASMParserJAE = 165
	MASMParserJNGE = 166
	MASMParserJA = 167
	MASMParserLOOPNE = 168
	MASMParserLOOPE = 169
	MASMParserJG = 170
	MASMParserJNLE = 171
	MASMParserJE = 172
	MASMParserJNC = 173
	MASMParserJC = 174
	MASMParserJNA = 175
	MASMParserJBE = 176
	MASMParserJLE = 177
	MASMParserJPE = 178
	MASMParserJNS = 179
	MASMParserJECXZ = 180
	MASMParserJNG = 181
	MASMParserMOVZX = 182
	MASMParserBSF = 183
	MASMParserBSR = 184
	MASMParserLES = 185
	MASMParserLEA = 186
	MASMParserLDS = 187
	MASMParserINS = 188
	MASMParserOUTS = 189
	MASMParserXADD = 190
	MASMParserCMPXCHG = 191
	MASMParserSHL = 192
	MASMParserSHR = 193
	MASMParserROR = 194
	MASMParserROL = 195
	MASMParserRCL = 196
	MASMParserSAL = 197
	MASMParserRCR = 198
	MASMParserSAR = 199
	MASMParserSHRD = 200
	MASMParserSHLD = 201
	MASMParserBTR = 202
	MASMParserBT = 203
	MASMParserBTC = 204
	MASMParserCALL = 205
	MASMParserINT = 206
	MASMParserRETN = 207
	MASMParserRET = 208
	MASMParserRETF = 209
	MASMParserIN = 210
	MASMParserOUT = 211
	MASMParserREP = 212
	MASMParserREPE = 213
	MASMParserREPZ = 214
	MASMParserREPNE = 215
	MASMParserREPNZ = 216
	MASMParserALPHA = 217
	MASMParserCONST = 218
	MASMParserCREF = 219
	MASMParserXCREF = 220
	MASMParserDATA = 221
	MASMParserDATA2 = 222
	MASMParserDOSSEG = 223
	MASMParserERR = 224
	MASMParserERR1 = 225
	MASMParserERR2 = 226
	MASMParserERRE = 227
	MASMParserERRNZ = 228
	MASMParserERRDEF = 229
	MASMParserERRNDEF = 230
	MASMParserERRB = 231
	MASMParserERRNB = 232
	MASMParserERRIDN = 233
	MASMParserERRDIF = 234
	MASMParserEVEN = 235
	MASMParserLIST = 236
	MASMParserWIDTH = 237
	MASMParserMASK = 238
	MASMParserSEQ = 239
	MASMParserXLIST = 240
	MASMParserEXIT = 241
	MASMParserMODEL = 242
	MASMParserBYTE = 243
	MASMParserSBYTE = 244
	MASMParserDB = 245
	MASMParserWORD = 246
	MASMParserSWORD = 247
	MASMParserDW = 248
	MASMParserDWORD = 249
	MASMParserSDWORD = 250
	MASMParserDD = 251
	MASMParserFWORD = 252
	MASMParserDF = 253
	MASMParserQWORD = 254
	MASMParserDQ = 255
	MASMParserTBYTE = 256
	MASMParserDT = 257
	MASMParserREAL4 = 258
	MASMParserREAL8 = 259
	MASMParserREAL = 260
	MASMParserFAR = 261
	MASMParserNEAR = 262
	MASMParserPROC = 263
	MASMParserQUESTION = 264
	MASMParserTIMES = 265
	MASMParserHexnum = 266
	MASMParserInteger = 267
	MASMParserOctalnum = 268
	MASMParserFloatingPointLiteral = 269
	MASMParserString = 270
	MASMParserEtiqueta = 271
	MASMParserSeparator = 272
	MASMParserWS = 273
	MASMParserLINE_COMMENT = 274
)

// MASMParser rules.
const (
	MASMParserRULE_compilationUnit = 0
	MASMParserRULE_segments = 1
	MASMParserRULE_proc = 2
	MASMParserRULE_code = 3
	MASMParserRULE_binary_exp1 = 4
	MASMParserRULE_unuary_exp1 = 5
	MASMParserRULE_unuary_exp2 = 6
	MASMParserRULE_binary_exp2 = 7
	MASMParserRULE_notarguments = 8
	MASMParserRULE_binary_exp3 = 9
	MASMParserRULE_unuary_exp3 = 10
	MASMParserRULE_binary_exp4 = 11
	MASMParserRULE_binary_exp5 = 12
	MASMParserRULE_binary_exp6 = 13
	MASMParserRULE_binary_exp7 = 14
	MASMParserRULE_binary_exp8 = 15
	MASMParserRULE_binary_exp9 = 16
	MASMParserRULE_unuary_exp4 = 17
	MASMParserRULE_unuary_exp5 = 18
	MASMParserRULE_binary_exp10 = 19
	MASMParserRULE_binary_exp11 = 20
	MASMParserRULE_binary_exp12 = 21
	MASMParserRULE_directive_exp1 = 22
	MASMParserRULE_variabledeclaration = 23
	MASMParserRULE_memory = 24
	MASMParserRULE_segmentos = 25
	MASMParserRULE_register = 26
	MASMParserRULE_o = 27
	MASMParserRULE_op = 28
	MASMParserRULE_ope = 29
	MASMParserRULE_oper = 30
	MASMParserRULE_opera = 31
	MASMParserRULE_operat = 32
	MASMParserRULE_operato = 33
	MASMParserRULE_operator = 34
	MASMParserRULE_l = 35
	MASMParserRULE_x = 36
	MASMParserRULE_s = 37
	MASMParserRULE_sh = 38
	MASMParserRULE_b = 39
	MASMParserRULE_call = 40
	MASMParserRULE_interruption = 41
	MASMParserRULE_in = 42
	MASMParserRULE_out = 43
	MASMParserRULE_re = 44
	MASMParserRULE_directives = 45
	MASMParserRULE_ty = 46
	MASMParserRULE_question = 47
	MASMParserRULE_time = 48
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_compilationUnit
	return p
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MASMParserIdentifier, 0)
}

func (s *CompilationUnitContext) AllSegments() []ISegmentsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISegmentsContext)(nil)).Elem())
	var tst = make([]ISegmentsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISegmentsContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) Segments(i int) ISegmentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISegmentsContext)
}

func (s *CompilationUnitContext) AllDirective_exp1() []IDirective_exp1Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirective_exp1Context)(nil)).Elem())
	var tst = make([]IDirective_exp1Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirective_exp1Context)
		}
	}

	return tst
}

func (s *CompilationUnitContext) Directive_exp1(i int) IDirective_exp1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirective_exp1Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirective_exp1Context)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MASMParserRULE_compilationUnit)
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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MASMParserIdentifier || ((((_la - 217)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 217))) & ((1 << (MASMParserALPHA - 217)) | (1 << (MASMParserCONST - 217)) | (1 << (MASMParserCREF - 217)) | (1 << (MASMParserXCREF - 217)) | (1 << (MASMParserDATA - 217)) | (1 << (MASMParserDATA2 - 217)) | (1 << (MASMParserDOSSEG - 217)) | (1 << (MASMParserERR - 217)) | (1 << (MASMParserERR1 - 217)) | (1 << (MASMParserERR2 - 217)) | (1 << (MASMParserERRE - 217)) | (1 << (MASMParserERRNZ - 217)) | (1 << (MASMParserERRDEF - 217)) | (1 << (MASMParserERRNDEF - 217)) | (1 << (MASMParserERRB - 217)) | (1 << (MASMParserERRNB - 217)) | (1 << (MASMParserERRIDN - 217)) | (1 << (MASMParserERRDIF - 217)) | (1 << (MASMParserEVEN - 217)) | (1 << (MASMParserLIST - 217)) | (1 << (MASMParserWIDTH - 217)) | (1 << (MASMParserMASK - 217)) | (1 << (MASMParserSEQ - 217)) | (1 << (MASMParserXLIST - 217)) | (1 << (MASMParserEXIT - 217)) | (1 << (MASMParserMODEL - 217)))) != 0) {
		p.SetState(100)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MASMParserIdentifier:
			{
				p.SetState(98)
				p.Segments()
			}


		case MASMParserALPHA, MASMParserCONST, MASMParserCREF, MASMParserXCREF, MASMParserDATA, MASMParserDATA2, MASMParserDOSSEG, MASMParserERR, MASMParserERR1, MASMParserERR2, MASMParserERRE, MASMParserERRNZ, MASMParserERRDEF, MASMParserERRNDEF, MASMParserERRB, MASMParserERRNB, MASMParserERRIDN, MASMParserERRDIF, MASMParserEVEN, MASMParserLIST, MASMParserWIDTH, MASMParserMASK, MASMParserSEQ, MASMParserXLIST, MASMParserEXIT, MASMParserMODEL:
			{
				p.SetState(99)
				p.Directive_exp1()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
		p.Match(MASMParserT__0)
	}
	{
		p.SetState(106)
		p.Match(MASMParserIdentifier)
	}



	return localctx
}


// ISegmentsContext is an interface to support dynamic dispatch.
type ISegmentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentsContext differentiates from other interfaces.
	IsSegmentsContext()
}

type SegmentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentsContext() *SegmentsContext {
	var p = new(SegmentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_segments
	return p
}

func (*SegmentsContext) IsSegmentsContext() {}

func NewSegmentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentsContext {
	var p = new(SegmentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_segments

	return p
}

func (s *SegmentsContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentsContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(MASMParserIdentifier)
}

func (s *SegmentsContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(MASMParserIdentifier, i)
}

func (s *SegmentsContext) AllCode() []ICodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICodeContext)(nil)).Elem())
	var tst = make([]ICodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICodeContext)
		}
	}

	return tst
}

func (s *SegmentsContext) Code(i int) ICodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *SegmentsContext) AllProc() []IProcContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProcContext)(nil)).Elem())
	var tst = make([]IProcContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProcContext)
		}
	}

	return tst
}

func (s *SegmentsContext) Proc(i int) IProcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProcContext)
}

func (s *SegmentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SegmentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterSegments(s)
	}
}

func (s *SegmentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitSegments(s)
	}
}

func (s *SegmentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitSegments(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Segments() (localctx ISegmentsContext) {
	localctx = NewSegmentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MASMParserRULE_segments)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(108)
		p.Match(MASMParserIdentifier)
	}
	{
		p.SetState(109)
		p.Match(MASMParserT__1)
	}
	{
		p.SetState(110)
		p.Match(MASMParserT__2)
	}
	{
		p.SetState(111)
		p.Match(MASMParserT__3)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(112)
					p.Code()
				}


			case 2:
				{
					p.SetState(113)
					p.Proc()
				}

			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	{
		p.SetState(119)
		p.Match(MASMParserIdentifier)
	}
	{
		p.SetState(120)
		p.Match(MASMParserT__4)
	}



	return localctx
}


// IProcContext is an interface to support dynamic dispatch.
type IProcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcContext differentiates from other interfaces.
	IsProcContext()
}

type ProcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcContext() *ProcContext {
	var p = new(ProcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_proc
	return p
}

func (*ProcContext) IsProcContext() {}

func NewProcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcContext {
	var p = new(ProcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_proc

	return p
}

func (s *ProcContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(MASMParserIdentifier)
}

func (s *ProcContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(MASMParserIdentifier, i)
}

func (s *ProcContext) AllCode() []ICodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICodeContext)(nil)).Elem())
	var tst = make([]ICodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICodeContext)
		}
	}

	return tst
}

func (s *ProcContext) Code(i int) ICodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *ProcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterProc(s)
	}
}

func (s *ProcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitProc(s)
	}
}

func (s *ProcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitProc(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Proc() (localctx IProcContext) {
	localctx = NewProcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MASMParserRULE_proc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(122)
		p.Match(MASMParserIdentifier)
	}
	{
		p.SetState(123)
		p.Match(MASMParserPROC)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(124)
				p.Code()
			}


		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	{
		p.SetState(130)
		p.Match(MASMParserRET)
	}
	{
		p.SetState(131)
		p.Match(MASMParserIdentifier)
	}
	{
		p.SetState(132)
		p.Match(MASMParserT__5)
	}



	return localctx
}


// ICodeContext is an interface to support dynamic dispatch.
type ICodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCodeContext differentiates from other interfaces.
	IsCodeContext()
}

type CodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeContext() *CodeContext {
	var p = new(CodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_code
	return p
}

func (*CodeContext) IsCodeContext() {}

func NewCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeContext {
	var p = new(CodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_code

	return p
}

func (s *CodeContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeContext) Binary_exp1() IBinary_exp1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp1Context)
}

func (s *CodeContext) Binary_exp10() IBinary_exp10Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp10Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp10Context)
}

func (s *CodeContext) Binary_exp11() IBinary_exp11Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp11Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp11Context)
}

func (s *CodeContext) Binary_exp12() IBinary_exp12Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp12Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp12Context)
}

func (s *CodeContext) Binary_exp2() IBinary_exp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp2Context)
}

func (s *CodeContext) Binary_exp3() IBinary_exp3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp3Context)
}

func (s *CodeContext) Binary_exp4() IBinary_exp4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp4Context)
}

func (s *CodeContext) Binary_exp5() IBinary_exp5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp5Context)
}

func (s *CodeContext) Binary_exp6() IBinary_exp6Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp6Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp6Context)
}

func (s *CodeContext) Binary_exp7() IBinary_exp7Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp7Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp7Context)
}

func (s *CodeContext) Binary_exp8() IBinary_exp8Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp8Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp8Context)
}

func (s *CodeContext) Binary_exp9() IBinary_exp9Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_exp9Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_exp9Context)
}

func (s *CodeContext) Unuary_exp1() IUnuary_exp1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnuary_exp1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnuary_exp1Context)
}

func (s *CodeContext) Unuary_exp2() IUnuary_exp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnuary_exp2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnuary_exp2Context)
}

func (s *CodeContext) Unuary_exp3() IUnuary_exp3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnuary_exp3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnuary_exp3Context)
}

func (s *CodeContext) Unuary_exp4() IUnuary_exp4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnuary_exp4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnuary_exp4Context)
}

func (s *CodeContext) Unuary_exp5() IUnuary_exp5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnuary_exp5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnuary_exp5Context)
}

func (s *CodeContext) Notarguments() INotargumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotargumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotargumentsContext)
}

func (s *CodeContext) Variabledeclaration() IVariabledeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariabledeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariabledeclarationContext)
}

func (s *CodeContext) Directive_exp1() IDirective_exp1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirective_exp1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirective_exp1Context)
}

func (s *CodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterCode(s)
	}
}

func (s *CodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitCode(s)
	}
}

func (s *CodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitCode(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Code() (localctx ICodeContext) {
	localctx = NewCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MASMParserRULE_code)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserMOV, MASMParserCMP, MASMParserTEST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Binary_exp1()
		}


	case MASMParserIN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Binary_exp10()
		}


	case MASMParserOUT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.Binary_exp11()
		}


	case MASMParserREP, MASMParserREPE, MASMParserREPZ, MASMParserREPNE, MASMParserREPNZ:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
			p.Binary_exp12()
		}


	case MASMParserXCHG:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(138)
			p.Binary_exp2()
		}


	case MASMParserADC, MASMParserADD, MASMParserSUB, MASMParserCBB, MASMParserXOR, MASMParserOR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(139)
			p.Binary_exp3()
		}


	case MASMParserMOVZX, MASMParserBSF, MASMParserBSR:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(140)
			p.Binary_exp4()
		}


	case MASMParserLES, MASMParserLEA, MASMParserLDS, MASMParserINS, MASMParserOUTS:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(141)
			p.Binary_exp5()
		}


	case MASMParserXADD, MASMParserCMPXCHG:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(142)
			p.Binary_exp6()
		}


	case MASMParserSHL, MASMParserSHR, MASMParserROR, MASMParserROL, MASMParserRCL, MASMParserSAL, MASMParserRCR, MASMParserSAR:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(143)
			p.Binary_exp7()
		}


	case MASMParserSHRD, MASMParserSHLD:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(144)
			p.Binary_exp8()
		}


	case MASMParserBTR, MASMParserBT, MASMParserBTC:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(145)
			p.Binary_exp9()
		}


	case MASMParserPUSH:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(146)
			p.Unuary_exp1()
		}


	case MASMParserPOP, MASMParserIDIV, MASMParserINC, MASMParserDEC, MASMParserNEG, MASMParserMUL, MASMParserDIV, MASMParserIMUL, MASMParserNOT, MASMParserSETPO, MASMParserSETAE, MASMParserSETNLE, MASMParserSETC, MASMParserSETNO, MASMParserSETNB, MASMParserSETP, MASMParserSETNGE, MASMParserSETL, MASMParserSETGE, MASMParserSETPE, MASMParserSETNL, MASMParserSETNZ, MASMParserSETNE, MASMParserSETNC, MASMParserSETBE, MASMParserSETNP, MASMParserSETNS, MASMParserSETO, MASMParserSETNA, MASMParserSETNAE, MASMParserSETZ, MASMParserSETLE, MASMParserSETNBE, MASMParserSETS, MASMParserSETE, MASMParserSETB, MASMParserSETA, MASMParserSETG, MASMParserSETNG:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(147)
			p.Unuary_exp2()
		}


	case MASMParserJNBE, MASMParserJNZ, MASMParserJPO, MASMParserJZ, MASMParserJS, MASMParserLOOPNZ, MASMParserJGE, MASMParserJB, MASMParserJNB, MASMParserJO, MASMParserJP, MASMParserJNO, MASMParserJNL, MASMParserJNAE, MASMParserLOOPZ, MASMParserJMP, MASMParserJNP, MASMParserLOOP, MASMParserJL, MASMParserJCXZ, MASMParserJAE, MASMParserJNGE, MASMParserJA, MASMParserLOOPNE, MASMParserLOOPE, MASMParserJG, MASMParserJNLE, MASMParserJE, MASMParserJNC, MASMParserJC, MASMParserJNA, MASMParserJBE, MASMParserJLE, MASMParserJPE, MASMParserJNS, MASMParserJECXZ, MASMParserJNG:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(148)
			p.Unuary_exp3()
		}


	case MASMParserCALL:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(149)
			p.Unuary_exp4()
		}


	case MASMParserINT, MASMParserRETN, MASMParserRET, MASMParserRETF:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(150)
			p.Unuary_exp5()
		}


	case MASMParserPOPAD, MASMParserAAA, MASMParserPOPA, MASMParserPOPFD, MASMParserCWD, MASMParserLAHF, MASMParserPUSHAD, MASMParserPUSHF, MASMParserAAS, MASMParserBSWAP, MASMParserPUSHFD, MASMParserCBW, MASMParserCWDE, MASMParserXLAT, MASMParserAAM, MASMParserAAD, MASMParserCDQ, MASMParserDAA, MASMParserSAHF, MASMParserDAS, MASMParserINTO, MASMParserIRET, MASMParserCLC, MASMParserSTC, MASMParserCMC, MASMParserCLD, MASMParserSTD, MASMParserCLI, MASMParserSTI, MASMParserMOVSB, MASMParserMOVSW, MASMParserMOVSD, MASMParserLODS, MASMParserLODSB, MASMParserLODSW, MASMParserLODSD, MASMParserSTOS, MASMParserSTOSB, MASMParserSTOSW, MASMParserSOTSD, MASMParserSCAS, MASMParserSCASB, MASMParserSCASW, MASMParserSCASD, MASMParserCMPS, MASMParserCMPSB, MASMParserCMPSW, MASMParserCMPSD, MASMParserINSB, MASMParserINSW, MASMParserINSD, MASMParserOUTSB, MASMParserOUTSW, MASMParserOUTSD:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(151)
			p.Notarguments()
		}


	case MASMParserIdentifier:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(152)
			p.Variabledeclaration()
		}


	case MASMParserALPHA, MASMParserCONST, MASMParserCREF, MASMParserXCREF, MASMParserDATA, MASMParserDATA2, MASMParserDOSSEG, MASMParserERR, MASMParserERR1, MASMParserERR2, MASMParserERRE, MASMParserERRNZ, MASMParserERRDEF, MASMParserERRNDEF, MASMParserERRB, MASMParserERRNB, MASMParserERRIDN, MASMParserERRDIF, MASMParserEVEN, MASMParserLIST, MASMParserWIDTH, MASMParserMASK, MASMParserSEQ, MASMParserXLIST, MASMParserEXIT, MASMParserMODEL:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(153)
			p.Directive_exp1()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IBinary_exp1Context is an interface to support dynamic dispatch.
type IBinary_exp1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp1Context differentiates from other interfaces.
	IsBinary_exp1Context()
}

type Binary_exp1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp1Context() *Binary_exp1Context {
	var p = new(Binary_exp1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp1
	return p
}

func (*Binary_exp1Context) IsBinary_exp1Context() {}

func NewBinary_exp1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp1Context {
	var p = new(Binary_exp1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp1

	return p
}

func (s *Binary_exp1Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp1Context) O() IOContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOContext)
}

func (s *Binary_exp1Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp1Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp1Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp1Context) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *Binary_exp1Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Binary_exp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp1(s)
	}
}

func (s *Binary_exp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp1(s)
	}
}

func (s *Binary_exp1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp1(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp1() (localctx IBinary_exp1Context) {
	localctx = NewBinary_exp1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MASMParserRULE_binary_exp1)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.O()
		}
		{
			p.SetState(157)
			p.Register()
		}
		{
			p.SetState(158)
			p.Match(MASMParserSeparator)
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
			{
				p.SetState(159)
				p.Register()
			}


		case MASMParserInteger:
			{
				p.SetState(160)
				p.Match(MASMParserInteger)
			}


		case MASMParserT__6:
			{
				p.SetState(161)
				p.Memory()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.O()
		}
		{
			p.SetState(165)
			p.Memory()
		}
		{
			p.SetState(166)
			p.Match(MASMParserSeparator)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
			{
				p.SetState(167)
				p.Register()
			}


		case MASMParserInteger:
			{
				p.SetState(168)
				p.Match(MASMParserInteger)
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}


	return localctx
}


// IUnuary_exp1Context is an interface to support dynamic dispatch.
type IUnuary_exp1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnuary_exp1Context differentiates from other interfaces.
	IsUnuary_exp1Context()
}

type Unuary_exp1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnuary_exp1Context() *Unuary_exp1Context {
	var p = new(Unuary_exp1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_unuary_exp1
	return p
}

func (*Unuary_exp1Context) IsUnuary_exp1Context() {}

func NewUnuary_exp1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unuary_exp1Context {
	var p = new(Unuary_exp1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_unuary_exp1

	return p
}

func (s *Unuary_exp1Context) GetParser() antlr.Parser { return s.parser }

func (s *Unuary_exp1Context) Op() IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *Unuary_exp1Context) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *Unuary_exp1Context) Register() IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Unuary_exp1Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Unuary_exp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unuary_exp1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Unuary_exp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterUnuary_exp1(s)
	}
}

func (s *Unuary_exp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitUnuary_exp1(s)
	}
}

func (s *Unuary_exp1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitUnuary_exp1(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Unuary_exp1() (localctx IUnuary_exp1Context) {
	localctx = NewUnuary_exp1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MASMParserRULE_unuary_exp1)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Op()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserInteger:
		{
			p.SetState(174)
			p.Match(MASMParserInteger)
		}


	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(175)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(176)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IUnuary_exp2Context is an interface to support dynamic dispatch.
type IUnuary_exp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnuary_exp2Context differentiates from other interfaces.
	IsUnuary_exp2Context()
}

type Unuary_exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnuary_exp2Context() *Unuary_exp2Context {
	var p = new(Unuary_exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_unuary_exp2
	return p
}

func (*Unuary_exp2Context) IsUnuary_exp2Context() {}

func NewUnuary_exp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unuary_exp2Context {
	var p = new(Unuary_exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_unuary_exp2

	return p
}

func (s *Unuary_exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *Unuary_exp2Context) Ope() IOpeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpeContext)
}

func (s *Unuary_exp2Context) Register() IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Unuary_exp2Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Unuary_exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unuary_exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Unuary_exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterUnuary_exp2(s)
	}
}

func (s *Unuary_exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitUnuary_exp2(s)
	}
}

func (s *Unuary_exp2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitUnuary_exp2(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Unuary_exp2() (localctx IUnuary_exp2Context) {
	localctx = NewUnuary_exp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MASMParserRULE_unuary_exp2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Ope()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(180)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(181)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IBinary_exp2Context is an interface to support dynamic dispatch.
type IBinary_exp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp2Context differentiates from other interfaces.
	IsBinary_exp2Context()
}

type Binary_exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp2Context() *Binary_exp2Context {
	var p = new(Binary_exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp2
	return p
}

func (*Binary_exp2Context) IsBinary_exp2Context() {}

func NewBinary_exp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp2Context {
	var p = new(Binary_exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp2

	return p
}

func (s *Binary_exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp2Context) Oper() IOperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperContext)
}

func (s *Binary_exp2Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp2Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp2Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp2Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Binary_exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp2(s)
	}
}

func (s *Binary_exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp2(s)
	}
}

func (s *Binary_exp2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp2(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp2() (localctx IBinary_exp2Context) {
	localctx = NewBinary_exp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MASMParserRULE_binary_exp2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Oper()
		}
		{
			p.SetState(185)
			p.Register()
		}
		{
			p.SetState(186)
			p.Match(MASMParserSeparator)
		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
			{
				p.SetState(187)
				p.Register()
			}


		case MASMParserT__6:
			{
				p.SetState(188)
				p.Memory()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Oper()
		}
		{
			p.SetState(192)
			p.Memory()
		}
		{
			p.SetState(193)
			p.Match(MASMParserSeparator)
		}
		{
			p.SetState(194)
			p.Register()
		}

	}


	return localctx
}


// INotargumentsContext is an interface to support dynamic dispatch.
type INotargumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotargumentsContext differentiates from other interfaces.
	IsNotargumentsContext()
}

type NotargumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotargumentsContext() *NotargumentsContext {
	var p = new(NotargumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_notarguments
	return p
}

func (*NotargumentsContext) IsNotargumentsContext() {}

func NewNotargumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotargumentsContext {
	var p = new(NotargumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_notarguments

	return p
}

func (s *NotargumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *NotargumentsContext) Opera() IOperaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperaContext)
}

func (s *NotargumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotargumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NotargumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterNotarguments(s)
	}
}

func (s *NotargumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitNotarguments(s)
	}
}

func (s *NotargumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitNotarguments(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Notarguments() (localctx INotargumentsContext) {
	localctx = NewNotargumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MASMParserRULE_notarguments)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Opera()
	}



	return localctx
}


// IBinary_exp3Context is an interface to support dynamic dispatch.
type IBinary_exp3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp3Context differentiates from other interfaces.
	IsBinary_exp3Context()
}

type Binary_exp3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp3Context() *Binary_exp3Context {
	var p = new(Binary_exp3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp3
	return p
}

func (*Binary_exp3Context) IsBinary_exp3Context() {}

func NewBinary_exp3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp3Context {
	var p = new(Binary_exp3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp3

	return p
}

func (s *Binary_exp3Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp3Context) Operat() IOperatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatContext)
}

func (s *Binary_exp3Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp3Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp3Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp3Context) AllMemory() []IMemoryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemoryContext)(nil)).Elem())
	var tst = make([]IMemoryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemoryContext)
		}
	}

	return tst
}

func (s *Binary_exp3Context) Memory(i int) IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Binary_exp3Context) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *Binary_exp3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp3(s)
	}
}

func (s *Binary_exp3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp3(s)
	}
}

func (s *Binary_exp3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp3(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp3() (localctx IBinary_exp3Context) {
	localctx = NewBinary_exp3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MASMParserRULE_binary_exp3)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Operat()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(201)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(202)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(205)
		p.Match(MASMParserSeparator)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(206)
			p.Register()
		}


	case MASMParserInteger:
		{
			p.SetState(207)
			p.Match(MASMParserInteger)
		}


	case MASMParserT__6:
		{
			p.SetState(208)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IUnuary_exp3Context is an interface to support dynamic dispatch.
type IUnuary_exp3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnuary_exp3Context differentiates from other interfaces.
	IsUnuary_exp3Context()
}

type Unuary_exp3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnuary_exp3Context() *Unuary_exp3Context {
	var p = new(Unuary_exp3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_unuary_exp3
	return p
}

func (*Unuary_exp3Context) IsUnuary_exp3Context() {}

func NewUnuary_exp3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unuary_exp3Context {
	var p = new(Unuary_exp3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_unuary_exp3

	return p
}

func (s *Unuary_exp3Context) GetParser() antlr.Parser { return s.parser }

func (s *Unuary_exp3Context) Operato() IOperatoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatoContext)
}

func (s *Unuary_exp3Context) Identifier() antlr.TerminalNode {
	return s.GetToken(MASMParserIdentifier, 0)
}

func (s *Unuary_exp3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unuary_exp3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Unuary_exp3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterUnuary_exp3(s)
	}
}

func (s *Unuary_exp3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitUnuary_exp3(s)
	}
}

func (s *Unuary_exp3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitUnuary_exp3(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Unuary_exp3() (localctx IUnuary_exp3Context) {
	localctx = NewUnuary_exp3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MASMParserRULE_unuary_exp3)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Operato()
	}
	{
		p.SetState(212)
		p.Match(MASMParserIdentifier)
	}



	return localctx
}


// IBinary_exp4Context is an interface to support dynamic dispatch.
type IBinary_exp4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp4Context differentiates from other interfaces.
	IsBinary_exp4Context()
}

type Binary_exp4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp4Context() *Binary_exp4Context {
	var p = new(Binary_exp4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp4
	return p
}

func (*Binary_exp4Context) IsBinary_exp4Context() {}

func NewBinary_exp4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp4Context {
	var p = new(Binary_exp4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp4

	return p
}

func (s *Binary_exp4Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp4Context) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *Binary_exp4Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp4Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp4Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp4Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Binary_exp4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp4(s)
	}
}

func (s *Binary_exp4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp4(s)
	}
}

func (s *Binary_exp4Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp4(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp4() (localctx IBinary_exp4Context) {
	localctx = NewBinary_exp4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MASMParserRULE_binary_exp4)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Operator()
	}
	{
		p.SetState(215)
		p.Register()
	}
	{
		p.SetState(216)
		p.Match(MASMParserSeparator)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(217)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(218)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IBinary_exp5Context is an interface to support dynamic dispatch.
type IBinary_exp5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp5Context differentiates from other interfaces.
	IsBinary_exp5Context()
}

type Binary_exp5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp5Context() *Binary_exp5Context {
	var p = new(Binary_exp5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp5
	return p
}

func (*Binary_exp5Context) IsBinary_exp5Context() {}

func NewBinary_exp5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp5Context {
	var p = new(Binary_exp5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp5

	return p
}

func (s *Binary_exp5Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp5Context) L() ILContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILContext)
}

func (s *Binary_exp5Context) Register() IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp5Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp5Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Binary_exp5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp5(s)
	}
}

func (s *Binary_exp5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp5(s)
	}
}

func (s *Binary_exp5Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp5(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp5() (localctx IBinary_exp5Context) {
	localctx = NewBinary_exp5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MASMParserRULE_binary_exp5)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(221)
		p.L()
	}
	{
		p.SetState(222)
		p.Register()
	}
	{
		p.SetState(223)
		p.Match(MASMParserSeparator)
	}
	{
		p.SetState(224)
		p.Memory()
	}



	return localctx
}


// IBinary_exp6Context is an interface to support dynamic dispatch.
type IBinary_exp6Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp6Context differentiates from other interfaces.
	IsBinary_exp6Context()
}

type Binary_exp6Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp6Context() *Binary_exp6Context {
	var p = new(Binary_exp6Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp6
	return p
}

func (*Binary_exp6Context) IsBinary_exp6Context() {}

func NewBinary_exp6Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp6Context {
	var p = new(Binary_exp6Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp6

	return p
}

func (s *Binary_exp6Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp6Context) X() IXContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXContext)
}

func (s *Binary_exp6Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp6Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp6Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp6Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Binary_exp6Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp6Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp6Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp6(s)
	}
}

func (s *Binary_exp6Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp6(s)
	}
}

func (s *Binary_exp6Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp6(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp6() (localctx IBinary_exp6Context) {
	localctx = NewBinary_exp6Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MASMParserRULE_binary_exp6)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(226)
		p.X()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(227)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(228)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(231)
		p.Match(MASMParserSeparator)
	}
	{
		p.SetState(232)
		p.Register()
	}



	return localctx
}


// IBinary_exp7Context is an interface to support dynamic dispatch.
type IBinary_exp7Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp7Context differentiates from other interfaces.
	IsBinary_exp7Context()
}

type Binary_exp7Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp7Context() *Binary_exp7Context {
	var p = new(Binary_exp7Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp7
	return p
}

func (*Binary_exp7Context) IsBinary_exp7Context() {}

func NewBinary_exp7Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp7Context {
	var p = new(Binary_exp7Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp7

	return p
}

func (s *Binary_exp7Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp7Context) S() ISContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISContext)
}

func (s *Binary_exp7Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp7Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp7Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp7Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Binary_exp7Context) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *Binary_exp7Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp7Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp7Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp7(s)
	}
}

func (s *Binary_exp7Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp7(s)
	}
}

func (s *Binary_exp7Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp7(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp7() (localctx IBinary_exp7Context) {
	localctx = NewBinary_exp7Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MASMParserRULE_binary_exp7)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(234)
		p.S()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(235)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(236)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(239)
		p.Match(MASMParserSeparator)
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserInteger:
		{
			p.SetState(240)
			p.Match(MASMParserInteger)
		}


	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(241)
			p.Register()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IBinary_exp8Context is an interface to support dynamic dispatch.
type IBinary_exp8Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp8Context differentiates from other interfaces.
	IsBinary_exp8Context()
}

type Binary_exp8Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp8Context() *Binary_exp8Context {
	var p = new(Binary_exp8Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp8
	return p
}

func (*Binary_exp8Context) IsBinary_exp8Context() {}

func NewBinary_exp8Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp8Context {
	var p = new(Binary_exp8Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp8

	return p
}

func (s *Binary_exp8Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp8Context) Sh() IShContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShContext)
}

func (s *Binary_exp8Context) AllSeparator() []antlr.TerminalNode {
	return s.GetTokens(MASMParserSeparator)
}

func (s *Binary_exp8Context) Separator(i int) antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, i)
}

func (s *Binary_exp8Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp8Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp8Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Binary_exp8Context) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *Binary_exp8Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp8Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp8(s)
	}
}

func (s *Binary_exp8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp8(s)
	}
}

func (s *Binary_exp8Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp8(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp8() (localctx IBinary_exp8Context) {
	localctx = NewBinary_exp8Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MASMParserRULE_binary_exp8)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Sh()
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(245)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(246)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(249)
		p.Match(MASMParserSeparator)
	}
	{
		p.SetState(250)
		p.Register()
	}
	{
		p.SetState(251)
		p.Match(MASMParserSeparator)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(252)
			p.Register()
		}


	case MASMParserInteger:
		{
			p.SetState(253)
			p.Match(MASMParserInteger)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IBinary_exp9Context is an interface to support dynamic dispatch.
type IBinary_exp9Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp9Context differentiates from other interfaces.
	IsBinary_exp9Context()
}

type Binary_exp9Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp9Context() *Binary_exp9Context {
	var p = new(Binary_exp9Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp9
	return p
}

func (*Binary_exp9Context) IsBinary_exp9Context() {}

func NewBinary_exp9Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp9Context {
	var p = new(Binary_exp9Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp9

	return p
}

func (s *Binary_exp9Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp9Context) B() IBContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBContext)
}

func (s *Binary_exp9Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp9Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp9Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp9Context) AllMemory() []IMemoryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemoryContext)(nil)).Elem())
	var tst = make([]IMemoryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemoryContext)
		}
	}

	return tst
}

func (s *Binary_exp9Context) Memory(i int) IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Binary_exp9Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp9Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp9Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp9(s)
	}
}

func (s *Binary_exp9Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp9(s)
	}
}

func (s *Binary_exp9Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp9(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp9() (localctx IBinary_exp9Context) {
	localctx = NewBinary_exp9Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MASMParserRULE_binary_exp9)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.B()
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(257)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(258)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(261)
		p.Match(MASMParserSeparator)
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(262)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(263)
			p.Memory()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IUnuary_exp4Context is an interface to support dynamic dispatch.
type IUnuary_exp4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnuary_exp4Context differentiates from other interfaces.
	IsUnuary_exp4Context()
}

type Unuary_exp4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnuary_exp4Context() *Unuary_exp4Context {
	var p = new(Unuary_exp4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_unuary_exp4
	return p
}

func (*Unuary_exp4Context) IsUnuary_exp4Context() {}

func NewUnuary_exp4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unuary_exp4Context {
	var p = new(Unuary_exp4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_unuary_exp4

	return p
}

func (s *Unuary_exp4Context) GetParser() antlr.Parser { return s.parser }

func (s *Unuary_exp4Context) Call() ICallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *Unuary_exp4Context) Register() IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Unuary_exp4Context) Memory() IMemoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *Unuary_exp4Context) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *Unuary_exp4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unuary_exp4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Unuary_exp4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterUnuary_exp4(s)
	}
}

func (s *Unuary_exp4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitUnuary_exp4(s)
	}
}

func (s *Unuary_exp4Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitUnuary_exp4(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Unuary_exp4() (localctx IUnuary_exp4Context) {
	localctx = NewUnuary_exp4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MASMParserRULE_unuary_exp4)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Call()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(267)
			p.Register()
		}


	case MASMParserT__6:
		{
			p.SetState(268)
			p.Memory()
		}


	case MASMParserInteger:
		{
			p.SetState(269)
			p.Match(MASMParserInteger)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IUnuary_exp5Context is an interface to support dynamic dispatch.
type IUnuary_exp5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnuary_exp5Context differentiates from other interfaces.
	IsUnuary_exp5Context()
}

type Unuary_exp5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnuary_exp5Context() *Unuary_exp5Context {
	var p = new(Unuary_exp5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_unuary_exp5
	return p
}

func (*Unuary_exp5Context) IsUnuary_exp5Context() {}

func NewUnuary_exp5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unuary_exp5Context {
	var p = new(Unuary_exp5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_unuary_exp5

	return p
}

func (s *Unuary_exp5Context) GetParser() antlr.Parser { return s.parser }

func (s *Unuary_exp5Context) Interruption() IInterruptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterruptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterruptionContext)
}

func (s *Unuary_exp5Context) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *Unuary_exp5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unuary_exp5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Unuary_exp5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterUnuary_exp5(s)
	}
}

func (s *Unuary_exp5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitUnuary_exp5(s)
	}
}

func (s *Unuary_exp5Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitUnuary_exp5(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Unuary_exp5() (localctx IUnuary_exp5Context) {
	localctx = NewUnuary_exp5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MASMParserRULE_unuary_exp5)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(272)
		p.Interruption()
	}
	{
		p.SetState(273)
		p.Match(MASMParserInteger)
	}



	return localctx
}


// IBinary_exp10Context is an interface to support dynamic dispatch.
type IBinary_exp10Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp10Context differentiates from other interfaces.
	IsBinary_exp10Context()
}

type Binary_exp10Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp10Context() *Binary_exp10Context {
	var p = new(Binary_exp10Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp10
	return p
}

func (*Binary_exp10Context) IsBinary_exp10Context() {}

func NewBinary_exp10Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp10Context {
	var p = new(Binary_exp10Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp10

	return p
}

func (s *Binary_exp10Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp10Context) In() IInContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInContext)
}

func (s *Binary_exp10Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp10Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp10Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp10Context) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *Binary_exp10Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp10Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp10Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp10(s)
	}
}

func (s *Binary_exp10Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp10(s)
	}
}

func (s *Binary_exp10Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp10(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp10() (localctx IBinary_exp10Context) {
	localctx = NewBinary_exp10Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MASMParserRULE_binary_exp10)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.In()
	}
	{
		p.SetState(276)
		p.Register()
	}
	{
		p.SetState(277)
		p.Match(MASMParserSeparator)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(278)
			p.Register()
		}


	case MASMParserInteger:
		{
			p.SetState(279)
			p.Match(MASMParserInteger)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IBinary_exp11Context is an interface to support dynamic dispatch.
type IBinary_exp11Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp11Context differentiates from other interfaces.
	IsBinary_exp11Context()
}

type Binary_exp11Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp11Context() *Binary_exp11Context {
	var p = new(Binary_exp11Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp11
	return p
}

func (*Binary_exp11Context) IsBinary_exp11Context() {}

func NewBinary_exp11Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp11Context {
	var p = new(Binary_exp11Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp11

	return p
}

func (s *Binary_exp11Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp11Context) Out() IOutContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOutContext)
}

func (s *Binary_exp11Context) Separator() antlr.TerminalNode {
	return s.GetToken(MASMParserSeparator, 0)
}

func (s *Binary_exp11Context) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *Binary_exp11Context) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *Binary_exp11Context) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *Binary_exp11Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp11Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp11Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp11(s)
	}
}

func (s *Binary_exp11Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp11(s)
	}
}

func (s *Binary_exp11Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp11(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp11() (localctx IBinary_exp11Context) {
	localctx = NewBinary_exp11Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MASMParserRULE_binary_exp11)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Out()
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(283)
			p.Register()
		}


	case MASMParserInteger:
		{
			p.SetState(284)
			p.Match(MASMParserInteger)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(287)
		p.Match(MASMParserSeparator)
	}
	{
		p.SetState(288)
		p.Register()
	}



	return localctx
}


// IBinary_exp12Context is an interface to support dynamic dispatch.
type IBinary_exp12Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_exp12Context differentiates from other interfaces.
	IsBinary_exp12Context()
}

type Binary_exp12Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_exp12Context() *Binary_exp12Context {
	var p = new(Binary_exp12Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_binary_exp12
	return p
}

func (*Binary_exp12Context) IsBinary_exp12Context() {}

func NewBinary_exp12Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exp12Context {
	var p = new(Binary_exp12Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_binary_exp12

	return p
}

func (s *Binary_exp12Context) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exp12Context) Re() IReContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReContext)
}

func (s *Binary_exp12Context) Opera() IOperaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperaContext)
}

func (s *Binary_exp12Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exp12Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Binary_exp12Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterBinary_exp12(s)
	}
}

func (s *Binary_exp12Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitBinary_exp12(s)
	}
}

func (s *Binary_exp12Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitBinary_exp12(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Binary_exp12() (localctx IBinary_exp12Context) {
	localctx = NewBinary_exp12Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MASMParserRULE_binary_exp12)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Re()
	}
	{
		p.SetState(291)
		p.Opera()
	}



	return localctx
}


// IDirective_exp1Context is an interface to support dynamic dispatch.
type IDirective_exp1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirective_exp1Context differentiates from other interfaces.
	IsDirective_exp1Context()
}

type Directive_exp1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirective_exp1Context() *Directive_exp1Context {
	var p = new(Directive_exp1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_directive_exp1
	return p
}

func (*Directive_exp1Context) IsDirective_exp1Context() {}

func NewDirective_exp1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Directive_exp1Context {
	var p = new(Directive_exp1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_directive_exp1

	return p
}

func (s *Directive_exp1Context) GetParser() antlr.Parser { return s.parser }

func (s *Directive_exp1Context) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *Directive_exp1Context) Identifier() antlr.TerminalNode {
	return s.GetToken(MASMParserIdentifier, 0)
}

func (s *Directive_exp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Directive_exp1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Directive_exp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterDirective_exp1(s)
	}
}

func (s *Directive_exp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitDirective_exp1(s)
	}
}

func (s *Directive_exp1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitDirective_exp1(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Directive_exp1() (localctx IDirective_exp1Context) {
	localctx = NewDirective_exp1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MASMParserRULE_directive_exp1)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(293)
			p.Directives()
		}
		{
			p.SetState(294)
			p.Match(MASMParserIdentifier)
		}


	case 2:
		{
			p.SetState(296)
			p.Directives()
		}

	}



	return localctx
}


// IVariabledeclarationContext is an interface to support dynamic dispatch.
type IVariabledeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariabledeclarationContext differentiates from other interfaces.
	IsVariabledeclarationContext()
}

type VariabledeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariabledeclarationContext() *VariabledeclarationContext {
	var p = new(VariabledeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_variabledeclaration
	return p
}

func (*VariabledeclarationContext) IsVariabledeclarationContext() {}

func NewVariabledeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariabledeclarationContext {
	var p = new(VariabledeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_variabledeclaration

	return p
}

func (s *VariabledeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariabledeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MASMParserIdentifier, 0)
}

func (s *VariabledeclarationContext) Ty() ITyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyContext)
}

func (s *VariabledeclarationContext) Question() IQuestionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuestionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuestionContext)
}

func (s *VariabledeclarationContext) String() antlr.TerminalNode {
	return s.GetToken(MASMParserString, 0)
}

func (s *VariabledeclarationContext) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *VariabledeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariabledeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VariabledeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterVariabledeclaration(s)
	}
}

func (s *VariabledeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitVariabledeclaration(s)
	}
}

func (s *VariabledeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitVariabledeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Variabledeclaration() (localctx IVariabledeclarationContext) {
	localctx = NewVariabledeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MASMParserRULE_variabledeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(MASMParserIdentifier)
	}
	{
		p.SetState(300)
		p.Ty()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserQUESTION:
		{
			p.SetState(301)
			p.Question()
		}


	case MASMParserString:
		{
			p.SetState(302)
			p.Match(MASMParserString)
		}


	case MASMParserInteger:
		{
			p.SetState(303)
			p.Match(MASMParserInteger)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IMemoryContext is an interface to support dynamic dispatch.
type IMemoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemoryContext differentiates from other interfaces.
	IsMemoryContext()
}

type MemoryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemoryContext() *MemoryContext {
	var p = new(MemoryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_memory
	return p
}

func (*MemoryContext) IsMemoryContext() {}

func NewMemoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemoryContext {
	var p = new(MemoryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_memory

	return p
}

func (s *MemoryContext) GetParser() antlr.Parser { return s.parser }

func (s *MemoryContext) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *MemoryContext) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *MemoryContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MASMParserIdentifier, 0)
}

func (s *MemoryContext) Integer() antlr.TerminalNode {
	return s.GetToken(MASMParserInteger, 0)
}

func (s *MemoryContext) Hexnum() antlr.TerminalNode {
	return s.GetToken(MASMParserHexnum, 0)
}

func (s *MemoryContext) Octalnum() antlr.TerminalNode {
	return s.GetToken(MASMParserOctalnum, 0)
}

func (s *MemoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MemoryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterMemory(s)
	}
}

func (s *MemoryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitMemory(s)
	}
}

func (s *MemoryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitMemory(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Memory() (localctx IMemoryContext) {
	localctx = NewMemoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MASMParserRULE_memory)
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
		p.Match(MASMParserT__6)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
		{
			p.SetState(307)
			p.Register()
		}


	case MASMParserIdentifier:
		{
			p.SetState(308)
			p.Match(MASMParserIdentifier)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MASMParserT__7 {
		{
			p.SetState(311)
			p.Match(MASMParserT__7)
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MASMParserAH, MASMParserAL, MASMParserAX, MASMParserBH, MASMParserBL, MASMParserBX, MASMParserCH, MASMParserCL, MASMParserCX, MASMParserDH, MASMParserDL, MASMParserDX, MASMParserSI, MASMParserDI, MASMParserSP, MASMParserBP, MASMParserEAX, MASMParserEBX, MASMParserECX, MASMParserEDX, MASMParserESI, MASMParserEDI, MASMParserESP, MASMParserEBP:
			{
				p.SetState(312)
				p.Register()
			}
			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			if _la == MASMParserT__7 {
				{
					p.SetState(313)
					p.Match(MASMParserT__7)
				}
				p.SetState(314)
				_la = p.GetTokenStream().LA(1)

				if !(((((_la - 266)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 266))) & ((1 << (MASMParserHexnum - 266)) | (1 << (MASMParserInteger - 266)) | (1 << (MASMParserOctalnum - 266)))) != 0)) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
				    p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}

			}



		case MASMParserInteger:
			{
				p.SetState(317)
				p.Match(MASMParserInteger)
			}


		case MASMParserHexnum:
			{
				p.SetState(318)
				p.Match(MASMParserHexnum)
			}


		case MASMParserOctalnum:
			{
				p.SetState(319)
				p.Match(MASMParserOctalnum)
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}
	{
		p.SetState(324)
		p.Match(MASMParserT__8)
	}



	return localctx
}


// ISegmentosContext is an interface to support dynamic dispatch.
type ISegmentosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentosContext differentiates from other interfaces.
	IsSegmentosContext()
}

type SegmentosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentosContext() *SegmentosContext {
	var p = new(SegmentosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_segmentos
	return p
}

func (*SegmentosContext) IsSegmentosContext() {}

func NewSegmentosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentosContext {
	var p = new(SegmentosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_segmentos

	return p
}

func (s *SegmentosContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentosContext) DS() antlr.TerminalNode {
	return s.GetToken(MASMParserDS, 0)
}

func (s *SegmentosContext) ES() antlr.TerminalNode {
	return s.GetToken(MASMParserES, 0)
}

func (s *SegmentosContext) CS() antlr.TerminalNode {
	return s.GetToken(MASMParserCS, 0)
}

func (s *SegmentosContext) SS() antlr.TerminalNode {
	return s.GetToken(MASMParserSS, 0)
}

func (s *SegmentosContext) GS() antlr.TerminalNode {
	return s.GetToken(MASMParserGS, 0)
}

func (s *SegmentosContext) FS() antlr.TerminalNode {
	return s.GetToken(MASMParserFS, 0)
}

func (s *SegmentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SegmentosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterSegmentos(s)
	}
}

func (s *SegmentosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitSegmentos(s)
	}
}

func (s *SegmentosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitSegmentos(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Segmentos() (localctx ISegmentosContext) {
	localctx = NewSegmentosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MASMParserRULE_segmentos)
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
	p.SetState(326)
	_la = p.GetTokenStream().LA(1)

	if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MASMParserDS) | (1 << MASMParserES) | (1 << MASMParserCS) | (1 << MASMParserSS) | (1 << MASMParserGS) | (1 << MASMParserFS))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IRegisterContext is an interface to support dynamic dispatch.
type IRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegisterContext differentiates from other interfaces.
	IsRegisterContext()
}

type RegisterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegisterContext() *RegisterContext {
	var p = new(RegisterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_register
	return p
}

func (*RegisterContext) IsRegisterContext() {}

func NewRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterContext {
	var p = new(RegisterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_register

	return p
}

func (s *RegisterContext) GetParser() antlr.Parser { return s.parser }

func (s *RegisterContext) AH() antlr.TerminalNode {
	return s.GetToken(MASMParserAH, 0)
}

func (s *RegisterContext) AL() antlr.TerminalNode {
	return s.GetToken(MASMParserAL, 0)
}

func (s *RegisterContext) AX() antlr.TerminalNode {
	return s.GetToken(MASMParserAX, 0)
}

func (s *RegisterContext) BH() antlr.TerminalNode {
	return s.GetToken(MASMParserBH, 0)
}

func (s *RegisterContext) BL() antlr.TerminalNode {
	return s.GetToken(MASMParserBL, 0)
}

func (s *RegisterContext) BX() antlr.TerminalNode {
	return s.GetToken(MASMParserBX, 0)
}

func (s *RegisterContext) CH() antlr.TerminalNode {
	return s.GetToken(MASMParserCH, 0)
}

func (s *RegisterContext) CL() antlr.TerminalNode {
	return s.GetToken(MASMParserCL, 0)
}

func (s *RegisterContext) CX() antlr.TerminalNode {
	return s.GetToken(MASMParserCX, 0)
}

func (s *RegisterContext) DH() antlr.TerminalNode {
	return s.GetToken(MASMParserDH, 0)
}

func (s *RegisterContext) DL() antlr.TerminalNode {
	return s.GetToken(MASMParserDL, 0)
}

func (s *RegisterContext) DX() antlr.TerminalNode {
	return s.GetToken(MASMParserDX, 0)
}

func (s *RegisterContext) SI() antlr.TerminalNode {
	return s.GetToken(MASMParserSI, 0)
}

func (s *RegisterContext) DI() antlr.TerminalNode {
	return s.GetToken(MASMParserDI, 0)
}

func (s *RegisterContext) SP() antlr.TerminalNode {
	return s.GetToken(MASMParserSP, 0)
}

func (s *RegisterContext) BP() antlr.TerminalNode {
	return s.GetToken(MASMParserBP, 0)
}

func (s *RegisterContext) EAX() antlr.TerminalNode {
	return s.GetToken(MASMParserEAX, 0)
}

func (s *RegisterContext) EBX() antlr.TerminalNode {
	return s.GetToken(MASMParserEBX, 0)
}

func (s *RegisterContext) ECX() antlr.TerminalNode {
	return s.GetToken(MASMParserECX, 0)
}

func (s *RegisterContext) EDX() antlr.TerminalNode {
	return s.GetToken(MASMParserEDX, 0)
}

func (s *RegisterContext) ESI() antlr.TerminalNode {
	return s.GetToken(MASMParserESI, 0)
}

func (s *RegisterContext) EDI() antlr.TerminalNode {
	return s.GetToken(MASMParserEDI, 0)
}

func (s *RegisterContext) ESP() antlr.TerminalNode {
	return s.GetToken(MASMParserESP, 0)
}

func (s *RegisterContext) EBP() antlr.TerminalNode {
	return s.GetToken(MASMParserEBP, 0)
}

func (s *RegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterRegister(s)
	}
}

func (s *RegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitRegister(s)
	}
}

func (s *RegisterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitRegister(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Register() (localctx IRegisterContext) {
	localctx = NewRegisterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MASMParserRULE_register)
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
	p.SetState(328)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 17)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 17))) & ((1 << (MASMParserAH - 17)) | (1 << (MASMParserAL - 17)) | (1 << (MASMParserAX - 17)) | (1 << (MASMParserBH - 17)) | (1 << (MASMParserBL - 17)) | (1 << (MASMParserBX - 17)) | (1 << (MASMParserCH - 17)) | (1 << (MASMParserCL - 17)) | (1 << (MASMParserCX - 17)) | (1 << (MASMParserDH - 17)) | (1 << (MASMParserDL - 17)) | (1 << (MASMParserDX - 17)) | (1 << (MASMParserSI - 17)) | (1 << (MASMParserDI - 17)) | (1 << (MASMParserSP - 17)) | (1 << (MASMParserBP - 17)) | (1 << (MASMParserEAX - 17)) | (1 << (MASMParserEBX - 17)) | (1 << (MASMParserECX - 17)) | (1 << (MASMParserEDX - 17)) | (1 << (MASMParserESI - 17)) | (1 << (MASMParserEDI - 17)) | (1 << (MASMParserESP - 17)) | (1 << (MASMParserEBP - 17)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IOContext is an interface to support dynamic dispatch.
type IOContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOContext differentiates from other interfaces.
	IsOContext()
}

type OContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOContext() *OContext {
	var p = new(OContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_o
	return p
}

func (*OContext) IsOContext() {}

func NewOContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OContext {
	var p = new(OContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_o

	return p
}

func (s *OContext) GetParser() antlr.Parser { return s.parser }

func (s *OContext) MOV() antlr.TerminalNode {
	return s.GetToken(MASMParserMOV, 0)
}

func (s *OContext) CMP() antlr.TerminalNode {
	return s.GetToken(MASMParserCMP, 0)
}

func (s *OContext) TEST() antlr.TerminalNode {
	return s.GetToken(MASMParserTEST, 0)
}

func (s *OContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterO(s)
	}
}

func (s *OContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitO(s)
	}
}

func (s *OContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitO(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) O() (localctx IOContext) {
	localctx = NewOContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MASMParserRULE_o)
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
	p.SetState(330)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 41)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 41))) & ((1 << (MASMParserMOV - 41)) | (1 << (MASMParserCMP - 41)) | (1 << (MASMParserTEST - 41)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) PUSH() antlr.TerminalNode {
	return s.GetToken(MASMParserPUSH, 0)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitOp(s)
	}
}

func (s *OpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitOp(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MASMParserRULE_op)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(332)
		p.Match(MASMParserPUSH)
	}



	return localctx
}


// IOpeContext is an interface to support dynamic dispatch.
type IOpeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpeContext differentiates from other interfaces.
	IsOpeContext()
}

type OpeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpeContext() *OpeContext {
	var p = new(OpeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_ope
	return p
}

func (*OpeContext) IsOpeContext() {}

func NewOpeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpeContext {
	var p = new(OpeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_ope

	return p
}

func (s *OpeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpeContext) POP() antlr.TerminalNode {
	return s.GetToken(MASMParserPOP, 0)
}

func (s *OpeContext) IDIV() antlr.TerminalNode {
	return s.GetToken(MASMParserIDIV, 0)
}

func (s *OpeContext) INC() antlr.TerminalNode {
	return s.GetToken(MASMParserINC, 0)
}

func (s *OpeContext) DEC() antlr.TerminalNode {
	return s.GetToken(MASMParserDEC, 0)
}

func (s *OpeContext) NEG() antlr.TerminalNode {
	return s.GetToken(MASMParserNEG, 0)
}

func (s *OpeContext) MUL() antlr.TerminalNode {
	return s.GetToken(MASMParserMUL, 0)
}

func (s *OpeContext) DIV() antlr.TerminalNode {
	return s.GetToken(MASMParserDIV, 0)
}

func (s *OpeContext) IMUL() antlr.TerminalNode {
	return s.GetToken(MASMParserIMUL, 0)
}

func (s *OpeContext) NOT() antlr.TerminalNode {
	return s.GetToken(MASMParserNOT, 0)
}

func (s *OpeContext) SETPO() antlr.TerminalNode {
	return s.GetToken(MASMParserSETPO, 0)
}

func (s *OpeContext) SETAE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETAE, 0)
}

func (s *OpeContext) SETNLE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNLE, 0)
}

func (s *OpeContext) SETC() antlr.TerminalNode {
	return s.GetToken(MASMParserSETC, 0)
}

func (s *OpeContext) SETNO() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNO, 0)
}

func (s *OpeContext) SETNB() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNB, 0)
}

func (s *OpeContext) SETP() antlr.TerminalNode {
	return s.GetToken(MASMParserSETP, 0)
}

func (s *OpeContext) SETNGE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNGE, 0)
}

func (s *OpeContext) SETL() antlr.TerminalNode {
	return s.GetToken(MASMParserSETL, 0)
}

func (s *OpeContext) SETGE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETGE, 0)
}

func (s *OpeContext) SETPE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETPE, 0)
}

func (s *OpeContext) SETNL() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNL, 0)
}

func (s *OpeContext) SETNZ() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNZ, 0)
}

func (s *OpeContext) SETNE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNE, 0)
}

func (s *OpeContext) SETNC() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNC, 0)
}

func (s *OpeContext) SETBE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETBE, 0)
}

func (s *OpeContext) SETNP() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNP, 0)
}

func (s *OpeContext) SETNS() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNS, 0)
}

func (s *OpeContext) SETO() antlr.TerminalNode {
	return s.GetToken(MASMParserSETO, 0)
}

func (s *OpeContext) SETNA() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNA, 0)
}

func (s *OpeContext) SETNAE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNAE, 0)
}

func (s *OpeContext) SETZ() antlr.TerminalNode {
	return s.GetToken(MASMParserSETZ, 0)
}

func (s *OpeContext) SETLE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETLE, 0)
}

func (s *OpeContext) SETNBE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNBE, 0)
}

func (s *OpeContext) SETS() antlr.TerminalNode {
	return s.GetToken(MASMParserSETS, 0)
}

func (s *OpeContext) SETE() antlr.TerminalNode {
	return s.GetToken(MASMParserSETE, 0)
}

func (s *OpeContext) SETB() antlr.TerminalNode {
	return s.GetToken(MASMParserSETB, 0)
}

func (s *OpeContext) SETA() antlr.TerminalNode {
	return s.GetToken(MASMParserSETA, 0)
}

func (s *OpeContext) SETG() antlr.TerminalNode {
	return s.GetToken(MASMParserSETG, 0)
}

func (s *OpeContext) SETNG() antlr.TerminalNode {
	return s.GetToken(MASMParserSETNG, 0)
}

func (s *OpeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OpeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterOpe(s)
	}
}

func (s *OpeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitOpe(s)
	}
}

func (s *OpeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitOpe(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Ope() (localctx IOpeContext) {
	localctx = NewOpeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MASMParserRULE_ope)
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
	p.SetState(334)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 45)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 45))) & ((1 << (MASMParserPOP - 45)) | (1 << (MASMParserIDIV - 45)) | (1 << (MASMParserINC - 45)) | (1 << (MASMParserDEC - 45)) | (1 << (MASMParserNEG - 45)) | (1 << (MASMParserMUL - 45)) | (1 << (MASMParserDIV - 45)) | (1 << (MASMParserIMUL - 45)) | (1 << (MASMParserNOT - 45)) | (1 << (MASMParserSETPO - 45)) | (1 << (MASMParserSETAE - 45)) | (1 << (MASMParserSETNLE - 45)) | (1 << (MASMParserSETC - 45)) | (1 << (MASMParserSETNO - 45)) | (1 << (MASMParserSETNB - 45)) | (1 << (MASMParserSETP - 45)) | (1 << (MASMParserSETNGE - 45)) | (1 << (MASMParserSETL - 45)) | (1 << (MASMParserSETGE - 45)) | (1 << (MASMParserSETPE - 45)) | (1 << (MASMParserSETNL - 45)) | (1 << (MASMParserSETNZ - 45)) | (1 << (MASMParserSETNE - 45)) | (1 << (MASMParserSETNC - 45)) | (1 << (MASMParserSETBE - 45)) | (1 << (MASMParserSETNP - 45)) | (1 << (MASMParserSETNS - 45)) | (1 << (MASMParserSETO - 45)) | (1 << (MASMParserSETNA - 45)) | (1 << (MASMParserSETNAE - 45)) | (1 << (MASMParserSETZ - 45)) | (1 << (MASMParserSETLE - 45)))) != 0) || ((((_la - 77)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 77))) & ((1 << (MASMParserSETNBE - 77)) | (1 << (MASMParserSETS - 77)) | (1 << (MASMParserSETE - 77)) | (1 << (MASMParserSETB - 77)) | (1 << (MASMParserSETA - 77)) | (1 << (MASMParserSETG - 77)) | (1 << (MASMParserSETNG - 77)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IOperContext is an interface to support dynamic dispatch.
type IOperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperContext differentiates from other interfaces.
	IsOperContext()
}

type OperContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperContext() *OperContext {
	var p = new(OperContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_oper
	return p
}

func (*OperContext) IsOperContext() {}

func NewOperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperContext {
	var p = new(OperContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_oper

	return p
}

func (s *OperContext) GetParser() antlr.Parser { return s.parser }

func (s *OperContext) XCHG() antlr.TerminalNode {
	return s.GetToken(MASMParserXCHG, 0)
}

func (s *OperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterOper(s)
	}
}

func (s *OperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitOper(s)
	}
}

func (s *OperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitOper(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Oper() (localctx IOperContext) {
	localctx = NewOperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MASMParserRULE_oper)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(336)
		p.Match(MASMParserXCHG)
	}



	return localctx
}


// IOperaContext is an interface to support dynamic dispatch.
type IOperaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperaContext differentiates from other interfaces.
	IsOperaContext()
}

type OperaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperaContext() *OperaContext {
	var p = new(OperaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_opera
	return p
}

func (*OperaContext) IsOperaContext() {}

func NewOperaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperaContext {
	var p = new(OperaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_opera

	return p
}

func (s *OperaContext) GetParser() antlr.Parser { return s.parser }

func (s *OperaContext) POPAD() antlr.TerminalNode {
	return s.GetToken(MASMParserPOPAD, 0)
}

func (s *OperaContext) AAA() antlr.TerminalNode {
	return s.GetToken(MASMParserAAA, 0)
}

func (s *OperaContext) POPA() antlr.TerminalNode {
	return s.GetToken(MASMParserPOPA, 0)
}

func (s *OperaContext) POPFD() antlr.TerminalNode {
	return s.GetToken(MASMParserPOPFD, 0)
}

func (s *OperaContext) CWD() antlr.TerminalNode {
	return s.GetToken(MASMParserCWD, 0)
}

func (s *OperaContext) LAHF() antlr.TerminalNode {
	return s.GetToken(MASMParserLAHF, 0)
}

func (s *OperaContext) PUSHAD() antlr.TerminalNode {
	return s.GetToken(MASMParserPUSHAD, 0)
}

func (s *OperaContext) PUSHF() antlr.TerminalNode {
	return s.GetToken(MASMParserPUSHF, 0)
}

func (s *OperaContext) AAS() antlr.TerminalNode {
	return s.GetToken(MASMParserAAS, 0)
}

func (s *OperaContext) BSWAP() antlr.TerminalNode {
	return s.GetToken(MASMParserBSWAP, 0)
}

func (s *OperaContext) PUSHFD() antlr.TerminalNode {
	return s.GetToken(MASMParserPUSHFD, 0)
}

func (s *OperaContext) CBW() antlr.TerminalNode {
	return s.GetToken(MASMParserCBW, 0)
}

func (s *OperaContext) CWDE() antlr.TerminalNode {
	return s.GetToken(MASMParserCWDE, 0)
}

func (s *OperaContext) XLAT() antlr.TerminalNode {
	return s.GetToken(MASMParserXLAT, 0)
}

func (s *OperaContext) AAM() antlr.TerminalNode {
	return s.GetToken(MASMParserAAM, 0)
}

func (s *OperaContext) AAD() antlr.TerminalNode {
	return s.GetToken(MASMParserAAD, 0)
}

func (s *OperaContext) CDQ() antlr.TerminalNode {
	return s.GetToken(MASMParserCDQ, 0)
}

func (s *OperaContext) DAA() antlr.TerminalNode {
	return s.GetToken(MASMParserDAA, 0)
}

func (s *OperaContext) SAHF() antlr.TerminalNode {
	return s.GetToken(MASMParserSAHF, 0)
}

func (s *OperaContext) DAS() antlr.TerminalNode {
	return s.GetToken(MASMParserDAS, 0)
}

func (s *OperaContext) INTO() antlr.TerminalNode {
	return s.GetToken(MASMParserINTO, 0)
}

func (s *OperaContext) IRET() antlr.TerminalNode {
	return s.GetToken(MASMParserIRET, 0)
}

func (s *OperaContext) CLC() antlr.TerminalNode {
	return s.GetToken(MASMParserCLC, 0)
}

func (s *OperaContext) STC() antlr.TerminalNode {
	return s.GetToken(MASMParserSTC, 0)
}

func (s *OperaContext) CMC() antlr.TerminalNode {
	return s.GetToken(MASMParserCMC, 0)
}

func (s *OperaContext) CLD() antlr.TerminalNode {
	return s.GetToken(MASMParserCLD, 0)
}

func (s *OperaContext) STD() antlr.TerminalNode {
	return s.GetToken(MASMParserSTD, 0)
}

func (s *OperaContext) CLI() antlr.TerminalNode {
	return s.GetToken(MASMParserCLI, 0)
}

func (s *OperaContext) STI() antlr.TerminalNode {
	return s.GetToken(MASMParserSTI, 0)
}

func (s *OperaContext) MOVSB() antlr.TerminalNode {
	return s.GetToken(MASMParserMOVSB, 0)
}

func (s *OperaContext) MOVSW() antlr.TerminalNode {
	return s.GetToken(MASMParserMOVSW, 0)
}

func (s *OperaContext) MOVSD() antlr.TerminalNode {
	return s.GetToken(MASMParserMOVSD, 0)
}

func (s *OperaContext) LODS() antlr.TerminalNode {
	return s.GetToken(MASMParserLODS, 0)
}

func (s *OperaContext) LODSB() antlr.TerminalNode {
	return s.GetToken(MASMParserLODSB, 0)
}

func (s *OperaContext) LODSW() antlr.TerminalNode {
	return s.GetToken(MASMParserLODSW, 0)
}

func (s *OperaContext) LODSD() antlr.TerminalNode {
	return s.GetToken(MASMParserLODSD, 0)
}

func (s *OperaContext) STOS() antlr.TerminalNode {
	return s.GetToken(MASMParserSTOS, 0)
}

func (s *OperaContext) STOSB() antlr.TerminalNode {
	return s.GetToken(MASMParserSTOSB, 0)
}

func (s *OperaContext) STOSW() antlr.TerminalNode {
	return s.GetToken(MASMParserSTOSW, 0)
}

func (s *OperaContext) SOTSD() antlr.TerminalNode {
	return s.GetToken(MASMParserSOTSD, 0)
}

func (s *OperaContext) SCAS() antlr.TerminalNode {
	return s.GetToken(MASMParserSCAS, 0)
}

func (s *OperaContext) SCASB() antlr.TerminalNode {
	return s.GetToken(MASMParserSCASB, 0)
}

func (s *OperaContext) SCASW() antlr.TerminalNode {
	return s.GetToken(MASMParserSCASW, 0)
}

func (s *OperaContext) SCASD() antlr.TerminalNode {
	return s.GetToken(MASMParserSCASD, 0)
}

func (s *OperaContext) CMPS() antlr.TerminalNode {
	return s.GetToken(MASMParserCMPS, 0)
}

func (s *OperaContext) CMPSB() antlr.TerminalNode {
	return s.GetToken(MASMParserCMPSB, 0)
}

func (s *OperaContext) CMPSW() antlr.TerminalNode {
	return s.GetToken(MASMParserCMPSW, 0)
}

func (s *OperaContext) CMPSD() antlr.TerminalNode {
	return s.GetToken(MASMParserCMPSD, 0)
}

func (s *OperaContext) INSB() antlr.TerminalNode {
	return s.GetToken(MASMParserINSB, 0)
}

func (s *OperaContext) INSW() antlr.TerminalNode {
	return s.GetToken(MASMParserINSW, 0)
}

func (s *OperaContext) INSD() antlr.TerminalNode {
	return s.GetToken(MASMParserINSD, 0)
}

func (s *OperaContext) OUTSB() antlr.TerminalNode {
	return s.GetToken(MASMParserOUTSB, 0)
}

func (s *OperaContext) OUTSW() antlr.TerminalNode {
	return s.GetToken(MASMParserOUTSW, 0)
}

func (s *OperaContext) OUTSD() antlr.TerminalNode {
	return s.GetToken(MASMParserOUTSD, 0)
}

func (s *OperaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterOpera(s)
	}
}

func (s *OperaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitOpera(s)
	}
}

func (s *OperaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitOpera(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Opera() (localctx IOperaContext) {
	localctx = NewOperaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MASMParserRULE_opera)
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
	p.SetState(338)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 85)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 85))) & ((1 << (MASMParserPOPAD - 85)) | (1 << (MASMParserAAA - 85)) | (1 << (MASMParserPOPA - 85)) | (1 << (MASMParserPOPFD - 85)) | (1 << (MASMParserCWD - 85)) | (1 << (MASMParserLAHF - 85)) | (1 << (MASMParserPUSHAD - 85)) | (1 << (MASMParserPUSHF - 85)) | (1 << (MASMParserAAS - 85)) | (1 << (MASMParserBSWAP - 85)) | (1 << (MASMParserPUSHFD - 85)) | (1 << (MASMParserCBW - 85)) | (1 << (MASMParserCWDE - 85)) | (1 << (MASMParserXLAT - 85)) | (1 << (MASMParserAAM - 85)) | (1 << (MASMParserAAD - 85)) | (1 << (MASMParserCDQ - 85)) | (1 << (MASMParserDAA - 85)) | (1 << (MASMParserSAHF - 85)) | (1 << (MASMParserDAS - 85)) | (1 << (MASMParserINTO - 85)) | (1 << (MASMParserIRET - 85)) | (1 << (MASMParserCLC - 85)) | (1 << (MASMParserSTC - 85)) | (1 << (MASMParserCMC - 85)) | (1 << (MASMParserCLD - 85)) | (1 << (MASMParserSTD - 85)) | (1 << (MASMParserCLI - 85)) | (1 << (MASMParserSTI - 85)) | (1 << (MASMParserMOVSB - 85)) | (1 << (MASMParserMOVSW - 85)) | (1 << (MASMParserMOVSD - 85)))) != 0) || ((((_la - 117)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 117))) & ((1 << (MASMParserLODS - 117)) | (1 << (MASMParserLODSB - 117)) | (1 << (MASMParserLODSW - 117)) | (1 << (MASMParserLODSD - 117)) | (1 << (MASMParserSTOS - 117)) | (1 << (MASMParserSTOSB - 117)) | (1 << (MASMParserSTOSW - 117)) | (1 << (MASMParserSOTSD - 117)) | (1 << (MASMParserSCAS - 117)) | (1 << (MASMParserSCASB - 117)) | (1 << (MASMParserSCASW - 117)) | (1 << (MASMParserSCASD - 117)) | (1 << (MASMParserCMPS - 117)) | (1 << (MASMParserCMPSB - 117)) | (1 << (MASMParserCMPSW - 117)) | (1 << (MASMParserCMPSD - 117)) | (1 << (MASMParserINSB - 117)) | (1 << (MASMParserINSW - 117)) | (1 << (MASMParserINSD - 117)) | (1 << (MASMParserOUTSB - 117)) | (1 << (MASMParserOUTSW - 117)) | (1 << (MASMParserOUTSD - 117)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IOperatContext is an interface to support dynamic dispatch.
type IOperatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatContext differentiates from other interfaces.
	IsOperatContext()
}

type OperatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatContext() *OperatContext {
	var p = new(OperatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_operat
	return p
}

func (*OperatContext) IsOperatContext() {}

func NewOperatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatContext {
	var p = new(OperatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_operat

	return p
}

func (s *OperatContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatContext) ADC() antlr.TerminalNode {
	return s.GetToken(MASMParserADC, 0)
}

func (s *OperatContext) ADD() antlr.TerminalNode {
	return s.GetToken(MASMParserADD, 0)
}

func (s *OperatContext) SUB() antlr.TerminalNode {
	return s.GetToken(MASMParserSUB, 0)
}

func (s *OperatContext) CBB() antlr.TerminalNode {
	return s.GetToken(MASMParserCBB, 0)
}

func (s *OperatContext) XOR() antlr.TerminalNode {
	return s.GetToken(MASMParserXOR, 0)
}

func (s *OperatContext) OR() antlr.TerminalNode {
	return s.GetToken(MASMParserOR, 0)
}

func (s *OperatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterOperat(s)
	}
}

func (s *OperatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitOperat(s)
	}
}

func (s *OperatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitOperat(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Operat() (localctx IOperatContext) {
	localctx = NewOperatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MASMParserRULE_operat)
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
	p.SetState(340)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 139)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 139))) & ((1 << (MASMParserADC - 139)) | (1 << (MASMParserADD - 139)) | (1 << (MASMParserSUB - 139)) | (1 << (MASMParserCBB - 139)) | (1 << (MASMParserXOR - 139)) | (1 << (MASMParserOR - 139)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IOperatoContext is an interface to support dynamic dispatch.
type IOperatoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatoContext differentiates from other interfaces.
	IsOperatoContext()
}

type OperatoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatoContext() *OperatoContext {
	var p = new(OperatoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_operato
	return p
}

func (*OperatoContext) IsOperatoContext() {}

func NewOperatoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatoContext {
	var p = new(OperatoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_operato

	return p
}

func (s *OperatoContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatoContext) JNBE() antlr.TerminalNode {
	return s.GetToken(MASMParserJNBE, 0)
}

func (s *OperatoContext) JNZ() antlr.TerminalNode {
	return s.GetToken(MASMParserJNZ, 0)
}

func (s *OperatoContext) JPO() antlr.TerminalNode {
	return s.GetToken(MASMParserJPO, 0)
}

func (s *OperatoContext) JZ() antlr.TerminalNode {
	return s.GetToken(MASMParserJZ, 0)
}

func (s *OperatoContext) JS() antlr.TerminalNode {
	return s.GetToken(MASMParserJS, 0)
}

func (s *OperatoContext) LOOPNZ() antlr.TerminalNode {
	return s.GetToken(MASMParserLOOPNZ, 0)
}

func (s *OperatoContext) JGE() antlr.TerminalNode {
	return s.GetToken(MASMParserJGE, 0)
}

func (s *OperatoContext) JB() antlr.TerminalNode {
	return s.GetToken(MASMParserJB, 0)
}

func (s *OperatoContext) JNB() antlr.TerminalNode {
	return s.GetToken(MASMParserJNB, 0)
}

func (s *OperatoContext) JO() antlr.TerminalNode {
	return s.GetToken(MASMParserJO, 0)
}

func (s *OperatoContext) JP() antlr.TerminalNode {
	return s.GetToken(MASMParserJP, 0)
}

func (s *OperatoContext) JNO() antlr.TerminalNode {
	return s.GetToken(MASMParserJNO, 0)
}

func (s *OperatoContext) JNL() antlr.TerminalNode {
	return s.GetToken(MASMParserJNL, 0)
}

func (s *OperatoContext) JNAE() antlr.TerminalNode {
	return s.GetToken(MASMParserJNAE, 0)
}

func (s *OperatoContext) LOOPZ() antlr.TerminalNode {
	return s.GetToken(MASMParserLOOPZ, 0)
}

func (s *OperatoContext) JMP() antlr.TerminalNode {
	return s.GetToken(MASMParserJMP, 0)
}

func (s *OperatoContext) JNP() antlr.TerminalNode {
	return s.GetToken(MASMParserJNP, 0)
}

func (s *OperatoContext) LOOP() antlr.TerminalNode {
	return s.GetToken(MASMParserLOOP, 0)
}

func (s *OperatoContext) JL() antlr.TerminalNode {
	return s.GetToken(MASMParserJL, 0)
}

func (s *OperatoContext) JCXZ() antlr.TerminalNode {
	return s.GetToken(MASMParserJCXZ, 0)
}

func (s *OperatoContext) JAE() antlr.TerminalNode {
	return s.GetToken(MASMParserJAE, 0)
}

func (s *OperatoContext) JNGE() antlr.TerminalNode {
	return s.GetToken(MASMParserJNGE, 0)
}

func (s *OperatoContext) JA() antlr.TerminalNode {
	return s.GetToken(MASMParserJA, 0)
}

func (s *OperatoContext) LOOPNE() antlr.TerminalNode {
	return s.GetToken(MASMParserLOOPNE, 0)
}

func (s *OperatoContext) LOOPE() antlr.TerminalNode {
	return s.GetToken(MASMParserLOOPE, 0)
}

func (s *OperatoContext) JG() antlr.TerminalNode {
	return s.GetToken(MASMParserJG, 0)
}

func (s *OperatoContext) JNLE() antlr.TerminalNode {
	return s.GetToken(MASMParserJNLE, 0)
}

func (s *OperatoContext) JE() antlr.TerminalNode {
	return s.GetToken(MASMParserJE, 0)
}

func (s *OperatoContext) JNC() antlr.TerminalNode {
	return s.GetToken(MASMParserJNC, 0)
}

func (s *OperatoContext) JC() antlr.TerminalNode {
	return s.GetToken(MASMParserJC, 0)
}

func (s *OperatoContext) JNA() antlr.TerminalNode {
	return s.GetToken(MASMParserJNA, 0)
}

func (s *OperatoContext) JBE() antlr.TerminalNode {
	return s.GetToken(MASMParserJBE, 0)
}

func (s *OperatoContext) JLE() antlr.TerminalNode {
	return s.GetToken(MASMParserJLE, 0)
}

func (s *OperatoContext) JPE() antlr.TerminalNode {
	return s.GetToken(MASMParserJPE, 0)
}

func (s *OperatoContext) JNS() antlr.TerminalNode {
	return s.GetToken(MASMParserJNS, 0)
}

func (s *OperatoContext) JECXZ() antlr.TerminalNode {
	return s.GetToken(MASMParserJECXZ, 0)
}

func (s *OperatoContext) JNG() antlr.TerminalNode {
	return s.GetToken(MASMParserJNG, 0)
}

func (s *OperatoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperatoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterOperato(s)
	}
}

func (s *OperatoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitOperato(s)
	}
}

func (s *OperatoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitOperato(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Operato() (localctx IOperatoContext) {
	localctx = NewOperatoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MASMParserRULE_operato)
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
	p.SetState(342)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 145)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 145))) & ((1 << (MASMParserJNBE - 145)) | (1 << (MASMParserJNZ - 145)) | (1 << (MASMParserJPO - 145)) | (1 << (MASMParserJZ - 145)) | (1 << (MASMParserJS - 145)) | (1 << (MASMParserLOOPNZ - 145)) | (1 << (MASMParserJGE - 145)) | (1 << (MASMParserJB - 145)) | (1 << (MASMParserJNB - 145)) | (1 << (MASMParserJO - 145)) | (1 << (MASMParserJP - 145)) | (1 << (MASMParserJNO - 145)) | (1 << (MASMParserJNL - 145)) | (1 << (MASMParserJNAE - 145)) | (1 << (MASMParserLOOPZ - 145)) | (1 << (MASMParserJMP - 145)) | (1 << (MASMParserJNP - 145)) | (1 << (MASMParserLOOP - 145)) | (1 << (MASMParserJL - 145)) | (1 << (MASMParserJCXZ - 145)) | (1 << (MASMParserJAE - 145)) | (1 << (MASMParserJNGE - 145)) | (1 << (MASMParserJA - 145)) | (1 << (MASMParserLOOPNE - 145)) | (1 << (MASMParserLOOPE - 145)) | (1 << (MASMParserJG - 145)) | (1 << (MASMParserJNLE - 145)) | (1 << (MASMParserJE - 145)) | (1 << (MASMParserJNC - 145)) | (1 << (MASMParserJC - 145)) | (1 << (MASMParserJNA - 145)) | (1 << (MASMParserJBE - 145)))) != 0) || ((((_la - 177)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 177))) & ((1 << (MASMParserJLE - 177)) | (1 << (MASMParserJPE - 177)) | (1 << (MASMParserJNS - 177)) | (1 << (MASMParserJECXZ - 177)) | (1 << (MASMParserJNG - 177)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) MOVZX() antlr.TerminalNode {
	return s.GetToken(MASMParserMOVZX, 0)
}

func (s *OperatorContext) BSF() antlr.TerminalNode {
	return s.GetToken(MASMParserBSF, 0)
}

func (s *OperatorContext) BSR() antlr.TerminalNode {
	return s.GetToken(MASMParserBSR, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (s *OperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MASMParserRULE_operator)
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
	p.SetState(344)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 182)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 182))) & ((1 << (MASMParserMOVZX - 182)) | (1 << (MASMParserBSF - 182)) | (1 << (MASMParserBSR - 182)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// ILContext is an interface to support dynamic dispatch.
type ILContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLContext differentiates from other interfaces.
	IsLContext()
}

type LContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLContext() *LContext {
	var p = new(LContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_l
	return p
}

func (*LContext) IsLContext() {}

func NewLContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LContext {
	var p = new(LContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_l

	return p
}

func (s *LContext) GetParser() antlr.Parser { return s.parser }

func (s *LContext) LES() antlr.TerminalNode {
	return s.GetToken(MASMParserLES, 0)
}

func (s *LContext) LEA() antlr.TerminalNode {
	return s.GetToken(MASMParserLEA, 0)
}

func (s *LContext) LDS() antlr.TerminalNode {
	return s.GetToken(MASMParserLDS, 0)
}

func (s *LContext) INS() antlr.TerminalNode {
	return s.GetToken(MASMParserINS, 0)
}

func (s *LContext) OUTS() antlr.TerminalNode {
	return s.GetToken(MASMParserOUTS, 0)
}

func (s *LContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterL(s)
	}
}

func (s *LContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitL(s)
	}
}

func (s *LContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitL(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) L() (localctx ILContext) {
	localctx = NewLContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MASMParserRULE_l)
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
	p.SetState(346)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 185)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 185))) & ((1 << (MASMParserLES - 185)) | (1 << (MASMParserLEA - 185)) | (1 << (MASMParserLDS - 185)) | (1 << (MASMParserINS - 185)) | (1 << (MASMParserOUTS - 185)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IXContext is an interface to support dynamic dispatch.
type IXContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXContext differentiates from other interfaces.
	IsXContext()
}

type XContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXContext() *XContext {
	var p = new(XContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_x
	return p
}

func (*XContext) IsXContext() {}

func NewXContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XContext {
	var p = new(XContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_x

	return p
}

func (s *XContext) GetParser() antlr.Parser { return s.parser }

func (s *XContext) XADD() antlr.TerminalNode {
	return s.GetToken(MASMParserXADD, 0)
}

func (s *XContext) CMPXCHG() antlr.TerminalNode {
	return s.GetToken(MASMParserCMPXCHG, 0)
}

func (s *XContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *XContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterX(s)
	}
}

func (s *XContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitX(s)
	}
}

func (s *XContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitX(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) X() (localctx IXContext) {
	localctx = NewXContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MASMParserRULE_x)
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
	p.SetState(348)
	_la = p.GetTokenStream().LA(1)

	if !(_la == MASMParserXADD || _la == MASMParserCMPXCHG) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_s
	return p
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) SHL() antlr.TerminalNode {
	return s.GetToken(MASMParserSHL, 0)
}

func (s *SContext) SHR() antlr.TerminalNode {
	return s.GetToken(MASMParserSHR, 0)
}

func (s *SContext) ROR() antlr.TerminalNode {
	return s.GetToken(MASMParserROR, 0)
}

func (s *SContext) ROL() antlr.TerminalNode {
	return s.GetToken(MASMParserROL, 0)
}

func (s *SContext) RCL() antlr.TerminalNode {
	return s.GetToken(MASMParserRCL, 0)
}

func (s *SContext) SAL() antlr.TerminalNode {
	return s.GetToken(MASMParserSAL, 0)
}

func (s *SContext) RCR() antlr.TerminalNode {
	return s.GetToken(MASMParserRCR, 0)
}

func (s *SContext) SAR() antlr.TerminalNode {
	return s.GetToken(MASMParserSAR, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitS(s)
	}
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MASMParserRULE_s)
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
	p.SetState(350)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 192)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 192))) & ((1 << (MASMParserSHL - 192)) | (1 << (MASMParserSHR - 192)) | (1 << (MASMParserROR - 192)) | (1 << (MASMParserROL - 192)) | (1 << (MASMParserRCL - 192)) | (1 << (MASMParserSAL - 192)) | (1 << (MASMParserRCR - 192)) | (1 << (MASMParserSAR - 192)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IShContext is an interface to support dynamic dispatch.
type IShContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShContext differentiates from other interfaces.
	IsShContext()
}

type ShContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShContext() *ShContext {
	var p = new(ShContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_sh
	return p
}

func (*ShContext) IsShContext() {}

func NewShContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShContext {
	var p = new(ShContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_sh

	return p
}

func (s *ShContext) GetParser() antlr.Parser { return s.parser }

func (s *ShContext) SHRD() antlr.TerminalNode {
	return s.GetToken(MASMParserSHRD, 0)
}

func (s *ShContext) SHLD() antlr.TerminalNode {
	return s.GetToken(MASMParserSHLD, 0)
}

func (s *ShContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterSh(s)
	}
}

func (s *ShContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitSh(s)
	}
}

func (s *ShContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitSh(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Sh() (localctx IShContext) {
	localctx = NewShContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, MASMParserRULE_sh)
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
	p.SetState(352)
	_la = p.GetTokenStream().LA(1)

	if !(_la == MASMParserSHRD || _la == MASMParserSHLD) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IBContext is an interface to support dynamic dispatch.
type IBContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBContext differentiates from other interfaces.
	IsBContext()
}

type BContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBContext() *BContext {
	var p = new(BContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_b
	return p
}

func (*BContext) IsBContext() {}

func NewBContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BContext {
	var p = new(BContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_b

	return p
}

func (s *BContext) GetParser() antlr.Parser { return s.parser }

func (s *BContext) BTR() antlr.TerminalNode {
	return s.GetToken(MASMParserBTR, 0)
}

func (s *BContext) BT() antlr.TerminalNode {
	return s.GetToken(MASMParserBT, 0)
}

func (s *BContext) BTC() antlr.TerminalNode {
	return s.GetToken(MASMParserBTC, 0)
}

func (s *BContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterB(s)
	}
}

func (s *BContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitB(s)
	}
}

func (s *BContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitB(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) B() (localctx IBContext) {
	localctx = NewBContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, MASMParserRULE_b)
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
	p.SetState(354)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 202)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 202))) & ((1 << (MASMParserBTR - 202)) | (1 << (MASMParserBT - 202)) | (1 << (MASMParserBTC - 202)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) CALL() antlr.TerminalNode {
	return s.GetToken(MASMParserCALL, 0)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitCall(s)
	}
}

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, MASMParserRULE_call)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(MASMParserCALL)
	}



	return localctx
}


// IInterruptionContext is an interface to support dynamic dispatch.
type IInterruptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterruptionContext differentiates from other interfaces.
	IsInterruptionContext()
}

type InterruptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterruptionContext() *InterruptionContext {
	var p = new(InterruptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_interruption
	return p
}

func (*InterruptionContext) IsInterruptionContext() {}

func NewInterruptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterruptionContext {
	var p = new(InterruptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_interruption

	return p
}

func (s *InterruptionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterruptionContext) INT() antlr.TerminalNode {
	return s.GetToken(MASMParserINT, 0)
}

func (s *InterruptionContext) RETN() antlr.TerminalNode {
	return s.GetToken(MASMParserRETN, 0)
}

func (s *InterruptionContext) RET() antlr.TerminalNode {
	return s.GetToken(MASMParserRET, 0)
}

func (s *InterruptionContext) RETF() antlr.TerminalNode {
	return s.GetToken(MASMParserRETF, 0)
}

func (s *InterruptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterruptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InterruptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterInterruption(s)
	}
}

func (s *InterruptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitInterruption(s)
	}
}

func (s *InterruptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitInterruption(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Interruption() (localctx IInterruptionContext) {
	localctx = NewInterruptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, MASMParserRULE_interruption)
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
	p.SetState(358)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 206)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 206))) & ((1 << (MASMParserINT - 206)) | (1 << (MASMParserRETN - 206)) | (1 << (MASMParserRET - 206)) | (1 << (MASMParserRETF - 206)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IInContext is an interface to support dynamic dispatch.
type IInContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInContext differentiates from other interfaces.
	IsInContext()
}

type InContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInContext() *InContext {
	var p = new(InContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_in
	return p
}

func (*InContext) IsInContext() {}

func NewInContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InContext {
	var p = new(InContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_in

	return p
}

func (s *InContext) GetParser() antlr.Parser { return s.parser }

func (s *InContext) IN() antlr.TerminalNode {
	return s.GetToken(MASMParserIN, 0)
}

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitIn(s)
	}
}

func (s *InContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitIn(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) In() (localctx IInContext) {
	localctx = NewInContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, MASMParserRULE_in)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(360)
		p.Match(MASMParserIN)
	}



	return localctx
}


// IOutContext is an interface to support dynamic dispatch.
type IOutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutContext differentiates from other interfaces.
	IsOutContext()
}

type OutContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutContext() *OutContext {
	var p = new(OutContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_out
	return p
}

func (*OutContext) IsOutContext() {}

func NewOutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutContext {
	var p = new(OutContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_out

	return p
}

func (s *OutContext) GetParser() antlr.Parser { return s.parser }

func (s *OutContext) OUT() antlr.TerminalNode {
	return s.GetToken(MASMParserOUT, 0)
}

func (s *OutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterOut(s)
	}
}

func (s *OutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitOut(s)
	}
}

func (s *OutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitOut(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Out() (localctx IOutContext) {
	localctx = NewOutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, MASMParserRULE_out)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(362)
		p.Match(MASMParserOUT)
	}



	return localctx
}


// IReContext is an interface to support dynamic dispatch.
type IReContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReContext differentiates from other interfaces.
	IsReContext()
}

type ReContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReContext() *ReContext {
	var p = new(ReContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_re
	return p
}

func (*ReContext) IsReContext() {}

func NewReContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReContext {
	var p = new(ReContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_re

	return p
}

func (s *ReContext) GetParser() antlr.Parser { return s.parser }

func (s *ReContext) REP() antlr.TerminalNode {
	return s.GetToken(MASMParserREP, 0)
}

func (s *ReContext) REPE() antlr.TerminalNode {
	return s.GetToken(MASMParserREPE, 0)
}

func (s *ReContext) REPZ() antlr.TerminalNode {
	return s.GetToken(MASMParserREPZ, 0)
}

func (s *ReContext) REPNE() antlr.TerminalNode {
	return s.GetToken(MASMParserREPNE, 0)
}

func (s *ReContext) REPNZ() antlr.TerminalNode {
	return s.GetToken(MASMParserREPNZ, 0)
}

func (s *ReContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterRe(s)
	}
}

func (s *ReContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitRe(s)
	}
}

func (s *ReContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitRe(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Re() (localctx IReContext) {
	localctx = NewReContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, MASMParserRULE_re)
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
	p.SetState(364)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 212)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 212))) & ((1 << (MASMParserREP - 212)) | (1 << (MASMParserREPE - 212)) | (1 << (MASMParserREPZ - 212)) | (1 << (MASMParserREPNE - 212)) | (1 << (MASMParserREPNZ - 212)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IDirectivesContext is an interface to support dynamic dispatch.
type IDirectivesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectivesContext differentiates from other interfaces.
	IsDirectivesContext()
}

type DirectivesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectivesContext() *DirectivesContext {
	var p = new(DirectivesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_directives
	return p
}

func (*DirectivesContext) IsDirectivesContext() {}

func NewDirectivesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectivesContext {
	var p = new(DirectivesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_directives

	return p
}

func (s *DirectivesContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectivesContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(MASMParserALPHA, 0)
}

func (s *DirectivesContext) CONST() antlr.TerminalNode {
	return s.GetToken(MASMParserCONST, 0)
}

func (s *DirectivesContext) CREF() antlr.TerminalNode {
	return s.GetToken(MASMParserCREF, 0)
}

func (s *DirectivesContext) XCREF() antlr.TerminalNode {
	return s.GetToken(MASMParserXCREF, 0)
}

func (s *DirectivesContext) DATA() antlr.TerminalNode {
	return s.GetToken(MASMParserDATA, 0)
}

func (s *DirectivesContext) DATA2() antlr.TerminalNode {
	return s.GetToken(MASMParserDATA2, 0)
}

func (s *DirectivesContext) DOSSEG() antlr.TerminalNode {
	return s.GetToken(MASMParserDOSSEG, 0)
}

func (s *DirectivesContext) ERR() antlr.TerminalNode {
	return s.GetToken(MASMParserERR, 0)
}

func (s *DirectivesContext) ERR1() antlr.TerminalNode {
	return s.GetToken(MASMParserERR1, 0)
}

func (s *DirectivesContext) ERR2() antlr.TerminalNode {
	return s.GetToken(MASMParserERR2, 0)
}

func (s *DirectivesContext) ERRE() antlr.TerminalNode {
	return s.GetToken(MASMParserERRE, 0)
}

func (s *DirectivesContext) ERRNZ() antlr.TerminalNode {
	return s.GetToken(MASMParserERRNZ, 0)
}

func (s *DirectivesContext) ERRDEF() antlr.TerminalNode {
	return s.GetToken(MASMParserERRDEF, 0)
}

func (s *DirectivesContext) ERRNDEF() antlr.TerminalNode {
	return s.GetToken(MASMParserERRNDEF, 0)
}

func (s *DirectivesContext) ERRB() antlr.TerminalNode {
	return s.GetToken(MASMParserERRB, 0)
}

func (s *DirectivesContext) ERRNB() antlr.TerminalNode {
	return s.GetToken(MASMParserERRNB, 0)
}

func (s *DirectivesContext) ERRIDN() antlr.TerminalNode {
	return s.GetToken(MASMParserERRIDN, 0)
}

func (s *DirectivesContext) ERRDIF() antlr.TerminalNode {
	return s.GetToken(MASMParserERRDIF, 0)
}

func (s *DirectivesContext) EVEN() antlr.TerminalNode {
	return s.GetToken(MASMParserEVEN, 0)
}

func (s *DirectivesContext) LIST() antlr.TerminalNode {
	return s.GetToken(MASMParserLIST, 0)
}

func (s *DirectivesContext) WIDTH() antlr.TerminalNode {
	return s.GetToken(MASMParserWIDTH, 0)
}

func (s *DirectivesContext) MASK() antlr.TerminalNode {
	return s.GetToken(MASMParserMASK, 0)
}

func (s *DirectivesContext) SEQ() antlr.TerminalNode {
	return s.GetToken(MASMParserSEQ, 0)
}

func (s *DirectivesContext) XLIST() antlr.TerminalNode {
	return s.GetToken(MASMParserXLIST, 0)
}

func (s *DirectivesContext) EXIT() antlr.TerminalNode {
	return s.GetToken(MASMParserEXIT, 0)
}

func (s *DirectivesContext) MODEL() antlr.TerminalNode {
	return s.GetToken(MASMParserMODEL, 0)
}

func (s *DirectivesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectivesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DirectivesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterDirectives(s)
	}
}

func (s *DirectivesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitDirectives(s)
	}
}

func (s *DirectivesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitDirectives(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Directives() (localctx IDirectivesContext) {
	localctx = NewDirectivesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, MASMParserRULE_directives)
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
	p.SetState(366)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 217)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 217))) & ((1 << (MASMParserALPHA - 217)) | (1 << (MASMParserCONST - 217)) | (1 << (MASMParserCREF - 217)) | (1 << (MASMParserXCREF - 217)) | (1 << (MASMParserDATA - 217)) | (1 << (MASMParserDATA2 - 217)) | (1 << (MASMParserDOSSEG - 217)) | (1 << (MASMParserERR - 217)) | (1 << (MASMParserERR1 - 217)) | (1 << (MASMParserERR2 - 217)) | (1 << (MASMParserERRE - 217)) | (1 << (MASMParserERRNZ - 217)) | (1 << (MASMParserERRDEF - 217)) | (1 << (MASMParserERRNDEF - 217)) | (1 << (MASMParserERRB - 217)) | (1 << (MASMParserERRNB - 217)) | (1 << (MASMParserERRIDN - 217)) | (1 << (MASMParserERRDIF - 217)) | (1 << (MASMParserEVEN - 217)) | (1 << (MASMParserLIST - 217)) | (1 << (MASMParserWIDTH - 217)) | (1 << (MASMParserMASK - 217)) | (1 << (MASMParserSEQ - 217)) | (1 << (MASMParserXLIST - 217)) | (1 << (MASMParserEXIT - 217)) | (1 << (MASMParserMODEL - 217)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// ITyContext is an interface to support dynamic dispatch.
type ITyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTyContext differentiates from other interfaces.
	IsTyContext()
}

type TyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyContext() *TyContext {
	var p = new(TyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_ty
	return p
}

func (*TyContext) IsTyContext() {}

func NewTyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TyContext {
	var p = new(TyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_ty

	return p
}

func (s *TyContext) GetParser() antlr.Parser { return s.parser }

func (s *TyContext) BYTE() antlr.TerminalNode {
	return s.GetToken(MASMParserBYTE, 0)
}

func (s *TyContext) SBYTE() antlr.TerminalNode {
	return s.GetToken(MASMParserSBYTE, 0)
}

func (s *TyContext) DB() antlr.TerminalNode {
	return s.GetToken(MASMParserDB, 0)
}

func (s *TyContext) WORD() antlr.TerminalNode {
	return s.GetToken(MASMParserWORD, 0)
}

func (s *TyContext) SWORD() antlr.TerminalNode {
	return s.GetToken(MASMParserSWORD, 0)
}

func (s *TyContext) DW() antlr.TerminalNode {
	return s.GetToken(MASMParserDW, 0)
}

func (s *TyContext) DWORD() antlr.TerminalNode {
	return s.GetToken(MASMParserDWORD, 0)
}

func (s *TyContext) SDWORD() antlr.TerminalNode {
	return s.GetToken(MASMParserSDWORD, 0)
}

func (s *TyContext) DD() antlr.TerminalNode {
	return s.GetToken(MASMParserDD, 0)
}

func (s *TyContext) FWORD() antlr.TerminalNode {
	return s.GetToken(MASMParserFWORD, 0)
}

func (s *TyContext) DF() antlr.TerminalNode {
	return s.GetToken(MASMParserDF, 0)
}

func (s *TyContext) QWORD() antlr.TerminalNode {
	return s.GetToken(MASMParserQWORD, 0)
}

func (s *TyContext) DQ() antlr.TerminalNode {
	return s.GetToken(MASMParserDQ, 0)
}

func (s *TyContext) TBYTE() antlr.TerminalNode {
	return s.GetToken(MASMParserTBYTE, 0)
}

func (s *TyContext) DT() antlr.TerminalNode {
	return s.GetToken(MASMParserDT, 0)
}

func (s *TyContext) REAL4() antlr.TerminalNode {
	return s.GetToken(MASMParserREAL4, 0)
}

func (s *TyContext) REAL8() antlr.TerminalNode {
	return s.GetToken(MASMParserREAL8, 0)
}

func (s *TyContext) REAL() antlr.TerminalNode {
	return s.GetToken(MASMParserREAL, 0)
}

func (s *TyContext) FAR() antlr.TerminalNode {
	return s.GetToken(MASMParserFAR, 0)
}

func (s *TyContext) NEAR() antlr.TerminalNode {
	return s.GetToken(MASMParserNEAR, 0)
}

func (s *TyContext) PROC() antlr.TerminalNode {
	return s.GetToken(MASMParserPROC, 0)
}

func (s *TyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterTy(s)
	}
}

func (s *TyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitTy(s)
	}
}

func (s *TyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitTy(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Ty() (localctx ITyContext) {
	localctx = NewTyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, MASMParserRULE_ty)
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
	p.SetState(368)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 243)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 243))) & ((1 << (MASMParserBYTE - 243)) | (1 << (MASMParserSBYTE - 243)) | (1 << (MASMParserDB - 243)) | (1 << (MASMParserWORD - 243)) | (1 << (MASMParserSWORD - 243)) | (1 << (MASMParserDW - 243)) | (1 << (MASMParserDWORD - 243)) | (1 << (MASMParserSDWORD - 243)) | (1 << (MASMParserDD - 243)) | (1 << (MASMParserFWORD - 243)) | (1 << (MASMParserDF - 243)) | (1 << (MASMParserQWORD - 243)) | (1 << (MASMParserDQ - 243)) | (1 << (MASMParserTBYTE - 243)) | (1 << (MASMParserDT - 243)) | (1 << (MASMParserREAL4 - 243)) | (1 << (MASMParserREAL8 - 243)) | (1 << (MASMParserREAL - 243)) | (1 << (MASMParserFAR - 243)) | (1 << (MASMParserNEAR - 243)) | (1 << (MASMParserPROC - 243)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IQuestionContext is an interface to support dynamic dispatch.
type IQuestionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuestionContext differentiates from other interfaces.
	IsQuestionContext()
}

type QuestionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuestionContext() *QuestionContext {
	var p = new(QuestionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MASMParserRULE_question
	return p
}

func (*QuestionContext) IsQuestionContext() {}

func NewQuestionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuestionContext {
	var p = new(QuestionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_question

	return p
}

func (s *QuestionContext) GetParser() antlr.Parser { return s.parser }

func (s *QuestionContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(MASMParserQUESTION, 0)
}

func (s *QuestionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuestionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *QuestionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterQuestion(s)
	}
}

func (s *QuestionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitQuestion(s)
	}
}

func (s *QuestionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitQuestion(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Question() (localctx IQuestionContext) {
	localctx = NewQuestionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, MASMParserRULE_question)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(370)
		p.Match(MASMParserQUESTION)
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
	p.RuleIndex = MASMParserRULE_time
	return p
}

func (*TimeContext) IsTimeContext() {}

func NewTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeContext {
	var p = new(TimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MASMParserRULE_time

	return p
}

func (s *TimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeContext) TIMES() antlr.TerminalNode {
	return s.GetToken(MASMParserTIMES, 0)
}

func (s *TimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MASMListener); ok {
		listenerT.ExitTime(s)
	}
}

func (s *TimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MASMVisitor:
		return t.VisitTime(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MASMParser) Time() (localctx ITimeContext) {
	localctx = NewTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, MASMParserRULE_time)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(372)
		p.Match(MASMParserTIMES)
	}



	return localctx
}


