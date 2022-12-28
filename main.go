package main

import (
	"fmt"
	"strings"
)

// task1
func solution(s string) int {
	// stringArrays := []string{}
	count := 0
	temp := 0
	check := true
	tempArr := ""
	length := len(s)
	for i := 0; i < length-1; i++ {
		fmt.Printf("%c\n", s[i])
		if !strings.Contains(tempArr, string(s[i])) {
			for j := i + 1; j < length; j++ {
				if s[i] == s[j] {
					temp += 1
					if !strings.Contains(tempArr, string(s[j])) {
						tempArr += string(s[i])
					}

				}
			}
			if temp%2 == 0 {
				if check == true {
					check = false
				} else {
					fmt.Printf("out: %c\n", s[i])
					fmt.Printf("%c\n", s[i])
					count += 1
				}
			}
		}
		temp = 0
	}
	fmt.Printf("string: %s", tempArr)
	return count
}

// task2
func solution2(message string, k int) string {
	index := 0
	if len(message) <= k {
		return message
	} else {

		for i := 0; i < k-3; i++ {
			if string(message[i]) == " " {
				index = i
				for j := i; j < k-3; j++ {
					if string(message[i]) == " " {
						index = j
						i = j
					}
				}
			}
		}

	}
	temp := ""
	if index == 0 {
		return "..."
	} else {
		temp = message[:index]
		temp += "..."
		return temp
	}
}

// task3
//func solution3(A [][]int) int {
//	result := 0
//	arrValue := 0
//	for i := 0; i < len(A); i++ {
//		for j := 0; j < len(A[i]); j++ {
//			arrValue += A[i][j]
//		}
//	}
//	for i := 0; i < len(A)-1; i++ {
//		for j := i; j < len(A[i]); j++ {
//			if j == (len(A[i]) - 1) {
//				if arrValue-(A[i][j]+A[i+1][j]) == A[i][j]+A[i+1][j] {
//					result += 1
//				}
//			} else {
//				if arrValue-(A[i][j]+A[i+1][i]) == A[i][j]+A[i+1][i] {
//					result += 1
//				}
//				if arrValue-(A[i][j]+A[i][i+1]) == A[i][j]+A[i][i+1] {
//					result += 1
//				}
//			}
//		}
//	}
//	return result
//}

func main() {
	//fmt.Printf("Result: %d", solution("erveriige"))
	//fmt.Printf("Result: %s", solution2("super dog", 4))
	//var A = [][]int{{1, 1, -2}, {3, 2, 4}, {-1, -2, -2}}
	//var A = [][]int{{1, 1, 2}, {2, 0, 0}}
	//fmt.Printf("Result: %d", solution3(A))
}
