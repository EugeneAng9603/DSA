// O(nm), O(nm)

// since we are calculating the prefix of strings, we can use a trie
// as we add the strings into trie, we also wanna keep track of fields
// to avoid going through the entire trie again
// we need: maxPrefixCount, maxPrefixLength, currString
// maxPrefixCount for every node for frequency
// maxPrefixLength to keep track curr longest prefix
// currstring is to update the result using maxPrefixLength
// "foods" and maxLength is 3, then curr answer is "foo"

// strings := []string{"algoexpert", "algorithm", "frontendexpert", "mlexpert"}
// expected := "algo"

//	{
//		"strings": ["hello", "world", "fossil", "worldly", "food", "forrest", "helium", "algoexpert", "algorithm"]
//	  }
//	  // expected = "fp"
package main

// O(nm), O(nm)
func LongestMostFrequentPrefix(strings []string) string {
	// Write your code here.
	trie := NewTrie()
	for _, str := range strings {
		trie.Insert(str)
	}
	return trie.maxPrefixFullString[:trie.maxPrefixLength]
}

// count is a field bind to the node to keep track of curr count
type TrieNode struct {
	children map[rune]*TrieNode
	count    int
}

type Trie struct {
	root                *TrieNode
	maxPrefixCount      int
	maxPrefixLength     int
	maxPrefixFullString string
}

// initialise Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}, maxPrefixCount: 0, maxPrefixLength: 0, maxPrefixFullString: ""}
}

// insert each string into trie with count
func (t *Trie) Insert(str string) {
	currNode := t.root
	for i, char := range str {
		// if not found then add new node
		if _, ok := currNode.children[char]; !ok {
			currNode.children[char] = &TrieNode{children: make(map[rune]*TrieNode), count: 0}
		}
		currNode = currNode.children[char]
		currNode.count += 1

		// update if found larger count
		if currNode.count > t.maxPrefixCount {
			t.maxPrefixCount = currNode.count
			t.maxPrefixLength = i + 1
			t.maxPrefixFullString = str
		} else if currNode.count == t.maxPrefixCount && i+1 > t.maxPrefixLength {
			t.maxPrefixLength = i + 1
			t.maxPrefixFullString = str
		}
	}
}

func (t *Trie) Display() {
    var dfs func(node *TrieNode, prefix string)
    dfs = func(node *TrieNode, prefix string) {
        if len(node.children) == 0 {
            fmt.Printf("%s (count: %d)\n", prefix, node.count)
            return
        }
        fmt.Printf("%s (count: %d)\n", prefix, node.count)
        for char, child := range node.children {
            dfs(child, prefix+string(char))
        }
    }
    dfs(t.root, "")
}

func main() {
  trie := NewTrie()
  trie.Insert("algoexpert")
  trie.Insert("algoexabc")
  trie.Display()
}
