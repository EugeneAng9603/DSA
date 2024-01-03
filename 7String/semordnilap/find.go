package main

// words = ["diaper", "repaid", "cba", "abc"]
// output = [["diaper", "repaid"], ["abc", "cba"]]
// have a map of word:true
// later check if reversed is found in map
// O(n*m), O(n*m)

func Semordnilap(words []string) [][]string {
	// Write your code here.
	wordSet := map[string]bool{}
	for _, word := range words {
		wordSet[word] = true
	}

	pairs := [][]string{}
	for _, word := range words {
		reversed := reverse(word)
		if _, found := wordSet[reversed]; found && reversed != word {
			pairs = append(pairs, []string{word, reversed})
			delete(wordSet, reversed)
			delete(wordSet, word)
		}
	}

	return pairs
}

func reverse(word string) string {
	reversed := []byte{}
	for i := len(word) - 1; i >= 0; i-- {
		reversed = append(reversed, word[i])
	}
	return string(reversed)
}
