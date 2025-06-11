# 🧠 Optimization Scenarios

## 1️⃣ All Numbers in [0, 100]

Use a **counting array** of size 101:

- **`AddNum()`**: Increment the count at index `num`
- **`FindMedian()`**: Walk through the count array to find the middle value(s)

➡️ **Time Complexity**:
- Insert: **O(1)**
- Median lookup: **O(100)** (constant for practical purposes)

---

## 2️⃣ 99% of Numbers in [0, 100]

Use a **hybrid approach**:

- Use **counting array** for numbers in `[0, 100]`
- Use a **heap or BST** for outliers (outside this range)
- **Track total count** to know whether the median lies in the array or in the heap/BST

➡️ **Benefits**:
- More efficient than using full heaps
- Reduces unnecessary overhead for the common case
