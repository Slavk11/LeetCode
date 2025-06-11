# 🧹 Remove Duplicates from Sorted Array

## 📚 Description

Given a **sorted** integer array `nums`, remove the **duplicates in-place** so that each unique element appears only once.  
The relative order of the elements must be preserved.

Your implementation should return the number of unique elements `k`.  
The first `k` elements of `nums` should contain the unique values in their original order.

> Elements beyond index `k-1` can be left unchanged — they will be ignored.

---

## ✅ Requirements

To be accepted, your solution must:

- Modify the input slice `nums` **in-place**
- Return `k` — the number of unique elements

The judge will verify your solution with the following logic (in Go):

```go
nums := []int{...}
expected := []int{...}

k := removeDuplicates(nums)

if k != len(expected) {
    panic("Length mismatch")
}
for i := 0; i < k; i++ {
    if nums[i] != expected[i] {
        panic("Element mismatch at index " + strconv.Itoa(i))
    }
}

💡 Examples
Example 1:

Input:  nums = []int{1, 1, 2}
Output: 2, nums = []int{1, 2, _}
