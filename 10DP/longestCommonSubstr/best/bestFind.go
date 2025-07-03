// 🔥 Advanced Optimal Solution: Suffix Automaton
// To get better than O(m × n) in practice, we can use a Suffix Automaton:

// Build automaton for one string (O(n) time).

// Walk through the other string to find the longest common substring in O(n) time.

// ✅ Time & Space Complexity
// Time: O(n + m) where n = len(s1), m = len(s2)

// Space: O(n) for the automaton

package best

type state struct {
	next map[byte]int
	len  int
	link int
}

type SuffixAutomaton struct {
	states []state
	last   int
}

func NewSuffixAutomaton(s string) *SuffixAutomaton {
	sa := &SuffixAutomaton{
		states: []state{{next: make(map[byte]int), len: 0, link: -1}}, // length 1
		last:   0,
	}
	for i := 0; i < len(s); i++ {
		sa.extend(s[i])
	}
	return sa
}

func (sa *SuffixAutomaton) extend(c byte) {
	cur := len(sa.states)
	sa.states = append(sa.states, state{
		next: make(map[byte]int),
		len:  sa.states[sa.last].len + 1,
	})
	p := sa.last
	for p != -1 && sa.states[p].next[c] == 0 {
		sa.states[p].next[c] = cur
		p = sa.states[p].link
	}
	if p == -1 {
		sa.states[cur].link = 0
	} else {
		q := sa.states[p].next[c]
		if sa.states[p].len+1 == sa.states[q].len {
			sa.states[cur].link = q
		} else {
			clone := len(sa.states)
			sa.states = append(sa.states, state{
				next: make(map[byte]int),
				len:  sa.states[p].len + 1,
				link: sa.states[q].link,
			})
			for k, v := range sa.states[q].next {
				sa.states[clone].next[k] = v
			}
			for p != -1 && sa.states[p].next[c] == q {
				sa.states[p].next[c] = clone
				p = sa.states[p].link
			}
			sa.states[q].link = clone
			sa.states[cur].link = clone
		}
	}
	sa.last = cur
}

func OptimisedLongestCommonSubstring(s1, s2 string) int {
	sa := NewSuffixAutomaton(s1)
	v := 0
	l := 0
	best := 0
	for i := 0; i < len(s2); i++ {
		c := s2[i]
		for v != -1 && sa.states[v].next[c] == 0 {
			v = sa.states[v].link
			if v != -1 {
				l = sa.states[v].len
			} else {
				l = 0
			}
		}
		if v == -1 {
			v = 0
			l = 0
		} else {
			v = sa.states[v].next[c]
			l++
			if l > best {
				best = l
			}
		}
	}
	return best
}
