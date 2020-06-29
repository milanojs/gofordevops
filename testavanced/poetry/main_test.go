package main

import (
	"testing"
)

func TestNumStanzas(t *testing.T) {
	p := Poem{}
	if p.NumStanzas() != 0 {
		t.Fatalf("empty Poem is not Empty")
	}
	p = Poem{{"one",
		"two",
		"three"}}

	//assert.Equal(t, 1, p.NumStanzas(), "there should be 3 stanzas")
	if p.NumStanzas() != 1 {
		t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
	}
}

func TestNumLines(t *testing.T) {
	p := Poem{}
	if p.NumLines() != 0 {
		t.Fatalf("empty Poem is not Empty")
	}
	p = Poem{{"one",
		"two",
		"three"}}

	//assert.Equal(t, 1, p.NumLines(), "there should be 3 stanzas")
	if p.NumLines() != 3 {
		t.Fatalf("Unexpected stanza count %d", p.NumLines())
	}
}

func TestNumNewPoem(t *testing.T) {
	if NewPoem() == nil {
		t.Fatalf("the struct is not empty")
	}
}

func TestStats(t *testing.T) {
	//empty string without slices
	p := Poem{}
	v, c, punt := p.Stats()
	if v != 0 || c != 0 || punt != 0 {
		t.Fatalf("There should be not vowels or consonats")
	}
	//empty string with slices
	p = Poem{{}}
	v, c, punt = p.Stats()
	if v != 0 || c != 0 || punt != 0 {
		t.Fatalf("There should be not vowels or consonats")
	}
	//empty string with slices and one word 2 vowels 1 consonant
	p = Poem{{"Hello"}}
	v, c, punt = p.Stats()
	if v != 2 || c != 3 || punt != 0 {
		t.Fatalf("There should be 2 vowels got %d and 3 consonats got %d ", v, c)
	}
	//more complex test of poems
	p = Poem{{"Hello, world!"}}
	v, c, punt = p.Stats()
	if v != 3 || c != 7 || punt != 3 {
		t.Fatalf("There should be 2 vowels got %d and 3 consonats got %d and 3 punts signs %d", v, c, punt)
	}

}

// func TestRunMain(t *testing.T) {
// 	main()
// }
