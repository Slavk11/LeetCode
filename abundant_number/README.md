# 💡 Abundant Number

## 📘 Challenge — *Easy*

Write a function that takes a positive integer and checks whether it is an **abundant number**.

An **abundant number** is a positive integer for which the **sum of its proper divisors** (excluding the number itself) is **greater than the number**.  
In other words, the number is *abundant* in terms of its divisors.

---

## 🧠 What are Proper Divisors?

Proper divisors of a number are all its positive divisors **excluding the number itself**.

---

## ✅ Examples

- **12**  
  Divisors: `1, 2, 3, 4, 6`  
  Sum: `1 + 2 + 3 + 4 + 6 = 16`  
  → `16 > 12` → **Abundant**

- **18**  
  Divisors: `1, 2, 3, 6, 9`  
  Sum: `1 + 2 + 3 + 6 + 9 = 21`  
  → `21 > 18` → **Abundant**

- **10**  
  Divisors: `1, 2, 5`  
  Sum: `8`  
  → `8 < 10` → Not Abundant

---

## 🛠 Usage

You can implement the function in your preferred language. Here's what it should do:

```plaintext
Input: 12
Output: "Abundant"
