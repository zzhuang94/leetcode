/**
 * Best Time to Buy and Sell Stock II
 *
 * Say you have an array for which the ith element is the price of a given stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete as many transactions
 * as you like (ie, buy one and sell one share of the stock multiple times).
 * However, you may not engage in multiple transactions at the same time
 * (ie, you must sell the stock before you buy again).
 */

package main

import "fmt"

func main() {
	prices := []int{100, 1, 1, 5, 6, 1, 1, -1, 6, 10, 4, 5, 5}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	len := len(prices)
	if len == 0 {
		return 0
	}
	sum, bought := 0, false
	l, h := 0, 0
	for i := 0; i < len-1; i++ {
		if !bought && prices[i] < prices[i+1] {
			l = prices[i]
			bought = true

		}
		if bought && prices[i] > prices[i+1] {
			h = prices[i]
			fmt.Println("sell", h)
			sum += h - l
			bought = false
		}
	}
	if bought {
		h = prices[len-1]
		sum += h - l
	}
	return sum
}
