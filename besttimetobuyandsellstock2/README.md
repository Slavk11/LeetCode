
---

# 💹 Best Time to Buy and Sell Stock II

## 🧠 Challenge

You are given an integer array `prices` where `prices[i]` is the price of a given stock on the `i`-th day.

On each day, you may decide to **buy and/or sell** the stock.
You can **only hold one share at a time**, but you may **buy and sell on the same day**.

Your goal is to **maximize your profit** by choosing the best days to buy and sell.

Return the **maximum profit** you can achieve.

---

## 🧾 Example

### Example 1

```
Input:  prices = [7,1,5,3,6,4]
Output: 7

Explanation:
Buy on day 2 (price = 1), sell on day 3 (price = 5), profit = 4  
Buy on day 4 (price = 3), sell on day 5 (price = 6), profit = 3  
Total Profit = 4 + 3 = 7
```

### Example 2

```
Input:  prices = [1,2,3,4,5]
Output: 4

Explanation:
Buy on day 1 (price = 1), sell on day 5 (price = 5), profit = 4
```

### Example 3

```
Input:  prices = [7,6,4,3,1]
Output: 0

Explanation:
No profitable transaction is possible. Profit = 0
```

---

## ✅ Constraints

* `1 <= prices.length <= 3 * 10⁴`
* `0 <= prices[i] <= 10⁴`

---

## 💡 Strategy

To maximize profit, simply sum up every **positive difference** between consecutive days:

If `prices[i] > prices[i - 1]`, then `prices[i] - prices[i - 1]` is a profit worth taking.

---

## 🧪 Sample Function Signature (Go)

```
func maxProfit(prices []int) int
```
