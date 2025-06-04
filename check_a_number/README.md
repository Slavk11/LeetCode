# 🔢 Number Classification — Abundant, Perfect, or Deficient

## 💡 Challenge

**Difficulty:** Easy

Write a function that checks whether a number is:

- **Abundant** — if the sum of its proper divisors is **greater** than the number
- **Perfect** — if the sum of its proper divisors is **equal** to the number
- **Deficient** — if the sum of its proper divisors is **less** than the number

---

## 📘 Definition

> A **proper divisor** of a number is a positive integer that divides the number **without a remainder**, excluding the number itself.

---

## ✅ Classification Rules

| Type       | Condition                         |
|------------|-----------------------------------|
| Abundant   | sum of divisors **>** number      |
| Perfect    | sum of divisors **==** number     |
| Deficient  | sum of divisors **<** number      |

---

## 🧪 Examples

- `12` → Divisors: `1, 2, 3, 4, 6` → Sum: `16` → **Abundant**
- `28` → Divisors: `1, 2, 4, 7, 14` → Sum: `28` → **Perfect**
- `10` → Divisors: `1, 2, 5` → Sum: `8` → **Deficient**


