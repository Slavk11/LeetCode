
---

# 🔡 String Compression

### Challenge: Medium

Implement a function that takes a string as input and returns a **compressed version** of that string.
Each character in the string should be followed by one or more `"#"` symbols.
The number of `"#"` reflects **how many times the character has appeared so far** (including the current one).

---

## 📘 Description

For every character in the input string:

* Append the character to the result
* Append `#` **N times**, where **N** is how many times this character has appeared **up to this point**

---

### 📥 Input

* A non-empty string `s` consisting of lowercase ASCII letters

### 📤 Output

* A new string representing the compressed version

---

## 💡 Example

### Input:

```text
aabccchbbccaaa
```

### Output:

```text
a##b#c###h#b##c##a###
```

---

## 🧠 Explanation

Step by step:

| Index | Char | Count So Far | Output Fragment |
| ----- | ---- | ------------ | --------------- |
| 0     | a    | 1            | a#              |
| 1     | a    | 2            | a##             |
| 2     | b    | 1            | b#              |
| 3     | c    | 1            | c#              |
| 4     | c    | 2            | c##             |
| 5     | c    | 3            | c###            |
| 6     | h    | 1            | h#              |
| 7     | b    | 2            | b##             |
| 8     | b    | 3            | b###            |
| 9     | c    | 4            | c####           |
| 10    | c    | 5            | c#####          |
| 11    | a    | 3            | a###            |
| 12    | a    | 4            | a####           |

Final result:

```text
a##b#c###h#b##c##a###
```

---

## 🛠 Sample Implementation (Go)

```
func CompressString(input string) string {
	if len(input) == 0 {
		return ""
	}

	var builder strings.Builder
	counts := make(map[rune]int)

	for _, ch := range input {
		counts[ch]++
		builder.WriteRune(ch)
		builder.WriteString(strings.Repeat("#", counts[ch]))
	}

	return builder.String()
}
```

---

## 🧪 Test Cases

| Input      | Expected Output  |
| ---------- | ---------------- |
| `"aab"`    | `a##b#`          |
| `"google"` | `g#o#o##g##l#e#` |
| `"coddy"`  | `c#o#d##y#`      |
| `"golang"` | `g#o#l#a#n#g##`  |

