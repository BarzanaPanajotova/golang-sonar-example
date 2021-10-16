package animal

type dog struct {
	words string
}

func NewDog(words string) dog {
	return dog{words}
}

func (a dog) Speak() string {
	return a.words
}
