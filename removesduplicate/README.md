# 🔁 Removes Duplicate

### 🧩 Challenge: Easy

Write a function that takes a string and **removes repeated adjacent duplicate characters** step by step — until no such duplicates remain.

Unlike simply removing all duplicates, this task removes **consecutive** duplicates **one pass at a time**, continuing until the string is fully reduced.

---

## ✏️ Description

Given a string, eliminate consecutive duplicate letters **iteratively**.  
In each pass, remove all groups of repeated characters that are adjacent.  
Repeat this process until there are no more such groups left.

---

### 🔍 Example

```txt
Input:  "aabccchbbccaaa"
Step 1: "abchbca"         ← "aa", "ccc", "bb", "cc", "aaa" removed
Output: "abchbca"
