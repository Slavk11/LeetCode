# 💠 Abundant Number

### 💡 Challenge: Easy

Write a function that takes an integer as input and checks if it is an **abundant number**.

---

## 📘 What is an Abundant Number?

An **abundant number** is a positive integer for which the **sum of its proper divisors** (excluding the number itself) is **greater than the number**.

In simpler terms:

> If the sum of all divisors (excluding the number) > the number itself → it's **abundant**.

---

## 🔍 Examples

### ✅ 12

- Proper divisors: `1, 2, 3, 4, 6`
- Sum: `1 + 2 + 3 + 4 + 6 = 16`
- Since `16 > 12`, → **12 is abundant**

---

### ✅ 18

- Proper divisors: `1, 2, 3, 6, 9`
- Sum: `1 + 2 + 3 + 6 + 9 = 21`
- Since `21 > 18`, → **18 is abundant**

---

### ❌ 10

- Proper divisors: `1, 2, 5`
- Sum: `1 + 2 + 5 = 8`
- Since `8 < 10`, → **10 is not abundant**

---

