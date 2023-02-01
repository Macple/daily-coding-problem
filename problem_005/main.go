package main

import (
	"fmt"
	"strconv"
)

// This problem was asked by Facebook.
// Given the mapping a : 1, b : 2, ... z : 26, and an encoded message, count the number of ways it can be decoded.
// For example, the message "111" would give 3, since it could be decoded as "aaa", "ka", and "ak".
// You can assume that the messages are decodable. For example, "001" is not allowed.

func main() {
	fmt.Println(codesCount("111"))
}

func codesCount(code string) int {
	n := len(code)
	dp := make([]int, n+1)
	dp[0] = 1

	if code[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		oneDigit, _ := strconv.Atoi(code[i-1 : i])
		twoDigits, _ := strconv.Atoi(code[i-2 : i])
		if oneDigit >= 1 && oneDigit <= 9 {
			dp[i] += dp[i-1]
		}
		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
