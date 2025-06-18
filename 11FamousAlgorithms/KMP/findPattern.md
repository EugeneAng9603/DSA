# 🔍 Knuth-Morris-Pratt (KMP) Algorithm Walkthrough

## 📘 Example

- **Text**: `"ababcabcabababd"`
- **Pattern (Substring)**: `"ababd"`

---

## 🔧 Step 1: Build the Pattern Table (LPS Array)

We construct the **LPS (Longest Prefix Suffix)** array for the pattern `"ababd"`.

### 🔁 Initialization

| Index | Char | LPS |
|-------|------|-----|
| 0     | a    | -1  |
| 1     | b    | -1  |
| 2     | a    | ?   |
| 3     | b    | ?   |
| 4     | d    | ?   |

### 🔄 LPS Construction Process

We initialize `i = 1`, `j = 0` and build the pattern:

| i | pattern[i] | pattern[j] | Match? | Action                                  | LPS[i] | Next i, j |
|---|------------|------------|--------|------------------------------------------|--------|-----------|
| 1 | b          | a          | ❌     | Set LPS[1] = -1                          | -1     | i=2, j=0   |
| 2 | a          | a          | ✅     | Set LPS[2] = j (0), then j++             | 0      | i=3, j=1   |
| 3 | b          | b          | ✅     | Set LPS[3] = j (1), then j++             | 1      | i=4, j=2   |
| 4 | d          | a          | ❌     | j = pattern[j - 1] + 1 = 0               | -      | j=0        |
| 4 | d          | a          | ❌     | Set LPS[4] = -1                          | -1     | i=5        |

### ✅ Final LPS Array


---

## 🔍 Step 2: Pattern Matching Using KMP

### Initialization

- `i = 0` (index in text)
- `j = 0` (index in pattern)

### Step-by-Step Matching

1. `text[0] == pattern[0]` → `'a' == 'a'` → ✅ → `i=1`, `j=1`
2. `text[1] == pattern[1]` → `'b' == 'b'` → ✅ → `i=2`, `j=2`
3. `text[2] == pattern[2]` → `'a' == 'a'` → ✅ → `i=3`, `j=3`
4. `text[3] == pattern[3]` → `'b' == 'b'` → ✅ → `i=4`, `j=4`
5. `text[4] == pattern[4]` → `'c' != 'd'` → ❌  
   - Jump to `j = LPS[3] + 1 = 2`

6. `text[4] == pattern[2]` → `'c' != 'a'` → ❌  
   - Jump to `j = LPS[1] + 1 = 0`

7. `text[4] == pattern[0]` → `'c' != 'a'` → ❌  
   - Increment `i = 5`

...

Eventually:

- Match occurs at `i = 10`  
- `text[10:15] == "ababd"` → ✅ Full match

---

## ✅ Result

- **Pattern found at index `10` in the text.**

---

## 🧠 Summary

### LPS Table for Pattern `"ababd"`

| Index | Char | LPS |
|-------|------|-----|
| 0     | a    | -1  |
| 1     | b    | -1  |
| 2     | a    | 0   |
| 3     | b    | 1   |
| 4     | d    | -1  |

### Final Match

- Pattern `"ababd"` is **found in the text** at index **10**.

---

📌 This walkthrough demonstrates how the KMP algorithm avoids re-checking characters and achieves **O(n + m)** time complexity.
