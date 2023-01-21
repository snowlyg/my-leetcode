//go:build 204
// +build 204

package main

import "log"

func main() {
	log.Println(countPrimes(34))
}

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	isPrimes := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrimes[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isPrimes[i] {
			for j := i * i; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			count++
		}
	}
	return count
}

// @lc code=end
