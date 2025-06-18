# 🔍 Knuth-Morris-Pratt (KMP) Algorithm — Walkthrough & Explanation

This document provides a step-by-step explanation of how the **KMP algorithm** works using a real example.

---

## 📘 Example

- **Text**: `ababcabcabababd`
- **Pattern**: `ababd`

---

## 🔧 Step 1: Build the Pattern Table (LPS Array)

The **LPS (Longest Prefix Suffix)** array stores the length of the longest prefix that is also a suffix for each sub-pattern. It allows us to skip unnecessary comparisons.

### LPS Table Construction

We compare characters in the pattern to find repeated prefixes/suffixes.

#### Pattern:

Index: 0 1 2 3 4
Pattern: a b a b d


### Step-by-Step Construction

We use two pointers:
- `i`: current character being evaluated
- `j`: length of the current matching prefix

| i | j | pattern[i] | pattern[j] | Match? | Action                            | LPS[i] | Next i, j |
|---|---|------------|------------|--------|------------------------------------|--------|-----------|
| 1 | 0 | b          | a          | ❌     | Set LPS[1] = -1                    | -1     | i=2, j=0   |
| 2 | 0 | a          | a          | ✅     | Set LPS[2] = 0, j++                | 0      | i=3, j=1   |
| 3 | 1 | b          | b          | ✅     | Set LPS[3] = 1, j++                | 1      | i=4, j=2   |
| 4 | 2 | d          | a          | ❌     | Set j = LPS[1] + 1 = 0             | -      | j=0        |
| 4 | 0 | d          | a          | ❌     | Set LPS[4] = -1, i++               | -1     | i=5, j=0   |

### ✅ Final LPS Array

| Index | Char | LPS |
|-------|------|-----|
| 0     | a    | -1  |
| 1     | b    | -1  |
| 2     | a    | 0   |
| 3     | b    | 1   |
| 4     | d    | -1  |

---

## 🔍 Step 2: Pattern Matching

We now use the LPS array to search for the pattern in the text efficiently.

### Initial State

- Start with `i = 0` (text index), `j = 0` (pattern index)

### Step-by-Step Matching

| Text[i] | Pattern[j] | Match? | Action                                           |
|---------|------------|--------|--------------------------------------------------|
| a       | a          | ✅     | i++, j++                                         |
| b       | b          | ✅     | i++, j++                                         |
| a       | a          | ✅     | i++, j++                                         |
| b       | b          | ✅     | i++, j++                                         |
| c       | d          | ❌     | Use LPS → j = LPS[3] + 1 = 2                     |
| c       | a          | ❌     | Use LPS → j = LPS[1] + 1 = 0                     |
| c       | a          | ❌     | i++                                              |
| a       | a          | ✅     | i++, j++                                         |
| b       | b          | ✅     | i++, j++                                         |
| c       | a          | ❌     | Use LPS → j = 0                                  |
| ...     | ...        | ...    | Continue until full pattern match at index 10   |
| d       | d          | ✅     | Pattern fully matched → return True             |

---

## ✅ Final Result

- **Pattern "ababd" found at index 10** in the text.

---

## 🧠 Summary

### LPS Table Recap

| Index | Char | LPS |
|-------|------|-----|
| 0     | a    | -1  |
| 1     | b    | -1  |
| 2     | a    | 0   |
| 3     | b    | 1   |
| 4     | d    | -1  |

### Key Concepts

- LPS allows skipping re-evaluation of previously matched characters.
- Time complexity is **O(n + m)** where `n = len(text)` and `m = len(pattern)`.
- Efficient for large strings and repetitive pattern searches.

---

📌 The KMP algorithm avoids redundant comparisons by reusing previously computed prefix information — making it much faster than brute-force matching.
