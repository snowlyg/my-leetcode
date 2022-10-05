package main

import "fmt"

func main() {
	fmt.Println(maxPoints([][]int{[]int{1, 1}, []int{2, 2}, []int{3, 3}}))
}

/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */

// @lc code=start
func maxPoints(points [][]int) int {
	ans := 0
	// 两点以下直接返回
	l := len(points)
	if l <= 2 {
		return l
	}

	for i, p := range points {

		// 超过一半，或者超过剩下的点 就可以确定当前直线为最多
		if ans >= l-i || ans > l/2 {
			break
		}

		cnt := map[int]int{}

		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			cnt[y+x*20001]++
		}

		// 比较直线点数量
		for _, v := range cnt {
			ans = max(ans, v+1)
		}
	}

	return ans
}

// 斜率
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
