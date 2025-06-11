# 🎯 Goal

We want to:

- **Keep data structured** such that the **median is always accessible in O(1)**
- **Support fast insertion** in **O(log n)** time

---

# 🧪 Example Walkthrough

## 👣 Step-by-Step Insertion
Step-by-step logic:
Push to low first

We always insert into low (max-heap) because we assume a new number might belong in the smaller half.

Push max(low) to high

This ensures that any number that should belong in the larger half is transferred to high.

Rebalance if needed

If high now has more elements than low, push the min(high) back to low.

✅ Why this guarantees correctness:
All elements in low ≤ all elements in high (since we moved the largest of low to high)

Sizes of the heaps are always balanced:

low.Len() == high.Len() or low.Len() == high.Len() + 1

The top of low and high will always give correct median

### Example 1: Insert `[1]`
- `low`: **[1]** (max-heap)  
- `high`: `[]`  
- **Median**: top of low → `1`

---

### Example 2: Insert `[1, 2]`
- Insert 2 → goes into `low`
- Push `max(low)` (2) into `high` to balance

**After balancing**:
- `low`: **[1]**
- `high`: **[2]**
- **Median**: (1 + 2) / 2 → `1.5`

---

### Example 3: Insert `[1, 2, 3]`
- Add 3 → insert into `low` → [3, 1]
- Push `max(low)` (3) to `high`
- Now, `high`: [2, 3] → too big → push `min(high)` back to `low`

**After balancing**:
- `low`: **[2, 1]**
- `high`: **[3]**
- **Median**: top of low → `2`

---

### Example 4: Insert `[1, 2, 3, 4]`
- Add 4 → insert into `low`
- Push `max(low)` (4) to `high`

**After balancing**:
- `low`: **[2, 1]**
- `high`: **[3, 4]**
- **Median**: (2 + 3) / 2 → `2.5`

---

# 🧮 What's Happening?

We always:
- Keep the **smaller half in `low`** (max-heap)
- Keep the **larger half in `high`** (min-heap)
- Maintain balance:  
  `|low.Len() - high.Len()| <= 1`

### This structure lets us:
- **Access median directly**:
  - If odd count → top of `low`
  - If even count → average of tops of `low` and `high`

---

# 🧠 Why 2 Heaps Work

### ✅ Benefits:
- Efficient **insert**: `O(log n)`
- Efficient **median lookup**: `O(1)`
- No need for full sorting or scanning
- Dynamically balances incoming data

---

# ✅ Summary Table

| Input Stream        | low (max-heap) | high (min-heap) | Median |
|---------------------|----------------|------------------|--------|
| `[1]`               | [1]            | []               | 1      |
| `[1, 2]`            | [1]            | [2]              | 1.5    |
| `[1, 2, 3]`         | [2, 1]         | [3]              | 2      |
| `[1, 2, 3, 4]`      | [2, 1]         | [3, 4]           | 2.5    |
| `[1, 2, 3, 4, 5]`   | [3, 1, 2]      | [4, 5]           | 3      |
