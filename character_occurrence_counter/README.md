# 🔍 Character Occurrence Counter

## 💡 Challenge — Medium

Given a string like:

# 🔍 Character Occurrence Counter

## 💡 Challenge — Medium

Given a string like:


Write a function that **counts how many times a given character or substring appears** in the input string.

---

## ✅ Requirements

- The function should take **two string parameters**:
    1. `searchStr` — the characters or substring to search for.
    2. `inputStr` — the string in which to search.
- It should return the **number of non-overlapping occurrences** of `searchStr` in `inputStr`.
- The comparison should be **case-insensitive**.

---

## 🧪 Example

```go
countOccurrences("avgofrcxgolangcoddy", "go") // returns 2
countOccurrences("ghojhfghdavhfdh", "h")       // returns 5
countOccurrences("HelloHELLOheLLo", "hello")   // returns 3
