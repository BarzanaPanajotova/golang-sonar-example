package animal

import "fmt"

type dog struct {
	words string
}

func NewDog(words string) dog {
	return dog{words}
}

func (d dog) Speak() string {
	return d.words
}

func (d dog) PlayFetch() {
	fmt.Println("Lets play fetch!")
}
