// Copyright 2021 Optakt Labs OÜ
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package zbor

// transactionDictionary is a byte slice that contains the result of running the Zstandard training mode
// on the transactions of the DPS index. This allows zstandard to achieve a better compression ratio, specifically for
// small data.
// See http://facebook.github.io/zstd/#small-data
// See https://github.com/facebook/zstd/blob/master/doc/zstd_compression_format.md#dictionary-format
var transactionDictionary = []byte{
	55, 164, 48, 236, 119, 253, 44, 5, 41, 16, 64, 93, 127, 215, 34, 140, 1, 22, 138, 91, 75, 104, 66, 58, 184, 71, 20, 131, 242, 139, 73, 63, 165, 100, 177, 98, 178, 161, 150, 126, 250, 233, 255, 244, 245, 209, 131, 85, 104, 107, 19, 8, 0, 0, 0, 7, 224, 0, 192, 202, 46, 0, 4, 0, 0, 0, 0, 0, 73, 0, 2, 0, 28, 29, 0, 0, 0, 26, 29, 30, 0, 64, 7, 0, 48, 92, 0, 0, 0, 0, 0, 0, 0, 56, 0, 7, 0, 56, 128, 19, 0, 0, 208, 1, 0, 0, 20, 52, 7, 0, 0, 0, 0, 128, 6, 0, 0, 0, 0, 12, 11, 0, 0, 0, 232, 0, 232, 0, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0, 0, 8, 0, 0, 0, 105, 99, 40, 34, 67, 111, 117, 108, 100, 32, 110, 111, 116, 32, 98, 111, 114, 114, 111, 119, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 116, 111, 32, 116, 104, 101, 32, 111, 119, 110, 101, 114, 39, 115, 32, 86, 97, 117, 108, 116, 33, 34, 41, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 47, 47, 32, 87, 105, 116, 104, 100, 114, 97, 119, 32, 116, 111, 107, 101, 110, 115, 32, 102, 114, 111, 109, 32, 116, 104, 101, 32, 115, 105, 103, 110, 101, 114, 39, 115, 32, 115, 116, 111, 114, 101, 100, 32, 118, 97, 117, 108, 116, 10, 32, 32, 32, 32, 32, 32, 32, 32, 115, 101, 108, 102, 46, 115, 101, 110, 116, 86, 97, 117, 108, 116, 32, 60, 45, 32, 118, 97, 117, 108, 116, 82, 101, 102, 46, 119, 105, 116, 104, 100, 114, 97, 119, 40, 97, 109, 111, 117, 110, 116, 58, 32, 97, 109, 111, 117, 110, 116, 41, 10, 32, 32, 32, 32, 125, 10, 10, 32, 32, 32, 32, 101, 120, 101, 99, 117, 116, 101, 32, 123, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 47, 47, 32, 71, 101, 116, 32, 97, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 116, 111, 32, 116, 104, 101, 32, 114, 101, 99, 105, 112, 105, 101, 110, 116, 39, 115, 32, 82, 101, 99, 101, 105, 118, 101, 114, 10, 32, 32, 32, 32, 32, 32, 32, 32, 108, 101, 116, 32, 114, 101, 99, 101, 105, 118, 101, 114, 82, 101, 102, 32, 61, 32, 32, 103, 101, 116, 65, 99, 99, 111, 117, 110, 116, 40, 116, 111, 41, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 46, 103, 101, 116, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 40, 47, 112, 117, 98, 108, 105, 99, 47, 102, 108, 111, 119, 84, 111, 107, 101, 110, 82, 101, 99, 101, 105, 118, 101, 114, 41, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 46, 98, 111, 114, 114, 111, 119, 60, 38, 123, 70, 117, 110, 103, 105, 98, 108, 101, 84, 111, 107, 101, 110, 46, 82, 101, 99, 101, 105, 118, 101, 114, 125, 62, 40, 41, 10, 9, 9, 9, 63, 63, 32, 112, 97, 110, 105, 99, 40, 34, 67, 111, 117, 108, 100, 32, 110, 111, 116, 32, 98, 111, 114, 114, 111, 119, 32, 114, 101, 99, 101, 105, 118, 101, 114, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 116, 111, 32, 116, 104, 101, 32, 114, 101, 99, 105, 112, 105, 101, 110, 116, 39, 115, 32, 86, 97, 117, 108, 116, 34, 41, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 47, 47, 32, 68, 101, 112, 111, 115, 105, 116, 32, 116, 104, 101, 32, 119, 105, 116, 104, 100, 114, 97, 119, 110, 32, 116, 111, 107, 101, 110, 115, 32, 105, 110, 32, 116, 104, 101, 32, 114, 101, 99, 105, 112, 105, 101, 110, 116, 39, 115, 32, 114, 101, 99, 101, 105, 118, 101, 114, 10, 32, 32, 32, 32, 32, 32, 32, 32, 114, 101, 99, 101, 105, 118, 101, 114, 82, 101, 102, 46, 100, 101, 112, 111, 115, 105, 116, 40, 102, 114, 111, 109, 58, 32, 60, 45, 115, 101, 108, 102, 46, 115, 101, 110, 116, 86, 97, 117, 108, 116, 41, 10, 32, 32, 32, 32, 125, 10, 125, 10, 104, 71, 97, 115, 76, 105, 109, 105, 116, 25, 39, 15, 105, 65, 114, 103, 117, 109, 101, 110, 116, 115, 130, 88, 39, 123, 34, 116, 121, 112, 101, 34, 58, 34, 85, 70, 105, 120, 54, 52, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 48, 46, 48, 48, 48, 48, 48, 48, 48, 49, 34, 125, 10, 88, 48, 123, 34, 116, 121, 112, 101, 34, 58, 34, 65, 100, 100, 114, 101, 115, 115, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 48, 120, 48, 57, 102, 56, 53, 53, 54, 97, 51, 53, 56, 57, 54, 48, 57, 50, 34, 125, 10, 107, 65, 117, 116, 104, 111, 114, 105, 122, 101, 114, 115, 129, 72, 9, 74, 246, 67, 169, 70, 212, 115, 107, 80, 114, 111, 112, 111, 115, 97, 108, 75, 101, 121, 163, 103, 65, 100, 100, 114, 101, 115, 115, 72, 9, 74, 246, 67, 169, 70, 212, 115, 104, 75, 101, 121, 73, 110, 100, 101, 120, 0, 110, 83, 101, 113, 117, 101, 110, 99, 101, 78, 117, 109, 98, 101, 114, 24, 66, 112, 82, 101, 102, 101, 114, 101, 110, 99, 101, 66, 108, 111, 99, 107, 73, 68, 88, 32, 109, 23, 229, 230, 64, 8, 190, 87, 111, 19, 114, 220, 63, 177, 26, 86, 156, 95, 136, 235, 251, 58, 161, 236, 164, 16, 230, 39, 43, 138, 81, 62, 113, 80, 97, 121, 108, 111, 97, 100, 83, 105, 103, 110, 97, 116, 117, 114, 101, 115, 246, 114, 69, 110, 118, 101, 108, 111, 112, 101, 83, 105, 103, 110, 97, 116, 117, 114, 101, 115, 129, 164, 103, 65, 100, 100, 114, 101, 115, 115, 72, 9, 74, 246, 67, 169, 70, 212, 115, 104, 75, 101, 121, 73, 110, 100, 101, 120, 0, 105, 83, 105, 103, 110, 97, 116, 117, 114, 101, 88, 64, 113, 105, 226, 197, 151, 129, 136, 5, 104, 75, 60, 31, 100, 204, 84, 112, 244, 223, 114, 15, 122, 89, 59, 91, 166, 246, 185, 199, 63, 236, 24, 232, 56, 214, 196, 222, 197, 39, 239, 225, 69, 24, 39, 229, 235, 57, 213, 57, 244, 161, 33, 88, 130, 194, 171, 226, 77, 73, 230, 3, 51, 115, 239, 141, 107, 83, 105, 103, 110, 101, 114, 73, 110, 100, 101, 120, 0, 169, 101, 80, 97, 121, 101, 114, 72, 12, 109, 76, 38, 211, 200, 50, 8, 102, 83, 99, 114, 105, 112, 116, 116, 117, 114, 101, 88, 64, 193, 185, 232, 83, 248, 71, 64, 213, 39, 41, 228, 15, 89, 252, 215, 3, 250, 170, 31, 213, 63, 223, 198, 121, 8, 170, 187, 47, 148, 1, 112, 23, 39, 185, 251, 251, 130, 210, 80, 2, 16, 188, 192, 129, 112, 95, 106, 154, 83, 166, 229, 5, 28, 199, 150, 140, 127, 46, 32, 141, 68, 158, 43, 242, 107, 83, 105, 103, 110, 101, 114, 73, 110, 100, 101, 120, 0, 169, 101, 80, 97, 121, 101, 114, 72, 248, 214, 224, 88, 107, 10, 32, 199, 102, 83, 99, 114, 105, 112, 116, 89, 2, 105, 10, 105, 109, 112, 111, 114, 116, 32, 70, 117, 110, 103, 105, 98, 108, 101, 84, 111, 107, 101, 110, 32, 102, 114, 111, 109, 32, 48, 120, 101, 101, 56, 50, 56, 53, 54, 98, 102, 50, 48, 101, 50, 97, 97, 54, 10, 105, 109, 112, 111, 114, 116, 32, 70, 108, 111, 119, 84, 111, 107, 101, 110, 32, 102, 114, 111, 109, 32, 48, 120, 48, 97, 101, 53, 51, 99, 98, 54, 101, 51, 102, 52, 50, 97, 55, 57, 10, 10, 116, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 40, 107, 101, 121, 58, 32, 83, 116, 114, 105, 110, 103, 44, 32, 97, 109, 111, 117, 110, 116, 58, 32, 85, 70, 105, 120, 54, 52, 41, 32, 123, 10, 9, 112, 114, 101, 112, 97, 114, 101, 40, 115, 105, 103, 110, 101, 114, 58, 32, 65, 117, 116, 104, 65, 99, 99, 111, 117, 110, 116, 41, 32, 123, 10, 9, 9, 108, 101, 116, 32, 97, 99, 99, 111, 117, 110, 116, 32, 61, 32, 65, 117, 116, 104, 65, 99, 99, 111, 117, 110, 116, 40, 112, 97, 121, 101, 114, 58, 32, 115, 105, 103, 110, 101, 114, 41, 10, 9, 9, 97, 99, 99, 111, 117, 110, 116, 46, 97, 100, 100, 80, 117, 98, 108, 105, 99, 75, 101, 121, 40, 107, 101, 121, 46, 100, 101, 99, 111, 100, 101, 72, 101, 120, 40, 41, 41, 10, 9, 9, 108, 101, 116, 32, 115, 101, 110, 100, 101, 114, 32, 61, 32, 115, 105, 103, 110, 101, 114, 46, 98, 111, 114, 114, 111, 119, 60, 38, 70, 108, 111, 119, 84, 111, 107, 101, 110, 46, 86, 97, 117, 108, 116, 62, 40, 102, 114, 111, 109, 58, 32, 47, 115, 116, 111, 114, 97, 103, 101, 47, 102, 108, 111, 119, 84, 111, 107, 101, 110, 86, 97, 117, 108, 116, 41, 10, 9, 9, 9, 63, 63, 32, 112, 97, 110, 105, 99, 40, 34, 99, 111, 117, 108, 100, 32, 110, 111, 116, 32, 98, 111, 114, 114, 111, 119, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 102, 111, 114, 32, 115, 101, 110, 100, 105, 110, 103, 32, 118, 97, 117, 108, 116, 34, 41, 10, 9, 9, 108, 101, 116, 32, 114, 101, 99, 101, 105, 118, 101, 114, 32, 61, 32, 97, 99, 99, 111, 117, 110, 116, 46, 103, 101, 116, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 40, 47, 112, 117, 98, 108, 105, 99, 47, 102, 108, 111, 119, 84, 111, 107, 101, 110, 82, 101, 99, 101, 105, 118, 101, 114, 41, 10, 9, 9, 9, 46, 98, 111, 114, 114, 111, 119, 60, 38, 123, 70, 117, 110, 103, 105, 98, 108, 101, 84, 111, 107, 101, 110, 46, 82, 101, 99, 101, 105, 118, 101, 114, 125, 62, 40, 41, 10, 9, 9, 9, 63, 63, 32, 112, 97, 110, 105, 99, 40, 34, 99, 111, 117, 108, 100, 32, 110, 111, 116, 32, 98, 111, 114, 114, 111, 119, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 102, 111, 114, 32, 114, 101, 99, 101, 105, 118, 105, 110, 103, 32, 118, 97, 117, 108, 116, 34, 41, 10, 9, 9, 114, 101, 99, 101, 105, 118, 101, 114, 46, 100, 101, 112, 111, 115, 105, 116, 40, 102, 114, 111, 109, 58, 32, 60, 45, 115, 101, 110, 100, 101, 114, 46, 119, 105, 116, 104, 100, 114, 97, 119, 40, 97, 109, 111, 117, 110, 116, 58, 32, 97, 109, 111, 117, 110, 116, 41, 41, 10, 9, 125, 10, 125, 10, 104, 71, 97, 115, 76, 105, 109, 105, 116, 0, 105, 65, 114, 103, 117, 109, 101, 110, 116, 115, 130, 88, 175, 123, 34, 116, 121, 112, 101, 34, 58, 34, 83, 116, 114, 105, 110, 103, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 102, 56, 52, 55, 98, 56, 52, 48, 98, 56, 99, 52, 54, 52, 48, 50, 56, 99, 100, 98, 98, 48, 49, 53, 54, 99, 98, 102, 57, 101, 99, 100, 49, 53, 49, 100, 55, 100, 51, 56, 48, 100, 57, 100, 53, 49, 52, 98, 98, 99, 52, 49, 55, 52, 48, 56, 56, 48, 56, 54, 51, 100, 51, 53, 55, 52, 97, 97, 100, 48, 97, 49, 101, 97, 48, 57, 53, 100, 52, 99, 50, 50, 52, 100, 51, 51, 51, 56, 55, 56, 54, 52, 57, 51, 56, 50, 102, 55, 57, 99, 48, 53, 54, 98, 97, 49, 57, 54, 97, 98, 102, 97, 99, 55, 56, 53, 53, 97, 51, 97, 54, 53, 54, 51, 51, 55, 57, 57, 51, 51, 54, 50, 101, 102, 50, 49, 48, 50, 48, 51, 56, 50, 48, 51, 101, 56, 34, 125, 10, 88, 39, 123, 34, 116, 121, 112, 101, 34, 58, 34, 85, 70, 105, 120, 54, 52, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 49, 46, 48, 48, 48, 48, 48, 48, 48, 48, 34, 125, 10, 107, 65, 117, 116, 104, 111, 114, 105, 122, 101, 114, 115, 129, 72, 248, 214, 224, 88, 107, 10, 32, 199, 107, 80, 114, 111, 112, 111, 115, 97, 108, 75, 101, 121, 163, 103, 65, 100, 100, 114, 101, 115, 115, 72, 248, 214, 224, 88, 107, 10, 32, 199, 104, 75, 101, 121, 73, 110, 100, 32, 47, 47, 32, 87, 105, 116, 104, 100, 114, 97, 119, 32, 116, 111, 107, 101, 110, 115, 32, 102, 114, 111, 109, 32, 116, 104, 101, 32, 115, 105, 103, 110, 101, 114, 39, 115, 32, 115, 116, 111, 114, 101, 100, 32, 118, 97, 117, 108, 116, 10, 32, 32, 32, 32, 32, 32, 32, 32, 115, 101, 108, 102, 46, 115, 101, 110, 116, 86, 97, 117, 108, 116, 32, 60, 45, 32, 118, 97, 117, 108, 116, 82, 101, 102, 46, 119, 105, 116, 104, 100, 114, 97, 119, 40, 97, 109, 111, 117, 110, 116, 58, 32, 97, 109, 111, 117, 110, 116, 41, 10, 32, 32, 32, 32, 125, 10, 10, 32, 32, 32, 32, 101, 120, 101, 99, 117, 116, 101, 32, 123, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 47, 47, 32, 71, 101, 116, 32, 97, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 116, 111, 32, 116, 104, 101, 32, 114, 101, 99, 105, 112, 105, 101, 110, 116, 39, 115, 32, 82, 101, 99, 101, 105, 118, 101, 114, 10, 32, 32, 32, 32, 32, 32, 32, 32, 108, 101, 116, 32, 114, 101, 99, 101, 105, 118, 101, 114, 82, 101, 102, 32, 61, 32, 32, 103, 101, 116, 65, 99, 99, 111, 117, 110, 116, 40, 116, 111, 41, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 46, 103, 101, 116, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 40, 47, 112, 117, 98, 108, 105, 99, 47, 102, 108, 111, 119, 84, 111, 107, 101, 110, 82, 101, 99, 101, 105, 118, 101, 114, 41, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 46, 98, 111, 114, 114, 111, 119, 60, 38, 123, 70, 117, 110, 103, 105, 98, 108, 101, 84, 111, 107, 101, 110, 46, 82, 101, 99, 101, 105, 118, 101, 114, 125, 62, 40, 41, 10, 9, 9, 9, 63, 63, 32, 112, 97, 110, 105, 99, 40, 34, 67, 111, 117, 108, 100, 32, 110, 111, 116, 32, 98, 111, 114, 114, 111, 119, 32, 114, 101, 99, 101, 105, 118, 101, 114, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 116, 111, 32, 116, 104, 101, 32, 114, 101, 99, 105, 112, 105, 101, 110, 116, 39, 115, 32, 86, 97, 117, 108, 116, 34, 41, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 47, 47, 32, 68, 101, 112, 111, 115, 105, 116, 32, 116, 104, 101, 32, 119, 105, 116, 104, 100, 114, 97, 119, 110, 32, 116, 111, 107, 101, 110, 115, 32, 105, 110, 32, 116, 104, 101, 32, 114, 101, 99, 105, 112, 105, 101, 110, 116, 39, 115, 32, 114, 101, 99, 101, 105, 118, 101, 114, 10, 32, 32, 32, 32, 32, 32, 32, 32, 114, 101, 99, 101, 105, 118, 101, 114, 82, 101, 102, 46, 100, 101, 112, 111, 115, 105, 116, 40, 102, 114, 111, 109, 58, 32, 60, 45, 115, 101, 108, 102, 46, 115, 101, 110, 116, 86, 97, 117, 108, 116, 41, 10, 32, 32, 32, 32, 125, 10, 125, 10, 104, 71, 97, 115, 76, 105, 109, 105, 116, 25, 39, 15, 105, 65, 114, 103, 117, 109, 101, 110, 116, 115, 130, 88, 39, 123, 34, 116, 121, 112, 101, 34, 58, 34, 85, 70, 105, 120, 54, 52, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 48, 46, 48, 48, 48, 48, 48, 48, 48, 49, 34, 125, 10, 88, 48, 123, 34, 116, 121, 112, 101, 34, 58, 34, 65, 100, 100, 114, 101, 115, 115, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 48, 120, 102, 98, 55, 57, 50, 97, 97, 100, 50, 49, 98, 56, 100, 101, 99, 100, 34, 125, 10, 107, 65, 117, 116, 104, 111, 114, 105, 122, 101, 114, 115, 129, 72, 231, 98, 59, 84, 77, 175, 144, 35, 107, 80, 114, 111, 112, 111, 115, 97, 108, 75, 101, 121, 163, 103, 65, 100, 100, 114, 101, 115, 115, 72, 231, 98, 59, 84, 77, 175, 144, 35, 104, 75, 101, 121, 73, 110, 100, 101, 120, 0, 110, 83, 101, 113, 117, 101, 110, 99, 101, 78, 117, 109, 98, 101, 114, 24, 51, 112, 82, 101, 102, 101, 114, 101, 110, 99, 101, 66, 108, 111, 99, 107, 73, 68, 88, 32, 238, 64, 182, 223, 179, 226, 138, 222, 253, 189, 166, 98, 15, 147, 114, 156, 255, 237, 64, 136, 214, 133, 30, 143, 32, 133, 53, 69, 110, 98, 35, 192, 113, 80, 97, 121, 108, 111, 97, 100, 83, 105, 103, 110, 97, 116, 117, 114, 101, 115, 246, 114, 69, 110, 118, 101, 108, 111, 112, 101, 83, 105, 103, 110, 97, 116, 117, 114, 101, 115, 129, 164, 103, 65, 100, 100, 114, 101, 115, 115, 72, 231, 98, 59, 84, 77, 175, 144, 35, 104, 75, 101, 121, 73, 110, 100, 101, 120, 0, 105, 83, 105, 103, 110, 97, 116, 117, 114, 101, 88, 64, 28, 224, 97, 58, 28, 108, 180, 147, 115, 177, 94, 84, 32, 166, 176, 184, 250, 80, 207, 1, 50, 115, 110, 117, 19, 27, 7, 169, 50, 192, 199, 221, 218, 8, 206, 246, 237, 45, 60, 77, 92, 21, 217, 196, 144, 187, 18, 33, 156, 21, 227, 166, 119, 195, 83, 147, 209, 196, 142, 73, 174, 197, 224, 57, 107, 83, 105, 103, 110, 101, 114, 73, 110, 100, 101, 120, 0, 169, 101, 80, 97, 121, 101, 114, 72, 98, 28, 208, 165, 249, 69, 7, 174, 102, 83, 99, 114, 105, 112, 116, 89, 4, 12, 10, 105, 109, 112, 111, 114, 116, 32, 70, 117, 110, 103, 105, 98, 108, 101, 84, 111, 107, 101, 110, 32, 102, 114, 111, 109, 32, 48, 120, 101, 101, 56, 50, 56, 53, 54, 98, 102, 50, 48, 101, 50, 97, 97, 54, 10, 105, 109, 112, 111, 114, 116, 32, 70, 108, 111, 119, 84, 111, 107, 111, 119, 84, 111, 107, 101, 110, 32, 102, 114, 111, 109, 32, 48, 120, 48, 97, 101, 53, 51, 99, 98, 54, 101, 51, 102, 52, 50, 97, 55, 57, 10, 10, 116, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 40, 97, 109, 111, 117, 110, 116, 58, 32, 85, 70, 105, 120, 54, 52, 44, 32, 116, 111, 58, 32, 65, 100, 100, 114, 101, 115, 115, 41, 32, 123, 10, 10, 32, 32, 32, 32, 47, 47, 32, 84, 104, 101, 32, 86, 97, 117, 108, 116, 32, 114, 101, 115, 111, 117, 114, 99, 101, 32, 116, 104, 97, 116, 32, 104, 111, 108, 100, 115, 32, 116, 104, 101, 32, 116, 111, 107, 101, 110, 115, 32, 116, 104, 97, 116, 32, 97, 114, 101, 32, 98, 101, 105, 110, 103, 32, 116, 114, 97, 110, 115, 102, 101, 114, 114, 101, 100, 10, 32, 32, 32, 32, 108, 101, 116, 32, 115, 101, 110, 116, 86, 97, 117, 108, 116, 58, 32, 64, 70, 117, 110, 103, 105, 98, 108, 101, 84, 111, 107, 101, 110, 46, 86, 97, 117, 108, 116, 10, 10, 32, 32, 32, 32, 112, 114, 101, 112, 97, 114, 101, 40, 115, 105, 103, 110, 101, 114, 58, 32, 65, 117, 116, 104, 65, 99, 99, 111, 117, 110, 116, 41, 32, 123, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 47, 47, 32, 71, 101, 116, 32, 97, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 116, 111, 32, 116, 104, 101, 32, 115, 105, 103, 110, 101, 114, 39, 115, 32, 115, 116, 111, 114, 101, 100, 32, 118, 97, 117, 108, 116, 10, 32, 32, 32, 32, 32, 32, 32, 32, 108, 101, 116, 32, 118, 97, 117, 108, 116, 82, 101, 102, 32, 61, 32, 115, 105, 103, 110, 101, 114, 46, 98, 111, 114, 114, 111, 119, 60, 38, 70, 108, 111, 119, 84, 111, 107, 101, 110, 46, 86, 97, 117, 108, 116, 62, 40, 102, 114, 111, 109, 58, 32, 47, 115, 116, 111, 114, 97, 103, 101, 47, 102, 108, 111, 119, 84, 111, 107, 101, 110, 86, 97, 117, 108, 116, 41, 10, 9, 9, 9, 63, 63, 32, 112, 97, 110, 105, 99, 40, 34, 67, 111, 117, 108, 100, 32, 110, 111, 116, 32, 98, 111, 114, 114, 111, 119, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 116, 111, 32, 116, 104, 101, 32, 111, 119, 110, 101, 114, 39, 115, 32, 86, 97, 117, 108, 116, 33, 34, 41, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 47, 47, 32, 87, 105, 116, 104, 100, 114, 97, 119, 32, 116, 111, 107, 101, 110, 115, 32, 102, 114, 111, 109, 32, 116, 104, 101, 32, 115, 105, 103, 110, 101, 114, 39, 115, 32, 115, 116, 111, 114, 101, 100, 32, 118, 97, 117, 108, 116, 10, 32, 32, 32, 32, 32, 32, 32, 32, 115, 101, 108, 102, 46, 115, 101, 110, 116, 86, 97, 117, 108, 116, 32, 60, 45, 32, 118, 97, 117, 108, 116, 82, 101, 102, 46, 119, 105, 116, 104, 100, 114, 97, 119, 40, 97, 109, 111, 117, 110, 116, 58, 32, 97, 109, 111, 117, 110, 116, 41, 10, 32, 32, 32, 32, 125, 10, 10, 32, 32, 32, 32, 101, 120, 101, 99, 117, 116, 101, 32, 123, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 47, 47, 32, 71, 101, 116, 32, 97, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 116, 111, 32, 116, 104, 101, 32, 114, 101, 99, 105, 112, 105, 101, 110, 116, 39, 115, 32, 82, 101, 99, 101, 105, 118, 101, 114, 10, 32, 32, 32, 32, 32, 32, 32, 32, 108, 101, 116, 32, 114, 101, 99, 101, 105, 118, 101, 114, 82, 101, 102, 32, 61, 32, 32, 103, 101, 116, 65, 99, 99, 111, 117, 110, 116, 40, 116, 111, 41, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 46, 103, 101, 116, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 40, 47, 112, 117, 98, 108, 105, 99, 47, 102, 108, 111, 119, 84, 111, 107, 101, 110, 82, 101, 99, 101, 105, 118, 101, 114, 41, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 46, 98, 111, 114, 114, 111, 119, 60, 38, 123, 70, 117, 110, 103, 105, 98, 108, 101, 84, 111, 107, 101, 110, 46, 82, 101, 99, 101, 105, 118, 101, 114, 125, 62, 40, 41, 10, 9, 9, 9, 63, 63, 32, 112, 97, 110, 105, 99, 40, 34, 67, 111, 117, 108, 100, 32, 110, 111, 116, 32, 98, 111, 114, 114, 111, 119, 32, 114, 101, 99, 101, 105, 118, 101, 114, 32, 114, 101, 102, 101, 114, 101, 110, 99, 101, 32, 116, 111, 32, 116, 104, 101, 32, 114, 101, 99, 105, 112, 105, 101, 110, 116, 39, 115, 32, 86, 97, 117, 108, 116, 34, 41, 10, 10, 32, 32, 32, 32, 32, 32, 32, 32, 47, 47, 32, 68, 101, 112, 111, 115, 105, 116, 32, 116, 104, 101, 32, 119, 105, 116, 104, 100,
}