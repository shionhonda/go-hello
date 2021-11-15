package intset

import (
	"bytes"
	"fmt"
)

// IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

const intSize = 32 << (^uint(0) >> 63)

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/intSize, uint(x%intSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/intSize, uint(x%intSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		for word != 0 {
			count++
			word = word & (word - 1)
		}
	}
	return count
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/intSize, uint(x%intSize)
	if word < len(s.words) {
		s.words[word] &= ^(1 << bit)
	}
	if len(s.words) > 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var newSet IntSet
	newSet.words = make([]uint, len(s.words))
	copy(newSet.words, s.words)
	return &newSet
}

// AddAll adds all elements of xs to the set
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith sets s to the intersection of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		}
	}
}

// SymmetricDifference sets s to the intersection of s and t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems returns a slice containing the elements of the set
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, intSize*i+j)
			}
		}
	}
	return elems
}
