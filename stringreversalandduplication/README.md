
# 🔁 String Reversal & Duplication

## 🧩 Challenge — Easy

We have a string like this: `"abcdfgh"`.

Write a function that takes a string as input, **reverses it**, and **duplicates each character** to create a new string.

## 📥 Input

A single string `s`, containing any characters.  
Example:
```

"abcdfgh"

```

---

## 📤 Output

A new string where:

- The input is reversed
- Each character in the reversed string appears **twice**

Example:
```

Input:  "abcdfgh"
Output: "hhggffddccbbaa"

```



## ✅ Requirements

- The original string must remain unchanged
- The result must be constructed efficiently
- Return a new string as output



## 💡 Example

```
func ReverseAndDouble(input string) string {
    n := len(input)
    result := make([]byte, 2*n)

    for i := 0; i < n; i++ {
        ch := input[n-1-i]
        result[2*i] = ch
        result[2*i+1] = ch
    }

    return string(result)
}
```


