package trie

import (
	"testing"
)

func TestTrieInsertAndSearch(t *testing.T) {
	for _, s := range []string{
		"cat",
		"dog",
		"coffee",
	} {
		trie := New()
		trie.Insert(s)
		v := trie.Search(s)
		if v == nil {
			t.Errorf("%s not exist", s)
		}

		if !v.IsEntry {
			t.Errorf("%s found node should have IsEntry=true, but got false", s)
		}
	}
}

func TestTrieInsertMultipleValues(t *testing.T) {
	// given
	cat := "cat"
	cattle := "cattle"

	trie := New()
	trie.Insert(cattle)

	v := trie.Search(cat)
	if v != nil {
		t.Errorf("%s should be nil", cat)
	}

	if v = trie.Search(cattle); v == nil {
		t.Errorf("%s not found in trie", cattle)
	}

	trie.Insert(cat)

	if v := trie.Search(cat); v == nil || !v.IsEntry {
		t.Errorf("%s has inserted, v should not be nil and IsEntry should be true", cat)
	}
}

func TestTrieDeleteNonEntry(t *testing.T) {
	trie := New()
	trie.Insert("cat")
	trie.Delete("ca")
	if trie.Search("cat") == nil {
		t.Error("cat should be alive")
	}
}

func TestTrieDelete(t *testing.T) {
	trie := New()
	for _, s := range []string{
		"cattle",
		"cat",
		"milk",
		"cattridge",
	} {
		trie.Insert(s)
		if trie.Search(s) == nil {
			t.Errorf("%s expected, but got nil", s)
		}

		// delete
		trie.Delete(s)
		if trie.Search(s) != nil {
			t.Errorf("%s deleted, but got non-nil value", s)
		}
	}
}

func TestLetterToIdx(t *testing.T) {
	for _, tt := range []struct {
		char     byte
		expected int
	}{
		{char: 'a', expected: 0},
		{char: 'b', expected: 1},
		{char: 'z', expected: 25},
	} {
		if val := letterToIdx(tt.char); val != tt.expected {
			t.Errorf("%d expected, got: %d", tt.expected, val)
		}
	}
}
