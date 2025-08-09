package main

func MaxProfit(prices []int) int {
	// 購入日のインデックスを初日で初期化
	left := 0
	// 当日に売却はできないため、売却日のインデックスを翌日で初期化
	right := left + 1

	maxProfit := 0

	for right < len(prices) {
		// 今日の価格が過去の購入価格より高い場合、利益を計算して最大利益を更新
		if prices[right] > prices[left] {
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			// 今日の価格が最小価格の場合、購入日を更新
			left = right
		}
		right++
	}
	return maxProfit
}

func MaxProfit2(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, todayPrice := range prices[1:] {
		if todayPrice < minPrice {
			minPrice = todayPrice
		} else if profit := todayPrice - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
