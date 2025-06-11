package main

import (
	"sort"
)

type MedianFinderCountArr struct {
	count [101]int
	size  int
}

func ConstructorCountArr() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinderCountArr) AddNumCountArr(num int) {
	mf.count[num]++
	mf.size++
}

func (mf *MedianFinderCountArr) FindMedianCountArr() float64 {
	mid1, mid2 := -1, -1
	countSoFar := 0

	for i, freq := range mf.count {
		countSoFar += freq
		if mf.size%2 == 1 {
			if countSoFar >= mf.size/2+1 {
				return float64(i)
			}
		} else {
			if mid1 == -1 && countSoFar >= mf.size/2 {
				mid1 = i
			}
			if countSoFar >= mf.size/2+1 {
				mid2 = i
				return float64(mid1+mid2) / 2
			}
		}
	}
	return 0.0 // should not reach here
}

// ---------------------------

type HybridMedianFinder struct {
	count        [101]int
	inRangeCount int
	outOfRange   []int
}

func HybridConstructor() HybridMedianFinder {
	return HybridMedianFinder{}
}

func (mf *HybridMedianFinder) AddNum(num int) {
	if num >= 0 && num <= 100 {
		mf.count[num]++
		mf.inRangeCount++
	} else {
		mf.outOfRange = append(mf.outOfRange, num)
	}
}

func (mf *HybridMedianFinder) FindMedian() float64 {
	total := mf.inRangeCount + len(mf.outOfRange)
	// mid1, mid2 := -1, -1

	allNums := make([]int, 0, total)
	for i, c := range mf.count {
		for j := 0; j < c; j++ {
			allNums = append(allNums, i)
		}
	}
	allNums = append(allNums, mf.outOfRange...)
	sort.Ints(allNums)

	if total%2 == 1 {
		return float64(allNums[total/2])
	}
	return float64(allNums[total/2-1]+allNums[total/2]) / 2
}

// func main() {
// 	mf := Constructor()
// 	mf.AddNum(10)
// 	mf.AddNum(20)
// 	fmt.Println(mf.FindMedian()) // 15.0
// 	mf.AddNum(30)
// 	fmt.Println(mf.FindMedian()) // 20.0
// }
