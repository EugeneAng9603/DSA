# 🔍 KMP Algorithm — Walkthrough for Sample Input

## 📘 Sample Input

- **Text**: `"aefoaefcdaefcdaed"`
- **Pattern**: `"aefcdaed"`

## ✅ Expected Output

- **Result**: `True`  
  (The pattern is present in the text.)

---

## 🧱 Step 1: Build the Pattern Array (LPS Table)

The **LPS (Longest Prefix Suffix)** array helps us skip unnecessary comparisons when a mismatch occurs during pattern matching.

### Pattern:
| a | e | f | c | d | a | e | d |
| - | - | - | - | - | - | - | - |
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |


### LPS Array Construction

| Index | Char | LPS | Reason                                                                 |
|-------|------|-----|------------------------------------------------------------------------|
| 0     | a    | -1  | No prefix or suffix                                                    |
| 1     | e    | -1  | No match with prefix                                                   |
| 2     | f    | -1  | No match with prefix                                                   |
| 3     | c    | -1  | No match with prefix                                                   |
| 4     | d    | -1  | No match with prefix                                                   |
| 5     | a    | 0   | Matches pattern[0] ('a') → LPS = 0                                     |
| 6     | e    | 1   | Matches pattern[1] ('e') → LPS = 1                                     |
| 7     | d    | -1  | Mismatch with pattern[2] ('f') → reset to -1                           |

### ✅ Final LPS Array

| Index   | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|---------|---|---|---|---|---|---|---|---|
| Pattern | a | e | f | c | d | a | e | d |
| LPS     | -1| -1| -1| -1| -1|  0|  1| -1|

---

## 🔍 Step 2: Pattern Matching Using LPS Array

We now slide the pattern over the text using two pointers:
- `i`: index in text
- `j`: index in pattern

### Matching Process

- Begin comparison at `i = 0`, `j = 0`
- Continue comparing `text[i]` with `pattern[j]`
- Use the LPS array to avoid rechecking known matched characters when mismatches occur

Eventually, at **index 9** in the text, a full match for the pattern starts:


Text: aefoaefcdaefcdaed
↑
Pattern: aefcdaed → matches fully

Characters match from:
- `text[9]` to `text[16]`
- with pattern indices `0` to `7`

---

## ✅ Final Result

- The pattern `"aefcdaed"` **exists** in the text `"aefoaefcdaefcdaed"`.
- **Output: `True`**

---

## 🧠 Summary

- The KMP algorithm efficiently found the pattern by using the **LPS array** to skip redundant comparisons.
- It completed the search in **linear time** relative to the length of the input and pattern.

### Time and Space Complexity

| Aspect   | Complexity |
|----------|------------|
| Time     | O(N + M)   |
| Space    | O(M)       |

Where:
- `N` = length of the text
- `M` = length of the pattern

---

📌 This test case confirms that KMP correctly handles complex, repeated character sequences and performs efficient substring searches.
