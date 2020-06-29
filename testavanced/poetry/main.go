package main

import "fmt"

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
	return Poem{}
}

func (p Poem) NumStanzas() int {
	return len(p)
}

func (s Stanza) NumLines() int {
	return len(s)
}
func (p Poem) NumLines() (count int) {
	for _, s := range p {
		count += s.NumLines()
	}
	return
}

func (p Poem) Stats() (numVowels, numConsonants, numPuntcs int) {
	for _, s := range p {
		for _, l := range s {
			for _, r := range l {
				switch r {
				case 'a', 'e', 'i', 'o', 'u':
					numVowels += 1
				case ',', ' ', '!':
					numPuntcs += 1
				default:
					numConsonants += 1
				}
			}
		}
	}
	return
}
func main() {
	p := Poem{{"one", "two"}}
	v, c, punt := p.Stats()
	fmt.Printf("the num of vowels: %d - Num of consonants %d, - Num of sign %d", v, c, punt)
	fmt.Printf("Stanzas: %d, Lines %d", p.NumStanzas(), p.NumLines())
}
