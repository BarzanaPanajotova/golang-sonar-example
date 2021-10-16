package animal

import "fmt"

type cat struct {
	words string
}

func NewCat(words string) cat {
	return cat{words}
}

func (c cat) Speak() string {
	return c.words
}

func (c cat) Purr() {
	fmt.Println("Purr")
}
